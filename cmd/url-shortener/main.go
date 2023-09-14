package main

import (
	"fmt"
	"github/SergeyMoshkov/url-shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
}
