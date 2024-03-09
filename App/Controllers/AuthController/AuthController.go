package AuthController

import (
	"encoding/json"
	"net/http"
	"seredaes/go-passmem/App/CryptPassword"
	"seredaes/go-passmem/App/DB"
	"seredaes/go-passmem/App/JWT"
	Logger "seredaes/go-passmem/App/Log"
	"seredaes/go-passmem/App/Response"
)

var userData struct {
	Email    string
	Password string
}

func Auth(response http.ResponseWriter, request *http.Request) {

	// GET DATA FROM POST
	err := json.NewDecoder(request.Body).Decode(&userData)
	if err != nil {
		IP_ADDRESS := request.RemoteAddr
		Logger.Log("[AUTH INCORRECT] IP: "+IP_ADDRESS, "warning")
		Response.RenderResponse(response, false, "Авторизация не прошла", nil, 422)
		return
	}

	// VALIDATION
	if userData.Email == "" {
		Response.RenderResponse(response, false, "Поле email обязательно для заполненения", nil, 422)
		return
	}

	if userData.Password == "" {
		Response.RenderResponse(response, false, "Поле password обязательно для заполненения", nil, 422)
		return
	}

	// CHECK IF EMAIL IS EXIST
	userExist := DB.UserExist(userData.Email)
	if !userExist {
		Response.RenderResponse(response, false, "Авторизация не прошла", nil, 422)
		return
	}

	userObject := DB.Getuser(userData.Email)

	passwordChecked := CryptPassword.CheckPassword(userData.Password, userObject.User.Password)

	if !passwordChecked {
		Response.RenderResponse(response, false, "Авторизация не прошла 1", nil, 422)
		return
	}

	jwtToken := map[string]string{"token": JWT.GenerateJWT(userData.Email)}

	Response.RenderResponse(response, true, "Авторизация успешно прошла", jwtToken, 200)

}
