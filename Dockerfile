FROM golang:1.20 AS build

WORKDIR /sneakers-api

RUN apt-get update && apt install -y protobuf-compiler

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

RUN export PATH="$PATH:$(go env GOPATH)/bin"

COPY . .

RUN make compile-proto-go

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o sneakers-api  cmd/main.go


FROM scratch

COPY --from=build /sneakers-api .

EXPOSE 5000

CMD ["./sneakers-api"]







