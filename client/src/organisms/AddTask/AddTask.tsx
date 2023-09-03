import { useMutation, useQueryClient } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { useEffect, useState } from "react";
import DatePicker from "react-date-picker";
import TimePicker from "react-time-picker";
import Button from "src/atoms/Button/Button";
import CharCounter from "src/atoms/CharCounter/CharCounter";
import Loading from "src/atoms/Loading/Loading";
import Typography from "src/atoms/Typography/Typography";
import { dateToString } from "src/helpers/date";
import useModal from "src/hooks/useModal";
import FormAreaControl from "src/molecules/FormAreaControl/FormAreaControl";
import FormControl from "src/molecules/FormControl/FormControl";
import FormControlWrapper from "src/molecules/FormControlWrapper/FormControlWrapper";
import { CreateTaskErrorResponse, createTask } from "src/services/taskService";

export type ValuePiece = Date | null;

export type Value = ValuePiece | [ValuePiece, ValuePiece];

export default function AddTask() {
    const { show } = useModal()

    const titleMax = 100
    const descMax = 200

    const now = new Date()
    const [hour, minute] = [now.getHours(), now.getMinutes()]

    const [isShow, setShow] = useState(false)
    const [title, setTitle] = useState<string>('');
    const [description, setDescription] = useState<string>('');
    const [date, setDate] = useState<Value>(now);
    const [time, setTime] = useState<string | null>(`${hour}:${minute}`);
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

            show(`Succesfully Creating Task: ${title}`, null)
        },
        onError: (err: AxiosError) => {
            const data = err.response?.data as CreateTaskErrorResponse
            alert(`Failed to Create Task: ${data.error}`)
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
                    <FormControl disabled={isLoading} title="Name" value={title} onChange={e => setTitle(e.target.value)} maxLength={titleMax} required />
                    <CharCounter current={title.length} max={titleMax}/>
                    <FormControlWrapper title="Due Date">
                        <div className="grid grid-cols-2 gap-4">
                            <div>
                                <DatePicker format="y-MM-dd" className="input input-bordered w-full" onChange={setDate} value={date} required={true}/>
                            </div>
                            <div>
                                <TimePicker format="HH:mm" disableClock={true} className="input input-bordered w-full" onChange={setTime} value={time} required={true}/>
                            </div>
                        </div>
                    </FormControlWrapper>
                    <FormAreaControl disabled={isLoading} title="Description" value={description} onChange={e => setDescription(e.target.value)} maxLength={descMax} required />
                    <CharCounter current={description.length} max={descMax}/>
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