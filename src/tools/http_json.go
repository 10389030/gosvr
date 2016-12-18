package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonResponse struct {
	HTTPStatus int         `json:"ret"`
	Message    interface{} `json:"msg"`
}

func WriteJsonResponse(rsp http.ResponseWriter, status int, msg interface{}) {
	rsp.Header().Set("Content-Type", "application/json; charset=utf-8")
	rsp.WriteHeader(status)

	jsonRsp := JsonResponse{HTTPStatus: status, Message: msg}
	jsonBody, _ := json.Marshal(jsonRsp)

	fmt.Fprintf(rsp, "%s", jsonBody)
}
