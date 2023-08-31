import DateAgo from "src/atoms/DateAgo/DateAgo";
import StatusIcon from "src/atoms/StatusIcon/StatusIcon";
import Typography from "src/atoms/Typography/Typography";
import useTask from "src/hooks/useTask";
import { Task } from "src/services/taskService";

export default function TaskRow({ task }: { task: Task }) {
    const { Title, Description, DueDate, CreatedAt, Status } = task
    const { edit } = useTask()

    const editTask = () => {
        edit(task.ID)
    }

    return (
        <div onClick={() => editTask()} className="p-6 shadow-md my-6 border border-slate-200 dark:border-slate-600 rounded-lg cursor-pointer hover:bg-base-200 ease-in bg-base-100 dark:bg-base-200 justify-center grid grid-cols-1 md:grid-cols-5 lg:grid-cols-6 items-center gap-4">
            <div className="text-center">
                <StatusIcon status={Status} />
            </div>
            <div className="grid grid-cols-1 lg:grid-cols-2 md:col-span-2 lg:col-span-3 gap-4">
                <Typography size="md">{Title}</Typography>
                <Typography truncate={true} size="sm">{Description}</Typography>
            </div>
            <Typography size="sm" className="flex">
                <span className="block md:hidden mr-2 italic">Due:</span><DateAgo date={DueDate}/>
            </Typography>
            <Typography size="sm" className="flex justify-end">
                <span className="block md:hidden mr-2 italic">Created on:</span><DateAgo date={CreatedAt}/>
            </Typography>
        </div>

    );
}