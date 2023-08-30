import AddTask from "src/organisms/AddTask/AddTask";
import ProfileBoard from "src/organisms/ProfileBoard/ProfileBoard";
import Tasks from "src/organisms/Tasks/Tasks";

const DashboardPage = () => {
    return (
        <div>
            <ProfileBoard />
            <div className="my-4 text-end">
                <AddTask />
            </div>
            <Tasks />
        </div>
    );
}

export default DashboardPage;