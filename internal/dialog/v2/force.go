package dialog

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func askForce(deletionErr error) (bool, error) {
	errConf := promptui.Prompt{
		Label:     fmt.Sprintf("Can not delete (%v), do you want to force", deletionErr),
		IsConfirm: true,
	}

	resForce, err := errConf.Run()
	if err != nil {
		if err == promptui.ErrEOF {
			return false, ErrCanceled
		}
		if err == promptui.ErrAbort {
			return false, nil
		}
		return false, err
	}
	return resForce == "y" || resForce == "Y", nil
}
