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

package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestM(t *testing.T) {
	t.Run("main fail on unknown command", func(t *testing.T) {
		var got int
		OsExit = func(code int) {
			got = code
		}
		os.Args = []string{"imps-injector", "serves"}
		main()
		assert.Equal(t, 1, got, "main should failed")
	})
	t.Run("show help", func(t *testing.T) {
		Help(Cmd, []string{"imps-injector", "help"})
	})

}
