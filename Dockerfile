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

# realize
# https://qiita.com/rin1208/items/64a6bc469d19ad0ec981
RUN go get -u github.com/oxequa/realize
ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64 \
  GO111MODULE=on
EXPOSE 8080
# ここのスペースほしいw
CMD ["realize", "start", "--build","--run"]

# webアプリに必要なライブラリのインストール
RUN go get github.com/gin-gonic/gin && \
  go get github.com/go-sql-driver/mysql && \
  go get github.com/jinzhu/gorm
