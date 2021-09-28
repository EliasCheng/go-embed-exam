package goembedexam

import "embed"

//go:embed embed.txt
var f embed.FS

// Read content of embeded txt file
func Read() (content []byte, err error) {
	return f.ReadFile("embed.txt")
}
