tmp_dir := ./tmp
pkg_base := github.com/s-shin/spelunker
cover_out := $(tmp_dir)/cover.out

build:
	go build

run: build
	./spelunker i

net_protoc:
	protoc -I net/shogi net/shogi/shogi.proto --go_out=plugins=grpc:net/shogi

test_main:
	go test -v -cover $(pkg_base)

test_shogi:
	go test -v -cover $(pkg_base)/shogi

test_parsec:
	go test -v -cover $(pkg_base)/parsec

test_shogi_record_ki2:
	go test -v -cover $(pkg_base)/shogi/record/ki2

test_shogi_script:
	go test -v -cover $(pkg_base)/shogi/script

coverage:
	go test -coverprofile=$(cover_out) $(pkg_shogi)
	go tool cover -func=$(cover_out)
