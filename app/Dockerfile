FROM golang:1.24 AS builder

ARG USERID=1000
ARG USERNAME="appuser"

RUN apt update
RUN apt install -y unzip

RUN useradd -u $USERID -m -U $USERNAME

USER $USERNAME

#install protoc
RUN wget -O protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v30.0-rc1/protoc-30.0-rc-1-linux-x86_64.zip && \
    unzip protoc.zip -d $HOME/protoc && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.5 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

ENV PATH="/home/${USERNAME}/protoc/bin:${PATH}"

WORKDIR /app

COPY --chown=$USERNAME:$USERNAME . .

RUN go get && \
    go build -o exec/app

FROM alpine:3.21.3 AS production

ARG GRPCPORT=8502
ARG GRPCSSLPORT=8503
ARG OPSPORT=8001
ARG HTTP=80
ARG HTTPS=443

EXPOSE $GRPCPORT
EXPOSE $GRPCSSLPORT
EXPOSE $OPSPORT
EXPOSE $HTTP
EXPOSE $HTTPS

RUN apk add libc6-compat

WORKDIR /app

COPY --chmod=110 --from=builder /app/exec/. /app

CMD /app/app