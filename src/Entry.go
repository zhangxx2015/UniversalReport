package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"sort"

	AppSettings "./AppSettings"
	CrossOrigin "./CrossOrigin"
	Dal "./DataAccessLayer"
	Directory "./Directory"
	. "./Exception"
	Http "./Http"
	JSON "./Json"
	Resp "./Resp"
	SQL "./SQL"
	Logicals "./SQL/Logicals"
	Relationals "./SQL/Relationals"
	. "./Utils"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/jmoiron/sqlx"

	"github.com/mahtuag/jwtplay/auth"
)

type Col struct {
	Raw   string
	Title string
	Index int
}
type colSlice []Col

// 这三个方法必须有，相当于实现了sort.Interface
func (s colSlice) Len() int {
	return len(s)
}
func (s colSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s colSlice) Less(i, j int) bool {
	return s[i].Index < s[j].Index // 这里是关键，我比较了Index这个字段
}

////////////////////////////////////////////////////////////////

func main() {
	conn := AppSettings.Config("./conn.json")
	Dal.DEBUG = conn[0].Debug
	if Dal.DEBUG {
		fmt.Println("!!DEBUG MODE")
	}
	dbx := func() *sqlx.DB {
		dbx, _ := Dal.Conn(conn[0].Driver, conn[0].Username, conn[0].Password, conn[0].Server, conn[0].Database)
		return dbx
	}

	port := flag.Int("p", 80, "http listen port")
	flag.Parse()
	fmt.Println("tiny http server, create by zhangxx(20437023)")
	if Dal.Exec(dbx(), conn[0].Test) {
		fmt.Printf("start service is succeed. listen port:[%d]\r\n", *port)
	}

	go func() { // self test
		time.Sleep(time.Second * 1)
		selfUrl := fmt.Sprintf("http://localhost:%d/ping?page=1&size=10", *port)
		fmt.Println("[!self test]>>[", selfUrl, "]")
		fmt.Println("[!self test]<<Response=", Http.Post(selfUrl, "text/plain", fmt.Sprintf(`{"username":"zhangxx","password":"123456"}`)))
	}()
	setts := AppSettings.Config("conf.json")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if "OPTIONS" == r.Method {
			CrossOrigin.Options(w)
			return
		}
		if strings.HasPrefix(r.RequestURI, "/ping") {
			fmt.Println("[PING]")
			fmt.Println("\tquery string:", "page=", QueryInt(r, "page", 0), "size=", QueryInt(r, "size", 10))
			fmt.Println("\tbody payload:", Payload(r))
			CrossOrigin.AppJson(w, Resp.Response{fmt.Sprintf("[%s]pong from golang", time.Now().Format("2006-01-02 15:04:05"))})
			return
		}
		authorized := false
		authorization := func() bool {
			result := false
			Try(func() {
				authorized = true
				//tokenString := "Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICI3SVJUX3N6bkVkc3hIZUl0V3NCZmZ5TTBvVEFvazg3aXl3SFRrbFRsSjM4In0.eyJqdGkiOiJmMzE5YWFiNy1kOTAwLTQ2OWMtYjkyZi1jMjJiYjg3NTIwMzIiLCJleHAiOjE1Njc1NjE3NzEsIm5iZiI6MCwiaWF0IjoxNTY3NTYxNDcxLCJpc3MiOiJodHRwczovL3Nzby5wb255dGVzdC5jb206ODQ0My9hdXRoL3JlYWxtcy9wb255dGVzdCIsImF1ZCI6ImlMYWItd2VieCIsInN1YiI6ImRkMWM0YzdlLTQ0ZmItNGU4OC05OWQ1LWNjZDA1N2M3M2JkMSIsInR5cCI6IkJlYXJlciIsImF6cCI6ImlMYWItd2VieCIsIm5vbmNlIjoiYjdkZDEwZTYtOTM1ZC00ZTNmLWI1YmEtMWZmMjIxZGQwOWNjIiwiYXV0aF90aW1lIjoxNTY3NTYwMjU5LCJzZXNzaW9uX3N0YXRlIjoiMTE0NjlmNDYtZTU3Yi00NTY3LWI4NTUtMDQwODEzNDg5YjRkIiwiYWNyIjoiMCIsImFsbG93ZWQtb3JpZ2lucyI6WyJodHRwOi8vaWxhYi5wb255dGVzdC5jb20iXSwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJyZWFsbS1tYW5hZ2VtZW50Ijp7InJvbGVzIjpbInZpZXctcmVhbG0iLCJ2aWV3LWlkZW50aXR5LXByb3ZpZGVycyIsInZpZXctZXZlbnRzIiwibWFuYWdlLXVzZXJzIiwicXVlcnktcmVhbG1zIiwidmlldy11c2VycyIsInZpZXctY2xpZW50cyIsInZpZXctYXV0aG9yaXphdGlvbiIsInF1ZXJ5LWNsaWVudHMiLCJxdWVyeS1ncm91cHMiLCJxdWVyeS11c2VycyJdfSwiaWxhYi1yZXBvcnQiOnsicm9sZXMiOlsiUk9MRV9VU0VSIl19LCJpTGFiLXNlcnZlciI6eyJyb2xlcyI6WyJST0xFX0FQSSIsIlJPTEVfVVNFUiJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJvcGVuaWQgcHJvZmlsZSBlbWFpbCIsInBvbnlfb3JnX2NvZGUiOiIxMDAxXzAxIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJuYW1lIjoi5bygIOa9h-a9hyIsInByZWZlcnJlZF91c2VybmFtZSI6InpoYW5neGlhb3hpYW8iLCJnaXZlbl9uYW1lIjoi5bygIiwibG9jYWxlIjoiemgtQ04iLCJmYW1pbHlfbmFtZSI6Iua9h-a9hyJ9.CKEIuKf7ovMK6zk5NcQY9QNDhRQrcIPXeAsK9n6Z8eYwFRr3qDCd6VE5UM7bXeYdDcunqP2WgThes9q73zVNUHEVPB5SDytcHS_UcwRw5IjPMm53z9TuYbXjHeHDs3UKp8k7y3jNVSY8Tw56wWwf4ZR9CCBjZfiKrInMOMiW53OTsHHLHQdQP2iVRH2eHvbh_cU0Nhbit_mkvUspSHNBvORp3rEYMjzF-Buq05EH1GDtRmASmW3iryb9NZVhxYk7_RUg-XwlG9wmL88K9EQDvq-0uK5g4KDNRwdVoAqBBzQ-0mbsZBZgOzRaPm03V9nhIu35yxYO0xbjYt4ElkB4ig"
				tokenString := r.Header.Get("Authorization")
				if Dal.DEBUG {
					fmt.Println("tokenString:", tokenString)
					if 1 > len(tokenString) {
						Throw("Unauthorization")
					}
				}

				tokenString = strings.TrimPrefix(tokenString, "Bearer ")
				claims, errClaims := auth.ParseClaims(tokenString)
				if nil != errClaims {
					Throw(fmt.Sprintln("ERROR:", errClaims.Error()))
				}
				// 调试输出
				// bytesJson, errMarshal := json.Marshal(claims)
				// if nil != errMarshal {
				// 	Throw(fmt.Sprintln("ERROR:", errMarshal.Error()))
				// }
				//fmt.Println("TOKEN INFO:", string(bytesJson)[0:20])
				// TODO:验证身份
				//{
				//  "resource_access": {
				//	  "ilab-report": {
				//		"roles": [
				//			"ROLE_USER"
				if val, ok := claims["resource_access"]; ok { //resource_access := claims["resource_access"].(map[string]interface{})
					resource_access := val.(map[string]interface{})
					if Dal.DEBUG {
						fmt.Println("\n resource_access:\n\t", resource_access)
					}
					if val, ok := resource_access["ilab-report-x"]; ok { //ilab_report := resource_access["ilab-report"].(map[string]interface{})
						ilab_report := val.(map[string]interface{})
						if Dal.DEBUG {
							fmt.Println("\n ilab_report:\n\t", ilab_report)
						}
						if val, ok := ilab_report["roles"]; ok { //roles := ilab_report["roles"].([]interface{})
							roles := val.([]interface{})
							if Dal.DEBUG {
								fmt.Println("\n roles:\n\t", roles)
							}
							if len(roles) > 0 {
								if Dal.DEBUG {
									fmt.Println("\n [0]:\n\t", roles[0])
								}
								result = "ROLE_USER" == roles[0]
								if false == result {
									if Dal.DEBUG {
										fmt.Println("\n Missing [ROLE_USER] Role\n\t")
									}
									Throw("Missing Role Authorization")
								}
							}
						}
					}
				}
			}, func(err error) {
				CrossOrigin.AppJson(w, Resp.Error(err.Error()))
			})
			return result
		}

		if strings.HasPrefix(r.RequestURI, "/listDsn") && authorization() {
			dsns := make([]map[string]interface{}, 0)
			for k, v := range setts {
				dsns = append(dsns, map[string]interface{}{"text": v.Driver, "id": k})
			}
			CrossOrigin.AppJson(w, Resp.JsonMaps(dsns, 1, 0, 10))
			return
		}
		if strings.HasPrefix(r.RequestURI, "/listBiz") && authorization() {
			Try(func() {
				cmd := SQL.New().FROM(`bizs`).ORDER("cts", false)
				page := QueryInt(r, "page", 0)
				size := QueryInt(r, "size", 10)
				keyword := PayloadMapString(PayloadMap(r), "keyword", "")
				if len(keyword) > 0 {
					cmd = cmd.WHERE(Logicals.AND, SQL.Exp(`name`, Relationals.LIKE, `%`, keyword, `%`))
				}
				rows := make([]map[string]interface{}, 0)
				Count := Dal.IntResult("COUNT", dbx(), cmd.COUNT().ToString())
				if Count > 0 {
					rows, _ = Dal.Query(dbx(), cmd.SELECT().ToString(), page*size, size)
				}
				CrossOrigin.AppJson(w, Resp.JsonMaps(rows, Count, page, size))
			}, func(err error) {
				CrossOrigin.AppJson(w, Resp.Error(err.Error()))
			})
			return
		}
		if strings.HasPrefix(r.RequestURI, "/readBiz") && authorization() {
			Try(func() {
				if !QueryContainskey(r, "uid") {
					Throw("[uid] cannot be empty")
				}
				uid := QueryString(r, "uid", "-1")
				rows, _ := Dal.Query(dbx(), SQL.New().SELECT().FROM(`bizs`).
					WHERE(Logicals.AND, SQL.Exp(`uid`, Relationals.EQ, uid)).ToString(), 0, 1)
				CrossOrigin.AppJson(w, Resp.JsonMaps(rows, 1, 0, 1))
			}, func(err error) {
				CrossOrigin.AppJson(w, Resp.Error(err.Error()))
			})
			return
		}
		if strings.HasPrefix(r.RequestURI, "/killBiz") && authorization() {
			Try(func() {
				if !QueryContainskey(r, "uid") {
					Throw("[uid] cannot be empty")
				}
				uid := QueryString(r, "uid", "-1")
				if Dal.Exec(dbx(), SQL.New().DELETE().FROM(`bizs`).
					WHERE(Logicals.AND, SQL.Exp(`uid`, Relationals.EQ, uid)).ToString()) {
					CrossOrigin.AppJson(w, Resp.Ok())
					return
				}
				CrossOrigin.AppJson(w, Resp.Error("unknow error"))
			}, func(err error) {
				CrossOrigin.AppJson(w, Resp.Error(err.Error()))
			})
			return
		}
		if strings.HasPrefix(r.RequestURI, "/saveBiz") && authorization() {
			Try(func() {
				body4 := PayloadMapNotContainskeys(r, func(key string, exist bool) {
					Throw(fmt.Sprintf("[%s] cannot be empty", key))
				}, "uid", "dsn", "name", "memo", "args", "code")
				uid := PayloadMapString(body4, "uid", "")
				delete(body4, "cts")
				delete(body4, "uid")
				cmdStr := ""
				if "" == uid {
					cmdStr = SQL.New().INSERT(body4).FROM(`bizs`).ToString()
				} else {
					cmdStr = SQL.New().UPDATE(body4).FROM(`bizs`).
						WHERE(Logicals.AND, SQL.Exp(`uid`, Relationals.EQ, uid)).ToString()
				}

				if Dal.Exec(dbx(), cmdStr) {
					name := PayloadMapString(body4, "name", "")
					rows, _ := Dal.Query(dbx(), SQL.New().SELECT().FROM(`bizs`).
						WHERE(Logicals.AND, SQL.Exp(`name`, Relationals.EQ, name)).ToString(), 0, 1)
					CrossOrigin.AppJson(w, Resp.JsonMaps(rows, 1, 0, 1))
					return
				}
				CrossOrigin.AppJson(w, Resp.Error("unknow error"))

			}, func(err error) {
				CrossOrigin.AppJson(w, Resp.Error(err.Error()))
			})
			return
		}
		if strings.HasPrefix(r.RequestURI, "/execBiz") && authorization() {
			Try(func() {
				body5 := PayloadMapNotContainskeys(r, func(key string, exist bool) {
					Throw(fmt.Sprintf("[%s] cannot be empty", key))
				}, "uid", "query")

				uid := PayloadMapString(body5, "uid", "-1")
				fmt.Println("uid:", uid)
				query := body5["query"].(map[string]interface{})
				fmt.Println("query:", query, fmt.Sprintf("%T", query))

				bizs, countBizs := Dal.Query(dbx(), SQL.New().SELECT().FROM(`bizs`).
					WHERE(Logicals.AND, SQL.Exp(`uid`, Relationals.EQ, uid)).ToString(), 0, 1)
				if 1 != countBizs {
					Throw(fmt.Sprintf("biz [%s] is not found", uid))
				}
				argsEndata := bizs[0]["args"].(string)
				decodeArgs, err := base64.StdEncoding.DecodeString(argsEndata)
				if err != nil {
					Throw(fmt.Sprintf("fail decode of field['args']"))
				}
				args := make([]map[string]interface{}, 0)
				if len(decodeArgs) > 0 {
					args = JSON.ParseArray(string(decodeArgs))
					fmt.Println("args:", args)
				}

				codeEndata, err := base64.StdEncoding.DecodeString(bizs[0]["code"].(string))
				if err != nil {
					Throw(fmt.Sprintf("fail decode of field['code']"))
				}
				if len(codeEndata) < 1 {
					Throw(fmt.Sprintf("['code'] cannot be empty"))
				}
				code := string(codeEndata)
				fmt.Println("code:", code)

				dsn := ParseIntDefault(ToString(bizs[0]["dsn"]), -1)
				fmt.Println("dsn:", dsn)
				biz := bizs[0]["name"].(string)
				fmt.Println("biz:", biz)

				pcode := code
				for i := 0; i < len(args); i++ {
					arg := args[i]
					for k, v := range arg {
						if "name" != k {
							continue
						}
						bingo := false
						for k1, v1 := range query {
							fmt.Println("Query:", k1, v1)
							if v != k1 {
								continue
							}
							toString := JSON.Stringify(v1, false)
							if len(toString) > 0 && strings.HasPrefix(toString, "\"") {
								toString = toString[1:]
							}
							if len(toString) > 0 && strings.HasSuffix(toString, "\"") {
								toString = toString[0 : len(toString)-1]
							}

							pcode = strings.Replace(pcode, fmt.Sprint(`/*--`, v, `--*/`), toString, -1)
							bingo = true
						}
						if bingo {
							continue
						}
						pcode = strings.Replace(pcode, fmt.Sprint(`/*--`, v, `--*/`), "", -1)
					}
				}
				fmt.Println("")
				fmt.Println("")
				fmt.Println("")
				fmt.Println("")
				fmt.Println("")
				fmt.Println("pcode:")
				fmt.Println(pcode)
				fmt.Println("")
				fmt.Println("")
				fmt.Println("")
				fmt.Println("")
				fmt.Println("")

				rows := make([]map[string]interface{}, 0)
				dbx, _ := Dal.Conn(setts[dsn].Driver, setts[dsn].Username, setts[dsn].Password, setts[dsn].Server, setts[dsn].Database)

				page := QueryInt(r, "page", 0)
				size := QueryInt(r, "size", 10)
				//fmt.Println("page:", page, "size:", size)
				export := QueryBool(r, "export", false)
				if export {
					size = 999999
				}

				rows, total := Dal.Query(dbx, pcode, page*size, size)
				if export {
					f := excelize.NewFile() // 创建一个工作表
					index := f.NewSheet("Sheet1")
					if len(rows) > 0 {
						keys := make(colSlice, 0)
						idx := 1
						for k, _ := range rows[0] {
							if -1 != strings.IndexAny(k, "@") {
								parts := strings.Split(k, "@")
								key := Col{Raw: k, Title: parts[0], Index: ParseIntDefault(parts[1], 0)}
								keys = append(keys, key)
							} else {
								key := Col{Raw: k, Title: k, Index: idx}
								keys = append(keys, key)
							}
							idx++
						}
						sort.Sort(keys)
						i := 1
						for j := 0; j < len(keys); j++ {
							f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", ToNumberSystem26(i), 1), keys[j].Title)
							i++
						}

						for j := 0; j < len(rows); j++ {
							row := rows[j]
							for i := 0; i < len(keys); i++ {
								if _, ok := row[keys[i].Raw]; ok {
									v := row[keys[i].Raw]
									axis := fmt.Sprintf("%s%d", ToNumberSystem26(1+i), 2+j)
									f.SetCellValue("Sheet1", axis, v)
								}
							}
						}
						f.SetColWidth("Sheet1", "A", ToNumberSystem26(len(keys)), 20)
					}

					f.SetActiveSheet(index) // 设置工作簿的默认工作表
					// errSaveAs := f.SaveAs("./Book1.xlsx")// 根据指定路径保存文件
					exportName := fmt.Sprintf("%s_%s.xlsx", biz, time.Now().Format("2006-01-02#15_04_05"))
					fmt.Println("exportName:", exportName)
					w.Header().Set("Content-Type", "application/zip")
					w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", exportName))
					_, errSaveAs := f.WriteTo(w)
					if errSaveAs != nil {
						Throw(fmt.Sprintln(errSaveAs))
					}
					return
				}

				if Dal.DEBUG {
					CrossOrigin.AppJson(w, Resp.JsonMapsDebug(rows, total, page, size, pcode))
				} else {
					CrossOrigin.AppJson(w, Resp.JsonMaps(rows, total, page, size))
				}
			}, func(err error) {
				CrossOrigin.AppJson(w, Resp.Error(err.Error()))
			})
			return
		}
		if strings.HasPrefix(r.RequestURI, "/execRaw") && authorization() {
			fmt.Println("next!!!")
			Try(func() {
				// PostData(`http://localhost/execRaw?page=0&size=100`,{"dsn":"0",query:`select * from
				//  xxt_bj.dbo.ilabusers`},function (resp){
				//     debugger
				// });
				body5 := PayloadMapNotContainskeys(r, func(key string, exist bool) {
					Throw(fmt.Sprintf("[%s] cannot be empty", key))
				}, "dsn", "query", "code")

				dsnStr := PayloadMapString(body5, "dsn", "-1")
				dsn := ParseIntDefault(dsnStr, -1)
				fmt.Println("dsn:", dsn)
				if -1 == dsn {
					Throw(fmt.Sprintf("[%d] dsn is illegal", dsn))
				}

				argsEndata := PayloadMapString(body5, "query", "")
				//fmt.Println("argsEndata:", argsEndata)
				decodeArgs, err := base64.StdEncoding.DecodeString(argsEndata)
				//fmt.Println("decodeArgs:", string(decodeArgs))
				if err != nil {
					Throw(fmt.Sprintf("fail decode of field['query']"))
				}
				args := make([]map[string]interface{}, 0)
				if len(decodeArgs) > 0 {
					args = JSON.ParseArray(string(decodeArgs))
					fmt.Println("query:", args)
				}

				code := PayloadMapString(body5, "code", "")
				fmt.Println("code:", code)

				pcode := code
				for i := 0; i < len(args); i++ {
					arg := args[i]
					for k, v := range arg {
						//fmt.Println("k:", k, "v:", v)
						toString := JSON.Stringify(v, false)
						if len(toString) > 0 && strings.HasPrefix(toString, "\"") {
							toString = toString[1:]
						}
						if len(toString) > 0 && strings.HasSuffix(toString, "\"") {
							toString = toString[0 : len(toString)-1]
						}
						pcode = strings.Replace(pcode, fmt.Sprint(`/*--`, k, `--*/`), toString, -1)
					}
				}
				fmt.Println("")
				fmt.Println("")
				fmt.Println("")
				fmt.Println("")
				fmt.Println("")
				fmt.Println("pcode:")
				fmt.Println(pcode)
				fmt.Println("")
				fmt.Println("")
				fmt.Println("")
				fmt.Println("")
				fmt.Println("")

				rows := make([]map[string]interface{}, 0)
				dbx, _ := Dal.Conn(setts[dsn].Driver, setts[dsn].Username, setts[dsn].Password, setts[dsn].Server, setts[dsn].Database)

				page := QueryInt(r, "page", 0)
				size := QueryInt(r, "size", 10)
				//fmt.Println("page:", page, "size:", size)

				rows, total := Dal.Query(dbx, pcode, page*size, size)

				if Dal.DEBUG {
					CrossOrigin.AppJson(w, Resp.JsonMapsDebug(rows, total, page, size, pcode))
				} else {
					CrossOrigin.AppJson(w, Resp.JsonMaps(rows, total, page, size))
				}
			}, func(err error) {
				CrossOrigin.AppJson(w, Resp.Error(err.Error()))
			})
			return
		}
		if !authorized {
			http.FileServer(http.Dir("./")).ServeHTTP(w, r)
		}
	})

	portVal := fmt.Sprintf(":%d", *port)
	var err error
	if withSSL, _ := Directory.PathExists("./ssl/"); !withSSL {
		err = http.ListenAndServe(portVal, nil)
	} else {
		err = http.ListenAndServeTLS(portVal, "./ssl/server.crt", "./ssl/server.key", nil)
	}
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
