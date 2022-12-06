import React from "react";

const Dice = ({ height, width }) => {
	return (
		<svg
			height={height}
			width={width}
			xmlns="http://www.w3.org/2000/svg"
			viewBox="0 0 24 24"
		>
			<g>
				<path
					d="M22,15.5H17.5a1,1,0,0,1,0-2h4.25a.25.25,0,0,0,.25-.25v-11A.25.25,0,0,0,21.75,2H11.25a.25.25,0,0,0-.25.25V6.5a1,1,0,0,1-2,0V2a2,2,0,0,1,2-2H22a2,2,0,0,1,2,2V13.5A2,2,0,0,1,22,15.5Z"
					fill="currentColor"
				></path>
				<circle cx="19" cy="10.5" r="1" fill="currentColor"></circle>
				<circle cx="19" cy="5" r="1" fill="currentColor"></circle>
				<circle cx="14.5" cy="5" r="1" fill="currentColor"></circle>
				<path
					d="M13.5,9H1.5A1.5,1.5,0,0,0,0,10.5v12A1.5,1.5,0,0,0,1.5,24h12A1.5,1.5,0,0,0,15,22.5v-12A1.5,1.5,0,0,0,13.5,9ZM3.5,20a1,1,0,1,1,1-1A1,1,0,0,1,3.5,20Zm0-5a1,1,0,1,1,1-1A1,1,0,0,1,3.5,15Zm4,5a1,1,0,1,1,1-1A1,1,0,0,1,7.5,20Zm0-5a1,1,0,1,1,1-1A1,1,0,0,1,7.5,15Zm4,5a1,1,0,1,1,1-1A1,1,0,0,1,11.5,20Zm0-5a1,1,0,1,1,1-1A1,1,0,0,1,11.5,15Z"
					fill="currentColor"
				></path>
			</g>
		</svg>
	);
};

export default Dice;
