// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type K8SInitParameters struct {

	// The gridscale's Kubernetes version of this instance (e.g. "1.21.14-gs1"). Define which gridscale k8s version will be used to create the cluster. For convenience, please use gscloud to get the list of available gridscale k8s version. NOTE: Either gsk_version or release is set at a time.
	// The gridscale k8s PaaS version (issued by gridscale) of this instance.
	GskVersion *string `json:"gskVersion,omitempty" tf:"gsk_version,omitempty"`

	// List of labels in the format [ "label1", "label2" ].
	// List of labels.
	// +listType=set
	Labels []*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.
	// The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Node pool's specification. NOTE: The node pool's specification is not yet mutable (except node_count).
	// Node pool's specification.
	NodePool []NodePoolInitParameters `json:"nodePool,omitempty" tf:"node_pool,omitempty"`

	// Custom CA from customer in pem format as string.
	// Custom CA from customer in pem format as string.
	OidcCAPem *string `json:"oidcCaPem,omitempty" tf:"oidc_ca_pem,omitempty"`

	// A client ID that all tokens must be issued for.
	// A client ID that all tokens must be issued for.
	OidcClientID *string `json:"oidcClientId,omitempty" tf:"oidc_client_id,omitempty"`

	// Enable OIDC for the k8s cluster.
	// Disable or enable OIDC
	OidcEnabled *bool `json:"oidcEnabled,omitempty" tf:"oidc_enabled,omitempty"`

	// JWT claim to use as the user's group.
	// JWT claim to use as the user's group.
	OidcGroupsClaim *string `json:"oidcGroupsClaim,omitempty" tf:"oidc_groups_claim,omitempty"`

	// Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.
	// Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.
	OidcGroupsPrefix *string `json:"oidcGroupsPrefix,omitempty" tf:"oidc_groups_prefix,omitempty"`

	// URL of the provider that allows the API server to discover public signing keys. Only URLs that use the https:// scheme are accepted.
	// URL of the provider that allows the API server to discover public signing keys. Only URLs that use the https:// scheme are accepted.
	OidcIssuerURL *string `json:"oidcIssuerUrl,omitempty" tf:"oidc_issuer_url,omitempty"`

	// A key=value pair that describes a required claim in the ID Token. Multiple claims can be set like this: key1=value1,key2=value2.
	// A key=value pair that describes a required claim in the ID Token. Multiple claims can be set like this: key1=value1,key2=value2
	OidcRequiredClaim *string `json:"oidcRequiredClaim,omitempty" tf:"oidc_required_claim,omitempty"`

	// The signing algorithms accepted. Default is 'RS256'. Other option is 'RS512'.
	// The signing algorithms accepted. Default is 'RS256'. Other option is 'RS512'.
	OidcSigningAlgs *string `json:"oidcSigningAlgs,omitempty" tf:"oidc_signing_algs,omitempty"`

	// JWT claim to use as the user name.
	// JWT claim to use as the user name.
	OidcUsernameClaim *string `json:"oidcUsernameClaim,omitempty" tf:"oidc_username_claim,omitempty"`

	// Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this flag isn't provided and --oidc-username-claim is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of --oidc-issuer-url. The value - can be used to disable all prefixing.
	// Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this flag isn't provided and --oidc-username-claim is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of --oidc-issuer-url. The value - can be used to disable all prefixing.
	OidcUsernamePrefix *string `json:"oidcUsernamePrefix,omitempty" tf:"oidc_username_prefix,omitempty"`

	// The Kubernetes release of this instance. Define which release will be used to create the cluster. For convenience, please use gscloud to get the list of available releases. NOTE: Either gsk_version or release is set at a time.
	// The k8s release of this instance.
	Release *string `json:"release,omitempty" tf:"release,omitempty"`

	// DEPRECATED  Security zone UUID linked to the Kubernetes resource. If security_zone_uuid is not set, the default security zone will be created (if it doesn't exist) and linked. A change of this argument necessitates the re-creation of the resource.
	// Security zone UUID linked to PaaS service.
	SecurityZoneUUID *string `json:"securityZoneUuid,omitempty" tf:"security_zone_uuid,omitempty"`
}

