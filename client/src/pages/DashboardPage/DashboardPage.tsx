import SearchTask from "src/molecules/SearchTask/SearchTask";
import AddTask from "src/organisms/AddTask/AddTask";
import EditTask from "src/organisms/EditTask/EditTask";
import ProfileBoard from "src/organisms/ProfileBoard/ProfileBoard";
import Tasks from "src/organisms/Tasks/Tasks";

const DashboardPage = () => {
    return (
        <div>
            <ProfileBoard />
            <div className="my-6 text-end flex justify-between">
                <SearchTask />
                <AddTask />
            </div>
            <Tasks />
            <EditTask />
        </div>
    );
}

export default DashboardPage;