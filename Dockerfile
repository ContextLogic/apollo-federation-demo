From registry-gitlab.i.wish.com/contextlogic/protobuf-go-base-image/release:golang1.16.5-protobuf3.15.7r1-grpc_gateway1.16.0 as build

COPY ./services /go/src/github.com/ContextLogic/services

ENV GOPROXY=https://athens.i.wish.com
ENV GONOSUMDB=github.com/ContextLogic

WORKDIR /go/src/github.com/ContextLogic/services/accounts
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o ../bin/ github.com/ContextLogic/accounts

WORKDIR /go/src/github.com/ContextLogic/services/inventory
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o ../bin/ github.com/ContextLogic/inventory

WORKDIR /go/src/github.com/ContextLogic/services/products
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o ../bin/ github.com/ContextLogic/products

WORKDIR /go/src/github.com/ContextLogic/services/reviews
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o ../bin/ github.com/ContextLogic/reviews

FROM registry-gitlab.i.wish.com/contextlogic/tooling-image/alpine:3.14

COPY --from=build /go/src/github.com/ContextLogic/services/bin /bin/services

EXPOSE 4001
EXPOSE 4002
EXPOSE 4003
EXPOSE 4004
