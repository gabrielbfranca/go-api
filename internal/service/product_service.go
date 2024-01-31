package service

import (
	"github.com/devfullcycle/imersao17/goapi/internal/database"
	"github.com/devfullcycle/imersao17/goapi/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(productDB database.ProductDB) *ProductService {
	return &ProductService{ProductDB: productDB}
}

func (ps *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := ps.ProductDB.GetProduct(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) GetProductByCategoryID(id string) ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProductByCategoryID(id)
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (ps *ProductService) CreateProduct(name, description, category_id string, image_url string, price float64) (*entity.Product, error) {
	product := entity.NewProduct(name, description, price, category_id, image_url)
	_, err := ps.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
