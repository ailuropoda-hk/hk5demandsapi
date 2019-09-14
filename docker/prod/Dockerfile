FROM golang:1.13rc2-buster as builder

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/hk5demandsapi
COPY ./controllers ./controllers
COPY ./lib ./lib
COPY ./routers ./routers
COPY ./main.go .
COPY ./Gopkg.lock .
COPY ./Gopkg.toml .
RUN dep ensure
RUN go build

FROM golang:1.13rc2-buster

RUN apt-get update
RUN apt-get install -y imagemagick
RUN apt-get install -y ffmpeg
RUN apt-get install -y youtube-dl

WORKDIR /usr/app/hk5demandsapi
COPY --from=builder /go/src/hk5demandsapi/hk5demandsapi .

WORKDIR /usr/local/bin
ENV PATH /usr/local/bin:$PATH
COPY ./docker_scripts/docker-entrypoint.sh .
RUN chmod +x docker-entrypoint.sh

WORKDIR /usr/app/hk5demandsapi

ENTRYPOINT ["docker-entrypoint.sh"]