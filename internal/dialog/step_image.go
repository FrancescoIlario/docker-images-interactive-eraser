package dialog

import (
	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/images"
	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/prompts"
	"github.com/manifoldco/promptui"
)

type imgSelection struct {
	img    *images.Image
	isQuit bool
}

func selectImage(imgs []images.Image) (*imgSelection, error) {
	promptImg := prompts.ImageSelector(imgs, txHeight)
	i, _, err := promptImg.Run()
	if err != nil {
		if err == promptui.ErrEOF {
			return &imgSelection{isQuit: true}, nil
		}
		return nil, err
	}

	if 0 <= i && i < len(imgs) {
		return &imgSelection{img: &imgs[i]}, nil
	}
	return &imgSelection{isQuit: true}, nil
}
