package qqwry

var buffer []byte

func init() {
	buffer = make([]byte, 4)
}

type helper interface {
	ReadAt(b []byte, off int64) (n int, err error)
}

func read1byte(file helper, offset int64) byte {
	file.ReadAt(buffer, offset)
	return buffer[0]
}

func read2byte(file helper, offset int64) (ret int64) {
	file.ReadAt(buffer, offset)
	ret = int64(buffer[0]) | int64(buffer[1])<<8
	return
}

func read3byte(file helper, offset int64) (ret int64) {
	file.ReadAt(buffer, offset)
	ret = int64(buffer[0]) | int64(buffer[1])<<8 | int64(buffer[2])<<16
	return
}

func read4byte(file helper, offset int64) (ret int64) {
	file.ReadAt(buffer, offset)
	ret = int64(buffer[0]) | int64(buffer[1])<<8 | int64(buffer[2])<<16 | int64(buffer[3])<<24
	return
}

/*
func readbytes(ip net.IP) int64 {
	v := ip.To4()
	return int64(v[0])<<24 | int64(v[1])<<16 | int64(v[2])<<8 | int64(v[3])
}

func readslice(v int64) net.IP {
	return net.IPv4(byte((v&0xFF000000)>>24), byte((v&0xFF0000)>>16), byte((v&0xFF00)>>8), byte(v&0xFF))
}
*/
