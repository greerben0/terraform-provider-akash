# Terraform Provider Akash

Run the following command to build the provider

```shell
go build -o terraform-provider-akash
```

## Test example configuration

First, build and install the provider locally.

```shell
make install
```

Then, run the following command to initialize the workspace and apply the sample configuration in the terraform project

```shell
terraform init && terraform apply
```