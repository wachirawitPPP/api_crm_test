package models

import (
	"encoding/json"
	"linecrmapi/configs"
	"linecrmapi/structs"
	"time"

	"gorm.io/gorm"
)

func AddLogOrder(log *structs.LogOrders) (err error) {
	query := configs.DBL1.Table("log_orders").Create(&log)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func SearchOrderCustomer(objPayload *structs.ObjPayloadSearchOrderCustomer, objQuery *[]structs.ObjResponseSearchOrderCustomer) error {
	query := configs.DB1.Table("customers")
	query = query.Select("customers.*, customer_groups.cg_name, customer_groups.cg_save_type, customer_groups.cg_save, ref_right_treatments.rt_code, ref_right_treatments.rt_name")
	query = query.Joins("LEFT JOIN customer_groups ON customer_groups.id = customers.customer_group_id")
	query = query.Joins("LEFT JOIN ref_right_treatments ON ref_right_treatments.id = customers.right_treatment_id")
	query = query.Where("customers.shop_mother_id = ?", objPayload.ShopId)
	query = query.Where("customers.ctm_is_active = ?", 1)

	// Optimize search
	searchPattern := "%" + *objPayload.SearchText + "%"
	query = query.Where("customers.ctm_id LIKE ? OR customers.ctm_fname LIKE ? OR customers.ctm_lname LIKE ? OR customers.ctm_nname LIKE ? OR customers.ctm_fname_en LIKE ? OR customers.ctm_lname_en LIKE ? OR customers.ctm_citizen_id LIKE ? OR customers.ctm_passport_id LIKE ? OR CONCAT(customers.ctm_fname, ' ', customers.ctm_lname) LIKE ?",
		searchPattern, searchPattern, searchPattern, searchPattern, searchPattern, searchPattern, searchPattern, searchPattern, searchPattern)

	query = query.Order("customers.ctm_fname ASC")
	query = query.Limit(30)

	// Use prepared statement
	stmt := query.Session(&gorm.Session{PrepareStmt: true}).Find(&objQuery)
	if stmt.Error != nil {
		return stmt.Error
	}
	return nil
}

func SearchOrderCustomerInArray(objPayload *structs.ObjPayloadSearchOrderCustomer, ShopId []int, objQuery *[]structs.ObjResponseSearchOrderCustomer) error {
	query := configs.DB1.Table("customers")
	query = query.Select("customers.*, customer_groups.cg_name, customer_groups.cg_save_type, customer_groups.cg_save, ref_right_treatments.rt_code, ref_right_treatments.rt_name")
	query = query.Joins("LEFT JOIN customer_groups ON customer_groups.id = customers.customer_group_id")
	query = query.Joins("LEFT JOIN ref_right_treatments ON ref_right_treatments.id = customers.right_treatment_id")
	query = query.Where("customers.shop_id IN ?", ShopId)
	query = query.Where("customers.shop_mother_id = ?", objPayload.ShopId)
	query = query.Where("customers.ctm_is_active = ?", 1)

	// Optimize search
	searchPattern := "%" + *objPayload.SearchText + "%"
	query = query.Where("customers.ctm_id LIKE ? OR customers.ctm_fname LIKE ? OR customers.ctm_lname LIKE ? OR customers.ctm_nname LIKE ? OR customers.ctm_fname_en LIKE ? OR customers.ctm_lname_en LIKE ? OR customers.ctm_citizen_id LIKE ? OR customers.ctm_passport_id LIKE ? OR CONCAT(customers.ctm_fname, ' ', customers.ctm_lname) LIKE ?",
		searchPattern, searchPattern, searchPattern, searchPattern, searchPattern, searchPattern, searchPattern, searchPattern, searchPattern)

	query = query.Order("LENGTH(customers.ctm_id), customers.ctm_id")
	query = query.Limit(30)

	// Use prepared statement
	stmt := query.Session(&gorm.Session{PrepareStmt: true}).Find(&objQuery)
	if stmt.Error != nil {
		return stmt.Error
	}
	return nil
}

func GetCartsList(filter structs.ObjPayloadSearchCart, isCount bool, dt *[]structs.Carts) error {
	// Build base query with table selection
	query := configs.DB1.Table("order_detail_temps")

	// Select fields based on count flag
	if isCount {
		query = query.Select(`
			order_detail_temps.*, 
			rooms.room_code, rooms.room_th, rooms.room_en, 
			room_types.room_type_th, room_types.room_type_en, 
			courses.course_code, courses.course_name, courses.course_unit, 
			checkings.checking_code, checkings.checking_name, checkings.checking_unit, 
			customers.ctm_fname, customers.ctm_lname, 
			products.pd_code, products.pd_name, 
			ref_units.u_name
		`)
	} else {
		query = query.Select("order_detail_temps.id")
	}

	// Add required conditions
	query = query.Where("order_detail_temps.customer_id = ?", filter.Customer_id).
		Where("order_detail_temps.shop_id = ?", filter.Shop_id).
		Where("order_detail_temps.ordt_is_active = 1")

	// Add search condition if present
	if *filter.Search != "" {
		searchPattern := "%" + *filter.Search + "%"
		query = query.Where("products.pd_code LIKE ? OR products.pd_name LIKE ?",
			searchPattern, searchPattern)
	}

	// Add all joins
	query = query.
		Joins("LEFT JOIN rooms ON order_detail_temps.room_id = rooms.id").
		Joins("LEFT JOIN room_types ON rooms.room_type_id = room_types.id").
		Joins("LEFT JOIN courses ON order_detail_temps.course_id = courses.id").
		Joins("LEFT JOIN checkings ON order_detail_temps.checking_id = checkings.id").
		Joins("INNER JOIN customers ON order_detail_temps.customer_id = customers.id").
		Joins("INNER JOIN products ON order_detail_temps.product_id = products.id").
		Joins("INNER JOIN product_units ON order_detail_temps.product_unit_id = product_units.id").
		Joins("INNER JOIN ref_units ON product_units.unit_id = ref_units.id")

	// Add pagination if counting
	if isCount {
		offset := filter.ActivePage * filter.PerPage
		query = query.Limit(filter.PerPage).Offset(offset)
	}

	return query.Scan(&dt).Error
}

func GetOrderList(filter structs.ObjPayloadSearchOrder, isCount bool, dt *[]structs.OrderLists) error {
	// Build base query with table selection
	query := configs.DB1.Table("orders")

	// Select fields based on count flag
	if isCount {
		query = query.Select(`
			orders.*, 
			customers.shop_id AS ctm_shop_id, 
			shops.shop_name AS ctm_shop_name, 
			users.user_fullname, users.user_fullname_en, 
			customers.ctm_id, customers.ctm_fname, customers.ctm_lname, 
			customers.ctm_fname_en, customers.ctm_lname_en, 
			queues.que_code
		`)
	} else {
		query = query.Select("orders.id")
	}

	// Add required conditions
	query = query.Where("orders.shop_id = ?", filter.Shop_id)

	// Add optional conditions
	if *filter.Customer_id != "" {
		query = query.Where("orders.customer_id = ?", filter.Customer_id)
	}
	if *filter.Or_is_active != "" {
		query = query.Where("orders.or_is_active = ?", *filter.Or_is_active)
	}
	if *filter.Or_datetime != "" {
		query = query.Where("DATE(orders.or_datetime) = ?", *filter.Or_datetime)
	}

	// Add search condition if present
	if *filter.Search != "" {
		searchPattern := "%" + *filter.Search + "%"
		query = query.Where(
			"queues.que_code LIKE ? OR "+
				"customers.ctm_id LIKE ? OR "+
				"orders.or_fullname LIKE ? OR "+
				"CONCAT(customers.ctm_fname,' ',customers.ctm_lname) LIKE ? OR "+
				"customers.ctm_fname LIKE ? OR "+
				"customers.ctm_lname LIKE ? OR "+
				"customers.ctm_fname_en LIKE ? OR "+
				"customers.ctm_lname_en LIKE ?",
			searchPattern, searchPattern, searchPattern, searchPattern,
			searchPattern, searchPattern, searchPattern, searchPattern,
		)
	}

	// Add joins
	query = query.
		Joins("INNER JOIN customers ON orders.customer_id = customers.id").
		Joins("INNER JOIN users ON orders.user_id = users.id").
		Joins("LEFT JOIN queues ON orders.queue_id = queues.id").
		Joins("LEFT JOIN shops ON customers.shop_id = shops.id")

	// Add pagination if counting
	if isCount {
		offset := filter.ActivePage * filter.PerPage
		query = query.Limit(filter.PerPage).Offset(offset)
	}

	// Add ordering
	query = query.
		Order("orders.id DESC").
		Order("orders.or_create DESC").
		Order("orders.or_is_active DESC")

	return query.Find(&dt).Error
}

func GetOrderDetail(orderId int, dt *structs.OrderDetail) (err error) {
	query := configs.DB1.Table("orders")
	query = query.Select("orders.*, users.user_fullname, users.user_fullname_en, customers.ctm_fname, customers.ctm_lname, customers.ctm_fname_en, customers.ctm_lname_en, queues.que_code")

	query = query.Where("orders.id = ?", orderId)
	// query = query.Where("orders.or_is_active = 1")
	query = query.Joins("INNER JOIN customers ON orders.customer_id = customers.id")
	query = query.Joins("INNER JOIN users ON orders.user_id = users.id")
	query = query.Joins("LEFT JOIN queues ON orders.queue_id = queues.id")

	if err = query.Scan(&dt).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderIDDetail(orderId int, shopId int, dt *structs.OrderDetailData) (err error) {
	query := configs.DB1.Table("orders")
	query = query.Select("orders.*, users.user_fullname, users.user_fullname_en, customers.ctm_fname, customers.ctm_lname, customers.ctm_fname_en, customers.ctm_lname_en, queues.que_code")

	query = query.Where("orders.id = ?", orderId)
	query = query.Where("orders.shop_id = ?", shopId)
	// query = query.Where("orders.or_is_active = 1")
	query = query.Joins("INNER JOIN customers ON orders.customer_id = customers.id")
	query = query.Joins("INNER JOIN users ON orders.user_id = users.id")
	query = query.Joins("LEFT JOIN queues ON orders.queue_id = queues.id")

	if err = query.Scan(&dt).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderSub(orderId int, dt *[]structs.OrderSub) (err error) {
	query := configs.DB1.Table("order_details")
	query = query.Select("order_details.*, rooms.room_code, rooms.room_th, rooms.room_en, room_types.room_type_th, room_types.room_type_en, courses.course_code, courses.course_name, courses.course_unit, checkings.checking_code, checkings.checking_name, checkings.checking_unit, products.pd_code, products.pd_name, ref_units.u_name, ref_units.u_name_en")
	query = query.Where("order_details.ord_is_active = 1")
	query = query.Where("order_details.order_id = ?", orderId)

	query = query.Joins("LEFT JOIN rooms ON order_details.room_id = rooms.id")
	query = query.Joins("LEFT JOIN room_types ON rooms.room_type_id = room_types.id")
	query = query.Joins("LEFT JOIN courses ON order_details.course_id = courses.id")
	query = query.Joins("LEFT JOIN checkings ON order_details.checking_id = checkings.id")
	query = query.Joins("LEFT JOIN products ON order_details.product_id = products.id")
	query = query.Joins("LEFT JOIN product_units ON order_details.product_unit_id = product_units.id")
	query = query.Joins("LEFT JOIN ref_units ON product_units.unit_id = ref_units.id")

	if err = query.Scan(&dt).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderSubId(subId int, dt *structs.OrderSub) (err error) {
	query := configs.DB1.Table("order_details")
	query = query.Select("order_details.*, rooms.room_code, rooms.room_th, rooms.room_en, room_types.room_type_th, room_types.room_type_en, courses.course_code, courses.course_name, courses.course_unit, checkings.checking_code, checkings.checking_name, checkings.checking_unit, products.pd_code, products.pd_name, ref_units.u_name")
	query = query.Where("order_details.ord_is_active = 1")
	query = query.Where("order_details.id = ?", subId)

	query = query.Joins("LEFT JOIN rooms ON order_details.room_id = rooms.id")
	query = query.Joins("LEFT JOIN room_types ON rooms.room_type_id = room_types.id")
	query = query.Joins("LEFT JOIN courses ON order_details.course_id = courses.id")
	query = query.Joins("LEFT JOIN checkings ON order_details.checking_id = checkings.id")
	query = query.Joins("LEFT JOIN products ON order_details.product_id = products.id")
	query = query.Joins("LEFT JOIN product_units ON order_details.product_unit_id = product_units.id")
	query = query.Joins("LEFT JOIN ref_units ON product_units.unit_id = ref_units.id")

	if err = query.Scan(&dt).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderTags(orderId int, dt *[]structs.OrderTags) (err error) {
	query := configs.DB1.Table("order_tags")
	query = query.Select("order_tags.*, tags.tag_name")
	query = query.Where("order_tags.order_id = ?", orderId)

	query = query.Joins("INNER JOIN tags ON order_tags.tags_id = tags.id")

	if err = query.Scan(&dt).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderTagId(tagId int, dt *structs.OrderTags) (err error) {
	query := configs.DB1.Table("order_tags")
	query = query.Select("order_tags.*, tags.tag_name")
	query = query.Where("order_tags.id = ?", tagId)

	query = query.Joins("INNER JOIN tags ON order_tags.tags_id = tags.id")

	if err = query.Scan(&dt).Error; err != nil {
		return err
	}
	return nil
}

func AddOrder(dataH *structs.OrderDetail) (id int, err error) {
	tx := configs.DB1.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return 0, err
	}

	// add checkings
	objH := structs.OrderAction{
		Id:                  0,
		Shop_id:             dataH.Shop_id,
		User_id:             dataH.User_id,
		Customer_id:         dataH.Customer_id,
		Customer_online_id:  dataH.Customer_online_id,
		Queue_id:            dataH.Queue_id,
		Or_fullname:         dataH.Or_fullname,
		Or_tel:              dataH.Or_tel,
		Or_email:            dataH.Or_email,
		Or_address:          dataH.Or_address,
		Or_district:         dataH.Or_district,
		Or_amphoe:           dataH.Or_amphoe,
		Or_province:         dataH.Or_province,
		Or_zipcode:          dataH.Or_zipcode,
		Or_comment:          dataH.Or_comment,
		Or_total_price:      dataH.Or_total_price,
		Or_discount_type_id: dataH.Or_discount_type_id,
		Or_discount_item:    dataH.Or_discount_item,
		Or_discount_value:   dataH.Or_discount_value,
		Or_discount:         dataH.Or_discount,
		Or_befor_vat:        dataH.Or_befor_vat,
		Tax_type_id:         dataH.Tax_type_id,
		Tax_rate:            dataH.Tax_rate,
		Or_vat:              dataH.Or_vat,
		Or_total:            dataH.Or_total,
		Or_is_active:        dataH.Or_is_active,
		Or_datetime:         dataH.Or_datetime,
		Or_create:           dataH.Or_create,
		Or_update:           dataH.Or_update,
		DpmId:               dataH.DpmId,
		Or_eclaim_id:        dataH.Or_eclaim_id,
		Or_eclaim_rate:      dataH.Or_eclaim_rate,
		Or_eclaim_over:      dataH.Or_eclaim_over,
		Or_eclaim_total:     dataH.Or_eclaim_total,
		Or_tele_code:        dataH.Or_tele_code,
	}

	if err = tx.Table("orders").Create(&objH).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	if dataH.Queue_id > 0 {
		if err = tx.Table("orders").Where("orders.id = ?", objH.Id).Updates(map[string]interface{}{"orders.queue_id": dataH.Queue_id}).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	//add subs
	for _, sub := range *dataH.Subs {
		var Queue_checking_id *int = nil
		var Queue_course_id *int = nil
		var Queue_product_id *int = nil
		if sub.Queue_id != nil {
			if sub.Ord_type_id == 1 {
				Queue_checking_id = sub.Queue_ord_id
			} else if sub.Ord_type_id == 2 {
				Queue_course_id = sub.Queue_ord_id
			} else if sub.Ord_type_id == 3 {
				Queue_product_id = sub.Queue_ord_id
			} else if sub.Ord_type_id == 5 {
				Queue_product_id = sub.Queue_ord_id
			}
		}
		objSub := structs.OrderSub{
			Id:                   0,
			Order_id:             objH.Id,
			Course_id:            sub.Course_id,
			Checking_id:          sub.Checking_id,
			Product_id:           sub.Product_id,
			Product_store_id:     sub.Product_store_id,
			Product_unit_id:      sub.Product_unit_id,
			Coin_id:              sub.Coin_id,
			Room_id:              sub.Room_id,
			Queue_id:             &dataH.Queue_id,
			Queue_ord_id:         sub.Queue_ord_id,
			Queue_checking_id:    Queue_checking_id,
			Queue_course_id:      Queue_course_id,
			Queue_product_id:     Queue_product_id,
			Ord_type_id:          sub.Ord_type_id,
			Ord_code:             sub.Ord_code,
			Ord_name:             sub.Ord_name,
			Ord_qty:              sub.Ord_qty,
			Ord_rate:             sub.Ord_rate,
			Ord_set_qty:          sub.Ord_set_qty,
			Ord_limit_qty:        sub.Ord_limit_qty,
			Ord_unit:             sub.Ord_unit,
			Ord_cost:             sub.Ord_cost,
			Ord_price:            sub.Ord_price,
			Ord_amount:           sub.Ord_amount,
			Ord_discount_type_id: sub.Ord_discount_type_id,
			Ord_discount_item:    sub.Ord_discount_item,
			Ord_discount:         sub.Ord_discount,
			Tax_type_id:          sub.Tax_type_id,
			Tax_rate:             sub.Tax_rate,
			Ord_vat:              sub.Ord_vat,
			Ord_total:            sub.Ord_total,
			Topical_id:           sub.Topical_id,
			Ord_topical:          sub.Ord_topical,
			Ord_direction:        sub.Ord_direction,
			Ord_is_set:           sub.Ord_is_set,
			Ord_id_set:           sub.Ord_id_set,
			Ord_is_use:           sub.Ord_is_use,
			Ord_is_active:        1,
			Ord_modify:           time.Now().Format("2006-01-02 15:04:05"),
			Ord_eclaim:           sub.Ord_eclaim,
		}

		var inInterface map[string]interface{}
		in, _ := json.Marshal(&objSub)
		json.Unmarshal(in, &inInterface)
		delete(inInterface, "room_code")
		delete(inInterface, "room_th")
		delete(inInterface, "room_en")
		delete(inInterface, "room_type_th")
		delete(inInterface, "room_type_en")
		delete(inInterface, "units")
		delete(inInterface, "u_name")
		delete(inInterface, "u_name_en")
		delete(inInterface, "balance")
		delete(inInterface, "label")
		delete(inInterface, "claim_price_ofc")
		delete(inInterface, "claim_price_lgo")
		delete(inInterface, "claim_price_ucs")
		delete(inInterface, "claim_price_sss")
		delete(inInterface, "claim_price_nhs")
		delete(inInterface, "claim_price_ssi")

		if sub.Course_id == nil || *sub.Course_id == 0 {
			delete(inInterface, "course_id")
		}
		if sub.Checking_id == nil || *sub.Checking_id == 0 {
			delete(inInterface, "checking_id")
		}
		if sub.Product_id == nil || *sub.Product_id == 0 {
			delete(inInterface, "product_id")
		}
		if sub.Product_unit_id == nil || *sub.Product_unit_id == 0 {
			delete(inInterface, "product_unit_id")
		}
		if sub.Product_store_id == nil || *sub.Product_store_id == 0 {
			delete(inInterface, "product_store_id")
		}
		if sub.Coin_id == nil || *sub.Coin_id == 0 {
			delete(inInterface, "coin_id")
		}

		if sub.Room_id == nil || *sub.Room_id == 0 {
			delete(inInterface, "room_id")
		}
		if sub.Queue_id == nil || *sub.Queue_id == 0 {
			delete(inInterface, "queue_id")
		}
		if sub.Topical_id == nil || *sub.Topical_id == 0 {
			delete(inInterface, "topical_id")
		}
		// if sub.Queue_ord_id == nil || *sub.Queue_ord_id == 0 {
		delete(inInterface, "queue_ord_id")
		// }

		if err = tx.Table("order_details").Create(inInterface).Error; err != nil {
			tx.Rollback()
			return objH.Id, err
		}
		if sub.Ord_type_id == 1 {
			queuechecking := structs.QueueCheckingByOrder{
				Queue_id:               dataH.Queue_id,
				User_id:                dataH.User_id,
				Checking_id:            *sub.Checking_id,
				Queci_code:             sub.Ord_code,
				Queci_name:             sub.Ord_name,
				Queci_qty:              sub.Ord_qty,
				Queci_unit:             sub.Ord_unit,
				Queci_cost:             sub.Ord_cost,
				Queci_price:            sub.Ord_price,
				Queci_discount_type_id: sub.Ord_discount_type_id,
				Queci_discount_item:    sub.Ord_discount_item,
				Queci_discount:         sub.Ord_discount,
				Queci_total:            sub.Ord_total,
				Queci_id_set:           sub.Ord_id_set,
				Queci_is_set:           sub.Ord_is_set,
				Queci_is_active:        1,
				Queci_modify:           time.Now().Format("2006-01-02 15:04:05"),
				DpmId:                  nil,
			}
			var inInterfaceQ map[string]interface{}
			inQ, _ := json.Marshal(&queuechecking)
			json.Unmarshal(inQ, &inInterfaceQ)
			if err = tx.Table("queue_checkings").Create(inInterfaceQ).Error; err != nil {
				tx.Rollback()
				return objH.Id, err
			}
		}
	}

	//add tags
	for _, tag := range *dataH.Tags {
		objTag := structs.OrderTags{
			Id:       0,
			Order_id: objH.Id,
			Tags_id:  tag.Tags_id,
			Tag_name: tag.Tag_name,
		}
		var inInterface map[string]interface{}
		in, _ := json.Marshal(&objTag)
		json.Unmarshal(in, &inInterface)
		delete(inInterface, "tag_name")
		if err = tx.Table("order_tags").Create(&inInterface).Error; err != nil {
			tx.Rollback()
			return objH.Id, err
		}
	}

	if objH.Queue_id > 0 {
		if err = tx.Table("queues").Where("queues.id = ?", objH.Queue_id).Updates(map[string]interface{}{"queues.que_status_id": 3, "queues.que_update": time.Now().Format("2006-01-02 15:04:05")}).Error; err != nil {
			tx.Rollback()
			return objH.Id, err
		}
	}

	tx.Commit()
	return objH.Id, nil
}

func UpdateOrder(dataH *structs.OrderDetail) (err error) {
	tx := configs.DB1.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	// add checkings
	objH := structs.OrderAction{
		Shop_id:             dataH.Shop_id,
		User_id:             dataH.User_id,
		Customer_id:         dataH.Customer_id,
		Queue_id:            dataH.Queue_id,
		Or_fullname:         dataH.Or_fullname,
		Or_tel:              dataH.Or_tel,
		Or_email:            dataH.Or_email,
		Or_address:          dataH.Or_address,
		Or_district:         dataH.Or_district,
		Or_amphoe:           dataH.Or_amphoe,
		Or_province:         dataH.Or_province,
		Or_zipcode:          dataH.Or_zipcode,
		Or_comment:          dataH.Or_comment,
		Or_total_price:      dataH.Or_total_price,
		Or_discount_type_id: dataH.Or_discount_type_id,
		Or_discount_item:    dataH.Or_discount_item,
		Or_discount_value:   dataH.Or_discount_value,
		Or_discount:         dataH.Or_discount,
		Or_befor_vat:        dataH.Or_befor_vat,
		Tax_type_id:         dataH.Tax_type_id,
		Tax_rate:            dataH.Tax_rate,
		Or_vat:              dataH.Or_vat,
		Or_total:            dataH.Or_total,
		Or_is_active:        dataH.Or_is_active,
		Or_datetime:         dataH.Or_datetime,
		Or_update:           dataH.Or_update,
		DpmId:               dataH.DpmId,
		Or_eclaim_id:        dataH.Or_eclaim_id,
		Or_eclaim_rate:      dataH.Or_eclaim_rate,
		Or_eclaim_over:      dataH.Or_eclaim_over,
		Or_eclaim_total:     dataH.Or_eclaim_total,
	}
	var HinInterface map[string]interface{}
	Hin, _ := json.Marshal(&objH)
	json.Unmarshal(Hin, &HinInterface)
	delete(HinInterface, "id")
	delete(HinInterface, "queue_id")
	delete(HinInterface, "customer_id")
	delete(HinInterface, "user_id")
	delete(HinInterface, "shop_id")
	delete(HinInterface, "or_is_active")
	// delete(HinInterface, "or_datetime")
	delete(HinInterface, "or_create")
	delete(HinInterface, "or_fullname")
	delete(HinInterface, "or_tele_code")

	if err = tx.Table("orders").Where("orders.id = ?", dataH.Id).Updates(&HinInterface).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, sub := range *dataH.Subs {
		objSub := structs.OrderSubUpdate{
			Tax_type_id:          sub.Tax_type_id,
			Tax_rate:             sub.Tax_rate,
			Ord_vat:              sub.Ord_vat,
			Ord_total:            sub.Ord_total,
			Ord_qty:              sub.Ord_qty,
			Ord_amount:           sub.Ord_qty * sub.Ord_price,
			Ord_discount_type_id: sub.Ord_discount_type_id,
			Ord_discount_item:    sub.Ord_discount_item,
			Ord_discount:         sub.Ord_discount,
		}
		var SubinInterface map[string]interface{}
		in, _ := json.Marshal(&objSub)
		json.Unmarshal(in, &SubinInterface)
		if err = tx.Table("order_details").Where("order_details.id = ?", sub.Id).Updates(&SubinInterface).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	//add tags
	for _, tag := range *dataH.Tags {
		objTag := structs.OrderTags{
			Id:       tag.Id,
			Order_id: dataH.Id,
			Tags_id:  tag.Tags_id,
			Tag_name: tag.Tag_name,
		}
		var inInterface map[string]interface{}
		in, _ := json.Marshal(&objTag)
		json.Unmarshal(in, &inInterface)
		delete(inInterface, "tag_name")
		if tag.Id > 0 {
			if err = tx.Table("order_tags").Where("order_tags.id = ?", tag.Id).Updates(&inInterface).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else {
			if err = tx.Table("order_tags").Create(&inInterface).Error; err != nil {
				tx.Rollback()
				return
			}
		}
	}

	tx.Commit()
	return nil
}

func CancelOrder(orderId int, Queue_id int, userId int, data *structs.OrderDetail) (err error) {
	tx := configs.DB1.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	if err = tx.Table("orders").Where("orders.id = ?", orderId).Updates(map[string]interface{}{"orders.or_is_active": 0, "orders.user_id_cancel": userId, "orders.or_update": time.Now().Format("2006-01-02 15:04:05")}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if Queue_id > 0 {
		if err = tx.Table("queues").Where("queues.id = ?", Queue_id).Updates(map[string]interface{}{"queues.que_status_id": 2, "queues.que_update": time.Now().Format("2006-01-02 15:04:05")}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func CancelOrderSub(ordersubId int, data *structs.OrderSub) (err error) {
	query := configs.DB1.Table("order_details")
	query = query.Where("id = ?", ordersubId)
	query = query.Model(&data)
	query = query.Updates(map[string]interface{}{"order_details.ord_is_active": 0, "order_details.ord_modify": time.Now().Format("2006-01-02 15:04:05")})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CancelOrderTag(tagId int, data *structs.OrderSub) (err error) {
	query := configs.DB1.Table("order_tags")
	query = query.Where("order_tags.id = ?", tagId)
	query = query.Delete(&data)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetItemCourseId(course_id int, objQuery *structs.ItemCourseOrder) (err error) {
	query := configs.DB1.Table("courses")
	query = query.Select("courses.*")
	query = query.Where("courses.id = ?", course_id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetItemCheckingId(checking_id int, objQuery *structs.ItemCheckingOrder) (err error) {
	query := configs.DB1.Table("checkings")
	query = query.Select("checkings.*")
	query = query.Where("checkings.id = ?", checking_id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetProductRate(product_unit_id int, pu *structs.ProductUnitList) (err error) {
	query := configs.DB1.Table("product_units")
	query = query.Select("product_units.pu_rate")
	query = query.Where("product_units.id = ?", product_unit_id)
	query = query.Find(&pu)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetOrderTopical(Shop_id int, pu *[]structs.OrderTopical) (err error) {
	query := configs.DB1.Table("topicals")
	query = query.Select("topicals.id,topicals.topical_name,topicals.topical_detail")
	query = query.Where("topicals.shop_id = ?", Shop_id)
	query = query.Where("topicals.topical_type_id = ?", 2)
	query = query.Where("topicals.topical_is_active = ?", 1)
	query = query.Where("topicals.topical_is_del = ?", 0)
	query = query.Find(&pu)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CheckTopical(Topical_id int, pu *[]structs.OrderTopical) (err error) {
	query := configs.DB1.Table("topicals")
	query = query.Select("topicals.id,topicals.topical_name,topicals.topical_detail")
	query = query.Joins("INNER JOIN topical_products ON topicals.id = topical_products.topical_id")
	query = query.Joins("INNER JOIN products ON topical_products.product_id = products.id")
	query = query.Where("topicals.id = ?", Topical_id)
	query = query.Where("topicals.topical_type_id = ?", 2)
	query = query.Where("topicals.topical_is_active = ?", 1)
	query = query.Where("topicals.topical_is_del = ?", 0)
	query = query.Where("products.pd_is_active = ?", 1)
	query = query.Where("products.pd_is_del = ?", 0)
	query = query.Find(&pu)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CheckOrderTopical(Order_id int, Topical_id int, pu *[]structs.OrderTopical) (err error) {
	query := configs.DB1.Table("orders")
	query = query.Select("topical_products.product_id AS id")
	query = query.Joins("INNER JOIN order_details ON orders.id = order_details.order_id")
	query = query.Joins("INNER JOIN topical_products ON order_details.product_id = topical_products.product_id")
	query = query.Joins("INNER JOIN topicals ON topical_products.topical_id = topicals.id")
	query = query.Joins("INNER JOIN products ON topical_products.product_id = products.id")
	query = query.Where("orders.id = ?", Order_id)
	query = query.Where("topicals.id = ?", Topical_id)
	query = query.Where("order_details.ord_qty >= topical_products.tpd_amount")
	query = query.Where("topicals.topical_type_id = ?", 2)
	query = query.Where("topicals.topical_is_active = ?", 1)
	query = query.Where("topicals.topical_is_del = ?", 0)
	query = query.Where("products.pd_is_active = ?", 1)
	query = query.Where("products.pd_is_del = ?", 0)
	query = query.Where("order_details.ord_is_active = ?", 1)
	query = query.Find(&pu)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetOrderByID(orderId int, dt *structs.Order) (err error) {
	query := configs.DB1.Table("orders")
	query = query.Select("orders.*")
	query = query.Where("orders.id = ?", orderId)
	if err = query.Scan(&dt).Error; err != nil {
		return err
	}
	return nil
}

func CheckQueueStatusID(queue_id int, data *structs.CheckQueueStatusID) error {
	query := configs.DB1.Table("queues")
	query = query.Select("queues.id, queues.shop_id, queues.user_id, queues.que_status_id, queues.que_user_fullname, queues.que_tele_code")
	query = query.Where("queues.id = ?", queue_id)
	query = query.Find(&data)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueIPDId(queue_id int, objQuery *structs.GetQueueId) (err error) {
	query := configs.DB1.Table("queues")
	query = query.Select("queues.*")
	query = query.Where("queues.id = ?", queue_id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetCountQueue(queue_id int, objQuery *structs.QueueMax) (err error) {
	query := configs.DB1.Table("queues")
	query = query.Select("Count(queues.id) AS id")
	query = query.Where("queues.que_ref_ipd = ?", queue_id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CreateQueueIPD(objQuery *structs.AddMoveQueueIPD) (err error) {
	query := configs.DB1.Table("queues").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func AddOrderIPD(dataH *structs.OrderDetail) (id int, err error) {
	tx := configs.DB1.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return 0, err
	}

	// add checkings
	objH := structs.OrderAction{
		Id:                  0,
		Shop_id:             dataH.Shop_id,
		User_id:             dataH.User_id,
		Customer_id:         dataH.Customer_id,
		Queue_id:            dataH.Queue_id,
		Or_fullname:         dataH.Or_fullname,
		Or_tel:              dataH.Or_tel,
		Or_email:            dataH.Or_email,
		Or_address:          dataH.Or_address,
		Or_district:         dataH.Or_district,
		Or_amphoe:           dataH.Or_amphoe,
		Or_province:         dataH.Or_province,
		Or_zipcode:          dataH.Or_zipcode,
		Or_comment:          dataH.Or_comment,
		Or_total_price:      dataH.Or_total_price,
		Or_discount_type_id: dataH.Or_discount_type_id,
		Or_discount_item:    dataH.Or_discount_item,
		Or_discount_value:   dataH.Or_discount_value,
		Or_discount:         dataH.Or_discount,
		Or_befor_vat:        dataH.Or_befor_vat,
		Tax_type_id:         dataH.Tax_type_id,
		Tax_rate:            dataH.Tax_rate,
		Or_vat:              dataH.Or_vat,
		Or_total:            dataH.Or_total,
		Or_is_active:        dataH.Or_is_active,
		Or_datetime:         dataH.Or_datetime,
		Or_create:           dataH.Or_create,
		Or_update:           dataH.Or_update,
		Or_tele_code:        dataH.Or_tele_code,
	}

	if err = tx.Table("orders").Create(&objH).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	if dataH.Queue_id > 0 {
		if err = tx.Table("orders").Where("orders.id = ?", objH.Id).Updates(map[string]interface{}{"orders.queue_id": dataH.Queue_id}).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	//add subs
	for _, sub := range *dataH.Subs {
		var Queue_checking_id *int = nil
		var Queue_course_id *int = nil
		var Queue_product_id *int = nil
		if sub.Queue_id != nil {
			if sub.Ord_type_id == 1 {
				Queue_checking_id = sub.Queue_ord_id
			} else if sub.Ord_type_id == 2 {
				Queue_course_id = sub.Queue_ord_id
			} else if sub.Ord_type_id == 3 {
				Queue_product_id = sub.Queue_ord_id
			} else if sub.Ord_type_id == 5 {
				Queue_product_id = sub.Queue_ord_id
			}
		}
		objSub := structs.OrderSub{
			Id:                   0,
			Order_id:             objH.Id,
			Course_id:            sub.Course_id,
			Checking_id:          sub.Checking_id,
			Product_id:           sub.Product_id,
			Product_store_id:     sub.Product_store_id,
			Product_unit_id:      sub.Product_unit_id,
			Coin_id:              sub.Coin_id,
			Room_id:              sub.Room_id,
			Queue_id:             &dataH.Queue_id,
			Queue_checking_id:    Queue_checking_id,
			Queue_course_id:      Queue_course_id,
			Queue_product_id:     Queue_product_id,
			Ord_type_id:          sub.Ord_type_id,
			Ord_code:             sub.Ord_code,
			Ord_name:             sub.Ord_name,
			Ord_qty:              sub.Ord_qty,
			Ord_rate:             sub.Ord_rate,
			Ord_set_qty:          sub.Ord_set_qty,
			Ord_limit_qty:        sub.Ord_limit_qty,
			Ord_unit:             sub.Ord_unit,
			Ord_cost:             sub.Ord_cost,
			Ord_price:            sub.Ord_price,
			Ord_amount:           sub.Ord_amount,
			Ord_discount_type_id: sub.Ord_discount_type_id,
			Ord_discount_item:    sub.Ord_discount_item,
			Ord_discount:         sub.Ord_discount,
			Tax_type_id:          sub.Tax_type_id,
			Tax_rate:             sub.Tax_rate,
			Ord_vat:              sub.Ord_vat,
			Ord_total:            sub.Ord_total,
			Topical_id:           sub.Topical_id,
			Ord_topical:          sub.Ord_topical,
			Ord_direction:        sub.Ord_direction,
			Ord_is_set:           sub.Ord_is_set,
			Ord_is_use:           sub.Ord_is_use,
			Ord_is_active:        1,
			Ord_modify:           time.Now().Format("2006-01-02 15:04:05"),
			Units:                &[]structs.ProductUnitList{},
		}
		var inInterface map[string]interface{}
		in, _ := json.Marshal(&objSub)
		json.Unmarshal(in, &inInterface)
		delete(inInterface, "room_code")
		delete(inInterface, "room_th")
		delete(inInterface, "room_en")
		delete(inInterface, "room_type_th")
		delete(inInterface, "room_type_en")
		delete(inInterface, "units")
		delete(inInterface, "u_name")
		delete(inInterface, "u_name_en")
		delete(inInterface, "balance")
		delete(inInterface, "label")
		delete(inInterface, "queue_ord_id")
		delete(inInterface, "claim_price_ofc")
		delete(inInterface, "claim_price_lgo")
		delete(inInterface, "claim_price_ucs")
		delete(inInterface, "claim_price_sss")
		delete(inInterface, "claim_price_nhs")
		delete(inInterface, "claim_price_ssi")

		if sub.Course_id == nil || *sub.Course_id == 0 {
			delete(inInterface, "course_id")
		}
		if sub.Checking_id == nil || *sub.Checking_id == 0 {
			delete(inInterface, "checking_id")
		}
		if sub.Product_id == nil || *sub.Product_id == 0 {
			delete(inInterface, "product_id")
		}
		if sub.Product_unit_id == nil || *sub.Product_unit_id == 0 {
			delete(inInterface, "product_unit_id")
		}
		if sub.Product_store_id == nil || *sub.Product_store_id == 0 {
			delete(inInterface, "product_store_id")
		}
		if sub.Coin_id == nil || *sub.Coin_id == 0 {
			delete(inInterface, "coin_id")
		}

		if sub.Room_id == nil || *sub.Room_id == 0 {
			delete(inInterface, "room_id")
		}
		if sub.Queue_id == nil || *sub.Queue_id == 0 {
			delete(inInterface, "queue_id")
		}
		if sub.Topical_id == nil || *sub.Topical_id == 0 {
			delete(inInterface, "topical_id")
		}
		// if sub.Ord_eclaim == nil || *sub.Ord_eclaim == 0 {
		// 	delete(inInterface, "ord_eclaim")
		// }
		if err = tx.Table("order_details").Create(inInterface).Error; err != nil {
			tx.Rollback()
			return objH.Id, err
		}

		if Queue_checking_id != nil && *Queue_checking_id > 0 {
			if err = tx.Table("queue_checkings").Where("id = ?", Queue_checking_id).Updates(map[string]interface{}{"queci_ipd_order": 2, "queci_modify": time.Now().Format("2006-01-02 15:04:05")}).Error; err != nil {
				tx.Rollback()
				return objH.Id, err
			}
		}

		if Queue_course_id != nil && *Queue_course_id > 0 {
			if err = tx.Table("queue_courses").Where("id = ?", Queue_course_id).Updates(map[string]interface{}{"quec_ipd_order": 2, "quec_modify": time.Now().Format("2006-01-02 15:04:05")}).Error; err != nil {
				tx.Rollback()
				return objH.Id, err
			}
		}

		if Queue_product_id != nil && *Queue_product_id > 0 {
			if err = tx.Table("queue_products").Where("id = ?", Queue_product_id).Updates(map[string]interface{}{"quep_ipd_order": 2, "quep_modify": time.Now().Format("2006-01-02 15:04:05")}).Error; err != nil {
				tx.Rollback()
				return objH.Id, err
			}
		}
	}

	//add tags
	for _, tag := range *dataH.Tags {
		objTag := structs.OrderTags{
			Id:       0,
			Order_id: objH.Id,
			Tags_id:  tag.Tags_id,
			Tag_name: tag.Tag_name,
		}
		var inInterface map[string]interface{}
		in, _ := json.Marshal(&objTag)
		json.Unmarshal(in, &inInterface)
		delete(inInterface, "tag_name")
		if err = tx.Table("order_tags").Create(&inInterface).Error; err != nil {
			tx.Rollback()
			return objH.Id, err
		}
	}

	if objH.Queue_id > 0 {
		if err = tx.Table("queues").Where("queues.id = ?", objH.Queue_id).Updates(map[string]interface{}{"queues.que_status_id": 3, "queues.que_update": time.Now().Format("2006-01-02 15:04:05")}).Error; err != nil {
			tx.Rollback()
			return objH.Id, err
		}
	}

	tx.Commit()
	return objH.Id, nil
}

func GetQueueProductIPD(queue_id int, objQuery *[]structs.QueueProduct) (err error) {
	query := configs.DB1.Table("queue_products")
	query = query.Select("queue_products.*")
	query = query.Where("queue_products.queue_id = ?", queue_id)
	// query = query.Where("queue_products.quep_type_id = ?", 1)
	query = query.Where("queue_products.quep_is_active = ?", 1)
	query = query.Where("queue_products.quep_ipd_order = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueProductUseIPD(queue_id int, objQuery *[]structs.QueueProduct) (err error) {
	query := configs.DB1.Table("queue_products")
	query = query.Select("queue_products.*")
	query = query.Where("queue_products.queue_id = ?", queue_id)
	query = query.Where("queue_products.quep_type_id = ?", 2)
	query = query.Where("queue_products.quep_is_active = ?", 1)
	query = query.Where("queue_products.quep_ipd_order = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueCourseIPD(queue_id int, objQuery *[]structs.QueueCourse) (err error) {
	query := configs.DB1.Table("queue_courses")
	query = query.Select("queue_courses.*")
	query = query.Where("queue_courses.queue_id = ?", queue_id)
	query = query.Where("queue_courses.quec_is_active = ?", 1)
	query = query.Where("queue_courses.quec_ipd_order = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueCheckingIPD(queue_id int, objQuery *[]structs.QueueChecking) (err error) {
	query := configs.DB1.Table("queue_checkings")
	query = query.Select("queue_checkings.*")
	query = query.Where("queue_checkings.queue_id = ?", queue_id)
	query = query.Where("queue_checkings.queci_is_active = ?", 1)
	query = query.Where("queue_checkings.queci_ipd_order = ?", 1)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CreateQueueItemProductIPD(objQuery *structs.QueueProductIPD) (err error) {
	query := configs.DB1.Table("queue_products").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CreateQueueItemCourseIPD(objQuery *structs.QueueCourseIPD) (err error) {
	query := configs.DB1.Table("queue_courses").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CreateQueueItemCheckingIPD(objQuery *structs.QueueCheckingIPD) (err error) {
	query := configs.DB1.Table("queue_checkings").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateQueueCheckingIPD(id int) (err error) {
	query := configs.DB1.Table("queue_checkings")
	query = query.Where("id = ?", id)
	query = query.Updates(map[string]interface{}{"queue_checkings.queci_is_active": 2, "queue_checkings.queci_modify": time.Now().Format("2006-01-02 15:04:05")})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateQueueCourseIPD(id int) (err error) {
	query := configs.DB1.Table("queue_courses")
	query = query.Where("id = ?", id)
	query = query.Updates(map[string]interface{}{"queue_courses.quec_is_active": 2, "queue_courses.quec_modify": time.Now().Format("2006-01-02 15:04:05")})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateQueueProductIPD(id int) (err error) {
	query := configs.DB1.Table("queue_products")
	query = query.Where("id = ?", id)
	query = query.Updates(map[string]interface{}{"queue_products.quep_is_active": 2, "queue_products.quep_modify": time.Now().Format("2006-01-02 15:04:05")})
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetQueueCheckingIPDRef(queci_id_ref int, objQuery *structs.QueueCheckingIPD) (err error) {
	query := configs.DB1.Table("queue_checkings")
	query = query.Select("queue_checkings.id,queue_checkings.checking_id")
	query = query.Where("queue_checkings.queci_id_ref = ?", queci_id_ref)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetQueueCourseIPDRef(queci_id_ref int, objQuery *structs.QueueCourseIPD) (err error) {
	query := configs.DB1.Table("queue_courses")
	query = query.Select("queue_courses.id,queue_courses.course_id")
	query = query.Where("queue_courses.quec_id_ref = ?", queci_id_ref)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

// labplus
func CheckingLabplus(queue_id int, data *[]structs.CheckingLabplus) (err error) {
	query := configs.DB1.Table("queues")
	query = query.Select("checkings.id,checkings.checking_code,checkings.checking_name")
	query = query.Joins("INNER JOIN queue_checkings ON queues.id = queue_checkings.queue_id")
	query = query.Joins("INNER JOIN checkings ON queue_checkings.checking_id = checkings.id")
	query = query.Where("queues.id = ?", queue_id)
	query = query.Where("checkings.checking_is_labplus = ?", 1)
	query = query.Where("queue_checkings.queci_is_active = ?", 1)
	query = query.Find(&data)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetCustomerLabplus(customer_id int, objQuery *structs.CustomerPatient) (err error) {
	query := configs.DB1.Table("customers")
	query = query.Select("customers.id,ctm_id,ctm_prefix,ctm_fname,ctm_lname,ctm_citizen_id,ctm_birthdate,ctm_gender")
	query = query.Where("customers.id = ?", customer_id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func AddQueuesLabplus(objQuery *structs.AddQueuesLabplus) (err error) {
	query := configs.DB1.Table("queues_labplus").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func DeleteQueuesLabplus(queue_id int, data *structs.QueuesLabplus) (err error) {
	query := configs.DB1.Table("queues_labplus")
	query = query.Where("queues_labplus.queue_id = ?", queue_id)
	query = query.Delete(&data)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetQueueIDLabplus(queue_id int, objQuery *structs.QueuesLabplus) (err error) {
	query := configs.DB1.Table("queues_labplus")
	query = query.Select("queues_labplus.*")
	query = query.Where("queues_labplus.queue_id = ?", queue_id)
	query = query.Find(&objQuery)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetCustomerShopById(ctm_citizen_id string, shop_id int, objQuery *structs.ObjQueryCustomer) error {
	query := configs.DB1.Table("customers")
	query = query.Select(`
		customers.*,
		customer_groups.cg_name,
		customer_groups.cg_save_type,
		customer_groups.cg_save,
		ref_right_treatments.rt_code,
		ref_right_treatments.rt_name,
		ref_right_treatments.rt_name_en
	`)
	query = query.Joins(`
		LEFT JOIN customer_groups ON customer_groups.id = customers.customer_group_id
		LEFT JOIN ref_right_treatments ON ref_right_treatments.id = customers.right_treatment_id
	`)
	query = query.Where("customers.ctm_citizen_id = ?", ctm_citizen_id)
	query = query.Where("customers.shop_id = ?", shop_id)
	return query.Find(objQuery).Error
}

func GetUserType1(shopId int, data *structs.UserType1) (err error) {
	query := configs.DB1.Table("users")
	query = query.Select("users.*")
	query = query.Joins("JOIN user_shops ON user_shops.user_id = users.id")
	query = query.Where("user_shops.shop_id = ?", shopId)
	query = query.Where("user_shops.us_is_owner = 1")
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetCustomerGroup(shopId int, data *structs.CustomerGroupId) (err error) {
	query := configs.DB1.Table("customer_groups")
	query = query.Select("customer_groups.id")
	query = query.Where("customer_groups.shop_id = ?", shopId)
	query = query.Where("customer_groups.cg_is_online = 1")
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func AddCustomerGroup(objQuery *structs.AddCustomerGroup) (err error) {
	query := configs.DB1.Table("customer_groups").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func CreateAccountList(objCreate *AccountList) error {
	query := configs.DB1.Table("account_lists").Create(&objCreate)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetAccountCode(shopId int, data *structs.AccountCode) (err error) {
	query := configs.DB1.Table("account_codes")
	query = query.Select("account_codes.id")
	query = query.Where("account_codes.shop_id = ?", shopId)
	query = query.Where("account_codes.account_type_id = 1")
	query = query.Where("account_codes.acc_is_del = 0")
	query = query.Order("account_codes.id ASC")
	query = query.Limit(1)
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetAccountList(shopId int, data *AccountList) (err error) {
	query := configs.DB1.Table("account_lists")
	query = query.Select("account_lists.id")
	query = query.Where("account_lists.shop_id = ?", shopId)
	query = query.Where("account_lists.acl_type_id = 4")
	query = query.Where("account_lists.acl_code = ?", "P001")
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetRoomBedId(shopId int, rooms *structs.RoomBedId) error {
	query := configs.DB1.Table("rooms")
	query = query.Select("rooms.id AS room_id,beds.id AS bed_id")
	query = query.Joins("INNER JOIN beds ON rooms.id = beds.room_id")
	query = query.Where("rooms.shop_id = ?", shopId)
	query = query.Where("rooms.room_type_id = ?", 1)
	query = query.Where("rooms.room_is_del = ?", 0)
	query = query.Where("beds.bed_is_del = ?", 0)
	query = query.Where("beds.bed_is_active  = ?", 1)
	query = query.Order("rooms.id ASC")
	query = query.Limit(1)
	query = query.Find(&rooms)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func CreateQueueByOrder(objQuery *structs.QueueByOrder) (err error) {
	query := configs.DB1.Table("queues").Create(&objQuery)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func GetOrderDetailProductCheck(order_id int, data *[]structs.OrderDetailCheck) error {
	return configs.DB1.Table("order_details").
		Select("order_details.id, order_details.product_id, order_details.product_store_id, order_details.ord_code, order_details.ord_name, SUM(order_details.ord_qty) AS ord_qty").
		Where("order_details.order_id = ? AND order_details.ord_is_active = ?",
			order_id, 1).
		Where("(order_details.ord_type_id = 3 OR order_details.ord_type_id = 5)").
		Group("order_details.product_store_id").
		Find(data).Error
}

func GetOrderDetailProductCheckCourse(order_id int, product_store_id int, data *structs.OrderDetailCheckCourse) (err error) {
	query := configs.DB1.Table("order_details")
	query = query.Select("order_details.id,SUM( order_details.ord_qty ) AS ord_qty ")
	query = query.Joins("INNER JOIN orders ON order_details.order_id = orders.id ")
	query = query.Where("order_details.order_id = ?", order_id)
	query = query.Where("order_details.product_store_id = ?", product_store_id)
	query = query.Where("order_details.ord_is_active = ?", 1)
	query = query.Where("order_details.course_id IS NOT NULL")
	query = query.Where("order_details.queue_id IS NULL")
	if err = query.Scan(data).Error; err != nil {
		return err
	}
	return nil
}
