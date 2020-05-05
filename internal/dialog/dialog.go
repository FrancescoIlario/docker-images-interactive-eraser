package dialog

import (
	"context"
	"fmt"

	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/images"
	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/tx"
)

const (
	txDiff = uint(9)
)

var (
	txHeight = 5
)

// DeleteImage ...
func DeleteImage(ctx context.Context, imgs []images.Image) (*DeletionResult, error) {
	txHeight = tx.CalculateHeight(txDiff)

	img, tag, err := selectImageTag(imgs)
	if err != nil {
		return nil, err
	}

	if confirm, err := confirmDeletion(img, tag); err != nil {
		return nil, fmt.Errorf("error confirming deletion: %v", err)
	} else if !confirm {
		return &DeletionResult{Canceled: true}, nil
	}

	prune := false
	if len(img.Tags) == 1 {
		if prune, err = askPrune(img, tag); err != nil {
			return nil, fmt.Errorf("error asking for prune choice: %v", err)
		}
	}

	id := getID(img, tag)
	if err := images.Delete(ctx, id, false, prune); err == nil {
		return nil, err
	}

	force, err := askForce(err)
	if err != nil {
		return nil, fmt.Errorf("error asking for force choice: %v", err)
	}

	if err := images.Delete(ctx, id, force, prune); err != nil {
		return nil, fmt.Errorf("error on force deletion: %v", err)
	}

	return &DeletionResult{
		Canceled:     false,
		ImageRemoved: img,
		TagRemoved:   tag,
	}, nil
}

func selectImageTag(imgs []images.Image) (*images.Image, *images.Tag, error) {
	for {
		img, err := selectImage(imgs)
		if err != nil {
			return nil, nil, fmt.Errorf("error selecting the image: %v", err)
		}

		var tag *images.Tag
		if len(img.Tags) > 0 {
			seltag, err := selectTag(img)
			if seltag.isBack {
				continue
			}
			tag = seltag.tag

			if err != nil {
				return nil, nil, fmt.Errorf("error selecting the tag: %v", err)
			}
		}

		return img, tag, err
	}
}

//DeletionResult ...
type DeletionResult struct {
	Canceled     bool
	ImageRemoved *images.Image
	TagRemoved   *images.Tag
}
