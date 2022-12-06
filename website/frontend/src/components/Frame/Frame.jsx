import React from "react";
import { Outlet } from "react-router-dom";
import Topbar from "./Topbar";

const Frame = () => {
	return (
		<div className="flex bg-vgray-900 w-screen h-screen text-light">
			<div className="flex flex-col h-screen w-full">
				<div className="z-50">
					<Topbar />
				</div>
				<div className="flex flex-grow overflow-y-auto">
					<div className="w-full">
						<div className="flex flex-col min-h-full w-full">
							<div className="flex flex-grow m-5 mt-2 lg:mt-4">
								<Outlet />
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	);
};

export default Frame;
