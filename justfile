target := "gigfun"

default: build

build:
    cd cmd && go build -o ../{{target}}
    strip -s {{target}}

clean:
    rm -f {{target}}

install destdir: build
    install -D -m 0755 {{target}} "{{destdir}}"/usr/bin/{{target}}
