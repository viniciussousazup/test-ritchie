// This is the formula implementation class.
// Where you will code your methods and manipulate the inputs to perform the specific operation you wish to automate.

package formula

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Formula struct {
	StreamTime string
}

func (f Formula) Run() {
	seconds, err := strconv.ParseInt(f.StreamTime, 10, 32)
	if err != nil {
		log.Fatalf("fail to parseInt: %s\nErro:%#v", f.StreamTime, err)
	}
	for i := int64(1); i <= seconds; i++ {
		time.Sleep(time.Second)
		fmt.Fprintf(os.Stdout, "Stdout:%d\n", i)
		fmt.Fprintf(os.Stderr, "Stderr:%d\n", i)
	}
}
