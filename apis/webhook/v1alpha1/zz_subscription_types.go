/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CustomHeaderInitParameters struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type CustomHeaderObservation struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type CustomHeaderParameters struct {

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DeliveryMethodInitParameters struct {

	// The custom_header of a webhook subscription define any optional headers that will be passed along with the payload to the destination URL.
	CustomHeader []CustomHeaderInitParameters `json:"customHeader,omitempty" tf:"custom_header,omitempty"`

	// Whether this webhook subscription is temporarily disabled. Becomes true if the delivery method URL is repeatedly rejected by the server.
	TemporarilyDisabled *bool `json:"temporarilyDisabled,omitempty" tf:"temporarily_disabled,omitempty"`

	// The type indicating the schema of the object. The provider sets this as webhook_subscription, which is currently the only acceptable value.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The destination URL for webhook delivery.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type DeliveryMethodObservation struct {

	// The custom_header of a webhook subscription define any optional headers that will be passed along with the payload to the destination URL.
	CustomHeader []CustomHeaderObservation `json:"customHeader,omitempty" tf:"custom_header,omitempty"`

	// Whether this webhook subscription is temporarily disabled. Becomes true if the delivery method URL is repeatedly rejected by the server.
	TemporarilyDisabled *bool `json:"temporarilyDisabled,omitempty" tf:"temporarily_disabled,omitempty"`

	// The type indicating the schema of the object. The provider sets this as webhook_subscription, which is currently the only acceptable value.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The destination URL for webhook delivery.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type DeliveryMethodParameters struct {

	// The custom_header of a webhook subscription define any optional headers that will be passed along with the payload to the destination URL.
	// +kubebuilder:validation:Optional
	CustomHeader []CustomHeaderParameters `json:"customHeader,omitempty" tf:"custom_header,omitempty"`

	// Whether this webhook subscription is temporarily disabled. Becomes true if the delivery method URL is repeatedly rejected by the server.
	// +kubebuilder:validation:Optional
	TemporarilyDisabled *bool `json:"temporarilyDisabled,omitempty" tf:"temporarily_disabled,omitempty"`

	// The type indicating the schema of the object. The provider sets this as webhook_subscription, which is currently the only acceptable value.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The destination URL for webhook delivery.
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type FilterInitParameters struct {

	// The id of the object being used as the filter. This field is required for all filter types except account_reference.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The type indicating the schema of the object. The provider sets this as webhook_subscription, which is currently the only acceptable value.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type FilterObservation struct {

	// The id of the object being used as the filter. This field is required for all filter types except account_reference.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The type indicating the schema of the object. The provider sets this as webhook_subscription, which is currently the only acceptable value.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type FilterParameters struct {

	// The id of the object being used as the filter. This field is required for all filter types except account_reference.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The type indicating the schema of the object. The provider sets this as webhook_subscription, which is currently the only acceptable value.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SubscriptionInitParameters struct {

	// Determines whether the subscription will produce webhook events.
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// The object describing where to send the webhooks.
	DeliveryMethod []DeliveryMethodInitParameters `json:"deliveryMethod,omitempty" tf:"delivery_method,omitempty"`

	// A short description of the webhook subscription
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A set of outbound event types the webhook will receive. The follow event types are possible:
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`

	// determines which events will match and produce a webhook. There are currently three types of filters that can be applied to webhook subscriptions: service_reference, team_reference and account_reference.
	Filter []FilterInitParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// The type indicating the schema of the object. The provider sets this as webhook_subscription, which is currently the only acceptable value.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SubscriptionObservation struct {

	// Determines whether the subscription will produce webhook events.
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// The object describing where to send the webhooks.
	DeliveryMethod []DeliveryMethodObservation `json:"deliveryMethod,omitempty" tf:"delivery_method,omitempty"`

	// A short description of the webhook subscription
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A set of outbound event types the webhook will receive. The follow event types are possible:
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`

	// determines which events will match and produce a webhook. There are currently three types of filters that can be applied to webhook subscriptions: service_reference, team_reference and account_reference.
	Filter []FilterObservation `json:"filter,omitempty" tf:"filter,omitempty"`

	// The id of the object being used as the filter. This field is required for all filter types except account_reference.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The type indicating the schema of the object. The provider sets this as webhook_subscription, which is currently the only acceptable value.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SubscriptionParameters struct {

	// Determines whether the subscription will produce webhook events.
	// +kubebuilder:validation:Optional
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// The object describing where to send the webhooks.
	// +kubebuilder:validation:Optional
	DeliveryMethod []DeliveryMethodParameters `json:"deliveryMethod,omitempty" tf:"delivery_method,omitempty"`

	// A short description of the webhook subscription
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A set of outbound event types the webhook will receive. The follow event types are possible:
	// +kubebuilder:validation:Optional
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`

	// determines which events will match and produce a webhook. There are currently three types of filters that can be applied to webhook subscriptions: service_reference, team_reference and account_reference.
	// +kubebuilder:validation:Optional
	Filter []FilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// The type indicating the schema of the object. The provider sets this as webhook_subscription, which is currently the only acceptable value.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// SubscriptionSpec defines the desired state of Subscription
type SubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubscriptionParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SubscriptionInitParameters `json:"initProvider,omitempty"`
}

// SubscriptionStatus defines the observed state of Subscription.
type SubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Subscription is the Schema for the Subscriptions API. Creates and manages a webhook subscription in PagerDuty.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerduty}
type Subscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.deliveryMethod) || has(self.initProvider.deliveryMethod)",message="deliveryMethod is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.events) || has(self.initProvider.events)",message="events is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.filter) || has(self.initProvider.filter)",message="filter is a required parameter"
	Spec   SubscriptionSpec   `json:"spec"`
	Status SubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubscriptionList contains a list of Subscriptions
type SubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subscription `json:"items"`
}

// Repository type metadata.
var (
	Subscription_Kind             = "Subscription"
	Subscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Subscription_Kind}.String()
	Subscription_KindAPIVersion   = Subscription_Kind + "." + CRDGroupVersion.String()
	Subscription_GroupVersionKind = CRDGroupVersion.WithKind(Subscription_Kind)
)

func init() {
	SchemeBuilder.Register(&Subscription{}, &SubscriptionList{})
}
