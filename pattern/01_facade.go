package pattern

import "fmt"

/*
	Интерфейсы
*/

type Storage interface{
	getUserId(/* тут могут быть какие-то важные переменные */)
}

type Client interface{
	getUserInformation(/* тут могут быть какие-то важные переменные */)
}

type EmailClient interface{
	sendToUserEmail(/* тут могут быть какие-то важные переменные */)
}

/* Фасад */

type UserService interface{
	getUserInfo(/* тут могут быть какие-то важные переменные */)
	/* еще какие-нибудь методы */
}

/*
	Реализации
*/

type MyStorage struct{}

func (m *MyStorage) getUserId(){
	fmt.Println("MyStorage getUserId()")
}

type MyClient struct{}

func (m *MyClient) getUserInformation(){
	fmt.Println("MyClient getUserInformation()")
}

type MyEmailClient struct{}

func (m *MyEmailClient) sendToUserEmail(){
	fmt.Println("MyEmailClient sendToUserEmail()")
}

/* Реализация фасада */

type MyUserService struct{
	/* создаю структуру, композирующую не конкретные реализации, а интерфейсы. Таким образом получается довольно гибкая структура */
	userClient Client
	emailClient EmailClient
	storage Storage
}

func (m *MyUserService) getUserInfo(){
	/* Именно в реализации MyUserService мы совершаем нужные нам бизнес логику, предоставляя внешне удобный интерфейс */
	fmt.Println("MyUserService getUserInfo()")
	m.storage.getUserId()
	m.emailClient.sendToUserEmail()
	m.userClient.getUserInformation()
}

/* Пример структуры клиента */
/* Писать интерфейс клиенту не стал, так как это в нашем случае не нужно */

type HttpClient struct{
	/* Не завишу от конкретной реализации */
	userService UserService
}

/* Предположим этот метот перехватывает GET запросы */
func (h *HttpClient) Get(){
	/* Клиенту все равно на внутреннюю реализацию фасада, у него есть лишь нужный ему интерфейс */
	h.userService.getUserInfo()
}

func main(){
	/* Только функция main будет зависеть от конкретных реализаций, чтобы распределить их */
	var userService UserService = &MyUserService{userClient: &MyClient{}, storage: &MyStorage{}, emailClient: &MyEmailClient{}}
	client := HttpClient{userService: userService}

	/* допустим тут запускается наш сервер */
	client.Get()
}