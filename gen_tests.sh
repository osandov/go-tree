#!/bin/sh

gawk -f gen_tests.awk < external_test.templ > external_test.go
gawk -f gen_tests.awk < benchmark_test.templ > benchmark_test.go
