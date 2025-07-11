package libs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"linecrmapi/structs"
	"mime/multipart"
	"net/http"
)

type AuthResponse struct {
	Status_code string `json:"status_code"`
	Error       string `json:"error"`
	Token       string `json:"token"`
}

type RequestOrder struct {
	Status_code string `json:"status_code"`
	Error       string `json:"error"`
	Message     string `json:"message"`
	Message_th  string `json:"message_th"`
	Request_no  string `json:"request_no"`
}

type ResponseCancel struct {
	Status_code string `json:"status_code"`
	Error       string `json:"error"`
	Message     string `json:"message"`
	Message_th  string `json:"message_th"`
}

type ResponseStatus struct {
	Status_code  string `json:"status_code"`
	Error        string `json:"error"`
	Process_code string `json:"process_code"`
	Process_name string `json:"process_name"`
	Message      string `json:"message"`
	Message_th   string `json:"message_th"`
}

type DataResult struct {
	ItemCode     string `json:"item_code"`
	ItemName     string `json:"item_name"`
	Result       string `json:"result"`
	ResultNormal string `json:"result_normal"`
	ItemUnit     string `json:"item_unit"`
	Remark       string `json:"remark"`
	Flag         string `json:"flag"`
}

// DataOrder represents the structure of lab_order field in the JSON response
type DataOrder struct {
	OrderCode  string       `json:"order_code"`
	OrderName  string       `json:"order_name"`
	DataResult []DataResult `json:"data_result"`
}

// DataPatient represents the structure of data_patient field in the JSON response
type DataPatient struct {
	HN           string `json:"hn"`
	IDCard       string `json:"idcard"`
	PatientName  string `json:"patient_name"`
	Sex          string `json:"sex"`
	ReceiveStaff string `json:"receive_staff"`
	ReceiveDate  string `json:"receive_date"`
	ReceiveTime  string `json:"receive_time"`
	ApproveStaff string `json:"approve_staff"`
	ApproveDate  string `json:"approve_date"`
	ApproveTime  string `json:"approve_time"`
}

// Response represents the structure of the entire JSON response
type ResponseResult struct {
	StatusCode  string      `json:"status_code"`
	Error       string      `json:"error"`
	Message     string      `json:"message"`
	Message_th  string      `json:"message_th"`
	DataPatient DataPatient `json:"data_patient"`
	DataOrder   []DataOrder `json:"data_order"`
}

func LabplusAuthen(link, username, password string) (*AuthResponse, error) {

	url := link + "/sign_in"
	method := "POST"

	// Create a buffer to hold the multipart form data
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	// Write fields to the multipart form data
	_ = writer.WriteField("username", username)
	_ = writer.WriteField("password", password)
	err := writer.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close writer: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set the Content-Type header to the form data content type
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("User-Agent", "APSTH")

	// Perform the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer res.Body.Close()

	// Read and return the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var result AuthResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return &result, nil
}

func LabplusRequestOrder(link string, token string, data *structs.LabplusRequestOrder) (*RequestOrder, error) {
	url := link + "/lab_order/request"

	// Convert data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "APSTH")

	// Perform the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer res.Body.Close()

	// Read and parse the response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var result RequestOrder
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return &result, nil

}

func LabpluscancelOrder(link, token, requestNo string) (*ResponseCancel, error) {
	url := link + "/lab_order/cancel_request/" + requestNo

	// Create a new HTTP request
	req, err := http.NewRequest("PATCH", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("User-Agent", "APSTH")

	// Perform the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer res.Body.Close()

	// Read and parse the response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var result ResponseCancel
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return &result, nil
}

func LabplusstatusOrder(link, token, requestNo string) (*ResponseStatus, error) {
	url := link + "/lab_order/status_request/" + requestNo

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("User-Agent", "APSTH")

	// Perform the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer res.Body.Close()

	// Read and parse the response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var result ResponseStatus
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return &result, nil
}

func LabplusgetResult(link, token, requestNo string) (*ResponseResult, error) {
	url := link + "/result/" + requestNo

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("User-Agent", "APSTH")

	// Perform the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer res.Body.Close()

	// Read and parse the response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var result ResponseResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return &result, nil
}
