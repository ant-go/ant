package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"hash/crc32"
	"io/ioutil"
	"strconv"
)

// MD5 md5()
func MD5(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

// MD5File md5_file()
func MD5File(filename string) string {
	s, _ := MD5FileE(filename)
	return s
}

// MD5FileE md5_file()
func MD5FileE(filename string) (s string, err error) {
	var data []byte
	if data, err = ioutil.ReadFile(filename); err != nil {
		return
	}
	sum := md5.Sum(data)
	s = hex.EncodeToString(sum[:])
	return
}

// SHA1 sha1()
func SHA1(s string) string {
	sum := sha1.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

// SHA1File sha1_file()
func SHA1File(filename string) string {
	s, _ := SHA1FileE(filename)
	return s
}

// SHA1FileE sha1_file()
func SHA1FileE(filename string) (s string, err error) {
	var data []byte
	if data, err = ioutil.ReadFile(filename); err != nil {
		return
	}
	sum := sha1.Sum(data)
	s = hex.EncodeToString(sum[:])
	return
}

// CRC32 crc32()
func CRC32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}

// CRC32String crc32()
func CRC32String(str string) string {
	return strconv.FormatUint(uint64(crc32.ChecksumIEEE([]byte(str))), 10)
}
