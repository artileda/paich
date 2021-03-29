## Paich

Simple code generator using Go's text/template library. Field descriptor inspired by Ruby on Rails.

Simple Usage
  - `./paich [model name] [field name]:[field type]:[field attr] --template=[template name]`
  - `./paich generate test test:test --template=hello-world`

Usage using `go run`
  - `go run main.go generate test name:Pavolia' 'Reine age:Expired  --template=json`
  -   Output
```
{ 
    "name":"Pavolia Reine",
    "age":"Expired", 
}
```

Folder template contain template following `text/template` format, and filename will be considered as template name.

Create your template and place it in `template` folder or assign environment variable `PAICH_TEMPLATE_DIR` with you template path then place you template in it.

```
PAICH_TEMPLATE_DIR="/home/user/template" ./paich 
```

### How to build

- `go build`
- `./paich`

#### Resource

- [Using Go Template, Gopher Academy](https://blog.gopheracademy.com/advent-2017/using-go-templates/)
- [Golang Templates Cheatsheet, Curtis Vermeeren](https://curtisvermeeren.github.io/2017/09/14/Golang-Templates-Cheatsheet)

### License

[MIT](./LICENSE)