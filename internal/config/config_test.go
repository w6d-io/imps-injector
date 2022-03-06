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

package config_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/w6d-io/imps-injector/internal/config"
)

var _ = Describe("Config", func() {
	Context("Init", func() {
		BeforeEach(func() {
		})
		AfterEach(func() {
		})
		It("succeed", func() {
			config.Init()
		})
	})
})
