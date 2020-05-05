package dialog

import (
	"fmt"

	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/images"
	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/prompts"
	"github.com/manifoldco/promptui"
)

func getID(img *images.Image, tag *images.Tag) string {
	if tag != nil {
		return string(*tag)
	}
	return string(img.ID)
}

func selectImage(imgs []images.Image) (img *images.Image, err error) {
	promptImg := prompts.ImageSelector(imgs, txHeight)
	if i, _, err := promptImg.Run(); err == nil {
		img = &imgs[i]
	}
	return
}

type tagSelection struct {
	tag    *images.Tag
	isBack bool
}

func selectTag(img *images.Image) (*tagSelection, error) {
	if lenTags := len(img.Tags); lenTags == 0 {
		return nil, fmt.Errorf("image seems to have no tags")
	} else if lenTags == 1 {
		return &tagSelection{tag: &img.Tags[0]}, nil
	}

	promptTag := prompts.TagSelector(img, txHeight)
	i, _, err := promptTag.Run()
	if err != nil {
		return nil, err
	}

	if i < len(img.Tags) {
		return &tagSelection{tag: &img.Tags[i]}, nil
	}
	return &tagSelection{isBack: true}, nil
}

func confirmDeletion(img *images.Image, tag *images.Tag) (confirm bool, err error) {
	id := img.ID
	if tag != nil {
		id = string(*tag)
	}

	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("Do you really want to delete tag %v", id),
		IsConfirm: true,
	}
	result, err := prompt.Run()
	if err == nil {
		confirm = result == "y"
	}
	return
}

func askPrune(img *images.Image, tag *images.Tag) (prune bool, err error) {
	id := img.ID
	if tag != nil {
		id = string(*tag)
	}

	pruneConfirm := promptui.Prompt{
		Label:     fmt.Sprintf("Do you want to prune the children of the image with tag %s", id),
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
