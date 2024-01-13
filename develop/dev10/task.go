package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Парсим аргументы командной строки
	host := flag.String("host", "", "Target host (IP or domain name)")
	port := flag.Int("port", 0, "Target port")
	timeout := flag.Duration("timeout", 10*time.Second, "Connection timeout")
	flag.Parse()

	// Проверяем наличие обязательных аргументов
	if *host == "" || *port == 0 {
		fmt.Println("Usage: go-telnet --host <host> --port <port> [--timeout <timeout>]")
		os.Exit(1)
	}

	// Формируем адрес сервера
	serverAddr := fmt.Sprintf("%s:%d", *host, *port)

	// Устанавливаем таймаут на подключение
	conn, err := net.DialTimeout("tcp", serverAddr, *timeout)
	if err != nil {
		fmt.Printf("Error connecting to %s: %v\n", serverAddr, err)
		os.Exit(1)
	}
	defer conn.Close()

	// Канал для отслеживания завершения программы
	done := make(chan struct{})

	// Запускаем горутину для обработки ввода из STDIN
	go func() {
		io.Copy(conn, os.Stdin)
		done <- struct{}{}
	}()

	// Запускаем горутину для обработки вывода в STDOUT
	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()

	// Ожидаем завершения программы по сигналам или закрытию сокета
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	select {
	case <-done:
	case <-signals:
	}
}
