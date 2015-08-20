# test go lang

Linux, OSX: [![Build Status](https://api.travis-ci.org/wookay/gocat.svg?branch=master)](https://travis-ci.org/wookay/gocat)
Windows: [![Build status](https://ci.appveyor.com/api/projects/status/epysdl4dds0pyhvn?svg=true)](https://ci.appveyor.com/project/wookay/gocat)

```shell
~/work/gocat master$ go test -v *.go
=== RUN   TestInt
--- PASS: TestInt (0.00s)
=== RUN   TestFloat
--- PASS: TestFloat (0.00s)
=== RUN   TestString
--- PASS: TestString (0.00s)
=== RUN   TestArray
--- PASS: TestArray (0.00s)
=== RUN   TestStruct
--- PASS: TestStruct (0.00s)
PASS
ok    command-line-arguments  0.013s
```
