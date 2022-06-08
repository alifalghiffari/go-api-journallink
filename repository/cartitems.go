package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type CartItemRepository struct {
	db *sql.DB
}

func NewCartItemRepository(db *sql.DB) *CartItemRepository {
	return &CartItemRepository{db: db}
}

func (c *CartItemRepository) FetchCartItems() ([]CartItem, error) {
	var sqlStatement string
	var cartItems []CartItem

	//TODO: add sql statement here
	//HINT: join table cart_items and products
	sqlStatement = `
		SELECT c.id, p.category, c.product_id, p.product_name, c.quantity, p.price
		FROM cart_items c
		JOIN products p ON c.product_id = p.id
	`

	rows, err := c.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	var cartItem CartItem
	for rows.Next() {
		err := rows.Scan(
			&cartItem.ID,
			&cartItem.Category,
			&cartItem.ProductID,
			&cartItem.ProductName,
			&cartItem.Quantity,
			&cartItem.Price,
		)
		if err != nil {
			return nil, err
		}
		cartItems = append(cartItems, cartItem)
	}

	return cartItems, nil
}

func (c *CartItemRepository) FetchCartByProductID(productID int64) (CartItem, error) {
	var cartItem CartItem
	var sqlStatement string
	//TODO : you must fetch the cart by product id
	//HINT : you can use the where statement
	sqlStatement = `SELECT c.id, p.category, c.product_id, p.product_name, c.quantity 
	FROM cart_items c
	JOIN products p ON c.product_id = p.id
	WHERE c.product_id = ?
	LIMIT 1;`

	row := c.db.QueryRow(sqlStatement, productID)
	err := row.Scan(
		&cartItem.ID,
		&cartItem.Category,
		&cartItem.ProductID,
		&cartItem.ProductName,
		&cartItem.Quantity,
	)
	if err != nil {
		return cartItem, err
	}

	return cartItem, nil
}

func (c *CartItemRepository) InsertCartItem(cartItem CartItem) error {
	// TODO: you must insert the cart item
	var sqlStatement string

	sqlStatement = `INSERT INTO cart_items (product_id, quantity) VALUES (?, ?);`

	_, err := c.db.Exec(sqlStatement, cartItem.ProductID, cartItem.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (c *CartItemRepository) IncrementCartItemQuantity(cartItem CartItem) error {
	//TODO : you must update the quantity of the cart item
	var sqlStatement string

	sqlStatement = `UPDATE cart_items SET quantity = quantity + 1 WHERE id = ?;`

	_, err := c.db.Exec(sqlStatement, cartItem.ID)
	if err != nil {
		return err
	}

	return nil
}

func (c *CartItemRepository) ResetCartItems() error {
	//TODO : you must reset the cart items
	//HINT : you can use the delete statement
	var sqlStatement string

	sqlStatement = `DELETE FROM cart_items;`

	_, err := c.db.Exec(sqlStatement)
	if err != nil {
		return err
	}

	return nil
}

func (c *CartItemRepository) TotalPrice() (int, error) {
	var sqlStatement string
	//TODO : you must calculate the total price of the cart items
	//HINT : you can use the sum statement

	sqlStatement = `SELECT SUM(p.price * c.quantity) 
	FROM cart_items c
	JOIN products p ON c.product_id = p.id;`

	var totalPrice int
	row := c.db.QueryRow(sqlStatement)
	err := row.Scan(&totalPrice)
	if err != nil {
		return 0, err
	}

	return totalPrice, nil
}
