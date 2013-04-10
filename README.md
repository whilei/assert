# Assert

Simple test assertions for Go. This is a fork of [bmizerany/assert][original]
with improved support for things like nil pointers, etc.

[Documentation][docs]

[original]: http://github.com/bmizerany/assert
[docs]: http://godoc.org/github.com/stvp/assert

Installation
------------

    $ go get github.com/stvp/assert

Example
-------

```go
package main

import "github.com/stvp/assert"

type CoolStruct struct{}

func TestThings(t *testing.T) {
  myString := "cool"

  assert.Equal(t, "cool", myString, "myString should be equal")
  assert.NotEqual(t, "nope", myString)

  var myStruct CoolStruct
  assert.Nil(t, myStruct)
}
```

See [assert_test.go][assert_test] for more usage examples.

Output
------

You can add extra information to test failures by passing in any number of extra
arguments:

```go
assert.Equal(t, "foo", myString, "Should set up a proper foo string")
```

```console
% go test
--- FAIL: TestImportantFeature (0.00 seconds)
	assert.go:18: /Users/tyson/go/src/github.com/foo/bar/main_test.go:31
	assert.go:38: 💩  Unexpected: "foo"
	assert.go:40: 💩  - Should set up a proper foo string
FAIL
exit status 1
FAIL	github.com/foo/bar	0.017s
```

[assert_test]: https://github.com/stvp/assert/blob/master/assert_test.go

License
-------

Copyright Blake Mizerany and Keith Rarick. Licensed under the [MIT
license](http://opensource.org/licenses/MIT). Additional modifications by
Stovepipe Studios.

