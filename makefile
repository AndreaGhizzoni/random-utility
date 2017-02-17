install:
	go install

test: pkgRandutil pkgSampleGen pkgOut

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

pkgOut:
	cd out;\
	go test -test.v=true > ../out.test ;\
	go test -bench=. -benchmem > ../out.bench ;\
	cd ..

