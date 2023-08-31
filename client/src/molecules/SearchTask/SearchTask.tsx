import { useState } from "react";
import Button from "src/atoms/Button/Button";
import Input from "src/atoms/Input/Input";
import SearchResult from "src/atoms/SearchResult/SearchResult";
import useTask from "src/hooks/useTask";

export default function SearchTask() {
    const { keyword, total } = useTask()
    const [ localKeyword, setKeyword] = useState('')
    const { search } = useTask()

    const submit = (e: React.FormEvent) => {
        e.preventDefault()

        search(localKeyword)
    }

    return (
        <div className="text-start">
            <form onSubmit={submit} className="flex justify-center">
                <Input type="text" value={localKeyword} onChange={e => setKeyword(e.target.value)} placeholder="Search task..." />
                <Button colorType="secondary" size="md">Search</Button>
            </form>
            <SearchResult keyword={keyword} total={total} />
        </div>
    );
}