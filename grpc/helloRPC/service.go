package helloRPC

type helloService interface {
	Hello(request string, reply *string) error
}
