package low_alloc

import (
	"fmt"
	t "github.com/chris-tomich/go-func/typed_example/low_alloc"
	"time"
	"math/rand"
	"testing"
)

func StandardLoopExampleWithConcurrency(bagOfWords t.StringList) {
	result1 := make(chan []byte)
	result2 := make(chan []byte)

	go (func () {
		generatedSentence1 := []byte("")

		for _, word := range bagOfWords {
			generatedSentence1 = append(generatedSentence1, append(word, []byte(" ")...)...)
		}

		result1 <- generatedSentence1
	})()

	go (func () {
		generatedSentence2 := []byte("")

		for _, word := range bagOfWords {
			if string(word) != Filter1 && string(word) != Filter2 && string(word) != Filter3 && string(word) != Filter4 {
				generatedSentence2 = append(generatedSentence2, append(word, []byte(" ")...)...)
			}
		}

		result2 <- generatedSentence2
	})()

	sentence1 := string(<- result1)
	sentence2 := string(<- result2)

	fmt.Sprintln(sentence1)
	fmt.Sprintln(sentence2)
}

func StandardLoopExampleWithoutConcurrency(bagOfWords t.StringList) {
	sentence1, sentence2 := []byte(""), []byte("")

	for _, word := range bagOfWords {
		sentence1 = append(sentence1, append(word, []byte(" ")...)...)

		if string(word) != Filter1 && string(word) != Filter2 && string(word) != Filter3 && string(word) != Filter4 {
			sentence2 = append(sentence2, append(word, []byte(" ")...)...)
		}
	}

	fmt.Sprintln(string(sentence1))
	fmt.Sprintln(string(sentence2))
}

func TypedExampleWithConcurrency(bagOfWords t.StringList) {
	query1 := bagOfWords.Map(func (word []byte) []byte { return append(word, []byte(" ")...) })

	query2 := query1.Filter(func (word []byte) bool { return string(word) != Filter1 })
	query2 = query2.Filter(func (word []byte) bool { return string(word) != Filter2 })
	query2 = query2.Filter(func (word []byte) bool { return string(word) != Filter3 })
	query2 = query2.Filter(func (word []byte) bool { return string(word) != Filter4 })

	result1 := make(chan []byte)
	result2 := make(chan []byte)

	go (func () {
		result1 <- query1.Reduce([]byte(""), func(init []byte, word []byte) []byte {
			return append(init, word...)
		})
	})()

	go (func () {
		result2 <- query2.Reduce([]byte(""), func(init []byte, word []byte) []byte {
			return append(init, word...)
		})
	})()

	sentence1 := string(<- result1)
	sentence2 := string(<- result2)
	fmt.Sprintln(sentence1)
	fmt.Sprintln(sentence2)
}

