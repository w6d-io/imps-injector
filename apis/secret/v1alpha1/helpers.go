/*
Copyright 2020 WILDCARD SA.

Licensed under the WILDCARD SA License, Version 1.0 (the "License");
WILDCARD SA is register in french corporation.
You may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.w6d.io/licenses/LICENSE-1.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is prohibited.
Created on 12/02/2022
*/

package v1alpha1

import (
	"context"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strings"

	corev1 "k8s.io/api/core/v1"

	"github.com/w6d-io/imps-injector/pkg/toolx"
	"github.com/w6d-io/x/logx"
)

func (s *serviceAccountList) OpIn(sa *corev1.ServiceAccount) bool {
	if sa == nil {
		return false
	}
	for _, a := range s.imps.LabelSelector {
		if a.OpIn(sa.Labels) {
			return true
		}
	}
	for _, a := range s.imps.AnnotationSelector {
		if a.OpIn(sa.Annotations) {
			return true
		}
	}
	ns := &corev1.Namespace{}
	if err := s.r.Get(context.Background(), types.NamespacedName{Name: sa.Namespace}, ns); err == nil {
		for _, a := range s.imps.NamespaceLabelSelector {
			if a.OpIn(ns.Labels) {
				return true
			}
		}
		for _, a := range s.imps.NamespaceAnnotationSelector {
			if a.OpIn(ns.Annotations) {
				return true
			}
		}
	}
	return false
}

func (s *serviceAccountList) OpNotIn(sa *corev1.ServiceAccount) bool {
	if sa == nil {
		return true
	}
	for _, a := range s.imps.LabelSelector {
		if !a.OpNotIn(sa.Labels) {
			return false
		}
	}
	for _, a := range s.imps.AnnotationSelector {
		if !a.OpNotIn(sa.Annotations) {
			return false
		}
	}
	ns := &corev1.Namespace{}
	if err := s.r.Get(context.Background(), types.NamespacedName{Name: sa.Namespace}, ns); err == nil {
		for _, a := range s.imps.NamespaceLabelSelector {
			if !a.OpNotIn(ns.Labels) {
				return false
			}
		}
		for _, a := range s.imps.NamespaceAnnotationSelector {
			if !a.OpNotIn(ns.Annotations) {
				return false
			}
		}
	}
	return true
}

func (s *serviceAccountList) OpExist(sa *corev1.ServiceAccount) bool {
	if sa == nil {
		return false
	}
	for _, a := range s.imps.LabelSelector {
		if a.OpExists(sa.Labels) {
			return true
		}
	}
	for _, a := range s.imps.AnnotationSelector {
		if a.OpExists(sa.Annotations) {
			return true
		}
	}
	ns := &corev1.Namespace{}

	if err := s.r.Get(context.Background(), types.NamespacedName{Name: sa.Namespace}, ns); err == nil {
		for _, a := range s.imps.NamespaceLabelSelector {
			if a.OpExists(ns.Labels) {
				return true
			}
		}
		for _, a := range s.imps.NamespaceAnnotationSelector {
			if a.OpExists(ns.Annotations) {
				return true
			}
		}
	}

	return false
}

func (s *serviceAccountList) OpDoesNotExist(sa *corev1.ServiceAccount) bool {
	if sa == nil {
		return true
	}
	for _, a := range s.imps.LabelSelector {
		if !a.OpDoesNotExist(sa.Labels) {
			return false
		}
	}
	for _, a := range s.imps.AnnotationSelector {
		if !a.OpDoesNotExist(sa.Annotations) {
			return false
		}
	}
	ns := &corev1.Namespace{}
	if err := s.r.Get(context.Background(), types.NamespacedName{Name: sa.Namespace}, ns); err == nil {
		for _, a := range s.imps.NamespaceLabelSelector {
			if !a.OpDoesNotExist(ns.Labels) {
				return false
			}
		}
		for _, a := range s.imps.NamespaceAnnotationSelector {
			if !a.OpDoesNotExist(ns.Annotations) {
				return false
			}
		}
	}
	return true
}

func (s *serviceAccountList) OpContains(sa *corev1.ServiceAccount) bool {
	if sa == nil {
		return false
	}
	for _, a := range s.imps.LabelSelector {
		if a.OpContains(sa.Labels) {
			return true
		}
	}
	for _, a := range s.imps.AnnotationSelector {
		if a.OpContains(sa.Annotations) {
			return true
		}
	}
	ns := &corev1.Namespace{}
	if err := s.r.Get(context.Background(), types.NamespacedName{Name: sa.Namespace}, ns); err == nil {
		for _, a := range s.imps.NamespaceLabelSelector {
			if a.OpContains(ns.Labels) {
				return true
			}
		}
		for _, a := range s.imps.NamespaceAnnotationSelector {
			if a.OpContains(ns.Annotations) {
				return true
			}
		}
	}
	return false
}

func (s *serviceAccountList) OpDoesNotContain(sa *corev1.ServiceAccount) bool {
	if sa == nil {
		return true
	}
	for _, a := range s.imps.LabelSelector {
		if !a.OpDoesNotContain(sa.Labels) {
			return false
		}
	}
	for _, a := range s.imps.AnnotationSelector {
		if !a.OpDoesNotContain(sa.Annotations) {
			return false
		}
	}
	ns := &corev1.Namespace{}
	if err := s.r.Get(context.Background(), types.NamespacedName{Name: sa.Namespace}, ns); err == nil {
		for _, a := range s.imps.NamespaceLabelSelector {
			if !a.OpDoesNotContain(ns.Labels) {
				return false
			}
		}
		for _, a := range s.imps.NamespaceAnnotationSelector {
			if !a.OpDoesNotContain(ns.Annotations) {
				return false
			}
		}
	}
	return true
}

