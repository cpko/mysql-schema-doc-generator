package main

import (
	"database/sql"
	"text/template"
	"os"
	"sort"
	"flag"
	"github.com/go-sql-driver/mysql"
	"strconv"
)

type DBInfo struct {
	host     string
	port     int
	username string
	password string
	database string
}

type Column struct {
	Name                   string
	DataType               string
	CharacterMaximumLength sql.NullInt64
	IsNullable             string
	IsIndexed              string
	Comment                string
}

type Table struct {
	TableName string
	Columns   []Column
}

func main() {
	dbInfo := parseDBInfoFromCommandLine();
	tables := fetchTableSchema(dbInfo)
	generate(tables)
}

func parseDBInfoFromCommandLine() DBInfo {
	host := flag.String("host", "127.0.0.1", "host")
	port := flag.Int("port", 3306, "port")
	username := flag.String("u", "root", "username")
	password := flag.String("p", "", "password")
	database := flag.String("d", "", "database")
	flag.Parse()

	if *database == "" {
		panic("database not provided,please use -d option to specify")
	}
	return DBInfo{*host, *port, *username, *password, *database}
}

func fetchTableSchema(dbInfo DBInfo) []Table {
	conf := mysql.Config{
		User:dbInfo.username,
		Passwd:dbInfo.password,
		Net:"tcp",
		Addr:dbInfo.host + ":" + strconv.Itoa(dbInfo.port),
		DBName:"information_schema",
		Params:map[string]string{"charset":"utf8"},
	};
	db, err := sql.Open("mysql", conf.FormatDSN())
	checkError(err)
	defer db.Close()

	rows, err := db.Query("select table_name,column_name,data_type,character_maximum_length,is_nullable,column_comment,column_key from columns where table_schema=?", dbInfo.database)
	checkError(err)

	tableColumns := make(map[string][]Column)
	for rows.Next() {
		var tableName string
		var columnName string
		var dataType string
		var characterMaximumLength sql.NullInt64
		var isNullable string
		var columnComment string
		var columnKey string
		err = rows.Scan(&tableName, &columnName, &dataType, &characterMaximumLength, &isNullable, &columnComment, &columnKey)
		checkError(err)

		var c Column
		c.Name = columnName
		c.DataType = dataType
		c.CharacterMaximumLength = characterMaximumLength
		c.IsNullable = isNullable
		if columnKey != "" {
			c.IsIndexed = "YES"
		} else {
			c.IsIndexed = "NO"
		}
		if columnComment != "" {
			c.Comment = columnComment;
		} else {
			if columnKey == "PRI" {
				c.Comment = "主键"
			}
		}

		columns := []Column{}
		existedColumns, ok := tableColumns[tableName]
		if ok {
			columns = append(existedColumns, c)
		} else {
			columns = append(columns, c)
		}
		tableColumns[tableName] = columns

	}

	//sort by table_name
	tableNames := make([]string, len(tableColumns))
	i := 0
	for tableName := range tableColumns {
		tableNames[i] = tableName
		i++
	}
	sort.Strings(tableNames)

	tables := make([]Table, len(tableNames))
	j := 0
	for _, tableName := range tableNames {
		table := Table{tableName, tableColumns[tableName]}
		tables[j] = table
		j++
	}

	return tables
}

func generate(tables []Table) {
	t, err := template.ParseFiles("./tables_desc_template.gtpl")
	checkError(err)

	file, err := os.Create("schema_doc.md")
	checkError(err)
	defer file.Close()

	execErr := t.Execute(file, tables)
	checkError(execErr)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}