package structs

type Patient struct {
	HN          string `json:"hn"`
	FullName    string `json:"full_name"`
	TName       string `json:"tname"`
	FName       string `json:"fname"`
	LName       string `json:"lname"`
	IDCard      string `json:"idcard"`
	Birthday    string `json:"birthday"`
	Sex         string `json:"sex"`
	DoctorName  string `json:"doctor_name"`
	RequestNote string `json:"request_note"`
}

// LabOrder represents a lab order in the request data
type LabOrder struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// Data represents the entire request data
type LabplusRequestOrder struct {
	Patient  Patient    `json:"patient"`
	LabOrder []LabOrder `json:"lab_order"`
}

type Labplus struct {
	Id                int    `json:"id"`
	Lapi_username     string `json:"lapi_username"`
	Lapi_password     string `json:"lapi_password"`
	Lapi_link         string `json:"lapi_link"`
	Lapi_is_active    int    `json:"lapi_is_active"`
	Lapi_token        string `json:"lapi_token"`
	Lapi_token_gen    string `json:"lapi_token_gen"`
	Lapi_token_expire string `json:"lapi_token_expire"`
}

type UpdateLabplus struct {
	Lapi_token        string `json:"lapi_token"`
	Lapi_token_gen    string `json:"lapi_token_gen"`
	Lapi_token_expire string `json:"lapi_token_expire"`
}

type RequestOrder struct {
	Status_code string `json:"status_code"`
	Error       string `json:"error"`
	Message     string `json:"message"`
	Message_th  string `json:"message_th"`
	Request_no  string `json:"request_no"`
}

type AddQueuesLabplus struct {
	Queue_id         int    `json:"queue_id"`
	Shop_id          int    `json:"shop_id"`
	Qlp_no           string `json:"qlp_no"`
	Qlp_message_th   string `json:"qlp_message_th"`
	Qlp_process_code string `json:"qlp_process_code"`
	Qlp_process_name string `json:"qlp_process_name"`
	Qlp_datetime     string `json:"qlp_datetime"`
	Qlp_update       string `json:"qlp_update"`
}

type QueuesLabplus struct {
	Id               int    `json:"id"`
	Queue_id         int    `json:"queue_id"`
	Shop_id          int    `json:"shop_id"`
	Qlp_no           string `json:"qlp_no"`
	Qlp_message_th   string `json:"qlp_message_th"`
	Qlp_process_code string `json:"qlp_process_code"`
	Qlp_process_name string `json:"qlp_process_name"`
	Qlp_datetime     string `json:"qlp_datetime"`
	Qlp_update       string `json:"qlp_update"`
}

type UpdateQueuesLabplus struct {
	Qlp_message_th       string `json:"qlp_message_th"`
	Qlp_process_code     string `json:"qlp_process_code"`
	Qlp_process_name     string `json:"qlp_process_name"`
	Qlp_receive_staff    string `json:"qlp_receive_staff"`
	Qlp_receive_datetime string `json:"qlp_receive_datetime"`
	Qlp_approve_staff    string `json:"qlp_approve_staff"`
	Qlp_approve_datetime string `json:"qlp_approve_datetime"`
	Qlp_update           string `json:"qlp_update"`
}

type PayloadQueuesLabplus struct {
	Id       int    `json:"id"`
	Queue_id int    `json:"queue_id"`
	Qlp_no   string `json:"qlp_no"`
}
