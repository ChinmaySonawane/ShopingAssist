package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"productReviewer/database"
	"productReviewer/database/product"
	"productReviewer/database/relation"
	"productReviewer/database/reviews"
	"strconv"
)

var db *sql.DB

//database initilization
func init() {
	d, err := database.Connect()
	if err != nil {
		fmt.Println("Cant cnnect")
		os.Exit(0)
	}
	db = d
	fmt.Println("conneted")
}

func main() {
	//close connnection
	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Println("Error while closing")
		}
		fmt.Println("closing")
	}()

	//hundle func
	r := mux.NewRouter()
	r.HandleFunc("/", ping)
	r.HandleFunc("/product", GetProducts).Methods("GET")
	r.HandleFunc("/product", PostProduct).Methods("POST")
	r.HandleFunc("/product/{id}", GetProduct).Methods("GET")
	r.HandleFunc("/product/{id}", DeleteProduct).Methods("DELETE")
	r.HandleFunc("/reviews", GetReviews).Methods("GET")
	r.HandleFunc("/reviews", PostReview).Methods("POST")
	r.HandleFunc("/reviews/{id}", GetReview).Methods("GET")
	r.HandleFunc("/reviews/{id}", DeleteReview).Methods("DELETE")
	r.HandleFunc("/relation", GetRelations).Methods("GET")

	http.Handle("/", r)
	fmt.Println("Starting up on 8080")
	http.ListenAndServe(":8080", nil)

	err := db.Ping()
	if err != nil {
		fmt.Println("Cannot ping db: ", err)
		os.Exit(0)
	}

	/*err = product.Update(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Update")*/
}

func ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "PONG")
}

func GetProducts(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "products list")
	//fmt.Println("get products")
	prod, err := product.List(db)
	if err != nil {
		fmt.Fprintln(w, "error while product listing")
	}
	for _, p := range prod {
		fmt.Fprintln(w, p)
	}
}

func GetRelations(w http.ResponseWriter, req *http.Request) {
	//fmt.Fprintln(w, "get all")
	//fmt.Println("get all")
	prod, err := relation.List(db)
	if err != nil {
		fmt.Fprintln(w, "error while product listing")
	}
	for _, p := range prod {
		fmt.Fprintln(w, "Product")
		fmt.Fprintln(w, p.Prod)
		for _, r := range p.Review {
			fmt.Fprintln(w, "Reviews")
			fmt.Fprintln(w, r)
		}
	}
}

func PostProduct(w http.ResponseWriter, req *http.Request) {
	//fmt.Println("post product")
	var prod product.Product
	/*b := Byte(req.Body)
	for key, _ := range req.Form {
		//log.Println(key)
		err := json.Unmarshal([]byte(key), &prod)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(prod)
		err = product.Insert(db, prod)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("inserted")
	}
	*/
	err := json.NewDecoder(req.Body).Decode(&prod)
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println(prod)
	err = product.Insert(db, prod)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintln(w, "inserted successfully")
}

//to parse id from url or request(url in request)
func getId(req *http.Request) (int, error) {
	vars := mux.Vars(req)
	id := vars["id"]
	fmt.Println(id)
	//fmt.Fprintln(w, id)
	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return i, nil
}

func GetProduct(w http.ResponseWriter, req *http.Request) {
	id, err := getId(req)
	if err != nil {
		fmt.Fprintln(w, "Error:", err)
		fmt.Println("Error:", err)
		return
	}
	prod, err := product.Select(db, id)
	if err != nil {
		fmt.Fprintln(w, "Error:", err)
		fmt.Println("Error:", err)
		return
	}
	fmt.Fprintln(w, prod)
}

func DeleteProduct(w http.ResponseWriter, req *http.Request) {
	id, err := getId(req)
	if err != nil {
		fmt.Fprintln(w, "Error:", err)
		fmt.Println("Error:", err)
		return
	}

	err = reviews.DeleteP(db, id)
	if err != nil {
		fmt.Fprintln(w, "Error:", err)
		fmt.Println("Error:", err)
		return
	}

	err = product.Delete(db, id)
	if err != nil {
		fmt.Fprintln(w, "Error:", err)
		fmt.Println("Error:", err)
		return
	}
	fmt.Fprintln(w, "delete successfully")
}

func GetReviews(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Reviews:\t")
	//fmt.Println("get ")
	review, err := reviews.List(db)
	if err != nil {
		fmt.Fprintln(w, "Error:", err)
		fmt.Println("Error:", err)
		return
	}
	for _, r := range review {
		fmt.Fprintln(w, r)
	}
}

func PostReview(w http.ResponseWriter, req *http.Request) {
	//fmt.Println("post reviews")
	var review reviews.Review
	/*b := Byte(req.Body)
	for key, _ := range req.Form {
		//log.Println(key)
		err := json.Unmarshal([]byte(key), &prod)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(prod)
		err = product.Insert(db, prod)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("inserted")
	}
	*/
	err := json.NewDecoder(req.Body).Decode(&review)
	if err != nil {
		fmt.Fprintln(w, "Error:", err)
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(review)
	err = reviews.Insert(db, review)
	if err != nil {
		fmt.Fprintln(w, "Error:", err)
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("reviews inserted successfully")
}

func GetReview(w http.ResponseWriter, req *http.Request) {
	id, err := getId(req)
	if err != nil {
		fmt.Fprintln(w, "Error:", err)
		fmt.Println("Error:", err)
		return
	}
	review, err := reviews.Select(db, id)
	if err != nil {
		fmt.Fprintln(w, "Error:", err)
		fmt.Println("Error:", err)
		return
	}
	fmt.Fprintln(w, "Reviews", review)
}

func DeleteReview(w http.ResponseWriter, req *http.Request) {
	id, err := getId(req)
	if err != nil {
		fmt.Fprintln(w, "Error:", err)
		fmt.Println("Error:", err)
		return
	}

	err = reviews.Delete(db, id)
	if err != nil {
		fmt.Fprintln(w, "Error:", err)
		fmt.Println("Error:", err)
		return
	}

	fmt.Fprintln(w, "Review delete successfully")
}
