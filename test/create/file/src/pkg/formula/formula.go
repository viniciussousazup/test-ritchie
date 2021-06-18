// This is the formula implementation class.
// Where you will code your methods and manipulate the inputs to perform the specific operation you wish to automate.

package formula

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"time"
)

type Formula struct {
	Name      string
	OutputDir string
}

func (f Formula) Run(writer io.Writer) {
	fmt.Println("simulation a creation of s3.")

	fmt.Printf("Creating s3 with name:%s\n", f.Name)

	if f.OutputDir != "" {
		outPutFilePath := fmt.Sprintf("%s/output.json", f.OutputDir)
		fmt.Printf("Writhing output to file:%s\n", outPutFilePath)
		outputData := map[string]string{
			"name": f.Name,
			"now-unix": strconv.Itoa(int(time.Now().Unix())),
		}
		outputJson , _ := json.Marshal(outputData)
		ioutil.WriteFile(outPutFilePath, outputJson,0777)
	}

}
