package storage

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"setting"
	"tools"
)

func DeleteAssets(rsp http.ResponseWriter, req *http.Request) {
	body_buf, err := ioutil.ReadAll(req.Body)

	var asset TblAssets
	err = bson.UnmarshalJSON(body_buf, &asset)
	if err != nil {
		_log.Printf("parameter error! %#v, %#v\n", asset, err)
		tools.WriteJsonResponse(rsp, http.StatusBadRequest, err.Error())
		return
	}

	// step: save to mongodb
	session, err := mgo.Dial(setting.MONGO_HOST)
	if err != nil {
		tools.WriteJsonResponse(rsp, http.StatusInternalServerError, err.Error())
		_log.Panic("connecting mgo error: %#v", err)
		return
	}
	defer session.Close()
	tbl_assets := session.DB(setting.MONGO_ASSETS_DB).C(setting.MONGO_ASSETS_TBL)
	filter := asset
	var find_rst []TblAssets
	err = tbl_assets.Find(filter).All(&find_rst)
	_log.Printf("%#v %#v\n", filter, find_rst)

	info, err := tbl_assets.RemoveAll(bson.M{"_id": asset.ObjectId})

	if err != nil {
		tools.WriteJsonResponse(rsp, http.StatusInternalServerError, err.Error())
		_log.Panic("unknown error: %#v", err)
		return
	} else {
		tools.WriteJsonResponse(rsp, http.StatusOK, info)
	}
}
