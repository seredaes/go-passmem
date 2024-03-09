package Response

import (
	"encoding/json"
	"net/http"
)

var ResponseBody struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func RenderResponse(response http.ResponseWriter, status bool, message string, data interface{}, rc int) {
	ResponseBody.Status = status
	ResponseBody.Message = message

	ResponseBody.Data = data
	response.Header().Set("Content-Type", "application/json; charset=utf-8")

	responseString, _ := json.Marshal(ResponseBody)

	if rc == 0 {
		rc = 200
	}

	if !status {
		response.WriteHeader(rc)
		response.Write(responseString)
	} else {
		response.Write(responseString)
	}
}
