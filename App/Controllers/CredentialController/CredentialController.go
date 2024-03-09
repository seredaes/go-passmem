package CredentialController

import (
	"encoding/json"
	"net/http"
	"seredaes/go-passmem/App/DB"
	"seredaes/go-passmem/App/JWT"
	"seredaes/go-passmem/App/Response"
	"seredaes/go-passmem/App/Types/UserType"
)

var credentialData UserType.Credentials

func CredentialList(response http.ResponseWriter, request *http.Request) {

	userObject, _ := JWT.GetUserFromJWT(request)

	Response.RenderResponse(response, true, "Success", userObject.Data, 200)

}

func validateCredential() (string, bool) {
	if credentialData.Login == "" {
		return "Поле Login обязательно для заполнения", false
	}
	if credentialData.Password == "" {
		return "Поле Password обязательно для заполнения", false
	}
	if credentialData.Description == "" {
		return "Поле Description обязательно для заполнения", false
	}

	return "", true
}

func CreateCredential(response http.ResponseWriter, request *http.Request) {
	// GET DATA FROM POST
	err := json.NewDecoder(request.Body).Decode(&credentialData)
	if err != nil {
		Response.RenderResponse(response, false, "Error decode request", nil, 422)
		return
	}

	msg, error := validateCredential()
	if !error {
		Response.RenderResponse(response, false, msg, nil, 422)
		return
	}

	userObject, _ := JWT.GetUserFromJWT(request)

	DB.CreateCredential(userObject.User.Email, credentialData)
	Response.RenderResponse(response, true, "Данные успешно сохранены", nil, 200)
}

func UpdateCredential(response http.ResponseWriter, request *http.Request) {
	err := json.NewDecoder(request.Body).Decode(&credentialData)
	if err != nil {
		Response.RenderResponse(response, false, "Error decode request", nil, 422)
		return
	}

	msg, error := validateCredential()
	if !error {
		Response.RenderResponse(response, false, msg, nil, 422)
		return
	}

	userObject, _ := JWT.GetUserFromJWT(request)

	status := DB.UpdateCredential(userObject.User.Email, credentialData)
	if status {
		Response.RenderResponse(response, true, "Данные успешно обновлены", nil, 200)
		return
	}

	Response.RenderResponse(response, false, "Данные не обновлены", nil, 422)

}

func DeleteCredential(response http.ResponseWriter, request *http.Request) {

	var ID UserType.CredentialID = UserType.CredentialID{}

	err := json.NewDecoder(request.Body).Decode(&ID)
	if err != nil {
		Response.RenderResponse(response, false, "Error decode request", nil, 422)
		return
	}

	userObject, _ := JWT.GetUserFromJWT(request)

	status := DB.DeleteCredential(userObject.User.Email, ID)
	if status {
		Response.RenderResponse(response, true, "Данные успешно удалены", nil, 200)
		return
	}

	Response.RenderResponse(response, false, "Данные не удалены", nil, 422)
}
