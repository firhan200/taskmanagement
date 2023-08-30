import Typography from "src/atoms/Typography/Typography";
import { formatDate } from "src/helpers/date";
import useTask from "src/hooks/useTask";
import { Task } from "src/services/taskService";

export default function TaskRow({ task }: { task: Task }) {
    const { Title, Description, DueDate, CreatedAt } = task
    const { edit } = useTask()

    const editTask = () => {
        edit(task.ID)
    }

    return (
        <div onClick={() => editTask()} className="p-6 shadow-md my-6 border border-slate-200 dark:border-slate-600 rounded-lg cursor-pointer hover:bg-base-200 ease-in bg-base-100 dark:bg-base-200 justify-center grid grid-cols-1 md:grid-cols-4 lg:grid-cols-4">
            <Typography size="md">{Title}</Typography>
            <Typography size="md">{Description}</Typography>
            <Typography size="sm" className="flex">
                <span className="block md:hidden mr-2 italic">Due:</span>{formatDate(DueDate)}
            </Typography>
            <Typography size="sm" className="flex justify-end">
                <span className="block md:hidden mr-2 italic">Created on:</span>{formatDate(CreatedAt)}
            </Typography>
        </div>
    );
}