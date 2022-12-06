const TickUp = ({ width, height, color }) => {
	return (
		<svg
			xmlns="http://www.w3.org/2000/svg"
			width={width}
			height={height}
			viewBox="0 0 5.373 3.134"
			className={color}
		>
			<path
				data-name="Icon color"
				d="M5.28,2.469a.313.313,0,0,1,0,.445l-.125.125a.307.307,0,0,1-.226.094H.444A.307.307,0,0,1,.218,3.04L.093,2.914a.313.313,0,0,1,0-.445L2.467.095a.3.3,0,0,1,.439,0Z"
				fill="currentColor"
			/>
		</svg>
	);
};

TickUp.defaultProps = {};

export default TickUp;
