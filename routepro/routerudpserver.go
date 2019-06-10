// routerudpserver
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/garyburd/redigo/redis"
)

var REDIS_CONNS = "192.168.3.242:6379"

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
	Packet       map[string]interface{}
}

var cmdSend = make(map[string]string)

func Init(){
	cmdSend["download"] = "js/download.json"
	cmdSend["set"] = "js/set.json"
	cmdSend["get"] = "js/get.json"
	cmdSend["file"] = "js/get.json"
}

func insert(addr string, value string) {
	c, err := redis.Dial("tcp", REDIS_CONNS)
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

func getRedisData(addr string) (int, string) {
	c, err := redis.Dial("tcp", REDIS_CONNS)
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

var msg1 = `{
    "name": "get",
    "version": "1.0.0",
    "serialnumber": "78A351323857",
    "keyname": "inform",
    "packet": {
        "black_mac": "1",
        "white_mac": "1",
        "black_url": "1",
        "white_url": "1"
    }
}`

var deviceId = "78A351323857"

type B struct {
	Keyname string `json:"keyname"`
	Name    string `json:"name"`
	Packet  struct {
		InternetGatewayDevice_DeviceInfo_HardwareVersion      string `json:"InternetGatewayDevice.DeviceInfo.HardwareVersion"`
		InternetGatewayDevice_DeviceInfo_Manufacturer         string `json:"InternetGatewayDevice.DeviceInfo.Manufacturer"`
		InternetGatewayDevice_DeviceInfo_ProductClass         string `json:"InternetGatewayDevice.DeviceInfo.ProductClass"`
		InternetGatewayDevice_DeviceInfo_SoftwareVersion      string `json:"InternetGatewayDevice.DeviceInfo.SoftwareVersion"`
		InternetGatewayDevice_DeviceInfo_UpTime               string `json:"InternetGatewayDevice.DeviceInfo.UpTime"`
		InternetGatewayDevice_DeviceInfo_clientSpeed          string `json:"InternetGatewayDevice.DeviceInfo.client_speed"`
		InternetGatewayDevice_DeviceInfo_cpuUtilization       string `json:"InternetGatewayDevice.DeviceInfo.cpu_utilization"`
		InternetGatewayDevice_DeviceInfo_flashUtilization     string `json:"InternetGatewayDevice.DeviceInfo.flash_utilization"`
		InternetGatewayDevice_DeviceInfo_memoryUtilization    string `json:"InternetGatewayDevice.DeviceInfo.memory_utilization"`
		InternetGatewayDevice_DeviceInfo_wanCurrentIPAddr     string `json:"InternetGatewayDevice.DeviceInfo.wan_current_ip_addr"`
		InternetGatewayDevice_DeviceInfo_workMode             string `json:"InternetGatewayDevice.DeviceInfo.work_mode"`
		InternetGatewayDevice_LANDevice_1_Wireless_WiFiClient []struct {
			Host string `json:"host"`
			IP   string `json:"ip"`
			Mac  string `json:"mac"`
			Type string `json:"type"`
		} `json:"InternetGatewayDevice.LANDevice.1.Wireless.WiFiClient"`
		ManufaInternetGatewayDevice_DeviceInfo_cturerOUI string `json:"ManufaInternetGatewayDevice.DeviceInfo.cturerOUI"`
		SerialNumber                                     string `json:"SerialNumber"`
		DevType                                          string `json:"dev_type"`
	} `json:"packet"`
	Serialnumber string `json:"serialnumber"`
	Version      string `json:"version"`
}

func sendCommand(dev string) {


	i, value := get(dev)
	if i == 0 {
		return
	}

	p1, _ := net.Dial("udp", value)
	fmt.Sprintf(msg, deviceId)
	fmt.Println(msg)
	p1.Write([]byte(msg1))
}

var returnstring = `{"name":"informResponse","version":"1.0.0","serialnumber":"78A351323857","keyname":"inform","packet":{"serialnumber":"78A351323857"}}`

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.PostForm)
	buf, _ := ioutil.ReadAll(r.Body)
	fmt.Println("jiangyibo" + string(buf))
	if r.Header.get("Content-Type") == "application/x-www-form-urlencoded" {
		if r.Method == "POST" {
			var (
				key   string = r.PostFormValue("key")
				value string = r.PostFormValue("value")
			)
			fmt.Printf("key is  : %s\n", key)
			fmt.Printf("value is: %s\n", value)
		}
	} else if r.Header.get("Content-Type") == "text/json" {
		var bb = &B{}
		json.Unmarshal(buf, bb)
		fmt.Println(bb.Keyname)

	}

	fmt.Fprintf(w, "Hi, This is an example of http service in golang!")
}

func apihandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.PostForm)
	buf, _ := ioutil.ReadAll(r.Body)
	fmt.Println("jiangyibo" + string(buf))
	if r.Header.get("Content-Type") == "application/x-www-form-urlencoded" {
		if r.Method == "POST" {
			var (
				key   string = r.PostFormValue("key")
				value string = r.PostFormValue("value")
				name string = r.PostFormValue("name")
			)
			fmt.Printf("key is  : %s\n", key)
			fmt.Printf("value is: %s\n", value)
			fmt.Printf("name is: %s\n", name)
			sendCommand(key)
		}
	} else if r.Header.get("Content-Type") == "text/json" {
		var bb = &B{}
		json.Unmarshal(buf, bb)
		fmt.Println(bb.Keyname)

	}

	fmt.Fprintf(w, "Hi, This is an example of http service in golang!")
}
func httpmain() {
	http.HandleFunc("/wifi/testapi", apihandler)
	http.HandleFunc("/wifi/acsjson", handler)	
	http.ListenAndServe(":8081", nil)
}

func main() {
	b1 := make([]byte, 2048)
	b2 := A{}
	go httpmain()
	 

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
			p1, _ := net.Dial("udp", adr.String())

			p1.Write([]byte(returnstring))
		}
	}
	fmt.Println("Hello World!")
}
