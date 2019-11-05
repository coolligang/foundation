package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"lhr/foundation/definitions"
	"lhr/foundation/helper"
	"net/http"
	"time"
)

const (
	urlStarBoxApiBase  string = "http://starbox.cloudwalk.work/capacity/sdk"
	urlStarBoxFileBase string = "http://ks3.kylin.cloudwalk.work/starbox-prd-ai"
)

var (
	apiGetDirId           = urlStarBoxApiBase + "/dir/id"
	apiGetScrollInDir     = urlStarBoxApiBase + "/file/scroll"
	apiGetNextScroll      = urlStarBoxApiBase + "/file/next_scroll"
	apiGetFileByStorageId = urlStarBoxFileBase
)

func NewStarBoxApi() *StarBoxApi {
	item := new(StarBoxApi)
	item.accessId, _ = helper.GetParameter(definitions.DLhrDatasetStarBoxAccessId)
	item.accessKey, _ = helper.GetParameter(definitions.DLhrDatasetStarBoxAccessKey)
	item.httpClient = http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	return item
}

type StarBoxApi struct {
	accessId   string
	accessKey  string
	httpClient http.Client
}

func (api *StarBoxApi) GetDirId(path string) (dirId json.Number, err error) {
	if api.accessId == "" || api.accessKey == "" {
		err = errors.New("cannot get accessId accessKey")
	} else {
		para := map[string]interface{}{
			"accessid":  api.accessId,
			"accesskey": api.accessKey,
			"path":      path,
		}

		req, err := json.Marshal(para)
		if err == nil {
			res, err := api.httpClient.Post(apiGetDirId, "application/json", bytes.NewBuffer(req))
			if err == nil {
				var result map[string]interface{}
				decoder := json.NewDecoder(res.Body)
				decoder.UseNumber()
				err = decoder.Decode(&result)
				if err == nil {
					dirId = result["data"].(map[string]interface{})["dirid"].(json.Number)
				}
			}
		}
	}
	return dirId, err
}

func (api *StarBoxApi) GetScrollInDir(dirIds json.Number) (scrollId string, totalNum json.Number, filename string, storageId string, err error) {
	if api.accessId == "" || api.accessKey == "" {
		err = errors.New("cannot get accessId accessKey")
	} else {
		para := map[string]interface{}{
			"accessid":   api.accessId,
			"accesskey":  api.accessKey,
			"dirIds":     []json.Number{dirIds},
			"includes":   []string{"filename", "storageId", "dirId"},
			"scrollSize": 1,
		}

		var req []byte
		req, err = json.Marshal(para)
		if err == nil {
			var res *http.Response
			res, err = api.httpClient.Post(apiGetScrollInDir, "application/json", bytes.NewBuffer(req))
			if err == nil {
				var result map[string]interface{}
				decoder := json.NewDecoder(res.Body)
				decoder.UseNumber()
				err = decoder.Decode(&result)
				if err == nil {
					if _, found := result["data"].(map[string]interface{})["dataList"]; found && len(result["data"].(map[string]interface{})["dataList"].([]interface{})) > 0 {
						scrollId = result["data"].(map[string]interface{})["scrollId"].(string)
						totalNum = result["data"].(map[string]interface{})["totalNum"].(json.Number)
						filename = result["data"].(map[string]interface{})["dataList"].([]interface{})[0].(map[string]interface{})["filename"].(string)
						storageId = result["data"].(map[string]interface{})["dataList"].([]interface{})[0].(map[string]interface{})["storageId"].(string)
					} else {
						err = errors.New("no data")
					}
				}
			}
		}
	}
	return scrollId, totalNum, filename, storageId, err
}

func (api *StarBoxApi) GetNextScroll(scrollId string) (totalNum json.Number, filename string, storageId string, err error) {
	if api.accessId == "" || api.accessKey == "" {
		err = errors.New("cannot get accessId accessKey")
	} else {
		para := map[string]interface{}{
			"accessid":  api.accessId,
			"accesskey": api.accessKey,
			"scrollId":  scrollId,
		}

		var req []byte
		req, err = json.Marshal(para)
		if err == nil {
			var res *http.Response
			res, err = api.httpClient.Post(apiGetNextScroll, "application/json", bytes.NewBuffer(req))
			if err == nil {
				var result map[string]interface{}
				decoder := json.NewDecoder(res.Body)
				decoder.UseNumber()
				err = decoder.Decode(&result)
				if err == nil {
					if _, found := result["data"].(map[string]interface{})["dataList"]; found && len(result["data"].(map[string]interface{})["dataList"].([]interface{})) > 0 {
						totalNum = result["data"].(map[string]interface{})["totalNum"].(json.Number)
						filename = result["data"].(map[string]interface{})["dataList"].([]interface{})[0].(map[string]interface{})["filename"].(string)
						storageId = result["data"].(map[string]interface{})["dataList"].([]interface{})[0].(map[string]interface{})["storageId"].(string)
					} else {
						err = errors.New("no data")
					}
				}
			}
		}
	}
	return totalNum, filename, storageId, err
}

func (api *StarBoxApi) GetFileByStorageId(storageId string) (content []byte, err error) {
	res, err := api.httpClient.Get(apiGetFileByStorageId + "/" + storageId)
	if err == nil {
		content, err = ioutil.ReadAll(res.Body)
	}
	return content, err
}

