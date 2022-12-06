import React from "react";

const Hearts = ({ height, width }) => {
	return (
		<svg
			xmlns="http://www.w3.org/2000/svg"
			height={height}
			width={width}
			fill="none"
			viewBox="0 0 14 14"
		>
			<path
				fill="currentColor"
				d="M7.51404 12.4771C7.19881 12.6722 6.80129 12.6722 6.48607 12.4771C3.98442 10.9292 0.391846 8.35487 0.391846 5.19604C0.391846 2.99432 2.17669 1.20947 4.37841 1.20947C5.38213 1.20947 6.29921 1.58041 7.00005 2.19268C7.70089 1.58041 8.61798 1.20947 9.6217 1.20947C11.8234 1.20947 13.6083 2.99432 13.6083 5.19604C13.6083 8.35487 10.0157 10.9292 7.51404 12.4771Z"
			></path>
		</svg>
	);
};

export default Hearts;
