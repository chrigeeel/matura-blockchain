import React from "react";
import ReactDOM from "react-dom/client";
import { createBrowserRouter, RouterProvider, Route } from "react-router-dom";
import "./css/index.css";
import Frame from "./components/Frame/Frame";
import Blockchain from "./components/Blockchain/Blockchain";
import Process from "./components/Process/Process";
import Home from "./components/Home/Home";

const router = createBrowserRouter([
	{
		path: "/",
		element: <Frame />,
		children: [
			{
				path: "",
				element: <Home />,
			},
			{
				path: "blockchain",
				element: <Blockchain />,
			},
			{
				path: "process",
				element: <Process />,
			},
		],
		errorElement: <div>Not found!</div>,
	},
]);

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
	<React.StrictMode>
		<RouterProvider router={router} />
	</React.StrictMode>
);
