// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	ipv4 "github.com/dNationCloud/provider-gridscale/internal/controller/gridscale/ipv4"
	ipv6 "github.com/dNationCloud/provider-gridscale/internal/controller/gridscale/ipv6"
	loadbalancer "github.com/dNationCloud/provider-gridscale/internal/controller/gridscale/loadbalancer"
	providerconfig "github.com/dNationCloud/provider-gridscale/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		ipv4.Setup,
		ipv6.Setup,
		loadbalancer.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
