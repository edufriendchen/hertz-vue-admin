package utils

import (
	"encoding/json"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"net/http"
)

type Result struct {
	Addr string `json:"addr"`
}

func GetAddressByIp(ip string) (string, error) {
	ip = "http://whois.pconline.com.cn/ipJson.jsp?ip=" + ip + "&json=true"
	request, _ := http.NewRequest("GET", ip, nil)
	client := &http.Client{}
	request.Header.Set("Content-Type", "application/json; charset=gbk")
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	var s Result
	json.Unmarshal([]byte(mahonia.NewDecoder("gbk").ConvertString(string(body))), &s)
	return s.Addr, nil
}
