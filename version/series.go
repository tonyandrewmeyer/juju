// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package version

import corebase "github.com/juju/juju/core/base"

// DefaultSupportedLTS returns the latest LTS that Juju supports and is
// compatible with.
func DefaultSupportedLTS() string {
	return "noble"
}

// DefaultSupportedLTSBase returns the latest LTS base that Juju supports
// and is compatible with.
func DefaultSupportedLTSBase() corebase.Base {
	return corebase.MakeDefaultBase(corebase.UbuntuOS, "24.04")
}
