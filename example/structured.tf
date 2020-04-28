locals {
  data = {
    "team1" = [
      "simple",
    ],
    "team2" = [
      "hello-world",
    ],
    "team3" = [
      "so-many-services",
      "yet-another-service",
      "not-invented-here",
      "dont-keep-it-simple",
      "distributed-monolith",
    ],
  }


  servicesMap = flatten([
    for key, services in local.data : [
      for item, service in local.data[key] : {
        team    = key
        service = service
      }
    ]
  ])
}

resource "yotascale_business_context" "root" {
  name     = "root"
  priority = 65535
}

resource "yotascale_business_context" "teams" {
  for_each = toset(keys(local.data))

  name   = each.key
  parent = yotascale_business_context.root.id
}

resource "yotascale_business_context" "services" {
  for_each = {
    for item in local.servicesMap : "${item.team}.${item.service}" => item
  }

  name   = each.value.service
  parent = yotascale_business_context.teams[each.value.team].id

  condition = "OR"
  group {
    condition = "AND"
    rule {
      key      = "team"
      operator = "Equal"
      value    = [each.value.team]
    }
    rule {
      key      = "service"
      operator = "Equal"
      value    = [each.value.service]
    }
  }

  group {
    condition = "AND"
    rule {
      key      = "team label"
      operator = "Equal"
      value    = [each.value.team]
    }
    rule {
      key      = "service label"
      operator = "Equal"
      value    = [each.value.service]
    }
  }
}

# Unallocated
resource "yotascale_business_context" "unallocated" {
  for_each = toset(keys(local.data))

  name   = "Unallocated"
  parent = yotascale_business_context.teams[each.key].id

  group {
    condition = "OR"
    rule {
      key      = "team"
      operator = "Equal"
      value    = [each.key]
    }
    rule {
      key      = "team label"
      operator = "Equal"
      value    = [each.key]
    }
  }
}
