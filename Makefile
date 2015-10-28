
all: public private
.PHONY: all

prep:
	rm -rf $GOPATH/pkg
	mkdir -p bin
	go get ./...

public: prep
	go build -o=./bin/bqpublic public/main.go

private: prep
	go build -o=./bin/bqprivate private/main.go

test: all
	go test -v ./...
	#bin/bqpublic --request_port=20000 --reply_port=30000 --http_port=9099 --gen_ssl &
	bin/bqpublic --request_port=20000 --reply_port=30000 --http_port=9099 &
	bin/bqprivate --request_port=20000 --reply_port=30000 --public_host=localhost &
	sleep 1
	#curl -vvv -k -H 'Host: google.com' https://127.0.0.1:9099/webhp?q=golang
	curl -vvv -k -H 'Host: google.com' http://127.0.0.1:9099/webhp?q=golang
	sleep 2
	killall bqprivate
	killall bqpublic
	echo Success

clean:
	rm -rf ./bin/
