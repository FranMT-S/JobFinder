package helpers

import (
	"fmt"
	"os"
)

// saveHTMLResponse is a function that saves the html response to a file for debugging
func SaveHTMLResponse(html string, filename string) {
	file, err := os.Create("tmp/" + filename)
	if err != nil {
		fmt.Println("cannot create the file", err)
		return
	}
	defer file.Close()
	file.WriteString(html)
}
