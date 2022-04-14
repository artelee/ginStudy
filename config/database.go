package config

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"log"
	"os"
	"sync"
)

type FirestoreDB struct {
	client *firestore.Client
	url    string
}

var dbInstance *FirestoreDB
var once sync.Once

func DBClient() *firestore.Client {
	once.Do(func() {
		fmt.Println("DB 연결 시작 --- FIREBASE ---")
		dbInstance = new(FirestoreDB)
		makeDBClient(dbInstance)
		fmt.Println("DB 연결 완료 --- FIREBASE ---")
	})
	return dbInstance.client
}

// makeDBClient Client 주입
func makeDBClient(rtb *FirestoreDB) {

	fmt.Println("dbstore name : " + os.Getenv("FIREBASE_FIRESTORE_DATABASE_PRODUCT_ID"))
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, os.Getenv("FIREBASE_FIRESTORE_DATABASE_PRODUCT_ID"))
	if err != nil {
		log.Fatal(err)
	}

	rtb.client = client
}
