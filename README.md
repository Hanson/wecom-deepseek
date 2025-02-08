# wecom-deepseek

本项目通过本地化部署ollama+deepseek后，通过 ollama 的接口以及第三方企微消息的接口，实现本地话 deepseek 自动回复。可以支持私聊以及群聊，用户可以自行修改 `main.go` 实现想要的不同效果。例如把思考过程 `<think>` 的部分移除等等。

## 项目依赖

* ollama
* deepseek
* [xbot 企微](https://www.apifox.cn/apidoc/shared-d478def0-67c1-4161-b385-eef8a94e9d17)

## 运行
`git clone` 后直接 `go run main.go`

## 效果示例
![49a8163c9674eb3ecc0d2e143294a1b](https://github.com/user-attachments/assets/b61b49c8-5b57-49b1-833b-3938fd3b4d75)

## 交流群
![image](https://github.com/user-attachments/assets/68cd30dd-aa7b-4e69-a75e-de00c0621b2e)
