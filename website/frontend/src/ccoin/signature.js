import BN from "bn.js";
import bs58 from "bs58";

export const SIGNATURE_LENGTH = 64;

export class Signature {
	_bn;

	constructor(value) {
		if (typeof value === "string") {
			const decoded = bs58.decode(value);
			if (decoded.length != SIGNATURE_LENGTH) {
				throw new Error("Invalid signature length");
			}
			this._bn = new BN(decoded);
		} else {
			this._bn = new BN(value);
		}

		if (this._bn.byteLength() > SIGNATURE_LENGTH) {
			throw new Error("Invalid signature length");
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
		if (b.length === SIGNATURE_LENGTH) {
			return b;
		}

		const zeroPad = Buffer.alloc(32);
		b.copy(zeroPad, 32 - b.length);
		return zeroPad;
	}
}
