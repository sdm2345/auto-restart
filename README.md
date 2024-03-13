# auto-restart

This is a simple script to restart a program when it crashes. It's useful for
long-running programs that you want to keep running even if they crash.

## build

```shell
git clone https://github.com/sdm2345
make install
```

## Usage

```shell
$ ./auto-restart 
Error: requires at least 1 arg(s), only received 0
Usage:
  auto-restart [flags]

Examples:
auto-restart -- curl https://google.com

Flags:
  -d, --debug            debug flag
  -t, --delay-time int   auto restart delay time (default 3)
  -f, --force            Force restart regardless of exit code 0
  -h, --help             help for auto-restart
  -m, --max-count int    max retry count, 0 means no limit

2024/03/13 18:17:05 requires at least 1 arg(s), only received 0

```

```shell
/03/13 18:17:55 auto-start will to start process with delay:3,try max count:0, args:[curl error]
2024/03/13 18:17:55 start process with count:1,try max count:0
curl: (6) Could not resolve host: error
2024/03/13 18:17:58 wait process get error: curl
2024/03/13 18:17:58 start process get error:exit status 6,last run time:3.72s
2024/03/13 18:18:01 start process with count:2,try max count:0
curl: (6) Could not resolve host: error
2024/03/13 18:18:01 wait process get error: curl
2024/03/13 18:18:01 start process get error:exit status 6,last run time:0.02s

```
