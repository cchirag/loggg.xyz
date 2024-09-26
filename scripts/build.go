package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	cwd, err := os.Getwd()

	if err != nil {
		log.Panic(err)
	}
	dest := filepath.Join(cwd, "_site/")

	source := filepath.Join(cwd, "public")

	cmd := exec.Command("cp", "--recursive", source, dest)
	cmd.Run()

}
