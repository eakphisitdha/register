package main

import (
	//"fmt"
	//gin
	"net/http"

	"github.com/gin-gonic/gin"
	//MongoDB
	//"context"
	//"log"
	//"time"
	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
)

//var collection *mongo.Collection

type User struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
}

var Users = []User{
	{
		Username: "Gantong",
		Phone:    "0970270138",
	},
	{
		Username: "Plam",
		Phone:    "0912345678",
	},
}

func getusers(c *gin.Context) {
	c.JSON(http.StatusOK, Users)
}

/*func insertuser(newuser User){
	inserted, err := collection.InsertOne(context.Background(), newuser)

	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println("Insert a newuser in DB : ",inserted)
}*/

func main() {

	// connectionURI := "mongodb://<host_ip>:<host_port>"
	//connectionURI := "mongodb://192.168.1.83:1"
	//client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//err = client.Connect(ctx)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer client.Disconnect(ctx)

	//err = client.Ping(ctx, readpref.Primary())
	//if err != nil {
	//	log.Fatal(err)
	//}

	//fmt.Println("Connected")

	//RESTful API
	r := gin.Default()

	/*r.GET("/app", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"username": "Gantong", "phone": "0970270138"})
	})*/

	r.GET("/users", getusers)

	r.POST("/register", func(c *gin.Context) {
		var newUser User
		// เรียก BindJSON เพื่อผูก JSON ที่รับมากับ newUser
		if err := c.BindJSON(&newUser); err != nil {
			return
		}
		// เพิ่ม User ใหม่เข้าไปใน slice
		Users = append(Users, newUser)
		c.JSON(http.StatusCreated, newUser)
		//เพิ่ม User ใหม่เข้าไปใน MongoDB
		//coll := client.Database("user").Collection("user")
		//doc := bson.D{{"Username", newUser.Username}, {"Phone", newUser.Phone}}
		//result, _ := coll.InsertOne(context.TODO(), doc)
		//fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	})

	r.Run() //change port here r.Run(:8011)
}
