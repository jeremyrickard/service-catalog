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

// This file was automatically generated by informer-gen

package internalversion

import (
	internalinterfaces "github.com/kubernetes-incubator/service-catalog/pkg/client/informers_generated/internalversion/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ClusterServiceBrokers returns a ClusterServiceBrokerInformer.
	ClusterServiceBrokers() ClusterServiceBrokerInformer
	// ClusterServiceClasses returns a ClusterServiceClassInformer.
	ClusterServiceClasses() ClusterServiceClassInformer
	// ClusterServicePlans returns a ClusterServicePlanInformer.
	ClusterServicePlans() ClusterServicePlanInformer
	// ServiceBindings returns a ServiceBindingInformer.
	ServiceBindings() ServiceBindingInformer
	// ServiceInstances returns a ServiceInstanceInformer.
	ServiceInstances() ServiceInstanceInformer
	// ServicePlans returns a ServicePlanInformer.
	ServicePlans() ServicePlanInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ClusterServiceBrokers returns a ClusterServiceBrokerInformer.
func (v *version) ClusterServiceBrokers() ClusterServiceBrokerInformer {
	return &clusterServiceBrokerInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterServiceClasses returns a ClusterServiceClassInformer.
func (v *version) ClusterServiceClasses() ClusterServiceClassInformer {
	return &clusterServiceClassInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterServicePlans returns a ClusterServicePlanInformer.
func (v *version) ClusterServicePlans() ClusterServicePlanInformer {
	return &clusterServicePlanInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ServiceBindings returns a ServiceBindingInformer.
func (v *version) ServiceBindings() ServiceBindingInformer {
	return &serviceBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceInstances returns a ServiceInstanceInformer.
func (v *version) ServiceInstances() ServiceInstanceInformer {
	return &serviceInstanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServicePlans returns a ServicePlanInformer.
func (v *version) ServicePlans() ServicePlanInformer {
	return &servicePlanInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
