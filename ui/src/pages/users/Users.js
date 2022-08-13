import Sidebar from "../../components/Sidebar";

export function Users() {
  return (
    <div className="flex flex-row dark:bg-black/95 w-screen">
      <div className="w-64">
        <Sidebar />
      </div>
      <div className="p-5 w-screen">
        <div className="flex flex-row align-middle justify-between">
          <div>
            <h1 className="text-3xl font-bold dark:text-white">Users</h1>
            <p className="text-base font-normal dark:text-gray-200">User Management</p>
          </div>
          <button type="button" class="text-white bg-orange-700 hover:bg-orange-800 focus:ring-2 focus:ring-orange-300 font-semibold rounded-md text-md px-5 py-2.5 mr-2 mb-2 dark:bg-orange-600 dark:hover:bg-orange-700 focus:outline-none dark:focus:ring-orange-800">New User</button>
        </div>
      </div>
    </div>
  )
}