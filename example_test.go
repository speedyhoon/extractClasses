package extractClasses_test

import (
	"fmt"

	"github.com/speedyhoon/extractClasses"
)

func ExampleExtract() {
	css := "input:out-of-range #id-name#second.third::selection input:out-of-range::selection"
	fmt.Printf("%q", extractClasses.Extract(css))
	// Output: ["#id-name" "#second" ".third"]
}

func ExampleExtractBytes() {
	css := []byte("input:out-of-range #id-name#second.third::selection input:out-of-range::selection")
	fmt.Printf("%q", extractClasses.ExtractBytes(css))
	// Output: ["#id-name" "#second" ".third"]
}
