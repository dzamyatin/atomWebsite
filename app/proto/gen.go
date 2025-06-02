package proto

////go:generate sh -c "find ./ -name '*.proto' | xargs protoc --go_out=../internal/grpc/generated/ --go_opt=paths=source_relative --go-grpc_out=../internal/grpc/generated/ --go-grpc_opt=paths=source_relative"
////go:generate sh -c "protoc --plugin=protoc-gen-ts=../../frontend/node_modules/.bin/protoc-gen-ts --js_out=import_style=commonjs,binary:../../frontend/src/gen --ts_out=service=grpc-web:../../frontend/src/gen -I ./ ./auth.proto"
////go:generate sh -c "protoc --plugin=protoc-gen-ts=../../frontend/node_modules/.bin/protoc-gen-ts --js_out=import_style=commonjs,binary:../../frontend/src/gen --ts_out=service=grpc-web:../../frontend/src/gen -I ./ ./auth.proto"
////go:generate sh -c "protoc --plugin=protoc-gen-ts=../../frontend/node_modules/.bin/protoc-gen-ts --js_out=import_style=commonjs,binary:../../frontend/src/gen --ts_out=service=grpc-web:../../frontend/src/gen -I ./ ./auth.proto"

////go:generate sh -c "protoc --plugin=protoc-gen-ts=../../frontend/node_modules/.bin/protoc-gen-ts --js_out=import_style=commonjs,binary:../../frontend/src/gen --ts_out=service=grpc-web:../../frontend/src/gen -I ./ ./auth.proto"
