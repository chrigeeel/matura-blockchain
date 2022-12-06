import React from "react";

const Conversation = ({ height, width }) => {
	return (
		<svg
			height={height}
			width={width}
			xmlns="http://www.w3.org/2000/svg"
			viewBox="0 0 24 24"
		>
			<g>
				<path
					d="M10.25,12.06a2.5,2.5,0,0,1,2.5-2.5H20a.25.25,0,0,0,.25-.25V3.06a1.5,1.5,0,0,0-1.5-1.5h-17a1.5,1.5,0,0,0-1.5,1.5v10a1.5,1.5,0,0,0,1.5,1.5h1.5v3.5a.5.5,0,0,0,.5.5.47.47,0,0,0,.35-.15L8,14.56h2a.25.25,0,0,0,.25-.25Z"
					fill="currentColor"
				></path>
				<path
					d="M22.75,11.06h-10a1,1,0,0,0-1,1v6a1,1,0,0,0,1,1h4.5l3.67,3.26a.5.5,0,0,0,.83-.38V19.06h1a1,1,0,0,0,1-1v-6A1,1,0,0,0,22.75,11.06Z"
					fill="currentColor"
				></path>
			</g>
		</svg>
	);
};

export default Conversation;
