package helper

import (
	"encoding/json"
	"github.com/Jeffail/gabs"
	"gopkg.in/gavv/httpexpect.v2"
	lhrHttpExpect "lhr/foundation/expect/httpexpect"
	"net/url"
	"strings"
	"testing"
	"time"
)

func FaceGoHttpPostRequest(t *testing.T, reqUrl string, reqPath string, reqJsonStr string) (*httpexpect.Response, error) {
	_, err := url.ParseRequestURI(reqUrl)

	if err != nil {
		t.Fail()
		return nil, err
	}

	if strings.HasSuffix(reqUrl, "/") {
		reqUrl = reqUrl[1 : len(reqUrl)-1]
	}
	if !strings.HasPrefix(reqPath, "/") {
		reqPath = "/" + reqPath
	}
	realReqJsonStr, err := FileInStringChanger(reqJsonStr)
	if err != nil {
		t.Fail()
		return nil, err
	}
	requestContent := make(map[string]interface{})
	jsonParsed, err := gabs.ParseJSON([]byte(realReqJsonStr))
	baseParsed, _ := gabs.ParseJSON([]byte(`{"app_id": "system","app_secret": "12345"}`))
	if err != nil {
		t.Fail()
		return nil, err
	}

	err = jsonParsed.Merge(baseParsed)
	if err != nil {
		t.Fail()
		return nil, err
	}

	err = json.Unmarshal([]byte(jsonParsed.String()), &requestContent)
	if err != nil {
		t.Fail()
		return nil, err
	}
	expect := lhrHttpExpect.NewHttpExpect(t, reqUrl)
	res := expect.POST(reqPath).WithHeader("Content-Type", "application/json").WithJSON(requestContent).Expect()
	return res, nil
}

func FaceGoHttpPostRequestWithTimeout(t *testing.T, reqUrl string, reqPath string, reqJsonStr string, timeout time.Duration) (*httpexpect.Response, error) {
	_, err := url.ParseRequestURI(reqUrl)

	if err != nil {
		t.Fail()
		return nil, err
	}

	if strings.HasSuffix(reqUrl, "/") {
		reqUrl = reqUrl[1 : len(reqUrl)-1]
	}
	if !strings.HasPrefix(reqPath, "/") {
		reqPath = "/" + reqPath
	}
	realReqJsonStr, err := FileInStringChanger(reqJsonStr)
	if err != nil {
		t.Fail()
		return nil, err
	}
	requestContent := make(map[string]interface{})
	jsonParsed, err := gabs.ParseJSON([]byte(realReqJsonStr))
	baseParsed, _ := gabs.ParseJSON([]byte(`{"app_id": "system","app_secret": "12345"}`))
	if err != nil {
		t.Fail()
		return nil, err
	}

	err = jsonParsed.Merge(baseParsed)
	if err != nil {
		t.Fail()
		return nil, err
	}

	err = json.Unmarshal([]byte(jsonParsed.String()), &requestContent)
	if err != nil {
		t.Fail()
		return nil, err
	}
	expect := lhrHttpExpect.NewHttpExpectWithTimeout(t, reqUrl, timeout)
	res := expect.POST(reqPath).WithHeader("Content-Type", "application/json").WithJSON(requestContent).Expect()
	return res, nil
}

func VmsHttpGetRequest(t *testing.T, reqUrl string, reqPath string, withQuery map[string]interface{}, cookies map[string]string, contentType string) (*httpexpect.Response, error) {
	_, err := url.ParseRequestURI(reqUrl)

	if err != nil {
		t.Fail()
		return nil, err
	}

	if strings.HasSuffix(reqUrl, "/") {
		reqUrl = reqUrl[1 : len(reqUrl)-1]
	}
	if !strings.HasPrefix(reqPath, "/") {
		reqPath = "/" + reqPath
	}

	expect := lhrHttpExpect.NewHttpExpect(t, reqUrl)
	var getRequest *httpexpect.Request
	getRequest = expect.GET(reqPath)
	if len(withQuery) > 0 {
		for key, _ := range withQuery {
			getRequest = getRequest.WithQuery(key, withQuery[key])
		}
	}
	if len(contentType) == 0 {
		res := getRequest.WithCookies(cookies).WithHeader("Content-Type", "application/json").Expect()

		return res, nil
	} else {
		res := getRequest.WithCookies(cookies).WithHeader("Content-Type", contentType).Expect()

		return res, nil
	}
}

