#crc32

A collection of crc32 algorithms for Go language.

Algorithms implemented:
* bitwise
* half-byte
* standard (1 byte)
* standard (no lookup table)
* slicing-by-4
* slicing-by-8
* slicing-by-16

Note, that it isn't a replacement for native Go's crc32 algorithms, since they are faster.

See benchmarks:
```
BenchmarkCrc32NativeSimple-8         200000      8397 ns/op    1024 B/op       1 allocs/op
BenchmarkCrc32NativeIEEE-8         20000000       102 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32NativeCastagnoli-8   20000000        65.1 ns/op     0 B/op       0 allocs/op
BenchmarkCrc32Bitwise-8              200000      8229 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32Halfbyte-8             500000      3827 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32Byte1-8               1000000      1665 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32Byte1Tableless-8       500000      3626 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32Bytes4-8              2000000       749 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32Bytes8-8              3000000       448 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32Bytes4x8-8            2000000       620 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32Bytes16-8             3000000       525 ns/op       0 B/op       0 allocs/op
```
