import Sidebar from '../components/sidebar'

const Dashboard: React.FC  = () => {
  return (
    <div className="flex flex-row dark:bg-black/95 w-screen">
      <div className="w-64">
        <Sidebar />
      </div>
      <div className="w-full">
        <div className="border-b border-gray-200 dark:border-black/30 h-16">
          <div className="h-16 px-4 flex flex-row justify-between items-center">
            <h1 className="text-2xl text-gray-900 font-medium dark:text-white">Home</h1>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Dashboard
