# go-func
A tool intended to be used with go generate to generate various functional features for custom types.

For more information about it and it's use, go to https://mymemorysucks.wordpress.com/2016/09/30/go-func-yourself-making-golang-a-little-more-functional-part-1/ .

To build go func, at a command-line use the following go build command.

```
go build -o $GOPATH/bin/go-func $GOPATH/src/github.com/chris-tomich/go-func/cmd/go_func.go
```

For an example of the go generate argument to use, look at "string_list_example/string_list.go".