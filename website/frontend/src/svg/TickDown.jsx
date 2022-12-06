const TickDown = ({ width, height, color }) => {
	return (
		<svg
			xmlns="http://www.w3.org/2000/svg"
			width={width}
			height={height}
			viewBox="0 0 4.036 2.354"
			className={`${color} rotate-180`}
		>
			<path
				data-name="Icon color"
				d="M3.967,1.855a.235.235,0,0,1,0,.334l-.094.094a.231.231,0,0,1-.169.071H.333a.231.231,0,0,1-.169-.071L.07,2.189a.235.235,0,0,1,0-.334L1.853.071a.226.226,0,0,1,.329,0Z"
				fill="currentColor"
			/>
		</svg>
	);
};

export default TickDown;
