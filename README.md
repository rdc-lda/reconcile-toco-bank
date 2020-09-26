# Reconcile TOCOnline & Millennium BCP

A utility to reconcile records from TOC Online (Excel file) and Millennium BCP (loosely formatted CSV file...).

## Prerequisites

Next to Go, yo need to install [glide](https://github.com/Masterminds/glide) for package management.

~~~bash
# For OSX use brew...
$ brew install glide
~~~

## Build

~~~bash
# Fetch libs, compile & install binary
$ make all
~~~

## Run

~~~bash
# The help function speaks for itself
$ ./bin/rtbutil -h
~~~
