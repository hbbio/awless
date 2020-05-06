package awsspec

import (
	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecr/ecriface"
	"github.com/hbbio/awless/cloud"
	"github.com/hbbio/awless/logger"
	"github.com/hbbio/awless/template/params"
)

type CreateRepository struct {
	_      string `action:"create" entity:"repository" awsAPI:"ecr" awsCall:"CreateRepository" awsInput:"ecr.CreateRepositoryInput" awsOutput:"ecr.CreateRepositoryOutput"`
	logger *logger.Logger
	graph  cloud.GraphAPI
	api    ecriface.ECRAPI
	Name   *string `awsName:"RepositoryName" awsType:"awsstr" templateName:"name"`
}

func (cmd *CreateRepository) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("name")))
}

func (cmd *CreateRepository) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*ecr.CreateRepositoryOutput).Repository.RepositoryArn)
}

type DeleteRepository struct {
	_       string `action:"delete" entity:"repository" awsAPI:"ecr" awsCall:"DeleteRepository" awsInput:"ecr.DeleteRepositoryInput" awsOutput:"ecr.DeleteRepositoryOutput"`
	logger  *logger.Logger
	graph   cloud.GraphAPI
	api     ecriface.ECRAPI
	Name    *string `awsName:"RepositoryName" awsType:"awsstr" templateName:"name"`
	Force   *bool   `awsName:"Force" awsType:"awsbool" templateName:"force"`
	Account *string `awsName:"RegistryId" awsType:"awsstr" templateName:"account"`
}

func (cmd *DeleteRepository) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("name"),
		params.Opt("account", "force"),
	))
}
