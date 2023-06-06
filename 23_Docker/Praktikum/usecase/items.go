package usecase

import (
	"errors"
	"fmt"
	"project_structure/model"
	"project_structure/repository/database"
)

type ProductUsecase interface {
	CreateItems(items *model.Order) error
	GetItem(id uint) (items model.Product, err error)
	GetListItems() (items []model.Product, err error)
	UpdateItems(items *model.Product) (err error)
	DeleteItems(id uint) (err error)
}

func CreateItem(name, description string, price, stock uint) (*model.Item, error) {
	items, err := database.GetItems()
	if err != nil {
		return nil, err
	}
	for _, v := range items {
		if v.Name == name {
			return nil, errors.New("item already exists")
		}
	}
	item := &model.Item{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
	}

	err = database.CreateItem(item)
	if err != nil {
		fmt.Println("error creating item")
		return nil, err
	}
	return item, nil
}

func GetItem(id uint) (item model.Item, err error) {
	item, err = database.GetItem(id)
	if err != nil {
		fmt.Println("GetItem: Error getting item from database")
		return
	}
	return
}

func GetListItems() (items []model.Item, err error) {
	items, err = database.GetItems()
	if err != nil {
		fmt.Println("GetListItems: Error getting items from database")
		return
	}
	return
}

func UpdateItem(item *model.Item, stock uint) (err error) {
	item.Stock = stock

	err = database.UpdateItem(item)
	if err != nil {
		fmt.Println("UpdateItem: Error updating item, err:", err)
		return
	}

	return
}

func DeleteItem(id uint) (err error) {
	item := model.Item{}
	item.ID = id
	err = database.DeleteItem(&item)
	if err != nil {
		fmt.Println("DeleteItem: error deleting item, err:", err)
		return
	}

	return
}
