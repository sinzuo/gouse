// routerudpserver
package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/garyburd/redigo/redis"
)

type WiFiClientt struct {
}

type Portstatet struct {
	port1 int
	port2 int
	port3 int
	port4 int
	port5 int
	port7 int
}

type Packett struct {
	ManufacturerOUI     string
	Memory_utilization  string
	Register            string
	ProductClass        string
	HardwareVersion     string
	Sn                  string
	Cpu_utilization     string
	Flash_utilization   string
	UpTime              string
	Work_mode           string
	DeviceDiqu          string
	Client_speed        string
	Dev_type            string
	SerialNumber        string
	Manufacturer        string
	SoftwareVersion     string
	Wan_current_ip_addr string
	WiFiClientt         []WiFiClientt
	Portstate           Portstatet
}
type A struct {
	Name         string
	Version      string
	Serialnumber string
	Commandkey   string
	Packet       Packett
}

func insert(addr string, value string) {
	c, err := redis.Dial("tcp", "192.168.3.242:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
	}
	defer c.Close()
	//取得hash
	r, err := redis.String(c.Do("HGet", "books", addr))
	if err != nil {
		fmt.Println("get abc failed,", err)

	}
	if value == r {
		fmt.Println("is same " + addr + "=" + value)
		return
	}
	//设置hashggsdf
	_, err = c.Do("HSet", "books", addr, value)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(r)
}

func get(addr string) (int, string) {
	c, err := redis.Dial("tcp", "192.168.3.242:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
	}
	defer c.Close()

	//取得hash
	r, err := redis.String(c.Do("HGet", "books", addr))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return 0, ""
	}
	fmt.Println(r)
	return 1, r
}

var msg = `{
    "name": "set",
    "version": "1.0.0",
    "serialnumber": "78A351323857",
    "keyname": "download",
    "packet": {
        "url": "1",
        "FileSize": "1",
        "md5": ""
    }
}`

var deviceId = "78A351323857"

func sendCommand() {
	time.Sleep(time.Second * 5)

	i, value := get("78A351323857")
	if i == 0 {
		return
	}

	p1, _ := net.Dial("udp", value)
	fmt.Sprintf(msg, deviceId)
	fmt.Println(msg)
	p1.Write([]byte(msg))
}

var returnstring = `{"name":"informResponse","version":"1.0.0","serialnumber":"78A351323857","keyname":"inform","packet":{"serialnumber":"78A351323857"}}`

func main() {
	b1 := make([]byte, 2048)
	b2 := A{}
	go sendCommand()
	p1, _ := net.ResolveUDPAddr("udp", ":8880")
	d1, _ := net.ListenUDP("udp", p1)
	for {
		n1, adr, e := d1.ReadFromUDP(b1)

		if e != nil {
			fmt.Println("error")
		} else {
			fmt.Println(string(b1))
			json.Unmarshal(b1[0:n1], &b2)
			insert(b2.Serialnumber, adr.String())
			fmt.Println(b2.Serialnumber, b2.Packet.HardwareVersion, adr.String())
			p1, _ := net.Dial("udp", adr)

			p1.Write([]byte(returnstring))
		}
	}
	fmt.Println("Hello World!")
}
