package dialog

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func askForce(deletionErr error) (force bool, err error) {
	errConf := promptui.Prompt{
		Label:     fmt.Sprintf("Can not delete (%v), do you want to force", deletionErr),
		IsConfirm: true,
	}
	if resForce, err := errConf.Run(); err == nil {
		force = resForce == "y"
	}
	return
}
