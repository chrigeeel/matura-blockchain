export const shortenAddress = (address, chars = 4) => {
	return `${address.slice(0, chars)}...${address.slice(-chars)}`;
};

export const formatCcoin = (fractions) => {
	return `${Math.round((fractions / 1_000_000_000) * 1_000) / 1_000} CCOIN`;
};
