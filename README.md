# siege-redis

<p align="center">
	<img src="https://camo.githubusercontent.com/5b13bf8be0d98cf8e2764ce07ca68ee02a273f63/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f676f6c616e672d312e31332d626c75652e7376673f7374796c653d666c6174">
	<a href="https://raw.githubusercontent.com/onevcat/Kingfisher/master/LICENSE"><img src="https://img.shields.io/cocoapods/l/Kingfisher.svg?style=flat"></a>
</p>

siege-redis is a tiny tool to test redis.

## Installation

```golang
go get github.com/Zfzlv/siege-redis

cd siege-redis

go build

```

## Usage

Example

./siege-redis -host 10.0.18.183 -port 6379 -a icR3TZHiL -m 1 -pool 1 -t 1m -w -o

It is as simple as doing this

## Help

./siege-redis -h

or

./siege-redis --help

## Thanks

[echosimple](https://www.zcool.com.cn/u/15027113)

[uuid](github.com/satori/go.uuid/)
[go-redis](github.com/go-redis/redis)
[progressbar](github.com/schollz/progressbar)

## License
[MIT](https://choosealicense.com/licenses/mit/)
