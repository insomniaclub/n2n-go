package main

import (
	"github.com/gensliu/n2n-go/internal/n2nserver/cmd"
)

var cfgFilePath = ""

func main() {
	cmd.Execute(cfgFilePath)
}