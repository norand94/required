## How to install

```shell
go get github.com/norand94/required   
```

## How to Use


```go
type T struct {
    Home     string `required:"t"`
    Optional string
}

var st = &T{"", ""}
err := Check(st)
fmt.Println(err) // required fields are empty: Home
```