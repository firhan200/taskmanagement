import { useInfiniteQuery } from "@tanstack/react-query";
import React from "react";
import { useState } from "react";
import Alert from "src/atoms/Alert/Alert";
import SkeletonLoading from "src/atoms/SkeletonLoading/SkeletonLoading";
import Typography from "src/atoms/Typography/Typography";
import TaskRow from "src/molecules/TaskRow/TaskRow";
import { TasksOrderBy, TasksSort, getTasks } from "src/services/taskService";

export default function Tasks() {
    const limit = 2;
    const [keyword, setKeyword] = useState<string>('');
    const [orderBy, setOrderBy] = useState<TasksOrderBy>({ OrderBy: "created_at" });
    const [sort, setSort] = useState<TasksSort>({ Sort: "desc" });

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

    const fetchTasks = async ({ pageParam = '' }) => {
        return await getTasks(
            pageParam,
            limit,
            orderBy,
            sort,
            keyword
        )
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
    } = useInfiniteQuery({
        queryKey: ['tasks'],
        queryFn: fetchTasks,
        getNextPageParam: (lastPage) => lastPage?.NextCursor,
    })

    if (isLoading) {
        return <LoadingState />
    }

    if (isError && error instanceof Error) {
        return <Alert type="error" text={error.message} />
    }

    if (data == null) {
        return <Alert type="error" text="Failed to get Tasks" />
    }

    const solidDate = data!

    return (
        <div>
            <Typography size="md">Tasks</Typography>
            {solidDate.pages.map((group, i) => (
                <React.Fragment  key={i}>
                    {group?.Data.map((task) => (
                        <TaskRow key={task.ID} task={task}/>
                    ))}
                </React.Fragment>
            ))}
            <div>
                <button
                    onClick={() => fetchNextPage()}
                    disabled={!hasNextPage || isFetchingNextPage}
                >
                    {isFetchingNextPage
                        ? 'Loading more...'
                        : hasNextPage
                            ? 'Load More'
                            : 'Nothing more to load'}
                </button>
            </div>
            <div>{isFetching && !isFetchingNextPage ? 'Fetching...' : null}</div>
        </div>
    );
}