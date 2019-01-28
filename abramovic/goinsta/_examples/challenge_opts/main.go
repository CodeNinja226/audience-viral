package main

import (
	"fmt"

	"github.com/abramovic/goinsta"
	"github.com/abramovic/goinsta/_examples/shared"
)

func main() {
	// we will not pass any options
	insta := goinsta.New("USERNAME", "PASSWORD", nil)

	err := insta.Login()
	shared.FatalOnErr(err)

	defer insta.Logout()

	instaPage, err := insta.GetUserByUsername("thedodo")
	shared.FatalOnErr(err)

	fmt.Printf("%+v\n", instaPage)
}
