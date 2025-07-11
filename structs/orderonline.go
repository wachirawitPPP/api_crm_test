package structs

import (
	"linecrmapi/middlewares"
	"time"
)

type OrderOnline struct {
	Id                  int     `json:"id"`
	Or_code             string  `json:"or_code"`
	Or_website_id       int     `json:"or_website_id"`
	Or_website_name     string  `json:"or_website_name"`
	Customer_online_id  int     `json:"customer_online_id"`
	Or_fullname         string  `json:"or_fullname"`
	Or_tel              string  `json:"or_tel"`
	Or_email            string  `json:"or_email"`
	Or_address          string  `json:"or_address"`
	Or_district         string  `json:"or_district"`
	Or_amphoe           string  `json:"or_amphoe"`
	Or_province         string  `json:"or_province"`
	Or_zipcode          string  `json:"or_zipcode"`
	Or_comment          string  `json:"or_comment"`
	Or_commission_price float64 `json:"or_commission_price"`
	Or_total_price      float64 `json:"or_total_price"`
	Or_discount_type_id float64 `json:"or_discount_type_id"`
	Or_discount_item    float64 `json:"or_discount_item"`
	Or_discount_value   float64 `json:"or_discount_value"`
	Or_discount         float64 `json:"or_discount"`
	User_id_cancel      int     `json:"user_id_cancel"`
	Payment_status      int     `json:"payment_status"`
	Payment_type_id     int     `json:"payment_type_id"`
	Payment_ref         string  `json:"payment_ref"`
	Or_date             string  `json:"or_date"`
	Or_create           string  `json:"or_create"`
	Or_update           string  `json:"or_update"`
	Or_is_active        int     `json:"or_is_active"`
	Or_is_del           int     `json:"or_is_del"`
}

type OrderOnlineDetail struct {
	Id                   int     `json:"id"`
	Order_online_id      int     `json:"order_online_id"`
	Shop_id              int     `json:"shop_id"`
	Ord_type_id          int     `json:"ord_type_id"`
	Ord_item_id          int     `json:"ord_item_id"`
	Ord_code             string  `json:"ord_code"`
	Ord_name             string  `json:"ord_name"`
	Ord_qty              float64 `json:"ord_qty"`
	Ord_set_qty          float64 `json:"ord_set_qty"`
	Ord_limit_qty        float64 `json:"ord_limit_qty"`
	Ord_rate             float64 `json:"ord_rate"`
	Ord_unit             string  `json:"ord_unit"`
	Ord_cost             float64 `json:"ord_cost"`
	Ord_price            float64 `json:"ord_price"`
	Ord_discount_type_id int     `json:"ord_discount_type_id"`
	Ord_discount_item    float64 `json:"ord_discount_item"`
	Ord_discount         float64 `json:"ord_discount"`
	Ord_commission       float64 `json:"ord_commission"`
	Ord_commission_price float64 `json:"ord_commission_price"`
	Ord_amount           float64 `json:"ord_amount"`
	Ord_total            float64 `json:"ord_total"`
	Topical_id           int     `json:"topical_id"`
	Ord_topical          string  `json:"ord_topical"`
	Ord_direction        string  `json:"ord_direction"`
	Ord_create           string  `json:"ord_create"`
	Ord_update           string  `json:"ord_update"`
	Ord_is_active        int     `json:"ord_is_active"`
	Ord_is_del           int     `json:"ord_is_del"`
}

type ObjPayloadCreateOrderOnline struct {
	Or_website_id       int                                 `json:"or_website_id"`
	Or_website_name     string                              `json:"or_website_name"`
	Customer_online_id  int                                 `json:"customer_online_id"`
	Or_fullname         string                              `json:"or_fullname"`
	Or_tel              string                              `json:"or_tel"`
	Or_email            string                              `json:"or_email"`
	Or_address          string                              `json:"or_address"`
	Or_district         string                              `json:"or_district"`
	Or_amphoe           string                              `json:"or_amphoe"`
	Or_province         string                              `json:"or_province"`
	Or_zipcode          string                              `json:"or_zipcode"`
	Or_comment          string                              `json:"or_comment"`
	Or_commission_price float64                             `json:"or_commission_price"`
	Or_total_price      float64                             `json:"or_total_price"`
	Or_discount_type_id float64                             `json:"or_discount_type_id"`
	Or_discount_item    float64                             `json:"or_discount_item"`
	Or_discount_value   float64                             `json:"or_discount_value"`
	Or_discount         float64                             `json:"or_discount"`
	User_id_cancel      int                                 `json:"user_id_cancel"`
	Payment_status      int                                 `json:"payment_status"`
	Payment_type_id     int                                 `json:"payment_type_id"`
	Payment_ref         string                              `json:"payment_ref"`
	Or_date             string                              `json:"or_date"`
	Or_detail           []ObjPayloadCreateOrderOnlineDetail `json:"or_detail"`
}