func VmsHttpPutRequest(t *testing.T, reqUrl string, reqPath string, reqJsonStr string, cookies map[string]string) (*httpexpect.Response, error) {
	_, err := url.ParseRequestURI(reqUrl)

	if err != nil {
		t.Fail()
		return nil, err
	}

	if strings.HasSuffix(reqUrl, "/") {
		reqUrl = reqUrl[1 : len(reqUrl)-1]
	}
	if !strings.HasPrefix(reqPath, "/") {
		reqPath = "/" + reqPath
	}
	realReqJsonStr, err := FileInStringChanger(reqJsonStr)
	if err != nil {
		t.Fail()
		return nil, err
	}
	requestContent := make(map[string]interface{})
	jsonParsed, err := gabs.ParseJSON([]byte(realReqJsonStr))
	if err != nil {
		t.Fail()
		return nil, err
	}

	err = json.Unmarshal([]byte(jsonParsed.String()), &requestContent)
	if err != nil {
		t.Fail()
		return nil, err
	}

	expect := lhrHttpExpect.NewHttpExpect(t, reqUrl)
	res := expect.PUT(reqPath).WithCookies(cookies).WithHeader("Content-Type", "application/json").WithJSON(requestContent).Expect()

	return res, nil
}

func VmsHttpPostRequest(t *testing.T, reqUrl string, reqPath string, reqJsonStr string, cookies map[string]string) (*httpexpect.Response, error) {
	_, err := url.ParseRequestURI(reqUrl)

	if err != nil {
		t.Fail()
		return nil, err
	}

	if strings.HasSuffix(reqUrl, "/") {
		reqUrl = reqUrl[1 : len(reqUrl)-1]
	}
	if !strings.HasPrefix(reqPath, "/") {
		reqPath = "/" + reqPath
	}
	//Post请求body为Null
	if len(reqJsonStr) == 0 {
		reqJsonStr = "{}"
	}
	//Post请求body为Not Null
	realReqJsonStr, err := FileInStringChanger(reqJsonStr)
	if err != nil {
		t.Fail()
		return nil, err
	}
	requestContent := make(map[string]interface{})
	jsonParsed, err := gabs.ParseJSON([]byte(realReqJsonStr))
	if err != nil {
		t.Fail()
		return nil, err
	}
	err = json.Unmarshal([]byte(jsonParsed.String()), &requestContent)
	if err != nil {
		t.Fail()
		return nil, err
	}
	//Post请求不带cookie
	if len(cookies) == 0 {
		expect := lhrHttpExpect.NewHttpExpect(t, reqUrl)
		res := expect.POST(reqPath).WithHeader("Content-Type", "application/json").WithJSON(requestContent).Expect()
		return res, nil
	}
	//Post请求带cookie
	expect := lhrHttpExpect.NewHttpExpect(t, reqUrl)
	res := expect.POST(reqPath).WithCookies(cookies).WithHeader("Content-Type", "application/json").WithJSON(requestContent).Expect()
	return res, nil
}

func VmsHttpDeleteRequest(t *testing.T, reqUrl string, reqPath string, reqJsonStr string, cookies map[string]string) (*httpexpect.Response, error) {
	_, err := url.ParseRequestURI(reqUrl)

	if err != nil {
		t.Fail()
		return nil, err
	}

	if strings.HasSuffix(reqUrl, "/") {
		reqUrl = reqUrl[1 : len(reqUrl)-1]
	}
	if !strings.HasPrefix(reqPath, "/") {
		reqPath = "/" + reqPath
	}
	realReqJsonStr, err := FileInStringChanger(reqJsonStr)
	if err != nil {
		t.Fail()
		return nil, err
	}
	requestContent := make(map[string]interface{})
	jsonParsed, err := gabs.ParseJSON([]byte(realReqJsonStr))
	if err != nil {
		t.Fail()
		return nil, err
	}

	err = json.Unmarshal([]byte(jsonParsed.String()), &requestContent)
	if err != nil {
		t.Fail()
		return nil, err
	}

	expect := lhrHttpExpect.NewHttpExpect(t, reqUrl)
	res := expect.DELETE(reqPath).WithCookies(cookies).WithHeader("Content-Type", "application/json").WithJSON(requestContent).Expect()

	return res, nil
}

