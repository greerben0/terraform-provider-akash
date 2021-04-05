terraform {
  required_providers {
    akash = {
      version = "0.0.1"
      source  = "bengineer.it/akash/akash"
    }
  }
}

provider "akash" {
  alias = "mainnet"
  network = "mainnet"

  akash_version = "0.10.2-rc1"
  chain_id = "akashnet-2"
}

data "akash_key" "tenent" {
  name = "tenent"
}

data akash_account tenent_account { 
  address =  data.akash_key.tenent.address
}

output "tenent_wallet_address" {
  value = data.akash_key.tenent.address
}

output tenent_wallet_balance { 
  value = data.akash_account.tenent_account.balance
}