type ObjPayloadCreateOrderOnlineDetail struct {
	Order_online_id      int     `json:"order_online_id"`
	Shop_id              int     `json:"shop_id"`
	Ord_type_id          int     `json:"ord_type_id"`
	Ord_item_id          int     `json:"ord_item_id"`
	Ord_code             string  `json:"ord_code"`
	Ord_name             string  `json:"ord_name"`
	Ord_qty              float64 `json:"ord_qty"`
	Ord_set_qty          float64 `json:"ord_set_qty"`
	Ord_limit_qty        float64 `json:"ord_limit_qty"`
	Ord_rate             float64 `json:"ord_rate"`
	Ord_unit             string  `json:"ord_unit"`
	Ord_cost             float64 `json:"ord_cost"`
	Ord_price            float64 `json:"ord_price"`
	Ord_discount_type_id int     `json:"ord_discount_type_id"`
	Ord_discount_item    float64 `json:"ord_discount_item"`
	Ord_discount         float64 `json:"ord_discount"`
	Ord_commission       float64 `json:"ord_commission"`
	Ord_commission_price float64 `json:"ord_commission_price"`
	Ord_amount           float64 `json:"ord_amount"`
	Ord_total            float64 `json:"ord_total"`
	Topical_id           int     `json:"topical_id"`
	Ord_topical          string  `json:"ord_topical"`
	Ord_direction        string  `json:"ord_direction"`
}

type ObjGetOrderOnline struct {
	Id                  int                       `json:"id"`
	Or_code             string                    `json:"or_code"`
	Or_website_id       int                       `json:"or_website_id"`
	Or_website_name     string                    `json:"or_website_name"`
	Customer_online_id  int                       `json:"customer_online_id"`
	Or_fullname         string                    `json:"or_fullname"`
	Or_tel              string                    `json:"or_tel"`
	Or_email            string                    `json:"or_email"`
	Or_address          string                    `json:"or_address"`
	Or_district         string                    `json:"or_district"`
	Or_amphoe           string                    `json:"or_amphoe"`
	Or_province         string                    `json:"or_province"`
	Or_zipcode          string                    `json:"or_zipcode"`
	Or_comment          string                    `json:"or_comment"`
	Or_commission_price float64                   `json:"or_commission_price"`
	Or_total_price      float64                   `json:"or_total_price"`
	Or_discount_type_id float64                   `json:"or_discount_type_id"`
	Or_discount_item    float64                   `json:"or_discount_item"`
	Or_discount_value   float64                   `json:"or_discount_value"`
	Or_discount         float64                   `json:"or_discount"`
	User_id_cancel      int                       `json:"user_id_cancel"`
	Payment_status      int                       `json:"payment_status"`
	Payment_type_id     int                       `json:"payment_type_id"`
	Payment_ref         string                    `json:"payment_ref"`
	Or_date             string                    `json:"or_date"`
	Or_create           string                    `json:"or_create"`
	Or_update           string                    `json:"or_update"`
	Or_is_active        int                       `json:"or_is_active"`
	Or_detail           []ObjGetOrderOnlineDetail `json:"or_detail"`
}

