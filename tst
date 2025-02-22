protoc --plugin=protoc-gen-ts=../../frontend/node_modules/.bin/protoc-gen-ts \
 --js_out=import_style=commonjs,binary:../../frontend/src/gen \
 --ts_out=service=grpc-web:../../frontend/src/gen \
 -I ./ ./auth.proto


 protoc --plugin=protoc-gen-ts=../../frontend/node_modules/.bin/protoc-gen-ts \
  --js_out=import_style=es6,binary:../../frontend/src/gen \
  --ts_out=service=grpc-web:../../frontend/src/gen \
  -I ./ ./auth.proto


   "dependencies": {
      "@grpc/grpc-js": "^1.13.3",
      "@improbable-eng/grpc-web": "^0.15.0",
      "@intlify/unplugin-vue-i18n": "^6.0.8",
      "buefy": "npm:@ntohq/buefy-next@^0.2.0",
      "google-protobuf": "^3.21.4",
      "grpc": "^1.24.11",
      "grpc-web-client": "^0.7.0",
      "pinia": "^3.0.1",
      "protoc-gen-ts": "^0.8.7",
      "vite-plugin-commonjs": "^0.10.4",
      "vue": "^3.5.13",
      "vue-i18n": "^9.14.4",
      "vue-router": "^4.5.0"
    },
    "devDependencies": {
      "@vitejs/plugin-vue": "^5.2.3",
      "browserify": "^16.2.2",
      "prettier": "3.5.3",
      "protobufjs-cli": "^1.1.3",
      "typescript": "^5.8.3",
      "vite": "^6.3.3",
      "vite-plugin-vue-devtools": "^7.7.2",
      "webpack": "^4.16.5",
      "webpack-cli": "^3.1.0"
    },