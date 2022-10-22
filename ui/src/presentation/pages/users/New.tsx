import { Link } from "react-router-dom";

import Sidebar from "../../components/sidebar";

const NewUser = () => {
  return (
    <div className="flex flex-row dark:bg-black/95 w-screen">
      <div className="w-64">
        <Sidebar />
      </div>
      <div className="w-full">
        <div className="border-b border-gray-200 dark:border-black/30 h-16">
          <div className="h-16 px-4 flex flex-row justify-between items-center">
            <h1 className="text-2xl text-gray-900 font-medium dark:text-white">New user</h1>
          </div>
        </div>

				<nav className="p-4" aria-label="Breadcrumb">
          <ol className="inline-flex items-center space-x-1 md:space-x-3">
            <li className="inline-flex items-center">
              <Link to="/" className="inline-flex items-center text-sm font-medium text-gray-700 hover:text-gray-900 dark:text-gray-400 dark:hover:text-white">
                <svg className="w-4 h-4 mr-2" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path d="M10.707 2.293a1 1 0 00-1.414 0l-7 7a1 1 0 001.414 1.414L4 10.414V17a1 1 0 001 1h2a1 1 0 001-1v-2a1 1 0 011-1h2a1 1 0 011 1v2a1 1 0 001 1h2a1 1 0 001-1v-6.586l.293.293a1 1 0 001.414-1.414l-7-7z"></path></svg>Home
              </Link>
            </li>
						<li>
							<div className="flex items-center">
								<svg className="w-6 h-6 text-gray-400" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd"></path></svg>
								<Link to="/users" className="ml-1 text-sm font-medium text-gray-700 hover:text-gray-900 md:ml-2 dark:text-gray-400 dark:hover:text-white">Users</Link>
							</div>
						</li>
            <li aria-current="page">
              <div className="flex items-center">
                <svg className="w-6 h-6 text-gray-400" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd"></path></svg>
                <span className="ml-1 text-sm font-medium text-gray-500 md:ml-2 dark:text-gray-400">New user</span>
              </div>
            </li>
          </ol>
        </nav>

				<div className="px-2 flex flex-wrap overflow-hidden">
					<div className="py-2 px-2 w-full overflow-hidden sm:w-full md:w-1/3 xl:w-1/3">
						<label className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">First name</label>
						<input type="text" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="name@flowbite.com" required/>
					</div>

					<div className="py-2 px-2 w-full overflow-hidden sm:w-full md:w-1/3 xl:w-1/3">
						<label className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">Last name</label>
						<input type="text" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="name@flowbite.com" required/>
					</div>

					<div className="py-2 px-2 w-full overflow-hidden sm:w-full md:w-1/3 xl:w-1/3">
						<label className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">Email</label>
						<input type="email" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="name@flowbite.com" required/>
					</div>
				</div>
      </div>
    </div>
  )
}

export default NewUser