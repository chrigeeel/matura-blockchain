import React, { useState } from "react";
import { Connection } from "../../ccoin/connection";
import { PublicKey } from "../../ccoin/keypair";
import { Transaction } from "../../ccoin/transaction";
import Button from "../Button";

const connection = new Connection(process.env.REACT_APP_RPC_URL);

const Send = ({ keypair, balance }) => {
	const [toSend, setToSend] = useState(0);
	const [receiver, setReceiver] = useState("");
	const [error, setError] = useState("");
	const [success, setSuccess] = useState("");

	const [signature, setSignature] = useState();

	const send = () => {
		setError("");
		const amount = toSend * 1_000_000_000;
		if (balance < amount) {
			setError("Balance not high enough!");
			return;
		}

		let toAccount;
		try {
			toAccount = new PublicKey(receiver);
		} catch {
			setError("Invalid public key!");
			return;
		}

		const transaction = new Transaction({
			sender: keypair.publicKey,
			receiver: toAccount,
			amount,
		});

		try {
			transaction.sign(keypair);
		} catch {
			setError("Signature failed!");
			return;
		}

		try {
			const resp = connection.sendTransaction(transaction);
			setSuccess("Successfully sent transaction!");

			setSignature(transaction._signature);
		} catch {
			setError("Could not send transaction!");
			return;
		}
	};

	return (
		<div className="flex">
			<div className="flex flex-col w-full">
				<span className="text-sm text-mute-active-2">
					Amount to send:
				</span>
				<input
					placeholder={"0.01 CCOIN"}
					className="bg-vgray-700 rounded-lg text-sm text-light py-2 px-3
					focus:ring-2 transition ring-mute-active ring-opacity-50
					focus:outline-none mt-0.5 w-full md:w-64"
					type="number"
					value={toSend}
					onChange={(e) => {
						setToSend(e.target.value);
					}}
				/>

				<span className="text-sm text-mute-active-2 mt-4">
					Destination public key:
				</span>
				<input
					placeholder={"Public key"}
					className="bg-vgray-700 rounded-lg text-sm text-light py-2 px-3
					focus:ring-2 transition ring-mute-active ring-opacity-50
					focus:outline-none mt-0.5 w-full md:w-64"
					type="text"
					value={receiver}
					onChange={(e) => {
						setReceiver(e.target.value);
					}}
				/>
				{error ? (
					<span className="text-red font-semibold mt-2 text-sm">
						{error}
					</span>
				) : null}
				{success ? (
					<span className="text-green font-semibold mt-2 text-sm">
						{success}
					</span>
				) : null}

				<Button
					onClick={() => {
						send();
					}}
					delay={200}
					name="Send!"
					className="mt-6 bg-green"
				/>
			</div>
		</div>
	);
};

export default Send;
