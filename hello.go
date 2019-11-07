package main

import (
	"fmt"
	"strings"
)

func main() {
	v := []string{"wocao","nima","aiyo"}

	fmt.Println(strings.Join(v, "-"))
}
