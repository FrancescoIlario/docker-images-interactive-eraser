package dialog

import (
	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/images"
	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/prompts"
)

type imgSelection struct {
	img    *images.Image
	isQuit bool
}

func selectImage(imgs []images.Image) (sel *imgSelection, err error) {
	promptImg := prompts.ImageSelector(imgs, txHeight)
	if i, _, err := promptImg.Run(); err == nil {
		if i == len(imgs) {
			sel = &imgSelection{isQuit: true}
		} else {
			sel = &imgSelection{img: &imgs[i]}
		}
	}
	return
}
