# https://blog.docker.com/2016/09/docker-golang/
# https://blog.golang.org/docker

# docker build -t wof-mmdb-server .
# docker run -it -p 6161:8080 -e HOST='0.0.0.0' wof-mmdb-server

FROM golang

ADD . /go-whosonfirst-mmdb

RUN cd /go-whosonfirst-mmdb; make bin

EXPOSE 8080

ENTRYPOINT /go-whosonfirst-mmdb/docker/entrypoint.sh

