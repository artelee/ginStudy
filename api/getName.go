package api

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
)

func GetWelcomeMessage(id string) string {
	print(`GetWelcomeMessage ============`)
	if id == `해리포터` {
		return `살아남은 아이의 접속을 환영합니다.`
	} else {
		return `머글은 돌아가세요.`
	}
}

type User struct {
	name string
	age  int
}

func GetUserList(client *firestore.Client) {
	ctx := context.Background()
	//var user User
	//var result map[string]User
	//if err := client.NewRef("/user/MZL2LmBrFzEW1Q3QCZcG").Get(ctx, &user); err != nil {
	//	log.Fatal(err)
	//}
	//
	//q := client.NewRef("user")
	//if err := q.Get(ctx, &result); err != nil {
	//	log.Fatal(err)
	//}
	//for key, acc := range result {
	//	log.Printf("%s => %v\n", key, acc)
	//}
	//log.Printf("%s has a balance of %f\n", user.name, user.age)

	///user/MZL2LmBrFzEW1Q3QCZcG
	states := client.Collection("user")
	tmp_ := states.Doc("MZL2LmBrFzEW1Q3QCZcG")
	// Or, in a single call:
	//tmp_ = client.Doc("/user/MZL2LmBrFzEW1Q3QCZcG")
	docsnap, err := tmp_.Get(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	dataMap := docsnap.Data()
	fmt.Println(dataMap)
}
