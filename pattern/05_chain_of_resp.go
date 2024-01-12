package pattern

import "fmt"

type Chain interface {
	next()
}

type BusinessLoginChain struct{}

func (c *BusinessLoginChain) next() {
	fmt.Println("business final chain")
}

type AuthentificationChain struct{
	chain Chain
}

func (c *AuthentificationChain) next() {
	c.chain.next()
}

type AuthorisationChain struct{
	chain Chain
}

func (c *AuthorisationChain) next() {
	c.chain.next()
}

/*
Этот паттерн позволяет передавать выполнение операции наследующее звено функции. В звеньях могут быть промежуточная логика.
Допустим эта цепочка может состоять из всех трех звеньев, однако цепочка бизнес логики будент последней. Ее можно сделать такой Авторизаци->Аутентификая->Бизнес
*/