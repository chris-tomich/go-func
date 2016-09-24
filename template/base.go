package template

const (
	PackageToken = "%%PACKAGE%%"
	SingularTypeToken = "%%SINGULAR_TYPE%%"
	CollectionTypeToken = "%%COLLECTION_TYPE%%"
)

const BaseTemplate = `package %%PACKAGE%%

type %%COLLECTION_TYPE%%MapFunc func (%%SINGULAR_TYPE%%) %%SINGULAR_TYPE%%
type %%COLLECTION_TYPE%%FilterFunc func (%%SINGULAR_TYPE%%) bool
type %%COLLECTION_TYPE%%ReduceFunc func (%%SINGULAR_TYPE%%, %%SINGULAR_TYPE%%) %%SINGULAR_TYPE%%

type %%COLLECTION_TYPE%% []%%SINGULAR_TYPE%%

func (l %%COLLECTION_TYPE%%) Map(fn %%COLLECTION_TYPE%%MapFunc) %%COLLECTION_TYPE%%Query {
	query := New%%COLLECTION_TYPE%%Query(l)

	query = query.Map(fn)

	return query
}

func (l %%COLLECTION_TYPE%%) Filter(fn %%COLLECTION_TYPE%%FilterFunc) %%COLLECTION_TYPE%%Query {
	query := New%%COLLECTION_TYPE%%Query(l)

	query = query.Filter(fn)

	return query
}

func (l %%COLLECTION_TYPE%%) Reduce(init %%SINGULAR_TYPE%%, fn %%COLLECTION_TYPE%%ReduceFunc) %%SINGULAR_TYPE%% {
	query := New%%COLLECTION_TYPE%%Query(l)

	return query.Reduce(init, fn)
}

type %%COLLECTION_TYPE%%Operation struct {
	Map       %%COLLECTION_TYPE%%MapFunc
	Filter    %%COLLECTION_TYPE%%FilterFunc
	Reduce    %%COLLECTION_TYPE%%ReduceFunc
	Operation string
}

type %%COLLECTION_TYPE%%Query struct {
	list %%COLLECTION_TYPE%%
	operations []*%%COLLECTION_TYPE%%Operation
}

func New%%COLLECTION_TYPE%%Query(list %%COLLECTION_TYPE%%) %%COLLECTION_TYPE%%Query {
	query := %%COLLECTION_TYPE%%Query{
		list: list,
		operations: []*%%COLLECTION_TYPE%%Operation{},
	}

	return query
}

func (l %%COLLECTION_TYPE%%Query) Map(fn %%COLLECTION_TYPE%%MapFunc) %%COLLECTION_TYPE%%Query {
	op := &%%COLLECTION_TYPE%%Operation{}
	op.Map = fn
	op.Operation = "Map"

	l.operations = append(l.operations, op)

	return l
}

func (l %%COLLECTION_TYPE%%Query) Filter(fn %%COLLECTION_TYPE%%FilterFunc) %%COLLECTION_TYPE%%Query {
	op := &%%COLLECTION_TYPE%%Operation{}
	op.Filter = fn
	op.Operation = "Filter"

	l.operations = append(l.operations, op)

	return l
}

func (l %%COLLECTION_TYPE%%Query) Reduce(init %%SINGULAR_TYPE%%, fn %%COLLECTION_TYPE%%ReduceFunc) %%SINGULAR_TYPE%% {
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

func (l %%COLLECTION_TYPE%%Query) ExecAll() %%COLLECTION_TYPE%% {
	new%%COLLECTION_TYPE%% := %%COLLECTION_TYPE%%{}

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
			new%%COLLECTION_TYPE%% = append(new%%COLLECTION_TYPE%%, value)
		}
	}

	return new%%COLLECTION_TYPE%%
}`
