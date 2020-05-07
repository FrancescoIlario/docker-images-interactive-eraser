package dialog

import (
	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/images"
	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/prompts"
	"github.com/manifoldco/promptui"
)

//ImgSelection ...
type ImgSelection struct {
	img    *images.Image
	isQuit bool
	cursor int
	scroll int
}

//SelectImage ...
func SelectImage(imgs []images.Image, cursor, scroll int) (*ImgSelection, error) {
	promptImg := prompts.ImageSelector(imgs, txHeight)
	i, _, err := promptImg.RunCursorAt(cursor, scroll)
	if err != nil {
		if err == promptui.ErrEOF {
			return &ImgSelection{
				isQuit: true,
				scroll: i,
				cursor: i,
			}, nil
		}
		return nil, err
	}

	if 0 <= i && i < len(imgs) {
		return &ImgSelection{
			img:    &imgs[i],
			scroll: i,
			cursor: i,
		}, nil
	}
	return &ImgSelection{isQuit: true}, nil
}
