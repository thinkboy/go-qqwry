package qqwry

import (
	"bytes"
	"fmt"
	iconv "github.com/zyxar/go-iconv"
	"net"
	"os"
)

const (
	_REDIRECT_MODE_1 = byte(0x01)
	_REDIRECT_MODE_2 = byte(0x02)
)

var locale string = "UTF-8"

type QQWry struct {
	*os.File
	first int64
	last  int64
}

func NewQQWry(path string) (*QQWry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return &QQWry{file, int64(read4byte(file, 0)), int64(read4byte(file, 4))}, nil
}

func (qw *QQWry) QueryIP(ip string) (countryStr, areaStr string) {
	var first, last, current int64

	addr := ip2Int64(ip)
	if addr == -1 {
		return "UNKNOWN", "UNKNOWN"
	}

	first = qw.first
	last = qw.last
	current = ((last-first)/7/2)*7 + first
	for current > first {
		if read4byte(qw, current) > addr {
			last = current
			current = ((last-first)/7/2)*7 + first
		} else {
			first = current
			current = ((last-first)/7/2)*7 + first
		}
	}

	offset := read3byte(qw, current+4)
	end := read4byte(qw, offset)
	if addr > end {
		return "UNKNOWN", "UNKNOWN"
	}

	offset += 4
	mode := read1byte(qw, offset)
	country := make([]byte, 256)
	var area []byte
	n := 0
	if mode == _REDIRECT_MODE_1 {
		offset = read3byte(qw, offset+1)
		if read1byte(qw, offset) == _REDIRECT_MODE_2 {
			off := read3byte(qw, offset+1)
			qw.ReadAt(country, off)
			n = bytes.IndexByte(country, 0x00)
			country = country[:n]
			offset += 4
		} else {
			qw.ReadAt(country, offset)
			n = bytes.IndexByte(country, 0x00)
			country = country[:n]
			offset += int64(n + 1)
		}
	} else if mode == _REDIRECT_MODE_2 {
		off := read3byte(qw, offset+1)
		qw.ReadAt(country, off)
		n = bytes.IndexByte(country, 0x00)
		country = country[:n]
		offset += 4
	} else {
		qw.ReadAt(country, offset)
		n = bytes.IndexByte(country, 0x00)
		country = country[:n]
		offset += int64(n + 1)
	}

	mode = read1byte(qw, offset)
	if mode == _REDIRECT_MODE_1 || mode == _REDIRECT_MODE_2 {
		offset = read3byte(qw, offset+1)
	}

	if offset != 0 {
		area = make([]byte, 256)
		qw.ReadAt(area, offset)
		n = bytes.IndexByte(area, 0x00)
		area = area[:n]
	}

	cz88 := []byte("CZ88.NET")
	if bytes.Compare(country[1:], cz88) != 0 {
		countryStr, _ = iconv.Conv(string(country), locale, "GBK")
	}

	if len(area) > 1 && bytes.Compare(area[1:], cz88) != 0 {
		areaStr, _ = iconv.Conv(string(area), locale, "GBK")
	}

	return countryStr, areaStr
}

func ip2Int64(ipStr string) int64 {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return -1
	}

	if len(ip) != 16 {
		return -1
	}

	return int64(ip[15]) | int64(ip[14])<<8 | int64(ip[13])<<16 | int64(ip[12])<<24
}
