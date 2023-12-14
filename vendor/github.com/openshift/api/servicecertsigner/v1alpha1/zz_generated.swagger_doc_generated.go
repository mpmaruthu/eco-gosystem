package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_ServiceCertSignerOperatorConfig = map[string]string{
	"":         "ServiceCertSignerOperatorConfig provides information to configure an operator to manage the service cert signing controllers\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
	"metadata": "metadata is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
}

func (ServiceCertSignerOperatorConfig) SwaggerDoc() map[string]string {
	return map_ServiceCertSignerOperatorConfig
}

var map_ServiceCertSignerOperatorConfigList = map[string]string{
	"":         "ServiceCertSignerOperatorConfigList is a collection of items\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
	"metadata": "metadata is the standard list's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "Items contains the items",
}

func (ServiceCertSignerOperatorConfigList) SwaggerDoc() map[string]string {
	return map_ServiceCertSignerOperatorConfigList
}

// AUTO-GENERATED FUNCTIONS END HERE
