FROM golang:1.13rc2-buster as builder

RUN apt-get update
RUN apt-get install -y imagemagick
RUN apt-get install -y ffmpeg
RUN apt-get install -y youtube-dl

WORKDIR /usr/local/bin
ENV PATH /usr/local/bin:$PATH
COPY ./docker_scripts/docker-entrypoint.sh .
RUN chmod +x docker-entrypoint.sh

WORKDIR /go/src/hk5demandsapi

ENTRYPOINT ["docker-entrypoint.sh"]