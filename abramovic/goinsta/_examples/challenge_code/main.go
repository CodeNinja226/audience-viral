package main

import (
	"fmt"

	"github.com/abramovic/goinsta"
	"github.com/abramovic/goinsta/_examples/shared"
)

func main() {
	insta := goinsta.New("USERNAME", "PASSWORD", &goinsta.Challenge{
		Delivery: goinsta.GOINSTA_CHALLENGE_SMS,
		Code:     "123456", // this could be an empty string (if Instagram has not provided a code)
	})

	err := insta.Login()
	shared.FatalOnErr(err)
	defer insta.Logout()

	instaPage, err := insta.GetUserByUsername("thedodo")
	shared.FatalOnErr(err)

	fmt.Printf("%+v\n", instaPage)
}
