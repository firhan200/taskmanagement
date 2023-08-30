import axios from "axios";
import { getToken } from "src/hooks/useAuth";

const apiUrl: string = import.meta.env.VITE_API_URL ?? "";

const token = getToken();

const addAuthorizationHeader = () => {
  return {
    Authorization: "Bearer " + token,
  };
};

export interface GetTasksResponse extends TasksSort {
  Data: Task[];
  Cursor: string;
  Limit: number;
  Search: string;
  Total: number;
  NextCursor: string;
}

export type Task = {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string;
  UserId: number;
  Title: string;
  Description: string;
  DueDate: string;
  Status: "Not urgent" | "Overdue" | "Due soon";
};

export type TasksSort = {
  OrderBy: "created_at" | "due_date";
  Sort: "desc" | "asc";
};

export const getTasks = async (
  cursor: string | null,
  limit: number,
  sort: TasksSort,
  search: string | ""
): Promise<GetTasksResponse | null> => {
  const res = await axios.get(`${apiUrl}tasks`, {
    params: {
      cursor: cursor,
      limit: limit,
      orderBy: sort.OrderBy,
      sort: sort.Sort,
      search: search,
    },
    headers: {
      Authorization: "Bearer " + token,
    },
  });

  const data = <GetTasksResponse>res.data;

  return data;
};

export const createTask = async (
  title: string,
  description: string,
  dueDate: string
) => {
  try {
    const res = await axios.post(
      `${apiUrl}tasks`,
      {
        title: title,
        description: description,
        dueDate: dueDate,
      },
      {
        headers: {
          ...addAuthorizationHeader(),
        },
      }
    );

    const data = res.data;

    return data;
  } catch {
    return null;
  }
};

export const getTaskById = async (id: number): Promise<Task | null> => {
  try {
    const res = await axios.get(`${apiUrl}tasks/${id}`, {
      headers: {
        ...addAuthorizationHeader(),
      },
    });

    const data = res.data;

    return data;
  } catch {
    return null;
  }
};

export const updateTask = async (
  id: number,
  title: string,
  description: string,
  dueDate: string
) => {
  try {
    const res = await axios.patch(
      `${apiUrl}tasks/${id}`,
      {
        title: title,
        description: description,
        dueDate: dueDate,
      },
      {
        headers: {
          ...addAuthorizationHeader(),
        },
      }
    );

    const data = res.data;

    return data;
  } catch {
    return null;
  }
};

type DeleteTaskResponse = {
    error: string | undefined
    task: Task | undefined
}

export const deleteTask = async (id: number): Promise<DeleteTaskResponse> => {
  const res = await axios.delete(`${apiUrl}tasks/${id}`, {
    headers: {
      ...addAuthorizationHeader(),
    },
  });

  const data = res.data;

  return data;
};
