package main

import (
	"fmt"
)

type StringList []string

func (l StringList) Map(fn MapFunc) Query {
	query := NewQuery(l)

	query = query.Map(fn)

	return query
}

func (l StringList) Filter(fn FilterFunc) Query {
	query := NewQuery(l)

	query = query.Filter(fn)

	return query
}

func (l StringList) Reduce(init string, fn ReduceFunc) string {
	query := NewQuery(l)

	return query.Reduce(init, fn)
}

type Operation struct {
	Map       MapFunc
	Filter    FilterFunc
	Reduce    ReduceFunc
	Operation OperationType
}

type OperationType string

type MapFunc func (string) string
type FilterFunc func (string) bool
type ReduceFunc func (string, string) string

const (
	Map OperationType = "Map"
	Filter OperationType = "Filter"
	Reduce OperationType = "Reduce"
)

type Query struct {
	list StringList
	operations []*Operation
}

func NewQuery(list StringList) Query {
	query := Query{
		list: list,
		operations: []*Operation{},
	}

	return query
}

func (l Query) Map(fn MapFunc) Query {
	op := &Operation{}
	op.Map = fn
	op.Operation = Map

	l.operations = append(l.operations, op)

	return l
}

func (l Query) Filter(fn FilterFunc) Query {
	op := &Operation{}
	op.Filter = fn
	op.Operation = Filter

	l.operations = append(l.operations, op)

	return l
}

func (l Query) Reduce(init string, fn ReduceFunc) string {
	reduction := init

	for _, elem := range l.list {
		value := elem
		isIncluded := true

		for _, op := range l.operations {
			switch (op.Operation) {
			case Map:
				value = op.Map(value)
			case Filter:
				isIncluded = isIncluded && op.Filter(elem)

				if !isIncluded {
					break
				}
			}
		}

		if isIncluded {
			reduction = fn(reduction, value)
		}
	}

	return reduction
}

/*func (l Query) ExecOne(op func (string)) {

}*/

func (l Query) ExecAll() StringList {
	newStringList := StringList{}

	for _, elem := range l.list {
		value := elem
		isIncluded := true

		for _, op := range l.operations {
			switch (op.Operation) {
			case Map:
				value = op.Map(value)
			case Filter:
				isIncluded = isIncluded && op.Filter(elem)

				if !isIncluded {
					break
				}
			}
		}

		if isIncluded {
			newStringList = append(newStringList, value)
		}
	}

	return newStringList
}

func main() {
	bagOfWords := StringList{ "This", "is", "a", "bag", "of", "words." }
	query1 := bagOfWords.Map(func (word string) string {
		return word + " "
	})

	query2 := query1.Filter(func (word string) bool {
		return word != "is"
	})

	result1 := query1.Reduce("", func (init string, word string) string {
		return init + word
	})

	result2 := query2.Reduce("", func (init string, word string) string {
		return init + word
	})

	fmt.Println(result1)
	fmt.Println(result2)
}
