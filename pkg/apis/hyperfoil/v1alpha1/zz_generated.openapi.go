// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/Hyperfoil/hyperfoil-operator/pkg/apis/hyperfoil/v1alpha1.Hyperfoil":       schema_pkg_apis_hyperfoil_v1alpha1_Hyperfoil(ref),
		"github.com/Hyperfoil/hyperfoil-operator/pkg/apis/hyperfoil/v1alpha1.HyperfoilSpec":   schema_pkg_apis_hyperfoil_v1alpha1_HyperfoilSpec(ref),
		"github.com/Hyperfoil/hyperfoil-operator/pkg/apis/hyperfoil/v1alpha1.HyperfoilStatus": schema_pkg_apis_hyperfoil_v1alpha1_HyperfoilStatus(ref),
	}
}

func schema_pkg_apis_hyperfoil_v1alpha1_Hyperfoil(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Hyperfoil is the Schema for the hyperfoils API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/Hyperfoil/hyperfoil-operator/pkg/apis/hyperfoil/v1alpha1.HyperfoilSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/Hyperfoil/hyperfoil-operator/pkg/apis/hyperfoil/v1alpha1.HyperfoilStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/Hyperfoil/hyperfoil-operator/pkg/apis/hyperfoil/v1alpha1.HyperfoilSpec", "github.com/Hyperfoil/hyperfoil-operator/pkg/apis/hyperfoil/v1alpha1.HyperfoilStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_hyperfoil_v1alpha1_HyperfoilSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "HyperfoilSpec defines the desired state of Hyperfoil",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_hyperfoil_v1alpha1_HyperfoilStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "HyperfoilStatus defines the observed state of Hyperfoil",
				Type:        []string{"object"},
			},
		},
	}
}
