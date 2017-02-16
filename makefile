install:
	go install

test: pkgRandutil pkgSampleGen

pkgRandutil:
	cd randutil ;\
	go test -test.v=true > ../randutil.test ;\
	go test -bench=. -benchmem > ../randutil.bench ;\
	cd ..

pkgSampleGen:
	cd samplegen ;\
	go test -test.v=true > ../samplegen.test ;\
	go test -bench=. -benchmem > ../samplegen.bench ;\
	cd ..


