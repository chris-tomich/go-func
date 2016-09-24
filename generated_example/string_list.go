package generated_example

//go:generate /home/chris/go-workspace/bin/go_func -singular=string -collection=StringList
type StringList []string

//go:generate /home/chris/go-workspace/bin/go_func -singular=[]byte -collection=LowAllocStringList
type LowAllocStringList [][]byte
