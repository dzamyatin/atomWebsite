import * as Auth from "../gen/client/src/index"
import {useLoginStore} from "@/stores/login";

const host = "http://localhost:8503"

function getAuthClient() {
    const store = useLoginStore();
    let client = new Auth.ApiClient(host)
    client.timeout = 60000;
    client.defaultHeaders = {
        'Authorization': 'Bearer ' + store.jwt,
        'User-Agent': 'OpenAPI-Generator/version not set/Javascript',
    }
    let auth = new Auth.AuthApi(client)
    return auth
}

function resolveFn(resolve) {
    return function (error, data, response) {
        resolve(new Result(
            error,
            data,
            response,
        ))
    }
}

export async function confirmEmail(
    email,
    code,
) {
    let promise = new Promise((resolve, reject) => {
        getAuthClient().authConfirmEmail(
            {
                'email': email,
                'code': code,
            },
            resolveFn(resolve)
        )
    });

    return promise
}

export async function sendEmailConfirmation(
    email,
) {
    let promise = new Promise((resolve, reject) => {
        getAuthClient().authSendEmailConfirmation(
            {
                'email': email,
            },
            resolveFn(resolve)
        )
    });

    return promise
}

export async function confirmPhone(
    phone,
    code,
) {
    let promise = new Promise((resolve, reject) => {
        getAuthClient().authConfirmPhone(
            {
                'phone': phone,
                'code': code,
            },
            resolveFn(resolve)
        )
    });

    return promise
}

export async function sendPhoneConfirmation(
    phone,
) {
    let promise = new Promise((resolve, reject) => {
        getAuthClient().authSendPhoneConfirmation(
            {
                'phone': phone,
            },
            resolveFn(resolve)
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
        getAuthClient().authRegister(
            {
                'email': email,
                'password': password,
                'phone': phone,
            },
            resolveFn(resolve)
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
        getAuthClient().authLogin(
            {
                'email': email,
                'password': password,
                'phone': phone,
            },
            resolveFn(resolve)
        )
    });

    return promise
}

export async function rememberPassword(
    email,
    phone
) {
    let promise = new Promise((resolve, reject) => {
        getAuthClient().authRememberPassword(
            {
                'email': email,
                'phone': phone,
            },
            resolveFn(resolve)
        )
    });

    return promise
}

export async function changePassword(
    newPassword,
    oldPassword,
    email,
    phone,
    code
) {
    let promise = new Promise((resolve, reject) => {
        getAuthClient().authChangePassword(
            {
                'newPassword': newPassword,
                'oldPassword': oldPassword,
                'email': email,
                'phone': phone,
                'code': code,
            },
            resolveFn(resolve)
        )
    });

    return promise
}

export async function getCurrentUser() {
    let promise = new Promise((resolve, reject) => {
        getAuthClient().authCurrent(
            resolveFn(resolve)
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
