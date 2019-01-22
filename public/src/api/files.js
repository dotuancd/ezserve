
import axios from 'axios'

export default {
    files(callback, options) {
        options = options || {};

        let params = {
            page: options.page || 1,
            per_page: options.perPage || 2
        }

        axios
            .get('http://localhost:8000/api/files?token=ebE3GB8Xsn3A5WOQ', {params})
            .then((response) => {
                callback(response.data)
            })
    },
    create(file, callback) {

        axios.post("http://localhost:8000/api/files?token=ebE3GB8Xsn3A5WOQ", file)
            .then((response) => {
                callback(response.data)
            })

    }
}