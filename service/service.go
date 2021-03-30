package service

import (
	"fmt"

	"product/config"
	"product/model"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllAlbums - fetch all Albums
func GetAllProducts(product *[]model.Register) (err error) {
	if err = config.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

//CreateAlbum - creates an album
func CreateProduct(product *model.Register) (err error) {
	if err = config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

//GetAlbum fetch one alb um
func GetProduct(product *model.Register, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}
func GetUserByName(user *model.Register, name string) (err error) {
	if err := config.DB.Where("first_name = ?", name).First(user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateAlbum - modifies an album
func UpdateProduct(product *model.Register, id string) (err error) {
	fmt.Println(product)
	config.DB.Save(product)
	return nil
}

//DeleteAlbum deletes a given album name
func DeleteProduct(product *model.Register, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(product)
	return nil
}
