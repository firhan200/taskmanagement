import moment from "moment";

export const formatDate = (dateString: string) => {
    if(dateString === ""){
        return;
    }

    const momentObj = moment(dateString, "YYYY-MM-DDThh:mm:ssZ");

    return momentObj.fromNow()
}

export const dateToString = (date: string, time: string) => {
    const dateObj = moment(date)

    return dateObj.format("YYYY-MM-DDT")+time+":00+7.00"
}