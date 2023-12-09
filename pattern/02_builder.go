package pattern

import "fmt"

//Интерфейс чего мы будем билдить

type Storage interface{
	PrintInformation()
}

//Реализация этого интерфейса

type PostgresStorage struct{
	driver string
	host string
	port string
	login string
	password string
}

func (p *PostgresStorage) PrintInformation(){
	fmt.Printf("driver: %s, host: %s, port: %s, login: %s, password: %s\n", p.driver, p.host, p.port, p.login, p.password)
}

//Лучше всего реализовать конструктор, чтобы внутренние переменные были закрыты
func NewPostgresStorage(driver, host, port, login, password string) *PostgresStorage{
	//тут может быть валидация входнх данных
	return &PostgresStorage{driver: driver, host: host, port: port, login: login, password: password}
}

//Сам сборщик, который будет постепенно собирать наше хранилище

type StorageBuilder interface{
	//Возвращаем инстанс сборщика, чтобы можно было сделать сборку более красивой
	SetDriver(driver string) StorageBuilder
	SetHost(host string) StorageBuilder
	SetPort(port string) StorageBuilder
	SetLogin(login string) StorageBuilder
	SetPassword(password string) StorageBuilder
	//Функция сборки, которая будет возвращать нам какую-то реализацию хранилища
	Build() (Storage, error)
}

//Реализация сборщика
type PostgresStorageBuilder struct{
	driver string
	host string
	port string
	login string
	password string
}

func (p *PostgresStorageBuilder) SetDriver(driver string) StorageBuilder{
	//Если захочется делать валидацию данных непосредтвенно в сборщике, то я бы завел поле с массивом ошибок.
	//При возникновении ошибки, добавлял бы ее в массив. Если в массиве есть ошибки, возращал бы их.
	p.driver = driver
	return p
}

func (p *PostgresStorageBuilder) SetHost(host string) StorageBuilder {
	p.host = host
	return p
}

func (p *PostgresStorageBuilder) SetPort(port string) StorageBuilder {
	p.port = port
	return p
}

func (p *PostgresStorageBuilder) SetLogin(login string) StorageBuilder{
	p.login = login
	return p
}

func (p *PostgresStorageBuilder) SetPassword(password string) StorageBuilder{
	p.password = password
	return p
}

//Функция сборщика
func (p *PostgresStorageBuilder) Build() (Storage, error) {
	return NewPostgresStorage(p.driver, p.host, p.port, p.login, p.password), nil
}

/* client function */
func getStorage(builder StorageBuilder) Storage{
	st, _ := builder.SetDriver("postgres").
					SetHost("localhost").
					SetLogin("postgres").
					SetPassword("postgres").
					SetPort("5432").Build()
	return  st
}

/*
	Builder позволяет избавиться от огромных методов-конструкторов, позволяя создавать экземпляр структуры по этапам. При этом у сброщика есть свой интерфейс, что позволяет не зависить от конкретной реализации сборщика.
	Минус. Для каждой конкретной реализации 'хранилища' нужно будет создавать собственный сборщик для этого 'хранилища', что будет увеличивать сложность программы.
*/
func main(){
	storage := getStorage(&PostgresStorageBuilder{})
	storage.PrintInformation()
}