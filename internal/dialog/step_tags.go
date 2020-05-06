package dialog

import (
	"fmt"

	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/images"
	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/prompts"
	"github.com/manifoldco/promptui"
)

type tagSelection struct {
	tags   []images.Tag
	isBack bool
}

func selectTags(img *images.Image) (*tagSelection, error) {
	if lenTags := len(img.Tags); lenTags == 0 {
		return nil, fmt.Errorf("image seems to have no tags")
	} else if lenTags == 1 {
		return &tagSelection{tags: []images.Tag{img.Tags[0]}}, nil
	}

	promptTag := prompts.TagSelector(img, txHeight)
	i, _, err := promptTag.Run()
	if err != nil {
		if err == promptui.ErrEOF {
			return &tagSelection{isBack: true}, nil
		}
		return nil, err
	}

	if i == 0 {
		return &tagSelection{tags: img.Tags}, nil
	} else if i <= len(img.Tags) {
		return &tagSelection{tags: []images.Tag{img.Tags[i]}}, nil
	}
	return &tagSelection{isBack: true}, nil
}
