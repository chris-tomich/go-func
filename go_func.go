package main

import "flag"

func main() {
	collectionType := flag.String("type", "[]interface{}", "This is the name of the type to extend with map/filter/reduce functionality.")

	flag.Parse()
}