package controllers

import (
	"linecrmapi/models"
	"linecrmapi/structs"
	"net/http"
	"sort"
	"sync"

	"github.com/gin-gonic/gin"
)

func ItemProductList(c *gin.Context) {
	var filter structs.ObjPayloaItem
	if errSBJ := c.ShouldBindJSON(&filter); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errSBJ.Error(),
		})
		return
	}

	var ItemList []structs.ItemProduct
	if errMD := models.GetItemProductList(filter, true, &ItemList); errMD != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    errMD.Error(),
		})
		return
	}

	if len(ItemList) == 0 {
		emptyList := []structs.ItemProduct{}
		c.JSON(200, gin.H{
			"status":  true,
			"message": "No Data",
			"data":    emptyList,
		})
		return
	}

	ItemListProduct := []structs.ItemProduct{}

	var wg sync.WaitGroup
	for i, st := range ItemList {
		wg.Add(1)
		go func(st string, i int) {
			defer wg.Done()
			if ItemList[i].Pd_type_id != 3 {
				var ProductSubSets []structs.ItemProductSub
				if errS := models.GetItemProductSetId(ItemList[i].Id, ItemList[i].Shop_id, &ProductSubSets); errS != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Product Set invalid.",
						"data":    errS.Error(),
					})
					return
				}

				var objQueryShopStore models.ShopStore
				errShopStore := models.GetShopStoreByIdType1(ItemList[i].Shop_id, &objQueryShopStore)
				if errShopStore != nil {
					c.JSON(http.StatusOK, gin.H{
						"status":  false,
						"message": "Get data shop store error!",
						"data":    errShopStore,
					})
				}

				var productStoreBalance structs.ObjQueryProductStoreBalance
				if errT := models.GetProductStoreBalance(objQueryShopStore.ID, ItemList[i].Product_id, &productStoreBalance); errT != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Product store balance invalid.",
						"data":    errT,
					})
					return
				}
				if len(ProductSubSets) > 0 {
					for _, ss := range ProductSubSets {
						var UnitSub []structs.ItemProductUnit
						if errU := models.GetItemProductUnitList(ss.Id, ItemList[i].Shop_id, &UnitSub); errU != nil {
							c.AbortWithStatusJSON(200, gin.H{
								"status":  false,
								"message": "Product Unit Sub invalid.",
								"data":    errU.Error(),
							})
							return
						}
						Topical_id := ss.Topical_id
						var Topical_detail string
						if Topical_id > 0 {
							var topical structs.ItemTopical
							if errT := models.GetItemTopicalId(Topical_id, &topical); errT != nil {
								c.AbortWithStatusJSON(200, gin.H{
									"status":  false,
									"message": "Topical invalid.",
									"data":    errT,
								})
								return
							}
							Topical_detail = topical.Topical_detail
						}
						ItemList[i].Subs = []structs.ItemProductSub{
							{
								Id:               ss.Id,
								Product_id:       ss.Product_id,
								Product_store_id: ss.Product_store_id,
								Product_units_id: ss.Product_units_id,
								Pd_code:          ss.Pd_code,
								Pd_name:          ss.Pd_name,
								U_name:           ss.U_name,
								Pu_amount:        ss.Pu_amount,
								Pu_rate:          ss.Pu_rate,
								Balance:          productStoreBalance.Pds_balance,
								Psp_price_ipd:    ss.Psp_price_ipd,
								Psp_price_opd:    ss.Psp_price_opd,
								Topical_id:       Topical_id,
								Topical_detail:   Topical_detail,
								Drug_direction:   ss.Drug_direction,
								Pd_image_1:       ss.Pd_image_1,
								Pd_image_2:       ss.Pd_image_2,
								Pd_image_3:       ss.Pd_image_3,
								Pd_image_4:       ss.Pd_image_4,
								Pd_detail:        ss.Pd_detail,
								Label:            "",
								Is_set:           0,
								Id_set:           nil,
								Units:            UnitSub,
							},
						}
					}
				}
			} else { //Set
				var ProductSets []structs.ItemProductSet
				ItemProductSub := []structs.ItemProductSub{}
				if errS := models.GetItemProductSetList(ItemList[i].Id, ItemList[i].Shop_id, &ProductSets); errS != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Product Set invalid.",
						"data":    errS.Error(),
					})
					return
				}

				if len(ProductSets) > 0 {
					for _, sset := range ProductSets {
						var ProductSubSets []structs.ItemProductSub
						if errS := models.GetItemProductSetId(sset.Id, ItemList[i].Shop_id, &ProductSubSets); errS != nil {
							c.AbortWithStatusJSON(200, gin.H{
								"status":  false,
								"message": "Product Set invalid.",
								"data":    errS.Error(),
							})
							return
						}
						var objQueryShopStore models.ShopStore
						errShopStore := models.GetShopStoreByIdType1(ItemList[i].Shop_id, &objQueryShopStore)
						if errShopStore != nil {
							c.JSON(http.StatusOK, gin.H{
								"status":  false,
								"message": "Get data shop store error!",
								"data":    errShopStore,
							})
						}
						var productStoreBalance structs.ObjQueryProductStoreBalance
						if errT := models.GetProductStoreBalance(objQueryShopStore.ID, sset.Id, &productStoreBalance); errT != nil {
							c.AbortWithStatusJSON(200, gin.H{
								"status":  false,
								"message": "Product store balance invalid.",
								"data":    errT,
							})
							return
						}
						Topical_id := sset.Topical_id
						var Topical_detail string
						if Topical_id > 0 {
							var topical structs.ItemTopical
							if errT := models.GetItemTopicalId(Topical_id, &topical); errT != nil {
								c.AbortWithStatusJSON(200, gin.H{
									"status":  false,
									"message": "Topical invalid.",
									"data":    errT,
								})
								return
							}
							Topical_detail = topical.Topical_detail
						}

						var UnitSub []structs.ItemProductUnit
						for _, ss := range ProductSubSets {
							if errU := models.GetItemProductUnitList(ss.Id, ItemList[i].Shop_id, &UnitSub); errU != nil {
								c.AbortWithStatusJSON(200, gin.H{
									"status":  false,
									"message": "Product Unit Sub invalid.",
									"data":    errU.Error(),
								})
								return
							}
							ProductSub := structs.ItemProductSub{
								Id:               ss.Id,
								Product_id:       ss.Product_id,
								Product_store_id: ss.Product_store_id,
								Product_units_id: ss.Product_units_id,
								Pd_code:          ss.Pd_code,
								Pd_name:          ss.Pd_name,
								U_name:           ss.U_name,
								Pu_amount:        sset.Product_amount,
								Pu_rate:          ss.Pu_rate,
								Balance:          productStoreBalance.Pds_balance,
								Psp_price_ipd:    sset.Product_list_ipd,
								Psp_price_opd:    sset.Product_list_opd,
								Topical_id:       Topical_id,
								Topical_detail:   Topical_detail,
								Drug_direction:   sset.Drug_direction,
								Pd_image_1:       ss.Pd_image_1,
								Pd_image_2:       ss.Pd_image_2,
								Pd_image_3:       ss.Pd_image_3,
								Pd_image_4:       ss.Pd_image_4,
								Pd_detail:        ss.Pd_detail,
								Label:            ItemList[i].Pd_name, //"Set",
								Is_set:           1,
								Id_set:           &ItemList[i].Product_id,
								Units:            UnitSub,
							}
							ItemProductSub = append(ItemProductSub, ProductSub)
						}
					}
					ItemList[i].Subs = ItemProductSub
				}
			}

			if len(ItemList[i].Subs) > 0 {
				ItemListProduct = append(ItemListProduct, ItemList[i])
			}
		}(st.Pd_code, i)
	}
	wg.Wait()

	sort.Slice(ItemListProduct, func(i, j int) bool {
		return ItemListProduct[i].Pd_code < ItemListProduct[j].Pd_code
	})

	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    ItemListProduct,
	})
}

