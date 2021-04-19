package config

var (
	SQLDriver   string
	DatabaseURL string
)

func init() {
	driver := "mysql"
	username := "root"
	password := "root"
	// dockerのネットワークを経由してアクセスするから、コンテナ名にしろ
	protocol := "tcp(go_myapp_db_1)"
	dbName := "go_database"

	database := username + ":" + password + "@" + protocol + "/" + dbName + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	SQLDriver = driver
	DatabaseURL = database
}
