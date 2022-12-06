import React from "react";

const Lightning = ({ height, width }) => {
	return (
		<svg
			width={width}
			height={height}
			xmlns="http://www.w3.org/2000/svg"
			viewBox="0 0 24 24"
		>
			<path
				d="M20.44,9.27A.48.48,0,0,0,20,9H12.62L14.49.61A.51.51,0,0,0,14.2,0a.5.5,0,0,0-.61.17l-10,14a.49.49,0,0,0,0,.52A.49.49,0,0,0,4,15h7.38L9.51,23.39A.51.51,0,0,0,9.8,24a.52.52,0,0,0,.61-.17l10-14A.49.49,0,0,0,20.44,9.27Z"
				fill="currentColor"
			></path>
		</svg>
	);
};

export default Lightning;
