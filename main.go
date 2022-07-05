package main

import (
	"fmt"
	"github.com/xfyun/asectl/cmd"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	command := cmd.NewDefaultCommand()

	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
