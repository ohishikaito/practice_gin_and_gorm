module app

// 使わないのを消すコマンド
// go mod tidy -v

go 1.15

require (
	firebase.google.com/go v3.13.0+incompatible
	firebase.google.com/go/v4 v4.5.0 // indirect
	github.com/auth0/go-jwt-middleware v1.0.0
	github.com/form3tech-oss/jwt-go v3.2.2+incompatible
	github.com/gin-contrib/cors v1.3.1 // indirect
	github.com/gin-gonic/gin v1.7.1
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/gorm v1.9.16
	github.com/kr/pretty v0.1.0 // indirect
	github.com/mattn/go-sqlite3 v2.0.1+incompatible // indirect
	golang.org/x/sys v0.0.0-20210423082822-04245dca01da // indirect
	google.golang.org/api v0.40.0
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)
