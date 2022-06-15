package repository

import "time"

type User struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Email   string `db:"email"`
	Password string `db:"password"`
	Role     string `db:"role"`
	Created_at time.Time `db:"created_at"`
	Updated_at time.Time `db:"updated_at"`
	Token    string `db:"token"`
}

type Journal struct {
	ID       int64  `db:"id"`
	UserID   int64  `db:"user_id"`
	Isi      string `db:"isi"`
	Status   string `db:"status"`
	DateSubmit time.Time `db:"date_submit"`
	Created_at time.Time `db:"created_at"`
	Updated_at time.Time `db:"updated_at"`
}

// type Product struct {
// 	ID          int64  `db:"id"`
// 	Category    string `db:"category"`
// 	ProductName string `db:"product_name"`
// 	Price       int    `db:"price"`
// 	Quantity    int    `db:"quantity"`
// }

// type CartItem struct {
// 	ID          int64  `db:"id"`
// 	Category    string `db:"category"`
// 	ProductID   int64  `db:"product_id"`
// 	ProductName string `db:"product_name"`
// 	Price       int    `db:"price"`
// 	Quantity    int    `db:"quantity"`
// }

// type Sales struct {
// 	ID          int64     `db:"id"`
// 	Date        time.Time `db:"date"`
// 	Category    string    `db:"category"`
// 	ProductID   int64     `db:"product_id"`
// 	ProductName string    `db:"product_name"`
// 	Price       int       `db:"price"`
// 	Quantity    int       `db:"quantity"`
// 	Total       int       `db:"total"`
// }

// type GetSalesRequest struct {
// 	StartPeriod *time.Time `db:"start_period"`
// 	EndPeriod   *time.Time `db:"end_period"`
// 	ProductID   int64      `db:"product_id"`
// 	ProductName string     `db:"product_name"`
// }

// func (r GetSalesRequest) IsEmptyRequest() bool {
// 	if r.StartPeriod == nil && r.EndPeriod == nil && r.ProductName == "" {
// 		return true
// 	}

// 	return false
// }

// func (r GetSalesRequest) IsValidRequest() bool {
// 	if r.StartPeriod == nil && r.EndPeriod != nil {
// 		return false
// 	}
// 	if r.StartPeriod != nil && r.EndPeriod == nil {
// 		return false
// 	}

// 	return true
// }
