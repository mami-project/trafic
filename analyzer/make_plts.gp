#!/usr/bin/env gnuplot
set term postscript eps color
set key outside
set xlabel "Time slots = 0.1 s"
set ylabel "Video bytes"

set output "abr-75.eps"
set title "lola-baseline-75 (ABR)"
plot for [col=3:4] "<./analyse-lola lola-baseline-75-1528175109" using 1:col with points title columnheader

set output "abr-80.eps"
set title "lola-baseline-80 (ABR)"
plot for [col=3:4] "<./analyse-lola lola-baseline-80-1528175180" using 1:col with points title columnheader

set output "abr-85.eps"
set title "lola-baseline-85 (ABR)"
plot for [col=3:4] "<./analyse-lola lola-baseline-85-1528175250" using 1:col with points title columnheader

set output "abr-90.eps"
set title "lola-baseline-90 (ABR)"
plot for [col=3:4] "<./analyse-lola lola-baseline-90-1528175321" using 1:col with points title columnheader

set output "abr-95.eps"
set title "lola-baseline-95 (ABR)"
plot for [col=3:4] "<./analyse-lola lola-baseline-95-1528175392" using 1:col with points title columnheader
