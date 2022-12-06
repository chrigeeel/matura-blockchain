import BN from "bn.js";
import bs58 from "bs58";
import { generateKeypair, generatePrivateKey, getPublicKey } from "./ed25519";
import { Buffer } from "buffer";

export const PUBLIC_KEY_LENGTH = 32;

export class PublicKey {
	_bn;

	constructor(value) {
		if (isPublicKeyData(value)) {
			this._bn = value._bn;
		} else {
			if (typeof value === "string") {
				const decoded = bs58.decode(value);
				if (decoded.length != PUBLIC_KEY_LENGTH) {
					throw new Error("Invalid public key length");
				}
				this._bn = new BN(decoded);
			} else {
				this._bn = new BN(value);
			}

			if (this._bn.byteLength() > PUBLIC_KEY_LENGTH) {
				throw new Error("Invalid public key length");
			}
		}
	}

	toBase58() {
		return bs58.encode(this.toBytes());
	}

	toJSON() {
		return this.toBase58();
	}

	toBytes() {
		return this.toBuffer();
	}

	toBuffer() {
		const b = this._bn.toArrayLike(Buffer);
		if (b.length === PUBLIC_KEY_LENGTH) {
			return b;
		}

		const zeroPad = Buffer.alloc(32);
		b.copy(zeroPad, 32 - b.length);
		return zeroPad;
	}
}

function isPublicKeyData(value) {
	return value._bn !== undefined;
}

export class Keypair {
	_keypair;

	constructor(keypair) {
		this._keypair = keypair ?? generateKeypair();
	}

	static generate() {
		return new Keypair(generateKeypair());
	}

	static fromBase58(value) {
		const decoded = bs58.decode(value);
		return this.fromSecretKey(decoded);
	}

	static fromSecretKey(secretKey, options) {
		if (secretKey.byteLength !== 64) {
			throw new Error("bad secret key size");
		}
		const publicKey = secretKey.slice(32, 64);
		if (!options || !options.skipValidation) {
			const privateScalar = secretKey.slice(0, 32);
			const computedPublicKey = getPublicKey(privateScalar);
			for (let ii = 0; ii < 32; ii++) {
				if (publicKey[ii] !== computedPublicKey[ii]) {
					throw new Error("provided secretKey is invalid");
				}
			}
		}
		return new Keypair({ publicKey, secretKey });
	}

	get publicKey() {
		return new PublicKey(this._keypair.publicKey);
	}

	get secretKey() {
		return new Uint8Array(this._keypair.secretKey);
	}
}
