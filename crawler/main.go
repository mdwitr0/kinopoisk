package main

import (
	"context"
	"github.com/mdwitr0/kinopoisk/crawler/external/kp"
)

func main() {

	ctx := context.Background()

	kpClient := kp.NewClient()

	res, err := kpClient.GetAllMovies(ctx, 50, 0)
	if err != nil {
		panic(err)
	}

	print(res)

}
