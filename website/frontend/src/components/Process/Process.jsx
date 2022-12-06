import React from "react";

const Process = () => {
	return (
		<div className="flex flex-col m-5 mx-auto max-w-5xl">
			<h2 className="font-bold text-2xl sm:text-4xl md:text-6xl underline underline-offset-4">
				Development process
			</h2>
			<div className="flex flex-col text-sm md:text-base">
				<p className="mt-6">
					First, a few words about my programming carreer so far. I
					have been developing software in different programming
					languages for over two years. I started out learning{" "}
					<span className="text-green font-semibold">Python</span>.
					Later, I started learning the newer language,{" "}
					<span className="text-green font-semibold">Golang</span>,
					which is designed for high performance, concurrency-based
					backend applications. I am in the crypto space myself and
					have developed other small apps incorporating blockchain
					technologies on commission.
				</p>
				<p className="mt-6">
					Now, onto the development process of my own cryptocurrency -
					ccoin.
					<br />
					Since the topic of my matura paper is the technology behind
					cryptocurrencies, I already had quite a good understanding
					of how things worked. I spent many hours researching for my
					matura paper and I was now able to directly convert that new
					knowledge into a coding project. I spent around{" "}
					<span className="text-green font-semibold">8 hours</span>
					writing the basic infrastructure surrounding keypairs and
					signatures. Of course, I was able to use{" "}
					<span className="text-green font-semibold">
						exisiting libraries
					</span>{" "}
					for elliptic curve implementations, but standardizing still
					needs to be done. The libraries however do the heavy
					mathematical lifting.
					<br />I then spent an additional{" "}
					<span className="text-green font-semibold">
						20-30 hours
					</span>{" "}
					building on that foundation I built. I decided to go with
					the proof of work consensus mechanism. This stage of
					development mostly consisted of implementing new features,
					causing bugs and then trying to fix those bugs. This stage
					was sometimes extremely frustrating and infurating, since I
					mostly was on my own and was not able to find a lot of help
					on the internet. After a few weeks of scattered work, I was
					finally done with the backend and most of the development
					process - or so I thought. The codebase up to this point was{" "}
					<span className="text-red font-semibold">
						over 1300 lines
					</span>{" "}
					long, only accounting for completely custom code, written
					solely by me.
				</p>
				<p className="mt-6">
					Next, I had to{" "}
					<span className="text-green font-semibold">deploy</span> the
					cryptocurrency somewhere. I chose Amazon Web Services since
					I am already familiar with their service. I spun up three
					small 2-core servers and deployed the source code I wrote on
					each of them. I started them one-by-one, and, of course,
					nothing seemed to work at first. What worked on my local PC
					did not work on the Linux servers and I spent quite a few
					hours figuring out what was wrong. After that additional
					time, I finally got the blockchain running on three nodes,
					controlled by me.
				</p>
				<p className="mt-6">
					Now, I ran into a problem. The blockchain was indeed running
					on those three servers, but who cares about a software
					running on three servers if one cannot interact with it. I
					decided the best way to make my blockchain usable was
					through a{" "}
					<span className="text-green font-semibold">website</span>{" "}
					with as little setting up required by the user as possible.
					Developing this website was not an easy task. I had to write
					all of the beautiful keypair and signature foundation I
					wrote in Golang again, in{" "}
					<span className="text-green font-semibold">Javascript</span>
					. Next, I made an extremely simple Demo page. On this demo
					page, a user gets credited 0.1 CCOIN by the network just for
					creating a keypair. This rule was decided by a majority vote
					in the network, meaning that over 50% of the people running
					nodes agreed with this change. That might sound impressive,
					but it actually was just me running those three nodes so I
					simply had to agree with myself and update the node's code.
					If more people had been running a node, this change would
					have been a lot harder. Consensus has to be reached to make
					a change. The longer chain is valid.
				</p>
				<p className="mt-6">
					And now, here we are. This was a very simplified rundown of
					the process of me creating my own blockchain. I did not go
					into technical detail, but rather wanted to share a bit of
					the experience.
				</p>
			</div>
		</div>
	);
};

export default Process;
