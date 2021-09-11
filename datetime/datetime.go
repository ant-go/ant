package datetime

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
	"unsafe"

	"github.com/spf13/cast"
)

/*
#cgo CFLAGS: -I./timelib
#cgo LDFLAGS: -L${SRCDIR}/timelib -lastro -linterval -lparsedate -lparseisointervals -lparsetz -ltimelib -ltm2unixtime -lunixtime2tm -ldow -lm
#include "time.h"
#include "timelib.h"
timelib_tzinfo *_date_parse_tzfile_wrapper(const char *formal_tzname, const timelib_tzdb *tzdb, int *dummy_error_code)
{
	int _dummy_error_code;
	return timelib_parse_tzfile(formal_tzname, tzdb, &_dummy_error_code);
}

long php_strtotime(const char *tz_name, const char *s, unsigned long long _len, signed long long timestamp)
{
	size_t len = (size_t) _len;
	int parse_error, epoch_does_not_fit_in_zend_long, dummy_error_code;
	const timelib_tzdb *tzdb;
	timelib_error_container *error;
	timelib_time *t, *now;
	timelib_tzinfo *tzi;
	timelib_sll ts;

	tzdb = timelib_builtin_db();
	tzi = timelib_parse_tzfile(tz_name, tzdb, &dummy_error_code);
	if (tzi == NULL) {
		return -1;
	}

	now = timelib_time_ctor();
	now->tz_info = tzi;
	now->zone_type = TIMELIB_ZONETYPE_ID;
	timelib_unixtime2local(now, (timelib_sll) timestamp);

	t = timelib_strtotime(s, len, &error, tzdb, _date_parse_tzfile_wrapper);
	parse_error = error->error_count;
	timelib_error_container_dtor(error);
	if (parse_error) {
		timelib_time_dtor(now);
		timelib_time_dtor(t);
		return -1;
	}

	timelib_fill_holes(t, now, TIMELIB_NO_CLONE);
	timelib_update_ts(t, tzi);
	ts = timelib_date_to_int(t, &epoch_does_not_fit_in_zend_long);

	timelib_time_dtor(now);
	timelib_time_dtor(t);

	if (epoch_does_not_fit_in_zend_long) {
		return -1;
	}

	return (long) ts;
}
*/
import "C"

// Time return current Unix timestamp (second, s)
//
// function time(): int
// https://www.php.net/manual/en/function.time.php
// https://www.php.net/manual/zh/function.time.php
func Time() int { return int(time.Now().Unix()) }

// TimeMs return current Unix timestamp (millisecond, ms)
func TimeMs() int { return int(time.Now().UnixNano() / 1000 / 1000) }

// TimeUs return current Unix timestamp(microsecond, us)
func TimeUs() int { return int(time.Now().UnixNano() / 1000) }

// TimeNs return current Unix timestamp(nanosecond, ns)
func TimeNs() int { return int(time.Now().UnixNano()) }

// StrToTime Parse about any English textual datetime description into a Unix timestamp
//
// function strtotime(string $datetime, ?int $baseTimestamp = null): int|false
// https://www.php.net/manual/en/function.strtotime
// https://www.php.net/manual/zh/function.strtotime
func StrToTime(s string, args ...interface{}) (result int) {
	if s == "" {
		return -1
	}
	tz := os.Getenv("TZ")
	if tz == "" {
		tz = "UTC"
	}
	timezone := C.CString(tz)
	input := C.CString(s)
	defer func() {
		C.free(unsafe.Pointer(input))
		C.free(unsafe.Pointer(timezone))
	}()
	length := C.ulonglong(len(s))
	var _timestamp int
	if len(args) > 0 {
		if i, err := cast.ToIntE(args[0]); err != nil {
			_timestamp = Time()
		} else {
			_timestamp = i
		}
	} else {
		_timestamp = Time()
	}
	timestamp := C.longlong(_timestamp)
	result = int(C.php_strtotime(timezone, input, length, timestamp))
	return
}

// Date Format a local time/date
//
// function date(string $format, ?int $timestamp = null): string
// https://www.php.net/manual/en/function.date
// https://www.php.net/manual/zh/function.date
func Date(format string, args ...interface{}) string {
	var t time.Time
	if len(args) > 0 {
		if i, err := cast.ToInt64E(args[0]); err != nil {
			t = time.Now()
		} else {
			t = time.Unix(i, 0)
		}
	} else {
		t = time.Now()
	}
	return dateFormat(format, true, t)
}

