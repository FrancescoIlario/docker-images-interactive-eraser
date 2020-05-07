package dialog

import (
	"fmt"

	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/images"
	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/prompts"
	"github.com/manifoldco/promptui"
)

//ErrImgSelCanceled error returned if image selection has been canceled by the user
var ErrImgSelCanceled = fmt.Errorf("image selection canceled")

//ErrTagSelCanceled error returned if tag selection has been canceled by the user
var ErrTagSelCanceled = fmt.Errorf("tag selection canceled")

type selectionForm struct {
	useTags  bool
	txHeight int
	imgs     []images.Image
	display  DisplayStatus
}

//DisplayStatus current scroll and cursor
type DisplayStatus struct {
	Scroll int
	Cursor int
}

//SelectionResult ...
type SelectionResult struct {
	Image *images.Image
	Tags  []images.Tag
}

//SelectionForm ...
type SelectionForm interface {
	SelectImageTags() (*SelectionResult, error)
	GetDisplayStatus() DisplayStatus
}

//NewSelectionForm ...
func NewSelectionForm(imgs []images.Image, useTags bool, txHeight int, display *DisplayStatus) SelectionForm {
	if display == nil {
		display = &DisplayStatus{}
	}

	return &selectionForm{
		imgs:     imgs,
		useTags:  useTags,
		txHeight: txHeight,
		display:  *display,
	}
}

//SelectImageTags ...
func (f *selectionForm) SelectImageTags() (*SelectionResult, error) {
	for {
		sel, err := f.selectImageTags()
		if err != ErrTagSelCanceled {
			return sel, err
		}
	}
}

func (f *selectionForm) GetDisplayStatus() DisplayStatus {
	return f.display
}

func (f *selectionForm) selectImageTags() (*SelectionResult, error) {
	img, err := f.selectImage()
	if err != nil {
		if err == promptui.ErrEOF {
			return nil, ErrImgSelCanceled
		}
		return nil, err
	}

	var tags []images.Tag
	if f.useTags && len(img.Tags) > 0 {
		tags, err = f.selectTags(img)
		if err != nil {
			if err == promptui.ErrEOF {
				return nil, ErrTagSelCanceled
			}
			return nil, err
		}
	}

	return &SelectionResult{Image: img, Tags: tags}, nil
}

func (f *selectionForm) selectImage() (*images.Image, error) {
	promptImg := prompts.ImageSelector(f.imgs, f.txHeight)
	i, _, err := promptImg.RunCursorAt(f.display.Cursor, f.display.Scroll)
	if err != nil {
		return nil, err
	}
	f.display.Cursor, f.display.Scroll = i, i

	return &f.imgs[i], nil
}

func (f *selectionForm) selectTags(img *images.Image) ([]images.Tag, error) {
	promptImg := prompts.TagSelector(img, f.txHeight)
	i, _, err := promptImg.Run()
	if err != nil {
		return nil, err
	}

	if i == 0 {
		return img.Tags, nil
	}
	return []images.Tag{img.Tags[i-1]}, nil
}
