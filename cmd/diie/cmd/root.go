package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/dialog/v2"
	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/images"
	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/tx"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   `diie`,
		Short: `Docker Image Interactive Eraser`,
		Long:  `Delete Images and Tags easily from CLI`,
		Run:   run,
	}

	useTags bool
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&useTags, "use-tags", "t", false, `Use tags selection`)
}

//Execute executes the root cmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	display := dialog.DisplayStatus{}

	for {
		imgs, err := images.GetImages(ctx)
		if err != nil {
			log.Printf("can not retrieve the list of Docker images: %v", err)
		}

		txHeight := tx.CalculateHeight(9)
		dg := dialog.NewDeleteImageDialog(ctx, imgs, useTags, txHeight, &display)
		if err := dg.DeleteImage(); err != nil {
			log.Fatalln(err)
		}
		display = dg.GetImageDisplayStatus()
	}
}
