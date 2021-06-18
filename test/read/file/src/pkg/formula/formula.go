// This is the formula implementation class.
// Where you will code your methods and manipulate the inputs to perform the specific operation you wish to automate.

package formula

import (
	"fmt"
	"io"
	"io/ioutil"
)

type Formula struct {
	OutputDir string
}

func (f Formula) Run(writer io.Writer) {
	fmt.Println("read-file:" + f.OutputDir)

	if f.OutputDir != "" {
		outPutFilePath := fmt.Sprintf("%s/output.json", f.OutputDir)
		fmt.Printf("Read output from file:%s\n", outPutFilePath)
		data, err := ioutil.ReadFile(outPutFilePath)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))
	}

}
