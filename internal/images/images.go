package images

import (
	"context"
	"math/big"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/dustin/go-humanize"
)

//Image ...
type Image struct {
	ID   string
	Size string
	Tags []Tag
}

//Tag ...
type Tag string

//GetImages ...
func GetImages(ctx context.Context) ([]Image, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		return nil, err
	}

	result := make([]Image, len(images))
	for idx, image := range images {
		id := image.ID[7:19]

		bb := big.NewInt(image.Size)
		size := humanize.BigBytes(bb)

		tags := make([]Tag, len(image.RepoTags))
		for i, v := range image.RepoTags {
			tags[i] = Tag(v)
		}

		result[idx] = Image{
			ID:   id,
			Size: size,
			Tags: tags,
		}
	}
	return result, nil
}

//Delete identifier may be an image id or a tag
func Delete(ctx context.Context, identifier string, pruneChildren, force bool) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	if _, err = cli.ImageRemove(ctx, identifier, types.ImageRemoveOptions{
		PruneChildren: pruneChildren,
		Force:         force,
	}); err != nil {
		return err
	}

	return nil
}
