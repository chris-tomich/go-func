package benchmarks

import (
	"fmt"
	t "github.com/chris-tomich/go-func/typed_example"
	"time"
	"math/rand"
	"testing"
)

func StandardLoopExample(bagOfWords t.StringList) {
	sentence1, sentence2 := "", ""

	for _, word := range bagOfWords {
		sentence1 = sentence1 + word + " "

		if word != Filter1 && word != Filter2 && word != Filter3 && word != Filter4 {
			sentence2 = sentence2 + word + " "
		}
	}

	fmt.Sprintln(sentence1)
	fmt.Sprintln(sentence2)
}

func TypedExampleWithConcurrency(bagOfWords t.StringList) {
	query1 := bagOfWords.Map(func (word string) string { return word + " " })

	query2 := query1.Filter(func (word string) bool { return word != Filter1 })
	query2 = query2.Filter(func (word string) bool { return word != Filter2 })
	query2 = query2.Filter(func (word string) bool { return word != Filter3 })
	query2 = query2.Filter(func (word string) bool { return word != Filter4 })

	result1 := make(chan string)
	result2 := make(chan string)

	go (func () {
		result1 <- query1.Reduce("", func(init string, word string) string {
			return init + word
		})
	})()

	go (func () {
		result2 <- query2.Reduce("", func(init string, word string) string {
			return init + word
		})
	})()

	sentence1 := <- result1
	sentence2 := <- result2
	fmt.Sprintln(sentence1)
	fmt.Sprintln(sentence2)
}

func TypedExampleWithoutConcurrency(bagOfWords t.StringList) {
	query1 := bagOfWords.Map(func (word string) string { return word + " " })

	query2 := query1.Filter(func (word string) bool { return word != Filter1 })
	query2 = query2.Filter(func (word string) bool { return word != Filter2 })
	query2 = query2.Filter(func (word string) bool { return word != Filter3 })
	query2 = query2.Filter(func (word string) bool { return word != Filter4 })

	sentence1 := query1.Reduce("", func(init string, word string) string { return init + word })

	sentence2 := query2.Reduce("", func(init string, word string) string { return init + word })

	fmt.Sprintln(sentence1)
	fmt.Sprintln(sentence2)
}

func GenerateRandomPosition(maxPosition int) int {
	randSource := rand.NewSource(time.Now().UnixNano())
	randPosition := rand.New(randSource).Int()

	if randPosition > maxPosition {
		randPosition = randPosition % maxPosition
	}

	return randPosition
}

const (
	Filter1 = "one"
	Filter2 = "ten"
	Filter3 = "hundred"
	Filter4 = "thousand"
)

