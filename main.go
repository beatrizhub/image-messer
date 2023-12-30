package main

import (
	"image-messer/images"
)

func main() {
	svc := images.NewService()
	svr := images.NewServer(":8080", svc)

	svr.Run()
}
