package config

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"log"
	"os"
	"sync"
)

type RealtimeDB struct {
	client *firestore.Client
	//client *db.Client
	url string
}

var dbInstance *RealtimeDB
var once sync.Once

func DBClient() *firestore.Client {
	once.Do(func() {
		fmt.Println("DB 연결 시작 --- FIREBASE ---")

		dbInstance = new(RealtimeDB)
		GetDBClient(dbInstance)
		fmt.Println("DB 연결 완료 --- FIREBASE ---")
	})
	return dbInstance.client
}

func GetDBClient(rtb *RealtimeDB) {

	fmt.Println("dbstore name : " + os.Getenv("FIREBASE_FIRESTORE_DATABASE_PRODUCT_ID"))
	ctx := context.Background()

	//config := &firebase.Config{
	//	ProjectID:   os.Getenv("FIREBASE_FIRESTORE_DATABASE_PRODUCT_ID"),
	//	DatabaseURL: rtb.url,
	//}

	//app, err := firebase.NewApp(ctx, config, option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")))
	//if err != nil {
	//	log.Fatal(err)
	//}
	client, err := firestore.NewClient(ctx, os.Getenv("FIREBASE_FIRESTORE_DATABASE_PRODUCT_ID"))
	if err != nil {
		log.Fatal(err)
	}
	//client, err := app.Database(ctx)
	//if err != nil {
	//	log.Fatal(err)
	//}
	rtb.client = client
}
