import React from "react";

const Diamonds = ({ height, width }) => {
	return (
		<svg
			height={height}
			width={width}
			xmlns="http://www.w3.org/2000/svg"
			fill="none"
			viewBox="0 0 14 14"
		>
			<path
				fill="#000000"
				d="M2.17805 6.27485C1.91163 6.69144 1.91183 7.28739 2.17737 7.70486C2.78369 8.6581 3.40109 9.6717 4.13436 10.5905C4.86763 11.5092 5.67659 12.2828 6.43738 13.0425C6.77057 13.3752 7.2462 13.3754 7.57869 13.0416C8.33469 12.2826 9.13528 11.5056 9.86565 10.5905C10.596 9.67534 11.2162 8.67224 11.822 7.72499C12.0884 7.3084 12.0882 6.71245 11.8226 6.29498C11.2163 5.34174 10.5989 4.32814 9.86565 3.40939C9.13238 2.49063 8.32341 1.71706 7.56262 0.95736C7.22944 0.624654 6.7538 0.624402 6.42132 0.958209C5.66531 1.71723 4.86472 2.49426 4.13436 3.40938C3.40399 4.3245 2.78383 5.3276 2.17805 6.27485Z"
			></path>
		</svg>
	);
};

export default Diamonds;
