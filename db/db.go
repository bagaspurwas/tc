package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"models"
)

type Block models.Block

type DB struct {
	Server string
	Database string
	Client mongo.Client
}

const (
	COLLECTION = "default"
)

func (db *DB) Connect(username string, password string) {
	url := "mongodb://" + username + ":" + password + "@" + db.server
	db.client, err := mongo.NewClient(options.Client().ApplyURI(url))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = db.client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
	    log.Fatal(err)
	}

}

func (db *DB) Disconnect() {
	err = db.client.Disconnect(context.TODO())

	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")	
}

func (db *DB) InsertBlock(block *BLock) {
	collection := client.Database(db.Database).Collection(COLLECTION)
	insert, err := collection.InsertOne(context.TODO(), block)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println("Inserted one Block _id:", insert.InsertedID)
}

func (db *DB) GetBlockchain() {
	collection := client.Database(db.Database).Collection(COLLECTION)

}