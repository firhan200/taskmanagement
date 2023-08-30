import { dateToReadableString, formatDate } from "src/helpers/date";

export default function DateAgo({ date }: { date: string }){
    return (
        <div className="tooltip" data-tip={dateToReadableString(date)}>
            { formatDate(date) }
        </div>
    );
}