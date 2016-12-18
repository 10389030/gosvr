package storage

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"setting"
	"tools"
)

func UploadSaveHandler(rsp http.ResponseWriter, req *http.Request) {
	// step: read request json into buffer
	body_buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		_log.Printf("read body error: %#v\n", err)
		tools.WriteJsonResponse(rsp, http.StatusBadRequest, err.Error())
		return
	}

	// step: parse buffer into bson object
	var new_asset TblAssets
	err = bson.UnmarshalJSON(body_buf, &new_asset)
	if err != nil {
		_log.Printf("unmarshal body error: %#v\n", err)
		tools.WriteJsonResponse(rsp, http.StatusBadRequest, err.Error())
		return
	}
	new_asset.ObjectId = bson.NewObjectId()
	new_asset.CreateTs = bson.MongoTimestamp(bson.Now().Unix())
	new_asset.UpdateTs = new_asset.CreateTs
	_log.Printf("%#v", new_asset)

	// step: save to mongodb
	session, err := mgo.Dial(setting.MONGO_HOST)
	if err != nil {
		tools.WriteJsonResponse(rsp, http.StatusInternalServerError, err.Error())
		_log.Panic("connecting mgo error: %#v", err)
		return
	}
	defer session.Close()
	tbl_assets := session.DB(setting.MONGO_ASSETS_DB).C(setting.MONGO_ASSETS_TBL)
	err = tbl_assets.Insert(new_asset)

	if err != nil {
		_log.Printf("error: %#v\n", err)
		tools.WriteJsonResponse(rsp, http.StatusInternalServerError, err.Error())
		return
	} else {
		_log.Printf("done")
		tools.WriteJsonResponse(rsp, http.StatusOK, new_asset)
		return
	}
}
