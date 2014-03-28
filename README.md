Hacking on Eli Bendersky's [TableGen lexer in Go](http://eli.thegreenplace.net/2014/03/27/rewriting-the-lexer-benchmark-in-go/)

Forked from https://gist.github.com/eliben/9740414

Run benchmark: `go test -bench=.`

Generate CPU profiler data: `go test -bench=. --cpuprofile=out.prof`

Render profiler data: `go tool pprof --svg tablegen-lexer.test out.prof > cpu.svg`
