package libs

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	Api = "https://api-v2.thaibulksms.com"
)

type Response struct {
	StatusCode int
	Data       SuccessResponse
	Error      struct {
		Code        int    `json:"code"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"error"`
}
type SuccessResponse struct {
	RemainingCredit int    `json:"remaining_credit"`
	TotalUseCredit  int    `json:"total_use_credit"`
	CreditType      string `json:"credit_type"`
	PhoneNumberList []struct {
		Number     string `json:"number"`
		MessageID  string `json:"message_id"`
		UsedCredit int    `json:"used_credit"`
	} `json:"phone_number_list"`
	BadPhoneNumberList []interface{} `json:"bad_phone_number_list"`
}

type ErrorResponse struct {
	Error struct {
		Code        int    `json:"code"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"error"`
}

type Sms struct {
	Msisdn  string
	Message string
}

func changeBase64(username string, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (s Sms) Send() (*Response, error) {
	ApiKey := os.Getenv("SMS_API_KEY")
	ApiSecretKey := os.Getenv("SMS_API_SECRET_KEY")
	Sender := os.Getenv("SMS_SENDER")
	payload := strings.NewReader("msisdn=" + s.Msisdn + "&message=" + s.Message + "&sender=" + Sender)
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	url := Api + "/sms"
	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Basic "+changeBase64(ApiKey, ApiSecretKey))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var res Response
	res.StatusCode = resp.StatusCode
	if resp.StatusCode == 201 {
		err := json.Unmarshal(body, &res.Data)
		if err != nil {
			return nil, err
		}
	} else {
		var e ErrorResponse
		err := json.Unmarshal(body, &e)
		res.Error = e.Error
		if err != nil {
			return nil, err
		}
	}
	return &res, nil
}
