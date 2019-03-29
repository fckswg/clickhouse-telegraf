package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/kshvakov/clickhouse"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	// get exec path
	ex, err := os.Executable()
	if err != nil {
		panic(fmt.Errorf("Error getting executable path: %v\n", err))
	}
	exPath := filepath.Dir(ex)

	// reading config.toml and setting default values for variables
	conf, err := readConfig("config", exPath, map[string]interface{}{
		"clickhouse.hostname": "127.0.0.1",
		"clickhouse.user":     "default",
		"clickhouse.password": "",
		"cliclhouse.database": "",
	})
	if err != nil {
		panic(fmt.Errorf("Error when reading config: %v\n", err))
	}

	// clickhouse credentials from file
	host := conf.GetString("clickhouse.hostname")
	user := conf.GetString("clickhouse.user")
	password := conf.GetString("clickhouse.password")
	database := conf.GetString("clickhouse.database")

	// prepare connection string
	dataSource := fmt.Sprintf("tcp://%s:9000?username=%s&password=%s&database=system",
		host, user, password)

	metrics := fillMetricsData(dataSource, "SELECT metric, value FROM system.metrics")
	asyncMetrics := fillMetricsData(dataSource, "SELECT * FROM system.asynchronous_metrics")
	events := fillMetricsData(dataSource, "SELECT event, value FROM system.events")
	processes := fillMetricsData(dataSource, "SELECT * FROM system.processes")
	replQueue := fillMetricsData(dataSource, "SELECT count(*) FROM system.replication_queue")

	listTablesQuery := fmt.Sprintf("show tables from %s", database)
	tablesSizes := fillMetricsData(dataSource, listTablesQuery)


	result := make(map[string]interface{})

	result["metrics"] = metrics
	result["asyncMetrics"] = asyncMetrics
	result["events"] = events
	result["processes"] = processes
	if len(replQueue) != 0 {
		result["replicationQueue"] = replQueue
	}
	result["tablesSizes"] = tablesSizes

	final, err := json.Marshal(result)
	fmt.Println(string(final))

}

func readConfig(filename, path string, defaults map[string]interface{}) (*viper.Viper, error) {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}
	v.SetConfigName(filename)
	v.AddConfigPath(path)
	v.AutomaticEnv()
	err := v.ReadInConfig()
	return v, err
}

func fillMetricsData(dataSource, query string) map[string]interface{} {

	data := make(map[string]interface{})

	connect, err := sql.Open("clickhouse", dataSource)
	if err != nil {
		panic(fmt.Errorf("Error while establish connection with clickhouse: %v\n", err))
	}

	count := 0

	if strings.Contains(query, "count") {
		row := connect.QueryRow(query)
		err := row.Scan(&count)
		if err != nil {
			return data
		}
		data["replicationQueue"] = count
	}

	if strings.Contains(query, "show") {
		rows, err := connect.Query(query)
		if err != nil {
			fmt.Println(err)
		}
		var tableList []string
		for rows.Next() {
			var table string
			rows.Scan(&table)
			tableList = append(tableList, table)
		}
		for _, t := range tableList {
			var size int64
			q := fmt.Sprintf("SELECT sum(rows) FROM system.parts WHERE (active = 1) AND (table = '%s')", t)
			row := connect.QueryRow(q)
			err := row.Scan(&size)
			if err != nil {
				return data
			}
			data[t] = size
		}
	}

	rows, err := connect.Query(query)
	if err != nil {
		panic(fmt.Errorf("Error while querying data from clickhouse: %v\n", err))
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	if len(columns) == 2 {
		for rows.Next() {
			var (
				metric string
				value  float64
			)
			if err := rows.Scan(&metric, &value); err != nil {
				log.Fatal(err)
			}
			data[metric] = value
		}
	} else {
		for rows.Next() {
			count += 1
		}
		data["processes"] = count
	}

	return data

}
