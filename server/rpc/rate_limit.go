package rpc

import (
	"strings"
	"tflgame/server/lib/limiter"
)

// Limiter implements a combination of rate limiters for a request.
type Limiter struct {
	IP limiter.Limiter
}

// Allow applies a rate limit based on ip, user, origin and a combination of
// ip and user.
func (l *Limiter) Allow(ip string) bool {
	_, ok := l.IP.Allow(ip)
	if !ok {
		return false
	}

	return true
}

func stripPort(addr string) string {
	if strings.Count(addr, ":") > 1 {
		if sq := strings.LastIndexByte(addr, ']'); sq > 1 {
			return addr[1:sq]
		}

		return addr
	}

	i := strings.LastIndexByte(addr, ':')
	if i == -1 {
		return addr
	}

	return addr[:i]
}
