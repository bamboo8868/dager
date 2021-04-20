package redis

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
)

type Config struct {
	Port int32
	Addr []byte
	Pass []byte
	DB   int32
}

type Client struct {
	Conn net.Conn
	rbuf *bufio.Reader
	wbuf *bufio.Writer
}

func handleConn(client *Client) {
	for {
		client.wbuf.Flush()

		buf := make([]byte,0,1024)
		length ,err := client.rbuf.Read(buf)

		if err != nil {
			log.Println(err)
		}

		log.Println(string(buf[:length]))

	}
}

func (client *Client) WriteBuf(buf []byte)  {
	client.wbuf.Write(buf)
}

func New(config *Config) *Client {
	addr := string(config.Addr) + ":" + strconv.Itoa(int(config.Port))
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		panic("error config")
	}
	client := &Client{}
	client.Conn = conn
	client.wbuf = bufio.NewWriter(conn)
	client.rbuf = bufio.NewReader(conn)
	go handleConn(client)
	return client
}

func (client *Client) WriteArgNums(num int)  {
	client.wbuf.Write([]byte("*"))
	client.wbuf.Write([]byte(strconv.Itoa(num)))
	client.wbuf.Write([]byte("\r\n"))
}

func (client *Client) WriteParam(val string )  {
	length := len(val)
	client.wbuf.Write([]byte("$"))
	client.wbuf.Write([]byte(strconv.Itoa(length)))
	client.wbuf.Write([]byte("\r\n"))
	client.wbuf.Write([]byte(val))
	client.wbuf.Write([]byte("\r\n"))
}
