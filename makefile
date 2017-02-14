# default option to make 
default : test

#go test ./... -test.v=true -bench=.
test: pkgRandutil pkgSampleGen

pkgRandutil:
	cd randutil; \
	go test -test.v=true > ../randutil.test; \
	go test -bench=. > ../randutil.bench ;\
	cd ..

pkgSampleGen:
	cd randutil; \
	go test -test.v=true > ../randutil.test; \
	go test -bench=. > ../randutil.bench ;\
	cd ..


