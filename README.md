# ant

<p align="center">
<a href="https://github.com/ant-go/ant/actions"><img src="https://github.com/ant-go/ant/workflows/tests/badge.svg" alt="Build Status"></a>
<a href="https://codecov.io/gh/ant-go/ant"><img src="https://codecov.io/gh/ant-go/ant/branch/main/graph/badge.svg" alt="codecov"></a>
<a href="https://goreportcard.com/report/github.com/ant-go/ant"><img src="https://goreportcard.com/badge/github.com/ant-go/ant" alt="Go Report Card"></a>
</p>

[ant](https://github.com/ant-go/ant) is a framework that can help PHP developers quickly switch to the Go language.

**It’s small, but powerful.**

Commonly used PHP's built-in functions will be implemented first in the plan, some modules of the most popular framework [Laravel](https://laravel.com/) will also be implemented later.

## Install

```bash
go mod download github.com/ant-go/ant
```

## Requirements

Go 1.7+.

## PHP built-in functions

#### Array(Slice/Map) Functions

* in_array()
* array_keys()
* array_values()

#### Hash Functions

* md5() | md5_file()
* sha1() | sha1_file()
* crc32()
* hash_equals()
* hash()
    * sha224
    * sha384
    * sha512/244
    * sha512/256
    * sha512
    * sha3-224
    * sha3-256
    * sha3-384
    * sha3-512
    * fnv132
    * fnv1a32
    * fnv164
    * fnv1a64

## LICENSE

MIT Licence.