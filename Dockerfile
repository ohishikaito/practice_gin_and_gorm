# # 2020/10/14最新versionを取得
FROM golang:1.15.2-alpine

# # アップデートとgitのインストール！！
RUN apk update && \
  apk add git

# # appディレクトリの作成
# RUN mkdir /go/src/app
# # ワーキングディレクトリの設定
# WORKDIR /go/src/app
# # ホストのファイルをコンテナの作業ディレクトリに移行
# ADD . /go/src/app

# 簡易設定ver
RUN mkdir /app
WORKDIR /app

# realize入れようとしたけど上手くいかん
# https://zenn.dev/h_sakano/articles/b38336d99f43e4e9e90b

# webアプリに必要なライブラリのインストール
RUN go get github.com/gin-gonic/gin && \
  # ginのホットリロード
  go get github.com/codegangsta/gin && \
  go get github.com/go-sql-driver/mysql && \
  go get github.com/jinzhu/gorm
