## Create a new GCP project 
まずプロジェクトを作成する

## GPC SDK導入
以下よりCloud Tools の最新バージョン（192.0.0）をインストールする
https://cloud.google.com/sdk/docs/

## App Engine SDK for Go導入
以下よりApp Engine SDK for Go を入手し導入
https://cloud.google.com/appengine/docs/standard/go/download

以下のようにパスを追加
```
# Google App Engine SDK
export PATH="/Users/yusukemisawa/dev/gcp/go_appengine:$PATH";
```

## app.yamlの作成
下記内容のapp.yamlを作成する。このファイルはGAEの設定ファイルと思っておけば良い
```
runtime: go
api_version: go1

handlers:
- url: /.*　        #任意のURL
  script: _go_app   #あとで意味を調べる
```

## ローカルでサーバー動かす
前提条件：app.yamlとmain関数が定義された.goファイルがあること

```
#ローカルサーバー起動
$ dev_appserver.py app.yaml

#デプロイ
$ gcloud app deploy

```