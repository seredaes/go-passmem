package JWT

import (
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"seredaes/go-passmem/App/DB"
	"seredaes/go-passmem/App/Env"
	"seredaes/go-passmem/App/Types/UserType"
	"strings"
)

func GenerateJWT(email string) string {
	header := b64.StdEncoding.EncodeToString([]byte("{\"alg\":\"sha256\"}"))
	payload := b64.StdEncoding.EncodeToString([]byte("{\"email\":\"" + email + "\"}"))
	signature := header + "." + payload

	secret, _ := Env.GetConfig("TOKEN_KEY")

	hmac := hmac.New(sha256.New, []byte(secret))
	hmac.Write([]byte(signature))
	dataHmac := hmac.Sum(nil)

	hmacHex := hex.EncodeToString(dataHmac)

	return header + "." + payload + "." + hmacHex
}

func CheckJWT(JWToken string) bool {

	var User struct {
		Email string
	}

	jwtArray := strings.Split(JWToken, ".")

	if len(jwtArray) != 3 {
		return false
	}

	payload, err := b64.StdEncoding.DecodeString(jwtArray[1])
	if err != nil {
		return false
	}

	err = json.Unmarshal(payload, &User)
	if err != nil {
		return false
	}

	if User.Email == "" {
		return false
	}

	newToken := GenerateJWT(User.Email)

	return strings.Compare(newToken, JWToken) == 0
}

func GetUserFromJWT(request *http.Request) (UserType.Userdata, bool) {

	JWToken := request.Header.Get("Authorization")
	JWToken = strings.Replace(JWToken, "Bearer ", "", -1)

	var User struct {
		Email string
	}

	jwtArray := strings.Split(JWToken, ".")

	if len(jwtArray) != 3 {
		return UserType.Userdata{}, false
	}

	payload, err := b64.StdEncoding.DecodeString(jwtArray[1])
	if err != nil {
		return UserType.Userdata{}, false
	}

	err = json.Unmarshal(payload, &User)
	if err != nil {
		return UserType.Userdata{}, false
	}

	return DB.Getuser(User.Email), true
}
