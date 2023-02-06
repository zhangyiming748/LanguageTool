package LanguageTool

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gookit/goutil/dump"
	"github.com/zhangyiming748/LanguageTool/curl"
	"github.com/zhangyiming748/log"
	"os"
)

type Result struct {
	Software struct {
		Name        string `json:"name"`
		Version     string `json:"version"`
		BuildDate   string `json:"buildDate"`
		ApiVersion  int    `json:"apiVersion"`
		Premium     bool   `json:"premium"`
		PremiumHint string `json:"premiumHint"`
		Status      string `json:"status"`
	} `json:"software"`
	Warnings struct {
		IncompleteResults bool `json:"incompleteResults"`
	} `json:"warnings"`
	Language struct {
		Name             string `json:"name"`
		Code             string `json:"code"`
		DetectedLanguage struct {
			Name       string  `json:"name"`
			Code       string  `json:"code"`
			Confidence float64 `json:"confidence"`
			Source     string  `json:"source"`
		} `json:"detectedLanguage"`
	} `json:"language"`
	Matches []struct {
		Message      string `json:"message"`
		ShortMessage string `json:"shortMessage"`
		Replacements []struct {
			Value string `json:"value"`
		} `json:"replacements"`
		Offset  int `json:"offset"`
		Length  int `json:"length"`
		Context struct {
			Text   string `json:"text"`
			Offset int    `json:"offset"`
			Length int    `json:"length"`
		} `json:"context"`
		Sentence string `json:"sentence"`
		Type     struct {
			TypeName string `json:"typeName"`
		} `json:"type"`
		Rule struct {
			Id          string `json:"id"`
			Description string `json:"description"`
			IssueType   string `json:"issueType"`
			Urls        []struct {
				Value string `json:"value"`
			} `json:"urls"`
			Category struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"category"`
			IsPremium bool `json:"isPremium"`
		} `json:"rule"`
		IgnoreForIncompleteSentence bool `json:"ignoreForIncompleteSentence"`
		ContextForSureMatch         int  `json:"contextForSureMatch"`
	} `json:"matches"`
	SentenceRanges [][]int `json:"sentenceRanges"`
}
type Available struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	LongCode string `json:"longCode"`
}

func GetAvailable() {
	var available []Available
	get, err := curl.HttpGet(nil, nil, "https://api.languagetoolplus.com/v2/languages")
	if err != nil {
		log.Warn.Println(err)
		return
	}
	err = json.Unmarshal(get, &available)
	if err != nil {
		log.Warn.Println(err)
		return
	}
	save(string(get))
	//log.Debug.Printf("%+v\n", available)
	dump.P(available)
	//s := string(get)
	//s = strings.Replace(s, "[", "", 1)
	//s = strings.Replace(s, "]", "", 1)
	//log.Debug.Println(s)
	//ss := strings.Split(s, "},{")
	//for i, v := range ss {
	//	log.Debug.Printf("%d.%s\n", i, v)
	//	//"name":"English (US)","code":"en","longCode":"en-US"
	//	split := strings.Split(v, "\"")
	//	a := &Available{
	//		Name:     split[3],
	//		Code:     split[7],
	//		LongCode: split[11],
	//	}
	//	available = append(available, *a)
	//}
	//log.Debug.Println(available)
}
func save(b string) {
	filePath := "example.json"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString(b)
	write.WriteString("\n")

	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}
