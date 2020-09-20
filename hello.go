package main

import (
	"fmt"
	
	"github.com/nowiseeyou/stringutil"
)

func main()  {
	s := "Hello , source,xi xi Msd"
	reverse := stringutil.Reverse(s)
	fmt.Printf(reverse)
}
