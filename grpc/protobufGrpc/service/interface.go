package service

const HelloServiceName = "HelloService"

type HelloService interface {
	Hello(*Request, *Reply) error
}
