package Http

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/axgle/mahonia"
)

func Post(url string, contentType string, body string) string {
	req, _ := http.NewRequest("POST", url, strings.NewReader(body))
	req.Header.Add("content-type", contentType)
	req.Header.Add("cache-control", "no-cache")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	respBody, _ := ioutil.ReadAll(res.Body)
	return string(respBody)
}

func Get(url string) string {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("Accept-Encoding", "")
	resp, _ := client.Do(request)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	var reader io.ReadCloser
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			panic(err.Error())
		}
	} else {
		reader = resp.Body
	}
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err.Error())
	}
	contentType := resp.Header.Get("Content-Type")
	if strings.Contains(contentType, "gbk") {
		return string(ConvertToByte(string(body), "gbk", "utf8"))
	}
	return string(body)
}
func ConvertToByte(src string, srcCode string, targetCode string) []byte {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(targetCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	return cdata
}

func Submit(uri string, params map[string]string, paramName, path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return "", err
	}
	_, err = io.Copy(part, file)
	err = writer.Close()
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("Accept", "application/json")

	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		return string(body), nil
	}
	return "", err
}
