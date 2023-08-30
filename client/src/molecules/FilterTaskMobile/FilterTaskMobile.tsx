import { useState } from "react";
import Select from "src/atoms/Select/Select";
import Typography from "src/atoms/Typography/Typography";
import useTask from "src/hooks/useTask";

export default function FilterTaskMobile() {
    const { sort, sorting } = useTask()

    const normalizeSortToSelect = () => {
        return sort.OrderBy + ":" + sort.Sort
    }

    const [selectedFilter, setSelectedFilter] = useState<string>(normalizeSortToSelect())

    const handleSelectChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        const selectedVal = e.target.value.split(":")
        const orderBy  = selectedVal[0] as "created_at" | "due_date"
        const sort = selectedVal[1] as "desc" | "asc"

        sorting({
            OrderBy: orderBy,
            Sort: sort
        });

        setSelectedFilter(e.target.value)
    }

    return (
        <div className="form-control w-full max-w-xs md:hidden">
            <label className="label">
                <Typography size="md">Filter Task:</Typography>
            </label>
           <Select value={selectedFilter} onChange={handleSelectChange} options={
            [
                {
                    key: "created_at:desc",
                    label: "Created At - Newest",
                },
                {
                    key: "created_at:asc",
                    label: "Created At - Oldest",
                },
                {
                    key: "due_date:asc",
                    label: "Due Date - Newest",
                },
                {
                    key: "due_date:desc",
                    label: "Due Date - Oldest",
                }
            ]
           }/>
        </div>
    );
}