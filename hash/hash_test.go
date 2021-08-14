package hash

import (
	"io/ioutil"
	"os"
	"testing"
)

func createFile(pattern string, contents string) *os.File {
	f, err := ioutil.TempFile(os.TempDir(), pattern)
	if err != nil {
		panic(err)
	}
	_, _ = f.WriteString(contents)
	return f
}

func TestMD5(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"md5_abc", args{"abc"}, "900150983cd24fb0d6963f7d28e17f72"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MD5(tt.args.s); got != tt.want {
				t.Errorf("MD5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMD5File(t *testing.T) {
	f := createFile(t.Name(), "abc")
	defer func() { _ = f.Close() }()

	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"md5_file_empty", args{""}, ""},
		{"md5_file_abc", args{f.Name()}, "900150983cd24fb0d6963f7d28e17f72"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MD5File(tt.args.filename); got != tt.want {
				t.Errorf("MD5File() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMD5FileE(t *testing.T) {
	f := createFile(t.Name(), "abc")
	defer func() { _ = f.Close() }()

	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantS   string
		wantErr bool
	}{
		{"md5_file_empty", args{""}, "", true},
		{"md5_file_abc", args{f.Name()}, "900150983cd24fb0d6963f7d28e17f72", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, err := MD5FileE(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("MD5FileE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotS != tt.wantS {
				t.Errorf("MD5FileE() gotS = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestSHA1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"sha1_abc", args{"abc"}, "a9993e364706816aba3e25717850c26c9cd0d89d"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA1(tt.args.s); got != tt.want {
				t.Errorf("SHA1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA1File(t *testing.T) {
	f := createFile(t.Name(), "abc")
	defer func() { _ = f.Close() }()

	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"sha1_file_empty", args{""}, ""},
		{"sha1_file_abc", args{f.Name()}, "a9993e364706816aba3e25717850c26c9cd0d89d"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA1File(tt.args.filename); got != tt.want {
				t.Errorf("SHA1File() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA1FileE(t *testing.T) {
	f := createFile(t.Name(), "abc")
	defer func() { _ = f.Close() }()

	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantS   string
		wantErr bool
	}{
		{"sha1_file_empty", args{""}, "", true},
		{"sha1_file_abc", args{f.Name()}, "a9993e364706816aba3e25717850c26c9cd0d89d", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, err := SHA1FileE(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("SHA1FileE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotS != tt.wantS {
				t.Errorf("SHA1FileE() gotS = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestCRC32(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"crc32_abc", args{"abc"}, 891568578},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CRC32(tt.args.str); got != tt.want {
				t.Errorf("CRC32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCRC32String(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"crc32_abc", args{"abc"}, "891568578"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CRC32String(tt.args.str); got != tt.want {
				t.Errorf("CRC32String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEquals(t *testing.T) {
	type args struct {
		known string
		user  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"hash_equals", args{"abc", "edf"}, false},
		{"hash_equals", args{"abc", "abc"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equals(tt.args.known, tt.args.user); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA224(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"sha224_abc", args{"abc"}, "23097d223405d8228642a477bda255b32aadbce4bda0b3f7e36c9da7"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA224(tt.args.s); got != tt.want {
				t.Errorf("SHA224() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA384(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"sha384_abc", args{"abc"}, "cb00753f45a35e8bb5a03d699ac65007272c32ab0eded1631a8b605a43ff5bed8086072ba1e7cc2358baeca134c825a7"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA384(tt.args.s); got != tt.want {
				t.Errorf("SHA384() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA512224(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"sha512/224_abc", args{"abc"}, "4634270f707b6a54daae7530460842e20e37ed265ceee9a43e8924aa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA512224(tt.args.s); got != tt.want {
				t.Errorf("SHA512224() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA512256(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"sha512/256_abc", args{"abc"}, "53048e2681941ef99b2e29b76b4c7dabe4c2d0c634fc6d46e0e2f13107e7af23"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA512256(tt.args.s); got != tt.want {
				t.Errorf("SHA512256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA512(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"sha512_abc", args{"abc"}, "ddaf35a193617abacc417349ae20413112e6fa4e89a97ea20a9eeee64b55d39a2192992a274fc1a836ba3c23a3feebbd454d4423643ce80e2a9ac94fa54ca49f"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA512(tt.args.s); got != tt.want {
				t.Errorf("SHA512() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA3224(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"sha3-224_abc", args{"abc"}, "e642824c3f8cf24ad09234ee7d3c766fc9a3a5168d0c94ad73b46fdf"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA3224(tt.args.s); got != tt.want {
				t.Errorf("SHA3224() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA3256(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"sha3-256_abc", args{"abc"}, "3a985da74fe225b2045c172d6bd390bd855f086e3e9d525b46bfe24511431532"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA3256(tt.args.s); got != tt.want {
				t.Errorf("SHA3256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA3384(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"sha3-384_abc", args{"abc"}, "ec01498288516fc926459f58e2c6ad8df9b473cb0fc08c2596da7cf0e49be4b298d88cea927ac7f539f1edf228376d25"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA3384(tt.args.s); got != tt.want {
				t.Errorf("SHA3384() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA3512(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"sha3-512_abc", args{"abc"}, "b751850b1a57168a5693cd924b6b096e08f621827444f70d884f5d0240d2712e10e116e9192af3c91a7ec57647e3934057340b4cf408d5a56592f8274eec53f0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA3512(tt.args.s); got != tt.want {
				t.Errorf("SHA3512() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFNV132(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"fnv132_abc", args{"abc"}, "439c2f4b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FNV132(tt.args.s); got != tt.want {
				t.Errorf("FNV132() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFNV1a32(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"fnv1a32_abc", args{"abc"}, "1a47e90b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FNV1a32(tt.args.s); got != tt.want {
				t.Errorf("FNV1a32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFNV164(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"fnv164_abc", args{"abc"}, "d8dcca186bafadcb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FNV164(tt.args.s); got != tt.want {
				t.Errorf("FNV164() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFNV1a64(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"fnv1a64_abc", args{"abc"}, "e71fa2190541574b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FNV1a64(tt.args.s); got != tt.want {
				t.Errorf("FNV1a64() = %v, want %v", got, tt.want)
			}
		})
	}
}
