package web

import (
	"ddns-go/v6/config"
	"fmt"
	"net"
	"net/http"
	"strings"
)

func Receiver(writer http.ResponseWriter, request *http.Request) {
	token := request.Header.Get("auth")
	time := request.PostFormValue("time")
	// todo 验证token

	ipaddr := getClientIP(request)
	config.HttpReceiveIp.Store("ipv4", ipaddr)
	fmt.Printf("收到请求: token:[%s],time:[%s],ipaddr:[%s] ", token, time, ipaddr)
	returnOK(writer, fmt.Sprintf("更新成功[%s]", ipaddr), nil)
}

func getClientIP(r *http.Request) string {
	// 获取 X-Forwarded-For 中的第一个 IP
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		// 可能有多个IP，用逗号分隔，第一个是原始客户端IP
		ips := strings.Split(xForwardedFor, ",")
		if len(ips) > 0 {
			return strings.TrimSpace(ips[0])
		}
	}

	// 检查 X-Real-Ip
	xRealIP := r.Header.Get("X-Real-IP")
	if xRealIP != "" {
		return xRealIP
	}

	// 最后回退到 RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr // 当无法解析时返回原始值
	}
	return ip
}
