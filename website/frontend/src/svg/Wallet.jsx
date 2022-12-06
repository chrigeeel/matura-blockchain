import React from "react";

const Wallet = ({ height, width }) => {
	return (
		<svg
			height={height}
			width={width}
			xmlns="http://www.w3.org/2000/svg"
			viewBox="0 0 24 24"
		>
			<g>
				<path
					d="M19.62,18.5a3.75,3.75,0,0,1,0-7.5h2.26a.25.25,0,0,0,.24-.25V8.25a2.5,2.5,0,0,0-2.5-2.5h-17a2.51,2.51,0,0,0-2.5,2.5v13a2.51,2.51,0,0,0,2.5,2.5h17a2.5,2.5,0,0,0,2.5-2.5v-2.5a.25.25,0,0,0-.24-.25Z"
					fill="currentColor"
				></path>
				<path
					d="M22.62,12.5h-3a2.25,2.25,0,0,0,0,4.5h3a1.51,1.51,0,0,0,1.26-1.5V14A1.51,1.51,0,0,0,22.62,12.5Z"
					fill="currentColor"
				></path>
				<path
					d="M19.81,1.18a1.21,1.21,0,0,0-.56-.76,1.24,1.24,0,0,0-.94-.13L5.44,3.76A.24.24,0,0,0,5.26,4a.25.25,0,0,0,.24.22l14.77,0a.24.24,0,0,0,.19-.1A.23.23,0,0,0,20.51,4Z"
					fill="currentColor"
				></path>
			</g>
		</svg>
	);
};

export default Wallet;
