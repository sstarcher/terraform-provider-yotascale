package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	api "github.com/sstarcher/yotascale-sdk-golang"
)

var client *api.Client
var clientErr error

func Provider() *schema.Provider {
	client, clientErr = api.NewClient()
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"yotascale_business_context": resourceBusinessContext(),
		},
	}
}
