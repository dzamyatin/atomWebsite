import {grpc} from "@improbable-eng/grpc-web";
import {Auth} from "../../ts/_proto/auth_pb_service";
import {
    RegisterRequest,
    RegisterResponse,
    LoginRequest,
    LoginResponse,
} from "../../ts/_proto/auth_pb";

const host = "http://localhost:9090";

function Register() {
    const req = new RegisterRequest();
    req.setEmail("daniil@dasdas.ru")
    req.setPhone("+79297145267")
    req.setPassword("hellowirld")

    grpc.unary(
        Auth,
        {
            request: getBookRequest,
            host: host,
            onEnd: res => {
                const { status, statusMessage, headers, message, trailers } = res;
                console.log("getBook.onEnd.status", status, statusMessage);
                console.log("getBook.onEnd.headers", headers);
                if (status === grpc.Code.OK && message) {
                    console.log("getBook.onEnd.message", message.toObject());
                }
                console.log("getBook.onEnd.trailers", trailers);
            }
        }
    )
}
Register()