package SQL

//package main

import (
	"encoding/json"
	"fmt"
	"strings"

	Instructs "./Instructs"
	Logicals "./Logicals"
	Relationals "./Relationals"
)

// SELECT[DISTINCT]
// FROM
// WHERE
// GROUP BY
// HAVING
// UNION
// ORDER BY

type SQLExpress struct {
	Field      string
	Relational Relationals.Relationals
	Value      string
}

func Exp(field string, relational Relationals.Relationals, value ...string) *SQLExpress {
	return &SQLExpress{Field: field, Relational: relational, Value: fmt.Sprintf(`'%s'`, strings.Join(value, ""))}
}
func (se *SQLExpress) ToString() string {
	return fmt.Sprint(se.Field, se.Relational.ToString(), se.Value)
}

type SQLInfo struct { // zhangxx @ 2019-08-12
	Inc    Instructs.Instructs
	Table  string
	Fields string
	Skip   int
	Size   int
	Orders []string
	Where  []string //[]SQLExpress
}

func New() *SQLInfo {
	return &SQLInfo{Table: "", Fields: "*", Size: -1, Skip: -1}
}

func (si *SQLInfo) ToString() string {
	//jsonData, _ := json.Marshal(si)
	//fmt.Println(string(jsonData))

	sql := ``
	switch si.Inc {
	case Instructs.SELECT:
		sql = fmt.Sprintf(`SELECT %s FROM %s`, si.Fields, si.Table)
		break
	case Instructs.COUNT:
		sql = fmt.Sprintf(`SELECT COUNT(1) AS COUNT FROM %s`, si.Table)
		break
	case Instructs.DELETE:
		sql = fmt.Sprintf(`DELETE FROM %s`, si.Table)
		break
	case Instructs.INSERT:
		sql = fmt.Sprintf(`INSERT INTO %s %s`, si.Table, si.Fields)
		break
	case Instructs.UPDATE:
		sql = fmt.Sprintf(`UPDATE %s SET %s`, si.Table, si.Fields)
		break
	}

	if len(si.Where) > 0 {
		sql = fmt.Sprintf("%s WHERE %s", sql, strings.Join(si.Where, " "))
	}
	if len(si.Orders) > 0 {
		sql = fmt.Sprintf("%s ORDER BY %s", sql, strings.Join(si.Orders, " "))
	}
	if si.Size > -1 {
		if si.Skip > -1 {
			sql = fmt.Sprintf("%s LIMIT %d, %d", sql, si.Skip, si.Size)
		} else {
			sql = fmt.Sprintf("%s LIMIT %d", sql, si.Size)
		}
	}
	return sql
}

func (si *SQLInfo) FROM(table string) *SQLInfo {
	si.Table = table
	return si
}

func (si *SQLInfo) COUNT() *SQLInfo {
	si.Inc = Instructs.COUNT
	si.Fields = "" // TODO:应该放到生成器函数ToString中
	si.Size = -1
	si.Skip = -1
	return si
}

func (si *SQLInfo) SELECT(fields ...string) *SQLInfo {
	si.Inc = Instructs.SELECT
	si.Fields = "*"
	if len(fields) > 0 {
		si.Fields = strings.Join(fields, ",")
	}
	return si
}
func (si *SQLInfo) DELETE() *SQLInfo {
	si.Inc = Instructs.DELETE
	si.Fields = ""
	return si
}

func (si *SQLInfo) INSERT(obj map[string]interface{}) *SQLInfo {
	si.Inc = Instructs.INSERT
	fields := []string{}
	vals := []string{}
	for key, val := range obj {
		fields = append(fields, key)
		jsonBytes, _ := json.Marshal(val)
		vals = append(vals, string(jsonBytes))
	}
	si.Fields = fmt.Sprintf(`(%s) VALUES(%s)`, strings.Join(fields, ","), strings.Join(vals, ","))
	return si
}
func (si *SQLInfo) UPDATE(obj map[string]interface{}) *SQLInfo {
	si.Inc = Instructs.UPDATE
	exps := []string{}
	for key, val := range obj {
		jsonBytes, _ := json.Marshal(val)
		val := string(jsonBytes)
		exps = append(exps, fmt.Sprintf(`%s = %s`, key, val))
	}
	si.Fields = strings.Join(exps, ",")
	return si
}

func (si *SQLInfo) LIMIT(size int) *SQLInfo {
	si.Size = size
	return si
}
func (si *SQLInfo) PAGING(skip int, size int) *SQLInfo {
	si.Size = size
	si.Skip = skip
	return si
}

func (si *SQLInfo) WHERE(logical Logicals.Logicals, express *SQLExpress) *SQLInfo {
	if len(si.Where) > 0 {
		if logical == Logicals.AND {
			si.Where = append(si.Where, "AND ")
		} else if logical == Logicals.OR {
			si.Where = append(si.Where, "OR ")
		}
	}
	si.Where = append(si.Where, express.ToString())
	return si
}

func (si *SQLInfo) ORDER(field string, desc bool) *SQLInfo {
	order := field
	if desc {
		order = fmt.Sprintf("%s DESC", order)
	}
	si.Orders = append(si.Orders, order)
	return si
}

func main() {
	page := 0
	size := 5
	keyword := "zh"

	cmd := New().FROM(`users`).PAGING(page*size, size).ORDER("uid", true)
	if len(keyword) > 0 {
		// cmd = cmd.WHERE(AND, fmt.Sprint(`name LIKE '%`, keyword, `%'`)).
		// 	WHERE(OR, fmt.Sprint(`account LIKE '%`, keyword, `%'`))
		cmd = cmd.WHERE(Logicals.AND, Exp(`name`, Relationals.LIKE, `%`, keyword, `%`)).
			WHERE(Logicals.OR, Exp(`account`, Relationals.LIKE, `%`, keyword, `%`))
	}

	fmt.Println(cmd.SELECT("name,account,age,gender,phone,mail,comment").ToString())
	fmt.Println(cmd.COUNT().ToString())
}