func ItemCourseList(c *gin.Context) {
	var filter structs.ObjPayloaItem
	if errSBJ := c.ShouldBindJSON(&filter); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errSBJ.Error(),
		})
		return
	}

	var ItemList []structs.ItemCourse
	if errMD := models.GetItemCourseList(filter, true, &ItemList); errMD != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    errMD.Error(),
		})
		return
	}

	if len(ItemList) == 0 {
		emptyList := []structs.ItemCourse{}
		c.JSON(200, gin.H{
			"status":  true,
			"message": "",
			"data":    emptyList,
		})
		return
	}

	ItemListCourse := []structs.ItemCourse{}

	var wg sync.WaitGroup
	for i, st := range ItemList {
		wg.Add(1)
		go func(st string, i int) {
			defer wg.Done()
			var ProductSub []structs.ItemProductSubSet
			if errU := models.GetItemCourseProduct(ItemList[i].Id, ItemList[i].Shop_id, &ProductSub); errU != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Course Product invalid.",
					"data":    errU.Error(),
				})
				return
			}

			if ItemList[i].Course_type_id == 1 {
				if len(ProductSub) == 0 {
					if ItemList[i].Course_lock_drug == 0 {
						if ItemList[i].Course_qtyset != 1 {
							ItemList[i].Course_amount = ItemList[i].Course_qtyset
						}
						ItemList[i].Subs = []structs.ItemCourseSub{
							{
								Id:             ItemList[i].Id,
								Course_type_id: ItemList[i].Course_type_id,
								Course_code:    ItemList[i].Course_code,
								Course_name:    ItemList[i].Course_name,
								Course_amount:  ItemList[i].Course_amount,
								Course_unit:    ItemList[i].Course_unit,
								Course_cost:    ItemList[i].Course_cost,
								Course_opd:     ItemList[i].Course_opd,
								Course_ipd:     ItemList[i].Course_ipd,
								Course_image_1: ItemList[i].Course_image_1,
								Course_image_2: ItemList[i].Course_image_2,
								Course_image_3: ItemList[i].Course_image_3,
								Course_image_4: ItemList[i].Course_image_4,
								Course_detail:  ItemList[i].Course_detail,
								Label:          "",
								Is_set:         0,
								Id_set:         nil,
								Products:       []structs.ItemProductSubSet{},
							},
						}
					}
				} else {
					for ip, productSubList := range ProductSub {

						var objQueryShopStore models.ShopStore
						errShopStore := models.GetShopStoreByIdType1(ItemList[i].Shop_id, &objQueryShopStore)
						if errShopStore != nil {
							c.JSON(http.StatusOK, gin.H{
								"status":  false,
								"message": "Get data shop store error!",
								"data":    errShopStore,
							})
						}

						ProductSub[ip].Balance = 0
						var productStoreBalance structs.ObjQueryProductStoreBalance
						if errT := models.GetProductStoreBalance(objQueryShopStore.ID, productSubList.Product_id, &productStoreBalance); errT != nil {
							c.AbortWithStatusJSON(200, gin.H{
								"status":  false,
								"message": "Product store balance invalid.",
								"data":    errT,
							})
							return
						}
						ProductSub[ip].Label = ItemList[i].Course_code + ":" + ItemList[i].Course_name
						ProductSub[ip].Balance = productStoreBalance.Pds_balance
					}
					if ItemList[i].Course_qtyset != 1 {
						ItemList[i].Course_amount = ItemList[i].Course_qtyset
					}
					ItemList[i].Subs = []structs.ItemCourseSub{
						{
							Id:             ItemList[i].Id,
							Course_type_id: ItemList[i].Course_type_id,
							Course_code:    ItemList[i].Course_code,
							Course_name:    ItemList[i].Course_name,
							Course_amount:  ItemList[i].Course_amount,
							Course_unit:    ItemList[i].Course_unit,
							Course_cost:    ItemList[i].Course_cost,
							Course_opd:     ItemList[i].Course_opd,
							Course_ipd:     ItemList[i].Course_ipd,
							Course_image_1: ItemList[i].Course_image_1,
							Course_image_2: ItemList[i].Course_image_2,
							Course_image_3: ItemList[i].Course_image_3,
							Course_image_4: ItemList[i].Course_image_4,
							Course_detail:  ItemList[i].Course_detail,
							Label:          "",
							Is_set:         0,
							Id_set:         nil,
							Products:       ProductSub,
						},
					}
				}
			} else {
				var CourseSets []structs.ItemCourseSub
				ItemCourseSub := []structs.ItemCourseSub{}
				if errS := models.GetItemCourseIdSet(ItemList[i].Id, &CourseSets); errS != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Course Set invalid.",
						"data":    errS.Error(),
					})
					return
				}

				if len(CourseSets) > 0 {
					for _, cset := range CourseSets {
						var ProductSub []structs.ItemProductSubSet
						if errU := models.GetItemCourseProduct(cset.Id, ItemList[i].Shop_id, &ProductSub); errU != nil {
							c.AbortWithStatusJSON(200, gin.H{
								"status":  false,
								"message": "Course Product invalid.",
								"data":    errU.Error(),
							})
							return
						}
						if len(ProductSub) == 0 {
							if cset.Course_lock_drug == 0 {
								if cset.Course_qtyset != 1 {
									cset.Course_amount = cset.Course_qtyset
								}
								CourseSub := structs.ItemCourseSub{
									Id:             cset.Id,
									Course_type_id: cset.Course_type_id,
									Course_code:    cset.Course_code,
									Course_name:    cset.Course_name,
									Course_amount:  cset.Course_list_qtyset,
									Course_unit:    cset.Course_unit,
									Course_cost:    cset.Course_cost,
									Course_opd:     cset.Course_list_opd, // * float64(cset.Course_list_qtyset),
									Course_ipd:     cset.Course_list_ipd, // * float64(cset.Course_list_qtyset),
									Course_image_1: cset.Course_image_1,
									Course_image_2: cset.Course_image_2,
									Course_image_3: cset.Course_image_3,
									Course_image_4: cset.Course_image_4,
									Course_detail:  cset.Course_detail,
									Label:          ItemList[i].Course_name, //"Set",
									Is_set:         1,
									Id_set:         &ItemList[i].Id,
									Products:       []structs.ItemProductSubSet{},
								}
								ItemCourseSub = append(ItemCourseSub, CourseSub)
							}
						} else {
							for ip, productSubList := range ProductSub {
								var objQueryShopStore models.ShopStore
								errShopStore := models.GetShopStoreByIdType1(ItemList[i].Shop_id, &objQueryShopStore)
								if errShopStore != nil {
									c.JSON(http.StatusOK, gin.H{
										"status":  false,
										"message": "Get data shop store error!",
										"data":    errShopStore,
									})
								}

								ProductSub[ip].Balance = 0
								var productStoreBalance structs.ObjQueryProductStoreBalance
								if errT := models.GetProductStoreBalance(objQueryShopStore.ID, productSubList.Product_id, &productStoreBalance); errT != nil {
									c.AbortWithStatusJSON(200, gin.H{
										"status":  false,
										"message": "Product store balance invalid.",
										"data":    errT,
									})
									return
								}
								ProductSub[ip].Balance = productStoreBalance.Pds_balance
							}
							if cset.Course_qtyset != 1 {
								cset.Course_amount = cset.Course_qtyset
							}
							CourseSub := structs.ItemCourseSub{
								Id:             cset.Id,
								Course_type_id: cset.Course_type_id,
								Course_code:    cset.Course_code,
								Course_name:    cset.Course_name,
								Course_amount:  cset.Course_list_qtyset,
								Course_unit:    cset.Course_unit,
								Course_cost:    cset.Course_cost,
								Course_opd:     cset.Course_list_opd, // * float64(cset.Course_list_qtyset),
								Course_ipd:     cset.Course_list_ipd, // * float64(cset.Course_list_qtyset),
								Course_image_1: cset.Course_image_1,
								Course_image_2: cset.Course_image_2,
								Course_image_3: cset.Course_image_3,
								Course_image_4: cset.Course_image_4,
								Course_detail:  cset.Course_detail,
								Label:          ItemList[i].Course_name, //"Set",
								Is_set:         1,
								Id_set:         &ItemList[i].Id,
								Products:       ProductSub,
							}
							ItemCourseSub = append(ItemCourseSub, CourseSub)
						}
					}
					ItemList[i].Subs = ItemCourseSub
				}
			}

			if len(ItemList[i].Subs) > 0 {
				ItemListCourse = append(ItemListCourse, ItemList[i])
			}
		}(st.Course_code, i)
	}
	wg.Wait()

	sort.Slice(ItemListCourse, func(i, j int) bool {
		return ItemListCourse[i].Course_code < ItemListCourse[j].Course_code
	})

	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    ItemListCourse,
	})
}