type ObjGetOrderOnlineDetail struct {
	Id                   int     `json:"id"`
	Order_online_id      int     `json:"order_online_id"`
	Shop_id              int     `json:"shop_id"`
	Ord_type_id          int     `json:"ord_type_id"`
	Ord_item_id          int     `json:"ord_item_id"`
	Ord_code             string  `json:"ord_code"`
	Ord_name             string  `json:"ord_name"`
	Ord_qty              float64 `json:"ord_qty"`
	Ord_set_qty          float64 `json:"ord_set_qty"`
	Ord_limit_qty        float64 `json:"ord_limit_qty"`
	Ord_rate             float64 `json:"ord_rate"`
	Ord_unit             string  `json:"ord_unit"`
	Ord_cost             float64 `json:"ord_cost"`
	Ord_price            float64 `json:"ord_price"`
	Ord_discount_type_id int     `json:"ord_discount_type_id"`
	Ord_discount_item    float64 `json:"ord_discount_item"`
	Ord_discount         float64 `json:"ord_discount"`
	Ord_commission       float64 `json:"ord_commission"`
	Ord_commission_price float64 `json:"ord_commission_price"`
	Ord_amount           float64 `json:"ord_amount"`
	Ord_total            float64 `json:"ord_total"`
	Topical_id           int     `json:"topical_id"`
	Ord_topical          string  `json:"ord_topical"`
	Ord_direction        string  `json:"ord_direction"`
	Ord_create           string  `json:"ord_create"`
	Ord_update           string  `json:"ord_update"`
	Ord_is_active        int     `json:"ord_is_active"`
	Ord_is_del           int     `json:"ord_is_del"`
}

type ObjPayloadSearchOrderOnline struct {
	Customer_online_id int    `json:"customer_online_id"`
	Order_online_id    int    `json:"order_online_id"`
	Search_text        string `json:"search_text"`
	Or_website_id      int    `json:"or_website_id"`
	Payment_status     int    `json:"payment_status"`
	Or_date            string `json:"or_date"`
	Or_is_active       int    `json:"or_is_active"`
	Active_page        int    `json:"active_page" binding:"required"`
	Per_page           int    `json:"per_page" binding:"required"`
}

type ObjPaginationOrderOnline struct {
	Result_data   []ObjGetOrderOnline `json:"result_data"`
	Count_of_page int                 `json:"count_of_page"`
	Count_all     int                 `json:"count_all"`
}

type UpdateOrderOnline struct {
	Id                  int     `json:"id"`
	Or_fullname         string  `json:"or_fullname"`
	Or_tel              string  `json:"or_tel"`
	Or_email            string  `json:"or_email"`
	Or_address          string  `json:"or_address"`
	Or_district         string  `json:"or_district"`
	Or_amphoe           string  `json:"or_amphoe"`
	Or_province         string  `json:"or_province"`
	Or_zipcode          string  `json:"or_zipcode"`
	Or_comment          string  `json:"or_comment"`
	Or_commission_price float64 `json:"or_commission_price"`
	Or_total_price      float64 `json:"or_total_price"`
	Or_discount_type_id float64 `json:"or_discount_type_id"`
	Or_discount_item    float64 `json:"or_discount_item"`
	Or_discount_value   float64 `json:"or_discount_value"`
	Or_discount         float64 `json:"or_discount"`
	User_id_cancel      int     `json:"user_id_cancel"`
	Payment_status      int     `json:"payment_status"`
	Payment_type_id     int     `json:"payment_type_id"`
	Payment_ref         string  `json:"payment_ref"`
	Or_date             string  `json:"or_date"`
	Or_update           string  `json:"or_update"`
	Or_is_active        int     `json:"or_is_active"`
}

type UpdateOrderOnlineDetail struct {
	Id                   int     `json:"id"`
	Ord_code             string  `json:"ord_code"`
	Ord_name             string  `json:"ord_name"`
	Ord_qty              float64 `json:"ord_qty"`
	Ord_set_qty          float64 `json:"ord_set_qty"`
	Ord_limit_qty        float64 `json:"ord_limit_qty"`
	Ord_rate             float64 `json:"ord_rate"`
	Ord_unit             string  `json:"ord_unit"`
	Ord_cost             float64 `json:"ord_cost"`
	Ord_price            float64 `json:"ord_price"`
	Ord_discount_type_id int     `json:"ord_discount_type_id"`
	Ord_discount_item    float64 `json:"ord_discount_item"`
	Ord_discount         float64 `json:"ord_discount"`
	Ord_commission       float64 `json:"ord_commission"`
	Ord_commission_price float64 `json:"ord_commission_price"`
	Ord_amount           float64 `json:"ord_amount"`
	Ord_total            float64 `json:"ord_total"`
	Topical_id           int     `json:"topical_id"`
	Ord_topical          string  `json:"ord_topical"`
	Ord_direction        string  `json:"ord_direction"`
	Ord_update           string  `json:"ord_update"`
	Ord_is_active        int     `json:"ord_is_active"`
}

