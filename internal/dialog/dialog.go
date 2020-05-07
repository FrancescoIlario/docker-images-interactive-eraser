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
func DeleteImage(ctx context.Context, imgs []images.Image, useTags bool) (*DeletionResult, error) {
	txHeight = tx.CalculateHeight(txDiff)

	sel, err := selectImageTag(imgs, useTags)
	if err != nil {
		return nil, err
	}
	if sel.isQuit {
		return &DeletionResult{
			Canceled: true,
		}, nil
	}

	img, tags := sel.img, sel.tags
	prune := false
	if !useTags || len(img.Tags) == 1 || len(img.Tags) == len(tags) {
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

		if err := images.Delete(ctx, id, force, prune); err != nil {
			return nil, fmt.Errorf("error on force deletion: %v", err)
		}
	}

	return &DeletionResult{
		Canceled:     false,
		ImageRemoved: img,
		TagsRemoved:  tags,
	}, nil
}

func selectImageTag(imgs []images.Image, useTags bool) (*imageTagsSelection, error) {
	cursor, scroll := 0, 0
	for {
		if li := len(imgs); cursor >= li || scroll >= li {
			li--
			cursor, scroll = li, li
		}

		sel, err := SelectImage(imgs, cursor, scroll)
		if err != nil {
			return nil, fmt.Errorf("error selecting the image: %v", err)
		}
		if sel.isQuit {
			return &imageTagsSelection{isQuit: true}, nil
		}
		cursor, scroll = sel.cursor, sel.scroll

		var tags []images.Tag
		if !useTags {
			tags = sel.img.Tags
		} else if len(sel.img.Tags) > 0 {
			seltags, err := selectTags(sel.img)
			if err != nil {
				return nil, fmt.Errorf("error selecting the tag: %v", err)
			}

			if seltags.isBack {
				continue
			}
			tags = seltags.tags
		}

		confirm, err := ConfirmDeletion(sel.img, tags)
		if err != nil {
			return nil, fmt.Errorf("error confirming deletion: %v", err)
		} else if !confirm {
			continue
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