func ItemCheckingList(c *gin.Context) {
	var filter structs.ObjPayloaItem
	if errSBJ := c.ShouldBindJSON(&filter); errSBJ != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Invalid request data.",
			"data":    errSBJ.Error(),
		})
		return
	}

	var ItemList []structs.ItemChecking
	if errMD := models.GetItemCheckingList(filter, true, &ItemList); errMD != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"status":  false,
			"message": "Something went wrong.",
			"data":    errMD.Error(),
		})
		return
	}

	if len(ItemList) == 0 {
		c.JSON(200, gin.H{
			"status":  true,
			"message": "",
			"data":    []structs.ItemChecking{},
		})
		return
	}

	ItemListChecking := []structs.ItemChecking{}
	var wg sync.WaitGroup
	for i, st := range ItemList {
		wg.Add(1)
		go func(st string, i int) {
			defer wg.Done()

			var objQueryShopStore models.ShopStore
			errShopStore := models.GetShopStoreByIdType1(ItemList[i].Shop_id, &objQueryShopStore)
			if errShopStore != nil {
				c.JSON(http.StatusOK, gin.H{
					"status":  false,
					"message": "Get data shop store error!",
					"data":    errShopStore,
				})
			}

			var ProductSub []structs.ItemProductSubSet
			if errU := models.GetItemCheckingProduct(ItemList[i].Id, ItemList[i].Shop_id, &ProductSub); errU != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status":  false,
					"message": "Checking Product invalid.",
					"data":    errU.Error(),
				})
				return
			}
			for ip, productSubList := range ProductSub {
				ProductSub[ip].Balance = 0
				var productStoreBalance structs.ObjQueryProductStoreBalance
				if errT := models.GetProductStoreBalance(objQueryShopStore.ID, productSubList.Product_id, &productStoreBalance); errT != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Product store balance invalid.",
						"data":    errT,
					})
					return
				}
				ProductSub[ip].Balance = productStoreBalance.Pds_balance
			}
			if ItemList[i].Checking_type_id != 4 {
				ItemList[i].Subs = []structs.ItemCheckingSub{
					{
						Id:               ItemList[i].Id,
						Checking_type_id: ItemList[i].Checking_type_id,
						Checking_code:    ItemList[i].Checking_code,
						Checking_name:    ItemList[i].Checking_name,
						Checking_amount:  ItemList[i].Checking_amount,
						Checking_unit:    ItemList[i].Checking_unit,
						Checking_cost:    ItemList[i].Checking_cost,
						Checking_opd:     ItemList[i].Checking_opd,
						Checking_ipd:     ItemList[i].Checking_ipd,
						Checking_image_1: ItemList[i].Checking_image_1,
						Checking_image_2: ItemList[i].Checking_image_2,
						Checking_image_3: ItemList[i].Checking_image_3,
						Checking_image_4: ItemList[i].Checking_image_4,
						Checking_detail:  ItemList[i].Checking_detail,
						Label:            "",
						Is_set:           0,
						Id_set:           nil,
						Products:         ProductSub,
					},
				}
			} else {
				var CheckingSets []structs.ItemCheckingSub
				ItemCheckingSub := []structs.ItemCheckingSub{}
				if errS := models.GetItemCheckingIdSet(ItemList[i].Id, &CheckingSets); errS != nil {
					c.AbortWithStatusJSON(200, gin.H{
						"status":  false,
						"message": "Checking Set invalid.",
						"data":    errS.Error(),
					})
					return
				}

				if len(CheckingSets) > 0 {
					for _, cset := range CheckingSets {
						var ProductSub []structs.ItemProductSubSet
						if errU := models.GetItemCheckingProduct(cset.Id, ItemList[i].Shop_id, &ProductSub); errU != nil {
							c.AbortWithStatusJSON(200, gin.H{
								"status":  false,
								"message": "Checking Product invalid.",
								"data":    errU.Error(),
							})
							return
						}
						if len(ProductSub) == 0 {
							CheckingSub := structs.ItemCheckingSub{
								Id:               cset.Id,
								Checking_type_id: cset.Checking_type_id,
								Checking_code:    cset.Checking_code,
								Checking_name:    cset.Checking_name,
								Checking_amount:  cset.Checking_amount,
								Checking_unit:    cset.Checking_unit,
								Checking_cost:    cset.Checking_cost,
								Checking_opd:     cset.Checking_list_opd,
								Checking_ipd:     cset.Checking_list_ipd,
								Checking_image_1: cset.Checking_image_1,
								Checking_image_2: cset.Checking_image_2,
								Checking_image_3: cset.Checking_image_3,
								Checking_image_4: cset.Checking_image_4,
								Checking_detail:  cset.Checking_detail,
								Label:            ItemList[i].Checking_name, //"Set",
								Is_set:           1,
								Id_set:           &ItemList[i].Id,
								Products:         []structs.ItemProductSubSet{},
							}
							ItemCheckingSub = append(ItemCheckingSub, CheckingSub)
						} else {
							for ip, productSubList := range ProductSub {

								ProductSub[ip].Balance = 0
								var productStoreBalance structs.ObjQueryProductStoreBalance
								if errT := models.GetProductStoreBalance(objQueryShopStore.ID, productSubList.Product_id, &productStoreBalance); errT != nil {
									c.AbortWithStatusJSON(200, gin.H{
										"status":  false,
										"message": "Product store balance invalid.",
										"data":    errT,
									})
									return
								}
								ProductSub[ip].Balance = productStoreBalance.Pds_balance
							}
							CheckingSub := structs.ItemCheckingSub{
								Id:               cset.Id,
								Checking_type_id: cset.Checking_type_id,
								Checking_code:    cset.Checking_code,
								Checking_name:    cset.Checking_name,
								Checking_amount:  cset.Checking_amount,
								Checking_unit:    cset.Checking_unit,
								Checking_cost:    cset.Checking_cost,
								Checking_opd:     cset.Checking_list_opd,
								Checking_ipd:     cset.Checking_list_ipd,
								Checking_image_1: cset.Checking_image_1,
								Checking_image_2: cset.Checking_image_2,
								Checking_image_3: cset.Checking_image_3,
								Checking_image_4: cset.Checking_image_4,
								Checking_detail:  cset.Checking_detail,
								Label:            ItemList[i].Checking_name, //"Set",
								Is_set:           1,
								Id_set:           &ItemList[i].Id,
								Products:         ProductSub,
							}
							ItemCheckingSub = append(ItemCheckingSub, CheckingSub)
						}
					}
					ItemList[i].Subs = ItemCheckingSub
				} else {
					if len(ProductSub) == 0 {
						ItemList[i].Subs = []structs.ItemCheckingSub{
							{
								Id:               ItemList[i].Id,
								Checking_type_id: ItemList[i].Checking_type_id,
								Checking_code:    ItemList[i].Checking_code,
								Checking_name:    ItemList[i].Checking_name,
								Checking_amount:  ItemList[i].Checking_amount,
								Checking_unit:    ItemList[i].Checking_unit,
								Checking_cost:    ItemList[i].Checking_cost,
								Checking_opd:     ItemList[i].Checking_opd,
								Checking_ipd:     ItemList[i].Checking_ipd,
								Checking_image_1: ItemList[i].Checking_image_1,
								Checking_image_2: ItemList[i].Checking_image_2,
								Checking_image_3: ItemList[i].Checking_image_3,
								Checking_image_4: ItemList[i].Checking_image_4,
								Checking_detail:  ItemList[i].Checking_detail,
								Label:            ItemList[i].Checking_name, //"Set",
								Is_set:           1,
								Id_set:           &ItemList[i].Id,
								Products:         []structs.ItemProductSubSet{},
							},
						}
					} else {
						for ip, productSubList := range ProductSub {
							ProductSub[ip].Balance = 0
							var productStoreBalance structs.ObjQueryProductStoreBalance
							if errT := models.GetProductStoreBalance(objQueryShopStore.ID, productSubList.Product_id, &productStoreBalance); errT != nil {
								c.AbortWithStatusJSON(200, gin.H{
									"status":  false,
									"message": "Product store balance invalid.",
									"data":    errT,
								})
								return
							}
							ProductSub[ip].Balance = productStoreBalance.Pds_balance
						}
						ItemList[i].Subs = []structs.ItemCheckingSub{
							{
								Id:               ItemList[i].Id,
								Checking_type_id: ItemList[i].Checking_type_id,
								Checking_code:    ItemList[i].Checking_code,
								Checking_name:    ItemList[i].Checking_name,
								Checking_amount:  ItemList[i].Checking_amount,
								Checking_unit:    ItemList[i].Checking_unit,
								Checking_cost:    ItemList[i].Checking_cost,
								Checking_opd:     ItemList[i].Checking_opd,
								Checking_ipd:     ItemList[i].Checking_ipd,
								Checking_image_1: ItemList[i].Checking_image_1,
								Checking_image_2: ItemList[i].Checking_image_2,
								Checking_image_3: ItemList[i].Checking_image_3,
								Checking_image_4: ItemList[i].Checking_image_4,
								Checking_detail:  ItemList[i].Checking_detail,
								Label:            ItemList[i].Checking_name, //"Set",
								Is_set:           1,
								Id_set:           &ItemList[i].Id,
								Products:         ProductSub,
							},
						}
					}
				}
			}

			if len(ItemList[i].Subs) > 0 {
				ItemListChecking = append(ItemListChecking, ItemList[i])
			}

		}(st.Checking_code, i)
	}
	wg.Wait()

	sort.Slice(ItemListChecking, func(i, j int) bool {
		return ItemListChecking[i].Checking_code < ItemListChecking[j].Checking_code
	})

	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    ItemListChecking,
	})
}
