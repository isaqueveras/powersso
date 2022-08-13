import Sidebar from "../../components/Sidebar";

export function Users() {
  return (
    <div className="flex flex-row dark:bg-black/95 w-screen">
      <div className="w-64">
        <Sidebar />
      </div>
      <div className="w-full">
        <div className="border-b border-gray-200 dark:border-black/30 h-16">
          <div className="h-16 px-4 flex flex-row justify-between items-center">
            <h1 className="text-2xl text-gray-900 font-medium dark:text-white/90">Users</h1>
            <button type="button" className="ml-2 shadow-sm text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-2 focus:ring-gray-200 font-medium rounded-lg text-sm px-5 py-2.5 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700">New User</button>
          </div>
        </div>
        <div className="overflow-x-auto relative">
          <table className="w-full text-sm text-left text-gray-500 dark:text-gray-400">
            <caption className="p-5 text-lg font-semibold text-left text-gray-900 bg-white dark:text-white/90 dark:bg-transparent">
              Our users
              <p className="mt-1 text-sm font-normal text-gray-500 dark:text-gray-400">List of the last users registered in the system.</p>
            </caption>
            <thead className="text-xs text-gray-700 uppercase dark:text-gray-400">
              <tr>
                <th scope="col" className="py-3 px-6">Name</th>
                <th scope="col" className="py-3 px-6">About</th>
                <th scope="col" className="py-3 px-6">Phone number</th>
                <th scope="col" className="py-3 px-6">City</th>
                <th scope="col" className="py-3 px-6">Country</th>
                <th scope="col" className="py-3 px-6">Birthday</th>
              </tr>
            </thead>
            <tbody>
              <tr className="bg-white border-b dark:bg-transparent dark:border-black/30">
                <th scope="row" className="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white/90">
                  <div className="text-base font-semibold">Isaque Veras</div>
                  <div className="font-normal text-gray-500">isaque@veras.com</div>
                </th>
                <td className="py-4 px-6">I'm a developer</td>
                <td className="py-4 px-6">(88) 9.9999.9999</td>
                <td className="py-4 px-6">Quixadá, CE</td>
                <td className="py-4 px-6">Brazil</td>
                <td className="py-4 px-6">04/02/2002</td>
              </tr>
              <tr className="bg-white border-b dark:bg-transparent dark:border-black/30">
                <th scope="row" className="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white/90">
                  <div className="text-base font-semibold">Isaque Veras</div>
                  <div className="font-normal text-gray-500">isaque@veras.com</div>
                </th>
                <td className="py-4 px-6">I'm a developer</td>
                <td className="py-4 px-6">(88) 9.9999.9999</td>
                <td className="py-4 px-6">Quixadá, CE</td>
                <td className="py-4 px-6">Brazil</td>
                <td className="py-4 px-6">04/02/2002</td>
              </tr>
              <tr className="bg-white border-b dark:bg-transparent dark:border-black/30">
                <th scope="row" className="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white/90">
                  <div className="text-base font-semibold">Isaque Veras</div>
                  <div className="font-normal text-gray-500">isaque@veras.com</div>
                </th>
                <td className="py-4 px-6">I'm a developer</td>
                <td className="py-4 px-6">(88) 9.9999.9999</td>
                <td className="py-4 px-6">Quixadá, CE</td>
                <td className="py-4 px-6">Brazil</td>
                <td className="py-4 px-6">04/02/2002</td>
              </tr>
            </tbody>
          </table>
        </div>
        <div className="px-6 py-3 flex justify-between items-center">
          <span className="text-sm text-gray-700 dark:text-gray-400">
            Showing <span className="font-semibold text-gray-900 dark:text-white">1</span> to <span className="font-semibold text-gray-900 dark:text-white">3</span> of <span className="font-semibold text-gray-900 dark:text-white">6</span> Entries
          </span>
          <div className="flex">
            <a href="?page=1" className="inline-flex items-center py-2 px-4 mr-3 text-sm font-medium text-gray-500 bg-white rounded-lg border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">
              <svg aria-hidden="true" className="mr-2 w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M7.707 14.707a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 1.414L5.414 9H17a1 1 0 110 2H5.414l2.293 2.293a1 1 0 010 1.414z" clip-rule="evenodd"></path></svg>
              Previous
            </a>
            <a href="?page=2" className="inline-flex items-center py-2 px-4 text-sm font-medium text-gray-500 bg-white rounded-lg border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">
              Next
              <svg aria-hidden="true" className="ml-2 w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M12.293 5.293a1 1 0 011.414 0l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-2.293-2.293a1 1 0 010-1.414z" clip-rule="evenodd"></path></svg>
            </a>
          </div>
        </div>
      </div>
    </div>
  )
}