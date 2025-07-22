package abc

import "fmt"

type ABCImpl struct{}

func (a *ABCImpl) DoSomething(x int) string {
	return fmt.Sprintf("Received %d", x)
}

func (a *ABCImpl) DoSomething2(x string) (string, error) {
	if x == "" {
		return "", fmt.Errorf("empty input")
	}
	return "Hello, " + x, nil
}
