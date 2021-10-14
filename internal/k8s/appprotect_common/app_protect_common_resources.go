package appprotect_common

import (
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// GetNsName gets the key of a resource in the format: "resNamespace/resName"
func GetNsName(obj *unstructured.Unstructured) string {
	return obj.GetNamespace() + "/" + obj.GetName()
}

// GetNamespacedName returns the fully namespace qualified name, using default if the namespace is not found.
func GetNamespacedName(namespaceName string) string {
	if !strings.Contains(namespaceName, "/") {
		return "default/" + namespaceName
	}
	return namespaceName
}

// ParseResourceReferenceAnnotation returns a namespace/name string
func ParseResourceReferenceAnnotation(ns, antn string) string {
	if !strings.Contains(antn, "/") {
		return ns + "/" + antn
	}
	return antn
}

// ParseResourceReferenceAnnotationList returns a slice of ns/names strings
func ParseResourceReferenceAnnotationList(ns, annotations string) []string {
	var out []string
	for _, antn := range strings.Split(annotations, ",") {
		out = append(out, ParseResourceReferenceAnnotation(ns, antn))
	}
	return out
}
