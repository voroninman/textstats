# Text statistics service

My attempt to do networking programming with Go.

### Build
```bash
go build textstats.go stat*.go
```

### Run
```bash
$ textstats --port=5555 --stat=8080 --debug=false
2015/05/19 15:30:26 Stats at http://localhost:8080/stats
2015/05/19 15:30:26 Listening TCP localhost:5555
```

### Usage
```bash
$ # Sending arbitrary text (supports Unicode)
$ # The service will skip all non-letter characters (⌘ will be skipped)
$ echo 'ΩΩΩ ⌘⌘⌘ 本語日' | nc localhost 5555
$
$ # Dump the statistics (default N=5)
$ # All input text is transformed to lowercase (Ω -> ω)
$ curl http://localhost:8080/stats?N=2
{"count":2,"top_2_letters":["ω","本"],"top_2_words":["ωωω","本語日"]}
```

### TODO

- Parse an input `data\n` interactively
- Add a simple functional test
- Add the functional test that checks against concurrency race conditions (maps aren't concurrency safe)
