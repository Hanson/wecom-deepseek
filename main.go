package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/http"
	"strings"
)

// 你的 ollama model 名称
var modelName = "deepseek-r1-14b:latest"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("err: %+v", err)
			return
		}

		typ := gjson.GetBytes(b, "type").String()
		userId := gjson.GetBytes(b, "user_id").String()
		sender := gjson.GetBytes(b, "data.sender").String()

		fmt.Println(typ)
		switch typ {
		case "WW_RECV_TEXT_MSG":
			conversationId := gjson.GetBytes(b, "data.conversation_id").String()

			// 私聊
			if strings.HasPrefix(conversationId, "S:") {
				fmt.Println(conversationId)

				// 非自己发的消息
				if sender != userId {
					fmt.Println(gjson.GetBytes(b, "data.content").String())
					resp := chat(gjson.GetBytes(b, "data.content").String())

					fmt.Println("rsp:" + resp)
					if resp != "" {
						sendText(conversationId, resp)
					}
				}
			} else if strings.HasPrefix(conversationId, "R:") {
				// 群聊
			}
		}
	})
	http.ListenAndServe(":8345", nil)
}

func sendText(conversationId, content string) {
	reqJson := fmt.Sprintf(`{"type":"WW_SEND_TEXT_V2_MSG","client_id":1,"data":{"content":"%s","conversation_id":"%s"}}`, content, conversationId)

	// xbot 监听的端口地址
	http.Post("http://127.0.0.1:5882", "application/json", strings.NewReader(reqJson))
}

func chat(content string) (aiRespContent string) {
	reqJson := fmt.Sprintf(`{"model":"%s","stream":false,"messages":[{"role":"user","content":"%s"}]}`, modelName, content)

	// ollama 的对话接口
	resp, err := http.Post("http://127.0.0.1:11434/api/chat", "application/json", strings.NewReader(reqJson))
	if err != nil {
		log.Printf("err: %+v", err)
		return
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("err: %+v", err)
		return
	}

	fmt.Println(string(b))

	aiRespContent = gjson.GetBytes(b, "message.content").String()

	return aiRespContent

}
