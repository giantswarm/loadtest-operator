package v1alpha1

 import (
	"fmt"

 	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

 const (
	kindLoadTest = "Loadtest"
)

 // NewLoadtestCRD returns a new custom resource definition for LoadTest. This
// might look something like the following.
//
//     apiVersion: apiextensions.k8s.io/v1beta1
//     kind: CustomResourceDefinition
//     metadata:
//       name: loadtest.core.giantswarm.io
//     spec:
//       group: core.giantswarm.io
//       scope: Namespaced
//       version: v1alpha1
//       names:
//         kind: Loadtest
//         plural: loadtests
//         singular: loadtest
//       subresources:
//         status: {}
//
func NewLoadtestCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextensionsv1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("loadtests.%s", group),
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   group,
			Scope:   "Namespaced",
			Version: version,
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Kind:     kindLoadtest,
				Plural:   "loadtests",
				Singular: "loadtest",
			},
			Subresources: &apiextensionsv1beta1.CustomResourceSubresources{
				Status: &apiextensionsv1beta1.CustomResourceSubresourceStatus{},
			},
		},
	}
}

 func NewLoadtestTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: version,
		Kind:       kindLoadtest,
	}
}

 // +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

 type Loadtest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              LoadtestSpec   `json:"spec"`
	Status            LoadtestStatus `json:"status"`
}

 // LoadtestSpec is the interface which defines the input parameters for
// a newly rendered g8s loadtest template.
type LoadtestSpec struct {
	Service  string `json:"service" yaml:"service"`
	Requests int `json:"requests" yaml:"requests"`
	Path string `json:"path" yaml:"path"`
}

 // LoadtestStatus holds the rendering result.
type LoadtestStatus struct {
	ConfigMap LoadtestConfigMap `json:"template" yaml:"template"`
}

 type LoadtestConfigMap struct {
	// Name is the name of the test.
	Name string `json:"name" yaml:"name"`
	// Namespace is the namespace of the test
	Namespace string `json:"namespace" yaml:"namespace"`
	// ResourceVersion is the Kubernetes resource version of the configmap.
	// Used to detect if the configmap has changed, e.g. 12345.
	ResourceVersion string `json:"resourceVersion" yaml:"resourceVersion"`
}

 // +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

 type LoadtestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Loadtest `json:"items"`
}