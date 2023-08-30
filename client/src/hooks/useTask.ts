import { TasksSort } from "src/services/taskService";
import { create } from "zustand";

type TaskStoreState = {
    editId: number | null
    keyword: string | ""
    sort: TasksSort
	search: (keyword: string) => void
	sorting: ( sort: TasksSort) => void
	edit: (id: number | null) => void
}

const taskStore = create<TaskStoreState>()((set) => ({
    editId: null,
    keyword: "",
	sort: { OrderBy: 'created_at', Sort: 'desc' },
    search: (keyword: string) => set(() => ({
		keyword: keyword,
	})),
	sorting: (sort: TasksSort) => set(() => ({
		sort: sort,
	})),
	edit: (id: number | null) => set(() => ({
		editId: id
	}))
}))

export default function useTask(){
	const [keyword, sort, search, sorting, edit, editId] = taskStore((state) => [
		state.keyword, 
		state.sort, 
		state.search, 
		state.sorting, 
		state.edit, 
		state.editId]) 

	return {
		keyword,
		sort,
		search,
		sorting,
		edit,
		editId
	}
}