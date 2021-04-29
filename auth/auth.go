package auth

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func NewFirebaseApp() *firebase.App {
	// firebaseからダウンロードした秘密鍵のjsonをぶち込んでる
	opt := option.WithCredentialsFile("config/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Errorf("error initializing app: %v", err)
		return nil
	}
	return app

	// // Initialize default app
	// app, err := firebase.NewApp(context.Background(), nil)
	// if err != nil {
	// 	fmt.Printf("error initializing app: %v\n", err)
	// }
	// return app
}

func NewAuthClient(app *firebase.App) *auth.Client {
	// Access auth service from the default app
	client, err := app.Auth(context.Background())
	if err != nil {
		fmt.Printf("error getting Auth client: %v\n", err)
	}
	return client
}

// -----------------------
// 以下、JWTの勉強用
// -----------------------

// import (
// 	"net/http"
// 	"time"

// 	jwtmiddleware "github.com/auth0/go-jwt-middleware"
// 	jwt "github.com/form3tech-oss/jwt-go"
// )

// // GetTokenHandler get token
// var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 	// headerのセット
// 	token := jwt.New(jwt.SigningMethodHS256)

// 	// claimsのセット
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["admin"] = true
// 	claims["sub"] = "54546557354"
// 	claims["name"] = "taro"
// 	claims["iat"] = time.Now()
// 	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

// 	// 電子署名
// 	tokenString, _ := token.SignedString([]byte("HOGE"))

// 	// JWTを返却
// 	w.Write([]byte(tokenString))
// })

// // JwtMiddleware check token
// var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
// 	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
// 		return []byte("HOGE"), nil
// 	},
// 	SigningMethod: jwt.SigningMethodHS256,
// })
