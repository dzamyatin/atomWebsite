go get -tool github.com/google/wire/cmd/wire@v0.6.0
go get tool
go install tool
//go:generate go tool stringer

GOPRIVATE=*.corp.example.com,*.research.example.com


protobuf format  https://protobuf.dev/reference/go/go-generated/#package
streaming example https://grpc.io/docs/languages/go/basics/

# TODO:
- make grpc server
- make wire by gooogle for di
- make main.go common entrypoint of apps

#
    ports:
      - "${OUT_GRPCPORT}:${GRPCPORT}"
      - "${OUT_GRPCSSLPORT}:${GRPCSSLPORT}"
      - "${OUT_OPSPORT}:${OPSPORT}"
      - "${OUT_HTTP}:${HTTP}"
      - "${OUT_HTTPS}:${HTTPS}"