// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	firewall "github.com/dNationCloud/provider-gridscale/internal/controller/gridscale/firewall"
	ipv4 "github.com/dNationCloud/provider-gridscale/internal/controller/gridscale/ipv4"
	ipv6 "github.com/dNationCloud/provider-gridscale/internal/controller/gridscale/ipv6"
	isoimage "github.com/dNationCloud/provider-gridscale/internal/controller/gridscale/isoimage"
	k8s "github.com/dNationCloud/provider-gridscale/internal/controller/gridscale/k8s"
	loadbalancer "github.com/dNationCloud/provider-gridscale/internal/controller/gridscale/loadbalancer"
	network "github.com/dNationCloud/provider-gridscale/internal/controller/gridscale/network"
	server "github.com/dNationCloud/provider-gridscale/internal/controller/gridscale/server"
	snapshot "github.com/dNationCloud/provider-gridscale/internal/controller/gridscale/snapshot"
	sshkey "github.com/dNationCloud/provider-gridscale/internal/controller/gridscale/sshkey"
	storage "github.com/dNationCloud/provider-gridscale/internal/controller/gridscale/storage"
	template "github.com/dNationCloud/provider-gridscale/internal/controller/gridscale/template"
	providerconfig "github.com/dNationCloud/provider-gridscale/internal/controller/providerconfig"
	certificate "github.com/dNationCloud/provider-gridscale/internal/controller/ssl/certificate"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		firewall.Setup,
		ipv4.Setup,
		ipv6.Setup,
		isoimage.Setup,
		k8s.Setup,
		loadbalancer.Setup,
		network.Setup,
		server.Setup,
		snapshot.Setup,
		sshkey.Setup,
		storage.Setup,
		template.Setup,
		providerconfig.Setup,
		certificate.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
