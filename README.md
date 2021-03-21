## Paich

Simple code generator using Go's text/template library. Field descriptor inspired by Ruby on Rails.

Simple Usage
  - `./paich [model name] [field name]:[field type]:[field attr] --template=[template name]`
  - `./paich generate test test:test --template=hello-world.paich`

Usage using `go run`
  - `go run main.go generate test name:Pavolia' 'Reine age:Expired  --template=json.paich`
  -   Output
```
{ 
    "name":"Pavolia Reine",
    "age":"Expired", 
}
```

Folder template contain template following `text/template` format, and filename will be considered as template name.

### How to build

- `go build`
- `./paich`

#### Resource

- [Using Go Template, Gopher Academy](https://blog.gopheracademy.com/advent-2017/using-go-templates/)
- [Golang Templates Cheatsheet, Curtis Vermeeren](https://curtisvermeeren.github.io/2017/09/14/Golang-Templates-Cheatsheet)

### License

[MIT](./LICENSE)