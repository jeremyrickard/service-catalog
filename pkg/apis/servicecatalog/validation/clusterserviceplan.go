/*
Copyright 2017 The Kubernetes Authors.

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

package validation

import (
	apivalidation "k8s.io/apimachinery/pkg/api/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"

	sc "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
)

// ValidateClusterServicePlan validates a ClusterServicePlan and returns a list of errors.
func ValidateClusterServicePlan(serviceplan *sc.ClusterServicePlan) field.ErrorList {
	return validateClusterServicePlan(serviceplan)
}

func validateClusterServicePlan(serviceplan *sc.ClusterServicePlan) field.ErrorList {
	allErrs := field.ErrorList{}

	allErrs = append(allErrs,
		apivalidation.ValidateObjectMeta(
			&serviceplan.ObjectMeta,
			false, /* namespace required */
			validateCommonServicePlanName,
			field.NewPath("metadata"))...)

	allErrs = append(allErrs, validateClusterServicePlanSpec(&serviceplan.Spec, field.NewPath("spec"))...)
	return allErrs
}

func validateClusterServicePlanSpec(spec *sc.ClusterServicePlanSpec, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	for _, msg := range validateCommonServiceClassName(spec.ClusterServiceClassRef.Name, false /* prefix */) {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("clusterServiceClassRef", "name"), spec.ClusterServiceClassRef.Name, msg))
	}

	if "" == spec.ClusterServiceBrokerName {
		allErrs = append(allErrs, field.Required(fldPath.Child("clusterServiceBrokerName"), "clusterServiceBrokerName is required"))
	}

	if "" == spec.ClusterServiceClassRef.Name {
		allErrs = append(allErrs, field.Required(fldPath.Child("clusterServiceClassRef"), "an owning ClusterServiceclass is required"))
	}

	allErrs = append(allErrs, validateCommonServicePlanSpec(spec.CommonServicePlanSpec, fldPath)...)
	return allErrs
}

// ValidateClusterServicePlanUpdate checks that when changing from an older
// ClusterServicePlan to a newer ClusterServicePlan is okay.
func ValidateClusterServicePlanUpdate(new *sc.ClusterServicePlan, old *sc.ClusterServicePlan) field.ErrorList {
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, validateClusterServicePlan(new)...)
	if new.Spec.ExternalID != old.Spec.ExternalID {
		allErrs = append(allErrs, field.Invalid(field.NewPath("spec").Child("externalID"), new.Spec.ExternalID, "externalID cannot change when updating a ClusterServicePlan"))
	}
	return allErrs
}
