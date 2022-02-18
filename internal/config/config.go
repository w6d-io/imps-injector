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

package config

import "github.com/spf13/viper"

var (
	// Version microservice version
	Version = ""

	// Revision git commit
	Revision = ""

	// Built Date built
	Built = ""

	// CfgFile contain the path of the config file
	CfgFile string

	//// OsExit is hack for unit-test
	//OsExit = os.Exit
	//
	//// SkipValidation toggling the config validation
	//SkipValidation bool
)

func setDefault() {
	viper.SetDefault(ViperKeyMetricsListen, ":8080")
	viper.SetDefault(ViperKeyProbeListen, ":8081")
	viper.SetDefault(ViperKeyLeader, false)
}

func Init() {
	setDefault()
}
