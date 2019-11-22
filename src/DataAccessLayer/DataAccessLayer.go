package DataAccessLayer

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"

	//_ "github.com/identitii/gdbc/postgresql"

	_ "github.com/mattn/go-sqlite3" // linux上缺乏cgo的问题
	//_ "github.com/iamacarpet/go-sqlite3-dynamic"//未通过验证
	//_ "crawshaw.io/sqlite" //未通过验证
	//_ "github.com/bvinc/go-sqlite-lite/sqlite3"

	_ "github.com/prestodb/presto-go-client/presto"

	_ "gopkg.in/goracle.v2" //oracle

	//_ "github.com/avct/prestgo"

	. "../Exception"

	. "../Utils"
	"github.com/jmoiron/sqlx"
)

func strify(rows *sql.Rows) []string {
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]interface{}, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	c := 0
	results := make(map[string]interface{})
	data := []string{}

	for rows.Next() {
		if c > 0 {
			data = append(data, ",")
		}

		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		for i, value := range values {
			switch value.(type) {
			case nil:
				results[columns[i]] = nil

			case []byte:
				s := string(value.([]byte))
				x, err := strconv.Atoi(s)

				if err != nil {
					results[columns[i]] = s
				} else {
					results[columns[i]] = x
				}

			default:
				results[columns[i]] = value
			}
		}

		b, _ := json.Marshal(results)
		data = append(data, strings.TrimSpace(string(b)))
		c++
	}

	return data
}

func jsonify(rows *sql.Rows, rowIndex int, rowSize int) ([]map[string]interface{}, int) { //[]string {
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]interface{}, len(columns))

	scanArgs := make([]interface{}, len(values))
	//for i := range values {
	for i := 0; i < len(values); i++ {
		scanArgs[i] = &values[i]
	}

	c := 0

	data := make([]map[string]interface{}, 0)

	for rows.Next() {
		if c >= rowIndex && rowSize > len(data) {
			results := make(map[string]interface{})
			// if c > 0 {
			// 	data = append(data, ",")
			// }

			err = rows.Scan(scanArgs...)
			if err != nil {
				panic(err.Error())
			}
			for i, value := range values {
				//fmt.Println("scan columns:", i, columns[i])

				switch value.(type) {
				case nil:
					results[columns[i]] = nil

				case []byte:
					s := string(value.([]byte))
					x, err := strconv.Atoi(s)

					if err != nil {
						results[columns[i]] = s
					} else {
						results[columns[i]] = x
					}

				default:
					results[columns[i]] = value
				}
			}

			// b, _ := json.Marshal(results)
			// data = append(data, strings.TrimSpace(string(b)))
			data = append(data, results)
		}
		c++

		// 内存限制
		memStat := new(runtime.MemStats)
		runtime.ReadMemStats(memStat)
		//fmt.Println(memStat.Alloc)
		if memStat.Alloc > 1024*1024*1024*4 {
			Throw(fmt.Sprintf("out of memory!total rows:[%d],memory:[%d]", c, memStat.Alloc))
			break
		}
		// //系统占用,仅linux/mac下有效
		// //system memory usage
		// sysInfo := new(syscall.Sysinfo_t)
		// err := syscall.Sysinfo(sysInfo)
		// if err == nil {
		// 	mem.All = sysInfo.Totalram * uint32(syscall.Getpagesize())
		// 	mem.Free = sysInfo.Freeram * uint32(syscall.Getpagesize())
		// 	mem.Used = mem.All - mem.Free
		// }

	}

	return data, c
}

func Conn(driver string, user string, pass string, host string, db string) (dbx *sqlx.DB, sql string) {
	dsn := ""
	debug := false
	switch driver {
	case "goracle":
		dsn = fmt.Sprintf("%s/%s@%s/%s", user, pass, host, db)
		if debug {
			dsn = "ilab/ilab@10.10.0.70:1521/oradb"
			sql = "select sysdate from dual"
		}
		break
	case "sqlite3":
		dsn = fmt.Sprintf("%s", db)
		if debug {
			dsn = "mydb.db3"
			sql = "SELECT name,account,job,age,phone FROM users"
		}
		break
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user, pass, host, db)
		if debug {
			dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "lims", "EegaSgHp9g54LnDJ", "192.168.25.11:3306", "lims")
			sql = "select eas_org,area,id,inner_code,name from org_site"
		}
		break
	case "presto":
		dsn = fmt.Sprintf("http://%s:%s@%s?catalog=ilab&schema=%s",
			user,
			pass,
			host,
			db)
		if debug {
			dsn = fmt.Sprintf("http://%s:%s@%s?catalog=ilab&schema=%s", "user", "test", "192.168.25.99:10010", "lims")
			sql = "SHOW TABLES"
			sql = "select name from ilab.lims.org_site"
			sql = "select billing_street from crm.turbocrm.tc_account WHERE crm.turbocrm.tc_account.billing_street IS NOT NULL limit 10"
		}
		break
	case "mssql":
		dsn = fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30&encrypt=disable",
			user,
			pass,
			host,
			db)
		if debug {
			dsn = fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30&encrypt=disable", "csj", "csj82618116", "192.168.25.15:1433", "LittleSystem_DEV")
			sql = "SELECT top 10 name,easDeptName,easOrgName FROM ilabUsers"
		}
		break
	default:
		panic(fmt.Sprintln("unknow driver:%s", driver))
		break
	}
	dbx, err := sqlx.Connect(driver, dsn)
	if err != nil {
		panic(err.Error())
	}
	return dbx, sql
}

