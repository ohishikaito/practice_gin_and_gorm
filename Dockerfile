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
# RUN go get -u github.com/oxequa/realize

# RUN apk add --no-cache alpine-sdk git && \
# お手本は↑だけどキャッシュ使いたいから↓にしたい
RUN apk add alpine-sdk git && \
  go get -u github.com/oxequa/realize
EXPOSE 8080
CMD ["realize", "start"]

# webアプリに必要なライブラリのインストール
RUN go get github.com/gin-gonic/gin && \
  go get github.com/go-sql-driver/mysql && \
  go get github.com/jinzhu/gorm
