import moment from "moment";

const stringToMoment = (dateString: string) => {
    const momentObj = moment(dateString, "YYYY-MM-DDThh:mm:ssZ");
    return momentObj
}

export const formatDate = (dateString: string) => {
    if(dateString === ""){
        return;
    }

    const momentObj = stringToMoment(dateString);

    return momentObj.fromNow()
}

export const dateToReadableString = (dateString: string) => {
    const momentObj = stringToMoment(dateString);

    return momentObj.format("DD MMMM YYYY, HH:mm")
}

export const dateToString = (date: Date, time: string) => {
    const [hour, minute] = time.split(':')
    date.setHours(parseInt(hour))
    date.setMinutes(parseInt(minute))
    return date.toISOString()
}

export const dateStringToDateAndTime = (date: string) => {
    const momentObj = moment(date)

    const dateValue = momentObj.toDate()
    const timeValue = momentObj.format("HH:mm")

    return {
        dateValue,
        timeValue
    }
}