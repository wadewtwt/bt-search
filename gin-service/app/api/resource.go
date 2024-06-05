package api

// 参考get/post https://www.cnblogs.com/zhaoyingjie/p/16145453.html
import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/form"
	"github.com/putyy/ai-share/app/models"
	"github.com/putyy/ai-share/app/servers"
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

func Search(c *gin.Context) {
	//1影视，2电影集，3福利，4综艺，5合集，6音乐，7港剧，8国产剧，9国产剧，10韩剧，11欧美，12日剧
	//13台剧，14泰剧，15纪录片，16国产动漫，17漫画，18韩漫，19日漫，20欧美动漫，21电子书，22图集 小说 课程 模板 软件 游戏 其它 未知

	ip := c.ClientIP()

	formData := form.SearchKeywordForm{}
	if err1 := c.ShouldBind(&formData); err1 != nil {
		ResponseError(c, "参数错误", err1.Error())
		return
	}
	insertRequestLog(formData.Keyword, ip)

	resultList := make([]map[string]interface{}, 0)

	funletuRes := servers.RequestFunletu(formData.Keyword, formData.PageNo)
	funletuResList := funletuRes.Data
	for i := 0; i < len(funletuResList); i++ {
		funletuItem := funletuResList[i]
		fmt.Println("title is:" + funletuItem.Title)
		tmpMap := map[string]interface{}{}
		tmpMap["title"] = funletuItem.Title
		tmpMap["hot"] = funletuItem.Views
		tmpMap["fileType"] = funletuItem.Filetype
		tmpMap["updateTime"] = funletuItem.UpdateTime
		//tmpMap["douban"] = 6.6

		resultList = append(resultList, tmpMap)
	}

	ResponseSuccess(c, resultList)
}

// 插入log
func insertRequestLog(keyword string, ip string) {
	model := models.RequestLog{
		Keyword: keyword,
		Ip:      ip,
	}
	err := models.Db().Create(&model).Error
	if err != nil {
		return
	}
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
