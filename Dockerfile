FROM golang

COPY . /glofox
WORKDIR /glofox

RUN go get -v all
RUN go build -o ./ ./cmd/glofox

EXPOSE 8000

CMD ["glofox"]