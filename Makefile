
all: public private

prep:
	rm -rf $GOPATH/pkg
	mkdir -p bin
	#go get ./...

public: prep
	go build -o=./bin/public public/main.go

private: prep
	go build -o=./bin/private private/main.go

test: all
	go test -v ./...
	bin/public &
	bin/private &
	sleep 1
	curl -XPOST -d"TEST" -H"Accept:application/json" http://localhost:9099/
	sleep 2
	killall private
	killall public
	echo Success

clean:
	rm -rf ./bin/