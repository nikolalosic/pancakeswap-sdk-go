package constants

// ChainID network chain id
type ChainID int

//go:generate stringer -type=ChainID -linecomment
const (
	Mainnet ChainID = 56
	Testnet ChainID = 97
)
