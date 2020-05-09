package prompts

import (
	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/images"
	"github.com/manifoldco/promptui"
)

// ImageSelector creates an image selector prompt
func ImageSelector(imgs []images.Image, txHeight int) *promptui.Select {
	templates := &promptui.SelectTemplates{
		Label: "{{ . }}",
		Active: "\U00002326   {{ .ID | cyan }} {{ .Size }}	{{ .Tags  }}",
		Inactive: "    {{ .ID | cyan }} {{ .Size }}	{{ .Tags }}",
		Selected: "\U00002326 {{ .ID | white | cyan }}",
		Details: `
--------- Image ----------
{{ "ID:" | faint }}	{{ .ID }}
{{ "Size:" | faint }}	{{ .Size }}
{{ "Tags:" | faint }}	{{ .Tags }}
`,
	}

	return &promptui.Select{
		Label:     "Choose one image",
		Items:     imgs,
		Templates: templates,
		Size:      txHeight,
	}
}

type imageTags struct {
	tag  string
	size string
	id   string
}
