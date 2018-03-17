package kafka

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"bootstrap_servers": &schema.Schema{
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Required:    true,
				Description: "The list of kafka brokers",
			},
			"timeout": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    false,
				Optional:    true,
				Default:     90,
				Description: "Timeout in seconds",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{},

		ResourcesMap: map[string]*schema.Resource{
			"kafak_topic": resourceKafkaTopic(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var brokers *[]string

	if brokersRaw, ok := d.GetOk("bootstrap_servers"); ok {
		brokerI := brokersRaw.([]interface{})
		log.Printf("[DEBUG] configuring provider with Brokers of size %d", len(brokerI))
		b := make([]string, len(brokerI))
		for i, v := range brokerI {
			b[i] = v.(string)
		}
		log.Printf("[DEBUG] b of size %d", len(b))
		brokers = &b
	} else {
		log.Printf("[ERROR] something wrong? %v , ", d.Get("timeout"))
		return nil, fmt.Errorf("brokers was not set")
	}

	log.Printf("[DEBUG] configuring provider with Brokers @ %v", brokers)
	timeout := d.Get("timeout").(int)

	config := &Config{
		BootstrapServers: brokers,
		Timeout:          timeout,
	}

	log.Printf("[DEBUG] Config @ %v", config)

	return NewClient(config)
}
