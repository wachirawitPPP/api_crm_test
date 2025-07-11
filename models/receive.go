package models

import (
	"encoding/json"
	"linecrmapi/configs"
	"linecrmapi/structs"
	"time"
)

func GetShopStoreReceive(shopId int, ssr *[]structs.ShopStoreReceive) (err error) {
	query := configs.DB1.Table("shop_stores")
	query = query.Select("shop_stores.*")
	query = query.Where("shop_stores.shop_id = ?", shopId)
	if err = query.Scan(&ssr).Error; err != nil {
		return err
	}
	return nil
}

func GetShopStoreIdReceive(Id int, ssr *structs.ShopStoreReceive) (err error) {
	query := configs.DB1.Table("shop_stores")
	query = query.Select("shop_stores.*")
	query = query.Where("shop_stores.id = ?", Id)
	if err = query.Scan(&ssr).Error; err != nil {
		return err
	}
	return nil
}

func GetProductIdReceive(product_id int, data *structs.ProductIdReceive) (err error) {
	query := configs.DB1.Table("products")
	query = query.Select("products.id,products.pd_code,products.pd_name")
	query = query.Where("products.id = ?", product_id)
	if err = query.Scan(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetPurchaseOrders(shopId int, purchase_orders_id int, po *structs.PurchaseOrderReceive) (err error) {
	query := configs.DB1.Table("purchase_orders")
	query = query.Select("purchase_orders.*")
	query = query.Where("purchase_orders.shop_id = ?", shopId)
	query = query.Where("purchase_orders.id = ?", purchase_orders_id)
	if err = query.Scan(&po).Error; err != nil {
		return err
	}
	return nil
}

func GetPurchaseOrdersDetail(purchase_orders_id int, pod *[]structs.PurchaseOrderDetailReceive) (err error) {
	query := configs.DB1.Table("purchase_order_details")
	query = query.Select("purchase_order_details.*, ref_units.u_name")
	query = query.Joins("INNER JOIN product_units ON purchase_order_details.product_unit_id = product_units.id")
	query = query.Joins("INNER JOIN ref_units ON product_units.unit_id = ref_units.id")
	query = query.Where("purchase_order_details.purchase_order_id = ?", purchase_orders_id)
	query = query.Order("purchase_order_details.id ASC")
	err = query.Scan(&pod).Error
	if err != nil {
		return err
	}
	return nil
}

func CheckPurchaseOrdersDetailExpire(purchase_orders_id int, pod *structs.PurchaseOrderDetailReceiveExpire) (err error) {
	query := configs.DB1.Table("purchase_order_details")
	query = query.Select("MIN(purchase_order_details.pd_expire) AS pd_expire")
	query = query.Where("purchase_order_details.purchase_order_id = ?", purchase_orders_id)
	err = query.Scan(&pod).Error
	if err != nil {
		return err
	}
	return nil
}

func GetProductStoresReceive(shop_store_id int, product_id int, pod *structs.ProductStoresReceive) (err error) {
	query := configs.DB1.Table("product_stores")
	query = query.Select("product_stores.*")
	query = query.Where("product_stores.shop_store_id = ?", shop_store_id)
	query = query.Where("product_stores.product_id = ?", product_id)
	err = query.Scan(&pod).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdatePOReceive(poupdate *structs.PurchaseOrderReceive) (err error) {
	// query := configs.DB1.Table("purchase_orders")
	// query = query.Where("id = ?", poupdate.Id)
	// query = query.Updates(&poupdate)
	// if err = query.Error; err != nil {
	// 	return err
	// }
	// return nil
	query := configs.DB1.Table("purchase_orders")
	query = query.Where("id = ?", poupdate.Id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&poupdate)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func AddProductStoresReceive(psadd *structs.ProductStoresReceive) (err error) {
	query := configs.DB1.Table("product_stores").Create(&psadd)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateProductStoresReceive(psupdate *structs.ProductStoresReceiveUpdate) (err error) {
	// query := configs.DB1.Table("product_stores")
	// query = query.Where("id = ?", psupdate.Id)
	// query = query.Updates(&psupdate)
	// if err = query.Error; err != nil {
	// 	return err
	// }
	// return nil
	query := configs.DB1.Table("product_stores")
	query = query.Where("id = ?", psupdate.Id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&psupdate)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func AddProductStoresOrderReceive(psoadd *structs.ProductStoresOrderReceive) (err error) {
	query := configs.DB1.Table("product_store_orders").Create(&psoadd)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func AddProductStoresHistoryReceive(pshadd *structs.ProductStoresHistoryReceive) (err error) {
	query := configs.DB1.Table("product_store_historys").Create(&pshadd)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdatePurchaseOrdersReceive(Shop_id int, Shop_store_id int, PO *structs.PurchaseOrderReceive, POD *[]structs.PurchaseOrderDetailReceive) (err error) {
	tx := configs.DB1.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	// update PO
	if err = tx.Table("purchase_orders").Where("purchase_orders.id = ?", PO.Id).Updates(&PO).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, checkps := range *POD {
		var PS structs.ProductStoresReceive
		if errPS := GetProductStoresReceive(Shop_store_id, checkps.Pd_id, &PS); errPS != nil || PS.Id == 0 {
			//Add ProductStoresReceive
			PSAdd := structs.ProductStoresReceive{
				Id:            0,
				Shop_store_id: Shop_store_id,
				Product_id:    checkps.Pd_id,
				Pds_barcode:   checkps.Pd_code,
				Pds_cost:      checkps.Pod_cost,
				Pds_in:        checkps.Pod_qty,
				Pds_out:       0.00,
				Pds_total:     checkps.Pod_qty,
				Pds_date:      time.Now().Format("2006-01-02 15:04:05"),
				Pds_comment:   "",
				Pds_is_active: 1,
				Pds_is_del:    0,
				Pds_create:    time.Now().Format("2006-01-02 15:04:05"),
				Pds_update:    time.Now().Format("2006-01-02 15:04:05"),
			}
			if err = tx.Table("product_stores").Create(&PSAdd).Error; err != nil {
				tx.Rollback()
				return err
			}
			//date to string
			dateTimeString := checkps.Pd_expire
			parsedTime, _ := time.Parse(time.RFC3339, dateTimeString)
			formattedString := parsedTime.Format("2006-01-02 15:04:05")
			//Add ProductStoresOrderReceive
			PSOAdd := structs.ProductStoresOrderReceive{
				Id:                       0,
				Product_store_id:         PSAdd.Id,
				Purchase_order_id:        checkps.Purchase_order_id,
				Purchase_order_detail_id: checkps.Id,
				Pdso_code:                PO.Po_code,
				Pdso_cost:                checkps.Pod_cost,
				Pdso_date:                time.Now().Format("2006-01-02 15:04:05"),
				Pdso_expire:              formattedString,
				Pdso_in:                  checkps.Pod_qty,
				Pdso_out:                 0,
				Pdso_use:                 0,
				Pdso_move:                0,
				Pdso_total:               checkps.Pod_qty,
				Pdso_is_active:           1,
				Pdso_update:              time.Now().Format("2006-01-02 15:04:05"),
				Pdso_create:              time.Now().Format("2006-01-02 15:04:05"),
			}
			if err = tx.Table("product_store_orders").Create(&PSOAdd).Error; err != nil {
				tx.Rollback()
				return err
			}

			//Add ProductStoresHistoryReceive
			PSHAdd := structs.ProductStoresHistoryReceive{
				Id:                       0,
				Shop_id:                  Shop_id,
				Shop_store_id:            Shop_store_id,
				Product_store_id:         PSAdd.Id,
				Pdsh_in:                  checkps.Pod_qty,
				Product_store_order_id:   PSOAdd.Id,
				Purchase_order_detail_id: checkps.Id,
				Pdsh_out:                 0,
				Pdsh_inout:               1,
				Pdsh_out_id:              0,
				Pdsh_amount:              checkps.Pod_qty,
				Pdsh_order_total:         checkps.Pod_qty,
				Pdsh_total:               checkps.Pod_qty,
				Pdsh_ref_doc_no:          PO.Po_code,
				Pdsh_comment:             "Add Receive New Stores",
				User_id:                  PO.User_id_receive,
				Product_id:               checkps.Pd_id,
				Pd_code:                  checkps.Pd_code,
				Pd_name:                  checkps.Pd_name,
				Pdsh_type_id:             1,
				Pdsh_date:                time.Now().Format("2006-01-02 15:04:05"),
				Pdsh_modify:              time.Now().Format("2006-01-02 15:04:05"),
			}
			if err = tx.Table("product_store_historys").Create(&PSHAdd).Error; err != nil {
				tx.Rollback()
				return err
			}

		} else {
			//Update ProductStoresReceive
			PSUpdate := structs.ProductStoresReceive{
				Id:         PS.Id,
				Pds_cost:   checkps.Pod_cost,
				Pds_in:     PS.Pds_in + checkps.Pod_qty,
				Pds_total:  PS.Pds_total + checkps.Pod_qty,
				Pds_update: time.Now().Format("2006-01-02 15:04:05"),
			}

			if err = tx.Table("product_stores").Where("product_stores.id = ?", PSUpdate.Id).Updates(&PSUpdate).Error; err != nil {
				tx.Rollback()
				return err
			}
			//date to string
			dateTimeString := checkps.Pd_expire
			parsedTime, _ := time.Parse(time.RFC3339, dateTimeString)
			formattedString := parsedTime.Format("2006-01-02 15:04:05")
			//Add ProductStoresOrderReceive
			PSOAdd := structs.ProductStoresOrderReceive{
				Id:                       0,
				Product_store_id:         PS.Id,
				Purchase_order_id:        checkps.Purchase_order_id,
				Purchase_order_detail_id: checkps.Id,
				Pdso_code:                PO.Po_code,
				Pdso_cost:                checkps.Pod_cost,
				Pdso_date:                time.Now().Format("2006-01-02 15:04:05"),
				Pdso_expire:              formattedString,
				Pdso_in:                  checkps.Pod_qty,
				Pdso_out:                 0,
				Pdso_use:                 0,
				Pdso_move:                0,
				Pdso_total:               checkps.Pod_qty,
				Pdso_is_active:           1,
				Pdso_update:              time.Now().Format("2006-01-02 15:04:05"),
				Pdso_create:              time.Now().Format("2006-01-02 15:04:05"),
			}
			if err = tx.Table("product_store_orders").Create(&PSOAdd).Error; err != nil {
				tx.Rollback()
				return err
			}

			//Add ProductStoresHistoryReceive
			PSHAdd := structs.ProductStoresHistoryReceive{
				Id:                       0,
				Shop_id:                  Shop_id,
				Shop_store_id:            Shop_store_id,
				Product_store_id:         PS.Id,
				Pdsh_in:                  checkps.Pod_qty,
				Product_store_order_id:   PSOAdd.Id,
				Purchase_order_detail_id: checkps.Id,
				Pdsh_out:                 0,
				Pdsh_inout:               1,
				Pdsh_out_id:              0,
				Pdsh_total_forward:       PS.Pds_total,
				Pdsh_amount:              checkps.Pod_qty,
				Pdsh_order_total:         PSOAdd.Pdso_total,
				Pdsh_total:               PSUpdate.Pds_total, //PS.Pds_total + checkps.Pod_qty,
				Pdsh_ref_doc_no:          PO.Po_code,
				Pdsh_comment:             "Add Receive",
				User_id:                  PO.User_id_receive,
				Product_id:               checkps.Pd_id,
				Pd_code:                  checkps.Pd_code,
				Pd_name:                  checkps.Pd_name,
				Pdsh_type_id:             1,
				Pdsh_date:                time.Now().Format("2006-01-02 15:04:05"),
				Pdsh_modify:              time.Now().Format("2006-01-02 15:04:05"),
			}
			if err = tx.Table("product_store_historys").Create(&PSHAdd).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	tx.Commit()
	return nil
}

// Transfers
func GetTransfers(shopId int, transfers_id int, tf *structs.TransferReceive) (err error) {
	query := configs.DB1.Table("transfers")
	query = query.Select("transfers.*")
	query = query.Where("transfers.shop_id_to = ?", shopId)
	query = query.Where("transfers.id = ?", transfers_id)
	if err = query.Scan(&tf).Error; err != nil {
		return err
	}
	return nil
}

func GetProductTransfers(transfers_id int, tfd *[]structs.TransferDetailReceive) (err error) {
	query := configs.DB1.Table("transfer_details")
	query = query.Select("transfer_details.*")
	query = query.Where("transfer_details.transfer_id = ?", transfers_id)
	if err = query.Scan(&tfd).Error; err != nil {
		return err
	}
	return nil
}

func CheckProductStoreOrderTransfer(product_store_order_id int, psoc *structs.ProductStoresOrderCheck) (err error) {
	query := configs.DB1.Table("product_store_orders")
	query = query.Select("product_store_orders.*")
	query = query.Where("product_store_orders.id = ?", product_store_order_id)
	if err = query.Scan(&psoc).Error; err != nil {
		return err
	}
	return nil
}

func GetProductStoreOrder(product_store_order_id int, psot *structs.ProductStoresOrderTransfer) (err error) {
	query := configs.DB1.Table("product_store_orders")
	query = query.Select("product_store_orders.*")
	query = query.Where("product_store_orders.id = ?", product_store_order_id)
	if err = query.Scan(&psot).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProductStoresOrderTransfer(psotupdate *structs.ProductStoresOrderTransfer) (err error) {
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&psotupdate)
	json.Unmarshal(in, &inInterface)

	query := configs.DB1.Table("product_store_orders")
	query = query.Where("id = ?", psotupdate.Id)
	query = query.Model(&psotupdate)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateTransferReceive(tfupdate *structs.TransferReceiveUpdate) (err error) {
	// query := configs.DB1.Table("transfers")
	// query = query.Where("id = ?", tfupdate.Id)
	// query = query.Updates(&tfupdate)
	// if err = query.Error; err != nil {
	// 	return err
	// }
	// return nil
	query := configs.DB1.Table("transfers")
	query = query.Where("id = ?", tfupdate.Id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&tfupdate)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func AddProductStoresHistoryTransfer(pshadd *structs.ProductStoresHistoryTransfer) (err error) {
	query := configs.DB1.Table("product_store_historys").Create(&pshadd)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func AddProductStoresOrderReceiveTransfer(psoadd *structs.ProductStoresOrderReceiveTransfer) (err error) {
	query := configs.DB1.Table("product_store_orders").Create(&psoadd)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func AddProductStoresHistoryReceiveTransfer(pshadd *structs.ProductStoresHistoryReceiveTransfer) (err error) {
	query := configs.DB1.Table("product_store_historys").Create(&pshadd)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

// // Rollback
// func GetProductStoreOrderRollback(tx *gorm.DB, productStoreOrderID int, PSOT *structs.ProductStoresOrderTransfer) error {
// 	if err := tx.Table("product_store_orders").
// 		Select("product_store_orders.*").
// 		Where("product_store_orders.id = ?", productStoreOrderID).
// 		Scan(PSOT).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func GetProductStoresReceiveRollback(tx *gorm.DB, shop_store_id int, product_id int, pod *structs.ProductStoresReceive) (err error) {
// 	if err := tx.Table("product_stores").
// 		Select("product_stores.*").
// 		Where("product_stores.shop_store_id = ?", shop_store_id).
// 		Where("product_stores.product_id = ?", product_id).
// 		Scan(&pod).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// Issue
func GetIssue(shopId int, issue_id int, iss *structs.IssueStock) (err error) {
	query := configs.DB1.Table("issues")
	query = query.Select("issues.*")
	query = query.Where("issues.shop_id = ?", shopId)
	query = query.Where("issues.id = ?", issue_id)
	if err = query.Scan(&iss).Error; err != nil {
		return err
	}
	return nil
}

func GetProductIssue(issue_id int, isds *[]structs.IssueDetailStock) (err error) {
	query := configs.DB1.Table("issue_details")
	query = query.Select("issue_details.*")
	query = query.Where("issue_details.issue_id = ?", issue_id)
	if err = query.Scan(&isds).Error; err != nil {
		return err
	}
	return nil
}

func GetProductStoresIssue(shop_store_id int, product_id int, psis *structs.ProductStoresIssues) (err error) {
	query := configs.DB1.Table("product_stores")
	query = query.Select("product_stores.*")
	query = query.Where("product_stores.shop_store_id = ?", shop_store_id)
	query = query.Where("product_stores.product_id = ?", product_id)
	err = query.Scan(&psis).Error
	if err != nil {
		return err
	}
	return nil
}

func GetProductStoreOrderIssue(product_store_order_id int, psois *structs.ProductStoresOrderIssues) (err error) {
	query := configs.DB1.Table("product_store_orders")
	query = query.Select("product_store_orders.*")
	query = query.Where("product_store_orders.id = ?", product_store_order_id)
	if err = query.Scan(&psois).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProductStoresIssue(psupdate *structs.ProductStoresIssuesUpdate) (err error) {
	// query := configs.DB1.Table("product_stores")
	// query = query.Where("id = ?", psupdate.Id)
	// query = query.Updates(&psupdate)
	// if err = query.Error; err != nil {
	// 	return err
	// }
	// return nil
	query := configs.DB1.Table("product_stores")
	query = query.Where("id = ?", psupdate.Id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&psupdate)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateProductStoresOrderIssue(psotupdate *structs.ProductStoresOrderIssues) (err error) {
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&psotupdate)
	json.Unmarshal(in, &inInterface)

	query := configs.DB1.Table("product_store_orders")
	query = query.Where("id = ?", psotupdate.Id)
	query = query.Model(&psotupdate)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func AddProductStoresHistoryIssue(pshadd *structs.ProductStoresHistoryIssues) (err error) {
	query := configs.DB1.Table("product_store_historys").Create(&pshadd)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateIssueConfirm(isuupdate *structs.IssueStockUpdate) (err error) {
	// query := configs.DB1.Table("issues")
	// query = query.Where("id = ?", isuupdate.Id)
	// query = query.Updates(&isuupdate)
	// if err = query.Error; err != nil {
	// 	return err
	// }
	// return nil
	query := configs.DB1.Table("issues")
	query = query.Where("id = ?", isuupdate.Id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&isuupdate)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func AddLogISS(log *structs.LogISS) (err error) {
	query := configs.DBL1.Table("log_issue").Create(&log)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

// Adjust
func GetAdjust(shopId int, adjust_id int, adjs *structs.AdjustStock) (err error) {
	query := configs.DB1.Table("adjusts")
	query = query.Select("adjusts.*")
	query = query.Where("adjusts.shop_id = ?", shopId)
	query = query.Where("adjusts.id = ?", adjust_id)
	if err = query.Scan(&adjs).Error; err != nil {
		return err
	}
	return nil
}

func GetProductAdjust(adjust_id int, adjds *[]structs.AdjustDetailStock) (err error) {
	query := configs.DB1.Table("adjust_details")
	query = query.Select("adjust_details.*")
	query = query.Where("adjust_details.adjust_id = ?", adjust_id)
	if err = query.Scan(&adjds).Error; err != nil {
		return err
	}
	return nil
}

func GetProductStoresAdjust(shop_store_id int, product_id int, psadj *structs.ProductStoresAdjust) (err error) {
	query := configs.DB1.Table("product_stores")
	query = query.Select("product_stores.*")
	query = query.Where("product_stores.shop_store_id = ?", shop_store_id)
	query = query.Where("product_stores.product_id = ?", product_id)
	err = query.Scan(&psadj).Error
	if err != nil {
		return err
	}
	return nil
}

func GetProductStoreOrderAdjust(product_store_order_id int, psoadj *structs.ProductStoresOrderAdjust) (err error) {
	query := configs.DB1.Table("product_store_orders")
	query = query.Select("product_store_orders.*")
	query = query.Where("product_store_orders.id = ?", product_store_order_id)
	if err = query.Scan(&psoadj).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProductStoresAdjust(psupdate *structs.ProductStoresAdjustUpdate) (err error) {
	// query := configs.DB1.Table("product_stores")
	// query = query.Where("id = ?", psupdate.Id)
	// query = query.Updates(&psupdate)
	// if err = query.Error; err != nil {
	// 	return err
	// }
	// return nil
	query := configs.DB1.Table("product_stores")
	query = query.Where("id = ?", psupdate.Id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&psupdate)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateProductStoresOrderAdjust(psotupdate *structs.ProductStoresOrderAdjust) (err error) {
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&psotupdate)
	json.Unmarshal(in, &inInterface)

	query := configs.DB1.Table("product_store_orders")
	query = query.Where("id = ?", psotupdate.Id)
	query = query.Model(&psotupdate)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func AddProductStoresHistoryAdjust(pshadd *structs.ProductStoresHistoryAdjust) (err error) {
	query := configs.DB1.Table("product_store_historys").Create(&pshadd)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateAdjustConfirm(adjupdate *structs.AdjustStock) (err error) {
	// query := configs.DB1.Table("adjusts")
	// query = query.Where("id = ?", adjupdate.Id)
	// query = query.Updates(&adjupdate)
	// if err = query.Error; err != nil {
	// 	return err
	// }
	// return nil
	query := configs.DB1.Table("adjusts")
	query = query.Where("id = ?", adjupdate.Id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&adjupdate)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func UpdateProductAdjustDetail(adjdupdate *structs.AdjustDetailUpdate) (err error) {
	// query := configs.DB1.Table("adjust_details")
	// query = query.Where("id = ?", adjdupdate.Id)
	// query = query.Updates(&adjdupdate)
	// if err = query.Error; err != nil {
	// 	return err
	// }
	// return nil
	query := configs.DB1.Table("adjust_details")
	query = query.Where("id = ?", adjdupdate.Id)
	var inInterface map[string]interface{}
	in, _ := json.Marshal(&adjdupdate)
	json.Unmarshal(in, &inInterface)
	query = query.Updates(&inInterface)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}

func AddLogADJ(log *structs.LogADJ) (err error) {
	query := configs.DBL1.Table("log_adjust").Create(&log)
	if err = query.Error; err != nil {
		return err
	}
	return nil
}
