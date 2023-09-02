# 02 Simple Example


[Video](https://www.youtube.com/watch?v=-EYNVEv-snE&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=3)

- Run tests with `go test`.

- `GOPATH` is vestigial
    Instead we use go modules, to version and manage dependencies.

- `go.mod` contains all your dependencies for the module.
    Will update automatically when you pull in packages.
    See also, `go tidy`. 