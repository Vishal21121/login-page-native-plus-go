package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Vishal21121/login-page-native-plus-go/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

const mongoUri = "mongodb://127.0.0.1:27017"
const dbName = "nativeUser"
const collectionName = "User"

var collection *mongo.Collection

func Init() {
	clientOption := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The mongodb is connected")
	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection Instance is created")

}

func insertOneUser(user model.User) bool {
	var userGot bson.M
	filter := bson.M{"email": user.Email}
	result := collection.FindOne(context.Background(), filter)
	fmt.Println(result)
	_ = result.Decode(&userGot)
	fmt.Println("the found user is: ", userGot)
	if userGot == nil {
		password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			log.Fatal(err)
		}
		user.Password = string(password)
		inserted, err := collection.InsertOne(context.Background(), user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted with id", inserted.InsertedID)
		return true
	}
	return false
}

func Createuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var user model.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	val := insertOneUser(user)
	if val {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"status": "201", "success": "true"})
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode("A user with this email already exits")
}
