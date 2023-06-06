package database

import (
	"project_structure/config"
	"project_structure/model"
)

func CreateItem(item *model.Item) error {
	if err := config.DB.Create(item).Error; err != nil {
		return err
	}
	return nil
}

func GetItems() (items []model.Item, err error) {
	if err = config.DB.Find(&items).Error; err != nil {
		return
	}
	return
}

func GetItem(id uint) (item model.Item, err error) {
	item.ID = id
	if err = config.DB.First(&item).Error; err != nil {
		return
	}
	return
}

func UpdateItem(item *model.Item) error {
	if err := config.DB.Updates(item).Error; err != nil {
		return err
	}
	return nil
}

func DeleteItem(item *model.Item) error {
	if err := config.DB.Delete(item).Error; err != nil {
		return err
	}
	return nil
}
