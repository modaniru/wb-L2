package pattern

import "fmt"

//Некий клиент, который мы не хотим расширять
type client interface{
	accept(visitor Visitor)
	getSenderEmail() string
}

//Его 'абстрактная' реализация, которая возвращает панику при реализованных ее методах
type BaseClient struct{
	senderEmail string
}

func (b *BaseClient) getSenderEmail() string {
	return b.senderEmail
}

func (b *BaseClient) accept(visitor Visitor) {
	panic("not implemented") // TODO: Implement
}

//Реализации, которые агрегируют BaseClient
//if not overlapping methond send email, throw panic if it execute
//has method getSenderEmail
type TaxiClient struct{
	BaseClient
}

func (t *TaxiClient) accept(visitor Visitor) {
	visitor.sendTaxiEmail(*t)
}

type RestaurantClient struct{
	BaseClient
}

func (r *RestaurantClient) accept(visitor Visitor) {
	visitor.sendRestaurantEmail(*r)
}

//Функции, которые мы хотим добавить к этой структуре. Тем самым только реализации visior'a открыты для модификации. Они могут быть и закрыты, ведь мы можем создать новую реализацию визитора.
type Visitor interface{
	sendAll(clients []client)
	sendTaxiEmail(client TaxiClient)
	sendRestaurantEmail(client RestaurantClient)
}

//Реализация посетителя
 type MyVisitor struct{

 }
//Проходим по всем клиентам, чтобы вызвать у них всех каждый своей метод, который использует посетитель
 func (m *MyVisitor) sendAll(clients []client){
	for _, v := range clients{
		v.accept(m)
	}
 }

 func (m *MyVisitor) sendTaxiEmail(client TaxiClient) {
	fmt.Println("send taxi email")
}

func (m *MyVisitor) sendRestaurantEmail(client RestaurantClient) {
	fmt.Println("send restaurant email")
}

func main(){
	//Кладем объект структуру в переменную интерфейса
	var visitor Visitor = &MyVisitor{}
	//Массив клиентов
	var clients []client = []client{&RestaurantClient{}, &TaxiClient{}}
	//Выполнить что-то у визитора, в контексте клиента
	clients[0].accept(visitor)
	//Послать всем клиентам сообщение
	visitor.sendAll(clients)

	/*
	
	В данном примере произойдет следующее sendAll()

	visitor.sendAll(clients) -> for client.accept(visitor Visitor) -> visiot.(Function on struct)

	*/
}

