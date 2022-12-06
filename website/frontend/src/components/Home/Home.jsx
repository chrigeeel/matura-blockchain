import React from "react";

import matura from "./matura-christian-tognazza-6dg-6.12.2022.pdf";

const Home = () => {
	return (
		<div className="flex flex-col m-5 mx-auto max-w-5xl w-full">
			<object data={matura} className="h-full"></object>
		</div>
	);
};

export default Home;
