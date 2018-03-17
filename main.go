package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-kafka/kafka"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: kafka.Provider})
}
