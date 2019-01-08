# pprof
```bash
$ go build main.go
$ ./main

$ curl -s http://localhost:6060/debug/pprof/profile > cpu.pprof
$ go tool pprof -png main cpu.pprof > out.png
```
