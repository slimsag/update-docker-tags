package main

import (
	"testing"
)

func Test_genericTag(t *testing.T) {
	tests := []struct {
		name  string
		image string
		repo  string
	}{
		{
			"prom",
			"FROM quay.io/prometheus/busybox-linux-amd64:latest@sha256:0c38f63cbe19e40123668a48c36466ef72b195e723cbfcbe01e9657a5f14cec6",
			"quay.io/prometheus/busybox-linux-amd64",
		},

		{
			"prom2",
			"FROM prom/prometheus:v2.16.0@sha256:e4ca62c0d62f3e886e684806dfe9d4e0cda60d54986898173c1083856cfda0f4 AS upstream",
			"prom/prometheus",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GENERIC_TAG_PATTERN.FindAllStringSubmatch(tt.image, -1)
			if got[0][1] != tt.repo {
				t.Errorf("got = %v, want %v", got, tt.repo)
			}
		})
	}
}
