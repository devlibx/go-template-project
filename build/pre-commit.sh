# This is a file which will do everything that is needed to be done before a commit
# This is a very extensive script to make sure everything is good
go install go.uber.org/mock/mockgen@latest
export E2E_TESTS_ENABLED=true
export INTEGRATION_TESTS_ENABLED=true
git pull
find . -type f -name '*.go-e' -exec rm {} \;
find . -type f -name '*_mock.go' -exec rm {} \;
go generate ./...
go fmt ./...
go test -count=1 ./...