type ObjPaymentOrderOnline struct {
	Order_online_id int    `json:"order_online_id"`
	Payment_status  int    `json:"payment_status"`
	Payment_type_id int    `json:"payment_type_id"`
	Payment_ref     string `json:"payment_ref"`
}

type ObjCancelOrderOnline struct {
	Order_online_id int `json:"order_online_id"`
	User_id_cancel  int `json:"user_id_cancel"`
}

func ConvertPayloadToOrderOnline(payload ObjPayloadCreateOrderOnline) OrderOnline {
	Or_code := "ORO_" + middlewares.GenerateDateTimeCode()
	return OrderOnline{
		Or_code:             Or_code,
		Or_website_id:       payload.Or_website_id,
		Or_website_name:     payload.Or_website_name,
		Customer_online_id:  payload.Customer_online_id,
		Or_fullname:         payload.Or_fullname,
		Or_tel:              payload.Or_tel,
		Or_email:            payload.Or_email,
		Or_address:          payload.Or_address,
		Or_district:         payload.Or_district,
		Or_amphoe:           payload.Or_amphoe,
		Or_province:         payload.Or_province,
		Or_zipcode:          payload.Or_zipcode,
		Or_comment:          payload.Or_comment,
		Or_commission_price: payload.Or_commission_price,
		Or_total_price:      payload.Or_total_price,
		Or_discount_type_id: payload.Or_discount_type_id,
		Or_discount_item:    payload.Or_discount_item,
		Or_discount_value:   payload.Or_discount_value,
		Or_discount:         payload.Or_discount,
		User_id_cancel:      payload.User_id_cancel,
		Payment_status:      payload.Payment_status,
		Payment_type_id:     payload.Payment_type_id,
		Payment_ref:         payload.Payment_ref,
		Or_date:             payload.Or_date,
		Or_is_active:        1,
		Or_is_del:           0,
		Or_create:           time.Now().Format("2006-01-02 15:04:05"),
		Or_update:           time.Now().Format("2006-01-02 15:04:05"),
	}
}

func ConvertPayloadDetailToOrderDetail(payload ObjPayloadCreateOrderOnlineDetail, orderID int) OrderOnlineDetail {
	return OrderOnlineDetail{
		Order_online_id:      orderID,
		Shop_id:              payload.Shop_id,
		Ord_type_id:          payload.Ord_type_id,
		Ord_item_id:          payload.Ord_item_id,
		Ord_code:             payload.Ord_code,
		Ord_name:             payload.Ord_name,
		Ord_qty:              payload.Ord_qty,
		Ord_set_qty:          payload.Ord_set_qty,
		Ord_limit_qty:        payload.Ord_limit_qty,
		Ord_rate:             payload.Ord_rate,
		Ord_unit:             payload.Ord_unit,
		Ord_cost:             payload.Ord_cost,
		Ord_price:            payload.Ord_price,
		Ord_discount_type_id: payload.Ord_discount_type_id,
		Ord_discount_item:    payload.Ord_discount_item,
		Ord_discount:         payload.Ord_discount,
		Ord_commission:       payload.Ord_commission,
		Ord_commission_price: payload.Ord_commission_price,
		Ord_amount:           payload.Ord_amount,
		Ord_total:            payload.Ord_total,
		Topical_id:           payload.Topical_id,
		Ord_topical:          payload.Ord_topical,
		Ord_direction:        payload.Ord_direction,
		Ord_is_active:        1,
		Ord_is_del:           0,
		Ord_create:           time.Now().Format("2006-01-02 15:04:05"),
		Ord_update:           time.Now().Format("2006-01-02 15:04:05"),
	}
}

