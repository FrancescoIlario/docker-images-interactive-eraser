package dialog

import "github.com/manifoldco/promptui"

func (d *deleteImageDialog) askPrune(sel *SelectionResult) (bool, error) {
	pruneConfirm := promptui.Prompt{
		Label:     "Do you want to prune the children of the image",
		IsConfirm: true,
	}
	pruneResult, err := pruneConfirm.Run()
	if err != nil {
		if err == promptui.ErrEOF {
			return false, ErrCanceled
		}
		if err == promptui.ErrAbort {
			return false, nil
		}
		return false, err
	}
	res := pruneResult == "y" || pruneResult == "Y"
	return res, nil
}
