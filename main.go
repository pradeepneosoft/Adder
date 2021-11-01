package main

import (
	"fmt"

	"github.com/pradeepneosoft/Add"
)

func main() {
	a := Add.Init{
		12.34,
		34.22,
	}
	fmt.Println(a.Add())

}
