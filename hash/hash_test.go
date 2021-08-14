package hash

import (
	"io/ioutil"
	"os"
	"testing"
)

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

func createFile(pattern string, contents string) *os.File {
	f, err := ioutil.TempFile(os.TempDir(), pattern)
	if err != nil {
		panic(err)
	}
	_, _ = f.WriteString(contents)
	return f
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
