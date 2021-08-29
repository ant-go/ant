package datetime

import (
	"os"
	"time"
	"unsafe"
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

long php_strtotime(const char *tz_name, const char *s, unsigned long long _len)
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

	now = timelib_time_ctor();
	now->tz_info = tzi;
	now->zone_type = TIMELIB_ZONETYPE_ID;
	timelib_unixtime2local(now, (timelib_sll) time(NULL));

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
func Time() int64 { return time.Now().Unix() }

// TimeMs return current Unix timestamp (millisecond, ms)
func TimeMs() int64 { return time.Now().UnixNano() / 1000 / 1000 }

// TimeUs return current Unix timestamp(microsecond, us)
func TimeUs() int64 { return time.Now().UnixNano() / 1000 }

// TimeNs return current Unix timestamp(nanosecond, ns)
func TimeNs() int64 { return time.Now().UnixNano() }

func StrToTime(s string) (result int) {
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
	result = int(C.php_strtotime(timezone, input, length))
	return
}
