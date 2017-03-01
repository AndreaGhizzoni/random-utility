# Project Zenium
Zenium it'a a utility program to generate random data files in the format 
specifid below.


## Get this project
```bash
cd $GOPATH/src
go get github.com/AndreaGhizzoni/zenium
```

## Build application
```bash
go install github.com/AndreaGhizzoni/zenium
../bin/zenium
```

## Structure of data files
All generated data fils will be in the following format:

### Random Slice (aka Vectors)
```bash
./$GOPATH/bin/zenium -generate rslice -l 50 -o vector.txt -m 1 -M 10
```
This command generate a file called `vector.txt` that contains `50` random 
numbers, from `1` to `10` using current time as seed to generate it.

The file format is the follow:
```
n
x1 x2 [...] xn
```
Where `n` is the given length of vector (50 in the example) and `x1 x2 [...] xn`
are all the generated numbers (with the properties see previously in the 
example) separated by space.

# TODO
- [ ] Ordered Vector
- [ ] Matrix
- [ ] Bound