module main

go 1.20

replace modules-and-tests/hello => ./hello

replace modules-and-tests/sum => ./sum

require (
	modules-and-tests/hello v0.0.0-00010101000000-000000000000
	modules-and-tests/sum v0.0.0-00010101000000-000000000000
)
