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

	img, tags, err := selectImageTag(imgs)
	if err != nil {
		return nil, err
	}

	if confirm, err := confirmDeletion(img, tags); err != nil {
		return nil, fmt.Errorf("error confirming deletion: %v", err)
	} else if !confirm {
		return &DeletionResult{Canceled: true}, nil
	}

	prune := false
	if lt := len(img.Tags); lt == 1 || lt == len(tags) {
		if prune, err = askPrune(img, tags); err != nil {
			return nil, fmt.Errorf("error asking for prune choice: %v", err)
		}
	}

	ids := getIDs(img, tags)
	for _, id := range ids {
		if err := images.Delete(ctx, id, false, prune); err == nil {
			continue
		}

		force, err := askForce(err)
		if err != nil {
			return nil, fmt.Errorf("error asking for force choice: %v", err)
		}

		for _, id := range ids {
			if err := images.Delete(ctx, id, force, prune); err != nil {
				return nil, fmt.Errorf("error on force deletion: %v", err)
			}
		}
	}

	return &DeletionResult{
		Canceled:     false,
		ImageRemoved: img,
		TagsRemoved:  tags,
	}, nil
}

func selectImageTag(imgs []images.Image) (*images.Image, []images.Tag, error) {
	for {
		img, err := selectImage(imgs)
		if err != nil {
			return nil, nil, fmt.Errorf("error selecting the image: %v", err)
		}

		var tags []images.Tag
		if len(img.Tags) > 0 {
			seltags, err := selectTags(img)
			if seltags.isBack {
				continue
			}
			tags = seltags.tags

			if err != nil {
				return nil, nil, fmt.Errorf("error selecting the tag: %v", err)
			}
		}

		return img, tags, err
	}
}

//DeletionResult ...
type DeletionResult struct {
	Canceled     bool
	ImageRemoved *images.Image
	TagsRemoved  []images.Tag
}
