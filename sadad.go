package gosadad

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

const TokenURL = "https://sadad.shaparak.ir/vpg/api/v0/Request/PaymentRequest"
const VerifyURL = "https://sadad.shaparak.ir/api/v0/Advice/Verify"

type TokenRequest struct {
	TerminalID    string `json:"TerminalId"`
	MerchantID    string `json:"MerchantId"`
	Amount        int    `json:"Amount"`
	SignData      string `json:"SignData"`
	ReturnURL     string `json:"ReturnUrl"`
	LocalDateTime string `json:"LocalDateTime"`
	OrderID       string `json:"OrderId"`
	Phone         string `json:"UserId"`
}

type TokenResponse struct {
	ResCode     string  `json:"ResCode"`
	Token       *string `json:"Token"`
	Description string  `json:"Description"`
}

//GetToken Get new token
func GetToken(request TokenRequest, key string) (string, error) {
	k, _ := base64.StdEncoding.DecodeString(key)
	//make SignData
	d, err := TripleEcbDesEncrypt([]byte(request.TerminalID+";"+request.OrderID+";"+strconv.Itoa(request.Amount)), k)
	if err != nil {
		return "", err
	}
	request.SignData = base64.StdEncoding.EncodeToString(d)
	request.LocalDateTime = time.Now().Format("01/02/2006 15:04:05 PM")

	payload, _ := json.Marshal(request)
	webResponse, err := call(TokenURL, bytes.NewBuffer(payload))
	if err != nil {
		return "", errors.New("خطا در ارتباط با بانک")
	}
	response := TokenResponse{}

	if json.Unmarshal(webResponse, &response) != nil {
		return "", errors.New("خطا در بازخوانی اطلاعات بانک")
	}

	if *response.Token != "" {
		return *response.Token, nil
	}
	return "", errors.New("خطا در ارتباط با بانک(" + response.Description + ")")
}

type VerifyResponse struct {
	ResCode       string `json:"ResCode"`
	Amount        int    `json:"Amount"`
	Description   string `json:"Description"`
	RetrivalRefNo string `json:"RetrivalRefNo"`
	SystemTraceNo string `json:"SystemTraceNo"`
	OrderID       int    `json:"OrderId"`
}

//Verify Verify by token
func Verify(token, key string) (*VerifyResponse, error) {
	request := struct {
		Token    string `json:"Token"`
		SignData string `json:"SignData"`
	}{}

	k, _ := base64.StdEncoding.DecodeString(key)
	//make SignData
	d, err := TripleEcbDesEncrypt([]byte(token), k)
	if err != nil {
		return nil, err
	}
	request.SignData = base64.StdEncoding.EncodeToString(d)

	payload, _ := json.Marshal(request)
	webResponse, err := call(VerifyURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, errors.New("خطا در ارتباط با بانک")
	}
	response := VerifyResponse{}
	if json.Unmarshal(webResponse, &response) != nil {
		return nil, errors.New("خطا در بازخوانی اطلاعات بانک")
	}

	return &response, nil
}
