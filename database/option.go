package database

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/canonical/go-dqlite/app"
	"github.com/juju/errors"

	"github.com/juju/juju/agent"
	"github.com/juju/juju/core/network"
)

const (
	dqliteDataDir = "dqlite"
	dqlitePort    = 17666
)

// OptionFactory creates Dqlite `App` initialisation arguments and options.
// These generally depend on a controller's agent config.
type OptionFactory struct {
	cfg            agent.Config
	port           int
	interfaceAddrs func() ([]net.Addr, error)

	bindAddress string
}

// NewOptionFactory returns a new OptionFactory reference
// based on the input agent configuration.
func NewOptionFactory(cfg agent.Config) *OptionFactory {
	return &OptionFactory{
		cfg:            cfg,
		port:           dqlitePort,
		interfaceAddrs: net.InterfaceAddrs,
	}
}

// EnsureDataDir ensures that a directory for Dqlite data exists at
// a path determined by the agent config, then returns that path.
func (f *OptionFactory) EnsureDataDir() (string, error) {
	dir := filepath.Join(f.cfg.DataDir(), dqliteDataDir)
	err := os.MkdirAll(dir, 0700)
	return dir, errors.Annotatef(err, "creating directory for Dqlite data")
}

// WithAddressOption returns a Dqlite application Option
// for specifying the local address:port to use.
// See the comment for ensureBindAddress below.
func (f *OptionFactory) WithAddressOption() (app.Option, error) {
	if err := f.ensureBindAddress(); err != nil {
		return nil, errors.Annotate(err, "ensuring Dqlite bind address")
	}

	return app.WithAddress(fmt.Sprintf("%s:%d", f.bindAddress, f.port)), nil
}

// WithTLSOption returns a Dqlite application Option for TLS encryption
// of traffic between clients and clustered application nodes.
func (f *OptionFactory) WithTLSOption() (app.Option, error) {
	stateInfo, ok := f.cfg.StateServingInfo()
	if !ok {
		return nil, errors.NotSupportedf("Dqlite node initialisation on non-controller machine/container")
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM([]byte(f.cfg.CACert()))

	controllerCert, err := tls.X509KeyPair([]byte(stateInfo.Cert), []byte(stateInfo.PrivateKey))
	if err != nil {
		return nil, errors.Annotate(err, "parsing controller certificate")
	}

	listen := &tls.Config{
		ClientCAs:    caCertPool,
		Certificates: []tls.Certificate{controllerCert},
	}

	dial := &tls.Config{
		RootCAs:      caCertPool,
		Certificates: []tls.Certificate{controllerCert},
		// We cannot provide a ServerName value here, so we rely on the
		// server validating the controller's client certificate.
		InsecureSkipVerify: true,
	}

	return app.WithTLS(listen, dial), nil
}

// ensureBindAddress sets the bind address, used by clients and peers.
// TODO (manadart 2022-09-30): For now, this is *similar* to the peergrouper
// logic in that we require a unique local-cloud IP. We will need to revisit
// this because at present it is not influenced by a configured juju-ha-space.
func (f *OptionFactory) ensureBindAddress() error {
	if f.bindAddress != "" {
		return nil
	}

	sysAddrs, err := f.interfaceAddrs()
	if err != nil {
		return errors.Annotate(err, "querying addresses of system NICs")
	}

	var addrs network.MachineAddresses
	for _, sysAddr := range sysAddrs {
		var host string

		switch v := sysAddr.(type) {
		case *net.IPNet:
			host = v.IP.String()
		case *net.IPAddr:
			host = v.IP.String()
		default:
			continue
		}

		addrs = append(addrs, network.NewMachineAddress(host))
	}

	cloudLocal := addrs.AllMatchingScope(network.ScopeMatchCloudLocal)
	if len(cloudLocal) == 0 {
		return errors.NotFoundf("suitable local address for advertising to Dqlite peers")
	}
	if len(cloudLocal) > 1 {
		return errors.Errorf(
			"multiple local-cloud addresses detected. Dqlite bootstrap requires a unique address; found %v", cloudLocal)
	}

	f.bindAddress = cloudLocal.Values()[0]
	return nil
}
