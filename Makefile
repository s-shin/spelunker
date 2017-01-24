tmp_dir := ./tmp
pkg_shogi := github.com/s-shin/spelunker/shogi
cover_out := $(tmp_dir)/cover.out

build:
	go build

run: build
	./spelunker

test:
	go test -v -cover $(pkg_shogi)

coverage:
	go test -coverprofile=$(cover_out) $(pkg_shogi)
	go tool cover -func=$(cover_out)

test_parsec:
	go test -v -cover github.com/s-shin/spelunker/parsec

test_shogi_record_ki2:
	go test -v -cover github.com/s-shin/spelunker/shogi/record/ki2
