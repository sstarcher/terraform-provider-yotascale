resource "yotascale_business_context" "company" {
  name     = "Company"
  priority = 65535
}

resource "yotascale_business_context" "team" {
  name   = "my-team"
  parent = yotascale_business_context.company.id
  condition = "OR"
  group {
    condition = "AND"
    rule {
      key      = "team"
      operator = "Equal"
      value    = ["my-team"]
    }
    rule {
      key      = "service"
      operator = "Equal"
      value    = ["my-service"]
    }
  }

  group {
    condition = "AND"
    rule {
      key      = "team"
      operator = "Equal"
      value    = ["my-team"]
    }
    rule {
      key      = "service"
      operator = "Equal"
      value    = ["my-service2"]
    }
  }
}
