import React from "react";
import ReactDOM from "react-dom/client";
import Home from "./modules/home/home.tsx";
import About from "./modules/about/about.tsx";
import Dashboard from "./modules/dashboard/dashboard.tsx";
import Legal from "./modules/legal/legal.tsx";
import "./App.css";

import "./index.css";

import { createBrowserRouter, RouterProvider } from "react-router-dom";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
  },
  {
    path: "/about",
    element: <About />,
  },
  {
    path: "/legal",
    element: <Legal />,
  },
  {
    path: "/dashboard",
    element: <Dashboard />,
  },
]);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <RouterProvider router={router} />
);
