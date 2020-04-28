package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/sstarcher/yotascale-sdk-golang/model"
)

func resourceBusinessContext() *schema.Resource {
	return &schema.Resource{
		Create: resourceBusinessContextCreate,
		Read:   resourceBusinessContextRead,
		Update: resourceBusinessContextUpdate,
		Delete: resourceBusinessContextDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"parent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"condition": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"condition": {
							Type:     schema.TypeString,
							Required: true,
						},
						"rule": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:     schema.TypeString,
										Required: true,
									},
									"operator": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeList,
										Required: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceBusinessContextCreate(d *schema.ResourceData, m interface{}) error {
	parent := d.Get("parent").(string)
	_, err := client.CreateContext(parent, terraformToYotascale(d))
	if err != nil {
		return err
	}

	return resourceBusinessContextRead(d, m)
}

func resourceBusinessContextRead(d *schema.ResourceData, m interface{}) error {
	contexts, err := client.ListContexts()
	if err != nil {
		return err
	}

	for _, item := range contexts {
		if item.UUID == d.Id() {
			d.Set("priority", item.Priority)
			d.Set("name", item.Name)
			d.Set("parent", item.ParentUUID)
			d.Set("condition", item.Criteria.Condition)

			var groups []interface{}
			for _, item := range item.Criteria.Rules {
				group := make(map[string]interface{})
				group["condition"] = item.Group.Condition

				var rules []interface{}
				for _, item := range item.Group.Rules {
					rule := make(map[string]interface{})
					rule["key"] = item.Key
					rule["operator"] = item.Operator

					var values []interface{}
					for _, item := range item.ValuesWrapper.Value {
						values = append(values, item)
					}

					rule["value"] = values
					rules = append(rules, rule)
				}
				group["rule"] = rules
				groups = append(groups, group)
			}
			d.Set("group", groups)
			break
		}
	}
	return nil
}

func resourceBusinessContextUpdate(d *schema.ResourceData, m interface{}) error {
	_, err := client.UpdateContext(terraformToYotascale(d))
	if err != nil {
		return err
	}

	return resourceBusinessContextRead(d, m)
}

func resourceBusinessContextDelete(d *schema.ResourceData, m interface{}) error {
	return client.DeleteContext(d.Id())
}

func terraformToYotascale(d *schema.ResourceData) model.InputBusinessContext {
	// d.Set("condition", item.Criteria.Condition)
	data := model.InputBusinessContext{
		UUID:       d.Id(),
		Name:       d.Get("name").(string),
		ParentUUID: d.Get("parent").(string),
	}

	if item, ok := d.GetOk("priority"); ok {
		data.Priority = item.(int32)
	}

	return data
}
