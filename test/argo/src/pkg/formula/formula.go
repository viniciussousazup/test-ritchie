// This is the formula implementation class.
// Where you will code your methods and manipulate the inputs to perform the specific operation you wish to automate.

package formula

import (
	"fmt"
	"io"

	"github.com/gookit/color"
)

type Formula struct {
	Text     string
	List     string
	Boolean  bool
	Password string
}

func (f Formula) Run(writer io.Writer) {
	var result string

	result += fmt.Sprintf("Hello World!\n")

	result += color.FgGreen.Render(fmt.Sprintf("My name is %s.\n", f.Text))

	if _, err := fmt.Fprintf(writer, result); err != nil {
		panic(err)
	}
}
