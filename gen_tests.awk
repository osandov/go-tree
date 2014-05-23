/^func test(\w*)\(t \*testing.T, tree Tree\) {$/ {
	f = gensub(/^func test(\w*)\(t \*testing.T, tree Tree\) {$/, "\\1", "", $0)
	tests[f] = 1
}

/^func benchmark(\w*)\(b \*testing.B, tree Tree\) {$/ {
	f = gensub(/^func benchmark(\w*)\(b \*testing.B, tree Tree\) {$/, "\\1", "", $0)
	benchmarks[f] = 1
}

{
	split($0, a, ": ")
	if (a[1] == "// TEST") {
		printf "// %s.\n", a[2]
		for (test in tests) {
			printf "func Test%s%s(t *testing.T) {\n", a[3], test
			printf "\ttest%s(t, %s)\n", test, a[4]
			printf "}\n"
		}
		for (benchmark in benchmarks) {
			printf "func Benchmark%s%s(b *testing.B) {\n", a[3], benchmark
			printf "\tbenchmark%s(b, %s)\n", benchmark, a[4]
			printf "}\n"
		}
	} else {
		print
	}
}
