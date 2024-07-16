package servers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RemoteResp struct {
	Message interface{} `json:"message"`
	Status  int         `json:"status"`
	Total   int         `jsoremoteRespn:"total"`
	Text    string      `json:"text"`
	Data    []struct {
		Id         int    `json:"id"`
		Category   string `json:"category"`
		CategoryId int    `json:"categoryid"`
		Course     string `json:"course"`
		CourseId   int    `json:"courseid"`
		CreateTime string `json:"createtime"`
		ExtCode    string `json:"extcode"`
		Filename   string `json:"filename"`
		Filetype   string `json:"filetype"`
		FileTypeId int    `json:"filetypeid"`
		Recommend  int    `json:"recommend"`
		Size       string `json:"size"`
		Sort       int    `json:"sort"`
		State      int    `json:"state"`
		Stated     string `json:"stated"`
		Title      string `json:"title"`
		Top        int    `json:"top"`
		UnzipCode  string `json:"unzipcode"`
		UpdateTime string `json:"updatetime"`
		UploaderId int    `json:"uploaderid"`
		Url        string `json:"url"`
		Valid      int    `json:"valid"`
		Views      int    `json:"views"`
		Violate    int    `json:"violate"`
	} `json:"data"`
}

func RequestFunletu(keyword string, pageNo int) (remoteResp RemoteResp) {

	data := make(map[string]interface{})
	data["style"] = "get"
	data["datasrc"] = "search"

	innerMap := map[string]interface{}{
		"id":         "",
		"datetime":   "",
		"commonid":   1,
		"parmid":     "",
		"fileid":     "",
		"reportid":   "",
		"validid":    "",
		"searchtext": keyword,
	}
	data["query"] = innerMap

	page := map[string]int{
		"pageSize":  10,
		"pageIndex": pageNo,
	}
	data["page"] = page

	order := map[string]string{
		"prop":  "id",
		"order": "desc",
	}
	data["order"] = order

	data["message"] = "请求资源列表数据"

	postBody, _ := json.Marshal(data)

	// 将数据转换为字节序列
	requestBody := bytes.NewBuffer(postBody)

	// 发送POST请求
	resp, err := http.Post("https://v.funletu.com/search", "application/json", requestBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Response body:", string(body))

	var res RemoteResp
	_ = json.Unmarshal(body, &res)
	return res
}
