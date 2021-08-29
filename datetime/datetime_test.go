package datetime

import "testing"

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
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty", args{""}, -1},
		{"UTC_2021-01-03", args{"2021-01-03"}, 1609603200},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrToTime(tt.args.s); got != tt.want {
				t.Errorf("StrToTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
