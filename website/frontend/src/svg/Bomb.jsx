import React from "react";

const Bomb = ({ height, width }) => {
	return (
		<svg
			height={height}
			width={width}
			xmlns="http://www.w3.org/2000/svg"
			viewBox="0 0 24 24"
		>
			<g>
				<path
					d="M14.38,7.6a1,1,0,0,1-.63-.93V6a1,1,0,0,0-1-1h-3a1,1,0,0,0-1,1v.67a1,1,0,0,1-.63.93,8.5,8.5,0,1,0,6.26,0Z"
					fill="currentColor"
				></path>
				<path
					d="M20.25,2.5a1,1,0,0,0-1,1,1,1,0,0,1-2,0,3.5,3.5,0,0,0-6.95-.58.46.46,0,0,0,.11.4.49.49,0,0,0,.38.18h1a.5.5,0,0,0,.49-.38,1.5,1.5,0,0,1,2.95.38,3,3,0,0,0,6,0A1,1,0,0,0,20.25,2.5Z"
					fill="currentColor"
				></path>
			</g>
		</svg>
	);
};

export default Bomb;
