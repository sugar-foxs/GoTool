// @program:     GoTool
// @file:        host.go
// @author:      sugar-foxs
// @create:      2021-06-26 13:55
// @description:
package system

import (
	"errors"
	"fmt"
	"net"
	"time"
)

func LocalIP() (net.IP, error) {
	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()
	iFaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	// 遍历全部网卡
	for _, iFace := range iFaces {
		if iFace.Flags&net.FlagUp == 0 {
			continue
		}
		if iFace.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrArr, err := iFace.Addrs()
		if err != nil {
			return nil, err
		}

		for _, addr := range addrArr {
			ip := GetIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("not connected to the network")
}

func GetIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}

	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}
