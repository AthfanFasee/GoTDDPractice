SHELL := powershell.exe
.SHELLFLAGS := -NoProfile -Command

test:
	go1.20 test
testC:
	go1.20 test -cover
testB:
	go1.20 test -bench="."