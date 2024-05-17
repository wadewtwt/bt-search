package api

// 参考get/post https://www.cnblogs.com/zhaoyingjie/p/16145453.html
import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/form"
	"io"
	"net/http"
)

type result struct {
	Msg  interface{} `json:"msg"`
	Code int         `json:"code"`
	Data []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

type remoteResp struct {
	Message interface{} `json:"message"`
	Status  int         `json:"status"`
	Total   int         `json:"total"`
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

func Search(c *gin.Context) {
	//1影视，2电影集，3福利，4综艺，5合集，6音乐，7港剧，8国产剧，9国产剧，10韩剧，11欧美，12日剧
	//13台剧，14泰剧，15纪录片，16国产动漫，17漫画，18韩漫，19日漫，20欧美动漫，21电子书，22图集 小说 课程 模板 软件 游戏 其它 未知

	formData := form.SearchKeywordForm{}
	if err1 := c.ShouldBind(&formData); err1 != nil {
		ResponseError(c, "参数错误", err1.Error())
		return
	}

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
		"searchtext": formData.Keyword,
	}
	data["query"] = innerMap

	page := map[string]int{
		"pageSize":  10,
		"pageIndex": formData.PageNo,
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

	var res remoteResp
	_ = json.Unmarshal(body, &res)

	ResponseSuccess(c, res.Data)
	return
}

func getTest(c *gin.Context) {
	rUrl := "http://127.0.0.1:12002/api/backend/configure/listApp"
	resp, err := http.Get(rUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	var res result
	_ = json.Unmarshal(body, &res)
	fmt.Printf("%#v", res)

	fmt.Printf("data is:%v", res.Code)
	ResponseSuccess(c, res.Data)
	return
}

func postTest() {
	postBody, _ := json.Marshal(map[string]int{
		"pageNo":   1,
		"pageSize": 20,
	})

	// 将数据转换为字节序列
	requestBody := bytes.NewBuffer(postBody)

	// 发送POST请求
	resp, err := http.Post("http://127.0.0.1:12002/api/backend/repairAddress/list", "application/json", requestBody)
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

	fmt.Println("Response status:", resp.Status)
	fmt.Println("Response body:", string(body))

}
