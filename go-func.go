package main

import "fmt"

type StringList []string

func (l StringList) Map(op MapFunc) Query {
	query := NewQuery(l)

	query = query.Map(op)

	return *query
}

func (l StringList) Filter(op FilterFunc) Query {
	query := NewQuery(l)

	query = query.Filter(op)

	return *query
}

func (l StringList) Reduce() Query {
	query := NewQuery(l)

	return *query
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

func NewQuery(list StringList) *Query {
	query := &Query{
		list: list,
		operations: []*Operation{},
	}

	return query
}

func (l Query) Map(fn MapFunc) *Query {
	op := &Operation{}
	op.Map = fn
	op.Operation = Map

	l.operations = append(l.operations, op)

	return &l
}

func (l Query) Filter(fn FilterFunc) *Query {
	op := &Operation{}
	op.Filter = fn
	op.Operation = Filter

	l.operations = append(l.operations, op)

	return &l
}

/*func (l Query) Reduce() string {

}*/

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
	query := bagOfWords.Map(func (word string) string {
		return word + " "
	})
	stringList := query.ExecAll()

	for _, word := range stringList {
		fmt.Print(word)
	}

	fmt.Println("")
}
