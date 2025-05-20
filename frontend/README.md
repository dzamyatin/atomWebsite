

https://v3.buefy.org/
https://github.com/ntohq/buefy-next
https://bulma.io/documentation/


# protogen
#sudo npm install -g protobufjs-cli
#npm install protobufjs-cli --save-dev

#pbjs -t static-module -w commonjs -o output.js example.proto
#pbjs -t static-module -w es6 -o output.js example.proto
#npm install google-protobuf



https://github.com/improbable-eng/grpc-web
https://github.com/improbable-eng/grpc-web/tree/master/client/grpc-web


https://github.com/improbable-eng/grpc-web/blob/master/client/grpc-web/docs/code-generation.md

> npm install protoc-gen-ts

```
protoc --plugin=protoc-gen-ts=../../frontend/node_modules/.bin/protoc-gen-ts --js_out=import_style=commonjs,binary:../../frontend/src/gen --ts_out=service=grpc-web:../../frontend/src/gen -I ./ ./auth.proto
```