type K8SObservation struct {

	// Defines the date and time of the last object change.
	// Time of the last change
	ChangeTime *string `json:"changeTime,omitempty" tf:"change_time,omitempty"`

	// The time the object was created.
	// Time this service was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// The gridscale's Kubernetes version of this instance (e.g. "1.21.14-gs1"). Define which gridscale k8s version will be used to create the cluster. For convenience, please use gscloud to get the list of available gridscale k8s version. NOTE: Either gsk_version or release is set at a time.
	// The gridscale k8s PaaS version (issued by gridscale) of this instance.
	GskVersion *string `json:"gskVersion,omitempty" tf:"gsk_version,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Private network UUID which k8s nodes are attached to. It can be used to attach other PaaS/VMs.
	// Private network UUID which k8s nodes are attached to. It can be used to attach other PaaS/VMs.
	K8SPrivateNetworkUUID *string `json:"k8sPrivateNetworkUuid,omitempty" tf:"k8s_private_network_uuid,omitempty"`

	// List of labels in the format [ "label1", "label2" ].
	// List of labels.
	// +listType=set
	Labels []*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The port number where this k8s service accepts connections.
	ListenPort []ListenPortObservation `json:"listenPort,omitempty" tf:"listen_port,omitempty"`

	// The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.
	// The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// DEPRECATED  Network UUID containing security zone, which is linked to the k8s cluster.
	// Network UUID containing security zone
	NetworkUUID *string `json:"networkUuid,omitempty" tf:"network_uuid,omitempty"`

	// Node pool's specification. NOTE: The node pool's specification is not yet mutable (except node_count).
	// Node pool's specification.
	NodePool []NodePoolObservation `json:"nodePool,omitempty" tf:"node_pool,omitempty"`

	// Custom CA from customer in pem format as string.
	// Custom CA from customer in pem format as string.
	OidcCAPem *string `json:"oidcCaPem,omitempty" tf:"oidc_ca_pem,omitempty"`

	// A client ID that all tokens must be issued for.
	// A client ID that all tokens must be issued for.
	OidcClientID *string `json:"oidcClientId,omitempty" tf:"oidc_client_id,omitempty"`

	// Enable OIDC for the k8s cluster.
	// Disable or enable OIDC
	OidcEnabled *bool `json:"oidcEnabled,omitempty" tf:"oidc_enabled,omitempty"`

	// JWT claim to use as the user's group.
	// JWT claim to use as the user's group.
	OidcGroupsClaim *string `json:"oidcGroupsClaim,omitempty" tf:"oidc_groups_claim,omitempty"`

	// Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.
	// Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.
	OidcGroupsPrefix *string `json:"oidcGroupsPrefix,omitempty" tf:"oidc_groups_prefix,omitempty"`

	// URL of the provider that allows the API server to discover public signing keys. Only URLs that use the https:// scheme are accepted.
	// URL of the provider that allows the API server to discover public signing keys. Only URLs that use the https:// scheme are accepted.
	OidcIssuerURL *string `json:"oidcIssuerUrl,omitempty" tf:"oidc_issuer_url,omitempty"`

	// A key=value pair that describes a required claim in the ID Token. Multiple claims can be set like this: key1=value1,key2=value2.
	// A key=value pair that describes a required claim in the ID Token. Multiple claims can be set like this: key1=value1,key2=value2
	OidcRequiredClaim *string `json:"oidcRequiredClaim,omitempty" tf:"oidc_required_claim,omitempty"`

	// The signing algorithms accepted. Default is 'RS256'. Other option is 'RS512'.
	// The signing algorithms accepted. Default is 'RS256'. Other option is 'RS512'.
	OidcSigningAlgs *string `json:"oidcSigningAlgs,omitempty" tf:"oidc_signing_algs,omitempty"`

	// JWT claim to use as the user name.
	// JWT claim to use as the user name.
	OidcUsernameClaim *string `json:"oidcUsernameClaim,omitempty" tf:"oidc_username_claim,omitempty"`

	// Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this flag isn't provided and --oidc-username-claim is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of --oidc-issuer-url. The value - can be used to disable all prefixing.
	// Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this flag isn't provided and --oidc-username-claim is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of --oidc-issuer-url. The value - can be used to disable all prefixing.
	OidcUsernamePrefix *string `json:"oidcUsernamePrefix,omitempty" tf:"oidc_username_prefix,omitempty"`

	// The Kubernetes release of this instance. Define which release will be used to create the cluster. For convenience, please use gscloud to get the list of available releases. NOTE: Either gsk_version or release is set at a time.
	// The k8s release of this instance.
	Release *string `json:"release,omitempty" tf:"release,omitempty"`

	// DEPRECATED  Security zone UUID linked to the Kubernetes resource. If security_zone_uuid is not set, the default security zone will be created (if it doesn't exist) and linked. A change of this argument necessitates the re-creation of the resource.
	// Security zone UUID linked to PaaS service.
	SecurityZoneUUID *string `json:"securityZoneUuid,omitempty" tf:"security_zone_uuid,omitempty"`

	// PaaS service template that k8s service uses.g. the k8s service is upgraded by gridscale staffs).
	// PaaS service template identifier for this service.
	ServiceTemplateUUID *string `json:"serviceTemplateUuid,omitempty" tf:"service_template_uuid,omitempty"`

	// status indicates the status of the object.
	// Current status of the service
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The amount of minutes the IP address has been in use.
	// Number of minutes that PaaS service is in use
	UsageInMinutes *float64 `json:"usageInMinutes,omitempty" tf:"usage_in_minutes,omitempty"`
}

type K8SParameters struct {

	// The gridscale's Kubernetes version of this instance (e.g. "1.21.14-gs1"). Define which gridscale k8s version will be used to create the cluster. For convenience, please use gscloud to get the list of available gridscale k8s version. NOTE: Either gsk_version or release is set at a time.
	// The gridscale k8s PaaS version (issued by gridscale) of this instance.
	// +kubebuilder:validation:Optional
	GskVersion *string `json:"gskVersion,omitempty" tf:"gsk_version,omitempty"`

	// List of labels in the format [ "label1", "label2" ].
	// List of labels.
	// +kubebuilder:validation:Optional
	// +listType=set
	Labels []*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.
	// The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Node pool's specification. NOTE: The node pool's specification is not yet mutable (except node_count).
	// Node pool's specification.
	// +kubebuilder:validation:Optional
	NodePool []NodePoolParameters `json:"nodePool,omitempty" tf:"node_pool,omitempty"`

	// Custom CA from customer in pem format as string.
	// Custom CA from customer in pem format as string.
	// +kubebuilder:validation:Optional
	OidcCAPem *string `json:"oidcCaPem,omitempty" tf:"oidc_ca_pem,omitempty"`

	// A client ID that all tokens must be issued for.
	// A client ID that all tokens must be issued for.
	// +kubebuilder:validation:Optional
	OidcClientID *string `json:"oidcClientId,omitempty" tf:"oidc_client_id,omitempty"`

	// Enable OIDC for the k8s cluster.
	// Disable or enable OIDC
	// +kubebuilder:validation:Optional
	OidcEnabled *bool `json:"oidcEnabled,omitempty" tf:"oidc_enabled,omitempty"`

	// JWT claim to use as the user's group.
	// JWT claim to use as the user's group.
	// +kubebuilder:validation:Optional
	OidcGroupsClaim *string `json:"oidcGroupsClaim,omitempty" tf:"oidc_groups_claim,omitempty"`

	// Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.
	// Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.
	// +kubebuilder:validation:Optional
	OidcGroupsPrefix *string `json:"oidcGroupsPrefix,omitempty" tf:"oidc_groups_prefix,omitempty"`

	// URL of the provider that allows the API server to discover public signing keys. Only URLs that use the https:// scheme are accepted.
	// URL of the provider that allows the API server to discover public signing keys. Only URLs that use the https:// scheme are accepted.
	// +kubebuilder:validation:Optional
	OidcIssuerURL *string `json:"oidcIssuerUrl,omitempty" tf:"oidc_issuer_url,omitempty"`

	// A key=value pair that describes a required claim in the ID Token. Multiple claims can be set like this: key1=value1,key2=value2.
	// A key=value pair that describes a required claim in the ID Token. Multiple claims can be set like this: key1=value1,key2=value2
	// +kubebuilder:validation:Optional
	OidcRequiredClaim *string `json:"oidcRequiredClaim,omitempty" tf:"oidc_required_claim,omitempty"`

	// The signing algorithms accepted. Default is 'RS256'. Other option is 'RS512'.
	// The signing algorithms accepted. Default is 'RS256'. Other option is 'RS512'.
	// +kubebuilder:validation:Optional
	OidcSigningAlgs *string `json:"oidcSigningAlgs,omitempty" tf:"oidc_signing_algs,omitempty"`

	// JWT claim to use as the user name.
	// JWT claim to use as the user name.
	// +kubebuilder:validation:Optional
	OidcUsernameClaim *string `json:"oidcUsernameClaim,omitempty" tf:"oidc_username_claim,omitempty"`

	// Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this flag isn't provided and --oidc-username-claim is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of --oidc-issuer-url. The value - can be used to disable all prefixing.
	// Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this flag isn't provided and --oidc-username-claim is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of --oidc-issuer-url. The value - can be used to disable all prefixing.
	// +kubebuilder:validation:Optional
	OidcUsernamePrefix *string `json:"oidcUsernamePrefix,omitempty" tf:"oidc_username_prefix,omitempty"`

	// The Kubernetes release of this instance. Define which release will be used to create the cluster. For convenience, please use gscloud to get the list of available releases. NOTE: Either gsk_version or release is set at a time.
	// The k8s release of this instance.
	// +kubebuilder:validation:Optional
	Release *string `json:"release,omitempty" tf:"release,omitempty"`

	// DEPRECATED  Security zone UUID linked to the Kubernetes resource. If security_zone_uuid is not set, the default security zone will be created (if it doesn't exist) and linked. A change of this argument necessitates the re-creation of the resource.
	// Security zone UUID linked to PaaS service.
	// +kubebuilder:validation:Optional
	SecurityZoneUUID *string `json:"securityZoneUuid,omitempty" tf:"security_zone_uuid,omitempty"`
}

type ListenPortInitParameters struct {
}

type ListenPortObservation struct {

	// The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}

type ListenPortParameters struct {
}

type NodePoolInitParameters struct {

	// (Immutable) The cluster CIDR that will be used to generate the CIDR of nodes, services, and pods. The allowed CIDR prefix length is /16. If the cluster CIDR is not set, the cluster will use "10.244.0.0/16" as it default (even though the cluster_cidr in the k8s resource is empty).
	// The cluster CIDR that will be used to generate the CIDR of nodes, services, and pods. The allowed CIDR prefix length is /16. If this field is empty, the default value is "10.244.0.0/16"
	ClusterCidr *string `json:"clusterCidr,omitempty" tf:"cluster_cidr,omitempty"`

	// Enables cluster encryption via wireguard if true. Only available for GSK version 1.29 and above. Default is false.
	// Enables cluster encryption via wireguard if true. Only available for GSK version 1.29 and above. Default is false.
	ClusterTrafficEncryption *bool `json:"clusterTrafficEncryption,omitempty" tf:"cluster_traffic_encryption,omitempty"`

	// Cores per worker node.
	// Cores per worker node.
	Cores *float64 `json:"cores,omitempty" tf:"cores,omitempty"`

	// Memory per worker node (in GiB).
	// Memory per worker node (in GiB).
	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`

	// The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.
	// Name of node pool.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Number of worker nodes.
	// Number of worker nodes.
	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	// Rocket storage per worker node (in GiB).
	// Rocket storage per worker node (in GiB).
	RocketStorage *float64 `json:"rocketStorage,omitempty" tf:"rocket_storage,omitempty"`

	// Storage per worker node (in GiB).
	// Storage per worker node (in GiB).
	Storage *float64 `json:"storage,omitempty" tf:"storage,omitempty"`

	// Storage type (one of storage, storage_high, storage_insane).
	// Storage type.
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// Enable surge node to avoid resources shortage during the cluster upgrade (Default: true).
	// Enable surge node to avoid resources shortage during the cluster upgrade.
	SurgeNode *bool `json:"surgeNode,omitempty" tf:"surge_node,omitempty"`
}

