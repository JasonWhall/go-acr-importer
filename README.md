# go-acr-importer

A Go CLI app to help import images into [Azure Container Registry](https://azure.microsoft.com/services/container-registry/)

This app will import public container images specified into your private ACR instance.

## Prequisites

- Azure Container Registry Instance
- Service Principal or User with Contributor permissions on the Container Registry Instance

## Usage

### Build the application

```sh
$ go build .
```

### Run the application

Using azure-cli authenticated session:
```sh
go-acr-importer --ResourceGroupName "MyResourceGroup" \
                --DestinationRegistryName "MyAcrName" \
                --ImageNames "docker.mycompany.io/example/app:v1, docker.mycompany.io/example/app:latest"
```

Using [MSI](https://docs.microsoft.com/azure/active-directory/managed-identities-azure-resources/overview):
```sh
go-acr-importer --ResourceGroupName "MyResourceGroup" \
                --DestinationRegistryName "MyAcrName" \
                --ImageNames "docker.mycompany.io/example/app:v1, docker.mycompany.io/example/app:latest"
                --MSIAuthentication
```

Using Explicit Service Principal Credentials:
```sh
go-acr-importer --ResourceGroupName "MyResourceGroup" \
                --DestinationRegistryName "MyAcrName" \
                --ImageNames "docker.mycompany.io/example/app:v1, docker.mycompany.io/example/app:latest" \
                --SubscriptionId "00000000-0000-0000-0000-000000000000" \
                --TenantId "00000000-0000-0000-0000-000000000000" \
                --ClientId "00000000-0000-0000-0000-000000000000" \
                --ClientSecret "SecretSauce"
```

Help:
```sh
go-acr-importer --help
```