import { Link, useLocation } from "react-router-dom";

import rychenbergLogo from "./rychenbergLogo.png";

const routes = [
	{
		path: "",
		name: "Home",
	},
	{
		path: "blockchain",
		name: "Demo",
	},
	{
		path: "process",
		name: "Process",
	},
];

const Topbar = () => {
	const location = useLocation();

	console.log(location);

	return (
		<div
			className="sticky justify-center md:justify-start flex items-center w-full h-[72px] md:h-20 bg-vgray-700
            border-vgray-300 border-b-2"
		>
			<div className="hidden md:flex items-center h-full border-x-2 border-vgray-300">
				<a
					href="https://www.krw.ch/"
					target="_blank"
					className="mx-3 h-14"
				>
					<img className="h-14" src={rychenbergLogo} alt="" />
				</a>
			</div>

			<div className="flex items-center px-8 gap-12 h-full">
				{routes.map((route) => {
					let className = "";
					if (location.pathname.slice(1) == route.path) {
						className = "!text-green";
					}

					return (
						<Link
							className={`font-semibold transition duration-300 
                            hover:text-opacity-100 text-light text-opacity-80 ${className}`}
							to={route.path}
						>
							{route.name}
						</Link>
					);
				})}
			</div>
		</div>
	);
};

export default Topbar;
