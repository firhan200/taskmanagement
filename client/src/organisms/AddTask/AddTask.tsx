import { useMutation } from "@tanstack/react-query";
import { useState } from "react";
import Button from "src/atoms/Button/Button";
import Input from "src/atoms/Input/Input";
import Loading from "src/atoms/Loading/Loading";
import Typography from "src/atoms/Typography/Typography";
import FormAreaControl from "src/molecules/FormAreaControl/FormAreaControl";
import FormControl from "src/molecules/FormControl/FormControl";
import { createTask } from "src/services/taskService";

export default function AddTask() {
    const [isShow, setShow] = useState(false)
    const [title, setTitle] = useState<string>('');
    const [description, setDescription] = useState<string>('');

    const { isLoading, mutate } = useMutation({
        mutationFn: () => {
            return createTask(title, description, '2023-10-03T10:30:00Z')
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