# Vault Utility
### Will do the following :
* used to authenticate with vault
    * Kubernetes authentication using (accountservice)
    * Token authentication ( for local test )
* get db secrets
* renew lease automatically until reach max_ttl
* initialize mongodb client automatically with new credentials
* plug multiple repositories to receive the pre-initialized mongodb client.
* read from multiple role (secrets path)


### check main.go on how it works.

### This Utility can be used to retrieve secrets from vault and integrate with any DB. 
