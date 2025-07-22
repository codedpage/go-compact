mockgen -source=abc.go -destination=mocks/mock_abc.go -package=mocks

go test -v -cover .

go test -coverprofile=coverage.out ../abc
go test -coverprofile=coverage.out .
go tool cover -html=coverage.out