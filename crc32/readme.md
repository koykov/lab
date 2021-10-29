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
BenchmarkCrc32/native-8                   187476              5574 ns/op            1024 B/op          1 allocs/op
BenchmarkCrc32/nativeIEEE-8             13999579             84.11 ns/op            0 B/op          0 allocs/op
BenchmarkCrc32/castagnoli-8             19790028             60.06 ns/op            0 B/op          0 allocs/op
BenchmarkCrc32/bitwise-8                  142256              8430 ns/op            0 B/op          0 allocs/op
BenchmarkCrc32/halfbyte-8                 313130              3836 ns/op            0 B/op          0 allocs/op
BenchmarkCrc32/byte4-8                   1578111             760.0 ns/op            0 B/op          0 allocs/op
BenchmarkCrc32/byte8-8                   2711008             439.1 ns/op            0 B/op          0 allocs/op
BenchmarkCrc32/byte4x8-8                 2452111             482.3 ns/op            0 B/op          0 allocs/op
BenchmarkCrc32/byte16-8                  2758572             445.4 ns/op            0 B/op          0 allocs/op
BenchmarkCrc32Long/native-8                  615           1935923 ns/op         1024 B/op          1 allocs/op
BenchmarkCrc32Long/nativeIEEE-8            26360             45075 ns/op            0 B/op          0 allocs/op
BenchmarkCrc32Long/castagnoli-8            35221             33677 ns/op            0 B/op          0 allocs/op
BenchmarkCrc32Long/bitwise-8                 142           8380055 ns/op            0 B/op          0 allocs/op
BenchmarkCrc32Long/byte1-8                   708           1713631 ns/op            0 B/op          0 allocs/op
BenchmarkCrc32Long/byte1Tableless-8          330           3621188 ns/op            0 B/op          0 allocs/op
BenchmarkCrc32Long/byte4-8                  1558            765062 ns/op            0 B/op          0 allocs/op
BenchmarkCrc32Long/byte8-8                  2800            445550 ns/op            0 B/op          0 allocs/op
BenchmarkCrc32Long/byte4x8-8                2666            440704 ns/op            0 B/op          0 allocs/op
BenchmarkCrc32Long/byte16-8                 3096            411865 ns/op            0 B/op          0 allocs/op
```
