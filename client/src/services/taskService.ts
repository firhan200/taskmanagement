import axios from "axios"
import { getToken } from "src/hooks/useAuth"

const apiUrl: string = import.meta.env.VITE_API_URL ?? ""

const token = getToken()

const addAuthorizationHeader = () => {
    return {
        "Authorization": "Bearer " + token
    }
}

export interface GetTasksResponse extends TasksOrderBy, TasksSort {
    Data: Task[]
    Cursor: string
    Limit: number
    Search: string
    Total: number
    NextCursor: string
}

export type Task = {
    ID: number
    CreatedAt: string
    UpdatedAt: string
    DeletedAt: any
    UserId: number
    Title: string
    Description: string
    DueDate: string
}

export type TasksOrderBy = {
    OrderBy: "created_at" | "due_date"
}

export type TasksSort = {
    Sort: "asc" | "desc"
}

export const getTasks = async (
    cursor: string | null,
    limit: number,
    orderBy: TasksOrderBy,
    sort: TasksSort,
    search: string | ""
): Promise<GetTasksResponse | null> => {
    try {
        // const test = await axios.get(`${apiUrl}test`, {
        //     headers: {
        //         "Authorization": "Bearer "+token,
                
        //     }
        // });

        const res = await axios.get(`${apiUrl}task`, {
            params: {
                cursor: "",
                limit: limit,
                orderBy: orderBy.OrderBy,
                sort: sort.Sort,
                search: "",
            },
            headers: {
                "Authorization": "Bearer "+token,
            }
        });

        const data = <GetTasksResponse>res.data

        return data
    } catch {
        return null
    }
}

export const createTask = async (title: string, description: string, dueDate: string) => {
    try {
        const res = await axios.post(`${apiUrl}tasks`, {
            title: title,
            description: description,
            dueDate: dueDate,
        },
            {
                headers: {
                    ...addAuthorizationHeader()
                }
            }
        );

        const data = res.data

        return data
    } catch {
        return null
    }
}