package hash

import (
	"testing"
)

func TestHmacMD5(t *testing.T) {
	type args struct {
		s string
		k string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"hmac_md5", args{"a", "b"}, "5e69fae25f4b4f3e8cc5de09a8163520"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HmacMD5(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("HmacMD5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHmacSHA1(t *testing.T) {
	type args struct {
		s string
		k string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"hmac_sha1", args{"a", "b"}, "8abe0fd691e3da3035f7b7ac91be45d99e942b9e"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HmacSHA1(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("HmacSHA1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHmacSHA224(t *testing.T) {
	type args struct {
		s string
		k string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"hmac_sha224", args{"a", "b"}, "a04a3802594be5c1ccaa49445f7a1170671c7d7f7eba996f05e6ca19"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HmacSHA224(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("HmacSHA224() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHmacSHA3224(t *testing.T) {
	type args struct {
		s string
		k string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"hmac_sha3224", args{"a", "b"}, "35cc1f9fc8d71d711580f4286702a42be679febe672511f30b2da378"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HmacSHA3224(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("HmacSHA3224() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHmacSHA3256(t *testing.T) {
	type args struct {
		s string
		k string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"hmac_sha3256", args{"a", "b"}, "877f24180180197a5501df0a2bd9bfe16141000f15036b80bfc984d9ecf0a193"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HmacSHA3256(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("HmacSHA3256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHmacSHA3384(t *testing.T) {
	type args struct {
		s string
		k string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"hmac_sha3384", args{"a", "b"}, "51720033fb14ce687ada5949250b2c01f825c642f48a63c932e7b9839154fea20e7c32e4dcca2a15877582b8c5a45f57"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HmacSHA3384(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("HmacSHA3384() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHmacSHA3512(t *testing.T) {
	type args struct {
		s string
		k string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"hmac_sha3512", args{"a", "b"}, "eacd6df033edbb2110c69ef1c7672b8b91ba581d940c111b4f068a4a6d0b7434ab6845c20df280e3c5c767ce25f49beff64b24a27543f3933ef77ca7b1dbb4e5"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HmacSHA3512(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("HmacSHA3512() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHmacSHA384(t *testing.T) {
	type args struct {
		s string
		k string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"hmac_sha384", args{"a", "b"}, "40292a79a280f609aef472272d2a2a0d148bfb83fa18fd2806440d07d16e31d9f0d398891063c11f6c73891e6f646106"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HmacSHA384(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("HmacSHA384() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHmacSHA512(t *testing.T) {
	type args struct {
		s string
		k string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"hmac_sha512", args{"a", "b"}, "68c1687fa7cb5170ff800580a0cec29dc0ccb515aaf95587bdfe5c923730a7852e2beefefd6be31d97aa612ad8b8569bba61ed2c339cd9b28409751b0b9e96a0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HmacSHA512(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("HmacSHA512() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHmacSHA512224(t *testing.T) {
	type args struct {
		s string
		k string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"hmac_sha512/244", args{"a", "b"}, "003a0e4031f8d193f27d260741e8fd08a3e54398c2e1f30f486877f2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HmacSHA512224(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("HmacSHA512224() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHmacSHA512256(t *testing.T) {
	type args struct {
		s string
		k string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"hmac_sha512/256", args{"a", "b"}, "776748e02b34f1335a1e4bf6dac40637e2321f5df0bb766b0b2de4f754a225db"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HmacSHA512256(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("HmacSHA512256() = %v, want %v", got, tt.want)
			}
		})
	}
}
