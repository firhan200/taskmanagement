import { TasksSort } from "src/services/taskService";
import { create } from "zustand";

type TaskStoreState = {
    editId: number | null
	currentTotal: number | null
	total: number
    keyword: string | ""
    sort: TasksSort
	search: (keyword: string) => void
	sorting: ( sort: TasksSort) => void
	edit: (id: number | null) => void
	setTotal: (total: number) => void
	setCurrentTotal: (totalFetchedData: number) => void
}

const taskStore = create<TaskStoreState>()((set) => ({
    editId: null,
	currentTotal: 0,
	total: 0,
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
	})),
	setTotal: (total: number) => set(() => ({
		total: total
	})),
	setCurrentTotal: (totalFetchedData: number) => set((state) => ({
		currentTotal: state.currentTotal ?? 0 + totalFetchedData
	}))
}))

export default function useTask(){
	const [
		keyword, 
		sort, 
		search, 
		sorting, 
		edit, 
		editId, 
		total, 
		currentTotal, 
		setTotal,  
		setCurrentTotal
	] = taskStore((state) => [
		state.keyword, 
		state.sort, 
		state.search, 
		state.sorting, 
		state.edit, 
		state.editId,
		state.total,
		state.currentTotal,
		state.setTotal,
		state.setCurrentTotal,
	]) 

	return {
		keyword,
		sort,
		search,
		sorting,
		edit,
		editId,
		total,
		currentTotal,
		setTotal,
		setCurrentTotal
	}
}