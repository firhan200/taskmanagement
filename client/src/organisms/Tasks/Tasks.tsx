import { useInfiniteQuery } from "@tanstack/react-query";
import { AxiosError } from "axios";
import React, { useEffect } from "react";
import Alert from "src/atoms/Alert/Alert";
import Button from "src/atoms/Button/Button";
import Loading from "src/atoms/Loading/Loading";
import SkeletonLoading from "src/atoms/SkeletonLoading/SkeletonLoading";
import Typography from "src/atoms/Typography/Typography";
import useAuth from "src/hooks/useAuth";
import useModal from "src/hooks/useModal";
import useTask from "src/hooks/useTask";
import useTheme from "src/hooks/useTheme";
import FilterTaskMobile from "src/molecules/FilterTaskMobile/FilterTaskMobile";
import TaskRow from "src/molecules/TaskRow/TaskRow";
import { getTasks } from "src/services/taskService";

export default function Tasks() {
    const { show } = useModal()
    const { logout } = useAuth()
    const { theme } = useTheme()
    const { sort, keyword, sorting } = useTask()
    const limit = 2;

    const ArrowUp = () => {
        return (
            <svg viewBox="0 0 24 24" width={20} fill="none" xmlns="http://www.w3.org/2000/svg"><g id="SVGRepo_bgCarrier" strokeWidth="0"></g><g id="SVGRepo_tracerCarrier" strokeLinecap="round" strokeLinejoin="round"></g><g id="SVGRepo_iconCarrier"> <path d="M7 3V21M7 3L11 7M7 3L3 7M14 3H21M14 9H19M14 15H17M14 21H15" stroke={theme == "dark" ? "#e6e6e6" : "#000"} strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"></path> </g></svg>
        )
    }

    const ArrowDown = () => {
        return (
            <svg viewBox="0 0 24 24" width={20} fill="none" xmlns="http://www.w3.org/2000/svg"><g id="SVGRepo_bgCarrier" strokeWidth="0"></g><g id="SVGRepo_tracerCarrier" strokeLinecap="round" strokeLinejoin="round"></g><g id="SVGRepo_iconCarrier"> <path d="M7 3V21M7 21L3 17M7 21L11 17M14 3H21M14 9H19M14 15H17M14 21H15" stroke={theme == "dark" ? "#e6e6e6" : "#000"} strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"></path> </g></svg>
        )
    }

    const LoadingState = () => {
        return (
            <div>
                {
                    [...Array(3)].map((_, key) => (
                        <div key={key} className="my-6">
                            <div className="grid grid-cols-4 gap-4 dark:border-slate-700 rounded-xl border p-6">
                                <SkeletonLoading height={30} type="square" isFull={true} />
                                <SkeletonLoading height={30} type="square" isFull={true} />
                                <SkeletonLoading height={30} type="square" isFull={true} />
                                <SkeletonLoading height={30} type="square" isFull={true} />
                            </div>
                        </div>
                    ))
                }
            </div>
        )
    }

    useEffect(() => {
        remove()
    }, [sort, keyword])

    const fetchTasks = async ({ pageParam = '' }) => {
        try {
            const res = await getTasks(
                pageParam,
                limit,
                sort,
                keyword
            )

            return res
        } catch (err) {
            const axiosError = err as AxiosError;
            const unauthorized = axiosError.response?.status === 401 ?? false
            if (unauthorized) {
                show("Unauthorized", logout)
            }
        }
    }

    const {
        data,
        error,
        isLoading,
        isError,
        fetchNextPage,
        hasNextPage,
        isFetching,
        isFetchingNextPage,
        remove
    } = useInfiniteQuery({
        queryKey: [`tasks`],
        queryFn: fetchTasks,
        getNextPageParam: (lastPage) => {
            if (lastPage == null) {
                return
            }
            return lastPage.NextCursor
        },
        refetchOnWindowFocus: false,
        cacheTime: 1
    })

    if (isLoading) {
        return <LoadingState />
    }

    if (isError && error instanceof AxiosError) {
        return <Alert type="error" text={error.message} />
    }

    if (data == null) {
        return <Alert type="error" text="Failed to get Tasks" />
    }

    const solidDate = data!

    const applyOrder = (orderBy: "created_at" | "due_date") => {
        sorting({
            OrderBy: orderBy,
            Sort: sort.Sort == "asc" ? "desc" : "asc"
        })
    }

    const RenderOrderBy = (title: string, orderBy: "created_at" | "due_date") => {
        return (
            <span onClick={() => applyOrder(orderBy)} className="cursor-pointer hover:underline flex gap-2">
                {title}
                {
                    orderBy == sort.OrderBy ? (
                        sort.Sort == "asc" ? <ArrowUp /> : <ArrowDown />
                    ) : null
                }

            </span>
        )
    }

    return (
        <div>
            <FilterTaskMobile />
            <div className="hidden md:block my-8">
                <div className="grid md:grid-cols-5 lg:grid-cols-6">
                    <div></div>
                    <div className="grid grid-cols-1 lg:grid-cols-2 md:col-span-2 lg:col-span-3">
                        <Typography size="md" className="font-bold">Name</Typography>
                        <Typography size="md" className="font-bold">Description</Typography>
                    </div>
                    <Typography size="md" className="font-bold">{RenderOrderBy("Due Date", "due_date")}</Typography>
                    <Typography size="md" className="font-bold flex justify-end">{RenderOrderBy("Created On", "created_at")}</Typography>
                </div>
            </div>

            {solidDate.pages.map((group, i) => (
                <React.Fragment key={i}>
                    {group?.Data.map((task) => (
                        <TaskRow key={task.ID} task={task} />
                    ))}
                </React.Fragment>
            ))}
            <div className="text-center">
                <Button
                    colorType="primary"
                    size="md"
                    onClick={() => fetchNextPage()}
                    disabled={!hasNextPage || isFetchingNextPage}
                >
                    {isFetchingNextPage
                        ? <Loading />
                        : hasNextPage
                            ? 'Load More'
                            : 'Nothing more to load'}
                </Button>
            </div>
            <div>{isFetching && !isFetchingNextPage ? 'Fetching...' : null}</div>
        </div>
    );
}