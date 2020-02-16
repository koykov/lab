# crc32

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
BenchmarkCrc32NativeSimple-8             200000      7455 ns/op    1024 B/op       1 allocs/op
BenchmarkCrc32NativeSimpleLong-8           1000   1948223 ns/op    1024 B/op       1 allocs/op
BenchmarkCrc32NativeIEEE-8             20000000       101 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32NativeIEEELong-8            30000     43664 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32NativeCastagnoli-8       20000000      65.2 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32NativeCastagnoliLong-8      50000     34471 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32Bitwise-8                  200000      8289 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32BitwiseLong-8                 200   8288214 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32Halfbyte-8                 500000      3837 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32HalfbyteLong-8                500   3862391 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32Byte1-8                   1000000      1663 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32Byte1Long-8                  1000   1691251 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32Byte1Tableless-8           500000      3632 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32Byte1TablelessLong-8          500   3634322 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32Bytes4-8                  2000000       761 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32Bytes4Long-8                 2000    760386 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32Bytes8-8                  3000000       460 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32Bytes8Long-8                 3000    443565 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32Bytes4x8-8                2000000       604 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32Bytes4x8Long-8               3000    576306 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32Bytes16-8                 3000000       510 ns/op       0 B/op       0 allocs/op
BenchmarkCrc32Bytes16Long-8                3000    480952 ns/op       0 B/op       0 allocs/op
```
