package coordinator

import "crypto/md5"

// Converts IpAddr and Port to a md5 hash identifier
func AddrToId(ip, port string) string {
	ipaddr := ip + ":" + port
	h := md5.New()
	h.Write([]byte(ipaddr))
	bs := h.Sum(nil)

	return string(bs)
}
