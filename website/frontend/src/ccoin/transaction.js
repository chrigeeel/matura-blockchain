import { sha256 } from "@noble/hashes/sha256";

import { sign } from "./ed25519";
import { Signature } from "./signature";

export class Transaction {
	_signature;
	_data;

	constructor(data) {
		data.nonce = randomNonce();
		this._data = data;
	}

	sign(keypair) {
		console.log(keypair);
		const signature = new Signature(
			sign(this.dataToSign(), keypair.secretKey)
		);
		this._signature = signature;

		return signature;
	}

	toJSON() {
		return {
			data: this._data,
			signature: this._signature,
		};
	}

	dataToSign() {
		return sha256(this.marshalData());
	}

	marshalData() {
		return JSON.stringify(this._data);
	}
}

const randomNonce = () => {
	const minimum = 0;
	const maximum = 4294967295;

	return Math.floor(Math.random() * (maximum - minimum + 1)) + minimum;
};
