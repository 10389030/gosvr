package storage

import (
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"setting"
)

// package storage settings

type Void struct{}

var (
	_pkgName = reflect.TypeOf(Void{}).PkgPath()
	_log     = setting.GetLogger(_pkgName)
)

// define data struct
// resources: img or other
const (
	ObjectStatusOK = 0
)

type TblBase struct {
	ObjectId bson.ObjectId       `bson:"_id" json:"_id"`
	CreateTs bson.MongoTimestamp `json:"createTs,int64"`
	UpdateTs bson.MongoTimestamp `json:"updateTs,int64"`
	OwnerId  bson.ObjectId       `bson:",omitempty" json:"ownerId"`
	Status   int16               `json:"status"`
}

type TblAssets struct {
	TblBase   `bson:",inline"`
	Key       string `json:"key"`
	Hash      string `json:"hash"`
	PicWidth  uint32 `json:"w"`
	PicHeight uint32 `json:"h"`
	FileSize  uint64 `json:"fsize"`
	MIMEType  string `json:"mimeType"`
	Fname     string `json:"fname"`
}
