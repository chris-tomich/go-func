package generated_example

type StringListMapFunc func (string) string
type StringListFilterFunc func (string) bool
type StringListReduceFunc func (string, string) string

type StringList []string

func (l StringList) Map(fn StringListMapFunc) StringListQuery {
	query := NewStringListQuery(l)

	query = query.Map(fn)

	return query
}

func (l StringList) Filter(fn StringListFilterFunc) StringListQuery {
	query := NewStringListQuery(l)

	query = query.Filter(fn)

	return query
}

func (l StringList) Reduce(init string, fn StringListReduceFunc) string {
	query := NewStringListQuery(l)

	return query.Reduce(init, fn)
}

type StringListOperation struct {
	Map       StringListMapFunc
	Filter    StringListFilterFunc
	Reduce    StringListReduceFunc
	Operation string
}

type StringListQuery struct {
	list StringList
	operations []*StringListOperation
}

func NewStringListQuery(list StringList) StringListQuery {
	query := StringListQuery{
		list: list,
		operations: []*StringListOperation{},
	}

	return query
}

func (l StringListQuery) Map(fn StringListMapFunc) StringListQuery {
	op := &StringListOperation{}
	op.Map = fn
	op.Operation = "Map"

	l.operations = append(l.operations, op)

	return l
}

func (l StringListQuery) Filter(fn StringListFilterFunc) StringListQuery {
	op := &StringListOperation{}
	op.Filter = fn
	op.Operation = "Filter"

	l.operations = append(l.operations, op)

	return l
}

func (l StringListQuery) Reduce(init string, fn StringListReduceFunc) string {
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

func (l StringListQuery) ExecAll() StringList {
	newStringList := StringList{}

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
			newStringList = append(newStringList, value)
		}
	}

	return newStringList
}