package main

import (
	"fmt"
	"github.com/jayson-hu/api-demo-go/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil{
		fmt.Println(err)
	}
}