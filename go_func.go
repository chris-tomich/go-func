package main

import (
	"fmt"
	t "github.com/chris-tomich/go-func/typed_example"
)

func BenchmarkTypedExample() {
	bagOfWords := t.StringList{ "This", "is", "a", "bag", "of", "words." }

	query1 := bagOfWords.Map(func (word string) string { return word + " " })

	query2 := query1.Filter(func (word string) bool { return word != "is" })

	result1 := query1.Reduce("", func (init string, word string) string { return init + word })

	result2 := query2.Reduce("", func (init string, word string) string { return init + word })

	fmt.Println(result1)
	fmt.Println(result2)
}

func main() {
	BenchmarkTypedExample()
}