func ConvertToObjGetOrderOnline(order OrderOnline, details []OrderOnlineDetail) ObjGetOrderOnline {
	return ObjGetOrderOnline{
		Id:                  order.Id,
		Or_code:             order.Or_code,
		Or_website_id:       order.Or_website_id,
		Or_website_name:     order.Or_website_name,
		Customer_online_id:  order.Customer_online_id,
		Or_fullname:         order.Or_fullname,
		Or_tel:              order.Or_tel,
		Or_email:            order.Or_email,
		Or_address:          order.Or_address,
		Or_district:         order.Or_district,
		Or_amphoe:           order.Or_amphoe,
		Or_province:         order.Or_province,
		Or_zipcode:          order.Or_zipcode,
		Or_comment:          order.Or_comment,
		Or_commission_price: order.Or_commission_price,
		Or_total_price:      order.Or_total_price,
		Or_discount_type_id: order.Or_discount_type_id,
		Or_discount_item:    order.Or_discount_item,
		Or_discount_value:   order.Or_discount_value,
		Or_discount:         order.Or_discount,
		User_id_cancel:      order.User_id_cancel,
		Payment_status:      order.Payment_status,
		Payment_type_id:     order.Payment_type_id,
		Payment_ref:         order.Payment_ref,
		Or_date:             order.Or_date,
		Or_create:           order.Or_create,
		Or_update:           order.Or_update,
		Or_is_active:        order.Or_is_active,
		Or_detail:           ConvertToObjGetOrderOnlineDetails(details),
	}
}

func ConvertToObjGetOrderOnlineDetails(details []OrderOnlineDetail) []ObjGetOrderOnlineDetail {
	var result []ObjGetOrderOnlineDetail
	for _, d := range details {
		result = append(result, ObjGetOrderOnlineDetail{
			Id:                   d.Id,
			Order_online_id:      d.Order_online_id,
			Shop_id:              d.Shop_id,
			Ord_type_id:          d.Ord_type_id,
			Ord_item_id:          d.Ord_item_id,
			Ord_code:             d.Ord_code,
			Ord_name:             d.Ord_name,
			Ord_qty:              d.Ord_qty,
			Ord_set_qty:          d.Ord_set_qty,
			Ord_limit_qty:        d.Ord_limit_qty,
			Ord_rate:             d.Ord_rate,
			Ord_unit:             d.Ord_unit,
			Ord_cost:             d.Ord_cost,
			Ord_price:            d.Ord_price,
			Ord_discount_type_id: d.Ord_discount_type_id,
			Ord_discount_item:    d.Ord_discount_item,
			Ord_discount:         d.Ord_discount,
			Ord_commission:       d.Ord_commission,
			Ord_commission_price: d.Ord_commission_price,
			Ord_amount:           d.Ord_amount,
			Ord_total:            d.Ord_total,
			Topical_id:           d.Topical_id,
			Ord_topical:          d.Ord_topical,
			Ord_direction:        d.Ord_direction,
			Ord_create:           d.Ord_create,
			Ord_update:           d.Ord_update,
			Ord_is_active:        d.Ord_is_active,
			Ord_is_del:           d.Ord_is_del,
		})
	}
	return result
}

func ConvertOrderToUpdateOrder(order OrderOnline) UpdateOrderOnline {
	return UpdateOrderOnline{
		Id:                  order.Id,
		Or_fullname:         order.Or_fullname,
		Or_tel:              order.Or_tel,
		Or_email:            order.Or_email,
		Or_address:          order.Or_address,
		Or_district:         order.Or_district,
		Or_amphoe:           order.Or_amphoe,
		Or_province:         order.Or_province,
		Or_zipcode:          order.Or_zipcode,
		Or_comment:          order.Or_comment,
		Or_commission_price: order.Or_commission_price,
		Or_total_price:      order.Or_total_price,
		Or_discount_type_id: order.Or_discount_type_id,
		Or_discount_item:    order.Or_discount_item,
		Or_discount_value:   order.Or_discount_value,
		Or_discount:         order.Or_discount,
		User_id_cancel:      order.User_id_cancel,
		Payment_status:      order.Payment_status,
		Payment_type_id:     order.Payment_type_id,
		Payment_ref:         order.Payment_ref,
		Or_date:             order.Or_date,
		Or_update:           "",
		Or_is_active:        order.Or_is_active,
	}
}

