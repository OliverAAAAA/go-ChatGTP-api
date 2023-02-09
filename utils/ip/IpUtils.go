package ip

import (
	"encoding/json"
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"net/http"
)

var API_URL = "http://whois.pconline.com.cn/ipJson.jsp?ip=%s&json=true"

func GetIpAddress(ip string) *string {

	client := &http.Client{}
	//生成Request对象
	req, err := http.NewRequest("GET", fmt.Sprintf(API_URL, ip), nil)
	if err != nil {
		fmt.Println(err)
	}
	//发起请求
	resp, err := client.Do(req)
	//关闭响应体
	defer resp.Body.Close()
	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	//解决中文乱码
	bodystr := mahonia.NewDecoder("gbk").ConvertString(string(body))
	//fmt.Println(bodystr)

	var result struct {
		Addr string `json:"addr"`
	}
	err = json.Unmarshal([]byte(bodystr), &result)
	return &result.Addr
}
