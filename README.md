## gomikuji
おみくじAPI

## Features
 * [x] JSON 形式でおみくじの結果を返す
 * [ ] 正月（1/1-1/3）だけ大吉にする
 * [ ] ハンドラのテストを書いてみる

## How to use
### 入手
```
$ go get github.com/yusukemisa/gomikuji
```

### GAEのSDKがある前提
```
#ローカルでサーバー起動
$ dev_appserver.py app.yaml
:
$ curl http://localhost:8080
{"text":"吉","attachments":[{"title":"あなたの運勢","color":"#bce2e8","image_url":"https://raw.githubusercontent.com/yusukemisa/gomikuji/master/omikuji_kichi.png"}]}
```

## call by Slack
![Call by Slack](https://raw.githubusercontent.com/yusukemisa/gomikuji/master/callBySlack.gif)

