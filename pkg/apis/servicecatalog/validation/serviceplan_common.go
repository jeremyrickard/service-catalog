/*
Copyright 2018 The Kubernetes Authors.

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
	"regexp"

	utilvalidation "k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"

	sc "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
)

const commonServicePlanNameFmt string = `[-.a-zA-Z0-9]+`
const commonServicePlanNameMaxLength int = 63

var servicePlanNameRegexp = regexp.MustCompile("^" + commonServicePlanNameFmt + "$")

// validateCommonServicePlanName is the validation function for both ClusterServicePlan
// and ServicePlan names.
func validateCommonServicePlanName(value string, prefix bool) []string {
	var errs []string
	if len(value) > commonServicePlanNameMaxLength {
		errs = append(errs, utilvalidation.MaxLenError(commonServicePlanNameMaxLength))
	}
	if !servicePlanNameRegexp.MatchString(value) {
		errs = append(errs, utilvalidation.RegexError(commonServicePlanNameFmt, "plan-name-40d-0983-1b89"))
	}

	return errs
}

func validateCommonServicePlanSpec(spec sc.CommonServicePlanSpec, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	if "" == spec.ExternalID {
		allErrs = append(allErrs, field.Required(fldPath.Child("externalID"), "externalID is required"))
	}

	if "" == spec.Description {
		allErrs = append(allErrs, field.Required(fldPath.Child("description"), "description is required"))
	}

	for _, msg := range validateCommonServicePlanName(spec.ExternalName, false /* prefix */) {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("externalName"), spec.ExternalName, msg))
	}

	for _, msg := range validateExternalID(spec.ExternalID) {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("externalID"), spec.ExternalID, msg))
	}

	return allErrs
}
