// This is the formula implementation class.
// Where you will code your methods and manipulate the inputs to perform the specific operation you wish to automate.

package formula

import (
	"fmt"
	"time"
)

type Formula struct {
	Token string
}

func (f Formula) Run() {
	// ctx := context.Background()
	// ts := oauth2.StaticTokenSource(
	// 	&oauth2.Token{AccessToken: f.Token},
	// )
	// tc := oauth2.NewClient(ctx, ts)

	// client := github.NewClient(tc)

	// // list all repositories for the authenticated user
	// repos, _, err := client.Repositories.List(ctx, "", nil)
	// if err != nil {
	// 	fmt.Println("list repos", err)
	// }
	test := 1
	for ok := true; ok; ok = test < 7 {
		fmt.Printf("STEP %d \n", test)
		time.Sleep(5 * time.Second)
		test++
	}
}
