package messaging

import (
	"encoding/json"
	"fmt"
	result "github.com/heaptracetechnology/microservice-messagebird/result"
	"github.com/messagebird/go-rest-api"
	"github.com/messagebird/go-rest-api/sms"
	"net/http"
	"os"
)

type SMS struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}

//Send SMS
func Send(responseWriter http.ResponseWriter, request *http.Request) {

	var apiKey = os.Getenv("API_KEY")

	decoder := json.NewDecoder(request.Body)
	var param SMS
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	client := messagebird.New(apiKey)

	msg, sendErr := sms.Create(client, param.From, []string{param.To}, param.Message, nil)
	if sendErr != nil {
		fmt.Println("sendErr ::", sendErr)
		result.WriteErrorResponse(responseWriter, sendErr)
		return
	}

	bytes, _ := json.Marshal(msg)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
