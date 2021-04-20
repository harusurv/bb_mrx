// +build unittest

package metrix

import (
	"encoding/hex"
	"math/big"
	"os"
	"reflect"
	"testing"

	"github.com/martinboehm/btcutil/chaincfg"
	"github.com/trezor/blockbook/bchain"
	"github.com/trezor/blockbook/bchain/coins/btc"
)

func TestMain(m *testing.M) {
	c := m.Run()
	chaincfg.ResetParams()
	os.Exit(c)
}

func Test_GetAddrDescFromAddress_Mainnet(t *testing.T) {

}

func Test_GetAddressesFromAddrDesc(t *testing.T) {

}

func Test_PackTx(t *testing.T) {

}

func Test_UnpackTx(t *testing.T) {

}
