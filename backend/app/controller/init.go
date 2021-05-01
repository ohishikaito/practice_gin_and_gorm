package controller

import (
	"app/auth"

	firebase "firebase.google.com/go"
)

var (
	firebaseApp *firebase.App
)

func init() {
	FirebaseInit()
}

func FirebaseInit() {
	firebaseApp = auth.NewFirebaseApp()
}
