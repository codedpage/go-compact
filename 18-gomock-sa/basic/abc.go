package abc

type ABCInterface interface {
	DoSomething(x int) string
	DoSomething2(x string) (string, error)
}
