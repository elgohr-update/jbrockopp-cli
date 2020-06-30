// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

import (
	"testing"
)

func TestConfig_Config_Validate(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		config  *Config
	}{
		{
			failure: false,
			config: &Config{
				Action: "generate",
				File:   "testdata/config.yml",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "view",
				File:   "testdata/config.yml",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
				File:   "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
				File:   "foo.txt",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.Validate()

		if test.failure {
			if err == nil {
				t.Errorf("Validate should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Validate returned err: %v", err)
		}
	}
}
