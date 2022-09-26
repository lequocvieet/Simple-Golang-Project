package routes

import (
	"context"
	"fmt"
	"golang_standard/contracts"
	"math/big"
	"net/http"
	_ "reflect"

	jwt "golang_standard/auth"
	db "golang_standard/database"
	"golang_standard/models"
	res "golang_standard/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllAlbums(w http.ResponseWriter, r *http.Request) {
	collection := db.ConnectAlbums()

	var result []bson.M
	data, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}

	defer data.Close(context.Background())
	for data.Next(context.Background()) {
		var elem bson.M
		err := data.Decode(&elem)

		if err != nil {
			res.JSON(w, 500, "Internal Server Error")
			return
		}

		result = append(result, elem)
	}

	res.JSON(w, 200, result)
}

func GetOnBlockChain(w http.ResponseWriter, r *http.Request) {
	//Get album onchain by index
	index := r.URL.Query().Get("index")
	fmt.Printf("index:%v\n", index)
	conn, _ := contracts.GetContract()
	n := new(big.Int)
	indexs, _ := n.SetString(index, 10)

	result, err := conn.GetInfo(&bind.CallOpts{}, indexs)

	if err != nil {
		res.JSON(w, 500, "Get album on chain has failed")
		return
	}

	res.JSON(w, 201, result)

}

func CreateAlbum(w http.ResponseWriter, r *http.Request) {
	creater, err := jwt.ExtractUsernameFromToken(r)

	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}

	title := r.PostFormValue("title")
	price := r.PostFormValue("price")

	if govalidator.IsNull(title) {
		res.JSON(w, 400, "Data can not empty")
		return
	}
	if govalidator.IsNull(price) {
		res.JSON(w, 400, "Data can not empty")
		return
	}

	title = models.Santize(title)
	uid := uuid.NewV4()

	id := fmt.Sprintf("%x-%x-%x-%x-%x", uid[0:4], uid[4:6], uid[6:8], uid[8:10], uid[10:])
	collection := db.ConnectAlbums()

	newAlbum := bson.M{"id": id, "price": price, "creater": creater, "title": title}

	_, errs := collection.InsertOne(context.TODO(), newAlbum)

	//create transaction on blockchain
	conn, auth := contracts.GetContract()
	n := new(big.Int)
	prices, _ := n.SetString(price, 10)
	result, err := conn.SetInfo(auth, prices, title, creater)

	if errs != nil {
		res.JSON(w, 500, "Create album has failed")
		return
	}

	res.JSON(w, 201, result)

}

/*
func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	id := ps.ByName("id")
	username, err := jwt.ExtractUsernameFromToken(r)
	collection := db.ConnectAlbums()

	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}

	var result bson.M
	errFind := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&result)

	if errFind != nil {
		res.JSON(w, 404, "Post Not Found")
		return
	}

	creater := fmt.Sprintf("%v", result["creater"])

	if username != creater {
		res.JSON(w, 403, "Permission Denied")
		return
	}

	errDelete := collection.FindOneAndDelete(context.TODO(), bson.M{"id": id}).Decode(&result)

	if errDelete != nil {
		res.JSON(w, 500, "Delete has failed")
		return
	}

	res.JSON(w, 200, "Delete Successfully")

}
*/
