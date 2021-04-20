package redis

import "time"

/**
	报文格式
*3
$3
SET
$3
foo
$3
bar

*/
func (client *Client) Set(key string, val string) bool {
	client.WriteArgNums(3)
	client.WriteParam("SET")
	client.WriteParam("name")
	client.WriteParam("1234")
	returnBuf, _, _ := client.rbuf.ReadLine()
	if string(returnBuf) == "+OK" {
		return true
	}
	return false
}

/**
set name age ex 1
*/
func (client *Client) SetTimeOut(key string, val string, timeout time.Duration) bool {
	client.WriteArgNums(5)
	client.WriteParam("SET")
	client.WriteParam(key)
	client.WriteParam(val)
	client.WriteParam("ex")
	client.WriteParam("10")
	return true
}

/**
get aaa
*/
func (client *Client) Get(key string) bool {
	client.WriteArgNums(2)
	client.WriteParam("GET")
	client.WriteParam(key)
	return true
}