func ConvertOrderDetailsToUpdateDetails(details []OrderOnlineDetail) []UpdateOrderOnlineDetail {
	var result []UpdateOrderOnlineDetail
	for _, d := range details {
		result = append(result, UpdateOrderOnlineDetail{
			Id:                   d.Id,
			Ord_code:             d.Ord_code,
			Ord_name:             d.Ord_name,
			Ord_qty:              d.Ord_qty,
			Ord_set_qty:          d.Ord_set_qty,
			Ord_limit_qty:        d.Ord_limit_qty,
			Ord_rate:             d.Ord_rate,
			Ord_unit:             d.Ord_unit,
			Ord_cost:             d.Ord_cost,
			Ord_price:            d.Ord_price,
			Ord_discount_type_id: d.Ord_discount_type_id,
			Ord_discount_item:    d.Ord_discount_item,
			Ord_discount:         d.Ord_discount,
			Ord_commission:       d.Ord_commission,
			Ord_commission_price: d.Ord_commission_price,
			Ord_amount:           d.Ord_amount,
			Ord_total:            d.Ord_total,
			Topical_id:           d.Topical_id,
			Ord_topical:          d.Ord_topical,
			Ord_direction:        d.Ord_direction,
			Ord_update:           d.Ord_update,
			Ord_is_active:        d.Ord_is_active,
		})
	}
	return result
}

func ConvertObjGetToUpdateOrder(obj ObjGetOrderOnline) UpdateOrderOnline {
	return UpdateOrderOnline{
		Id:                  obj.Id,
		Or_fullname:         obj.Or_fullname,
		Or_tel:              obj.Or_tel,
		Or_email:            obj.Or_email,
		Or_address:          obj.Or_address,
		Or_district:         obj.Or_district,
		Or_amphoe:           obj.Or_amphoe,
		Or_province:         obj.Or_province,
		Or_zipcode:          obj.Or_zipcode,
		Or_comment:          obj.Or_comment,
		Or_commission_price: obj.Or_commission_price,
		Or_total_price:      obj.Or_total_price,
		Or_discount_type_id: obj.Or_discount_type_id,
		Or_discount_item:    obj.Or_discount_item,
		Or_discount_value:   obj.Or_discount_value,
		Or_discount:         obj.Or_discount,
		User_id_cancel:      obj.User_id_cancel,
		Payment_status:      obj.Payment_status,
		Payment_type_id:     obj.Payment_type_id,
		Payment_ref:         obj.Payment_ref,
		Or_date:             obj.Or_date,
		Or_update:           "",
		Or_is_active:        obj.Or_is_active,
	}
}
func ConvertObjGetDetailToOrderDetail(d ObjGetOrderOnlineDetail) OrderOnlineDetail {
	return OrderOnlineDetail{
		Id:                   d.Id,
		Order_online_id:      d.Order_online_id,
		Shop_id:              d.Shop_id,
		Ord_type_id:          d.Ord_type_id,
		Ord_item_id:          d.Ord_item_id,
		Ord_code:             d.Ord_code,
		Ord_name:             d.Ord_name,
		Ord_qty:              d.Ord_qty,
		Ord_set_qty:          d.Ord_set_qty,
		Ord_limit_qty:        d.Ord_limit_qty,
		Ord_rate:             d.Ord_rate,
		Ord_unit:             d.Ord_unit,
		Ord_cost:             d.Ord_cost,
		Ord_price:            d.Ord_price,
		Ord_discount_type_id: d.Ord_discount_type_id,
		Ord_discount_item:    d.Ord_discount_item,
		Ord_discount:         d.Ord_discount,
		Ord_commission:       d.Ord_commission,
		Ord_commission_price: d.Ord_commission_price,
		Ord_amount:           d.Ord_amount,
		Ord_total:            d.Ord_total,
		Topical_id:           d.Topical_id,
		Ord_topical:          d.Ord_topical,
		Ord_direction:        d.Ord_direction,
		Ord_create:           d.Ord_create,
		Ord_update:           d.Ord_update,
		Ord_is_active:        d.Ord_is_active,
		Ord_is_del:           d.Ord_is_del,
	}
}

