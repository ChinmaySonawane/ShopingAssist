package relation

import (
	"database/sql"
	"fmt"
	"productReviewer/database/product"
	"productReviewer/database/reviews"
)

type Relation struct {
	Prod   product.Product
	Review []reviews.Review
}

func List(db *sql.DB) ([]Relation, error) {
	relas := make([]Relation, 0)
	sqlStatement := `SELECT * FROM product`
	row, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println("1")
		return relas, err
	}
	reviewsql := `
	SELECT * FROM reviews 
	WHERE product_id=$1`
	for row.Next() {
		var rela Relation
		err = row.Scan(&rela.Prod.ProductID, &rela.Prod.Name, &rela.Prod.Price, &rela.Prod.ProductType, &rela.Prod.Rating)
		if err != nil {
			fmt.Println("2")
			return relas, err
		}
		row, err := db.Query(reviewsql, rela.Prod.ProductID)
		if err != nil {
			fmt.Println("3")
			return relas, err
		}
		//relas.Review = make([]reviews.Review, 0)
		for row.Next() {
			var r reviews.Review
			err := row.Scan(&r.ReviewID, &r.ReviewerName, &r.Review, &r.Rating, &r.ProductID)
			if err != nil {
				fmt.Println("4")
				return relas, err
			}
			//fmt.Println(review)
			rela.Review = append(rela.Review, r)
		}
		//fmt.Println(review)
		relas = append(relas, rela)
	}
	return relas, nil
}
