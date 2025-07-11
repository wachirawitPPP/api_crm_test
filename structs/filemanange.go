package structs

type ResSpace struct {
	All_storage     string  `json:"all_storage"`
	Use_storage     string  `json:"use_storage"`
	Per_storage_use float32 `json:"per_storage_use"`
	Last_use_date   string  `json:"last_use_date"`
	Is_upload       int     `json:"is_upload"`
}

type ObjPayloadSearchFile struct {
	Search      *string `json:"search" binding:"required,omitempty"`
	Date        *string `json:"date" binding:"required,omitempty"`
	Que_type_id *string `json:"que_type_id" binding:"required,omitempty"`
	Shop_id     int     `json:"shop_id"`
	ActivePage  int     `json:"active_page" binding:"required"`
	PerPage     int     `json:"per_page" binding:"required"`
}

type ResponsePaginationFile struct {
	Result_data   []QueueFiles `json:"result_data"`
	Count_of_page int          `json:"count_of_page"`
	Count_all     int          `json:"count_all"`
}

type QueueFiles struct {
	Id              int        `json:"id"`
	Que_code        string     `json:"que_code"`
	Queue_id        int        `json:"queue_id"`
	Quef_path       string     `json:"quef_path"`
	Quef_size       int        `json:"quef_size"`
	Quef_is_use     int        `json:"quef_is_use"`
	Quef_modify     string     `json:"quef_modify"`
	Que_create      string     `json:"que_create"`
	Que_update      string     `json:"que_update"`
	Ctm_fullname    string     `json:"ctm_fullname"`
	Ctm_fullname_en string     `json:"ctm_fullname_en"`
	Ctm_fname       string     `json:"ctm_fname"`
	Ctm_lname       string     `json:"ctm_lname"`
	Que_type_id     int        `json:"que_type_id"`
	Tags            *[]QueTags `json:"tags" gorm:"-"`
}

type ObjPayloadSearchFileXray struct {
	Search     *string `json:"search" binding:"required,omitempty"`
	Date       *string `json:"date" binding:"required,omitempty"`
	Shop_id    int     `json:"shop_id"`
	ActivePage int     `json:"active_page" binding:"required"`
	PerPage    int     `json:"per_page" binding:"required"`
}

type ResponsePaginationFileXray struct {
	Result_data   []QueueFilesXray `json:"result_data"`
	Count_of_page int              `json:"count_of_page"`
	Count_all     int              `json:"count_all"`
}

type QueueFilesXray struct {
	Id              int        `json:"id"`
	Que_code        string     `json:"que_code"`
	Queue_id        int        `json:"queue_id"`
	Chk_upload      string     `json:"chk_upload"`
	Chk_upload_size string     `json:"chk_upload_size"`
	Chk_create      string     `json:"chk_create"`
	Chk_update      string     `json:"chk_update"`
	Ctm_fullname    string     `json:"ctm_fullname"`
	Ctm_fullname_en string     `json:"ctm_fullname_en"`
	Ctm_fname       string     `json:"ctm_fname"`
	Ctm_lname       string     `json:"ctm_lname"`
	Tags            *[]QueTags `json:"tags" gorm:"-"`
}

type QueTags struct {
	Id       int    `json:"id"`
	Tag_id   int    `json:"tag_id"`
	Tag_name string `json:"tag_name"`
}
