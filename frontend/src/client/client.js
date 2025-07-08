import * as Auth from "../gen/client/src/index"

const host = "http://localhost:8503"

let client = new Auth.ApiClient(host)
client.timeout = 60000;

let auth = new Auth.AuthApi(client)

export async function confirmEmail(
    email,
    code,
) {
    let promise = new Promise((resolve, reject) => {
        auth.authConfirmEmail(
            {
                'email': email,
                'code': code,
            },
            function callback(error, data, response) {
                resolve(new Result(
                    error,
                    data,
                    response,
                ))
            }
        )
    });

    return promise
}

export async function register(
    email,
    password,
    phone
) {
    let promise = new Promise((resolve, reject) => {
        auth.authRegister(
            {
                'email': email,
                'password': password,
                'phone': phone,
            },
            function callback(error, data, response) {
                resolve(new Result(
                    error,
                    data,
                    response,
                ))
            }
        )
    });

    return promise
}

export async function login(
    email,
    password,
    phone
) {
    let promise = new Promise((resolve, reject) => {
        auth.authLogin(
            {
                'email': email,
                'password': password,
                'phone': phone,
            },
            function callback(error, data, response) {
                resolve(new Result(
                    error,
                    data,
                    response,
                ))
            }
        )
    });

    return promise
}

class Result {
    constructor(
        error,
        data,
        response,
    ) {
        this.error = error
        this.data = data
        this.response = response
    }
}
