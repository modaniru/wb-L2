package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		
		if !scanner.Scan(){
			break
		}

		input := scanner.Text()
		commands := strings.Split(input, " | ")
		for _, command := range commands{
			execCommand(command)
		}
	}
}

func execCommand(command string){
	args := strings.Split(command, " ")
	switch args[0]{
	case "cs":
	case "ls":
	case "pwd":
	case "echo":
	case "kill":
	case "ps":
	case "quit":
		fmt.Println("Bye Bye!")
		os.Exit(0)
	default:
		fmt.Println("unexcepted command!")
		return
	}

	cmd := exec.Command(args[0], args[1:]...)
	err := cmd.Err
	if err != nil{
		log.Fatal(err.Error())
	}

	bytes, err := cmd.Output()
	if err != nil{
		log.Fatal(err.Error())
	}

	str := string(bytes)
	fmt.Println(str[:len(str) - 1])
}