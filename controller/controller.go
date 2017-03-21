package controller

import (
	"errors"
	"fmt"

	"github.com/kubernetes-incubator/service-catalog/contrib/pkg/broker/controller"
	"github.com/kubernetes-incubator/service-catalog/pkg/brokerapi"
	minio "github.com/minio/minio-go"
)

type minioController struct {
	bucketPrefix    string
	minioClient     *minio.Client
	region          string
	accessKeyID     string
	secretAccessKey string
	endpoint        string
}

// CreateController returns a Minio Broker Controller
func CreateController(c Config) (controller.Controller, error) {
	minioClient, err := minio.New(c.Endpoint, c.AccessKeyID, c.SecretAccessKey, c.UseSSL)
	if err != nil {
		return nil, err
	}
	return &minioController{
		bucketPrefix:    c.BucketPrefix,
		minioClient:     minioClient,
		region:          c.Region,
		accessKeyID:     c.AccessKeyID,
		secretAccessKey: c.SecretAccessKey,
		endpoint:        c.Endpoint,
	}, nil
}

// Catalog returns the Minio Broker catalog entries
func (c *minioController) Catalog() (*brokerapi.Catalog, error) {
	return &brokerapi.Catalog{
		Services: []*brokerapi.Service{
			{
				Name:        "minio",
				ID:          "4e689cf1-3861-4a9e-9946-f9e1b7b957cf",
				Description: "Minio",
				Plans: []brokerapi.ServicePlan{
					{
						Name:        "default",
						ID:          "473671df-a9f2-4eb1-92cf-f1497d073771",
						Description: "Minio",
						Free:        true,
					},
				},
			},
		},
	}, nil
}

// CreateServiceInstance
func (c *minioController) CreateServiceInstance(id string, req *brokerapi.CreateServiceInstanceRequest) (*brokerapi.CreateServiceInstanceResponse, error) {
	bucketName := c.bucketName(id)

	found, err := c.minioClient.BucketExists(bucketName)
	if err != nil {
		return nil, err
	}
	if !found {
		err = c.minioClient.MakeBucket(bucketName, c.region)
		if err != nil {
			return nil, err
		}
	}

	return &brokerapi.CreateServiceInstanceResponse{}, nil
}

// GetServiceInstance
func (c *minioController) GetServiceInstance(id string) (string, error) {
	return "", errors.New("Unimplemented")
}

// RemoveServiceInstance
func (c *minioController) RemoveServiceInstance(id string) (*brokerapi.DeleteServiceInstanceResponse, error) {
	bucketName := c.bucketName(id)

	found, err := c.minioClient.BucketExists(bucketName)
	if err != nil {
		return nil, err
	}
	if found {
		err := c.minioClient.RemoveBucket(bucketName)
		if err != nil {
			return nil, err
		}
	}

	return &brokerapi.DeleteServiceInstanceResponse{}, nil
}

// Bind
func (c *minioController) Bind(instanceID string, bindingID string, req *brokerapi.BindingRequest) (*brokerapi.CreateServiceBindingResponse, error) {
	bucketName := c.bucketName(instanceID)

	bindingResponse := &brokerapi.CreateServiceBindingResponse{
		Credentials: brokerapi.Credential{
			"name":     bucketName,
			"region":   c.region,
			"username": c.accessKeyID,
			"password": c.secretAccessKey,
			"endpoint": c.endpoint,
		},
	}
	return bindingResponse, nil
}

// UnBind
func (c *minioController) UnBind(instanceID string, bindingID string) error {
	return nil
}

func (c *minioController) bucketName(instanceID string) string {
	return fmt.Sprintf("%s-%s", c.bucketPrefix, instanceID)
}
