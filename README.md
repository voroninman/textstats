# Word statistics service

My attempt to do networking programming with Go.

### Build
```
go build textstats.go stat.go stat_letters.go stat_words.go
```

### Run
```
$ textstats --port=5555 --stat=8080
2015/05/17 21:04:33 Waiting for data at http://localhost:5555
2015/05/17 21:04:33 Stats at http://localhost:8080
```

### Usage
```
$ # Sending arbitrary text
$ curl -X POST -d 'bar foo buzz foo buzz bar bar bar' http://localhost:5555
$ # Dump the statistics (default N=5)
$ curl http://localhost:8080/stats?N=3
{"count":8,"top_3_letters":["b","a","r"],"top_3_words":["bar","buzz","foo"]}
```

### TODO

- Accept input data over TCP not over HTTP
- Read a request body interactively
