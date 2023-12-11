package models

import "github.com/beriloqueiroz/simple-go-catalog/db"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func ListAll() []Product {
	conn := db.InitConnection()
	allProducts, err := conn.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}
	p := Product{}
	products := []Product{}

	for allProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = allProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Id = id
		p.Quantity = quantity
		p.Price = price

		products = append(products, p)
	}
	defer conn.Close()
	return products
}
