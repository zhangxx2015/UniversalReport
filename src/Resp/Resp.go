package Contract

//////////////////////////////////// API 契约包装
type Resp struct { // zhangxx @ 2019-06-12 兼容
	// 记录总数
	Total int `json:"totalCount"` //`json:"total"`
	// 分页上限
	PageMax int `json:"totalPage"` //`json:"PageMax"`
	// 分页索引
	PageIndex int `json:"page"` //`json:"pageIndex"`
	// 分页大小
	PageSize int `json:"size"` //`json:"pageSize"`
	//// 请求内容
	//Request string `json:"Request"`
	// 返回状态
	Code   string `json:"code"`
	Status string `json:"Status"`
	// 返回内容
	ResType string `json:"ResType"`
	// 数据
	Data interface{} `json:"data"` //`json:"Data"`
	// 错误
	Error string `json:"message"` //`json:"Error"`
}

//////////////////////////////////// Resp 辅助函数

type Response struct {
	Data string `json:"data"`
}

func Ok() Resp {
	return Resp{Total: 0, PageMax: 0, PageIndex: 0, PageSize: 0, Status: "OK", Code: "0", ResType: "OK", Data: []string{}, Error: ""}
}

func Error(err string) Resp {
	return Resp{Total: 0, PageMax: 0, PageIndex: 0, PageSize: 0, Status: "Error", Code: "1", ResType: "Error", Data: nil, Error: err}
}

func JsonData(obj interface{}) Resp {
	return Resp{Total: 0, PageMax: 0, PageIndex: 0, PageSize: 0, Status: "OK", Code: "0", ResType: "JsonData", Data: obj, Error: ""}
}

func JsonDatas(objs []string, total int, pageIndex int, pageSize int) Resp {
	pageMax := total / pageSize
	if pageSize > total {
		pageMax = 1
	}
	return Resp{Total: total, PageMax: pageMax, PageIndex: pageIndex, PageSize: pageSize, Status: "OK", Code: "0", ResType: "JsonData", Data: objs, Error: ""}
}
func JsonMaps(objs []map[string]interface{}, total int, pageIndex int, pageSize int) Resp {
	pageMax := total / pageSize
	if pageSize > total {
		pageMax = 1
	}
	return Resp{Total: total, PageMax: pageMax, PageIndex: pageIndex, PageSize: pageSize, Status: "OK", Code: "0", ResType: "JsonData", Data: objs, Error: ""}
}
func JsonMapsDebug(objs []map[string]interface{}, total int, pageIndex int, pageSize int, message string) Resp {
	pageMax := total / pageSize
	if pageSize > total {
		pageMax = 1
	}
	return Resp{Total: total, PageMax: pageMax, PageIndex: pageIndex, PageSize: pageSize, Status: "OK", Code: "0", ResType: "JsonData", Data: objs, Error: message}
}

// func resp(w http.ResponseWriter, obj interface{}) {
// 	appJson(w, JsonData(obj))
// }
