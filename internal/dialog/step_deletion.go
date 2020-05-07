package dialog

import (
	"fmt"

	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/images"
	"github.com/manifoldco/promptui"
)

//ConfirmDeletion ...
func ConfirmDeletion(img *images.Image, tags []images.Tag) (bool, error) {
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
	if err != nil {
		if err == promptui.ErrAbort {
			return false, nil
		}
		return false, err
	}

	return result == "y", nil
}
