FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o webservice .
EXPOSE 8000
CMD ["/app/webservice"]

