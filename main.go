package main

import "github.com/mad-czarls/go-api-user/router"

func main() {
	//db := database.Client()
	//
	////@TODO example usage below
	//ctx := context.Background()
	//
	//err := db.Set(ctx, "test_key", "test_value2222", 0).Err()
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//val, err := db.Get(ctx, "test_key").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("test_key value is: ", val)

	r := router.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
