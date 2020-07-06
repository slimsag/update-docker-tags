package main

import (
	"regexp"
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
		{
			"golang",
			"FROM golang:1.13-alpine@sha256:ed003971a4809c9ae45afe2d318c24b9e3f6b30864a322877c69a46c504d852c AS builder",
			"golang:1.13-alpine",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := regexp.MustCompile(defaultPattern).FindAllStringSubmatch(tt.image, -1)
			if got[0][1] != tt.repo {
				t.Errorf("got = %v, want %v", got, tt.repo)
			}
		})
	}
}
