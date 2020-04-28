
all: terraform-provider-yotascale

rebuild: remove terraform-provider-yotascale

remove:
	rm -f terraform-provider-yotascale

terraform-provider-yotascale:
	go build .

test: remove all
	terraform init
	TF_LOG=1 terraform plan


