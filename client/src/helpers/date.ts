import moment from "moment";

export const formatDate = (dateString: string) => {
    if(dateString === ""){
        return;
    }

    const momentObj = moment(dateString, "YYYY-MM-DDThh:mm:ssZ");

    return momentObj.fromNow()
}

export const dateToString = (date: Date, time: string) => {
    const [hour, minute] = time.split(':')
    date.setHours(parseInt(hour))
    date.setMinutes(parseInt(minute))
    return date.toISOString()
}