package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IP struct {
	Query string
}

func GetSelfPeer() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}

	var ip IP
	json.Unmarshal(body, &ip)

	return ip.Query
}

func RegisterAsPeer(peers []string) error {
	ip := GetSelfPeer()

	for _, peer := range peers {
		if peer == ip {
			continue
		}
		baseUrl := fmt.Sprintf("http://%s:%d", peer, NodePort)
		resp, err := http.Post(fmt.Sprintf("%s/peer/%s", baseUrl, ip), "", nil)
		if err != nil {
			return err
		}

		fmt.Println(resp.StatusCode)

		if resp.StatusCode != 200 {
			return errors.New("failed")
		}
	}

	return nil
}

func ReplacePeers(peers []string) error {
	j, err := json.MarshalIndent(peers, "", "    ")
	if err != nil {
		return err
	}

	WritePeer(GetSelfPeer())

	return ioutil.WriteFile("./peers.json", j, 0644)
}

func ReadPeers() ([]string, error) {
	var peers []string

	f, err := ioutil.ReadFile("./peers.json")
	if err != nil {
		return []string{}, err
	}

	err = json.Unmarshal(f, &peers)
	return peers, err
}

func WritePeer(peer string) error {
	peers, err := ReadPeers()
	if err != nil {
		return err
	}

	for _, p := range peers {
		if p == peer {
			return errors.New("peer already registered")
		}
	}

	peers = append(peers, peer)
	j, err := json.MarshalIndent(peers, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile("./peers.json", j, 0644)
}
