import React, { useEffect, useState } from "react";
import bs58 from "bs58";
import { Keypair } from "../../ccoin/keypair";
import { formatCcoin, shortenAddress } from "../../ccoin/fmt";
import Decentralised from "../../svg/Decentralised";
import Button from "../Button";
import { Connection } from "../../ccoin/connection";
import { copyTextToClipboard } from "../../helpers/helpers";
import Send from "./Send";

const connection = new Connection(process.env.REACT_APP_RPC_URL);

const Blockchain = () => {
	const [keypair, setKeypair] = useState();
	const [balance, setBalance] = useState(0);

	useEffect(() => {
		const secretKey = localStorage.getItem("secretKey");
		try {
			setKeypair(Keypair.fromBase58(secretKey));
		} catch {
			setKeypair(Keypair.generate());
		}
	}, []);

	useEffect(() => {
		if (!keypair) {
			return;
		}

		localStorage.setItem("secretKey", bs58.encode(keypair.secretKey));
		console.log(keypair.publicKey.toBase58());

		(async () => {
			setBalance(
				(await connection.getBalance(keypair.publicKey)).balance
			);
		})();
		const interval = setInterval(async () => {
			setBalance(
				(await connection.getBalance(keypair.publicKey)).balance
			);
		}, 2000);

		return () => {
			clearInterval(interval);
		};
	}, [keypair]);

	return (
		<div className="flex flex-col m-5 mx-auto max-w-5xl">
			<h2 className="font-bold text-2xl sm:text-4xl md:text-6xl underline underline-offset-4">
				My own Blockchain - ccoin
			</h2>
			<div className="flex flex-col text-sm md:text-base">
				<p className="mt-6">
					For my matura project 2022, I developed a{" "}
					<span className="text-green font-semibold">
						cryptocurrency
					</span>{" "}
					with{" "}
					<span className="text-green font-semibold">
						blockchain technologies
					</span>{" "}
					by myself and decided to call it "
					<span className="text-green font-semibold">ccoin</span>
					".
				</p>
				<p className="mt-2">
					One of the{" "}
					<span className="text-red font-semibold">
						biggest obstacles
					</span>{" "}
					to overcome was however not the development of the
					blockchain itself, but making it accessible for anyone to
					use through this website. Using my blockchain should feel
					intuitive and one should not require any extensive
					technological knowledge to make use of it.
				</p>
				<p className="mt-2">
					To solve this user-experience problem, I decided to create
					this easy-to-use website. <br />
					Anyone can create an "account" on the blockchain by clicking
					the buttons below. Upon account creation, one instantly owns{" "}
					<span className="text-green font-semibold">0.1 ccoin</span>!
					Your private keys are never shared with any server, but
					rather stored in your browser's local storage.
				</p>
			</div>

			<div className="w-full h-[2px] bg-mute bg-opacity-70 my-4 md:my-6"></div>

			{!keypair ? (
				<div>Loading...</div>
			) : (
				<div className="flex flex-col">
					<div className="flex items-center">
						<span className="font-semibold text-lg md:text-xl">
							Your account:
						</span>
						<Button
							name="Generate new account"
							className="bg-red ml-auto text-sm"
							onClick={() => {
								setKeypair(Keypair.generate());
							}}
						/>
					</div>
					<div className="flex flex-col mt-6">
						<span className="text-mute-active-2 text-sm">
							Your secret key:
						</span>
						<button
							onClick={() => {
								copyTextToClipboard(
									bs58.encode(keypair.secretKey)
								);
							}}
							className="mt-2 mr-auto"
						>
							<span
								className="md:hidden rounded-lg px-4 py-2
                             bg-vgray-300 font-semibold text-lg 
                             whitespace-nowrap overflow-ellipsis overflow-hidden"
							>
								{shortenAddress(
									bs58.encode(keypair.secretKey),
									12
								)}
							</span>
							<span
								className="hidden md:block lg:hidden rounded-lg px-4 py-2
                             bg-vgray-300 font-semibold text-lg 
                             whitespace-nowrap overflow-ellipsis overflow-hidden"
							>
								{shortenAddress(
									bs58.encode(keypair.secretKey),
									24
								)}
							</span>
							<span
								className="hidden lg:block rounded-lg px-4 py-2
                             bg-vgray-300 font-semibold text-lg 
                             whitespace-nowrap overflow-ellipsis overflow-hidden"
							>
								{shortenAddress(
									bs58.encode(keypair.secretKey),
									32
								)}
							</span>
						</button>
					</div>
					<div className="flex flex-col mt-6">
						<span className="text-mute-active-2 text-sm">
							Your public key:
						</span>
						<button
							onClick={() => {
								copyTextToClipboard(
									keypair.publicKey.toBase58()
								);
							}}
							className="mt-2 mr-auto"
						>
							<span className="hidden md:block rounded-lg px-4 py-2 bg-vgray-300 font-semibold text-lg">
								{keypair.publicKey.toBase58()}
							</span>
							<span className="block md:hidden rounded-lg px-4 py-2 bg-vgray-300 font-semibold text-base">
								{shortenAddress(
									keypair.publicKey.toBase58(),
									12
								)}
							</span>
						</button>
					</div>

					<div className="flex flex-col mt-10">
						<span className="text-base font-medium">
							Your balance:
						</span>
						<span className="text-3xl text-green font-bold mt-1">
							{formatCcoin(balance)}
						</span>
					</div>
					<div className="mt-6">
						<Send balance={balance} keypair={keypair} />
					</div>
				</div>
			)}
		</div>
	);
};

export default Blockchain;
