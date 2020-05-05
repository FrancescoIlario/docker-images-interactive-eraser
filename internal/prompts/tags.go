package prompts

import (
	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/images"
	"github.com/manifoldco/promptui"
)

type opTag struct {
	images.Tag
	Description string
}

func createOpTags(img *images.Image) []opTag {
	opTags := make([]opTag, len(img.Tags)+2)
	opTags[0] = opTag{
		Tag:         "+ All",
		Description: "Select all tags",
	}
	for i, tag := range img.Tags {
		opTags[i+1] = opTag{
			Tag:         tag,
			Description: string(tag),
		}
	}
	opTags[len(img.Tags)+1] = opTag{
		Tag:         "<-- Back",
		Description: "Back to the previous menu",
	}
	return opTags
}

//TagSelector creates a tag selector prompt
func TagSelector(img *images.Image, txHeight int) *promptui.Select {
	opTags := createOpTags(img)
	templates := &promptui.SelectTemplates{
		Label:    "Image: {{ .ID }} {{ .Size }}",
		Active:   "\U00002326   {{ .Tag | cyan }}",
		Inactive: "    {{ .Tag | cyan }}",
		Selected: "\U00002326 {{ .Tag | red | cyan }}",
		Details: `
--------- Image ----------
{{ "Description:" | faint }}	{{ .Description }}
`,
	}

	prompt := promptui.Select{
		Label:     img,
		Items:     opTags,
		Templates: templates,
		Size:      txHeight,
	}

	return &prompt
}
