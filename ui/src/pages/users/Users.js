import { Link } from "react-router-dom";

import Sidebar from "../../components/Sidebar";

export function Users() {
  return (
    <>
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
          <nav class="p-4" aria-label="Breadcrumb">
            <ol class="inline-flex items-center space-x-1 md:space-x-3">
              <li class="inline-flex items-center">
                <Link to="/" class="inline-flex items-center text-sm font-medium text-gray-700 hover:text-gray-900 dark:text-gray-400 dark:hover:text-white">
                  <svg class="w-4 h-4 mr-2" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path d="M10.707 2.293a1 1 0 00-1.414 0l-7 7a1 1 0 001.414 1.414L4 10.414V17a1 1 0 001 1h2a1 1 0 001-1v-2a1 1 0 011-1h2a1 1 0 011 1v2a1 1 0 001 1h2a1 1 0 001-1v-6.586l.293.293a1 1 0 001.414-1.414l-7-7z"></path></svg>
                  Home
                </Link>
              </li>
              <li aria-current="page">
                <div class="flex items-center">
                  <svg class="w-6 h-6 text-gray-400" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd"></path></svg>
                  <span class="ml-1 text-sm font-medium text-gray-500 md:ml-2 dark:text-gray-400">Users</span>
                </div>
              </li>
            </ol>
          </nav>
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
      
      <div id="drawer-example" class="fixed z-40 h-screen p-4 overflow-y-auto bg-white w-80 dark:bg-gray-800" tabindex="-1" aria-labelledby="drawer-label">
        <h5 id="drawer-label" class="inline-flex items-center mb-4 text-base font-semibold text-gray-500 dark:text-gray-400"><svg class="w-5 h-5 mr-2" aria-hidden="true" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd"></path></svg>Info</h5>
        <button type="button" data-drawer-dismiss="drawer-example" aria-controls="drawer-example" class="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 absolute top-2.5 right-2.5 inline-flex items-center dark:hover:bg-gray-600 dark:hover:text-white" >
          <svg aria-hidden="true" class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd"></path></svg>
          <span class="sr-only">Close menu</span>
        </button>
        <p class="mb-6 text-sm text-gray-500 dark:text-gray-400">Supercharge your hiring by taking advantage of our <a href="/" class="text-blue-600 underline dark:text-blue-500 hover:no-underline">limited-time sale</a> for Flowbite Docs + Job Board. Unlimited access to over 190K top-ranked candidates and the #1 design job board.</p>
        <div class="grid grid-cols-2 gap-4">
          <a href="/" class="px-4 py-2 text-sm font-medium text-center text-gray-900 bg-white border border-gray-200 rounded-lg focus:outline-none hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-gray-200 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700">Learn more</a>
          <a href="/" class="inline-flex items-center px-4 py-2 text-sm font-medium text-center text-white bg-blue-700 rounded-lg hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800">Get access <svg class="w-4 h-4 ml-2" aria-hidden="true" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M12.293 5.293a1 1 0 011.414 0l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-2.293-2.293a1 1 0 010-1.414z" clip-rule="evenodd"></path></svg></a>
        </div>
      </div>
    </>
  )
}