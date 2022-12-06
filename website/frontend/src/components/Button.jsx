import { useState, useEffect, useRef } from "react";
import Spinner from "../svg/Spinner";

const Button = ({
	className,
	name,
	icon,
	onClickControl,
	onClick,
	disabled,
	delay,
	spinner,
}) => {
	const [processing, setProcessing] = useState(false);
	const componentIsMounted = useRef(true);
	useEffect(() => {
		return () => {
			componentIsMounted.current = false;
		};
	}, []);

	return (
		<button
			className={
				`
            relative 
			before:bg-light before:transition-all before:absolute before:w-full before:h-full before:opacity-0 
			 before:rounded-lg
            flex items-center justify-center py-2 px-4 rounded-lg font-semibold ${className} ` +
				(() => {
					if (disabled) {
						return " !bg-mute cursor-not-allowed";
					}
					return "hover:before:opacity-5";
				})()
			}
			disabled={disabled || processing}
			onClick={async (e) => {
				e.stopPropagation();
				e.preventDefault();
				setProcessing(true);
				const passed = await onClickControl();
				if (!passed) {
					if (componentIsMounted.current) {
						setProcessing(false);
					}
					return;
				}
				await setTimeout(async () => {
					await onClick();
					setProcessing(false);
				}, delay);
			}}
		>
			{spinner && processing ? (
				<span className="text-light">
					<Spinner height="24px" />
				</span>
			) : (
				<>
					{icon ? icon : null}
					<span
						className={
							`font-semibold ` +
							(() => {
								if (icon) {
									return "ml-1.5";
								}
								return "";
							})()
						}
					>
						{name}
					</span>
				</>
			)}
		</button>
	);
};

Button.defaultProps = {
	name: "",
	delay: 0,
	className: "h-11 bg-vgray-700 text-light",
	onClickControl: () => {
		return true;
	},
	onClick: () => {
		console.log("click");
	},
	spinner: true,
};

export default Button;
