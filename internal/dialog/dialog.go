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

//DeletionResult ...
type DeletionResult struct {
	Canceled     bool
	ImageRemoved *images.Image
	TagsRemoved  []images.Tag
}

// DeleteImage ...
func DeleteImage(ctx context.Context, imgs []images.Image) (*DeletionResult, error) {
	txHeight = tx.CalculateHeight(txDiff)

	sel, err := selectImageTag(imgs)
	if err != nil {
		return nil, err
	}
	if sel.isQuit {
		return &DeletionResult{
			Canceled: true,
		}, nil
	}

	img, tags := sel.img, sel.tags
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

func selectImageTag(imgs []images.Image) (*imageTagsSelection, error) {
	for {
		sel, err := selectImage(imgs)
		if err != nil {
			return nil, fmt.Errorf("error selecting the image: %v", err)
		}
		if sel.isQuit {
			return &imageTagsSelection{isQuit: true}, nil
		}

		var tags []images.Tag
		if len(sel.img.Tags) > 0 {
			seltags, err := selectTags(sel.img)
			if seltags.isBack {
				continue
			}
			tags = seltags.tags

			if err != nil {
				return nil, fmt.Errorf("error selecting the tag: %v", err)
			}
		}

		return &imageTagsSelection{
			img:  sel.img,
			tags: tags,
		}, err
	}
}

type imageTagsSelection struct {
	img    *images.Image
	tags   []images.Tag
	isQuit bool
}
