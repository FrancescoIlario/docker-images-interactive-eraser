package greeter

import (
	"fmt"

	"github.com/FrancescoIlario/docker-images-interactive-eraser/pkg/envext"
)

//Greeter ...
type Greeter interface {
	//Greet ...
	Greet()
}

type greeter struct{}

//New ...
func New() Greeter {
	return &greeter{}
}

func (g *greeter) Greet() {
	greet := envext.GetEnvOrDefault("GREETER_GREET", "Hello World!")
	fmt.Println(greet)
}
