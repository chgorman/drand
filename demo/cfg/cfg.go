package cfg

import (
	"github.com/drand/drand/v2/crypto"
	"github.com/drand/drand/v2/internal/chain"
)

// Config stores configuration for the orchestrator.
// It's in a separate package to avoid import cycles.
type Config struct {
	N            int
	Thr          int
	Period       string
	Binary       string
	WithCurl     bool
	Scheme       *crypto.Scheme
	BeaconID     string
	IsCandidate  bool
	DBEngineType chain.StorageType
	PgDSN        func() string
	MemDBSize    int
	Offset       int
	BasePath     string
}
