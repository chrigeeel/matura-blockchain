import React from "react";

const Hamburger = ({ height, width }) => {
	return (
		<svg
			height={height}
			width={width}
			xmlns="http://www.w3.org/2000/svg"
			viewBox="0 0 24 24"
		>
			<g>
				<rect
					x="3"
					y="3"
					width="18"
					height="4"
					rx="1.5"
					fill="currentColor"
				></rect>
				<rect
					x="3"
					y="10"
					width="18"
					height="4"
					rx="1.5"
					fill="currentColor"
				></rect>
				<rect
					x="3"
					y="17"
					width="18"
					height="4"
					rx="1.5"
					fill="currentColor"
				></rect>
			</g>
		</svg>
	);
};

export default Hamburger;
