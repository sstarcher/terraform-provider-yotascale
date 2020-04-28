module github.com/sstarcher/terraform-provider-yotascale

go 1.14

require (
	github.com/hashicorp/terraform-plugin-sdk v1.9.1
	github.com/sirupsen/logrus v1.5.0 // indirect
	github.com/sstarcher/yotascale-sdk-golang v0.0.0-00010101000000-000000000000
)

replace github.com/sstarcher/yotascale-sdk-golang => ./../yotascale-sdk-golang
