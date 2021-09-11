package datetime

import (
	"math"
	"os"
	"strconv"
	"testing"
	"time"
)

const format = "Y-m-d H:i:s"

func TestTime(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"time():second", true},
		{"time():millisecond", true},
		{"time():microsecond", true},
		{"time():nanosecond", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Time() > 0; got != tt.want {
				t.Errorf("Time() = %v, want %v", got, tt.want)
			}
			if got := TimeMs() > 0; got != tt.want {
				t.Errorf("TimeMs() = %v, want %v", got, tt.want)
			}
			if got := TimeUs() > 0; got != tt.want {
				t.Errorf("TimeUs() = %v, want %v", got, tt.want)
			}
			if got := TimeNs() > 0; got != tt.want {
				t.Errorf("TimeNs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrToTime(t *testing.T) {
	type args struct {
		s    string
		args []interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty", args{"", nil}, -1},
		{"now", args{"now", nil}, Time()},
		{"now", args{"now", []interface{}{"?"}}, Time()},
		{"2010-10-10 10:10:10 +2 weeks", args{"+2 weeks", []interface{}{"1286676610"}}, 1286676610 + 24*3600*7*2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got int
			if tt.args.args == nil {
				got = StrToTime(tt.args.s)
			} else {
				got = StrToTime(tt.args.s, tt.args.args...)
			}
			if got == tt.want {
				return
			}
			if math.Abs(float64(got)-float64(tt.want)) <= 1 {
				return
			}
			t.Errorf("StrToTime() = %v, want [%v, %v]", got, tt.want-1, tt.want+1)
		})
	}
}

func TestDateGMT0(t *testing.T) {
	timezone := "GMT0"
	_ = os.Setenv("TZ", timezone)
	if loc, err := time.LoadLocation(timezone); err == nil {
		time.Local = loc
	}
	y := strconv.FormatInt(int64(time.Now().Year()), 10)

	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1999-10-13 00:00:00", args{format, []interface{}{StrToTime("1999-10-13")}}, "1999-10-13 00:00:00"},
		{"1999-10-13 00:00:00", args{format, []interface{}{StrToTime("Oct 13  1999")}}, "1999-10-13 00:00:00"},
		{"2000-01-19 00:00:00", args{format, []interface{}{StrToTime("2000-01-19")}}, "2000-01-19 00:00:00"},
		{"2000-01-19 00:00:00", args{format, []interface{}{StrToTime("Jan 19  2000")}}, "2000-01-19 00:00:00"},
		{"2001-12-21 00:00:00", args{format, []interface{}{StrToTime("2001-12-21")}}, "2001-12-21 00:00:00"},
		{"2001-12-21 00:00:00", args{format, []interface{}{StrToTime("Dec 21  2001")}}, "2001-12-21 00:00:00"},
		{"2001-12-21 12:16:00", args{format, []interface{}{StrToTime("2001-12-21 12:16")}}, "2001-12-21 12:16:00"},
		{"2001-12-21 12:16:00", args{format, []interface{}{StrToTime("Dec 21 2001 12:16")}}, "2001-12-21 12:16:00"},
		{y + "-12-21 12:16:00", args{format, []interface{}{StrToTime("Dec 21  12:16")}}, y + "-12-21 12:16:00"},
		{"2001-10-22 21:19:58", args{format, []interface{}{StrToTime("2001-10-22 21:19:58")}}, "2001-10-22 21:19:58"},
		{"2001-10-22 23:19:58", args{format, []interface{}{StrToTime("2001-10-22 21:19:58-02")}}, "2001-10-22 23:19:58"},
		{"2001-10-22 23:32:58", args{format, []interface{}{StrToTime("2001-10-22 21:19:58-0213")}}, "2001-10-22 23:32:58"},
		{"2001-10-22 19:19:58", args{format, []interface{}{StrToTime("2001-10-22 21:19:58+02")}}, "2001-10-22 19:19:58"},
		{"2001-10-22 19:06:58", args{format, []interface{}{StrToTime("2001-10-22 21:19:58+0213")}}, "2001-10-22 19:06:58"},
		{"2001-10-23 01:00:58", args{format, []interface{}{StrToTime("2001-10-22T21:20:58-03:40")}}, "2001-10-23 01:00:58"},
		{"2001-10-22 23:19:58", args{format, []interface{}{StrToTime("2001-10-22T211958-2")}}, "2001-10-22 23:19:58"},
		{"2001-10-22 19:06:58", args{format, []interface{}{StrToTime("20011022T211958+0213")}}, "2001-10-22 19:06:58"},
		{"2001-10-22 19:05:00", args{format, []interface{}{StrToTime("20011022T21:20+0215")}}, "2001-10-22 19:05:00"},
		{"1996-12-30 00:00:00", args{format, []interface{}{StrToTime("1997W011")}}, "1996-12-30 00:00:00"},
		{"2004-03-01 05:00:00", args{format, []interface{}{StrToTime("2004W101T05:00+0")}}, "2004-03-01 05:00:00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Date(tt.args.format, tt.args.args...); got != tt.want {
				t.Errorf("Date() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateUSEastern(t *testing.T) {
	timezone := "US/Eastern"
	_ = os.Setenv("TZ", timezone)
	if loc, err := time.LoadLocation(timezone); err == nil {
		time.Local = loc
	}
	y := strconv.FormatInt(int64(time.Now().Year()), 10)

	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1999-10-13 00:00:00", args{format, []interface{}{StrToTime("1999-10-13")}}, "1999-10-13 00:00:00"},
		{"1999-10-13 00:00:00", args{format, []interface{}{StrToTime("Oct 13  1999")}}, "1999-10-13 00:00:00"},
		{"2000-01-19 00:00:00", args{format, []interface{}{StrToTime("2000-01-19")}}, "2000-01-19 00:00:00"},
		{"2000-01-19 00:00:00", args{format, []interface{}{StrToTime("Jan 19  2000")}}, "2000-01-19 00:00:00"},
		{"2001-12-21 00:00:00", args{format, []interface{}{StrToTime("2001-12-21")}}, "2001-12-21 00:00:00"},
		{"2001-12-21 00:00:00", args{format, []interface{}{StrToTime("Dec 21  2001")}}, "2001-12-21 00:00:00"},
		{"2001-12-21 12:16:00", args{format, []interface{}{StrToTime("2001-12-21 12:16")}}, "2001-12-21 12:16:00"},
		{"2001-12-21 12:16:00", args{format, []interface{}{StrToTime("Dec 21 2001 12:16")}}, "2001-12-21 12:16:00"},
		{y + "-12-21 12:16:00", args{format, []interface{}{StrToTime("Dec 21  12:16")}}, y + "-12-21 12:16:00"},
		{"2001-10-22 21:19:58", args{format, []interface{}{StrToTime("2001-10-22 21:19:58")}}, "2001-10-22 21:19:58"},
		{"2001-10-22 19:19:58", args{format, []interface{}{StrToTime("2001-10-22 21:19:58-02")}}, "2001-10-22 19:19:58"},
		{"2001-10-22 19:32:58", args{format, []interface{}{StrToTime("2001-10-22 21:19:58-0213")}}, "2001-10-22 19:32:58"},
		{"2001-10-22 15:19:58", args{format, []interface{}{StrToTime("2001-10-22 21:19:58+02")}}, "2001-10-22 15:19:58"},
		{"2001-10-22 15:06:58", args{format, []interface{}{StrToTime("2001-10-22 21:19:58+0213")}}, "2001-10-22 15:06:58"},
		{"2001-10-22 21:00:58", args{format, []interface{}{StrToTime("2001-10-22T21:20:58-03:40")}}, "2001-10-22 21:00:58"},
		{"2001-10-22 19:19:58", args{format, []interface{}{StrToTime("2001-10-22T211958-2")}}, "2001-10-22 19:19:58"},
		{"2001-10-22 15:06:58", args{format, []interface{}{StrToTime("20011022T211958+0213")}}, "2001-10-22 15:06:58"},
		{"2001-10-22 15:05:00", args{format, []interface{}{StrToTime("20011022T21:20+0215")}}, "2001-10-22 15:05:00"},
		{"1996-12-30 00:00:00", args{format, []interface{}{StrToTime("1997W011")}}, "1996-12-30 00:00:00"},
		{"2004-03-01 00:00:00", args{format, []interface{}{StrToTime("2004W101T05:00+0")}}, "2004-03-01 00:00:00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Date(tt.args.format, tt.args.args...); got != tt.want {
				t.Errorf("Date() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGMDate(t *testing.T) {
	var tmp *time.Location
	time.Local, tmp = time.UTC, time.Local
	y := strconv.FormatInt(int64(time.Now().Year()), 10)

	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Y-m-d H:i:s", args{"Y-m-d H:i:s", []interface{}{1631200022}}, "2021-09-09 15:07:02"},
		{y, args{"Y", nil}, y},
		{y, args{"Y", []interface{}{"?"}}, y},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			if tt.args.args == nil {
				got = GMDate(tt.args.format)
			} else {
				got = GMDate(tt.args.format, tt.args.args...)
			}
			if got != tt.want {
				t.Errorf("GMDate() = %v, want %v", got, tt.want)
			}
		})
	}
	time.Local = tmp
}

func TestDate(t *testing.T) {
	y := strconv.FormatInt(int64(time.Now().Year()), 10)
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{y, args{"Y", nil}, y},
		{y, args{"Y", []interface{}{"?"}}, y},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			if tt.args.args == nil {
				got = Date(tt.args.format)
			} else {
				got = Date(tt.args.format, tt.args.args...)
			}
			if got != tt.want {
				t.Errorf("Date() = %v, want %v", got, tt.want)
			}
		})
	}
}