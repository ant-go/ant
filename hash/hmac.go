package hash

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

// HmacMD5 hmac('md5', $data, $key)
func HmacMD5(s, k string) string {
	o := hmac.New(md5.New, []byte(k))
	_, _ = o.Write([]byte(s))
	sum := o.Sum(nil)
	return hex.EncodeToString(sum)
}

// HmacSHA1 hmac('sha1', $data, $key)
func HmacSHA1(s, k string) string {
	o := hmac.New(sha1.New, []byte(k))
	_, _ = o.Write([]byte(s))
	sum := o.Sum(nil)
	return hex.EncodeToString(sum)
}

// HmacSHA224 hmac('sha224', $data, $key)
func HmacSHA224(s, k string) string {
	o := hmac.New(sha256.New224, []byte(k))
	_, _ = o.Write([]byte(s))
	sum := o.Sum(nil)
	return hex.EncodeToString(sum)
}

// HmacSHA384 hmac('sha384', $data, $key)
func HmacSHA384(s, k string) string {
	o := hmac.New(sha512.New384, []byte(k))
	_, _ = o.Write([]byte(s))
	sum := o.Sum(nil)
	return hex.EncodeToString(sum)
}

// HmacSHA512224 hmac('sha512/224', $data, $key)
func HmacSHA512224(s, k string) string {
	o := hmac.New(sha512.New512_224, []byte(k))
	_, _ = o.Write([]byte(s))
	sum := o.Sum(nil)
	return hex.EncodeToString(sum)
}

// HmacSHA512256 hmac('sha512/256', $data, $key)
func HmacSHA512256(s, k string) string {
	o := hmac.New(sha512.New512_256, []byte(k))
	_, _ = o.Write([]byte(s))
	sum := o.Sum(nil)
	return hex.EncodeToString(sum)
}

// HmacSHA512 hmac('sha512', $data, $key)
func HmacSHA512(s, k string) string {
	o := hmac.New(sha512.New, []byte(k))
	_, _ = o.Write([]byte(s))
	sum := o.Sum(nil)
	return hex.EncodeToString(sum)
}

// HmacSHA3224 hmac('sha3-224', $data, $key)
func HmacSHA3224(s, k string) string {
	o := hmac.New(sha3.New224, []byte(k))
	_, _ = o.Write([]byte(s))
	sum := o.Sum(nil)
	return hex.EncodeToString(sum)
}

// HmacSHA3256 hmac('sha3-256', $data, $key)
func HmacSHA3256(s, k string) string {
	o := hmac.New(sha3.New256, []byte(k))
	_, _ = o.Write([]byte(s))
	sum := o.Sum(nil)
	return hex.EncodeToString(sum)
}

// HmacSHA3384 hmac('sha3-384', $data, $key)
func HmacSHA3384(s, k string) string {
	o := hmac.New(sha3.New384, []byte(k))
	_, _ = o.Write([]byte(s))
	sum := o.Sum(nil)
	return hex.EncodeToString(sum)
}

// HmacSHA3512 hmac('sha3-512', $data, $key)
func HmacSHA3512(s, k string) string {
	o := hmac.New(sha3.New512, []byte(k))
	_, _ = o.Write([]byte(s))
	sum := o.Sum(nil)
	return hex.EncodeToString(sum)
}
