
import http from 'axios'

export default {
    login(credentials, callback) {
        http.post("http://localhost:8000/api", credentials)
            .then((response) => {
                callback(response.data)
            })
    }
}