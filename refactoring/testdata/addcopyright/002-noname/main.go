package main

import "fmt"

func main() { //<<<<<addcopyright,5,1,5,1,,pass
	msg := `This program contains the word Copyright in a string, but this
should not prevent the tool from adding a copyright header.  It will only
refuse to add the header if the word Copyright appears in a comment.
Also, notice that the refactored file is formatted like gofmt.`
	fmt.Println(msg)
}
