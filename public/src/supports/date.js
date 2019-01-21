
import moment from 'moment'

export default {
    humanDate(dateTime) {
        console.log(dateTime)
        return moment(dateTime).fromNow()
    }
}