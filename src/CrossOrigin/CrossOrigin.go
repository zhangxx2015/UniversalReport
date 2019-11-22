package Utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "PUT,GET,POST,DELETE,OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Server", "goHttpd/1.0")
	w.Header().Add("X-Powered-By", "zfx4go<20437023@qq.com>")
	w.Header().Add("Content-Type", contentType)

	w.Header().Add("access-control-expose-headers", "Authorization")


	access-control-allow-credentials: true
	access-control-allow-origin: http://ilab.ponytest.com
	access-control-expose-headers: Access-Control-Allow-Methods
	cache-control: no-store



	access-control-allow-credentials: true
	access-control-allow-origin: http://ilab.ponytest.com
	access-control-expose-headers: Access-Control-Allow-Methods


	Access-Control-Allow-Headers: Origin, X-Requested-With, Content-Type, Accept
*/

func Options(w http.ResponseWriter) {
	// EnableAllCorsAttribute

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Methods", "PUT,GET,POST,DELETE,OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Add("Content-Length", "0")
	w.WriteHeader(http.StatusOK)
}

func AppJson(w http.ResponseWriter, obj interface{}) {
	if nil != obj {
		json, err := json.Marshal(obj)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// EnableAllCorsAttribute
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "PUT,GET,POST,DELETE,OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "*")

		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Cache-Control", "no-cache")
		w.Header().Add("Server", "goHttpd/1.0")
		w.Header().Add("X-Powered-By", "zfx4go<20437023@qq.com>")
		w.Header().Add("Content-Type", "application/json;charset=UTF-8")

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}

func JsonResponse(response interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func AppText(w http.ResponseWriter, text string, contentType string) {
	// EnableAllCorsAttribute
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "PUT,GET,POST,DELETE,OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Server", "goHttpd/1.0")
	w.Header().Add("X-Powered-By", "zfx4go<20437023@qq.com>")
	w.Header().Add("Content-Type", contentType)

	fmt.Fprintf(w, text)
	return

	w.WriteHeader(200)
}
