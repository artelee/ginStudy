package api

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
)

func GetWelcomeMessage(id string) string {
	print(`GetWelcomeMessage ============`)
	if id == `해리포터` {
		return `살아남은 아이의 접속을 환영합니다.`
	} else {
		return `머글은 돌아가세요.`
	}
}

// User firestore user collection fields
type User struct {
	name string `mapstructure:"name"`
	age  int    `mapstructure:"age"`
}

func GetUserList(client *firestore.Client) {

	ctx := context.Background()
	users := client.Collection("user")
	tmp_ := users.Doc("MZL2LmBrFzEW1Q3QCZcG")

	// Or, in a single call:
	//tmp_ = client.Doc("/user/MZL2LmBrFzEW1Q3QCZcG")

	docsnap, err := tmp_.Get(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	dataMap := docsnap.Data()
	fmt.Println(dataMap)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
	//myData := make(map[string]interface{})
	//myData["name"] = "헤르미온느 그레인저"
	//myData["age"] = 15
	//result := &User{}
	//if err := mapstructure.Decode(myData, &result); err != nil {
	//	fmt.Println(err)
	//}
	//_, _, adderr := users.Add(ctx, myData)
	//if adderr != nil {
	//	// Handle any errors in an appropriate way, such as returning them.
	//	log.Printf("An error has occurred: %s", adderr)
	//}

	fmt.Println("All users:")
	iter := users.Documents(ctx)
	for {
		doc, errIter := iter.Next()
		if errIter == iterator.Done {
			break
		}
		if errIter != nil {
			fmt.Println(`errIter : `, errIter)
		}
		fmt.Println(doc.Data())
	}

}
