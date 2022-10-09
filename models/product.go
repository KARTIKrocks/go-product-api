package models

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Stock int     `json:"stock"`
	Price float32 `json:"price"`
}

// CRUD

func CreateProduct(products *[]Product, newProduct *Product) (err error) {
	*products = append(*products, *newProduct)

	return nil
}
func ReadProductById(products *[]Product, id int) (product Product) {
	for _, item := range *products {
		if item.Id == id {
			return item
		}
	}
	return Product{}
}
func UpdateProductById(products *[]Product, updateProduct *Product) (product Product) {
	for index, item := range *products {
		if item.Id == (*updateProduct).Id {
			(*products)[index].Name = (*updateProduct).Name
			(*products)[index].Stock = (*updateProduct).Stock
			(*products)[index].Price = (*updateProduct).Price
			return (*products)[index]
		}
	}
	return Product{}
}
func DeleteProductById(products *[]Product, id int) (product Product) {
	var deletedProduct Product

	for index, item := range *products {
		if item.Id == id {
			n := len(*products)
			deletedProduct = item
			(*products)[index] = (*products)[n-1]
			*products = (*products)[:n-1]

			return deletedProduct
		}
	}
	return Product{}
}
