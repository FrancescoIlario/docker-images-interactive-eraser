package prompts

import (
	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/images"
	"github.com/manifoldco/promptui"
)

// ImageSelector creates an image selector prompt
func ImageSelector(imgs []images.Image, txHeight int) *promptui.Select {
	data := append(imgs, images.Image{
		ID:   "Quit",
		Size: "<-- Exit from the application",
	})

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "\U00002326   {{ .ID | cyan }} {{ .Size | white }} {{ .Tags | white }}",
		Inactive: "    {{ .ID | cyan }} {{ .Size | white }} {{ .Tags | white }}",
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
		Items:     data,
		Templates: templates,
		Size:      txHeight,
	}
}

type imageTags struct {
	tag  string
	size string
	id   string
}
