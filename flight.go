package smart_plane

import (
	l "github.com/rafa-mori/logz"
	t "github.com/rafa-mori/smart_plane/types"

	fgi "github.com/rafa-mori/gobe/factory/gateway"
	sp "github.com/rafa-mori/smart_plane/internal/smart_contracts"
)

type Flight struct {
	Logger                l.Logger `json:"-"`
	*t.Mutexes            `json:"-"`
	*t.Reference          `json:"reference,omitempty"`
	*sp.BlockchainManager `json:"-"`
	*fgi.AuthManager      `json:"-"`
}
