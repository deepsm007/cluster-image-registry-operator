//go:build tools
// +build tools

// Official workaround to track tool dependencies with go modules:
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

package tools

import (
	// Makefile
	_ "github.com/openshift/api/imageregistry/v1/zz_generated.crd-manifests"
	_ "github.com/openshift/build-machinery-go"
)
