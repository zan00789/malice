package container

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/reference"
	"github.com/docker/docker/registry"
	apiclient "github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/container"
	networktypes "github.com/docker/engine-api/types/network"
	"github.com/maliceio/malice/malice/docker/client"
)

func pullImage(ctx context.Context, docker *client.Docker, image string, out io.Writer) error {
	ref, err := reference.ParseNamed(image)
	if err != nil {
		return err
	}

	// Resolve the Repository name from fqn to RepositoryInfo
	_, err = registry.ParseRepositoryInfo(ref)
	if err != nil {
		return err
	}

	// authConfig := dockerCli.ResolveAuthConfig(ctx, repoInfo.Index)
	// encodedAuth, err := client.EncodeAuthToBase64(authConfig)
	// if err != nil {
	// 	return err
	// }

	options := types.ImageCreateOptions{
	// RegistryAuth: encodedAuth,
	}

	responseBody, err := docker.Client.ImageCreate(ctx, image, options)
	if err != nil {
		return err
	}
	defer responseBody.Close()

	return jsonmessage.DisplayJSONMessagesStream(
		responseBody,
		out,
		os.Stdout.Fd(),
		true,
		nil)
}

type cidFile struct {
	path    string
	file    *os.File
	written bool
}

func (cid *cidFile) Close() error {
	cid.file.Close()

	if !cid.written {
		if err := os.Remove(cid.path); err != nil {
			return fmt.Errorf("failed to remove the CID file '%s': %s \n", cid.path, err)
		}
	}

	return nil
}

func (cid *cidFile) Write(id string) error {
	if _, err := cid.file.Write([]byte(id)); err != nil {
		return fmt.Errorf("Failed to write the container ID to the file: %s", err)
	}
	cid.written = true
	return nil
}

func newCIDFile(path string) (*cidFile, error) {
	if _, err := os.Stat(path); err == nil {
		return nil, fmt.Errorf("Container ID file found, make sure the other container isn't running or delete %s", path)
	}

	f, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to create the container ID file: %s", err)
	}

	return &cidFile{path: path, file: f}, nil
}

// createContainer
func createContainer(docker *client.Docker, ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *networktypes.NetworkingConfig, cidfile, name string) (*types.ContainerCreateResponse, error) {
	stderr := os.Stderr

	var containerIDFile *cidFile
	if cidfile != "" {
		var err error
		if containerIDFile, err = newCIDFile(cidfile); err != nil {
			return nil, err
		}
		defer containerIDFile.Close()
	}

	// var trustedRef reference.Canonical
	_, ref, err := reference.ParseIDOrReference(config.Image)
	if err != nil {
		return nil, err
	}
	// if ref != nil {
	// 	ref = reference.WithDefaultTag(ref)

	// 	if ref, ok := ref.(reference.NamedTagged); ok {
	// 		var err error
	// 		trustedRef, err = docker.TrustedReference(ctx, ref)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		config.Image = trustedRef.String()
	// 	}
	// }

	//create the container
	response, err := docker.Client.ContainerCreate(ctx, config, hostConfig, networkingConfig, name)

	//if image not found try to pull it
	if err != nil {
		if apiclient.IsErrImageNotFound(err) && ref != nil {
			fmt.Fprintf(stderr, "Unable to find image '%s' locally\n", ref.String())

			// we don't want to write to stdout anything apart from container.ID
			if err = pullImage(ctx, docker, config.Image, stderr); err != nil {
				return nil, err
			}
			// if ref, ok := ref.(reference.NamedTagged); ok != nil {
			// 	if err := docker.TagTrusted(ctx, trustedRef, ref); err != nil {
			// 	return nil, err
			// 	}
			// }
			// Retry
			var retryErr error
			response, retryErr = docker.Client.ContainerCreate(ctx, config, hostConfig, networkingConfig, name)
			if retryErr != nil {
				return nil, retryErr
			}
		} else {
			return nil, err
		}
	}

	for _, warning := range response.Warnings {
		fmt.Fprintf(stderr, "WARNING: %s\n", warning)
	}
	if containerIDFile != nil {
		if err = containerIDFile.Write(response.ID); err != nil {
			return nil, err
		}
	}
	return &response, nil
}