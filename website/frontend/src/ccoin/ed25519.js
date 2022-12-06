import { sha512 } from "@noble/hashes/sha512";
import * as ed25519 from "@noble/ed25519";

ed25519.utils.sha512Sync = (...m) => sha512(ed25519.utils.concatBytes(...m));

export const generatePrivateKey = ed25519.utils.randomPrivateKey;
export const generateKeypair = () => {
	const privateScalar = ed25519.utils.randomPrivateKey();
	const publicKey = getPublicKey(privateScalar);
	const secretKey = new Uint8Array(64);
	secretKey.set(privateScalar);
	secretKey.set(publicKey, 32);
	return {
		publicKey,
		secretKey,
	};
};

export const getPublicKey = ed25519.sync.getPublicKey;

export const sign = (message, secretKey) =>
	ed25519.sync.sign(message, secretKey.slice(0, 32));

export const verify = ed25519.sync.verify;
