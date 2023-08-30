import { useState } from "react";
import Button from "src/atoms/Button/Button";
import Input from "src/atoms/Input/Input";
import useTask from "src/hooks/useTask";

export default function SearchTask() {
    const [keyword, setKeyword] = useState('')
    const { search } = useTask()

    const submit = (e: React.FormEvent) => {
        e.preventDefault()

        search(keyword)
    }

    return (
        <form onSubmit={submit} className="flex justify-center">
            <Input type="text" value={keyword} onChange={e => setKeyword(e.target.value)} placeholder="Search task..." />
            <Button colorType="secondary" size="md">Search</Button>
        </form>
    );
}