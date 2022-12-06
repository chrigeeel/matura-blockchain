import React from "react";

const Check = ({ height, width }) => {
	return (
		<svg
			height={height}
			width={width}
			xmlns="http://www.w3.org/2000/svg"
			viewBox="0 0 24 24"
		>
			<g>
				<path
					d="M9.22,14.79a1.24,1.24,0,0,0,1.74,0l7.11-6.69a1.25,1.25,0,0,0-1.72-1.82l-6.21,5.86L8.63,10.64A1.25,1.25,0,1,0,6.85,12.4Z"
					fill="currentColor"
				></path>
				<path
					d="M20.05,0H4A2,2,0,0,0,2,2V8a18,18,0,0,0,9.53,15.93,1,1,0,0,0,.94,0A18,18,0,0,0,22,8V2A2,2,0,0,0,20.05,0ZM20,8a16,16,0,0,1-8,13.91A16,16,0,0,1,4,8V2.5A.5.5,0,0,1,4.5,2l15,0a.5.5,0,0,1,.5.5Z"
					fill="currentColor"
				></path>
			</g>
		</svg>
	);
};

export default Check;
