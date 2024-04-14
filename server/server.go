package main

import (
	"log"
	"math"
	"net"
	"net/http"
	"net/rpc"
)

const addr = ":12345"
const network = "tcp4"

// Тип данных для RPC сервера
type Server int

type Point struct {
	X, Y float64
}

type Points struct {
	A, B Point
}

// Метод Dist доступен для удаленного вызова
func (s *Server) Dist(msg Points, resp *float64) error {
	*resp = math.Sqrt(math.Pow((msg.A.X-msg.B.X), 2) + math.Pow((msg.A.Y-msg.B.Y), 2))
	return nil
}

func main() {
	// Создаем указатель на переменную типа Server
	srv := new(Server)
	// Регистрируем методы типа Server в службе RPC
	rpc.Register(srv)
	// Регистрируем HTTP-обработчик для службы RPC
	rpc.HandleHTTP()
	// Создаём сетевую службу
	listener, err := net.Listen(network, addr)
	if err != nil {
		log.Fatal(err)
	}
	http.Serve(listener, nil)
}
