package main

import (
	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/greeter"
)

func main() {
	gr := greeter.New()
	gr.Greet()
}
