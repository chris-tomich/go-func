package main

import (
	"flag"
	"os"
	"fmt"
	"strings"
	"github.com/chris-tomich/go-func/template"
)

func main() {
	//TODO: Develop a version that looks at the line of the go generate argument and, if no type is given, look at the following line in code.
	singularType := flag.String("singular", "[]interface{}", "This is the name of the type contained within the collection.")
	collectionType := flag.String("collection", "[]interface{}", "This is the name of the type to extend with map/filter/reduce functionality.")

	flag.Parse()

	// The package name will be the package declared at the top of this file. For example, in this file it would be "main".
	// It does not include the full package location.
	packageName := os.Getenv("GOPACKAGE")
	fileName := os.Getenv("GOFILE") + "_" + *collectionType + "_generated.go"

	replacer := strings.NewReplacer(template.PackageToken, packageName, template.SingularTypeToken, *singularType, template.CollectionTypeToken, *collectionType)

	fd, fileOpenErr := os.Create(fileName)

	if fileOpenErr != nil {
		panic(fileOpenErr)
	}

	n, fileWriteErr := replacer.WriteString(fd, template.BaseTemplate)

	if fileWriteErr != nil {
		panic(fileWriteErr)
	}

	fmt.Printf("%d total replacements made.\n", n)
}