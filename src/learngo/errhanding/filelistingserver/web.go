package main

import (
	"log"
	"net/http"
	"os"

	"github.com/whhcs/learngo/errhanding/filelistingserver/filelisting"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

func errWrapper(
	handler appHandler) func(
	http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter,
		request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			log.Printf("Error occurred "+
				"handling request: %s",
				err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,
				http.StatusText(code), code)
		}
	}
}

func main() {
	// url 里输入一个文件的地址，显示文件的内容
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
