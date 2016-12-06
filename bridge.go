package bridge;

func SayHello(name string) (greeting string, err error) {
	greeting = "Hello, " + name
	return
}