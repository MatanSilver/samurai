ls -d * | entr find . -name "*.go" | xargs -n 1 gofmt -w
