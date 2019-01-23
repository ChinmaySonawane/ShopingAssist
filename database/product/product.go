package product

import (
	"database/sql"
	_ "github.com/derekparker/delve/pkg/proc"
	_ "github.com/lib/pq"
)

//product struct
type Product struct {
	ProductID   int     `json:"pid"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	ProductType string  `json:"type"`
	Rating      int     `json:"rating"`
}

//insert for product
func Insert(db *sql.DB, proc Product) error {
	//fmt.Println("In inset")
	sqlStatement := `
	INSERT INTO product (name , price, type, rating)
	VALUES ($1, $2, $3, $4) RETURNING product_id`
	id := 0
	err := db.QueryRow(sqlStatement, proc.Name, proc.Price, proc.ProductType, proc.Rating).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

//selecting perticular product
func Select(db *sql.DB, id int) (Product, error) {
	var prod Product
	sqlStatement := `
	SELECT * FROM product 
	WHERE product_id=$1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&prod.ProductID, &prod.Name, &prod.Price, &prod.ProductType, &prod.Rating)
	if err != nil {
		return prod, err
	}
	return prod, nil
}

//selecting all products
func List(db *sql.DB) ([]Product, error) {
	prods := make([]Product, 0)
	sqlStatement := `SELECT * FROM product`
	row, err := db.Query(sqlStatement)
	if err != nil {
		return prods, err
	}

	for row.Next() {
		var prod Product
		err = row.Scan(&prod.ProductID, &prod.Name, &prod.Price, &prod.ProductType, &prod.Rating)
		if err != nil {
			return prods, err
		}
		prods = append(prods, prod)
	}
	return prods, nil
}

//delete particular product
func Delete(db *sql.DB, id int) error {
	sqlStatement := `DELETE FROM product WHERE product_id = $1;`
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	return nil
}

/*func Update(db *sql.DB) error {
	sqlStatement := `UPDATE product
	SET rating = $2
	WHERE product_id = $1`
	_, err := db.Exec(sqlStatement, 1, "type", 1)
	if err != nil {
		return err
	}
	return nil
}*/
