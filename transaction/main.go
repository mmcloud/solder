package main

import (
	"context"
	"log"
)

// var interfaces.Service _ := (*RecommendationService)(nil)

// type RecommendationService struct {
// }

func main() {
	ctx := context.Background()
	sold, err := Solder(ctx)
	if err != nil {
		log.Fatal(err)
	}
	sold.Start(ctx)

}
