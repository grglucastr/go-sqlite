package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	// Migrate the schema
	db.AutoMigrate(&Product{})

	if err != nil {
		panic("Failed to connect database")
	}

	keepRunning := true

	for keepRunning {

		menu := "\n------- MENU -------\n\n1 - Insert\n2 - Search\n3 - Update\n4 - Delete\n5 - Exit\nChoose option: "

		reader := bufio.NewReader(os.Stdin)
		opt, _ := getInput(menu, reader)

		fmt.Println()

		switch opt {
		case "1":
			processInsert(db)
		case "2":
			processSearch(db)
		case "3":
			processUpdate(db)
		case "4":
			processDelete(db)
		case "5":
			keepRunning = false
		}
	}
}

func processInsert(db *gorm.DB) {
	reader := bufio.NewReader(os.Stdin)
	code, _ := getInput("Product Code: ", reader)
	p, _ := getInput("Product Price: ", reader)

	price, _ := strconv.ParseUint(p, 10, 32)

	pro := &Product{
		Code:  code,
		Price: uint(price),
	}

	db.Create(pro)
}

func processSearch(db *gorm.DB) *Product {
	reader := bufio.NewReader(os.Stdin)
	code, _ := getInput("Product Code: ", reader)

	var product Product
	db.First(&product, "code=?", code)

	printProduct(&product)

	return &product
}

func processUpdate(db *gorm.DB) {

	product := processSearch(db)

	if product == nil {
		return
	}

	reader := bufio.NewReader(os.Stdin)
	cm := "Product Code (" + product.Code + "): "
	code, _ := getInput(cm, reader)

	pm := "Product Price (" + strconv.FormatUint(uint64(product.Price), 10) + "): "
	price, _ := getInput(pm, reader)

	pr, _ := strconv.ParseUint(price, 10, 32)
	prc := uint(pr)

	db.Model(&product).Updates(Product{Code: code, Price: prc})

	fmt.Println("Product updated.")
}

func processDelete(db *gorm.DB) {
	product := processSearch(db)

	if product == nil {
		return
	}

	printProduct(product)

	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Proceed delete? (s/N): ", reader)

	fmt.Println("opt: ", opt)

	if opt == "" {
		fmt.Println("Cancelled.")
		return
	}

	db.Delete(product, product.ID)
	fmt.Print("Deleted.")
}

func printProduct(product *Product) {
	code := product.Code
	price := strconv.FormatUint(uint64(product.Price), 10)

	fmt.Println("Code: " + code)
	fmt.Println("Price: " + price)
	fmt.Println()
}
