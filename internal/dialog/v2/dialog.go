package dialog

import (
	"context"
	"fmt"

	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/dialog"
	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/images"
)

type deleteImageDialog struct {
	ctx           context.Context
	useTags       bool
	imgs          []images.Image
	txHeight      int
	selectionForm SelectionForm
}

//DeleteImageDialog ...
type DeleteImageDialog interface {
	DeleteImage() error
	GetImageDisplayStatus() DisplayStatus
}

//NewDeleteImageDialog ...
func NewDeleteImageDialog(ctx context.Context, imgs []images.Image, useTags bool, txHeight int, display *DisplayStatus) DeleteImageDialog {
	selForm := NewSelectionForm(imgs, useTags, txHeight, display)

	return &deleteImageDialog{
		ctx:           ctx,
		useTags:       useTags,
		imgs:          imgs,
		txHeight:      txHeight,
		selectionForm: selForm,
	}
}

//GetDisplayStatus ...
func (d *deleteImageDialog) GetImageDisplayStatus() DisplayStatus {
	return d.selectionForm.GetDisplayStatus()
}

//ErrNotConfirmed error returned when the deletion was not confirmed
var ErrNotConfirmed = fmt.Errorf("deletion not confirmed")

//DeleteImage ...
func (d *deleteImageDialog) DeleteImage() error {
	sel, err := d.selectionForm.SelectImageTags()
	if err != nil {
		return err
	}

	confirm, err := d.confirmDeletion(sel)
	if err != nil {
		return err
	} else if !confirm {
		return ErrNotConfirmed
	}

	prune := true
	if it := len(sel.Image.Tags); it <= 1 || it == len(sel.Tags) {
		if prune, err = d.askPrune(sel); err != nil {
			return err
		}
	}

	ids, ctx := getIDs(sel), context.Background()
	for _, id := range ids {
		err := images.Delete(ctx, id, false, prune)
		if err == nil {
			continue
		}

		force, err := askForce(err)
		if err != nil {
			return err
		}
		if force {
			if err := images.Delete(ctx, id, force, prune); err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *deleteImageDialog) confirmDeletion(sel *SelectionResult) (bool, error) {
	return dialog.ConfirmDeletion(sel.Image, sel.Tags)
}

func getIDs(sel *SelectionResult) []string {
	lt := len(sel.Tags)
	if lt == 0 {
		return []string{sel.Image.ID}
	}

	ids := make([]string, lt)
	for idx, tag := range sel.Tags {
		ids[idx] = string(tag)
	}
	return ids
}
