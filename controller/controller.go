package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	model "github.com/anshukr21/blogapp/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://kranshu:Anish321@cluster0.lckjoq5.mongodb.net/?retryWrites=true&w=majority"
const dbName = "blog"
const colName = "page"

var collection *mongo.Collection

// runs at start
func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb //using TODO is preferrable
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")
}

// helpers

// insert 1 blogpost
func insertOneblogpost(blogpost model.Blog) {
	inserted, err := collection.InsertOne(context.Background(), blogpost)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 blogpost in db with id: ", inserted.InsertedID)
}

// update 1 blogpost
func updateOneblogpost(blogpostId string, blogpost model.Blog) {
	id, _ := primitive.ObjectIDFromHex(blogpostId)
	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	Updated, err := collection.InsertOne(context.Background(), blogpost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated 1 blogpost in db with id: ", Updated.InsertedID)
}

// delete 1 blogpost
func deleteOneblogpost(blogpostId string) {
	id, _ := primitive.ObjectIDFromHex(blogpostId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("blogpost got delete with delete count: ", deleteCount)
}

// delete all blogposts from mongodb
func deleteAllblogpost() int64 {

	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of blogposts delete: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all blogposts from database
func getAllblogposts() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var blogposts []primitive.M

	for cur.Next(context.Background()) {
		var blogpost bson.M
		err := cur.Decode(&blogpost)
		if err != nil {
			log.Fatal(err)
		}
		blogposts = append(blogposts, blogpost)
	}

	defer cur.Close(context.Background())
	return blogposts
}

func GetAllBlogposts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allblogposts := getAllblogposts()
	json.NewEncoder(w).Encode(allblogposts)
}

func CreateBlogpost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var blogpost model.Blog
	_ = json.NewDecoder(r.Body).Decode(&blogpost)
	insertOneblogpost(blogpost)
	json.NewEncoder(w).Encode(blogpost)
}

func UpdateBlogpost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	var blogpost model.Blog
	params := mux.Vars(r)
	_ = json.NewDecoder(r.Body).Decode(&blogpost)
	updateOneblogpost(params["id"], blogpost)
	json.NewEncoder(w).Encode(blogpost)
}

func DeleteBlogpost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneblogpost(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllBlogposts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllblogpost()
	json.NewEncoder(w).Encode(count)
}
