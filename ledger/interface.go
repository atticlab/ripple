package ledger

import (
	"github.com/atticlab/ripple/data"
)

type Sync interface {
	Current(uint32)
	Missing(*data.LedgerRange) *data.Work
	Submit([]data.Hashable)
	Copy() *RadixMap
}
