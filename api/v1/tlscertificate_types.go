/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TLSCertificateSpec defines the desired state of TLSCertificate
type TLSCertificateSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// The following markers will use OpenAPI v3 schema to validate the value
	// More info: https://book.kubebuilder.io/reference/markers/crd-validation.html

	// CommonName is the CN (Common Name) of the certificate
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=64
	CommonName string `json:"commonName"`

	// DNSNames is a list of DNS names to be included in the certificate's Subject Alternative Names
	// +optional
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=100
	DNSNames []string `json:"dnsNames,omitempty"`

	// SecretName is the name of the Secret resource where the certificate will be stored
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:Pattern=`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`
	SecretName string `json:"secretName"`

	// ValidityDays specifies the number of days the certificate should be valid for
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=825
	// +kubebuilder:default=90
	ValidityDays int32 `json:"validityDays"`

	// RenewBefore specifies how long before expiry the certificate should be renewed
	// +optional
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:default=30
	RenewBefore *int32 `json:"renewBefore,omitempty"`
}

// TLSCertificateStatus defines the observed state of TLSCertificate.
type TLSCertificateStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Conditions represent the latest available observations of the Certificate's state
	// +optional
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// NotBefore is the earliest time the certificate is valid from
	// +optional
	NotBefore *metav1.Time `json:"notBefore,omitempty"`

	// NotAfter is the time at which the certificate expires
	// +optional
	NotAfter *metav1.Time `json:"notAfter,omitempty"`

	// RenewalTime is the time at which the certificate should be renewed
	// +optional
	RenewalTime *metav1.Time `json:"renewalTime,omitempty"`

	// SerialNumber is the serial number of the issued certificate
	// +optional
	SerialNumber string `json:"serialNumber,omitempty"`

	// Revision tracks the number of times the certificate has been issued
	// +optional
	Revision int `json:"revision,omitempty"`

	// LastIssuedTime is the timestamp when the certificate was last issued
	// +optional
	LastIssuedTime *metav1.Time `json:"lastIssuedTime,omitempty"`
}

// Condition types for Certificate
const (
	// CertificateConditionReady indicates that the certificate is ready for use
	CertificateConditionReady string = "Ready"

	// CertificateConditionIssuing indicates that certificate issuance is in progress
	CertificateConditionIssuing string = "Issuing"

	// CertificateConditionInvalidConfiguration indicates that the certificate has invalid configuration
	CertificateConditionInvalidConfiguration string = "InvalidConfiguration"
)

// Condition reasons for Certificate
const (
	// ReasonIssued indicates certificate was successfully issued
	ReasonIssued string = "Issued"

	// ReasonRenewing indicates certificate is being renewed
	ReasonRenewing string = "Renewing"

	// ReasonFailed indicates certificate issuance failed
	ReasonFailed string = "Failed"

	// ReasonPending indicates certificate issuance is pending
	ReasonPending string = "Pending"

	// ReasonInvalidSpec indicates the spec is invalid
	ReasonInvalidSpec string = "InvalidSpec"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,shortName=cert;certs
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Secret",type="string",JSONPath=".spec.secretName"
// +kubebuilder:printcolumn:name="NotAfter",type="date",JSONPath=".status.notAfter"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// TLSCertificate is the Schema for the tlscertificates API
type TLSCertificate struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty,omitzero"`

	// spec defines the desired state of TLSCertificate
	// +required
	Spec TLSCertificateSpec `json:"spec"`

	// status defines the observed state of TLSCertificate
	// +optional
	Status TLSCertificateStatus `json:"status,omitempty,omitzero"`
}

// +kubebuilder:object:root=true

// TLSCertificateList contains a list of TLSCertificate
type TLSCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TLSCertificate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TLSCertificate{}, &TLSCertificateList{})
}
