package main

import (
	"github.com/gruntwork-io/terratest/modules/helm"
	"testing"
)

func TestHelm(t *testing.T) {

	options := &helm.Options{}
	helm.AddRepo(t, options, "keda", "https://kedacore.github.io/charts")
	helm.RunHelmCommandAndGetOutputE(t, options, "repo", "update")
	helm.RunHelmCommandAndGetStdOutE(t, options, "template", "keda", "kedacore/keda", "--namespace", "keda", "--version", "2.12.0")
}
