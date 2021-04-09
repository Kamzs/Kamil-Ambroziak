FROM golang:1.15.7-buster
RUN go get -u github.com/Kamzs/Kamil-Ambroziak
WORKDIR /app
RUN cd /go/src/github.com/Kamzs/Kamil-Ambroziak/main && go build -o app && mv app /app
EXPOSE 8080
CMD /app/app
