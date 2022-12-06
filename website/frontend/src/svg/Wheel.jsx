import React from "react";

const Wheel = ({ height, width }) => {
	return (
		<svg
			height={height}
			width={width}
			xmlns="http://www.w3.org/2000/svg"
			viewBox="0 0 24 24"
		>
			<path
				d="M12,0A12,12,0,1,0,24,12,12,12,0,0,0,12,0ZM22,12a9.79,9.79,0,0,1-.35,2.61.23.23,0,0,1-.13.16.28.28,0,0,1-.2,0l-5.67-2.35a.24.24,0,0,1-.15-.24v-.38a.24.24,0,0,1,.15-.24l5.67-2.35a.28.28,0,0,1,.2,0,.23.23,0,0,1,.13.16A9.79,9.79,0,0,1,22,12ZM2,12a9.79,9.79,0,0,1,.35-2.61.23.23,0,0,1,.13-.16.28.28,0,0,1,.2,0l5.67,2.35a.24.24,0,0,1,.15.24v.38a.24.24,0,0,1-.15.24L2.68,14.78a.28.28,0,0,1-.2,0,.23.23,0,0,1-.13-.16A9.79,9.79,0,0,1,2,12ZM20.67,7a.25.25,0,0,1-.12.36L14.88,9.72a.24.24,0,0,1-.28-.06l-.26-.26a.24.24,0,0,1-.06-.28l2.34-5.67a.29.29,0,0,1,.15-.14.3.3,0,0,1,.21,0A10,10,0,0,1,20.67,7ZM14.61,2.35a.23.23,0,0,1,.16.13.28.28,0,0,1,0,.2L12.42,8.37a.22.22,0,0,1-.21.14h-.4a.24.24,0,0,1-.24-.15L9.22,2.68a.28.28,0,0,1,0-.2.23.23,0,0,1,.16-.13,9.91,9.91,0,0,1,5.22,0ZM7,3.33a.3.3,0,0,1,.21,0,.29.29,0,0,1,.15.14L9.72,9.12a.24.24,0,0,1-.06.28l-.26.26a.24.24,0,0,1-.28.06L3.45,7.38A.25.25,0,0,1,3.33,7,10,10,0,0,1,7,3.33ZM3.33,17a.25.25,0,0,1,.12-.36l5.77-2.39a.11.11,0,0,1,.11,0,3.51,3.51,0,0,0,.33.34.24.24,0,0,1,.06.28L7.38,20.55a.29.29,0,0,1-.15.14.3.3,0,0,1-.21,0A10,10,0,0,1,3.33,17Zm6.06,4.67a.23.23,0,0,1-.16-.13.28.28,0,0,1,0-.2l2.35-5.67a.24.24,0,0,1,.24-.15h.4a.22.22,0,0,1,.21.14l2.36,5.69a.28.28,0,0,1,0,.2.23.23,0,0,1-.16.13,9.91,9.91,0,0,1-5.22,0Zm7.59-1a.3.3,0,0,1-.21,0,.29.29,0,0,1-.15-.14l-2.34-5.67a.24.24,0,0,1,.06-.28l.26-.26a.24.24,0,0,1,.28-.06l5.67,2.34a.25.25,0,0,1,.12.36A10,10,0,0,1,17,20.67Z"
				fill="currentColor"
			></path>
		</svg>
	);
};

export default Wheel;
