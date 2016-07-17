テストデータ

```
$ sh makeTestFiles.sh
$ du -hd1 ./testdata
2.5G    ./testdata/dir0
2.5G    ./testdata/dir1
2.5G    ./testdata/dir10
2.5G    ./testdata/dir2
2.5G    ./testdata/dir3
2.5G    ./testdata/dir4
2.5G    ./testdata/dir5
2.5G    ./testdata/dir6
2.5G    ./testdata/dir7
2.5G    ./testdata/dir8
2.5G    ./testdata/dir9
 27G    ./testdata
```

```
$ go test -bench . -benchmem
BenchmarkCompress-4            1        460285076137 ns/op      75295706376 B/op           44162 allocs/op
BenchmarkCompressNew-4         1        4492852287 ns/op         1045384 B/op      11474 allocs/op
ok      _/Users/shinofara/work/compress 464.881s
```