func VmsHttpPatchRequest(t *testing.T, reqUrl string, reqPath string, reqJsonStr string, cookies map[string]string) (*httpexpect.Response, error) {
	_, err := url.ParseRequestURI(reqUrl)

	if err != nil {
		t.Fail()
		return nil, err
	}

	if strings.HasSuffix(reqUrl, "/") {
		reqUrl = reqUrl[1 : len(reqUrl)-1]
	}
	if !strings.HasPrefix(reqPath, "/") {
		reqPath = "/" + reqPath
	}

	//Patch没有请求body
	if len(reqJsonStr) == 0 {
		expect := lhrHttpExpect.NewHttpExpect(t, reqUrl)
		res := expect.PATCH(reqPath).WithCookies(cookies).WithHeader("Content-Type", "application/json").Expect()

		return res, nil
	}

	//Patch请求有请求body
	realReqJsonStr, err := FileInStringChanger(reqJsonStr)
	if err != nil {
		t.Fail()
		return nil, err
	}
	requestContent := make(map[string]interface{})
	jsonParsed, err := gabs.ParseJSON([]byte(realReqJsonStr))
	if err != nil {
		t.Fail()
		return nil, err
	}

	err = json.Unmarshal([]byte(jsonParsed.String()), &requestContent)
	if err != nil {
		t.Fail()
		return nil, err
	}

	expect := lhrHttpExpect.NewHttpExpect(t, reqUrl)
	res := expect.PATCH(reqPath).WithCookies(cookies).WithHeader("Content-Type", "application/json").WithJSON(requestContent).Expect()

	return res, nil
}

func AiCameraHttpDeleteRequest(t *testing.T, reqUrl string, reqPath string, reqJsonStr string, cookies map[string]string) (*httpexpect.Response, error) {
	_, err := url.ParseRequestURI(reqUrl)

	if err != nil {
		t.Fail()
		return nil, err
	}

	if strings.HasSuffix(reqUrl, "/") {
		reqUrl = reqUrl[1 : len(reqUrl)-1]
	}
	if !strings.HasPrefix(reqPath, "/") {
		reqPath = "/" + reqPath
	}
	realReqJsonStr, err := FileInStringChanger(reqJsonStr)
	if err != nil {
		t.Fail()
		return nil, err
	}
	requestContent := make(map[string]interface{})
	jsonParsed, err := gabs.ParseJSON([]byte(realReqJsonStr))
	if err != nil {
		t.Fail()
		return nil, err
	}

	err = json.Unmarshal([]byte(jsonParsed.String()), &requestContent)
	if err != nil {
		t.Fail()
		return nil, err
	}
	expect := lhrHttpExpect.NewHttpExpect(t, reqUrl)
	res := expect.DELETE(reqPath).WithCookies(cookies).WithHeader("Content-Type", "application/json").WithJSON(requestContent).Expect()
	return res, nil
}
func AiCameraHttpGetRequest(t *testing.T, reqUrl string, reqPath string, cookies map[string]string) (*httpexpect.Response, error) {
	_, err := url.ParseRequestURI(reqUrl)

	if err != nil {
		t.Fail()
		return nil, err
	}

	if strings.HasSuffix(reqUrl, "/") {
		reqUrl = reqUrl[1 : len(reqUrl)-1]
	}
	if !strings.HasPrefix(reqPath, "/") {
		reqPath = "/" + reqPath
	}
	expect := lhrHttpExpect.NewHttpExpect(t, reqUrl)
	res := expect.GET(reqPath).WithCookies(cookies).WithHeader("Content-Type", "text/html").Expect()
	return res, nil
}
func AiCameraHttpGetRequest_Image(t *testing.T, reqUrl string, reqPath string, reqJsonStr string, cookies map[string]string) (*httpexpect.Response, error) {
	_, err := url.ParseRequestURI(reqUrl)

	if err != nil {
		t.Fail()
		return nil, err
	}

	if strings.HasSuffix(reqUrl, "/") {
		reqUrl = reqUrl[1 : len(reqUrl)-1]
	}
	if !strings.HasPrefix(reqPath, "/") {
		reqPath = "/" + reqPath
	}
	realReqJsonStr, err := FileInStringChanger(reqJsonStr)
	if err != nil {
		t.Fail()
		return nil, err
	}
	requestContent := make(map[string]interface{})
	jsonParsed, err := gabs.ParseJSON([]byte(realReqJsonStr))
	if err != nil {
		t.Fail()
		return nil, err
	}

	err = json.Unmarshal([]byte(jsonParsed.String()), &requestContent)
	if err != nil {
		t.Fail()
		return nil, err
	}
	expect := lhrHttpExpect.NewHttpExpect(t, reqUrl)
	res := expect.GET(reqPath).WithCookies(cookies).WithHeader("Content-Type", "text/html").WithJSON(requestContent).Expect()
	return res, nil
}

