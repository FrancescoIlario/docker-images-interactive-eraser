package dialog

import (
	"fmt"

	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/images"
	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/prompts"
	"github.com/manifoldco/promptui"
)

func getIDs(img *images.Image, tags []images.Tag) []string {
	lt := len(tags)
	if lt == 0 {
		return []string{img.ID}
	}

	ids := make([]string, lt)
	for idx, tag := range tags {
		ids[idx] = string(tag)
	}
	return ids
}

func selectImage(imgs []images.Image) (img *images.Image, err error) {
	promptImg := prompts.ImageSelector(imgs, txHeight)
	if i, _, err := promptImg.Run(); err == nil {
		img = &imgs[i]
	}
	return
}

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
		return nil, err
	}

	if i == 0 {
		return &tagSelection{tags: img.Tags}, nil
	} else if i < len(img.Tags) {
		return &tagSelection{tags: []images.Tag{img.Tags[i]}}, nil
	}
	return &tagSelection{isBack: true}, nil
}

func confirmDeletion(img *images.Image, tags []images.Tag) (confirm bool, err error) {
	var label string
	if tags == nil || len(tags) == 0 {
		label = fmt.Sprintf("Do you really want to delete image %v", img.ID)
	} else if len(tags) == 1 {
		label = fmt.Sprintf("Do you really want to delete image %v (tag: %v)", img.ID, tags[0])
	} else {
		label = fmt.Sprintf("Do you really want to delete tags %v (image %v)", tags, img.ID)
	}

	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}
	result, err := prompt.Run()
	if err == nil {
		confirm = result == "y"
	}
	return
}

func askPrune(img *images.Image, tags []images.Tag) (prune bool, err error) {
	pruneConfirm := promptui.Prompt{
		Label:     "Do you want to prune the children of the image?",
		IsConfirm: true,
	}
	pruneResult, err := pruneConfirm.Run()
	if err == nil {
		prune = pruneResult == "y"
	}
	return
}

func askForce(deletionErr error) (force bool, err error) {
	errConf := promptui.Prompt{
		Label:     fmt.Sprintf("Can not delete (%v), do you want to force?", deletionErr),
		IsConfirm: true,
	}
	if resForce, err := errConf.Run(); err == nil {
		force = resForce == "y"
	}
	return
}
