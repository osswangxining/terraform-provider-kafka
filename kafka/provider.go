package kafka

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"kafka_bin_path": &schema.Schema{
				Type:        schema.TypeString,
				Required:    false,
				Optional:    true,
				Default:     "",
				Description: "custom path of Kafka executables",
			},
			"zookeeper": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "zookeeper address (host:port)",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{},

		ResourcesMap: map[string]*schema.Resource{
			"kafak_topic": resoureKafkaTopic(),
		},

		ConfigureFunc: providerConfigure,
	}
}
