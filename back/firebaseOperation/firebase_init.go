package firebaseOperation

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
)

func Init()*firebase.App {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	fmt.Println(app)
	fmt.Println("jgirhgi")
	return app
}
