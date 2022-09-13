FROM golang:alpine

ENV GIN_MODE=release
ENV PORT=3004

WORKDIR /usr/src/app

COPY . /usr/src/app

RUN go build /usr/src/app

EXPOSE $PORT

ENTRYPOINT ["./Static-Site-Generator"]