// GMDate Format a GMT/UTC date/time
//
// function gmdate(string $format, ?int $timestamp = null): string
// https://www.php.net/manual/en/function.gmdate
// https://www.php.net/manual/zh/function.gmdate
func GMDate(format string, args ...interface{}) string {
	var t time.Time
	if len(args) > 0 {
		if i, err := cast.ToInt64E(args[0]); err != nil {
			t = time.Now()
		} else {
			t = time.Unix(i, 0)
		}
	} else {
		t = time.Now()
	}
	t = t.In(time.UTC)
	return dateFormat(format, false, t)
}

// ---------------------------------------------------------------------------------------------------------------------

var (
	monFullNames = []string{
		"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December",
	}
	monShortNames = []string{
		"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
	}
	dayFullNames = []string{
		"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday",
	}
	dayShortNames = []string{
		"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat",
	}

	dTableCommon  = [13]int{0, 0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}
	dTableLeap    = [13]int{0, 0, 31, 60, 91, 121, 152, 182, 213, 244, 274, 305, 335}
	mlTableCommon = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	mlTableLeap   = [13]int{0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
)

func dateFormat(format string, localtime bool, t time.Time) string {
	var (
		rfcColon bool
		buf      strings.Builder
	)
	for _, char := range format {
		switch char {
		// day
		case 'd':
			buf.WriteString(fmt.Sprintf("%02d", t.Day()))
		case 'D':
			buf.WriteString(dayShortNames[t.Weekday()])
		case 'j':
			buf.WriteString(fmt.Sprintf("%d", t.Day()))
		case 'l':
			buf.WriteString(dayFullNames[t.Weekday()])
		case 'S':
			buf.WriteString(englishSuffix(t.Day()))
		case 'w':
			buf.WriteString(fmt.Sprintf("%d", t.Weekday()))
		case 'N':
			w := int(t.Weekday())
			if w == 0 {
				w = 7
			}
			buf.WriteString(fmt.Sprintf("%d", w))
		case 'z':
			buf.WriteString(fmt.Sprintf("%d", dayOfYear(t.Year(), int(t.Month()), t.Day())))
		// week
		case 'W':
			_, w := t.ISOWeek()
			buf.WriteString(fmt.Sprintf("%02d", w))
		case 'o':
			y, _ := t.ISOWeek()
			buf.WriteString(fmt.Sprintf("%d", y))
		// month
		case 'F':
			buf.WriteString(monFullNames[t.Month()-1])
		case 'm':
			buf.WriteString(fmt.Sprintf("%02d", int(t.Month())))
		case 'M':
			buf.WriteString(monShortNames[t.Month()-1])
		case 'n':
			buf.WriteString(fmt.Sprintf("%d", int(t.Month())))
		case 't':
			buf.WriteString(fmt.Sprintf("%d", daysInMonth(t.Year(), int(t.Month()))))
		// year
		case 'L':
			i := 0
			if isLeap(t.Year()) {
				i = 1
			}
 			buf.WriteString(fmt.Sprintf("%d", i))
		case 'y':
			buf.WriteString(fmt.Sprintf("%02d", t.Year()%100))
		case 'Y':
			y := t.Year()
			var s string
			if y < 0 {
				s = "-"
			}
			buf.WriteString(fmt.Sprintf("%s%04d", s, int(math.Abs(float64(y)))))
		// time
		case 'a':
			h := "am"
			if t.Hour() >= 12 {
				h = "pm"
			}
			buf.WriteString(h)
		case 'A':
			h := "AM"
			if t.Hour() >= 12 {
				h = "PM"
			}
			buf.WriteString(h)
		case 'B':
			sse := t.Unix()
			i := (sse - (sse - ((sse % 86400) + 3600))) * 10
			if i < 0 {
				i += 864000
			}
			i = (i / 864) % 1000
			buf.WriteString(fmt.Sprintf("%03d", i))
		case 'g':
			h, v := t.Hour(), 12
			if h%12 != 0 {
				v = h % 12
			}
			buf.WriteString(fmt.Sprintf("%d", v))
		case 'G':
			buf.WriteString(fmt.Sprintf("%d", t.Hour()))
		case 'h':
			h, v := t.Hour(), 12
			if h%12 != 0 {
				v = h % 12
			}
			buf.WriteString(fmt.Sprintf("%02d", v))
		case 'H':
			buf.WriteString(fmt.Sprintf("%02d", t.Hour()))
		case 'i':
			buf.WriteString(fmt.Sprintf("%02d", t.Minute()))
		case 's':
			buf.WriteString(fmt.Sprintf("%02d", t.Second()))
		case 'u':
			buf.WriteString(fmt.Sprintf("%06d", t.Nanosecond()/1000))
		case 'v':
			buf.WriteString(fmt.Sprintf("%03d", t.Nanosecond()/1000/1000))
		// timezone
		case 'I':
			i := 0
			if isTimeDST(t) {
				i = 1
			}
			buf.WriteString(fmt.Sprintf("%d", i))
		case 'p':
			zone, offset := t.Zone()
			if !localtime || zone == "UTC" || offset == 0 {
				buf.WriteString("Z")
				continue
			}
			// break intentionally missing
			fallthrough
		case 'P':
			rfcColon = true
			fallthrough
		case 'O':
			_, offset := t.Zone()
			c := "+"
			d1, d2, s := 0, 0, ""
			if localtime {
				if offset < 0 {
					c = "-"
				}
				d1 = int(math.Abs(float64(offset) / 3600.0))
				if rfcColon {
					s = ":"
				}
				d2 = int(math.Abs(float64(offset % 3600.0 / 60.0)))
			}
			buf.WriteString(fmt.Sprintf("%s%02d%s%02d", c, d1, s, d2))
		case 'T':
			zone, _ := t.Zone()
			if !localtime {
				zone = "GMT"
			}
			buf.WriteString(zone)
		case 'e':
			if !localtime {
				buf.WriteString("UTC")
			} else {
				zone, _ := t.Zone()
				buf.WriteString(zone)
			}
		case 'Z':
			_, offset := t.Zone()
			if !localtime {
				offset = 0
			}
			buf.WriteString(fmt.Sprintf("%d", offset))
		// full date/time
		case 'c':
			_, offset := t.Zone()
			s, d1, d2 := "+", 0, 0
			if localtime {
				if offset < 0 {
					s = "-"
				}
				d1 = int(math.Abs(float64(offset) / 3600.0))
				d2 = int(math.Abs(float64(offset % 3600.0 / 60.0)))
			}
			buf.WriteString(fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02d%s%02d:%02d",
				t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), s, d1, d2))
		case 'r':
			_, offset := t.Zone()
			s, d1, d2 := "+", 0, 0
			if localtime {
				if offset < 0 {
					s = "-"
				}
				d1 = int(math.Abs(float64(offset) / 3600.0))
				d2 = int(math.Abs(float64(offset % 3600.0 / 60.0)))
			}
			buf.WriteString(fmt.Sprintf("%3s, %02d %3s %04d %02d:%02d:%02d %s%02d%02d",
				dayShortNames[t.Weekday()],
				t.Day(), monShortNames[t.Month()-1],
				t.Year(), t.Hour(), t.Minute(), t.Second(), s, d1, d2))
		case 'U':
			buf.WriteString(fmt.Sprintf("%d", t.Unix()))
		case '\\':
			fallthrough
		default:
			buf.WriteRune(char)
		}
	}
	return buf.String()
}

