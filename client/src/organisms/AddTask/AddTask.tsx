import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useEffect, useState } from "react";
import DatePicker from "react-date-picker";
import TimePicker from "react-time-picker";
import Button from "src/atoms/Button/Button";
import Loading from "src/atoms/Loading/Loading";
import Typography from "src/atoms/Typography/Typography";
import { dateToString } from "src/helpers/date";
import FormAreaControl from "src/molecules/FormAreaControl/FormAreaControl";
import FormControl from "src/molecules/FormControl/FormControl";
import FormControlWrapper from "src/molecules/FormControlWrapper/FormControlWrapper";
import { createTask } from "src/services/taskService";

export type ValuePiece = Date | null;

export type Value = ValuePiece | [ValuePiece, ValuePiece];

export default function AddTask() {
    const [isShow, setShow] = useState(false)
    const [title, setTitle] = useState<string>('');
    const [description, setDescription] = useState<string>('');
    const [date, setDate] = useState<Value>(new Date());
    const [time, setTime] = useState("10:00");
    const [dueDate, setDueDate] = useState<string>('');

    useEffect(() => {
        if(date == null){
            return
        }

        if(time == null){
            return
        }

        const dueDateResult = dateToString(date as Date, time.toString())
        setDueDate(dueDateResult)
    }, [date, time])

    const queryClient = useQueryClient()

    const { isLoading, mutate } = useMutation({
        mutationFn: () => {
            return createTask(title, description, dueDate)
        },
        onSuccess: () => {
            queryClient.invalidateQueries({
                queryKey: ['tasks']
            })

            setTitle('')
            setDescription('')
            setShow(false)
        }
    })

    const add = (e: React.FormEvent) => {
        e.preventDefault()

        mutate();
    }

    return (
        <>
            {/* The button to open modal */}
            <button className="btn" onClick={() => setShow(true)}>Create Task +</button>

            <dialog open={isShow} id="my_modal_4" className="modal text-start">
                <form onSubmit={add} method="dialog" className="modal-box w-11/12 max-w-5xl">
                    <Typography size="md">Create New Task</Typography>
                    <FormControl disabled={isLoading} title="Title" value={title} onChange={e => setTitle(e.target.value)} required />
                    <FormControlWrapper title="Due Date">
                        <div className="grid grid-cols-2 gap-4">
                            <div>
                                <DatePicker format="y-MM-dd" className="input input-bordered w-full" onChange={setDate} value={date} required={true}/>
                            </div>
                            <div>
                                <TimePicker format="HH:mm" disableClock={true} className="input input-bordered w-full" onChange={setTime} value={time} required={true}/>
                            </div>
                            { dueDate }
                        </div>
                    </FormControlWrapper>
                    <FormAreaControl disabled={isLoading} title="Description" value={description} onChange={e => setDescription(e.target.value)} required />
                    {
                        isLoading ? (
                            <Loading />
                        ) : (
                            <div className="modal-action">
                                <Button size="md" type="submit" colorType="primary">Submit</Button>
                                <Button type="button" size="md" onClick={() => setShow(false)} colorType="primary">Close</Button>
                            </div>
                        )
                    }

                </form>
            </dialog>
        </>
    )
}