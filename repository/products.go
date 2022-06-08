package repository

import "database/sql"

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) FetchProductByID(id int64) (Product, error) {
	//TODO: You must implement this function fot fetch product by id
	var sqlStatement string
	var product Product

	sqlStatement = `SELECT category, product_name, price, quantity FROM products WHERE id = ?;`

	row := p.db.QueryRow(sqlStatement, id)
	product.ID = id
	err := row.Scan(
		&product.Category,
		&product.ProductName,
		&product.Price,
		&product.Quantity,
	)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (p *ProductRepository) FetchProductByName(productName string) (Product, error) {
	// TODO: You must implement this function for fetch product by name
	var sqlStatement string
	var product Product

	sqlStatement = `SELECT id, category, price, quantity FROM products WHERE product_name = ?;`

	row := p.db.QueryRow(sqlStatement, productName)
	product.ProductName = productName
	err := row.Scan(
		&product.ID,
		&product.Category,
		&product.Price,
		&product.Quantity,
	)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (p *ProductRepository) FetchProducts() ([]Product, error) {
	// TODO: You must implement this function for fetch all products
	var sqlStatement string
	var products []Product

	sqlStatement = `SELECT id, category, product_name, price, quantity FROM products`

	rows, err := p.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	var product Product
	for rows.Next() {
		err := rows.Scan(
			&product.ID,
			&product.Category,
			&product.ProductName,
			&product.Price,
			&product.Quantity,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
