## Course
https://www.udemy.com/course/go-programming-golang-the-complete-developers-guide

## Packages
using packages:

```go
import "name"

// or import multiple packeges
import (
  "name"
  "namespace/packageName"
)

// import . "name" - all of methods from the package "name", destructured
// import pk - alias, referring to the long name "namespace/packageName"
import (
  . "name"
  pk "namespace/packageName"
)
```

## Modules
- Created by having `go.mod` file in the root directory
- `go.mod` can be managed by GO CLI
- `go.mod` contains information ablut your project: dependencies, go version, package info
- all Go projects have `go.mod` file

example of the `go.mod` file:

1. namespace: example.com, module name: practice
2. go version 1.19
3. require block, multiple packages, each line: external package name and version, ( for example: namespace: github.com/alexflint, packageName: go-arg, version: v1.4.2 )

```mod
module example.com/practice

go 1.19

require (
  github.com/alexflint/go-arg v1.4.2
  github.com/fatih/color v1.13.0
)
```
