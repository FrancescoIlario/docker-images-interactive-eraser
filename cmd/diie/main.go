package main

import (
	"context"
	"log"

	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/dialog"
	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/images"
)

const (
	txDiff = uint(9)
)

var (
	txHeight = 5
)

func main() {
	ctx := context.Background()

	for {
		imgs, err := images.GetImages(ctx)
		if err != nil {
			log.Printf("can not retrieve the list of Docker images: %v", err)
		}

		dl, err := dialog.DeleteImage(ctx, imgs)
		if err != nil {
			log.Fatalln(err)
		}
		if dl.Canceled {
			return
		}
	}
}
