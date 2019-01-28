package shared

import (
	"log"

	"github.com/abramovic/goinsta"
)

func FatalOnErr(err error) {
	if err == nil {
		// do nothing
		return
	}
	switch err {
	case goinsta.ErrChallengeOptionsRequired:
		log.Fatal("we did not pass any options to goinsta.New() so we could not send a challenge code ")
	case goinsta.ErrChallengeCodeRequired:
		log.Fatal("a challenge is required. please check your phone/email")
	case goinsta.ErrChallengeCodeInvalid:
		log.Fatal("The challenge code provided above is invalid")
	default:
		log.Fatal(err)
	}
}
