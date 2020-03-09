package service

import (
	"bloger/dao/db"
	"bloger/model"
)

//获取所有分类
func GetALLCategoryList() (categoryList []*model.Category,err error) {
	categoryList, err = db.GetAllCategoryList()
	if err != nil {
		return
	}
	return


}







