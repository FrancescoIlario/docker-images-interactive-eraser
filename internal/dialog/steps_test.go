package dialog

import (
	"testing"

	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/images"
)

func Test_getIDs_NoTags(t *testing.T) {
	img := &images.Image{
		ID:   "img",
		Size: "200 MB",
		Tags: []images.Tag{"aa", "bb", "cc"},
	}

	tags := []images.Tag{"aa", "bb"}

	ids := getIDs(img, tags)
	if len(ids) != 2 {
		t.Fatalf("expected 2 ids, obtained %v: %v", len(ids), ids)
	}
	if ids[0] != "aa" {
		t.Fatalf("expected ids[0] == %v, obtained %v", tags[0], ids[0])
	}
	if ids[1] != "bb" {
		t.Fatalf("expected ids[0] == %v, obtained %v", tags[1], ids[0])
	}
}
func Test_getIDs(t *testing.T) {
	img := &images.Image{
		ID:   "img",
		Size: "200 MB",
	}

	ids := getIDs(img, nil)
	if len(ids) != 1 {
		t.Fatalf("expected 1 id, obtained %v: %v", len(ids), ids)
	}
	if ids[0] != img.ID {
		t.Fatalf("expected ids[0] == %v, obtained %v", img.ID, ids[0])
	}
}
