go get -tool github.com/google/wire/cmd/wire@v0.6.0
go get tool
go install tool
//go:generate go tool stringer

GOPRIVATE=*.corp.example.com,*.research.example.com