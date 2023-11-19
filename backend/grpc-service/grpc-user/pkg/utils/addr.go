package utils

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/17 14:48
//
import (
	"fmt"
	"net"
	"strings"
)

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

// GetLocalIPs 获取本地局域网IP地址列表
func GetLocalIPs() (string, error) {
	var localIPs []string

	// 获取所有网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", fmt.Errorf("获取网络接口失败: %v", err)
	}

	// 遍历网络接口
	for _, iface := range interfaces {
		// 排除一些特殊的接口
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		// 获取接口的地址信息
		addrs, err := iface.Addrs()
		if err != nil {
			return "", fmt.Errorf("获取地址信息失败: %v", err)
		}

		// 遍历地址信息
		for _, addr := range addrs {
			// 判断是否是 IP 地址
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				// 判断是否是 IPv4 地址
				if ipnet.IP.To4() != nil {
					localIPs = append(localIPs, ipnet.IP.String())
				}
			}
		}
	}

	return strings.Join(localIPs, "."), nil
}
