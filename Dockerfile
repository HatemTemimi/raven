FROM golang:1.20.5-bullseye
 
WORKDIR /

COPY . .

RUN go get

RUN go build .

ENTRYPOINT [ "/Raven" ] 
