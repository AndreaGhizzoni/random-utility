# default option to make 
default : test

test: 
	go test ./... -test.v=true -bench=.
