package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type JinShanResp struct {
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

type DeelResp struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Translations []struct {
			Beams []struct {
				Sentences []struct {
					Text string `json:"text"`
					Ids  []int  `json:"ids"`
				} `json:"sentences"`
				NumSymbols int `json:"num_symbols"`
			} `json:"beams"`
			Quality string `json:"quality"`
		} `json:"translations"`
		TargetLang            string `json:"target_lang"`
		SourceLang            string `json:"source_lang"`
		SourceLangIsConfident bool   `json:"source_lang_is_confident"`
		DetectedLanguages     struct {
		} `json:"detectedLanguages"`
	} `json:"result"`
}

func UserInputForDic() string {
	fmt.Println("请输入需要查询的单词")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读输入 发生错误")
	}
	input = strings.Replace(input, "\n", "", -1)
	return input
}
func CreateRequestDeepl(ctx context.Context, word string, text chan<- []byte, engineType chan<- string) {
	// https://www.deepl.com/zh/translator#en/zh-hant/chinese
	client := &http.Client{}
	var data = strings.NewReader(`{"jsonrpc":"2.0","method": "LMT_handle_jobs","params":{"jobs":[{"kind":"default","sentences":[{"text":"` + word + `","id":1,"prefix":""}],"raw_en_context_before":[],"raw_en_context_after":[],"preferred_num_beams":4}],"lang":{"target_lang":"ZH","preference":{"weight":{},"default":"default"},"source_lang_computed":"EN"},"priority":-1,"commonJobParams":{"quality":"fast","regionalVariant":"zh-Hant","mode":"translate","browserType":1,"textType":"plaintext"},"timestamp":1723013205878},"id":77130010}`)
	req, err := http.NewRequest("POST", "https://www2.deepl.com/jsonrpc?method=LMT_handle_jobs", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", "INGRESSCOOKIE=ea454a92cb4ed072cda1d584ad5c80bf|a6d4ac311669391fc997a6a267dc91c0; userCountry=TW; dapUid=9f1b746b-d54e-4876-90ec-3d60eb29b961; releaseGroups=3961.B2B-663.2.3_4322.DWFA-689.2.2_10449.DF-3959.2.2_12316.DM-1846.2.1_12541.AAEXP-10770.1.1_220.DF-1925.1.9_2055.DM-814.2.3_2962.DF-3552.2.6_5562.DWFA-732.2.2_6359.DM-1411.2.10_7759.DWFA-814.2.2_12499.MTD-319.1.4_12534.AAEXP-10763.1.1_12535.AAEXP-10764.2.1_1483.DM-821.2.2_2373.DM-1113.2.4_12560.AAEXP-10789.1.1_6732.DF-3818.2.4_7616.DWFA-777.2.2_8776.DM-1442.2.2_12540.AAEXP-10769.2.1_12548.AAEXP-10777.1.1_12556.AAEXP-10785.1.1_12537.AAEXP-10766.2.1_12644.WTT-1296.2.2_12687.TACO-153.1.1_9129.DM-1419.2.2_10382.DF-3962.2.2_11591.TACO-210.2.3_11861.CEX-77.2.4_12542.AAEXP-10771.1.1_12645.DAL-1151.2.1_3283.DWFA-661.2.2_11549.DM-1149.2.2_12532.AAEXP-10761.1.1_12543.AAEXP-10772.1.1_12555.AAEXP-10784.1.1_12559.AAEXP-10788.1.1_8635.DM-1158.2.3_10551.DAL-1134.2.1_10753.TACO-145.2.4_11548.DM-1613.1.1_12533.AAEXP-10762.2.1_12558.AAEXP-10787.1.1_4121.WDW-356.2.5_4854.DM-1255.2.5_7584.TACO-60.2.2_8287.TC-1035.2.5_10550.DWFA-884.2.2_12500.DF-3968.1.1_12538.AAEXP-10767.1.1_12550.AAEXP-10779.2.1_12552.AAEXP-10781.1.1_2413.DWFA-524.2.4_6402.DWFA-716.2.3_8391.DM-1630.2.2_11547.DF-3929.2.2_12539.AAEXP-10768.1.1_12545.AAEXP-10774.2.1_12553.AAEXP-10782.1.1_975.DM-609.2.3_8041.DM-1581.2.2_9683.SEO-747.2.2_9824.AP-523.2.3_10752.TACO-109.2.3_11072.B2B-1154.2.5_12546.AAEXP-10775.1.1_12561.AAEXP-10790.1.1_5719.DWFA-761.2.2_12547.AAEXP-10776.1.1_12551.AAEXP-10780.2.1_12557.AAEXP-10786.1.1_12642.DM-1870.2.1_2455.DPAY-2828.2.2_12074.DEM-1270.1.1_12457.WDW-653.2.1_8393.DPAY-3431.2.2_12536.AAEXP-10765.2.1_1583.DM-807.2.5_3613.WDW-267.2.2_12544.AAEXP-10773.2.1_12549.AAEXP-10778.1.1_12646.MTD-779.2.1_2656.DM-1177.2.2_10794.DF-3869.2.1_12003.DM-1857.2.1_12498.DM-1867.2.2_12554.AAEXP-10783.1.1; LMTBID=v2|edf82a56-039b-44ba-a424-417d3589e051|34586eb73fb0fbf16adc292f7a0b0cf0; privacySettings=%7B%22v%22%3A2%2C%22t%22%3A1722988800%2C%22m%22%3A%22LAX_AUTO%22%2C%22consent%22%3A%5B%22NECESSARY%22%2C%22PERFORMANCE%22%2C%22COMFORT%22%2C%22MARKETING%22%5D%7D; dapVn=1; dapSid=%7B%22sid%22%3A%2231c4c779-b0b2-4a90-b52b-93457a2b3c97%22%2C%22lastUpdate%22%3A1723013205%7D")
	req.Header.Set("origin", "https://www.deepl.com")
	req.Header.Set("priority", "u=1, i")
	req.Header.Set("referer", "https://www.deepl.com/")
	req.Header.Set("sec-ch-ua", `"Not)A;Brand";v="99", "Microsoft Edge";v="127", "Chromium";v="127"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 Edg/127.0.0.0")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//SleepDuration := time.Duration(rand.Intn(3)+1) * time.Second // 随机生成 1 到 7 秒之间的时间
	//time.Sleep(SleepDuration)                                    // 睡眠随机时间
	select {
	case text <- bodyText:
		engineType <- "DeepL"
	case <-ctx.Done():
		return
	}
}
func CreateRequest(ctx context.Context, word string, text chan<- []byte, engineType chan<- string) {
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
	//SleepDuration := time.Duration(rand.Intn(4)+3) * time.Second // 随机生成 1 到 7 秒之间的时间
	//time.Sleep(SleepDuration)                                    // 睡眠随机时间
	select {
	case text <- bodyText:
		engineType <- "金山"
	case <-ctx.Done():
		return
	}

}
func HandleJinShanByteStream(bodyText []byte) {
	//fmt.Printf("%s\n", bodyText)
	// Json-> Go struct
	var jinShanResp JinShanResp
	err := json.Unmarshal(bodyText, &jinShanResp)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%#v\n", jinShanResp.Message)
	// 输出关键信息
	//fmt.Println(jinShanResp.Message[0].Key, jinShanResp.Message[0].Paraphrase)
	for _, v := range jinShanResp.Message {
		fmt.Println("word:", v.Key)
		for _, means := range v.Means {
			fmt.Println(means.Part, means.Means)
		}
		fmt.Println("===================================")
	}
}
func HandleDeepLByteStream(bodyText []byte) {
	//fmt.Printf("%s\n", bodyText)
	// Json-> Go struct
	var resp DeelResp
	err := json.Unmarshal(bodyText, &resp)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range resp.Result.Translations {
		for _, means := range v.Beams {
			fmt.Println(means.Sentences[0].Text)
		}
		fmt.Println("===================================")
	}
}
func HandleAndOutput(bodyText []byte, engineName string) {

	switch engineName {
	case "金山":
		HandleJinShanByteStream(bodyText)
	case "DeepL":
		HandleDeepLByteStream(bodyText)
	}
}
func FuncSearch(word string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	text := make(chan []byte)
	engineType := make(chan string)
	// 请求
	go CreateRequest(ctx, word, text, engineType)
	go CreateRequestDeepl(ctx, word, text, engineType)
	// 等待第一个响应
	textResult := <-text
	engineName := <-engineType
	cancel()
	// 清理剩余的请求
	select {
	case <-text:
		// 另一个引擎的响应被取消或超时
	case <-time.After(500 * time.Millisecond):
		// 等待一段时间，确保取消请求
	}
	fmt.Println(engineName)
	fmt.Println(textResult)
	// 结果解析
	HandleAndOutput(textResult, engineName)

}
func main() {
	for {
		//命令窗口用户输入
		word := UserInputForDic()
		//args形式输入
		//if len(os.Args) != 2 {
		//	fmt.Fprintf(os.Stderr, `usage:simpleDict WORD example: simpleDict hello`)
		//	os.Exit(1)
		//}
		//word := os.Args[1]
		FuncSearch(word)
		//CreateRequestDeepl()
	}
}
