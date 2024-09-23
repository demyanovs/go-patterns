package main

type customError struct {
}

func (e customError) Error() string {
	return "my error"
}

func main() {
	println(handle().Error())
}

func handle() error {
	return customError{}
}
