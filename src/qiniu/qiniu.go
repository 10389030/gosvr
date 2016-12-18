package qiniu

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type UploadPolicy struct {
	Scope               string `json:"scope,omitempty"`
	Deadline            int64  `json:"deadline,omitempty"`
	InsertOnly          int    `json:"insertOnly,omitempty"`
	EndUser             string `json:"endUser,omitempty"`
	ReturnUrl           string `json:"returnUrl,omitempty"`
	ReturnBody          string `json:"returnBody,omitempty"`
	CallbackUrl         string `json:"callbackUrl,omitempty"`
	CallbackHost        string `json:"callbackHost,omitempty"`
	CallbackBody        string `json:"callbackBody,omitempty"`
	CallbackBodyType    string `json:"callbackBodyType,omitempty"`
	CallbackFetchKey    int    `json:"callbackFetchKey,omitempty"`
	PersistentOps       string `json:"persistentOps,omitempty"`
	PersistentNotifyUrl string `json:"persistentNotifyUrl,omitempty"`
	PersistentPipeline  string `json:"persistentPipeline,omitempty"`
	SaveKey             string `json:"saveKey,omitempty"`
	FsizeMin            int64  `json:"fsizeMin,omitempty"`
	FsizeLimit          int64  `json:"fsizeLimit,omitempty"`
	DetectMime          int    `json:"detectMime,omitempty"`
	MimeLimit           string `json:"mimeLimit,omitempty"`
	DeleteAfterDays     int    `json:"deleteAfterDays,omitempty"`
}

type UploadToken struct {
	UpToken string `json:"uptoken"`
}

func (token UploadToken) Get(rsp http.ResponseWriter, req *http.Request) {
	// TODO: set client info into DefaultUploadPolicy
	uploadPolicy := DefaultUploadPolicy
	uploadPolicy.Deadline = time.Now().Unix() + 7*24*3600
	// uploadPolicy.EndUser = //
	// to json
	json_buf, _ := json.Marshal(uploadPolicy)
	//_log.Printf("policy: %s\n", json_buf)

	// base64 policy
	base_buf := base64.URLEncoding.EncodeToString(json_buf)
	// _log.Printf("policy base64: %s\n", base_buf)
	// HMAC-SHA1
	h := hmac.New(sha1.New, []byte(QN_SecretKey))
	h.Write([]byte(base_buf))
	sign := h.Sum(nil)
	// base64 hmac_sha1
	encoded_sign := base64.URLEncoding.EncodeToString(sign)
	// _log.Printf("sign: %s\n", encoded_sign)

	// _log.Printf("keys: accesskey: %s, secrect keys: %s\n", QN_AccessKey, QN_SecretKey)

	// join
	parts := []string{QN_AccessKey, encoded_sign, base_buf}
	token.UpToken = strings.Join(parts, ":")

	jsonBody, _ := json.Marshal(token)
	fmt.Fprintf(rsp, "%s", jsonBody)
}
