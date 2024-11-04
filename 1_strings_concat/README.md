# 如何高效拼接字符串？

Stack Overflow上最热门的问题。
https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go

方案：
- 未知数目/长度
  - 直接用`+`
  - `bytes.Buffer`
  - `strings.Builder` 
- 已知数目/长度
  - `append`(不预先分配内存)
  - `append`(预先分配内存) 
  - `copy`(预先分配内存) 

实际结果
```shell
BenchmarkConcat-8                1000000             21705 ns/op          503995 B/op          1 allocs/op
BenchmarkBuffer-8               295922799                4.094 ns/op           3 B/op          0 allocs/op
BenchmarkStringBuilder-8        491725101                2.428 ns/op           5 B/op          0 allocs/op
BenchmarkAppendEmpty-8          1000000000               1.190 ns/op           6 B/op          0 allocs/op
BenchmarkAppendPreAlloc-8       1000000000               0.4369 ns/op          0 B/op          0 allocs/op
BenchmarkCopy-8                 420070317                2.899 ns/op           0 B/op          0 allocs/op
```

结论：
不知道长度用`strings.Builder`，知道长度用`append`。