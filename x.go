package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: x cmd [args]")
		os.Exit(1)
	}
	name := os.Args[1]
	args := []string{name}
	args = append(args, os.Args[2:]...)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		cmd := exec.Command(name)
		cmd.Args = args
		cmd.Stdin = strings.NewReader(s.Text() + "\n")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}