func ExecBatch(dbx *sqlx.DB, sqls []string) bool {
	defer func() {
		if DEBUG {
			fmt.Println("Disconnect...")
		}
		dbx.Close()
	}()
	tx := dbx.MustBegin()
	//tx.MustExec(`INSERT INTO student VALUES ('1', 'Jack', 'Jack', 'England', '', '', 'http://img2.imgtn.bdimg.com/it/u=3588772980,2454248748&fm=27&gp=0.jpg', '1', '2018-06-26 17:08:35');`)
	//tx.MustExec(`INSERT INTO student VALUES ('2', 'Emily', 'Emily', 'England', '', '', 'http://img2.imgtn.bdimg.com/it/u=3588772980,2454248748&fm=27&gp=0.jpg', '2', null);`)
	for i := 0; i < len(sqls); i++ {
		ret := tx.MustExec(sqls[i])
		if DEBUG {
			fmt.Println(ret.RowsAffected())
			fmt.Println(ret.LastInsertId())
		}
	}
	err := tx.Commit()
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}
	return true
}

func ExecBatchDetail(driver string, user string, pass string, host string, db string, sqls []string) bool {
	dbx, _ := Conn(driver, user, pass, host, db)
	defer func() {
		if DEBUG {
			fmt.Println("Disconnect...")
		}
		dbx.Close()
	}()
	return ExecBatch(dbx, sqls)
}

func Exec(dbx *sqlx.DB, sql string, args ...interface{}) bool {
	defer func() {
		if DEBUG {
			fmt.Println("Disconnect...")
		}
		dbx.Close()
	}()
	tx := dbx.MustBegin()
	//tx.MustExec(`INSERT INTO student VALUES ('1', 'Jack', 'Jack', 'England', '', '', 'http://img2.imgtn.bdimg.com/it/u=3588772980,2454248748&fm=27&gp=0.jpg', '1', '2018-06-26 17:08:35');`)
	//tx.MustExec(`INSERT INTO student VALUES ('2', 'Emily', 'Emily', 'England', '', '', 'http://img2.imgtn.bdimg.com/it/u=3588772980,2454248748&fm=27&gp=0.jpg', '2', null);`)

	ret := tx.MustExec(sql, args...)
	if DEBUG {
		fmt.Println(ret.RowsAffected())
		fmt.Println(ret.LastInsertId())
	}

	err := tx.Commit()
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}
	return true
}

func ExecDetail(driver string, user string, pass string, host string, db string, sql string, args ...interface{}) bool {
	dbx, _ := Conn(driver, user, pass, host, db)
	defer func() {
		if DEBUG {
			fmt.Println("Disconnect...")
		}
		dbx.Close()
	}()
	return Exec(dbx, sql, args...)
}

var DEBUG bool

func Query(dbx *sqlx.DB, sql string, rowIndex int, rowSize int, args ...interface{}) ([]map[string]interface{}, int) { //[]string {
	defer func() {
		if DEBUG {
			fmt.Println("Disconnect...")
		}
		dbx.Close()
	}()
	if DEBUG {
		fmt.Println(sql)
	}
	rows, err := dbx.Query(sql, args...)
	if nil != err {
		Throw(fmt.Sprintln("ERROR!", err))
	}
	if DEBUG {
		fmt.Printf("datasource:%s\r\n", dbx.DriverName())
		fmt.Println(fmt.Sprintf("SQL:[%s]", sql), args)
		fmt.Print("RESULT:")
		fmt.Println(json.Marshal(rows))
	}

	if nil == rows {
		return make([]map[string]interface{}, 0), 0
	}
	//return make([]map[string]interface{}, 0)
	return jsonify(rows, rowIndex, rowSize)
	/*
		jsonData := jsonify(rows)
		// if DEBUG {
		// 	fmt.Println(jsonData)
		// }
		// defer func() {
		// 	if rows != nil {
		// 		rows.Close() //可以关闭掉未scan连接一直占用
		// 	}
		// }()
		// if DEBUG {
		// 	fmt.Println("done")
		// 	fmt.Scanln()
		// }
		// //return strings.Join(jsonData, "")
		return jsonData
	*/
}

func QueryDetail(driver string, user string, pass string, host string, db string, sql string, rowIndex int, rowSize int, args ...interface{}) ([]map[string]interface{}, int) { //[]string {
	dbx, _ := Conn(driver, user, pass, host, db)
	defer func() {
		if DEBUG {
			fmt.Println("Disconnect...")
		}
		dbx.Close()
	}()
	return Query(dbx, sql, rowIndex, rowSize, args...)
}

func IntResult(key string, dbx *sqlx.DB, sql string, args ...interface{}) int {
	return ParseInt(StringResult(key, dbx, sql, args...))
}
func StringResult(key string, dbx *sqlx.DB, sql string, args ...interface{}) string {
	return ToString(TopRow(dbx, sql, args...)[key])
}
func TopRow(dbx *sqlx.DB, sql string, args ...interface{}) map[string]interface{} { //[]string {
	defer func() {
		if DEBUG {
			fmt.Println("Disconnect...")
		}
		dbx.Close()
	}()
	datas, _ := Query(dbx, sql, 0, 1, args...)
	for _, val := range datas {
		return val
	}
	return make(map[string]interface{})
}
func TopRowDetail(driver string, user string, pass string, host string, db string, sql string, rowIndex int, rowSize int, args ...interface{}) map[string]interface{} { //[]string {
	datas, _ := QueryDetail(driver, user, pass, host, db, sql, rowIndex, rowSize, args...)
	for _, val := range datas {
		return val
	}
	return make(map[string]interface{})
}
