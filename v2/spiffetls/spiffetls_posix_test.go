//go:build !windows
// +build !windows

package spiffetls_test

import (
	"github.com/damarescavalcante/go-spiffe/v2/spiffetls"
	"github.com/damarescavalcante/go-spiffe/v2/spiffetls/tlsconfig"
)

func listenAndDialCasesOS() []listenAndDialCase {
	return []listenAndDialCase{
		{
			name:             "Wrong workload API server socket",
			dialMode:         spiffetls.TLSClient(tlsconfig.AuthorizeID(serverID)),
			defaultWlAPIAddr: "wrong-socket-path",
			dialErr:          "spiffetls: cannot create X.509 source: workload endpoint socket URI must have a \"tcp\" or \"unix\" scheme",
			listenErr:        "spiffetls: cannot create X.509 source: workload endpoint socket URI must have a \"tcp\" or \"unix\" scheme",
		},
	}
}
