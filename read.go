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