func englishSuffix(number int) string {
	if number >= 10 && number <= 19 {
		return "th"
	}
	switch number % 10 {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}
}

func isLeap(y int) bool {
	return y%4 == 0 && (y%100 != 0 || y%400 == 0)
}

func dayOfYear(y, m, d int) int {
	if isLeap(y) {
		return dTableLeap[m] + d - 1
	}
	return dTableCommon[m] + d - 1
}

func daysInMonth(y, m int) int {
	if isLeap(y) {
		return mlTableLeap[m]
	}
	return mlTableCommon[m]
}

// isTimeDST returns true if time t occurs within daylight saving time
// for its time zone.
// https://stackoverflow.com/questions/53046636/how-to-check-whether-current-local-time-is-dst
func isTimeDST(t time.Time) bool {
	// If the most recent (within the last year) clock change
	// was forward then assume the change was from std to dst.
	hh, mm, _ := t.UTC().Clock()
	tClock := hh*60 + mm
	for m := -1; m > -12; m-- {
		// assume dst lasts for least one month
		hh, mm, _ := t.AddDate(0, m, 0).UTC().Clock()
		clock := hh*60 + mm
		if clock != tClock {
			if clock > tClock {
				// std to dst
				return true
			}
			// dst to std
			return false
		}
	}
	// assume no dst
	return false
}
