import Typography from "src/atoms/Typography/Typography";
import { formatDate } from "src/helpers/date";
import { Task } from "src/services/taskService";

export default function TaskRow({ task }: { task: Task }) {
    const { Title, Description, DueDate, CreatedAt } = task
    return (
        <div className="p-6 shadow-md my-6 border border-slate-200 rounded-lg cursor-pointer hover:bg-base-200 ease-in bg-base-100 dark:bg-base-200 justify-center grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4">
            <Typography className="font-bold" size="md">{Title}</Typography>
            <Typography className="text-slate-500" size="md">{Description}</Typography>
            <Typography size="sm">{ formatDate(DueDate) }</Typography>
            <Typography size="sm">{ formatDate(CreatedAt) }</Typography>
        </div>
    );
}