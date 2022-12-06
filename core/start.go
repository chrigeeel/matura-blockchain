package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func StartCore() {
	StartNodeServer()
	StartRPCServer()

	time.Sleep(time.Second * 3)

	j, err := ioutil.ReadFile("./chain.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(j, &ThisChain)

	peers, err := ReadPeers()
	if err != nil {
		panic(err)
	}

	if len(peers) == 0 {
		peers = PermanentPeers
	}

	func() {
		for _, peer := range peers {
			fmt.Println(peer)

			chain, err := FetchPeerChain(peer)
			if err != nil {
				fmt.Println(err)
				continue
			}

			mempool, err := FetchPeerMempool(peer)
			if err != nil {
				fmt.Println(err)
				continue
			}

			peers, err := FetchPeerPeers(peer)
			if err != nil {
				fmt.Println(err)
				continue
			}

			if len(peers) == 0 {
				peers = PermanentPeers
			}

			err = ReplacePeers(peers)
			if err != nil {
				fmt.Println(err)
				continue
			}

			peers = append(peers, GetSelfPeer())

			ThisChain = chain
			WriteChain()

			Mempool = mempool

			err = RegisterAsPeer(peers)
			if err != nil {
				fmt.Println(err)
			}
			return
		}
		panic("no peer was able to provide startup info")
	}()
}

func FetchPeerChain(peer string) (Chain, error) {
	baseUrl := fmt.Sprintf("http://%s:%d", peer, NodePort)
	resp, err := http.Get(fmt.Sprintf("%s/chain", baseUrl))
	if err != nil {
		return Chain{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Chain{}, err
	}

	var chain Chain
	err = json.Unmarshal(body, &chain)
	if err != nil {
		return Chain{}, err
	}

	return chain, nil
}

func FetchPeerMempool(peer string) ([]Transaction, error) {
	baseUrl := fmt.Sprintf("http://%s:%d", peer, NodePort)
	resp, err := http.Get(fmt.Sprintf("%s/mempool", baseUrl))
	if err != nil {
		return []Transaction{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Transaction{}, err
	}
	var mempool []Transaction
	err = json.Unmarshal(body, &mempool)
	if err != nil {
		return []Transaction{}, err
	}

	return mempool, nil
}

func FetchPeerPeers(peer string) ([]string, error) {
	baseUrl := fmt.Sprintf("http://%s:%d", peer, NodePort)
	resp, err := http.Get(fmt.Sprintf("%s/peers", baseUrl))
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}
	var peers []string
	err = json.Unmarshal(body, &peers)
	if err != nil {
		return []string{}, err
	}

	return peers, nil
}
