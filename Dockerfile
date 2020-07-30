FROM golang

COPY . /hexa
WORKDIR /hexa

RUN go get -v all
RUN go build -o ./ ./cmd/hexa

EXPOSE 8000

CMD ["hexa"]