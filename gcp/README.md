Usage
==========================

# prepare

- create default credentials via https://developers.google.com/identity/protocols/application-default-credentials
- download json key, for example `hyper-test-4be009a5b924.json`

# run

```
$ export http_proxy=http://127.0.0.1:8118
$ export https_proxy=http://127.0.0.1:8118

$ export GOOGLE_APPLICATION_CREDENTIALS=`pwd`/hyper-test-4be009a5b924.json
$ go run disk.go
```