type NodePoolObservation struct {

	// (Immutable) The cluster CIDR that will be used to generate the CIDR of nodes, services, and pods. The allowed CIDR prefix length is /16. If the cluster CIDR is not set, the cluster will use "10.244.0.0/16" as it default (even though the cluster_cidr in the k8s resource is empty).
	// The cluster CIDR that will be used to generate the CIDR of nodes, services, and pods. The allowed CIDR prefix length is /16. If this field is empty, the default value is "10.244.0.0/16"
	ClusterCidr *string `json:"clusterCidr,omitempty" tf:"cluster_cidr,omitempty"`

	// Enables cluster encryption via wireguard if true. Only available for GSK version 1.29 and above. Default is false.
	// Enables cluster encryption via wireguard if true. Only available for GSK version 1.29 and above. Default is false.
	ClusterTrafficEncryption *bool `json:"clusterTrafficEncryption,omitempty" tf:"cluster_traffic_encryption,omitempty"`

	// Cores per worker node.
	// Cores per worker node.
	Cores *float64 `json:"cores,omitempty" tf:"cores,omitempty"`

	// Memory per worker node (in GiB).
	// Memory per worker node (in GiB).
	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`

	// The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.
	// Name of node pool.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Number of worker nodes.
	// Number of worker nodes.
	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	// Rocket storage per worker node (in GiB).
	// Rocket storage per worker node (in GiB).
	RocketStorage *float64 `json:"rocketStorage,omitempty" tf:"rocket_storage,omitempty"`

	// Storage per worker node (in GiB).
	// Storage per worker node (in GiB).
	Storage *float64 `json:"storage,omitempty" tf:"storage,omitempty"`

	// Storage type (one of storage, storage_high, storage_insane).
	// Storage type.
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// Enable surge node to avoid resources shortage during the cluster upgrade (Default: true).
	// Enable surge node to avoid resources shortage during the cluster upgrade.
	SurgeNode *bool `json:"surgeNode,omitempty" tf:"surge_node,omitempty"`
}

type NodePoolParameters struct {

	// (Immutable) The cluster CIDR that will be used to generate the CIDR of nodes, services, and pods. The allowed CIDR prefix length is /16. If the cluster CIDR is not set, the cluster will use "10.244.0.0/16" as it default (even though the cluster_cidr in the k8s resource is empty).
	// The cluster CIDR that will be used to generate the CIDR of nodes, services, and pods. The allowed CIDR prefix length is /16. If this field is empty, the default value is "10.244.0.0/16"
	// +kubebuilder:validation:Optional
	ClusterCidr *string `json:"clusterCidr,omitempty" tf:"cluster_cidr,omitempty"`

	// Enables cluster encryption via wireguard if true. Only available for GSK version 1.29 and above. Default is false.
	// Enables cluster encryption via wireguard if true. Only available for GSK version 1.29 and above. Default is false.
	// +kubebuilder:validation:Optional
	ClusterTrafficEncryption *bool `json:"clusterTrafficEncryption,omitempty" tf:"cluster_traffic_encryption,omitempty"`

	// Cores per worker node.
	// Cores per worker node.
	// +kubebuilder:validation:Optional
	Cores *float64 `json:"cores" tf:"cores,omitempty"`

	// Memory per worker node (in GiB).
	// Memory per worker node (in GiB).
	// +kubebuilder:validation:Optional
	Memory *float64 `json:"memory" tf:"memory,omitempty"`

	// The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.
	// Name of node pool.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Number of worker nodes.
	// Number of worker nodes.
	// +kubebuilder:validation:Optional
	NodeCount *float64 `json:"nodeCount" tf:"node_count,omitempty"`

	// Rocket storage per worker node (in GiB).
	// Rocket storage per worker node (in GiB).
	// +kubebuilder:validation:Optional
	RocketStorage *float64 `json:"rocketStorage,omitempty" tf:"rocket_storage,omitempty"`

	// Storage per worker node (in GiB).
	// Storage per worker node (in GiB).
	// +kubebuilder:validation:Optional
	Storage *float64 `json:"storage" tf:"storage,omitempty"`

	// Storage type (one of storage, storage_high, storage_insane).
	// Storage type.
	// +kubebuilder:validation:Optional
	StorageType *string `json:"storageType" tf:"storage_type,omitempty"`

	// Enable surge node to avoid resources shortage during the cluster upgrade (Default: true).
	// Enable surge node to avoid resources shortage during the cluster upgrade.
	// +kubebuilder:validation:Optional
	SurgeNode *bool `json:"surgeNode,omitempty" tf:"surge_node,omitempty"`
}

// K8SSpec defines the desired state of K8S
type K8SSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     K8SParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider K8SInitParameters `json:"initProvider,omitempty"`
}

// K8SStatus defines the observed state of K8S.
type K8SStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        K8SObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// K8S is the Schema for the K8Ss API. Manages a k8s cluster in gridscale.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gridscale}
type K8S struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.nodePool) || (has(self.initProvider) && has(self.initProvider.nodePool))",message="spec.forProvider.nodePool is a required parameter"
	Spec   K8SSpec   `json:"spec"`
	Status K8SStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K8SList contains a list of K8Ss
type K8SList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8S `json:"items"`
}

// Repository type metadata.
var (
	K8S_Kind             = "K8S"
	K8S_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: K8S_Kind}.String()
	K8S_KindAPIVersion   = K8S_Kind + "." + CRDGroupVersion.String()
	K8S_GroupVersionKind = CRDGroupVersion.WithKind(K8S_Kind)
)

func init() {
	SchemeBuilder.Register(&K8S{}, &K8SList{})
}