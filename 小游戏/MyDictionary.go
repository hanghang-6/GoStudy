package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type AutoGenerated struct {
	Message []struct {
		Key        string `json:"key"`
		Paraphrase string `json:"paraphrase"`
		Value      int    `json:"value"`
		Means      []struct {
			Part  string   `json:"part"`
			Means []string `json:"means"`
		} `json:"means"`
	} `json:"message"`
	Status int `json:"status"`
}

func UserInputForDic() string {
	fmt.Println("请输入需要查询的单词")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读输入 发生错误")
	}
	input = strings.Replace(input, "\n", "", -1)
	// input = strings.TrimSuffix(input,"\n")
	if err != nil {
		fmt.Println("非法输入，你输入的不是一个int")
	}
	return input
}
func main() {
	word := UserInputForDic()
	client := &http.Client{}
	// 创建请求
	req, err := http.NewRequest("GET", "https://dict.iciba.com/dictionary/word/suggestion?word="+word+"&nums=5&ck=709a0db45332167b0e2ce1868b84773e&timestamp=1722860711643&client=6&uid=123123&key=1000006&is_need_mean=1&signature=dcf7c628e9f52e8325cb85d124bff5f0", nil)
	if err != nil {
		log.Fatal(err)
	}
	// 设置请求头
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Origin", "https://www.iciba.com")
	req.Header.Set("Referer", "https://www.iciba.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 Edg/127.0.0.0")
	req.Header.Set("sec-ch-ua", `"Not)A;Brand";v="99", "Microsoft Edge";v="127", "Chromium";v="127"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// 读取响应
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// 错误日志
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	//fmt.Printf("%s\n", bodyText)
	var autoGenerated AutoGenerated
	err = json.Unmarshal(bodyText, &autoGenerated)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%#v\n", autoGenerated.Message)
	// 输出关键信息
	//fmt.Println(autoGenerated.Message[0].Key, autoGenerated.Message[0].Paraphrase)
	for _, v := range autoGenerated.Message {
		fmt.Println("word:", v.Key)
		for _, means := range v.Means {
			fmt.Println(means.Part, means.Means)
		}
		fmt.Println("===================================")
	}
}
