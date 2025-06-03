package main

import (
	"fmt"
	"github.com/iamthefamous/gator.git/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}

	if err := cfg.SetUser("asylbek"); err != nil {
		panic(err)
	}

	updateCfg, err := config.Read()
	if err != nil {
		panic(err)
	}
	fmt.Println(updateCfg)
}
