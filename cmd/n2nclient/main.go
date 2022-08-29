package main

import "github.com/gensliu/n2n-go/internal/n2nclient/cmd"

var cfgFilePath = ""

func main() {
	cmd.Execute(cfgFilePath)
}