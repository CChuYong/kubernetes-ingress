package v1beta1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:validation:Optional
// +kubebuilder:resource:shortName=pr

// DosProtectedResource defines a collection of Dos protected resources.
// status: preview
type DosProtectedResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec DosProtectedResourceSpec `json:"spec"`
}

type DosProtectedResourceSpec struct {
	Enable           bool            `json:"enable"`
	Name             string          `json:"name"`
	ApDosMonitor     string          `json:"apDosMonitor"`
	DosAccessLogDest string          `json:"dosAccessLogDest"`
	ApDosPolicy      string          `json:"apDosPolicy"`
	DosSecurityLog   *DosSecurityLog `json:"dosSecurityLog"`
}

// DosSecurityLog defines the security log of a Dos policy.
type DosSecurityLog struct {
	Enable       bool   `json:"enable"`
	ApDosLogConf string `json:"apDosLogConf"`
	DosLogDest   string `json:"dosLogDest"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DosProtectedResourceList is a list of the DosProtectedResource resources.
type DosProtectedResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []DosProtectedResource `json:"items"`
}
