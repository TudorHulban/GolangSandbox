#### Install
```bash
go get github.com/securego/gosec/cmd/gosec
```
Copy from $GOPATH/bin to /usr/local/go/bin.

```bash
gosec -no-fail ./... 2>&1 | tee ~/ram/sast-gosec.log
```
