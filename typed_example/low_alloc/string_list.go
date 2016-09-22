package low_alloc

type OperationType string

const (
	Map OperationType = "Map"
	Filter OperationType = "Filter"
	Reduce OperationType = "Reduce"
)

type MapFunc func ([]byte) []byte
type FilterFunc func ([]byte) bool
type ReduceFunc func ([]byte, []byte) []byte

type StringList [][]byte

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

func (l StringList) Reduce(init []byte, fn ReduceFunc) []byte {
	query := NewQuery(l)

	return query.Reduce(init, fn)
}

type Operation struct {
	Map       MapFunc
	Filter    FilterFunc
	Reduce    ReduceFunc
	Operation OperationType
}

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

func (l Query) Reduce(init []byte, fn ReduceFunc) []byte {
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