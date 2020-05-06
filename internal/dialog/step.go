package dialog

import (
	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/images"
)

func getIDs(img *images.Image, tags []images.Tag) []string {
	lt := len(tags)
	if lt == 0 {
		return []string{img.ID}
	}

	ids := make([]string, lt)
	for idx, tag := range tags {
		ids[idx] = string(tag)
	}
	return ids
}