func GenerateBagOfWords() t.StringList {
	bagOfWords := t.StringList{ "Lorem", "ipsum", "dolor", "sit", "amet,", "consectetur", "adipiscing", "elit.", "Sed", "elementum", "lacinia", "nisi,", "vitae", "aliquam", "quam", "aliquam", "eu.", "Proin", "dignissim", "metus", "massa,", "vitae", "scelerisque", "nibh", "malesuada", "sed.", "Etiam", "sodales", "quis", "orci", "eu", "elementum.", "Maecenas", "at", "leo", "ipsum.", "Phasellus", "sed", "maximus", "justo.", "Cras", "dictum", "felis", "id", "mauris", "laoreet", "euismod.", "Duis", "eu", "risus", "id", "enim", "porttitor", "rhoncus.", "Suspendisse", "vulputate", "maximus", "diam,", "ac", "consectetur", "magna.", "Ut", "sit", "amet", "erat", "id", "dui", "consequat", "laoreet.", "Mauris", "eget", "nibh", "vitae", "massa", "commodo", "sagittis", "vel", "vel", "nisl.", "Aliquam", "faucibus", "neque", "ac", "nibh", "pharetra", "pretium.", "Cum", "sociis", "natoque", "penatibus", "et", "magnis", "dis", "parturient", "montes,", "nascetur", "ridiculus", "mus.", "Ut", "mattis", "tempus", "orci,", "et", "ornare", "odio", "ullamcorper", "a.", "Aenean", "nec", "lobortis", "justo.", "Curabitur", "tempus", "arcu", "vel", "libero", "aliquam,", "sed", "scelerisque", "ante", "tincidunt.", "Etiam", "quis", "lectus", "velit.", "Etiam", "vitae", "leo", "nec", "mauris", "iaculis", "vulputate", "tristique", "vel", "nunc.", "Nulla", "facilisis", "feugiat", "mollis.", "Fusce", "elementum", "dignissim", "convallis.", "Aliquam", "erat", "volutpat.", "Proin", "nec", "condimentum", "arcu,", "a", "congue", "mauris.", "Nunc", "est", "ante,", "accumsan", "id", "lacinia", "sed,", "imperdiet", "porttitor", "urna.", "Duis", "risus", "tortor,", "scelerisque", "sed", "volutpat", "a,", "scelerisque", "in", "leo.", "Nullam", "ipsum", "urna,", "maximus", "ac", "egestas", "id,", "auctor", "in", "augue.", "Praesent", "pellentesque", "nec", "risus", "eu", "hendrerit.", "Vestibulum", "ante", "ipsum", "primis", "in", "faucibus", "orci", "luctus", "et", "ultrices", "posuere", "cubilia", "Curae;", "Curabitur", "est", "magna,", "vestibulum", "et", "malesuada", "quis,", "dapibus", "nec", "nibh.", "Vivamus", "eu", "tristique", "orci.", "Etiam", "convallis", "est", "arcu,", "in", "tempor", "sem", "lacinia", "ut.", "Sed", "in", "orci", "pharetra,", "suscipit", "justo", "a,", "dictum", "odio.", "Integer", "pulvinar", "est", "in", "porta", "pellentesque.", "Sed", "suscipit", "placerat", "libero", "quis", "placerat.", "Duis", "tincidunt", "nulla", "arcu,", "vel", "vestibulum", "est", "elementum", "eu.", "Mauris", "sit", "amet", "bibendum", "eros,", "eu", "tempus", "mauris.", "Etiam", "ex", "mauris,", "ullamcorper", "volutpat", "dui", "et,", "scelerisque", "efficitur", "ex.", "Nam", "dignissim", "vehicula", "pulvinar.", "Nam", "libero", "ligula,", "ultricies", "vel", "cursus", "eu,", "iaculis", "ac", "nibh.", "Sed", "eget", "lectus", "sagittis,", "placerat", "neque", "a,", "gravida", "arcu.", "Curabitur", "nisi", "enim,", "lobortis", "mollis", "odio", "ac,", "posuere", "volutpat", "enim.", "Phasellus", "sollicitudin", "eros", "mollis", "est", "venenatis", "elementum.", "Cras", "tempor", "luctus", "vestibulum.", "Suspendisse", "dapibus", "magna", "et", "orci", "pretium,", "vel", "auctor", "mi", "cursus.", "Nam", "ultrices,", "tortor", "eu", "convallis", "porttitor,", "magna", "turpis", "fringilla", "ipsum,", "et", "blandit", "nibh", "lacus", "eu", "eros.", "Proin", "ac", "enim", "vel", "diam", "aliquam", "ullamcorper", "eget", "in", "lorem.", "Vivamus", "arcu", "sapien,", "facilisis", "vel", "dapibus", "nec,", "mollis", "eget", "magna.", "Maecenas", "hendrerit", "ligula", "lacinia,", "tincidunt", "neque", "id,", "consectetur", "enim.", "Quisque", "in", "nibh", "efficitur,", "pulvinar", "magna", "id,", "vulputate", "quam.", "Cras", "a", "dictum", "magna.", "Nam", "iaculis", "justo", "id", "metus", "aliquam,", "quis", "tincidunt", "ante", "consequat.", "Sed", "congue", "id", "velit", "sit", "amet", "convallis.", "Proin", "quis", "tortor", "sit", "amet", "purus", "fermentum", "rutrum", "at", "id", "ligula.", "Interdum", "et", "malesuada", "fames", "ac", "ante", "ipsum", "primis", "in", "faucibus.", "Cras", "quis", "sollicitudin", "urna.", "Quisque", "quis", "velit", "at", "nibh", "sodales", "tempus", "et", "egestas", "diam.", "Suspendisse", "in", "mattis", "lectus.", "Nulla", "orci", "arcu,", "vestibulum", "eu", "turpis", "quis,", "pulvinar", "aliquam", "metus.", "Integer", "facilisis", "ipsum", "augue,", "at", "rutrum", "nisi", "cursus", "et.", "Curabitur", "mauris", "libero,", "varius", "ac", "laoreet", "et,", "dignissim", "molestie", "magna.", "Vivamus", "vel", "ex", "a", "magna", "porta", "dapibus", "eu", "at", "magna.", "Aenean", "accumsan,", "dolor", "quis", "vulputate", "tempus,", "metus", "diam", "consectetur", "nibh,", "vel", "dignissim", "velit", "ex", "quis", "massa.", "Pellentesque", "vitae", "consectetur", "mauris.", "In", "eu", "tempus", "sapien.", "Ut", "sed", "sapien", "venenatis,", "malesuada", "magna", "id,", "fringilla", "elit.", "Curabitur", "aliquet", "enim", "at", "orci", "pharetra", "rhoncus", "sed", "vel", "dolor.", "In", "ultricies,", "enim", "nec", "hendrerit", "ultricies,", "dolor", "libero", "tempor", "nibh,", "at", "imperdiet", "ligula", "erat", "eget", "nulla.", "Fusce", "ut", "felis", "id", "quam", "finibus", "mattis", "a", "vitae", "est.", "Pellentesque", "sit", "amet", "varius", "eros.", "Mauris", "dapibus", "sem", "metus,", "a", "maximus", "velit", "dignissim", "ac.", "Aenean", "volutpat", "laoreet", "felis", "ac", "pulvinar.", "Praesent", "enim", "sem,", "congue", "quis", "porta", "sit", "amet,", "ullamcorper", "egestas", "lacus.", "Maecenas", "quam", "nisl,", "eleifend", "ac", "justo", "id,", "mattis", "laoreet", "odio.", "Sed", "eget", "odio", "ullamcorper,", "semper", "tellus", "vitae,", "vulputate", "ante.", "Nam", "vitae", "nunc", "rhoncus,", "pulvinar", "elit", "aliquet,", "finibus", "massa.", "Nullam", "quis", "convallis", "augue.", "Maecenas", "lacinia", "enim", "orci,", "a", "lobortis", "eros", "pretium", "non.", "Praesent", "rhoncus", "porta", "magna.", "Curabitur", "dictum", "elementum", "lectus,", "ultricies", "finibus", "eros", "elementum", "id.", "Praesent", "non", "porttitor", "diam,", "et", "vehicula", "massa.", "Fusce", "aliquet", "nisi", "eget", "tincidunt", "maximus.", "Sed", "lacinia", "augue", "sollicitudin", "ornare", "cursus.", "Lorem", "ipsum", "dolor", "sit", "amet,", "consectetur", "adipiscing", "elit.", "In", "ligula", "lectus,", "ullamcorper", "nec", "interdum", "in,", "pharetra", "nec", "orci.", "Nulla", "luctus", "et", "urna", "vel", "ultrices.", "Ut", "sit", "amet", "tincidunt", "purus.", "Phasellus", "efficitur,", "quam", "sit", "amet", "sagittis", "malesuada,", "sem", "est", "vulputate", "enim,", "ac", "finibus", "lacus", "risus", "in", "libero.", "Mauris", "ac", "orci", "cursus,", "lobortis", "ante", "eget,", "tempus", "velit.", "Vestibulum", "at", "sagittis", "dui.", "Sed", "sit", "amet", "euismod", "ante.", "Curabitur", "in", "purus", "at", "erat", "lobortis", "accumsan.", "Sed", "ut", "venenatis", "nibh.", "Duis", "libero", "orci,", "bibendum", "id", "volutpat", "ut,", "volutpat", "et", "lacus.", "Mauris", "id", "sodales", "orci.", "Vestibulum", "eleifend,", "lorem", "non", "egestas", "eleifend,", "magna", "arcu", "placerat", "orci,", "vitae", "euismod", "dui", "nunc", "ut", "augue.", "Donec", "commodo", "nisi", "in", "scelerisque", "ornare.", "Etiam", "sed", "ligula", "posuere", "augue", "auctor", "ultrices", "in", "ut", "nibh.", "Maecenas", "risus", "libero,", "aliquam", "eu", "sodales", "gravida,", "auctor", "a", "purus.", "In", "hac", "habitasse", "platea", "dictumst.", "Proin", "nec", "elementum", "nulla.", "Suspendisse", "at", "dictum", "purus.", "Cras", "dictum", "vel", "velit", "in", "congue.", "Morbi", "ut", "vestibulum", "lacus.", "Curabitur", "ac", "magna", "ex.", "Nulla", "fringilla", "dapibus", "lacus,", "eu", "commodo", "nisi.", "Vestibulum", "dignissim", "vitae", "sem", "molestie", "ullamcorper.", "Phasellus", "nibh", "libero,", "aliquet", "vitae", "nisl", "id,", "mattis", "rhoncus", "sapien.", "Etiam", "fringilla,", "lectus", "a", "cursus", "malesuada,", "diam", "metus", "auctor", "magna,", "nec", "molestie", "sem", "metus", "in", "felis.", "Nunc", "nec", "malesuada", "metus.", "Phasellus", "sit", "amet", "eros", "vitae", "orci", "scelerisque", "malesuada.", "Mauris", "tempus", "eget", "quam", "ut", "ultrices.", "Cras", "nec", "arcu", "tellus.", "Nullam", "volutpat", "nunc", "vitae", "lorem", "elementum", "suscipit.", "Aenean", "auctor", "sed", "felis", "nec", "venenatis.", "Integer", "congue,", "dolor", "a", "tristique", "pellentesque,", "ante", "justo", "sollicitudin", "nisl,", "in", "tempor", "erat", "orci", "et", "ex.", "Curabitur", "eu", "commodo", "nisi.", "Pellentesque", "habitant", "morbi", "tristique", "senectus", "et", "netus", "et", "malesuada", "fames", "ac", "turpis", "egestas.", "Aliquam", "sagittis", "erat", "felis,", "nec", "mollis", "justo", "dignissim", "ac.", "Nunc", "finibus", "leo", "turpis,", "sit", "amet", "dapibus", "nunc", "volutpat", "ut.", "Nullam", "convallis", "porta", "semper.", "Integer", "scelerisque,", "sapien", "et", "sodales", "hendrerit,", "tellus", "quam", "vulputate", "massa,", "non", "tempus", "enim", "lectus", "at", "ex.", "Pellentesque", "habitant", "morbi", "tristique", "senectus", "et", "netus", "et", "malesuada", "fames", "ac", "turpis", "egestas.", "Nam", "mattis", "iaculis", "quam", "id", "vestibulum.", "Donec", "pharetra", "euismod", "diam", "ut", "gravida.", "Aenean", "at", "felis", "molestie,", "tempor", "sapien", "ac,", "euismod", "nisi.", "In", "aliquet", "rutrum", "sodales.", "Sed", "eleifend", "mollis", "eros,", "nec." }

	bagSize := len(bagOfWords) - 1

	bagOfWords[GenerateRandomPosition(bagSize)] = Filter1
	bagOfWords[GenerateRandomPosition(bagSize)] = Filter1
	bagOfWords[GenerateRandomPosition(bagSize)] = Filter1
	bagOfWords[GenerateRandomPosition(bagSize)] = Filter2
	bagOfWords[GenerateRandomPosition(bagSize)] = Filter2
	bagOfWords[GenerateRandomPosition(bagSize)] = Filter2
	bagOfWords[GenerateRandomPosition(bagSize)] = Filter3
	bagOfWords[GenerateRandomPosition(bagSize)] = Filter3
	bagOfWords[GenerateRandomPosition(bagSize)] = Filter3
	bagOfWords[GenerateRandomPosition(bagSize)] = Filter4
	bagOfWords[GenerateRandomPosition(bagSize)] = Filter4
	bagOfWords[GenerateRandomPosition(bagSize)] = Filter4

	return bagOfWords
}

func BenchmarkTypedExampleWithConcurrency(b *testing.B) {
	bagOfWords := GenerateBagOfWords()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TypedExampleWithConcurrency(bagOfWords)
	}
}

func BenchmarkTypedExampleWithoutConcurrency(b *testing.B) {
	bagOfWords := GenerateBagOfWords()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TypedExampleWithoutConcurrency(bagOfWords)
	}
}

func BenchmarkStandardLoopExample(b *testing.B) {
	bagOfWords := GenerateBagOfWords()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StandardLoopExample(bagOfWords)
	}
}
