# Reconcile TOCOnline & Millennium BCP

A utility to reconcile records downloaded for the same period from TOC Online (XSLT) and Millennium BCP (CSV).

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
Usage: rtbutil [--tocsheet TOCSHEET] [--bcpsheet BCPSHEET]

Options:
  --tocsheet TOCSHEET, -t TOCSHEET
                         full path to TOC Online Excel file [default: ./testdata/extratocontas_514106654_25_09_2020.xlsx]
  --bcpsheet BCPSHEET, -b BCPSHEET
                         full path to Millennium BCP Excel file [default: ./testdata/EXTMV120825856552.xls]
  --help, -h             display this help and exit

# Example
$ rtbutil -t ./2020-06v2\ extratocontas_514106654_26_09_2020.xlsx \
   -b ./EXTMV120826185594.xls

These amounts were not found TOC online (compared to BCP)
+-----+--------+------------+------------------------------------------------+
| Ref | Amount | Date       | Description                                    |
+-----+--------+------------+------------------------------------------------+
| 1   | -59.99 | 2020-06-08 | MDB 857107 WORTEN 2724-520 AMA                 |
| 2   | 59.99  | 2020-06-10 | MDB 857107 WORTEN 2724-520 AMA                 |
| 3   | -30.00 | 2020-06-26 | DD PT34100781 EDP COMERCIAL C 16010003738930   |
| 4   | -90.00 | 2020-06-26 | DD PT34100781 EDP COMERCIAL C P1610003242953   |
| 5   | -5.00  | 2020-06-30 | COM.MANUTENCAO CONTA PACOTE FREQUENTE NEGOCIOS |
| 6   | -0.20  | 2020-06-30 | IMPOSTO SELO ART 17.3.4                        |
+-----+--------+------------+------------------------------------------------+



These amounts were not found in BCP Millennium (compared to TOC)
+-----+---------+------------+------------------------------------+
| Ref | Amount  | Date       | Description                        |
+-----+---------+------------+------------------------------------+
| 1   | -120.00 | 2020-06-26 | Pagamento a fornecedor PF 2020/150 |
+-----+---------+------------+------------------------------------+
~~~
