package server

import (
	"math/big"
	"net/http"

	"github.com/caoxuwen/go/services/bifrost/bitcoin"
	"github.com/caoxuwen/go/services/bifrost/config"
	"github.com/caoxuwen/go/services/bifrost/database"
	"github.com/caoxuwen/go/services/bifrost/ethereum"
	"github.com/caoxuwen/go/services/bifrost/queue"
	"github.com/caoxuwen/go/services/bifrost/sse"
	"github.com/caoxuwen/go/services/bifrost/stellar"
	"github.com/caoxuwen/go/support/log"
)

// ProtocolVersion is the version of the protocol that Bifrost server and
// JS SDK use to communicate.
const ProtocolVersion int = 2

type Server struct {
	BitcoinListener            *bitcoin.Listener            `inject:""`
	BitcoinAddressGenerator    *bitcoin.AddressGenerator    `inject:""`
	Config                     *config.Config               `inject:""`
	Database                   database.Database            `inject:""`
	EthereumListener           *ethereum.Listener           `inject:""`
	EthereumAddressGenerator   *ethereum.AddressGenerator   `inject:""`
	StellarAccountConfigurator *stellar.AccountConfigurator `inject:""`
	TransactionsQueue          queue.Queue                  `inject:""`
	SSEServer                  sse.ServerInterface          `inject:""`

	MinimumValueBtc string
	MinimumValueEth string
	SignerPublicKey string

	minimumValueSat int64
	minimumValueWei *big.Int
	httpServer      *http.Server
	log             *log.Entry
}

type GenerateAddressResponse struct {
	ProtocolVersion int    `json:"protocol_version"`
	Chain           string `json:"chain"`
	Address         string `json:"address"`
	Signer          string `json:"signer"`
}
