package RegistrationController

import (
	"encoding/json"
	"net/http"
	"seredaes/go-passmem/App/DB"
	"seredaes/go-passmem/App/Response"
)

var userData struct {
	Email    string
	Password string
}

func Registration(response http.ResponseWriter, request *http.Request) {

	// GET DATA FROM POST
	err := json.NewDecoder(request.Body).Decode(&userData)
	if err != nil {
		Response.RenderResponse(response, false, "Error decode request", nil, 422)
		return
	}

	// VALIDATION
	// CHECK IF EMAIL EXIST
	if userData.Email == "" {
		Response.RenderResponse(response, false, "Поле email обязательно для заполненения", nil, 422)
		return
	}
	// CHECK IF PASSWORD EXIST
	if userData.Password == "" {
		Response.RenderResponse(response, false, "Поле password обязательно для заполненения", nil, 422)
		return
	}

	// CHECK IF EMAIL IS FREE
	userExist := DB.UserExist(userData.Email)
	if userExist {
		Response.RenderResponse(response, false, "Указанный email не может быть использован", nil, 422)
		return
	}

	// CREATE USER
	DB.CreateUser(userData.Email, userData.Password)
	Response.RenderResponse(response, true, "Пользователь успешно зарегистрирован", nil, 200)

}