func ConvertObjGetDetailsToUpdateDetails(details []ObjGetOrderOnlineDetail) ([]UpdateOrderOnlineDetail, []OrderOnlineDetail) {
	var resultUpdate []UpdateOrderOnlineDetail
	var resultCreate []OrderOnlineDetail
	for _, d := range details {
		if d.Id > 0 {
			resultUpdate = append(resultUpdate, UpdateOrderOnlineDetail{
				Id:                   d.Id,
				Ord_code:             d.Ord_code,
				Ord_name:             d.Ord_name,
				Ord_qty:              d.Ord_qty,
				Ord_set_qty:          d.Ord_set_qty,
				Ord_limit_qty:        d.Ord_limit_qty,
				Ord_rate:             d.Ord_rate,
				Ord_unit:             d.Ord_unit,
				Ord_cost:             d.Ord_cost,
				Ord_price:            d.Ord_price,
				Ord_discount_type_id: d.Ord_discount_type_id,
				Ord_discount_item:    d.Ord_discount_item,
				Ord_discount:         d.Ord_discount,
				Ord_commission:       d.Ord_commission,
				Ord_commission_price: d.Ord_commission_price,
				Ord_amount:           d.Ord_amount,
				Ord_total:            d.Ord_total,
				Topical_id:           d.Topical_id,
				Ord_topical:          d.Ord_topical,
				Ord_direction:        d.Ord_direction,
				Ord_update:           time.Now().Format("2006-01-02 15:04:05"),
				Ord_is_active:        d.Ord_is_active,
			})
		} else {
			resultCreate = append(resultCreate, ConvertObjGetDetailToOrderDetail(d))
		}

	}
	return resultUpdate, resultCreate
}

func ConvertUpdateDetailToOrderDetail(d UpdateOrderOnlineDetail) OrderOnlineDetail {
	return OrderOnlineDetail{
		Id:                   d.Id,
		Ord_code:             d.Ord_code,
		Ord_name:             d.Ord_name,
		Ord_qty:              d.Ord_qty,
		Ord_set_qty:          d.Ord_set_qty,
		Ord_limit_qty:        d.Ord_limit_qty,
		Ord_rate:             d.Ord_rate,
		Ord_unit:             d.Ord_unit,
		Ord_cost:             d.Ord_cost,
		Ord_price:            d.Ord_price,
		Ord_discount_type_id: d.Ord_discount_type_id,
		Ord_discount_item:    d.Ord_discount_item,
		Ord_discount:         d.Ord_discount,
		Ord_commission:       d.Ord_commission,
		Ord_commission_price: d.Ord_commission_price,
		Ord_amount:           d.Ord_amount,
		Ord_total:            d.Ord_total,
		Topical_id:           d.Topical_id,
		Ord_topical:          d.Ord_topical,
		Ord_direction:        d.Ord_direction,
		Ord_update:           d.Ord_update,
		Ord_is_active:        d.Ord_is_active,
	}
}

func ConvertToUpdateOrderOnlineDetail(o OrderOnlineDetail) UpdateOrderOnlineDetail {
	return UpdateOrderOnlineDetail{
		Id:                   o.Id,
		Ord_code:             o.Ord_code,
		Ord_name:             o.Ord_name,
		Ord_qty:              o.Ord_qty,
		Ord_set_qty:          o.Ord_set_qty,
		Ord_limit_qty:        o.Ord_limit_qty,
		Ord_rate:             o.Ord_rate,
		Ord_unit:             o.Ord_unit,
		Ord_cost:             o.Ord_cost,
		Ord_price:            o.Ord_price,
		Ord_discount_type_id: o.Ord_discount_type_id,
		Ord_discount_item:    o.Ord_discount_item,
		Ord_discount:         o.Ord_discount,
		Ord_commission:       o.Ord_commission,
		Ord_commission_price: o.Ord_commission_price,
		Ord_amount:           o.Ord_amount,
		Ord_total:            o.Ord_total,
		Topical_id:           o.Topical_id,
		Ord_topical:          o.Ord_topical,
		Ord_direction:        o.Ord_direction,
		Ord_update:           "",
		Ord_is_active:        o.Ord_is_active,
	}
}
