# 前提
## containerの立ち上げ
docker-compose build

docker-compose up -d

## 単体で実行
### only gin
docker-compose exec app go run main.go

# デバッカー
`.vscode/launch.json`で
`${workspaceFolder}/app`って書いて、
appディレクトリ配下のファイルを読むようにしてる

## 変更履歴
GOPATHを`/Users/ohishikaido/go`から`/Users/ohishikaido/Desktop/golang`に変更
↓
変えたらbinとかが移動したからやめた