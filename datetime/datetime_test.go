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

func TestDateDST(t *testing.T) {
	tz := os.Getenv("TZ")
	_ = os.Setenv("TZ", "America/Toronto")
	var tmp = time.Local
	time.Local, _ = time.LoadLocation("America/Toronto")

	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Y-m-d H:i:sp", args{"Y-m-d H:i:sp", []interface{}{StrToTime("2010-10-10 10:10:11")}}, "2010-10-10 10:10:11-04:00"},
		{"I", args{"I", []interface{}{StrToTime("2010-10-10 10:10:10")}}, "1"},
		{"c", args{"c", []interface{}{StrToTime("2010-10-10 10:10:10")}}, "2010-10-10T10:10:10-04:00"},
		{"r", args{"r", []interface{}{StrToTime("2010-10-10 10:10:10")}}, "Sun, 10 Oct 2010 10:10:10 -0400"},
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
	time.Local = tmp
	_ = os.Setenv("TZ", tz)
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
		{"Y-m-d T H:i:s", args{"Y-m-d \\T H:i:s", []interface{}{1631200022}}, "2021-09-09 T 15:07:02"},
		{"Y-m-d H:i:s", args{"Y-m-d H:i:s", []interface{}{1631200022}}, "2021-09-09 15:07:02"},
		{y, args{"Y", nil}, y},
		{y, args{"Y", []interface{}{"?"}}, y},
		{"T", args{"T", nil}, "GMT"},
		{"e", args{"e", nil}, "UTC"},
		{"Z", args{"Z", nil}, "0"},
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
	var tmp *time.Location
	time.Local, tmp = time.UTC, time.Local

	parse := func(v string) int {
		_t, _ := time.Parse("2006-01-02 15:04:05", v)
		return int(_t.Unix())
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
		{y, args{"Y", nil}, y},
		{y, args{"Y", []interface{}{"?"}}, y},
		{"D", args{"D", []interface{}{parse("2010-10-10 10:10:10")}}, "Sun"},
		{"j", args{"j", []interface{}{parse("2010-10-01 10:10:10")}}, "1"},
		{"l", args{"l", []interface{}{parse("2010-10-10 10:10:10")}}, "Sunday"},
		{"S_st", args{"S", []interface{}{parse("2010-10-01 10:10:10")}}, "st"},
		{"S_nd", args{"S", []interface{}{parse("2010-10-02 10:10:10")}}, "nd"},
		{"S_rd", args{"S", []interface{}{parse("2010-10-03 10:10:10")}}, "rd"},
		{"S_th", args{"S", []interface{}{parse("2010-10-10 10:10:10")}}, "th"},
		{"S_th", args{"S", []interface{}{parse("2010-10-04 10:10:10")}}, "th"},
		{"w", args{"w", []interface{}{parse("2010-10-10 10:10:10")}}, "0"},
		{"N", args{"N", []interface{}{parse("2010-10-10 10:10:10")}}, "7"},
		{"z", args{"z", []interface{}{parse("2010-10-10 10:10:10")}}, "282"},
		{"z_leep", args{"z", []interface{}{parse("2000-10-10 10:10:10")}}, "283"},
		{"W", args{"W", []interface{}{parse("2010-10-10 10:10:10")}}, "40"},
		{"o", args{"o", []interface{}{parse("2010-10-10 10:10:10")}}, "2010"},
		{"F", args{"F", []interface{}{parse("2010-10-10 10:10:10")}}, "October"},
		{"m", args{"m", []interface{}{parse("2010-10-10 10:10:10")}}, "10"},
		{"M", args{"M", []interface{}{parse("2010-10-10 10:10:10")}}, "Oct"},
		{"n", args{"n", []interface{}{parse("2010-10-10 10:10:10")}}, "10"},
		{"t", args{"t", []interface{}{parse("2010-10-10 10:10:10")}}, "31"},
		{"t_leep", args{"t", []interface{}{parse("2000-02-10 10:10:10")}}, "29"},
		{"L_0", args{"L", []interface{}{parse("2010-10-10 10:10:10")}}, "0"},
		{"L_1", args{"L", []interface{}{parse("2000-10-10 10:10:10")}}, "1"},
		{"y", args{"y", []interface{}{parse("2010-10-10 10:10:10")}}, "10"},
		{"Y_+", args{"Y", []interface{}{parse("2010-10-10 10:10:10")}}, "2010"},
		{"Y_-", args{"Y", []interface{}{StrToTime("-2010-10-10 10:10:10")}}, "-2010"},
		{"a_am", args{"a", []interface{}{parse("2010-10-10 10:10:10")}}, "am"},
		{"a_pm", args{"a", []interface{}{parse("2010-10-10 22:10:10")}}, "pm"},
		{"A_AM", args{"A", []interface{}{parse("2010-10-10 10:10:10")}}, "AM"},
		{"A_PM", args{"A", []interface{}{parse("2010-10-10 22:10:10")}}, "PM"},
		{"B", args{"B", []interface{}{parse("2010-10-10 10:10:10")}}, "465"},
		{"B_-", args{"B", []interface{}{parse("1800-10-10 10:10:10")}}, "465"},
		{"g", args{"g", []interface{}{parse("2010-10-10 22:10:10")}}, "10"},
		{"G", args{"G", []interface{}{parse("2010-10-10 22:10:10")}}, "22"},
		{"h", args{"h", []interface{}{parse("2010-10-10 01:10:10")}}, "01"},
		{"H", args{"H", []interface{}{parse("2010-10-10 01:10:10")}}, "01"},
		{"i", args{"i", []interface{}{parse("2010-10-10 01:10:10")}}, "10"},
		{"s", args{"s", []interface{}{parse("2010-10-10 01:10:11")}}, "11"},
		{"u", args{"u", []interface{}{parse("2010-10-10 01:10:11")}}, "000000"},
		{"v", args{"v", []interface{}{parse("2010-10-10 01:10:11")}}, "000"},
		{"I", args{"I", []interface{}{parse("2010-10-10 01:10:11")}}, "0"},
		{"Y-m-d H:i:sP", args{"Y-m-d H:i:sP", []interface{}{parse("2010-10-10 01:10:11")}}, "2010-10-10 01:10:11+00:00"},
		{"Y-m-d H:i:sp", args{"Y-m-d H:i:sp", []interface{}{parse("2010-10-10 01:10:11")}}, "2010-10-10 01:10:11Z"},
		{"T", args{"T", []interface{}{parse("2010-10-10 01:10:11")}}, "UTC"},
		{"e", args{"e", []interface{}{parse("2010-10-10 01:10:11")}}, "UTC"},
		{"Z", args{"Z", []interface{}{parse("2010-10-10 01:10:11")}}, "0"},
		{"c", args{"c", []interface{}{parse("2010-10-10 01:10:11")}}, "2010-10-10T01:10:11+00:00"},
		{"r", args{"r", []interface{}{parse("2010-10-10 01:10:11")}}, "Sun, 10 Oct 2010 01:10:11 +0000"},
		{"U", args{"U", []interface{}{1286557811}}, "1286557811"},
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
	time.Local = tmp
}
