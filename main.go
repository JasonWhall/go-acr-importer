package main

import (
	"context"
	"flag"
	"log"
	"strings"
	"sync"
)

var (
	wg                sync.WaitGroup
	ctx               = context.Background()
	subscriptionID    = flag.String("SubscriptionId", "", "Subscription ID")
	tenantId          = flag.String("TenantId", "", "Tenant ID")
	clientId          = flag.String("ClientId", "", "Client App ID")
	clientSecret      = flag.String("ClientSecret", "", "Client App Secret")
	resourceGroupName = flag.String("ResourceGroupName", "", "Resource Group Of Destination ACR")
	registryName      = flag.String("DestinationRegistryName", "", "Name of Destination ACR")
	images            = flag.String("ImageNames", "", "Names of images. Expected in format \"example.acr/image:test,example.cr/images/image:latest\"")
	useMsi            = flag.Bool("MSIAuthentication", false, "Use MSI endpoint authentication")
)

func main() {

	flag.Parse()

	// Login to ACR
	acrClient, err := NewContainerRegistryClient()
	if err != nil {
		log.Fatalf("Unable to Log into Container Registry: %v", err)
	}

	// Find Container Registry Desintation
	acr, err := acrClient.Get(ctx, *resourceGroupName, *registryName)
	if err != nil {
		log.Fatalf("Unable to find Container Registry: %v", err)
	}

	log.Printf("Found Registry: %s", *acr.Name)

	// Loop through images to import
	imagesArray := strings.Split(*images, ",")

	for _, image := range imagesArray {

		wg.Add(1)

		log.Printf("Importing Image: %s", image)

		go func(image string) {

			defer wg.Done()

			err := ImportImage(image, acrClient)
			if err != nil {
				log.Printf("Failed to import image: %s\n %v", image, err)
			}

		}(image)
	}

	wg.Wait()
}
