import {grpc} from "@improbable-eng/grpc-web";
import * as AuthService from "../../ts/_proto/auth_pb_service";
const Auth = AuthService.Auth;
import * as AuthPb from "../../ts/_proto/auth_pb";

const host = "http://localhost:8502";
// const host = "localhost:8502";

console.log("tst1")

function Register() {
    const req = new AuthPb.RegisterRequest();
    req.setEmail("daniil@dasdas.ru")
    req.setPhone("+79297145267")
    req.setPassword("hellowirld")

    console.log("tst2")

    console.log(grpc)
    grpc.unary(
        Auth.Register,
        {
            request: req,
            host: host,
            onEnd: res => {
                console.log("tst3")
                const {
                    status,
                    statusMessage,
                    headers,
                    message,
                    trailers
                } = res;
                console.log("Register.onEnd.status", status, statusMessage);
                console.log("Register.onEnd.headers", headers);
                if (status === grpc.Code.OK && message) {
                    console.log("Register.onEnd.message", message.toObject());
                }
                console.log("Register.onEnd.trailers", trailers);
            }
        }
    )
}

Register()
