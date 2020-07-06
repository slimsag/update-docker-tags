package main

import (
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_genericTag(t *testing.T) {
	tests := []struct {
		name  string
		image string
		want  [][]string
	}{
		{
			"prom",
			"FROM quay.io/prometheus/busybox-linux-amd64:latest@sha256:0c38f63cbe19e40123668a48c36466ef72b195e723cbfcbe01e9657a5f14cec6",
			[][]string{{
				" quay.io/prometheus/busybox-linux-amd64:latest@sha256:0c38f63cbe19e40123668a48c36466ef72b195e723cbfcbe01e9657a5f14cec6",
				"quay.io/prometheus/busybox-linux-amd64", "latest",
				"sha256:0c38f63cbe19e40123668a48c36466ef72b195e723cbfcbe01e9657a5f14cec6",
			}},
		},
		{
			"prom2",
			"FROM prom/prometheus:v2.16.0@sha256:e4ca62c0d62f3e886e684806dfe9d4e0cda60d54986898173c1083856cfda0f4 AS upstream",
			[][]string{{
				" prom/prometheus:v2.16.0@sha256:e4ca62c0d62f3e886e684806dfe9d4e0cda60d54986898173c1083856cfda0f4",
				"prom/prometheus", "v2.16.0",
				"sha256:e4ca62c0d62f3e886e684806dfe9d4e0cda60d54986898173c1083856cfda0f4",
			}},
		},
		{
			"golang",
			"FROM golang:1.13-alpine@sha256:ed003971a4809c9ae45afe2d318c24b9e3f6b30864a322877c69a46c504d852c AS builder",
			[][]string{{
				" golang:1.13-alpine@sha256:ed003971a4809c9ae45afe2d318c24b9e3f6b30864a322877c69a46c504d852c",
				"golang", "1.13-alpine",
				"sha256:ed003971a4809c9ae45afe2d318c24b9e3f6b30864a322877c69a46c504d852c",
			}},
		},
		{
			"dont patch filepaths",
			`import("foo/bar")`,
			nil,
		},
	}
	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := regexp.MustCompile(defaultPattern).FindAllStringSubmatch(tst.image, -1)
			if diff := cmp.Diff(tst.want, got); diff != "" {
				t.Errorf("%v", diff)
			}
		})
	}
}
