package main

import (
	"fmt"
	"github.com/huandu/xstrings"
	"github.com/passionde/rebraincourse/GO-03/random"
)

func main() {
	fmt.Println(random.City())
	fmt.Println(random.Digit())
	fmt.Println(xstrings.Shuffle(random.City()))
}
