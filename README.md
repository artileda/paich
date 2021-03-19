## Paich

Simple Go code generator for web api. Currently underdevelopment. Field descriptor inspired by Ruby on Rails.

Simple Usage
  - `./paich [model name] [field name]:[field type]:[field attr] --template=[template name]`
  - `./paich generate test test:test --template=hello-world.paich`

Folder template contain template following `text/template` format, and filename will be considered as template name.

### How to build

- `go build`
- `./paich`


### License

[MIT](./LICENSE)