import React from "react";

const Card = ({ height, width }) => {
	return (
		<svg
			height={height}
			width={width}
			xmlns="http://www.w3.org/2000/svg"
			viewBox="0 0 24 24"
		>
			<g>
				<path
					d="M15.17,11.4,12.8,8.23a1,1,0,0,0-1.6,0L8.83,11.4a1,1,0,0,0,0,1.2l2.37,3.17a1,1,0,0,0,1.6,0l2.37-3.17A1,1,0,0,0,15.17,11.4Z"
					fill="currentColor"
				></path>
				<path
					d="M18.5,24H5.5a3,3,0,0,1-3-3V3a3,3,0,0,1,3-3h13a3,3,0,0,1,3,3V21A3,3,0,0,1,18.5,24ZM5.5,2a1,1,0,0,0-1,1V21a1,1,0,0,0,1,1h13a1,1,0,0,0,1-1V3a1,1,0,0,0-1-1Z"
					fill="currentColor"
				></path>
				<circle
					cx="7.25"
					cy="4.75"
					r="1.25"
					fill="currentColor"
				></circle>
				<circle
					cx="16.75"
					cy="19.25"
					r="1.25"
					fill="currentColor"
				></circle>
			</g>
		</svg>
	);
};

export default Card;
