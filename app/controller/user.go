package controller

// var userService service.UserService

// func (app *App) UserIndexHandler(w http.ResponseWriter, r *http.Request) {
// 	users, _ := userService.GetUserIndex()
// 	type data struct {
// 		Data []model.User `json:"data"`
// 	}
// 	body := data{
// 		Data: users,
// 	}
// 	app.successResponse(w, body, http.StatusOK)
// }

// func (app *App) userCreateHandler(w http.ResponseWriter, r *http.Request) {
// 	var user model.User
// 	json.NewDecoder(r.Body).Decode(&user)
// 	if err := userService.CreateUser(&user); err != nil {
// 		app.errorResponse(w, err.Error(), http.StatusUnprocessableEntity)
// 		return
// 	}
// 	type data struct {
// 		Data model.User `json:"data"`
// 	}
// 	body := data{
// 		Data: user,
// 	}
// 	app.successResponse(w, body, http.StatusCreated)
// }
