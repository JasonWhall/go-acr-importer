package main

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/containerregistry/mgmt/containerregistry"
)

func ImportImage(image string, acrClient *containerregistry.RegistriesClient) error {

	// Builds Params to pass to import request
	importParams := ImportImagePreparer(image)
	log.Printf("Attempting to import image: %s", image)

	future, err := acrClient.ImportImage(
		ctx,
		*resourceGroupName,
		*registryName,
		importParams,
	)
	if err != nil {
		return err
	}

	// Adds a timeout as invalid registries are not checked by Azure after import is created
	deadline, cancel := context.WithTimeout(ctx, 120*time.Second)
	defer cancel()

	err = future.WaitForCompletionRef(deadline, acrClient.Client)
	if err != nil {
		return err
	}

	log.Println("Image imported:", image)
	return nil
}

func ImportImagePreparer(image string) containerregistry.ImportImageParameters {

	// Split String on Registry URL and image name with paths
	registryUri := strings.Split(image, "/")[0]
	imageName := strings.SplitN(image, "/", 2)[1]

	importSource := containerregistry.ImportSource{
		RegistryURI: &registryUri,
		SourceImage: &imageName,
	}

	importParams := containerregistry.ImportImageParameters{
		Source:     &importSource,
		Mode:       "Force", // Force import to overwrite existing images
		TargetTags: &[]string{imageName},
	}

	return importParams
}