func AiCameraHttpPutRequest(t *testing.T, reqUrl string, reqPath string, reqJsonStr string, cookies map[string]string) (*httpexpect.Response, error) {
	_, err := url.ParseRequestURI(reqUrl)

	if err != nil {
		t.Fail()
		return nil, err
	}

	if strings.HasSuffix(reqUrl, "/") {
		reqUrl = reqUrl[1 : len(reqUrl)-1]
	}
	if !strings.HasPrefix(reqPath, "/") {
		reqPath = "/" + reqPath
	}
	realReqJsonStr, err := FileInStringChanger(reqJsonStr)
	if err != nil {
		t.Fail()
		return nil, err
	}
	requestContent := make(map[string]interface{})
	jsonParsed, err := gabs.ParseJSON([]byte(realReqJsonStr))
	if err != nil {
		t.Fail()
		return nil, err
	}

	err = json.Unmarshal([]byte(jsonParsed.String()), &requestContent)
	if err != nil {
		t.Fail()
		return nil, err
	}
	expect := lhrHttpExpect.NewHttpExpect(t, reqUrl)
	res := expect.PUT(reqPath).WithCookies(cookies).WithHeader("Content-Type", "application/json").WithJSON(requestContent).Expect()
	return res, nil
}
func AiCameraHttpPutRtsp(t *testing.T, reqUrl string, reqPath string, reqJsonStr string, cookies map[string]string) (*httpexpect.Response, error) {
	_, err := url.ParseRequestURI(reqUrl)

	if err != nil {
		t.Fail()
		return nil, err
	}

	if strings.HasSuffix(reqUrl, "/") {
		reqUrl = reqUrl[1 : len(reqUrl)-1]
	}
	if !strings.HasPrefix(reqPath, "/") {
		reqPath = "/" + reqPath
	}
	realReqJsonStr, err := FileInStringChanger(reqJsonStr)
	if err != nil {
		t.Fail()
		return nil, err
	}
	requestContent := make(map[string]interface{})
	jsonParsed, err := gabs.ParseJSON([]byte(realReqJsonStr))
	if err != nil {
		t.Fail()
		return nil, err
	}

	err = json.Unmarshal([]byte(jsonParsed.String()), &requestContent)
	if err != nil {
		t.Fail()
		return nil, err
	}
	expect := lhrHttpExpect.NewHttpExpect(t, reqUrl)
	res := expect.PUT(reqPath).WithCookies(cookies).WithHeader("Content-Type", "text/html").WithJSON(requestContent).Expect()
	return res, nil
}

func AiCameraHttpPostRequest(t *testing.T, reqUrl string, reqPath string, reqJsonStr string, cookies map[string]string) (*httpexpect.Response, error) {
	_, err := url.ParseRequestURI(reqUrl)

	if err != nil {
		t.Fail()
		return nil, err
	}

	if strings.HasSuffix(reqUrl, "/") {
		reqUrl = reqUrl[1 : len(reqUrl)-1]
	}
	if !strings.HasPrefix(reqPath, "/") {
		reqPath = "/" + reqPath
	}
	realReqJsonStr, err := FileInStringChanger(reqJsonStr)
	if err != nil {
		t.Fail()
		return nil, err
	}
	requestContent := make(map[string]interface{})
	jsonParsed, err := gabs.ParseJSON([]byte(realReqJsonStr))
	if err != nil {
		t.Fail()
		return nil, err
	}

	err = json.Unmarshal([]byte(jsonParsed.String()), &requestContent)
	if err != nil {
		t.Fail()
		return nil, err
	}
	expect := lhrHttpExpect.NewHttpExpect(t, reqUrl)
	res := expect.POST(reqPath).WithCookies(cookies).WithHeader("Content-Type", "application/json").WithJSON(requestContent).Expect()
	return res, nil
}
func AiCameraHttpGetRequest_GetExt(t *testing.T, reqUrl string, reqPath string, withQuery map[string]interface{}, cookies map[string]string) (*httpexpect.Response, error) {
	_, err := url.ParseRequestURI(reqUrl)

	if err != nil {
		t.Fail()
		return nil, err
	}

	if strings.HasSuffix(reqUrl, "/") {
		reqUrl = reqUrl[1 : len(reqUrl)-1]
	}
	if !strings.HasPrefix(reqPath, "/") {
		reqPath = "/" + reqPath
	}
	if err != nil {
		t.Fail()
		return nil, err
	}
	if err != nil {
		t.Fail()
		return nil, err
	}

	expect := lhrHttpExpect.NewHttpExpect(t, reqUrl)
	var res *httpexpect.Response
	request := expect.GET(reqPath)
	for k, _ := range withQuery {
		//追加get参数
		request = request.WithQuery(k, withQuery[k])
	}

	res = request.WithCookies(cookies).WithHeader("Content-Type", "application/json").Expect()

	return res, nil
}
