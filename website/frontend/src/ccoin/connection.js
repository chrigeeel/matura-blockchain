import axios from "axios";

export class Connection {
	_url;

	constructor(url) {
		this._url = url;
	}

	async getBalance(publicKey) {
		const path = `${this._url}/balance/${publicKey.toBase58()}`;
		const resp = await axios.get(path);
		return resp.data;
	}

	async getTransaction(signature) {
		const path = `${this._url}/transaction/${signature.toBase58()}`;
		const resp = await axios.get(path);
		return resp.data;
	}

	async sendTransaction(transaction) {
		const path = `${this._url}/transaction`;
		const resp = await axios.post(path, transaction);
		return resp.data;
	}
}
