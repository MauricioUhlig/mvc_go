package models

import (
	"errors"

	"github.com/MauricioUhlig/mvc_go/db"
)

type Product struct {
	Id                int
	Name, Description string
	Quantity          int
	Price             float32
}

func (p *Product) SetPrice(price float32) error {
	if price < 0 {
		return errors.New("price must be >= 0")
	}
	p.Price = price
	return nil
}

func (p *Product) GetPrice() float32 {
	return p.Price
}

func GetProducts() []Product {
	db := db.ConnectToDB()
	rows, err := db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}
	var result []Product

	for rows.Next() {
		var id, quantity int
		var name, description string
		var price float32
		rows.Scan(&id, &name, &description, &quantity, &price)

		product := Product{id, name, description, quantity, price}
		result = append(result, product)
	}
	defer db.Close()
	return result
}

func InsertProduct(name, description string, quantity int, price float32) {
	db := db.ConnectToDB()
	insert, err := db.Prepare("insert into products (name, description, quantity, price) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(name, description, quantity, price)
	defer db.Close()
}

func DeleteProduct(id int) {
	db := db.ConnectToDB()

	delete, err := db.Prepare("delete from products where id = $1")
	if err != nil {
		panic(err.Error())
	}
	delete.Exec(id)
	defer db.Close()
}

func GetProduct(id int) Product {
	db := db.ConnectToDB()
	rows := db.QueryRow("select * from products where id = $1", id)

	var quantity int
	var name, description string
	var price float32
	rows.Scan(&id, &name, &description, &quantity, &price)

	product := Product{id, name, description, quantity, price}

	defer db.Close()
	return product

}

func UpdateProduct(id int, name, description string, quantity int, price float32) {
	db := db.ConnectToDB()
	insert, err := db.Prepare("update products set name = $2, description = $3, quantity = $4, price = $5 where id = $1")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(id, name, description, quantity, price)
	defer db.Close()
}
