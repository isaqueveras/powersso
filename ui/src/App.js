import { Routes, Route } from "react-router-dom";

import { Dashboard, Projects } from "./pages";
import { Users, NewUser } from "./pages/users";

export default function App() {
  return (
    <div className="h-screen">
      <Routes>
        <Route path="/" element={<Dashboard />} />
        <Route path="projects" element={<Projects />} />
        <Route path="users" element={<Users />} />
        <Route path="users/new" element={<NewUser />} />
      </Routes>
    </div>
  )
}
