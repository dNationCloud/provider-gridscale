// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/dNationCloud/provider-gridscale/apis/ssl/v1alpha1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Filesystem.
func (mg *Filesystem) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkUUID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkUUIDRef,
		Selector:     mg.Spec.ForProvider.NetworkUUIDSelector,
		To: reference.To{
			List:    &NetworkList{},
			Managed: &Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkUUID")
	}
	mg.Spec.ForProvider.NetworkUUID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkUUIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NetworkUUID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.NetworkUUIDRef,
		Selector:     mg.Spec.InitProvider.NetworkUUIDSelector,
		To: reference.To{
			List:    &NetworkList{},
			Managed: &Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NetworkUUID")
	}
	mg.Spec.InitProvider.NetworkUUID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NetworkUUIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Loadbalancer.
func (mg *Loadbalancer) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.ForwardingRule); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ForwardingRule[i3].CertificateUUID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ForwardingRule[i3].CertificateUUIDRef,
			Selector:     mg.Spec.ForProvider.ForwardingRule[i3].CertificateUUIDSelector,
			To: reference.To{
				List:    &v1alpha1.CertificateList{},
				Managed: &v1alpha1.Certificate{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ForwardingRule[i3].CertificateUUID")
		}
		mg.Spec.ForProvider.ForwardingRule[i3].CertificateUUID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ForwardingRule[i3].CertificateUUIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ListenIPv4UUID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ListenIPv4UUIDRef,
		Selector:     mg.Spec.ForProvider.ListenIPv4UUIDSelector,
		To: reference.To{
			List:    &IPv4List{},
			Managed: &IPv4{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ListenIPv4UUID")
	}
	mg.Spec.ForProvider.ListenIPv4UUID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ListenIPv4UUIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ListenIPv6UUID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ListenIPv6UUIDRef,
		Selector:     mg.Spec.ForProvider.ListenIPv6UUIDSelector,
		To: reference.To{
			List:    &IPv6List{},
			Managed: &IPv6{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ListenIPv6UUID")
	}
	mg.Spec.ForProvider.ListenIPv6UUID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ListenIPv6UUIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.ForwardingRule); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ForwardingRule[i3].CertificateUUID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ForwardingRule[i3].CertificateUUIDRef,
			Selector:     mg.Spec.InitProvider.ForwardingRule[i3].CertificateUUIDSelector,
			To: reference.To{
				List:    &v1alpha1.CertificateList{},
				Managed: &v1alpha1.Certificate{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.ForwardingRule[i3].CertificateUUID")
		}
		mg.Spec.InitProvider.ForwardingRule[i3].CertificateUUID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.ForwardingRule[i3].CertificateUUIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ListenIPv4UUID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ListenIPv4UUIDRef,
		Selector:     mg.Spec.InitProvider.ListenIPv4UUIDSelector,
		To: reference.To{
			List:    &IPv4List{},
			Managed: &IPv4{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ListenIPv4UUID")
	}
	mg.Spec.InitProvider.ListenIPv4UUID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ListenIPv4UUIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ListenIPv6UUID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ListenIPv6UUIDRef,
		Selector:     mg.Spec.InitProvider.ListenIPv6UUIDSelector,
		To: reference.To{
			List:    &IPv6List{},
			Managed: &IPv6{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ListenIPv6UUID")
	}
	mg.Spec.InitProvider.ListenIPv6UUID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ListenIPv6UUIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Server.
func (mg *Server) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IPv4),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.IPv4Ref,
		Selector:     mg.Spec.ForProvider.IPv4Selector,
		To: reference.To{
			List:    &IPv4List{},
			Managed: &IPv4{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IPv4")
	}
	mg.Spec.ForProvider.IPv4 = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IPv4Ref = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IPv6),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.IPv6Ref,
		Selector:     mg.Spec.ForProvider.IPv6Selector,
		To: reference.To{
			List:    &IPv6List{},
			Managed: &IPv6{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IPv6")
	}
	mg.Spec.ForProvider.IPv6 = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IPv6Ref = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Isoimage),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.IsoimageRef,
		Selector:     mg.Spec.ForProvider.IsoimageSelector,
		To: reference.To{
			List:    &IsoimageList{},
			Managed: &Isoimage{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Isoimage")
	}
	mg.Spec.ForProvider.Isoimage = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IsoimageRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Network); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Network[i3].FirewallTemplateUUID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Network[i3].FirewallTemplateUUIDRef,
			Selector:     mg.Spec.ForProvider.Network[i3].FirewallTemplateUUIDSelector,
			To: reference.To{
				List:    &FirewallList{},
				Managed: &Firewall{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Network[i3].FirewallTemplateUUID")
		}
		mg.Spec.ForProvider.Network[i3].FirewallTemplateUUID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Network[i3].FirewallTemplateUUIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Network); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Network[i3].ObjectUUID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Network[i3].ObjectUUIDRef,
			Selector:     mg.Spec.ForProvider.Network[i3].ObjectUUIDSelector,
			To: reference.To{
				List:    &NetworkList{},
				Managed: &Network{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Network[i3].ObjectUUID")
		}
		mg.Spec.ForProvider.Network[i3].ObjectUUID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Network[i3].ObjectUUIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Storage); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Storage[i3].ObjectUUID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Storage[i3].ObjectUUIDRef,
			Selector:     mg.Spec.ForProvider.Storage[i3].ObjectUUIDSelector,
			To: reference.To{
				List:    &StorageList{},
				Managed: &Storage{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Storage[i3].ObjectUUID")
		}
		mg.Spec.ForProvider.Storage[i3].ObjectUUID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Storage[i3].ObjectUUIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IPv4),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.IPv4Ref,
		Selector:     mg.Spec.InitProvider.IPv4Selector,
		To: reference.To{
			List:    &IPv4List{},
			Managed: &IPv4{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IPv4")
	}
	mg.Spec.InitProvider.IPv4 = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IPv4Ref = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IPv6),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.IPv6Ref,
		Selector:     mg.Spec.InitProvider.IPv6Selector,
		To: reference.To{
			List:    &IPv6List{},
			Managed: &IPv6{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IPv6")
	}
	mg.Spec.InitProvider.IPv6 = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IPv6Ref = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Isoimage),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.IsoimageRef,
		Selector:     mg.Spec.InitProvider.IsoimageSelector,
		To: reference.To{
			List:    &IsoimageList{},
			Managed: &Isoimage{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Isoimage")
	}
	mg.Spec.InitProvider.Isoimage = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IsoimageRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Network); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Network[i3].FirewallTemplateUUID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.Network[i3].FirewallTemplateUUIDRef,
			Selector:     mg.Spec.InitProvider.Network[i3].FirewallTemplateUUIDSelector,
			To: reference.To{
				List:    &FirewallList{},
				Managed: &Firewall{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Network[i3].FirewallTemplateUUID")
		}
		mg.Spec.InitProvider.Network[i3].FirewallTemplateUUID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Network[i3].FirewallTemplateUUIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Network); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Network[i3].ObjectUUID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.Network[i3].ObjectUUIDRef,
			Selector:     mg.Spec.InitProvider.Network[i3].ObjectUUIDSelector,
			To: reference.To{
				List:    &NetworkList{},
				Managed: &Network{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Network[i3].ObjectUUID")
		}
		mg.Spec.InitProvider.Network[i3].ObjectUUID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Network[i3].ObjectUUIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Storage); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Storage[i3].ObjectUUID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.Storage[i3].ObjectUUIDRef,
			Selector:     mg.Spec.InitProvider.Storage[i3].ObjectUUIDSelector,
			To: reference.To{
				List:    &StorageList{},
				Managed: &Storage{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Storage[i3].ObjectUUID")
		}
		mg.Spec.InitProvider.Storage[i3].ObjectUUID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Storage[i3].ObjectUUIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Snapshot.
func (mg *Snapshot) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageUUID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.StorageUUIDRef,
		Selector:     mg.Spec.ForProvider.StorageUUIDSelector,
		To: reference.To{
			List:    &StorageList{},
			Managed: &Storage{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageUUID")
	}
	mg.Spec.ForProvider.StorageUUID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageUUIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageUUID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.StorageUUIDRef,
		Selector:     mg.Spec.InitProvider.StorageUUIDSelector,
		To: reference.To{
			List:    &StorageList{},
			Managed: &Storage{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.StorageUUID")
	}
	mg.Spec.InitProvider.StorageUUID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.StorageUUIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Storage.
func (mg *Storage) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Template); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Template[i3].Sshkeys),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.Template[i3].SshkeysRefs,
			Selector:      mg.Spec.ForProvider.Template[i3].SshkeysSelector,
			To: reference.To{
				List:    &SshkeyList{},
				Managed: &Sshkey{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Template[i3].Sshkeys")
		}
		mg.Spec.ForProvider.Template[i3].Sshkeys = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.Template[i3].SshkeysRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Template); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Template[i3].TemplateUUID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Template[i3].TemplateUUIDRef,
			Selector:     mg.Spec.ForProvider.Template[i3].TemplateUUIDSelector,
			To: reference.To{
				List:    &TemplateList{},
				Managed: &Template{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Template[i3].TemplateUUID")
		}
		mg.Spec.ForProvider.Template[i3].TemplateUUID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Template[i3].TemplateUUIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Template); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Template[i3].Sshkeys),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.Template[i3].SshkeysRefs,
			Selector:      mg.Spec.InitProvider.Template[i3].SshkeysSelector,
			To: reference.To{
				List:    &SshkeyList{},
				Managed: &Sshkey{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Template[i3].Sshkeys")
		}
		mg.Spec.InitProvider.Template[i3].Sshkeys = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.Template[i3].SshkeysRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Template); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Template[i3].TemplateUUID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.Template[i3].TemplateUUIDRef,
			Selector:     mg.Spec.InitProvider.Template[i3].TemplateUUIDSelector,
			To: reference.To{
				List:    &TemplateList{},
				Managed: &Template{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Template[i3].TemplateUUID")
		}
		mg.Spec.InitProvider.Template[i3].TemplateUUID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Template[i3].TemplateUUIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Template.
func (mg *Template) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SnapshotUUID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SnapshotUUIDRef,
		Selector:     mg.Spec.ForProvider.SnapshotUUIDSelector,
		To: reference.To{
			List:    &SnapshotList{},
			Managed: &Snapshot{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SnapshotUUID")
	}
	mg.Spec.ForProvider.SnapshotUUID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SnapshotUUIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SnapshotUUID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.SnapshotUUIDRef,
		Selector:     mg.Spec.InitProvider.SnapshotUUIDSelector,
		To: reference.To{
			List:    &SnapshotList{},
			Managed: &Snapshot{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SnapshotUUID")
	}
	mg.Spec.InitProvider.SnapshotUUID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SnapshotUUIDRef = rsp.ResolvedReference

	return nil
}
