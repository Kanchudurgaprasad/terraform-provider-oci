// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	networkLoadBalancersPolicyDataSourceRepresentation = map[string]interface{}{}

	NetworkLoadBalancersPolicyResourceConfig = ""
)

func TestNetworkLoadBalancerNetworkLoadBalancersPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkLoadBalancerNetworkLoadBalancersPolicyResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	datasourceName := "data.oci_network_load_balancer_network_load_balancers_policies.test_network_load_balancers_policies"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_network_load_balancer_network_load_balancers_policies", "test_network_load_balancers_policies", Required, Create, networkLoadBalancersPolicyDataSourceRepresentation) +
					NetworkLoadBalancersPolicyResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(

					resource.TestCheckResourceAttrSet(datasourceName, "network_load_balancers_policy_collection.#"),
					resource.TestCheckResourceAttr(datasourceName, "network_load_balancers_policy_collection.0.items.#", "3"),
				),
			},
		},
	})
}
