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

func Insert(product Product) {
	conn := db.InitConnection()
	stmt, err := conn.Prepare("insert into products (name, description, price, quantity) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	stmt.Exec(product.Name, product.Description, product.Price, product.Quantity)
	defer conn.Close()
}

func Update(product Product) {
	conn := db.InitConnection()
	stmt, err := conn.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	stmt.Exec(product.Name, product.Description, product.Price, product.Quantity, product.Id)
	defer conn.Close()
}

func Delete(id string) {
	conn := db.InitConnection()
	delete, err := conn.Prepare("delete from products where Id=$1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)
	defer conn.Close()
}

func Find(id string) Product {
	conn := db.InitConnection()
	productQuery, err := conn.Prepare("select * from products where Id=$1")
	if err != nil {
		panic(err.Error())
	}
	product, err := productQuery.Query(id)
	if err != nil {
		panic(err.Error())
	}
	p := Product{}
	for product.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = product.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Id = id
		p.Quantity = quantity
		p.Price = price
	}
	defer conn.Close()
	return p
}
