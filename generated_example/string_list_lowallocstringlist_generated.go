package generated_example

type LowAllocStringListMapFunc func ([]byte) []byte
type LowAllocStringListFilterFunc func ([]byte) bool
type LowAllocStringListReduceFunc func ([]byte, []byte) []byte

type LowAllocStringList [][]byte

func (l LowAllocStringList) Map(fn LowAllocStringListMapFunc) LowAllocStringListQuery {
	query := NewLowAllocStringListQuery(l)

	query = query.Map(fn)

	return query
}

func (l LowAllocStringList) Filter(fn LowAllocStringListFilterFunc) LowAllocStringListQuery {
	query := NewLowAllocStringListQuery(l)

	query = query.Filter(fn)

	return query
}

func (l LowAllocStringList) Reduce(init []byte, fn LowAllocStringListReduceFunc) []byte {
	query := NewLowAllocStringListQuery(l)

	return query.Reduce(init, fn)
}

type LowAllocStringListOperation struct {
	Map       LowAllocStringListMapFunc
	Filter    LowAllocStringListFilterFunc
	Reduce    LowAllocStringListReduceFunc
	Operation string
}

type LowAllocStringListQuery struct {
	list LowAllocStringList
	operations []*LowAllocStringListOperation
}

func NewLowAllocStringListQuery(list LowAllocStringList) LowAllocStringListQuery {
	query := LowAllocStringListQuery{
		list: list,
		operations: []*LowAllocStringListOperation{},
	}

	return query
}

func (l LowAllocStringListQuery) Map(fn LowAllocStringListMapFunc) LowAllocStringListQuery {
	op := &LowAllocStringListOperation{}
	op.Map = fn
	op.Operation = "Map"

	l.operations = append(l.operations, op)

	return l
}

func (l LowAllocStringListQuery) Filter(fn LowAllocStringListFilterFunc) LowAllocStringListQuery {
	op := &LowAllocStringListOperation{}
	op.Filter = fn
	op.Operation = "Filter"

	l.operations = append(l.operations, op)

	return l
}

func (l LowAllocStringListQuery) Reduce(init []byte, fn LowAllocStringListReduceFunc) []byte {
	reduction := init

	for _, elem := range l.list {
		value := elem
		isIncluded := true

		for _, op := range l.operations {
			switch (op.Operation) {
			case "Map":
				value = op.Map(value)
			case "Filter":
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

func (l LowAllocStringListQuery) ExecAll() LowAllocStringList {
	newLowAllocStringList := LowAllocStringList{}

	for _, elem := range l.list {
		value := elem
		isIncluded := true

		for _, op := range l.operations {
			switch (op.Operation) {
			case "Map":
				value = op.Map(value)
			case "Filter":
				isIncluded = isIncluded && op.Filter(elem)

				if !isIncluded {
					break
				}
			}
		}

		if isIncluded {
			newLowAllocStringList = append(newLowAllocStringList, value)
		}
	}

	return newLowAllocStringList
}