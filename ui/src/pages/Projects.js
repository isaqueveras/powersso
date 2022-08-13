import Sidebar from "../components/Sidebar";

export function Projects() {
  return (
    <div className="flex flex-row dark:bg-black/95 w-screen">
      <div className="w-64">
        <Sidebar />
      </div>
      <span className="p-5 text-white">Projects</span>
    </div>
  )
}