# Specifies a parent image
FROM golang:1.20.5-bullseye
 
WORKDIR /app
COPY go.mod ./
RUN go mod download
RUN go get github.com/HatemTemimi/Raven/raven
COPY *.go ./
RUN go build -o /raven
CMD [ “/raven” ]
