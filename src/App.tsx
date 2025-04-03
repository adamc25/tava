import "./styles.css";
import logo from "./assets/logo.png";

import { BrowserRouter, Route, Routes, NavLink } from "react-router-dom";

import EmployeeList from './employee/EmployeeList';
import EditEmployee from './employee/EditEmployee'
import Payroll from "./Payroll";

export default function App() {
  return (
    <BrowserRouter>
      <nav>
        <div className="logo-container">
          <img src={logo} alt="logo" />
        </div>
        <div>
          <NavLink
          	to="/payroll" 
          	className={({ isActive }) => isActive ? "router active-link" : "router nav-link"}
          >Payroll
        </NavLink>
        </div>
        <div>
          <NavLink
          	to="/employees"
          	className={({ isActive }) => isActive ? "router active-link" : "router nav-link"}
          >Employees
         </NavLink>
        </div>
      </nav>
      <main>
        <Routes>
          <Route path="/payroll" element={<Payroll />} />
          <Route path="/employees" element={<EmployeeList />} />
          <Route path="/employees/:id" element={<EditEmployee />} />
        </Routes>
      </main>
    </BrowserRouter>
  );
}
