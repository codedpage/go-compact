go test -v               (no %)
go test -v -cover        (with %)

go test
go test -run .


go test -coverprofile=coverage.out
go tool cover -html=coverage.out 


go test -coverprofile=coverage.out ../abc-full 


go test -timeout 30s -run ^TestABCImpl_DoSomething2$ abc
go test -run ^TestABCImpl_DoSomething2$ 