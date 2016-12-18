package main

import (
	"fmt"
	"net/http"
	"qiniu"
	"setting"
	"storage"
	"time"
	"tools"
)

var _log = setting.GetLogger("main")

func main() {
	// load configure file

	// regist routing
	mux := tools.NewRESTServeMux()
	mux.HandleRESTful("/qiniu_uptoken", qiniu.UploadToken{})
	mux.HandleFunc("/upload_save", storage.UploadSaveHandler)
	mux.HandleFunc("/assets_list", storage.AssetsListHandle)
	mux.HandleFunc("/upload_delete", storage.DeleteAssets)
	mux.Handle("/", http.FileServer(http.Dir("/Users/junzexu/gocode/static")))

	// serving

	server := &http.Server{
		Addr:           setting.HOST,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	ret := server.ListenAndServe()
	_log.Printf("serving: %s, %#v", setting.HOST, ret)
}

func HelloHanlder(rsp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rsp, "Hello World.")
}
