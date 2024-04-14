package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Point struct {
	X, Y float64
}

type Points struct {
	A, B Point
}

const addr = ":12345"
const network = "tcp4"

func main() {
	// Создаем клиента службы RPC
	client, err := rpc.DialHTTP(network, addr)
	if err != nil {
		log.Fatal(err)
	}
	var resp float64
	tochki := Points{A: Point{1, 1}, B: Point{4, 5}}

	// Удаленный вызов процедуры Server.Dist
	err = client.Call("Server.Dist", tochki, &resp)
	if err != nil {
		fmt.Println("ошибка:", err)
	}
	fmt.Println("distance:", resp)
}
