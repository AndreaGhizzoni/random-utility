# Project Zenium
Zenium it'a a utility program to generate random data files.

## Introductions
* [Dependencies](##Dependencies)
* [Get this project](##Get this project)
* [Structure of Data](##Structure of random data)
    - [Random Slice](###Slice (aka Vectors))
    - [Random Matrix](###Matrix)
    
## Dependencies
 - [cli framework](https://github.com/urfave/cli)

## Get this project
Keep in mind that this project is still WORK IN PROGRESS, so changes will be
made.
```bash
cd $GOPATH/src
go get github.com/AndreaGhizzoni/zenium
./$GOPATH/bin/zenium -h
```

## Structure of random data
All generated data files will be in the following format:

### Slice (aka Vectors)
```bash
./$GOPATH/bin/zenium rslice -l 50 -o vector.txt -m 1 -M 10
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


### Matrix
```bash
./$GOPATH/bin/zenium matrix -o matrix.txt -c 10 -r 10 -m 1 -M 10
```
This command generate a file called `matrix.txt` that contains a `10x10` matrix
of random numbers, from `1` to `10` using current time as seed to generate it.

The file format is the follow:
```
m n
x11 [...] x1n
x21
[...]
xm1 [...] xmn
```
Where `m n` is the given size of matrix (10 10 in the example) and listed below
that all the elements of matrix: columns separated by space and rows by new line.

# TODO
- [x] Slice (aka Vector)
- [x] Matrix
- [x] Complete -output|-o flag
- [ ] Ordered Slice
- [ ] Bound
- [ ] Multiple generation with a single command