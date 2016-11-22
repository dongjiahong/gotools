package es

import (
	"log"
	"testing"
)

func Test_easyServer(t *testing.T) {
	var easy EasyServer
	err := easy.NewServer("../conf/easy_server.conf", "first", " TEST ")
	if err != nil {
		log.Fatalln("create server err: ", err)
	}
	easy.l.Println("yes , i'am ok!")
	easy.l.Println("yes , i will sining again! ")
}
