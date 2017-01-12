package oracle

import "github.com/juju/juju/environs"

func init() {
	// this will init the provider map of the internal globalProviderRegistry
	// when the oracle provider package will be used
	// the location when the package is called first is in ../all/all.go
	environs.RegisterProvider(providerType, environProvider{}, providerAliases...)
}
