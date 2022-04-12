package config

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"fmt"
	"log"
	"os"
	"sync"
)

type RealtimeDB struct {
	client *db.Client
	url    string
}

var dbInstance *RealtimeDB
var once sync.Once

func DBClient() *db.Client {
	once.Do(func() {
		fmt.Println("DB 연결 시작 --- FIREBASE REALTIME ---")
		dbInstance = new(RealtimeDB)
		GetDBConnenction(dbInstance)
		fmt.Println("DB 연결 완료 --- FIREBASE REALTIME ---")
	})
	return dbInstance.client
}

func GetDBConnenction(rtb *RealtimeDB) {
	rtb.url = `https://` + os.Getenv("FIREBASE_REALTIME_DATABASE_PRODUCT_ID") + `.firebaseio.com`

	ctx := context.Background()
	config := &firebase.Config{
		DatabaseURL: rtb.url,
	}
	app, err := firebase.NewApp(ctx, config)
	if err != nil {
		log.Fatal(err)
	}
	client, err := app.Database(ctx)
	if err != nil {
		log.Fatal(err)
	}
	rtb.client = client
}
