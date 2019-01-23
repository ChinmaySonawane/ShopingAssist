package reviews

import (
	"database/sql"
	"fmt"
)

type Review struct {
	ReviewID     int    `json:"rid"`
	ReviewerName string `json:"name"`
	Review       string `json:"text"`
	Rating       int    `json:"rating"`
	ProductID    int    `json:"pid"`
}

func Insert(db *sql.DB, review Review) error {
	sqlStatement := `
	INSERT INTO reviews ( reviewer_name, review, rating, product_id)
	VALUES ($1, $2, $3, $4) RETURNING review_id`
	id := 0
	err := db.QueryRow(sqlStatement, review.ReviewerName, review.Review, review.Rating, review.ProductID).Scan(&id)
	if err != nil {
		return err
	}
	fmt.Println("insert in reviews")
	return nil
}

func Select(db *sql.DB, id int) (Review, error) {
	var review Review
	sqlStatement := `
	SELECT * FROM reviews 
	WHERE review_id=$1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&review.ReviewID, &review.ReviewerName, &review.Review, &review.Rating, &review.ProductID)
	if err != nil {
		return review, err
	}
	fmt.Println(review)
	return review, nil
}

func List(db *sql.DB) ([]Review, error) {
	review := make([]Review, 0)
	sqlStatement := `SELECT * FROM reviews`
	row, err := db.Query(sqlStatement)
	if err != nil {
		return review, err
	}
	for row.Next() {
		var r Review
		err := row.Scan(&r.ReviewID, &r.ReviewerName, &r.Review, &r.Rating, &r.ProductID)
		if err != nil {
			return review, err
		}
		//fmt.Println(review)
		review = append(review, r)
	}
	return review, nil
}

func DeleteP(db *sql.DB, no int) error {
	sqlStatement := `DELETE FROM reviews WHERE product_id = $1;`
	_, err := db.Exec(sqlStatement, no)
	if err != nil {
		return err
	}
	return nil
}

func Delete(db *sql.DB, no int) error {
	sqlStatement := `DELETE FROM reviews WHERE review_id = $1;`
	_, err := db.Exec(sqlStatement, no)
	if err != nil {
		return err
	}
	return nil
}