func TypedExampleWithoutConcurrency(bagOfWords t.StringList) {
	query1 := bagOfWords.Map(func (word []byte) []byte { return append(word, []byte(" ")...) })

	query2 := query1.Filter(func (word []byte) bool { return string(word) != Filter1 })
	query2 = query2.Filter(func (word []byte) bool { return string(word) != Filter2 })
	query2 = query2.Filter(func (word []byte) bool { return string(word) != Filter3 })
	query2 = query2.Filter(func (word []byte) bool { return string(word) != Filter4 })

	sentence1 := string(query1.Reduce([]byte(""), func(init []byte, word []byte) []byte { return append(init, word...) }))

	sentence2 := string(query2.Reduce([]byte(""), func(init []byte, word []byte) []byte { return append(init, word...) }))

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
	bagOfWords := t.StringList{ []byte("Lorem"), []byte("ipsum"), []byte("dolor"), []byte("sit"), []byte("amet,"), []byte("consectetur"), []byte("adipiscing"), []byte("elit."), []byte("Sed"), []byte("elementum"), []byte("lacinia"), []byte("nisi,"), []byte("vitae"), []byte("aliquam"), []byte("quam"), []byte("aliquam"), []byte("eu."), []byte("Proin"), []byte("dignissim"), []byte("metus"), []byte("massa,"), []byte("vitae"), []byte("scelerisque"), []byte("nibh"), []byte("malesuada"), []byte("sed."), []byte("Etiam"), []byte("sodales"), []byte("quis"), []byte("orci"), []byte("eu"), []byte("elementum."), []byte("Maecenas"), []byte("at"), []byte("leo"), []byte("ipsum."), []byte("Phasellus"), []byte("sed"), []byte("maximus"), []byte("justo."), []byte("Cras"), []byte("dictum"), []byte("felis"), []byte("id"), []byte("mauris"), []byte("laoreet"), []byte("euismod."), []byte("Duis"), []byte("eu"), []byte("risus"), []byte("id"), []byte("enim"), []byte("porttitor"), []byte("rhoncus."), []byte("Suspendisse"), []byte("vulputate"), []byte("maximus"), []byte("diam,"), []byte("ac"), []byte("consectetur"), []byte("magna."), []byte("Ut"), []byte("sit"), []byte("amet"), []byte("erat"), []byte("id"), []byte("dui"), []byte("consequat"), []byte("laoreet."), []byte("Mauris"), []byte("eget"), []byte("nibh"), []byte("vitae"), []byte("massa"), []byte("commodo"), []byte("sagittis"), []byte("vel"), []byte("vel"), []byte("nisl."), []byte("Aliquam"), []byte("faucibus"), []byte("neque"), []byte("ac"), []byte("nibh"), []byte("pharetra"), []byte("pretium."), []byte("Cum"), []byte("sociis"), []byte("natoque"), []byte("penatibus"), []byte("et"), []byte("magnis"), []byte("dis"), []byte("parturient"), []byte("montes,"), []byte("nascetur"), []byte("ridiculus"), []byte("mus."), []byte("Ut"), []byte("mattis"), []byte("tempus"), []byte("orci,"), []byte("et"), []byte("ornare"), []byte("odio"), []byte("ullamcorper"), []byte("a."), []byte("Aenean"), []byte("nec"), []byte("lobortis"), []byte("justo."), []byte("Curabitur"), []byte("tempus"), []byte("arcu"), []byte("vel"), []byte("libero"), []byte("aliquam,"), []byte("sed"), []byte("scelerisque"), []byte("ante"), []byte("tincidunt."), []byte("Etiam"), []byte("quis"), []byte("lectus"), []byte("velit."), []byte("Etiam"), []byte("vitae"), []byte("leo"), []byte("nec"), []byte("mauris"), []byte("iaculis"), []byte("vulputate"), []byte("tristique"), []byte("vel"), []byte("nunc."), []byte("Nulla"), []byte("facilisis"), []byte("feugiat"), []byte("mollis."), []byte("Fusce"), []byte("elementum"), []byte("dignissim"), []byte("convallis."), []byte("Aliquam"), []byte("erat"), []byte("volutpat."), []byte("Proin"), []byte("nec"), []byte("condimentum"), []byte("arcu,"), []byte("a"), []byte("congue"), []byte("mauris."), []byte("Nunc"), []byte("est"), []byte("ante,"), []byte("accumsan"), []byte("id"), []byte("lacinia"), []byte("sed,"), []byte("imperdiet"), []byte("porttitor"), []byte("urna."), []byte("Duis"), []byte("risus"), []byte("tortor,"), []byte("scelerisque"), []byte("sed"), []byte("volutpat"), []byte("a,"), []byte("scelerisque"), []byte("in"), []byte("leo."), []byte("Nullam"), []byte("ipsum"), []byte("urna,"), []byte("maximus"), []byte("ac"), []byte("egestas"), []byte("id,"), []byte("auctor"), []byte("in"), []byte("augue."), []byte("Praesent"), []byte("pellentesque"), []byte("nec"), []byte("risus"), []byte("eu"), []byte("hendrerit."), []byte("Vestibulum"), []byte("ante"), []byte("ipsum"), []byte("primis"), []byte("in"), []byte("faucibus"), []byte("orci"), []byte("luctus"), []byte("et"), []byte("ultrices"), []byte("posuere"), []byte("cubilia"), []byte("Curae;"), []byte("Curabitur"), []byte("est"), []byte("magna,"), []byte("vestibulum"), []byte("et"), []byte("malesuada"), []byte("quis,"), []byte("dapibus"), []byte("nec"), []byte("nibh."), []byte("Vivamus"), []byte("eu"), []byte("tristique"), []byte("orci."), []byte("Etiam"), []byte("convallis"), []byte("est"), []byte("arcu,"), []byte("in"), []byte("tempor"), []byte("sem"), []byte("lacinia"), []byte("ut."), []byte("Sed"), []byte("in"), []byte("orci"), []byte("pharetra,"), []byte("suscipit"), []byte("justo"), []byte("a,"), []byte("dictum"), []byte("odio."), []byte("Integer"), []byte("pulvinar"), []byte("est"), []byte("in"), []byte("porta"), []byte("pellentesque."), []byte("Sed"), []byte("suscipit"), []byte("placerat"), []byte("libero"), []byte("quis"), []byte("placerat."), []byte("Duis"), []byte("tincidunt"), []byte("nulla"), []byte("arcu,"), []byte("vel"), []byte("vestibulum"), []byte("est"), []byte("elementum"), []byte("eu."), []byte("Mauris"), []byte("sit"), []byte("amet"), []byte("bibendum"), []byte("eros,"), []byte("eu"), []byte("tempus"), []byte("mauris."), []byte("Etiam"), []byte("ex"), []byte("mauris,"), []byte("ullamcorper"), []byte("volutpat"), []byte("dui"), []byte("et,"), []byte("scelerisque"), []byte("efficitur"), []byte("ex."), []byte("Nam"), []byte("dignissim"), []byte("vehicula"), []byte("pulvinar."), []byte("Nam"), []byte("libero"), []byte("ligula,"), []byte("ultricies"), []byte("vel"), []byte("cursus"), []byte("eu,"), []byte("iaculis"), []byte("ac"), []byte("nibh."), []byte("Sed"), []byte("eget"), []byte("lectus"), []byte("sagittis,"), []byte("placerat"), []byte("neque"), []byte("a,"), []byte("gravida"), []byte("arcu."), []byte("Curabitur"), []byte("nisi"), []byte("enim,"), []byte("lobortis"), []byte("mollis"), []byte("odio"), []byte("ac,"), []byte("posuere"), []byte("volutpat"), []byte("enim."), []byte("Phasellus"), []byte("sollicitudin"), []byte("eros"), []byte("mollis"), []byte("est"), []byte("venenatis"), []byte("elementum."), []byte("Cras"), []byte("tempor"), []byte("luctus"), []byte("vestibulum."), []byte("Suspendisse"), []byte("dapibus"), []byte("magna"), []byte("et"), []byte("orci"), []byte("pretium,"), []byte("vel"), []byte("auctor"), []byte("mi"), []byte("cursus."), []byte("Nam"), []byte("ultrices,"), []byte("tortor"), []byte("eu"), []byte("convallis"), []byte("porttitor,"), []byte("magna"), []byte("turpis"), []byte("fringilla"), []byte("ipsum,"), []byte("et"), []byte("blandit"), []byte("nibh"), []byte("lacus"), []byte("eu"), []byte("eros."), []byte("Proin"), []byte("ac"), []byte("enim"), []byte("vel"), []byte("diam"), []byte("aliquam"), []byte("ullamcorper"), []byte("eget"), []byte("in"), []byte("lorem."), []byte("Vivamus"), []byte("arcu"), []byte("sapien,"), []byte("facilisis"), []byte("vel"), []byte("dapibus"), []byte("nec,"), []byte("mollis"), []byte("eget"), []byte("magna."), []byte("Maecenas"), []byte("hendrerit"), []byte("ligula"), []byte("lacinia,"), []byte("tincidunt"), []byte("neque"), []byte("id,"), []byte("consectetur"), []byte("enim."), []byte("Quisque"), []byte("in"), []byte("nibh"), []byte("efficitur,"), []byte("pulvinar"), []byte("magna"), []byte("id,"), []byte("vulputate"), []byte("quam."), []byte("Cras"), []byte("a"), []byte("dictum"), []byte("magna."), []byte("Nam"), []byte("iaculis"), []byte("justo"), []byte("id"), []byte("metus"), []byte("aliquam,"), []byte("quis"), []byte("tincidunt"), []byte("ante"), []byte("consequat."), []byte("Sed"), []byte("congue"), []byte("id"), []byte("velit"), []byte("sit"), []byte("amet"), []byte("convallis."), []byte("Proin"), []byte("quis"), []byte("tortor"), []byte("sit"), []byte("amet"), []byte("purus"), []byte("fermentum"), []byte("rutrum"), []byte("at"), []byte("id"), []byte("ligula."), []byte("Interdum"), []byte("et"), []byte("malesuada"), []byte("fames"), []byte("ac"), []byte("ante"), []byte("ipsum"), []byte("primis"), []byte("in"), []byte("faucibus."), []byte("Cras"), []byte("quis"), []byte("sollicitudin"), []byte("urna."), []byte("Quisque"), []byte("quis"), []byte("velit"), []byte("at"), []byte("nibh"), []byte("sodales"), []byte("tempus"), []byte("et"), []byte("egestas"), []byte("diam."), []byte("Suspendisse"), []byte("in"), []byte("mattis"), []byte("lectus."), []byte("Nulla"), []byte("orci"), []byte("arcu,"), []byte("vestibulum"), []byte("eu"), []byte("turpis"), []byte("quis,"), []byte("pulvinar"), []byte("aliquam"), []byte("metus."), []byte("Integer"), []byte("facilisis"), []byte("ipsum"), []byte("augue,"), []byte("at"), []byte("rutrum"), []byte("nisi"), []byte("cursus"), []byte("et."), []byte("Curabitur"), []byte("mauris"), []byte("libero,"), []byte("varius"), []byte("ac"), []byte("laoreet"), []byte("et,"), []byte("dignissim"), []byte("molestie"), []byte("magna."), []byte("Vivamus"), []byte("vel"), []byte("ex"), []byte("a"), []byte("magna"), []byte("porta"), []byte("dapibus"), []byte("eu"), []byte("at"), []byte("magna."), []byte("Aenean"), []byte("accumsan,"), []byte("dolor"), []byte("quis"), []byte("vulputate"), []byte("tempus,"), []byte("metus"), []byte("diam"), []byte("consectetur"), []byte("nibh,"), []byte("vel"), []byte("dignissim"), []byte("velit"), []byte("ex"), []byte("quis"), []byte("massa."), []byte("Pellentesque"), []byte("vitae"), []byte("consectetur"), []byte("mauris."), []byte("In"), []byte("eu"), []byte("tempus"), []byte("sapien."), []byte("Ut"), []byte("sed"), []byte("sapien"), []byte("venenatis,"), []byte("malesuada"), []byte("magna"), []byte("id,"), []byte("fringilla"), []byte("elit."), []byte("Curabitur"), []byte("aliquet"), []byte("enim"), []byte("at"), []byte("orci"), []byte("pharetra"), []byte("rhoncus"), []byte("sed"), []byte("vel"), []byte("dolor."), []byte("In"), []byte("ultricies,"), []byte("enim"), []byte("nec"), []byte("hendrerit"), []byte("ultricies,"), []byte("dolor"), []byte("libero"), []byte("tempor"), []byte("nibh,"), []byte("at"), []byte("imperdiet"), []byte("ligula"), []byte("erat"), []byte("eget"), []byte("nulla."), []byte("Fusce"), []byte("ut"), []byte("felis"), []byte("id"), []byte("quam"), []byte("finibus"), []byte("mattis"), []byte("a"), []byte("vitae"), []byte("est."), []byte("Pellentesque"), []byte("sit"), []byte("amet"), []byte("varius"), []byte("eros."), []byte("Mauris"), []byte("dapibus"), []byte("sem"), []byte("metus,"), []byte("a"), []byte("maximus"), []byte("velit"), []byte("dignissim"), []byte("ac."), []byte("Aenean"), []byte("volutpat"), []byte("laoreet"), []byte("felis"), []byte("ac"), []byte("pulvinar."), []byte("Praesent"), []byte("enim"), []byte("sem,"), []byte("congue"), []byte("quis"), []byte("porta"), []byte("sit"), []byte("amet,"), []byte("ullamcorper"), []byte("egestas"), []byte("lacus."), []byte("Maecenas"), []byte("quam"), []byte("nisl,"), []byte("eleifend"), []byte("ac"), []byte("justo"), []byte("id,"), []byte("mattis"), []byte("laoreet"), []byte("odio."), []byte("Sed"), []byte("eget"), []byte("odio"), []byte("ullamcorper,"), []byte("semper"), []byte("tellus"), []byte("vitae,"), []byte("vulputate"), []byte("ante."), []byte("Nam"), []byte("vitae"), []byte("nunc"), []byte("rhoncus,"), []byte("pulvinar"), []byte("elit"), []byte("aliquet,"), []byte("finibus"), []byte("massa."), []byte("Nullam"), []byte("quis"), []byte("convallis"), []byte("augue."), []byte("Maecenas"), []byte("lacinia"), []byte("enim"), []byte("orci,"), []byte("a"), []byte("lobortis"), []byte("eros"), []byte("pretium"), []byte("non."), []byte("Praesent"), []byte("rhoncus"), []byte("porta"), []byte("magna."), []byte("Curabitur"), []byte("dictum"), []byte("elementum"), []byte("lectus,"), []byte("ultricies"), []byte("finibus"), []byte("eros"), []byte("elementum"), []byte("id."), []byte("Praesent"), []byte("non"), []byte("porttitor"), []byte("diam,"), []byte("et"), []byte("vehicula"), []byte("massa."), []byte("Fusce"), []byte("aliquet"), []byte("nisi"), []byte("eget"), []byte("tincidunt"), []byte("maximus."), []byte("Sed"), []byte("lacinia"), []byte("augue"), []byte("sollicitudin"), []byte("ornare"), []byte("cursus."), []byte("Lorem"), []byte("ipsum"), []byte("dolor"), []byte("sit"), []byte("amet,"), []byte("consectetur"), []byte("adipiscing"), []byte("elit."), []byte("In"), []byte("ligula"), []byte("lectus,"), []byte("ullamcorper"), []byte("nec"), []byte("interdum"), []byte("in,"), []byte("pharetra"), []byte("nec"), []byte("orci."), []byte("Nulla"), []byte("luctus"), []byte("et"), []byte("urna"), []byte("vel"), []byte("ultrices."), []byte("Ut"), []byte("sit"), []byte("amet"), []byte("tincidunt"), []byte("purus."), []byte("Phasellus"), []byte("efficitur,"), []byte("quam"), []byte("sit"), []byte("amet"), []byte("sagittis"), []byte("malesuada,"), []byte("sem"), []byte("est"), []byte("vulputate"), []byte("enim,"), []byte("ac"), []byte("finibus"), []byte("lacus"), []byte("risus"), []byte("in"), []byte("libero."), []byte("Mauris"), []byte("ac"), []byte("orci"), []byte("cursus,"), []byte("lobortis"), []byte("ante"), []byte("eget,"), []byte("tempus"), []byte("velit."), []byte("Vestibulum"), []byte("at"), []byte("sagittis"), []byte("dui."), []byte("Sed"), []byte("sit"), []byte("amet"), []byte("euismod"), []byte("ante."), []byte("Curabitur"), []byte("in"), []byte("purus"), []byte("at"), []byte("erat"), []byte("lobortis"), []byte("accumsan."), []byte("Sed"), []byte("ut"), []byte("venenatis"), []byte("nibh."), []byte("Duis"), []byte("libero"), []byte("orci,"), []byte("bibendum"), []byte("id"), []byte("volutpat"), []byte("ut,"), []byte("volutpat"), []byte("et"), []byte("lacus."), []byte("Mauris"), []byte("id"), []byte("sodales"), []byte("orci."), []byte("Vestibulum"), []byte("eleifend,"), []byte("lorem"), []byte("non"), []byte("egestas"), []byte("eleifend,"), []byte("magna"), []byte("arcu"), []byte("placerat"), []byte("orci,"), []byte("vitae"), []byte("euismod"), []byte("dui"), []byte("nunc"), []byte("ut"), []byte("augue."), []byte("Donec"), []byte("commodo"), []byte("nisi"), []byte("in"), []byte("scelerisque"), []byte("ornare."), []byte("Etiam"), []byte("sed"), []byte("ligula"), []byte("posuere"), []byte("augue"), []byte("auctor"), []byte("ultrices"), []byte("in"), []byte("ut"), []byte("nibh."), []byte("Maecenas"), []byte("risus"), []byte("libero,"), []byte("aliquam"), []byte("eu"), []byte("sodales"), []byte("gravida,"), []byte("auctor"), []byte("a"), []byte("purus."), []byte("In"), []byte("hac"), []byte("habitasse"), []byte("platea"), []byte("dictumst."), []byte("Proin"), []byte("nec"), []byte("elementum"), []byte("nulla."), []byte("Suspendisse"), []byte("at"), []byte("dictum"), []byte("purus."), []byte("Cras"), []byte("dictum"), []byte("vel"), []byte("velit"), []byte("in"), []byte("congue."), []byte("Morbi"), []byte("ut"), []byte("vestibulum"), []byte("lacus."), []byte("Curabitur"), []byte("ac"), []byte("magna"), []byte("ex."), []byte("Nulla"), []byte("fringilla"), []byte("dapibus"), []byte("lacus,"), []byte("eu"), []byte("commodo"), []byte("nisi."), []byte("Vestibulum"), []byte("dignissim"), []byte("vitae"), []byte("sem"), []byte("molestie"), []byte("ullamcorper."), []byte("Phasellus"), []byte("nibh"), []byte("libero,"), []byte("aliquet"), []byte("vitae"), []byte("nisl"), []byte("id,"), []byte("mattis"), []byte("rhoncus"), []byte("sapien."), []byte("Etiam"), []byte("fringilla,"), []byte("lectus"), []byte("a"), []byte("cursus"), []byte("malesuada,"), []byte("diam"), []byte("metus"), []byte("auctor"), []byte("magna,"), []byte("nec"), []byte("molestie"), []byte("sem"), []byte("metus"), []byte("in"), []byte("felis."), []byte("Nunc"), []byte("nec"), []byte("malesuada"), []byte("metus."), []byte("Phasellus"), []byte("sit"), []byte("amet"), []byte("eros"), []byte("vitae"), []byte("orci"), []byte("scelerisque"), []byte("malesuada."), []byte("Mauris"), []byte("tempus"), []byte("eget"), []byte("quam"), []byte("ut"), []byte("ultrices."), []byte("Cras"), []byte("nec"), []byte("arcu"), []byte("tellus."), []byte("Nullam"), []byte("volutpat"), []byte("nunc"), []byte("vitae"), []byte("lorem"), []byte("elementum"), []byte("suscipit."), []byte("Aenean"), []byte("auctor"), []byte("sed"), []byte("felis"), []byte("nec"), []byte("venenatis."), []byte("Integer"), []byte("congue,"), []byte("dolor"), []byte("a"), []byte("tristique"), []byte("pellentesque,"), []byte("ante"), []byte("justo"), []byte("sollicitudin"), []byte("nisl,"), []byte("in"), []byte("tempor"), []byte("erat"), []byte("orci"), []byte("et"), []byte("ex."), []byte("Curabitur"), []byte("eu"), []byte("commodo"), []byte("nisi."), []byte("Pellentesque"), []byte("habitant"), []byte("morbi"), []byte("tristique"), []byte("senectus"), []byte("et"), []byte("netus"), []byte("et"), []byte("malesuada"), []byte("fames"), []byte("ac"), []byte("turpis"), []byte("egestas."), []byte("Aliquam"), []byte("sagittis"), []byte("erat"), []byte("felis,"), []byte("nec"), []byte("mollis"), []byte("justo"), []byte("dignissim"), []byte("ac."), []byte("Nunc"), []byte("finibus"), []byte("leo"), []byte("turpis,"), []byte("sit"), []byte("amet"), []byte("dapibus"), []byte("nunc"), []byte("volutpat"), []byte("ut."), []byte("Nullam"), []byte("convallis"), []byte("porta"), []byte("semper."), []byte("Integer"), []byte("scelerisque,"), []byte("sapien"), []byte("et"), []byte("sodales"), []byte("hendrerit,"), []byte("tellus"), []byte("quam"), []byte("vulputate"), []byte("massa,"), []byte("non"), []byte("tempus"), []byte("enim"), []byte("lectus"), []byte("at"), []byte("ex."), []byte("Pellentesque"), []byte("habitant"), []byte("morbi"), []byte("tristique"), []byte("senectus"), []byte("et"), []byte("netus"), []byte("et"), []byte("malesuada"), []byte("fames"), []byte("ac"), []byte("turpis"), []byte("egestas."), []byte("Nam"), []byte("mattis"), []byte("iaculis"), []byte("quam"), []byte("id"), []byte("vestibulum."), []byte("Donec"), []byte("pharetra"), []byte("euismod"), []byte("diam"), []byte("ut"), []byte("gravida."), []byte("Aenean"), []byte("at"), []byte("felis"), []byte("molestie,"), []byte("tempor"), []byte("sapien"), []byte("ac,"), []byte("euismod"), []byte("nisi."), []byte("In"), []byte("aliquet"), []byte("rutrum"), []byte("sodales."), []byte("Sed"), []byte("eleifend"), []byte("mollis"), []byte("eros,"), []byte("nec.") }

	bagSize := len(bagOfWords) - 1

	bagOfWords[GenerateRandomPosition(bagSize)] = []byte(Filter1)
	bagOfWords[GenerateRandomPosition(bagSize)] = []byte(Filter1)
	bagOfWords[GenerateRandomPosition(bagSize)] = []byte(Filter1)
	bagOfWords[GenerateRandomPosition(bagSize)] = []byte(Filter2)
	bagOfWords[GenerateRandomPosition(bagSize)] = []byte(Filter2)
	bagOfWords[GenerateRandomPosition(bagSize)] = []byte(Filter2)
	bagOfWords[GenerateRandomPosition(bagSize)] = []byte(Filter3)
	bagOfWords[GenerateRandomPosition(bagSize)] = []byte(Filter3)
	bagOfWords[GenerateRandomPosition(bagSize)] = []byte(Filter3)
	bagOfWords[GenerateRandomPosition(bagSize)] = []byte(Filter4)
	bagOfWords[GenerateRandomPosition(bagSize)] = []byte(Filter4)
	bagOfWords[GenerateRandomPosition(bagSize)] = []byte(Filter4)

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

func BenchmarkStandardLoopExampleWithConcurrency(b *testing.B) {
	bagOfWords := GenerateBagOfWords()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StandardLoopExampleWithConcurrency(bagOfWords)
	}
}

func BenchmarkStandardLoopExampleWithoutConcurrency(b *testing.B) {
	bagOfWords := GenerateBagOfWords()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StandardLoopExampleWithoutConcurrency(bagOfWords)
	}
}
