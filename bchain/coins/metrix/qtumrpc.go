package metrix

import (
	"encoding/json"
	"math/big"

	"github.com/golang/glog"
	"github.com/trezor/blockbook/bchain"
	"github.com/trezor/blockbook/bchain/coins/btc"
)

// MetrixRPC is an interface to JSON-RPC bitcoind service.
type MetrixRPC struct {
	*btc.BitcoinRPC
	minFeeRate *big.Int // satoshi per kb
}

// NewMetrixRPC returns new MetrixRPC instance.
func NewMetrixRPC(config json.RawMessage, pushHandler func(bchain.NotificationType)) (bchain.BlockChain, error) {
	b, err := btc.NewBitcoinRPC(config, pushHandler)
	if err != nil {
		return nil, err
	}

	s := &MetrixRPC{
		b.(*btc.BitcoinRPC),
		big.NewInt(400000),
	}
	s.RPCMarshaler = btc.JSONMarshalerV1{}
	s.ChainConfig.SupportsEstimateSmartFee = true

	return s, nil
}

// Initialize initializes MetrixRPC instance.
func (b *MetrixRPC) Initialize() error {
	ci, err := b.GetChainInfo()
	if err != nil {
		return err
	}
	chainName := ci.Chain

	params := GetChainParams(chainName)

	// always create parser
	b.Parser = NewMetrixParser(params, b.ChainConfig)

	// parameters for getInfo request
	if params.Net == MainnetMagic {
		b.Testnet = false
		b.Network = "livenet"
	} else {
		b.Testnet = true
		b.Network = "testnet"
	}

	glog.Info("rpc: block chain ", params.Name)

	return nil
}

// GetTransactionForMempool returns a transaction by the transaction ID
// It could be optimized for mempool, i.e. without block time and confirmations
func (b *MetrixRPC) GetTransactionForMempool(txid string) (*bchain.Tx, error) {
	return b.GetTransaction(txid)
}

// EstimateSmartFee returns fee estimation
func (b *MetrixRPC) EstimateSmartFee(blocks int, conservative bool) (big.Int, error) {
	feeRate, err := b.BitcoinRPC.EstimateSmartFee(blocks, conservative)
	if err != nil {
		if b.minFeeRate.Cmp(&feeRate) == 1 {
			feeRate = *b.minFeeRate
		}
	}
	return feeRate, err
}
