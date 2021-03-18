// This is the formula implementation class.
// Where you will code your methods and manipulate the inputs to perform the specific operation you wish to automate.

package formula

import (
	"context"
	"fmt"

	"github.com/jenkins-x/go-scm/scm"
	"github.com/jenkins-x/go-scm/scm/factory"
)

type Formula struct {
	Token string
}

func (f Formula) Run() {
	client, err := factory.NewClient("github", "https://api.github.com", f.Token)
	ctx := context.Background()
	if err != nil {
		fmt.Println(err)
	}
	repositories, _, err := client.Repositories.List(ctx, scm.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}

	for _, repository := range repositories {
		fmt.Printf(repository.Name)
	}
	// fmt.Fprintf(os.Stdout, "Stdout:%s\n", f.Token)
}
