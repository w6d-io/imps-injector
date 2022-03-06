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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Helper unit-test", func() {
	Context("OpIn", func() {
		It("succeed for LabelSelector", func() {
			in := v1alpha1.ImagePullSecretInjectorSpec{
				LabelSelector: []v1alpha1.ServiceAccountSelector{
					// Op : In
					{
						Key:      "test-op-in",
						Operator: v1alpha1.OpIn,
						Values: []string{
							"test-ok",
						},
					},
				},
			}
			sal := &corev1.ServiceAccountList{
				Items: []corev1.ServiceAccount{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "test-op-in-label",
							Namespace: "unit-test",
							Labels: map[string]string{
								"test-op-in": "test-ok",
							},
						},
					},
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "test-op-in-label",
							Namespace: "unit-test",
							Labels: map[string]string{
								"test-op-in": "test-ko",
							},
						},
					},
				},
			}
			res := in.Matches(ctx, k8sClient, sal)
			Expect(res).ToNot(BeNil())
			Expect(len(res.Items)).To(Equal(1))
		})
	})
	Context("OpContains", func() {
		It("succeed for LabelSelector", func() {
			in := v1alpha1.ImagePullSecretInjectorSpec{
				LabelSelector: []v1alpha1.ServiceAccountSelector{
					// Op : In
					{
						Key:      "test-op-in",
						Operator: v1alpha1.OpContains,
						Values: []string{
							"ok",
						},
					},
				},
			}
			sal := &corev1.ServiceAccountList{
				Items: []corev1.ServiceAccount{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "test-op-in-label",
							Namespace: "unit-test",
							Labels: map[string]string{
								"test-op-in": "test-ok",
							},
						},
					},
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "test-op-in-label",
							Namespace: "unit-test",
							Labels: map[string]string{
								"test-op-in": "test-ko",
							},
						},
					},
				},
			}
			res := in.Matches(ctx, k8sClient, sal)
			Expect(res).ToNot(BeNil())
			Expect(len(res.Items)).To(Equal(1))
		})
	})
	Context("OpStartWith", func() {
		It("succeed for LabelSelector", func() {
			in := v1alpha1.ImagePullSecretInjectorSpec{
				LabelSelector: []v1alpha1.ServiceAccountSelector{
					// Op : In
					{
						Key:      "test-op-in",
						Operator: v1alpha1.OpStartWith,
						Values: []string{
							"test",
						},
					},
				},
			}
			sal := &corev1.ServiceAccountList{
				Items: []corev1.ServiceAccount{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "test-op-in-label",
							Namespace: "unit-test",
							Labels: map[string]string{
								"test-op-in": "test-ok",
							},
						},
					},
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "test-op-in-label",
							Namespace: "unit-test",
							Labels: map[string]string{
								"test-op-in": "test-ko",
							},
						},
					},
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "test-op-in-label",
							Namespace: "unit-test",
							Labels: map[string]string{
								"test-op-in": "ko-test",
							},
						},
					},
				},
			}
			res := in.Matches(ctx, k8sClient, sal)
			Expect(res).ToNot(BeNil())
			Expect(len(res.Items)).To(Equal(2))
		})
	})
	Context("OpExists", func() {
		It("succeed for LabelSelector", func() {
			in := v1alpha1.ImagePullSecretInjectorSpec{
				LabelSelector: []v1alpha1.ServiceAccountSelector{
					// Op : In
					{
						Key:      "test-op-exist",
						Operator: v1alpha1.OpExists,
					},
				},
			}
			sal := &corev1.ServiceAccountList{
				Items: []corev1.ServiceAccount{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "test-op-exist-label",
							Namespace: "unit-test",
							Labels: map[string]string{
								"test-op-exist": "test-ok",
							},
						},
					},
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "test-op-exist-label",
							Namespace: "unit-test",
							Labels: map[string]string{
								"test-op-exist": "test-ko",
							},
						},
					},
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "test-op-exist-label",
							Namespace: "unit-test",
							Labels: map[string]string{
								"test-op-exist-1": "test-ko",
							},
						},
					},
				},
			}
			res := in.Matches(ctx, k8sClient, sal)
			Expect(res).ToNot(BeNil())
			Expect(len(res.Items)).To(Equal(2))
		})
	})
	Context("LabelSelector", func() {
		It("succeed for OpNotIn", func() {
			in := v1alpha1.ImagePullSecretInjectorSpec{
				LabelSelector: []v1alpha1.ServiceAccountSelector{
					// Op : NotIn
					{
						Key:      "test-op-in",
						Operator: v1alpha1.OpNotIn,
						Values: []string{
							"test-ok",
						},
					},
				},
			}
			sal := &corev1.ServiceAccountList{
				Items: []corev1.ServiceAccount{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "test-not-op-in-label",
							Namespace: "unit-test",
							Labels: map[string]string{
								"test-not-op-in": "test-ok",
							},
						},
					},
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "test-op-in-label",
							Namespace: "unit-test",
							Labels: map[string]string{
								"test-op-in": "test-ok",
							},
						},
					},
				},
			}
			res := in.Matches(ctx, k8sClient, sal)
			Expect(res).ToNot(BeNil())
			Expect(len(res.Items)).To(Equal(1))
		})
	})
	Context("LabelSelector", func() {
		It("succeed for OpDoesNotContain", func() {
			in := v1alpha1.ImagePullSecretInjectorSpec{
				LabelSelector: []v1alpha1.ServiceAccountSelector{
					// Op : NotIn
					{
						Key:      "test-op-in",
						Operator: v1alpha1.OpDoesNotContain,
						Values: []string{
							"ok",
						},
					},
				},
			}
			sal := &corev1.ServiceAccountList{
				Items: []corev1.ServiceAccount{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "test-not-op-in-label",
							Namespace: "unit-test",
							Labels: map[string]string{
								"test-op-in": "test-ok",
							},
						},
					},
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "test-op-in-label",
							Namespace: "unit-test",
							Labels: map[string]string{
								"test-op-in": "test-ko",
							},
						},
					},
				},
			}
			res := in.Matches(ctx, k8sClient, sal)
			Expect(res).ToNot(BeNil())
			Expect(len(res.Items)).To(Equal(1))
		})
	})
	Context("LabelSelector", func() {
		It("succeed for OpDoesNotStartWith", func() {
			in := v1alpha1.ImagePullSecretInjectorSpec{
				LabelSelector: []v1alpha1.ServiceAccountSelector{
					// Op : NotIn
					{
						Key:      "test-op-in",
						Operator: v1alpha1.OpDoesNotStartWith,
						Values: []string{
							"test",
						},
					},
				},
			}
			sal := &corev1.ServiceAccountList{
				Items: []corev1.ServiceAccount{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "test-not-op-in-label",
							Namespace: "unit-test",
							Labels: map[string]string{
								"test-op-in": "test-ok",
							},
						},
					},
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "test-op-in-label",
							Namespace: "unit-test",
							Labels: map[string]string{
								"test-op-in": "test-ko",
							},
						},
					},
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "test-op-in-label",
							Namespace: "unit-test",
							Labels: map[string]string{
								"test-op-in": "ko-test-ko",
							},
						},
					},
				},
			}
			res := in.Matches(ctx, k8sClient, sal)
			Expect(res).ToNot(BeNil())
			Expect(len(res.Items)).To(Equal(1))
		})
	})
	Context("LabelSelector", func() {
		It("succeed for OpDoesNotExist", func() {
			in := v1alpha1.ImagePullSecretInjectorSpec{
				LabelSelector: []v1alpha1.ServiceAccountSelector{
					// Op : NotIn
					{
						Key:      "test-op-in",
						Operator: v1alpha1.OpDoesNotExist,
					},
				},
			}
			sal := &corev1.ServiceAccountList{
				Items: []corev1.ServiceAccount{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "test-not-op-in-label",
							Namespace: "unit-test",
							Labels: map[string]string{
								"test-not-op-in": "test-ok",
							},
						},
					},
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "test-op-in-label",
							Namespace: "unit-test",
							Labels: map[string]string{
								"test-op-in": "test-ok",
							},
						},
					},
				},
			}
			res := in.Matches(ctx, k8sClient, sal)
			Expect(res).ToNot(BeNil())
			Expect(len(res.Items)).To(Equal(1))
		})
	})
	Context("Namespaces", func() {
		It("succeed for OpDoesNotExist", func() {
			in := v1alpha1.ImagePullSecretInjectorSpec{
				Namespaces: []string{
					"unit-test",
					"default",
				},
			}
			sa := &corev1.ServiceAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-not-op-in-label",
					Namespace: "unit-test",
					Labels: map[string]string{
						"test-not-op-in": "test-ok",
					},
				},
			}
			res := in.Match(ctx, k8sClient, sa)
			Expect(res).ToNot(BeNil())
			Expect(len(res.Items)).To(Equal(1))
		})
	})
})
