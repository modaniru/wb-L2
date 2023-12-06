package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	//получаем время от сервера 'ntp0.ntp-servers.net'
	ntpTime, err := ntp.Time("ntp0.ntp-servers.net")
	if err != nil {
		//записываем ошибку в stderr
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
	//выводим результаты
	fmt.Printf("local time: %s\n", time.Now())
	fmt.Printf("ntp time: %s\n", ntpTime)
}
