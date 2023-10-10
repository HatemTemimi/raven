FROM golang:1.21-alpine3.18
 
WORKDIR /

COPY . .

RUN go get

RUN go build .

ENTRYPOINT [ "/Raven" ] 
