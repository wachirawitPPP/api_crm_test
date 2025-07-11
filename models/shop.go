package models

import (
	"linecrmapi/configs"
	"linecrmapi/structs"
	"time"
)

func GetShopById(shopId int, objResponse *structs.ShopReadResponse) error {
	query := configs.DB1.Table("shops")
	query = query.Select("shops.*")
	// query = query.Joins("INNER JOIN doc_settings ON doc_settings.shop_id = shops.id")
	query = query.Where("shops.id = ?", shopId)
	query = query.Where("shops.shop_status_id = ?", 1)
	query = query.Find(&objResponse)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetShopCustomerById(shop_mother_id int, citizen_id string, result *Customer) (err error) {
	query := configs.DB1.Table("customers")
	query = query.Select("customers.id")
	query = query.Joins("INNER JOIN shops ON shops.id = customers.shop_id")
	query = query.Where("customers.ctm_citizen_id = ?", citizen_id)
	// query = query.Where("customers.shop_id = ?", shop_id)
	query = query.Where("customers.shop_mother_id = ?", shop_mother_id)
	query = query.Where("customers.ctm_is_active = 1")
	query = query.Where("customers.ctm_is_del = 0")
	query = query.Where("shops.shop_status_id = 1")
	query = query.Where("shops.shop_expire > ?", time.Now().Format("2006-01-02"))
	query = query.Group("customers.id")
	query = query.Order("shops.id ASC")
	query = query.Limit(1)
	if err := query.Scan(&result).Error; err != nil {
		return err
	}
	return nil
}

func GetInShopList(citizen_id string) ([]int64, error) {
	var mother_id []int64
	query := configs.DB1.Table("customers")
	query = query.Select("customers.shop_mother_id")
	query = query.Where("customers.ctm_citizen_id = ?", citizen_id)
	query = query.Group("customers.shop_mother_id")
	query = query.Find(&mother_id)
	if query.Error != nil {
		return mother_id, query.Error
	}
	return mother_id, nil
}

func GetInShopListByMother(shop_mother_id []int64, objResponse *[]structs.InShopList) error {
	query := configs.DB1.Table("shops")
	query = query.Select("shops.id, shops.shop_code as code,shops.shop_name as name,shops.shop_mother_id,shops.shop_nature as nature,shops.shop_province as province,shops.shop_latlong as latlong,shops.shop_image as image,nature_types.nature_type_name as nature_type")
	query = query.Joins("INNER JOIN nature_types ON nature_types.id = shops.nature_type_id")
	query = query.Where("shops.shop_mother_id IN ?", shop_mother_id)
	query = query.Where("shops.shop_status_id = ?", 1)
	query = query.Where("shops.shop_expire > ?", time.Now().Format("2006-01-02"))
	query = query.Group("shops.id")
	query = query.Order("shops.shop_code ASC")
	query = query.Find(&objResponse)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
