package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/subtle"
	"encoding/hex"
	"hash/crc32"
	"hash/fnv"
	"io/ioutil"
	"strconv"

	"golang.org/x/crypto/sha3"
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

// Equals hash_equals()
func Equals(known, user string) bool {
	return subtle.ConstantTimeCompare([]byte(known), []byte(user)) == 1
}

// SHA224 hash('sha224', $data)
func SHA224(s string) string {
	o := sha256.New224()
	_, _ = o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

// SHA384 hash('sha384', $data)
func SHA384(s string) string {
	o := sha512.New384()
	_, _ = o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

// SHA512224 hash('sha512/224', $data)
func SHA512224(s string) string {
	o := sha512.New512_224()
	_, _ = o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

// SHA512256 hash('sha512/256', $data)
func SHA512256(s string) string {
	o := sha512.New512_256()
	_, _ = o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

// SHA512 hash('sha512', $data)
func SHA512(s string) string {
	o := sha512.New()
	_, _ = o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

// SHA3224 hash('sha3-224', $data)
func SHA3224(s string) string {
	o := sha3.New224()
	_, _ = o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

// SHA3256 hash('sha3-256', $data)
func SHA3256(s string) string {
	o := sha3.New256()
	_, _ = o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

// SHA3384 hash('sha3-384', $data)
func SHA3384(s string) string {
	o := sha3.New384()
	_, _ = o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

// SHA3512 hash('sha3-512', $data)
func SHA3512(s string) string {
	o := sha3.New512()
	_, _ = o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

// FNV132 hash('fnv132', $data)
func FNV132(s string) string {
	o := fnv.New32()
	_, _ = o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

// FNV1a32 hash('fnv1a32', $data)
func FNV1a32(s string) string {
	o := fnv.New32a()
	_, _ = o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

// FNV164 hash('fnv164', $data)
func FNV164(s string) string {
	o := fnv.New64()
	_, _ = o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

// FNV1a64 hash('fnv1a64', $data)
func FNV1a64(s string) string {
	o := fnv.New64a()
	_, _ = o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}
