// This is the formula implementation class.
// Where you will code your methods and manipulate the inputs to perform the specific operation you wish to automate.

package formula

import (
	"context"
	"fmt"

	"github.com/google/go-github/v33/github"
	"golang.org/x/oauth2"
)

type Formula struct {
	Token string
}

func (f Formula) Run() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: f.Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		fmt.Println(err)
	}
	for _, repository := range repos {
		fmt.Println(*repository.Name)
	}
}
