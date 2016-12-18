package storage

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"setting"
	"tools"
)

func AssetsListHandle(rsp http.ResponseWriter, req *http.Request) {
	// step: query mongodb
	session, err := mgo.Dial(setting.MONGO_HOST)
	if err != nil {
		tools.WriteJsonResponse(rsp, http.StatusInternalServerError, err.Error())
		_log.Panic("connecting mgo error: %#v", err)
		return
	}
	defer session.Close()
	tbl_assets := session.DB(setting.MONGO_ASSETS_DB).C(setting.MONGO_ASSETS_TBL)

	var assets_list []TblAssets
	err = tbl_assets.Find(bson.M{}).All(&assets_list)

	rst := map[string]interface{}{
		"assets_list": assets_list,
		"total":       len(assets_list),
	}
	tools.WriteJsonResponse(rsp, http.StatusOK, rst)
}
