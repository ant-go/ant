timelibDir = ./datetime/timelib
timelib:
	gcc -c -o $(timelibDir)/astro.o               $(timelibDir)/astro.c
	gcc -c -o $(timelibDir)/dow.o                 $(timelibDir)/dow.c
	gcc -c -o $(timelibDir)/interval.o            $(timelibDir)/interval.c
	gcc -c -o $(timelibDir)/parse_date.o          $(timelibDir)/parse_date.c
	gcc -c -o $(timelibDir)/parse_iso_intervals.o $(timelibDir)/parse_iso_intervals.c
	gcc -c -o $(timelibDir)/parse_tz.o            $(timelibDir)/parse_tz.c
	gcc -c -o $(timelibDir)/timelib.o             $(timelibDir)/timelib.c
	gcc -c -o $(timelibDir)/tm2unixtime.o         $(timelibDir)/tm2unixtime.c
	gcc -c -o $(timelibDir)/unixtime2tm.o         $(timelibDir)/unixtime2tm.c
	ar rcs $(timelibDir)/libastro.a               $(timelibDir)/astro.o
	ar rcs $(timelibDir)/libdow.a                 $(timelibDir)/dow.o
	ar rcs $(timelibDir)/libinterval.a            $(timelibDir)/interval.o
	ar rcs $(timelibDir)/libparsedate.a           $(timelibDir)/parse_date.o
	ar rcs $(timelibDir)/libparseisointervals.a   $(timelibDir)/parse_iso_intervals.o
	ar rcs $(timelibDir)/libparsetz.a             $(timelibDir)/parse_tz.o
	ar rcs $(timelibDir)/libtimelib.a             $(timelibDir)/timelib.o
	ar rcs $(timelibDir)/libtm2unixtime.a         $(timelibDir)/tm2unixtime.o
	ar rcs $(timelibDir)/libunixtime2tm.a         $(timelibDir)/unixtime2tm.o

test:
	make timelib
	sh coverage.sh