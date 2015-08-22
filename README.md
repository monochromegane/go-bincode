# go-bincode

A tool for embedding Go code in a Go binary using go-bindata.

## Usage

1. Install go-bindata and go-bincode.
2. Execute `go-bincode` command at your repository.

```sh
$ go-bincode
```

You can find `bincode.go` and `bincoder.go`.
So, your Go program get {show,list,restore}-code options.

### List code

`list-code` list your code.

```sh
$ my-app -list-code
app.go
my-app/main.go
```
**Note** code-list don't contain `bincode.go` and `bincoder.go`.

### Show code

`show-code` show your code.

```sh
$ my-app -show-code my-app/main.go
package main

func main() {
    my-app.Do()
}
```

### Restore code

`restore-code` restore your code.

```sh
$ my-app -restore-code
```

## Generated code dependency

Generated code (`bincode.go` and `bincoder.go`) don't depends on your code, only provides {show,list,restore}-code options at init() function.

## Installation

```sh
$ go get -u github.com/jteeuwen/go-bindata/... # Requirement
$ go get -u github.com/monochromegane/go-bincode/...
```

## Contribution

1. Fork it
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Run test suite with the `go generate && go test ./...` command and confirm that it passes
6. Run `gofmt -s`
7. Create new Pull Request

## License

[MIT](https://github.com/monochromegane/go-bincode/blob/master/LICENSE)

## Author

[monochromegane](https://github.com/monochromegane)

