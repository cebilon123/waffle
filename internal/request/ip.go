package request

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
)

const (
	HeaderXForwardedFor    = "X-Forwarded-For"
	HeaderXForwarded       = "X-Forwarded"
	HeaderForwardedFor     = "Forwarded-For"
	HeaderForwarded        = "Forwarded"
	HeaderXClientIP        = "X-Client-Ip"
	HeaderCFConnectingIP   = "Cf-Connecting-Ip"
	HeaderFastlyClientIP   = "Fastly-Client-Ip"
	HeaderTrueClientIP     = "True-Client-Ip"
	HeaderXRealIP          = "X-Real-Ip"
	HeaderXClusterClientIP = "X-Cluster-Client-Ip"
)

// headers are the Request headers that can provide us with the ip addresses
var headers = []string{
	HeaderXForwardedFor,
	HeaderXForwarded,
	HeaderForwardedFor,
	HeaderForwarded,
	HeaderXClientIP,
	HeaderCFConnectingIP,
	HeaderFastlyClientIP,
	HeaderTrueClientIP,
	HeaderXRealIP,
	HeaderXClusterClientIP,
}

// GetRealIPAddress client can connect use directly, but can also
// be forwarded here by the i.e. Cloudflare, AWS etc.
func GetRealIPAddress(r http.Request) (net.IP, error) {
	addrStr, err := getIPStringFromHeaders(r)
	if err != nil {
		addrStr, err = getIPStringFromRequestRemoteAddress(r)
		if err != nil {
			return nil, fmt.Errorf("get ip string from Request remote address")
		}
	}

	if ipAddr := net.ParseIP(addrStr); ipAddr != nil {
		return ipAddr, nil
	}

	return nil, errors.New("cannot get ip address from the Request")
}

func getIPStringFromHeaders(r http.Request) (string, error) {
	for _, header := range headers {
		if header == HeaderXForwardedFor {
			v, ok := getFirstClientAddress(r.Header.Get(HeaderXForwardedFor))
			if !ok {
				continue
			}

			return v, nil
		}

		if addr := r.Header.Get(header); len(addr) != 0 {
			return r.Header.Get(header), nil
		}
	}

	return "", errors.New("there is no header with ip address")
}

func getIPStringFromRequestRemoteAddress(r http.Request) (string, error) {
	ipAddr, err := net.ResolveIPAddr("ip4", r.RemoteAddr)
	if err != nil {
		ipAddr, err = net.ResolveIPAddr("ip6", r.RemoteAddr)
		if err != nil {
			return "", fmt.Errorf("resolve ip address for ip4 and ip6: %w", err)
		}
	}

	return ipAddr.String(), nil
}

func getFirstClientAddress(header string) (string, bool) {
	if header == "" {
		return "", false
	}
	// x-forwarded-for may return multiple IP addresses in the format: "client IP, proxy 1 IP, proxy 2 IP"
	// Therefore, the right-most IP address is the IP address of the most recent proxy
	// and the left-most IP address is the IP address of the originating client.
	forwardedIps := strings.Split(header, ",")
	for _, ip := range forwardedIps {
		// header can contain spaces too, strip those out.
		ip = strings.TrimSpace(ip)
		// make sure we only use this if it's ipv4 (ip:port)
		if splitted := strings.Split(ip, ":"); len(splitted) == 2 {
			ip = splitted[0]
		}
		if isCorrectIP(ip) {
			return ip, true
		}
	}
	return "", false
}

func isCorrectIP(ip string) bool {
	return net.ParseIP(ip) != nil
}
