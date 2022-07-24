import Sidebar from "../components/Sidebar";

export function Dashboard() {
  return (
    <div className="flex flex-row">
      <Sidebar />
      <span>Hello!</span>
    </div>
  )
}