func (s *serviceAccountList) OpStartWith(sa *corev1.ServiceAccount) bool {
	if sa == nil {
		return false
	}
	for _, a := range s.imps.LabelSelector {
		if a.OpStartWith(sa.Labels) {
			return true
		}
	}
	for _, a := range s.imps.AnnotationSelector {
		if a.OpStartWith(sa.Annotations) {
			return true
		}
	}
	ns := &corev1.Namespace{}
	if err := s.r.Get(context.Background(), types.NamespacedName{Name: sa.Namespace}, ns); err == nil {
		for _, a := range s.imps.NamespaceLabelSelector {
			if a.OpStartWith(ns.Labels) {
				return true
			}
		}
		for _, a := range s.imps.NamespaceAnnotationSelector {
			if a.OpStartWith(ns.Annotations) {
				return true
			}
		}
	}
	return false
}

func (s *serviceAccountList) OpDoesNotStartWith(sa *corev1.ServiceAccount) bool {
	if sa == nil {
		return true
	}

	for _, a := range s.imps.LabelSelector {
		if !a.OpDoesNotStartWith(sa.Labels) {
			return false
		}
	}
	for _, a := range s.imps.AnnotationSelector {
		if !a.OpDoesNotStartWith(sa.Annotations) {
			return false
		}
	}
	ns := &corev1.Namespace{}
	if err := s.r.Get(context.Background(), types.NamespacedName{Name: sa.Namespace}, ns); err == nil {
		for _, a := range s.imps.NamespaceLabelSelector {
			if !a.OpDoesNotStartWith(ns.Labels) {
				return false
			}
		}
		for _, a := range s.imps.NamespaceAnnotationSelector {
			if !a.OpDoesNotStartWith(ns.Annotations) {
				return false
			}
		}
	}
	return true
}

func (s *ServiceAccountSelector) OpIn(m map[string]string) bool {
	if s.Operator != OpIn {
		return false
	}
	for _, value := range s.Values {
		if v, ok := m[s.Key]; ok && v == value {
			return true
		}
	}
	return false
}

func (s *ServiceAccountSelector) OpNotIn(m map[string]string) bool {
	if s.Operator != OpNotIn {
		return true
	}
	for _, value := range s.Values {
		if v, ok := m[s.Key]; ok && v == value {
			return false
		}
	}
	return true
}

func (s *ServiceAccountSelector) OpExists(m map[string]string) bool {
	if s.Operator != OpExists {
		return false
	}
	_, ok := m[s.Key]
	return ok
}

func (s *ServiceAccountSelector) OpDoesNotExist(m map[string]string) bool {
	if s.Operator != OpDoesNotExist {
		return true
	}
	_, ok := m[s.Key]
	return !ok
}

func (s *ServiceAccountSelector) OpContains(m map[string]string) bool {
	if s.Operator != OpContains {
		return false
	}
	for _, value := range s.Values {
		if v, ok := m[s.Key]; ok && strings.Contains(v, value) {
			return true
		}
	}
	return false
}

func (s *ServiceAccountSelector) OpDoesNotContain(m map[string]string) bool {
	if s.Operator != OpDoesNotContain {
		return true
	}
	for _, value := range s.Values {
		if v, ok := m[s.Key]; ok && strings.Contains(v, value) {
			return false
		}
	}
	return true
}

func (s *ServiceAccountSelector) OpStartWith(m map[string]string) bool {
	if s.Operator != OpStartWith {
		return false
	}
	for _, value := range s.Values {
		if v, ok := m[s.Key]; ok && strings.HasPrefix(v, value) {
			return true
		}
	}
	return false
}

func (s *ServiceAccountSelector) OpDoesNotStartWith(m map[string]string) bool {
	if s.Operator != OpDoesNotStartWith {
		return true
	}
	for _, value := range s.Values {
		if v, ok := m[s.Key]; ok && strings.HasPrefix(v, value) {
			return false
		}
	}
	return true
}

func (in ImagePullSecretInjectorSpec) Match(ctx context.Context, r client.Client, s *corev1.ServiceAccount) *corev1.ServiceAccountList {
	return in.Matches(ctx, r, &corev1.ServiceAccountList{Items: []corev1.ServiceAccount{
		*s,
	}})
}

func (in ImagePullSecretInjectorSpec) Matches(ctx context.Context, r client.Client, l *corev1.ServiceAccountList) *corev1.ServiceAccountList {
	log := logx.WithName(ctx, "Match")
	sal := &serviceAccountList{
		r:    r,
		imps: in,
		list: &corev1.ServiceAccountList{
			TypeMeta: l.TypeMeta,
			ListMeta: l.ListMeta,
		},
	}

	for i := range l.Items {

		if (sal.OpIn(&l.Items[i]) || sal.OpExist(&l.Items[i]) ||
			sal.OpContains(&l.Items[i]) || sal.OpStartWith(&l.Items[i]) ||
			toolx.InArray(l.Items[i].Namespace, in.Namespaces)) ||
			(sal.OpNotIn(&l.Items[i]) && sal.OpDoesNotExist(&l.Items[i]) &&
				sal.OpDoesNotContain(&l.Items[i]) && sal.OpDoesNotStartWith(&l.Items[i])) {
			log.V(3).Info("add sa in list", "name", l.Items[i].Name, "namespace", l.Items[i].Namespace)
			sal.list.Items = append(sal.list.Items, l.Items[i])
			continue
		}
	}
	return sal.list
}
