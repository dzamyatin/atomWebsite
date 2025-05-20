
// import "google-protobuf"
// import * as goog from "goog"
import {grpc} from "@improbable-eng/grpc-web";

import * as module from  "../gen/auth_pb.js"

grpc.unary(RegisterRequest, {
    request: getBookRequest,
    host: host,
    onEnd: res => {
        const { status, statusMessage, headers, message, trailers } = res;
        if (status === grpc.Code.OK && message) {
            console.log("all ok. got book: ", message.toObject());
        }
    }
});