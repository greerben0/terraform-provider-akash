package akash

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

type TerraformAkashClient struct {
	Network        string
	Home           string
	AkashVersion   string
	AkashChainID   string
	AkashNode      string
	KeyringBackend string
	LogLevel       string
	DefaultFees    string
}

func NewAkashClient(network, home, version, chainID, keyringBackend *string) (*TerraformAkashClient, error) {
	c := TerraformAkashClient{}

	if network == nil {
		c.Network = "mainnet"
	} else {
		c.Network = *network
	}

	akashNet := "https://raw.githubusercontent.com/ovrclk/net/master/" + c.Network

	if home == nil {
		home, err := os.UserHomeDir()

		if err != nil {
			return nil, err
		}
		c.Home = home
	} else {
		c.Home = *home
	}

	if version == nil {
		version, err := getVersion(akashNet)

		if err != nil {
			return nil, err
		}

		c.AkashVersion = version
	} else {
		c.AkashVersion = *version
	}

	if &chainID == nil {
		chainID, err := getChainID(akashNet)

		if err != nil {
			return nil, err
		}

		c.AkashChainID = chainID
	} else {
		c.AkashChainID = *chainID
	}

	if keyringBackend == nil {
		c.KeyringBackend = "os"
	} else {
		c.KeyringBackend = *keyringBackend
	}

	akashNode, err := getAkashNode(akashNet)

	if err != nil {
		return nil, err
	}

	c.AkashNode = akashNode

	c.LogLevel = "info"
	c.DefaultFees = "5000uakt"

	log.Println(c)

	return &c, nil
}

func getVersion(akashNet string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", akashNet, "version.txt"))
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	return string(body), nil
}

func getChainID(akashNet string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", akashNet, "chain-id.txt"))
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	chainID := string(body)

	return chainID, nil
}

func getAkashNode(akashNet string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", akashNet, "rpc-nodes.txt"))
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	nodes := strings.Split(string(body), "\n")

	randomIndex := rand.Intn(len(nodes)) - 1
	log.Println(randomIndex)
	akashNode := nodes[randomIndex]

	return akashNode, nil
}

func (c *TerraformAkashClient) GetKey(name string) (Key, error) {
	//TODO Implement
	key := Key{
		Name:    name,
		Type:    "local",
		Address: "akash123",
		Pubkey:  "akashpub123",
	}

	return key, nil
}

func (c *TerraformAkashClient) GetAccount(address string) (Account, error) {
	// TODO Implement
	account := Account{
		Address: address,
		Amount:  "1234567890",
		Denom:   "uakt",
	}

	return account, nil
}
