import * as Auth from "../gen/client/src/index"

const host = "http://localhost:8503"

let client = new Auth.ApiClient(host)

let auth = new Auth.AuthApi(client)

auth.authRegister(
    {
        'email': 'testemaildz@ya.ru',
        'password': 'pwd123',
        'phone': '+79291112233',
    },
    function callback(error, data, response) {
        console.log('test')
        console.log(error)
        console.log(data)
        console.log(response)
    }
)
