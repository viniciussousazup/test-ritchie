// This is the formula implementation class.
// Where you will code your methods and manipulate the inputs to perform the specific operation you wish to automate.

package formula

import (
	"fmt"
	"os"
)

type Formula struct {
	Token string
}

func (f Formula) Run() {
	fmt.Fprintf(os.Stdout, "Stdout:%s\n", f.Token)
}
