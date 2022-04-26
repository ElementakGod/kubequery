package terminal

import (
	"database/sql"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"strings"
)

type Table struct {
	Headers []string
	Data    [][]string
}

func PrintTable(resultTable Table) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(resultTable.Headers)
	table.SetAutoFormatHeaders(false)
	for _, v := range resultTable.Data {
		table.Append(v)
	}
	table.Render()
}

func ShowTables(db *sql.DB) {
	tableNamesQuery := "SELECT name FROM sqlite_master"
	tableNames := GetData(db, tableNamesQuery)
	result := make([][]string, 0)
	resultTable := Table{
		Headers: []string{"table", "columns"},
	}

	for _, table := range tableNames.Data {
		columnNamesQuery := "SELECT name FROM pragma_table_info('" + table[0] + "')"
		columnNames := GetData(db, columnNamesQuery)
		var columns []string

		for _, v1 := range columnNames.Data {
			s := strings.Join(v1, ", ")
			columns = append(columns, s)
		}
		result = append(result, []string{table[0], strings.Join(columns, ", ")})
	}

	resultTable.Data = result
	PrintTable(resultTable)
}

func RunQuery(db *sql.DB, query string) (int64, error) {
	tx, _ := db.Begin()
	statement, err := tx.Prepare(query)
	if err != nil {
		fmt.Printf("Error in %v : %v\n", query, err.Error())
		return 0, nil
	}
	defer statement.Close()
	res, err := statement.Exec()
	if err != nil {
		tx.Rollback()
		fmt.Printf("Error in %v : %v\n", query, err.Error())
		return 0, nil
	}
	defer tx.Commit()
	return res.RowsAffected()
}

func GetData(db *sql.DB, query string) Table {
	table := Table{}
	row, err := db.Query(query)
	if err != nil {
		fmt.Printf("Error in %v : %v\n", query, err.Error())
		return table
	}
	defer row.Close()

	columns, err := row.Columns()
	if err != nil {
		fmt.Printf("Error reading columns %v : %v\n", query, err.Error())
		return table
	}

	output := make([][]string, 0)
	rawResult := make([][]byte, len(columns))
	dest := make([]interface{}, len(columns))
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	for row.Next() {
		row.Scan(dest...)
		res := make([]string, 0)
		for _, raw := range rawResult {
			if raw != nil {
				res = append(res, string(raw))
			}
		}
		output = append(output, res)
	}

	table.Headers = columns
	table.Data = output
	return table
}
