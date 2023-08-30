import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useEffect, useState } from "react";
import DatePicker from "react-date-picker";
import TimePicker from "react-time-picker";
import Button from "src/atoms/Button/Button";
import Loading from "src/atoms/Loading/Loading";
import Typography from "src/atoms/Typography/Typography";
import { dateStringToDateAndTime, dateToString } from "src/helpers/date";
import useTask from "src/hooks/useTask";
import FormAreaControl from "src/molecules/FormAreaControl/FormAreaControl";
import FormControl from "src/molecules/FormControl/FormControl";
import FormControlWrapper from "src/molecules/FormControlWrapper/FormControlWrapper";
import { getTaskById, updateTask } from "src/services/taskService";

export type ValuePiece = Date | null;

export type Value = ValuePiece | [ValuePiece, ValuePiece];

export default function EditTask() {
    const { editId, edit } = useTask()
    console.log(editId)

    const [title, setTitle] = useState<string>('');
    const [description, setDescription] = useState<string>('');
    const [date, setDate] = useState<Value>(new Date());
    const [time, setTime] = useState("10:00");
    const [dueDate, setDueDate] = useState<string>('');
    const [loadingData, setLoading] = useState(false);

    const LoadingState = () => {
        return (
            <div className="justify-center flex gap-4">
                <Typography size="md">Getting Data</Typography>
                <Loading />
            </div>
        )
    }

    useEffect(() => {
        if (editId === null) {
            return
        }

        async function getTask() {
            setLoading(true)

            const res = await getTaskById(editId!)
            if (res == null) {
                //not found
            }

            const { dateValue, timeValue } = dateStringToDateAndTime(res!.DueDate)

            //parse due date
            setTitle(res!.Title)
            setDescription(res!.Title)
            setDate(dateValue)
            setTime(timeValue)

            setLoading(false)
        }

        getTask();
    }, [editId])

    useEffect(() => {
        if (date == null) {
            return
        }

        if (time == null) {
            return
        }

        const dueDateResult = dateToString(date as Date, time.toString())
        setDueDate(dueDateResult)
    }, [date, time])

    const queryClient = useQueryClient()

    const { isLoading, mutate } = useMutation({
        mutationFn: async () => {
            if(editId == null){
                return null
            }
            return await updateTask(editId!, title, description, dueDate)
        },
        onSuccess: () => {
            queryClient.invalidateQueries({
                queryKey: ['tasks']
            })

            window.alert("success")

            resetForm()
            edit(null)
        }
    })

    const submit = (e: React.FormEvent) => {
        e.preventDefault()

        mutate();
    }

    const resetForm = () => {
        setTitle('')
        setDescription('')
        setDate(new Date())
        setTime('10:00')
    }

    const closeModal = () => {
        resetForm()

        edit(null)
    }

    return (
        <dialog open={editId !== null} id="my_modal_4" className="modal text-start">
            <form onSubmit={submit} method="dialog" className="modal-box w-11/12 max-w-5xl">
                {
                    loadingData ? <LoadingState /> : (
                        <>
                            <Typography size="md">{ title }</Typography>
                            <FormControl disabled={isLoading} title="Title" value={title} onChange={e => setTitle(e.target.value)} required />
                            <FormControlWrapper title="Due Date">
                                <div className="grid grid-cols-2 gap-4">
                                    <div>
                                        <DatePicker format="y-MM-dd" className="input input-bordered w-full" onChange={setDate} value={date} required={true} />
                                    </div>
                                    <div>
                                        <TimePicker format="HH:mm" disableClock={true} className="input input-bordered w-full" onChange={setTime} value={time} required={true} />
                                    </div>
                                    {dueDate}
                                </div>
                            </FormControlWrapper>
                            <FormAreaControl disabled={isLoading} title="Description" value={description} onChange={e => setDescription(e.target.value)} required />
                            {
                                isLoading ? (
                                    <Loading />
                                ) : (
                                    <div className="modal-action">
                                        <Button type="button" size="md" onClick={() => closeModal()} colorType="danger">Delete</Button>
                                        <Button size="md" type="submit" colorType="primary">Save</Button>
                                        <Button type="button" size="md" onClick={() => closeModal()} colorType="primary">Close</Button>
                                    </div>
                                )
                            }
                        </>
                    )
                }

            </form>
        </dialog>
    )
}