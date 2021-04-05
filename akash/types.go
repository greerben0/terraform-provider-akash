package akash

type Key struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Address string `json:"address"`
	Pubkey  string `json:"pubkey"`
}

type Account struct {
	Address string
	Amount  string
	Denom   string
}
