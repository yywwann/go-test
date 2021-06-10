package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	addr, _ := net.InterfaceAddrs()
	fmt.Println("---InterfaceAddrs---")
	fmt.Println(addr) //[127.0.0.1/8 10.236.15.24/24 ::1/128 fe80::3617:ebff:febe:f123/64],本地地址,ipv4和ipv6地址,这些信息可通过ifconfig命令看到

	fmt.Println("---Interfaces---")
	interfaces, _ := net.Interfaces()
	fmt.Printf("%+v\n", interfaces) //[{1 65536 lo  up|loopback} {2 1500 eth0 34:17:eb:be:f1:23 up|broadcast|multicast}] 类型:Index, MTU(最大传输单元), Name网络接口名, Flags支持状态

	fmt.Println("---JoinHostPort---")
	hp := net.JoinHostPort("127.0.0.1", "8080")
	fmt.Println(hp) //127.0.0.1:8080,根据ip和端口组成一个addr字符串表示

	fmt.Println("---LookupAddr---")
	lt, _ := net.LookupAddr("127.0.0.1")
	fmt.Println(lt) //[localhost],根据地址查找到改地址的一个映射列表

	fmt.Println("---LookupCNAME---")
	cname, _ := net.LookupCNAME("www.baidu.com")
	fmt.Println(cname) //www.a.shifen.com,查找规范的dns主机名字

	fmt.Println("---LookupHost---")
	host, _ := net.LookupHost("www.baidu.com")
	fmt.Println(host) //[111.13.100.92 111.13.100.91],查找给定域名的host名称

	fmt.Println("---LookupIP---")
	ip, _ := net.LookupIP("www.baidu.com")
	fmt.Println(ip) //[111.13.100.92 111.13.100.91],查找给定域名的ip地址,可通过nslookup www.baidu.com进行查找操作.


	fmt.Println("---InternalIP---")
	fmt.Println(InternalIP())
}

func InternalIP() string {
	inters, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, inter := range inters {
		if inter.Flags&net.FlagUp != 0 && !strings.HasPrefix(inter.Name, "lo") {
			addrs, err := inter.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}
	return ""
}