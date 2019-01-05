FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o main .
HEALTHCHECK CMD curl --fail http://localhost:1234/ || exit 1
CMD ["/app/main"]
