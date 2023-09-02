package main

import (
	"example/user/hello/morestrings"
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func main() {
	// fmt.Println("Hello, world.")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}

// Running:
// $ go install example/user/hello
// will produce an executable binary until $HOME/go/bin/hello

// This path can be set and unset using:
// $ go env -w GOBIN=/somewhere/else/bin
// $ go env -u GOBIN

// An import path can describe how to obtain the package source code using a revision control system such as Git or Mercurial.
// The go tool uses this property to automatically fetch packages from remote repositories.
// For instance, to use of github.com/google/go-cmp/cmp in the program above.

// Now that you have a dependency on an external module,
// you need to download that module and record its version in your go.mod file.
// The go mod tidy command adds missing module requirements for imported packages and removes requirements on modules that aren't used anymore.

// Module dependencies are automatically downloaded to the pkg/mod subdirectory of the directory indicated by the GOPATH environment variable

// To remove all downloaded modules, you can pass the -modcache flag to go clean:
// $ go clean -modcache
