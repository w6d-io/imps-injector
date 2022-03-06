/*
Copyright 2020 WILDCARD SA.

Licensed under the WILDCARD SA License, Version 1.0 (the "License");
WILDCARD SA is register in french corporation.
You may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.w6d.io/licenses/LICENSE-1.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is prohibited.
Created on 06/03/2022
*/

package v1alpha1_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/w6d-io/imps-injector/apis/secret/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
)

var _ = Describe("DeepCopy Generated", func() {
	Context("", func() {
		var (
			ipsiSpec   *v1alpha1.ImagePullSecretInjectorSpec
			ipsiStatus *v1alpha1.ImagePullSecretInjectorStatus
			ipsi       *v1alpha1.ImagePullSecretInjector
			ipsil      *v1alpha1.ImagePullSecretInjectorList
			sc         *v1alpha1.SecretConfig
			sas        *v1alpha1.ServiceAccountSelector
		)
		BeforeEach(func() {
			sc = &v1alpha1.SecretConfig{
				Name: "unit-test",
				Labels: map[string]string{
					"label": "value",
				},
				Annotations: map[string]string{
					"annotation": "value",
				},
			}
			sas = &v1alpha1.ServiceAccountSelector{
				Key:      "unit-test-key",
				Operator: v1alpha1.OpContains,
				Values: []string{
					"value",
				},
			}
			ipsiSpec = &v1alpha1.ImagePullSecretInjectorSpec{
				Secrets: []v1alpha1.SecretConfig{
					*sc,
				},
				Namespaces: []string{"unit-test"},
				AnnotationSelector: []v1alpha1.ServiceAccountSelector{
					*sas,
				},
				LabelSelector: []v1alpha1.ServiceAccountSelector{
					*sas,
				},
				NamespaceAnnotationSelector: []v1alpha1.ServiceAccountSelector{
					*sas,
				},
				NamespaceLabelSelector: []v1alpha1.ServiceAccountSelector{
					*sas,
				},
			}
			ipsiStatus = &v1alpha1.ImagePullSecretInjectorStatus{
				Status: pointer.String("unit-test"),
				Conditions: []metav1.Condition{
					{
						Type: "Test",
					},
				},
			}
			ipsi = &v1alpha1.ImagePullSecretInjector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "unit-test",
				},
				Spec:   *ipsiSpec,
				Status: *ipsiStatus,
			}
			ipsil = &v1alpha1.ImagePullSecretInjectorList{}
			ipsil.Items = append(ipsil.Items, *ipsi)
		})
		AfterEach(func() {
			ipsi = nil
			ipsil.Items = []v1alpha1.ImagePullSecretInjector{}
		})
		It("DeepCopyInto", func() {
			By("ipsi", func() {
				out := &v1alpha1.ImagePullSecretInjector{}
				ipsi.DeepCopyInto(out)
				Expect(out).ToNot(BeNil())
			})
			By("ipsil", func() {
				out := &v1alpha1.ImagePullSecretInjectorList{}
				ipsil.DeepCopyInto(out)
				Expect(out).ToNot(BeNil())
			})
			By("ipsiSpec", func() {
				out := &v1alpha1.ImagePullSecretInjectorSpec{}
				ipsiSpec.DeepCopyInto(out)
				Expect(out).ToNot(BeNil())
			})
			By("ipsiStatus", func() {
				out := &v1alpha1.ImagePullSecretInjectorStatus{}
				ipsiStatus.DeepCopyInto(out)
				Expect(out).ToNot(BeNil())
			})
		})
		It("DeepCopy", func() {
			By("ipsi", func() {
				out := ipsi.DeepCopy()
				Expect(out).ToNot(BeNil())
			})
			By("ipsil", func() {
				out := ipsil.DeepCopy()
				Expect(out).ToNot(BeNil())
			})
			By("ipsiSpec", func() {
				out := ipsiSpec.DeepCopy()
				Expect(out).ToNot(BeNil())
			})
			By("ipsiStatus", func() {
				out := ipsiStatus.DeepCopy()
				Expect(out).ToNot(BeNil())
			})
			By("sc", func() {
				out := sc.DeepCopy()
				Expect(out).ToNot(BeNil())
			})
			By("sas", func() {
				out := sas.DeepCopy()
				Expect(out).ToNot(BeNil())
			})
			By("ipsiSpec", func() {
				var n *v1alpha1.ImagePullSecretInjectorSpec
				out := n.DeepCopy()
				Expect(out).To(BeNil())
			})
			By("ipsiStatus", func() {
				var n *v1alpha1.ImagePullSecretInjectorStatus
				out := n.DeepCopy()
				Expect(out).To(BeNil())
			})
			By("sc", func() {
				var n *v1alpha1.SecretConfig
				out := n.DeepCopy()
				Expect(out).To(BeNil())
			})
			By("sas", func() {
				var n *v1alpha1.ServiceAccountSelector
				out := n.DeepCopy()
				Expect(out).To(BeNil())
			})
		})
		It("DeepCopyObject", func() {
			By("ipsi", func() {
				out := ipsi.DeepCopyObject()
				Expect(out).ToNot(BeNil())
			})
			By("ipsil", func() {
				out := ipsil.DeepCopyObject()
				Expect(out).ToNot(BeNil())
			})
		})
		It("DeepCopyObject", func() {
			By("ipsi", func() {
				var n *v1alpha1.ImagePullSecretInjector
				out := n.DeepCopyObject()
				Expect(out).To(BeNil())
			})
			By("ipsil", func() {
				var n *v1alpha1.ImagePullSecretInjectorList
				out := n.DeepCopyObject()
				Expect(out).To(BeNil())
			})
		})
	})
})
