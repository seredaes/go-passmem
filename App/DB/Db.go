package DB

import (
	"encoding/json"
	"fmt"
	"os"
	CryptPassword "seredaes/go-passmem/App/CryptPassword"
	"seredaes/go-passmem/App/EncryptData"
	"seredaes/go-passmem/App/Env"
	Logger "seredaes/go-passmem/App/Log"
	"seredaes/go-passmem/App/Types/UserType"

	"github.com/google/uuid"
)

var data []UserType.Userdata = []UserType.Userdata{}

const DB_FILE = "./DB/DATA_MODEL.esdata"

func CreateUser(email string, password string) {
	hashedPassword, _ := CryptPassword.HashPassword(password)
	data = append(data, UserType.Userdata{User: UserType.User{Email: email, Password: hashedPassword}})
	StoreDB()
}

func UserExist(email string) bool {
	for _, item := range data {
		if item.User.Email == email {
			return true
		}
	}

	return false
}

func Getuser(email string) UserType.Userdata {
	for _, item := range data {
		if item.User.Email == email {
			return item
		}
	}

	return UserType.Userdata{}
}

func StoreDB() {
	jsonDB, err := json.Marshal(data)
	if err != nil {
		Logger.Log(fmt.Sprint(err), "danger")
	}

	enkKey, _ := Env.GetConfig("ENCRYPT_KEY")
	encryptedDB, _ := EncryptData.Encrypt(string(jsonDB), enkKey)

	if err := os.WriteFile(DB_FILE, []byte(encryptedDB), 0666); err != nil {
		Logger.Log(fmt.Sprint(err), "danger")
	}
}

func RestoreDB() {
	_, err := os.Open(DB_FILE)
	if err != nil {
		StoreDB()
	}

	file, err := os.ReadFile(DB_FILE)
	if err != nil {
		Logger.Log(fmt.Sprint(err), "danger")
	}

	enkKey, _ := Env.GetConfig("ENCRYPT_KEY")
	decryptedDB, _ := EncryptData.Decrypt(string(file), enkKey)

	err = json.Unmarshal([]byte(decryptedDB), &data)
	if err != nil {
		Logger.Log(fmt.Sprint(err), "danger")
	}
}

func CreateCredential(email string, CredentialData UserType.Credentials) {
	for index, item := range data {
		if item.User.Email == email {
			CredentialData.Id = uuid.New()
			data[index].Data = append(item.Data, CredentialData)
		}
	}

	StoreDB()
}

func UpdateCredential(email string, CredentialData UserType.Credentials) bool {

	var userIndex, dataIndex int = -1, -1

	for index, item := range data {
		if item.User.Email == email {
			userIndex = index
		}
	}

	for index, item := range data[userIndex].Data {
		if item.Id == CredentialData.Id {
			dataIndex = index
		}
	}

	if userIndex >= 0 && dataIndex >= 0 {
		data[userIndex].Data[dataIndex].Login = CredentialData.Login
		data[userIndex].Data[dataIndex].Password = CredentialData.Password
		data[userIndex].Data[dataIndex].Description = CredentialData.Description
		data[userIndex].Data[dataIndex].Link = CredentialData.Link
		data[userIndex].Data[dataIndex].Phone = CredentialData.Phone
		StoreDB()
		return true
	}

	return false

}

func DeleteCredential(email string, CredentialData UserType.CredentialID) bool {

	var userIndex, dataIndex int = -1, -1

	for index, item := range data {
		if item.User.Email == email {
			userIndex = index
		}
	}

	for index, item := range data[userIndex].Data {
		if item.Id == CredentialData.Id {
			dataIndex = index
		}
	}

	if userIndex >= 0 && dataIndex >= 0 {
		test := data[userIndex].Data[0:dataIndex]
		test = append(test, data[userIndex].Data[dataIndex+1:len(data[userIndex].Data)]...)
		data[userIndex].Data = test
		StoreDB()
		return true
	}

	return false

}
