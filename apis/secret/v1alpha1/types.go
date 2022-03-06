/*
Copyright 2020 WILDCARD SA.

Licensed under the WILDCARD SA License, Version 1.0 (the "License");
WILDCARD SA is register in french corporation.
You may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.w6d.io/licenses/LICENSE-1.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is prohibited.
Created on 08/02/2022
*/

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ServiceAccountSelector represents the AND of the selectors represented
// by the scoped-resource selector terms.
type ServiceAccountSelector struct {
	Key string `json:"key,omitempty"`
	// Operator represents a key's relationship to the value.
	// Valid operators are In and NotIn. Defaults to In.
	Operator Operator `json:"operator,omitempty"`
	// Values An array of string values. If the operator is In or NotIn,
	// the values array must be non-empty.
	Values []string `json:"value,omitempty"`
}

// SecretConfig describes the properties of the secrets created in each selected namespace
type SecretConfig struct {
	// Name specifies the name of the secret object
	Name string `json:"name"`
	// Labels specifies additional labels to be put on the Secret object
	Labels map[string]string `json:"labels,omitempty"`
	// Annotations specifies additional annotations to be put on the Secret object
	Annotations map[string]string `json:"annotations,omitempty"`
}

type Operator string
type Status string

// These are valid values for Operator
const (
	OpIn        Operator = "In"
	OpExists    Operator = "Exists"
	OpContains  Operator = "Contains"
	OpStartWith Operator = "StartWith"

	OpNotIn            Operator = "NotIn"
	OpDoesNotExist     Operator = "DoesNotExist"
	OpDoesNotContain   Operator = "DoesNotContain"
	OpDoesNotStartWith Operator = "DoesNotStartWith"

	StatusSynchronized          Status = "Synchronized"
	StatusDisSynchronized       Status = "DisSynchronized"
	StatusPartiallySynchronized Status = "PartiallySynchronized"
)

type serviceAccountList struct {
	r    client.Client
	list *corev1.ServiceAccountList
	imps ImagePullSecretInjectorSpec
}
