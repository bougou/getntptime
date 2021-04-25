# A tiny tool to get the time from a ntp server

## Compile

```bash
GOOS=linux GOARCH=amd64 go build -v -x
```

## Run

```bash
$ ./getntptime --help
Usage of ./getntptime:
      --format string   output format, one of: unixnsec|unixmsec|unixsec|datetime (default "datetime")
      --server string   ntpd server (default "0.pool.ntp.org")

$ ./getntptime
2021-04-26 00:38:29.560562583 +0800 CST m=+0.309613078

$ ./getntptime --server 1.pool.ntp.org --format unixsec
1619368840

$ ./getntptime --server 1.pool.ntp.org --format unixmsec
1619368843191

$ ./getntptime --server 1.pool.ntp.org --format unixnsec
1619368845138040574
```
