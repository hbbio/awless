package awsspec

import (
	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"

	"github.com/hbbio/awless/cloud"
	"github.com/hbbio/awless/logger"
	"github.com/hbbio/awless/template/params"
)

type CreateSnapshot struct {
	_           string `action:"create" entity:"snapshot" awsAPI:"ec2" awsCall:"CreateSnapshot" awsInput:"ec2.CreateSnapshotInput" awsOutput:"ec2.Snapshot" awsDryRun:""`
	logger      *logger.Logger
	graph       cloud.GraphAPI
	api         ec2iface.EC2API
	Volume      *string `awsName:"VolumeId" awsType:"awsstr" templateName:"volume"`
	Description *string `awsName:"Description" awsType:"awsstr" templateName:"description"`
}

func (cmd *CreateSnapshot) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("volume"),
		params.Opt("description"),
	))
}

func (cmd *CreateSnapshot) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*ec2.Snapshot).SnapshotId)
}

type DeleteSnapshot struct {
	_      string `action:"delete" entity:"snapshot" awsAPI:"ec2" awsCall:"DeleteSnapshot" awsInput:"ec2.DeleteSnapshotInput" awsOutput:"ec2.DeleteSnapshotOutput" awsDryRun:""`
	logger *logger.Logger
	graph  cloud.GraphAPI
	api    ec2iface.EC2API
	Id     *string `awsName:"SnapshotId" awsType:"awsstr" templateName:"id"`
}

func (cmd *DeleteSnapshot) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("id")))
}

type CopySnapshot struct {
	_            string `action:"copy" entity:"snapshot" awsAPI:"ec2" awsCall:"CopySnapshot" awsInput:"ec2.CopySnapshotInput" awsOutput:"ec2.CopySnapshotOutput" awsDryRun:""`
	logger       *logger.Logger
	graph        cloud.GraphAPI
	api          ec2iface.EC2API
	SourceId     *string `awsName:"SourceSnapshotId" awsType:"awsstr" templateName:"source-id"`
	SourceRegion *string `awsName:"SourceRegion" awsType:"awsstr" templateName:"source-region"`
	Encrypted    *bool   `awsName:"Encrypted" awsType:"awsbool" templateName:"encrypted"`
	Description  *string `awsName:"Description" awsType:"awsstr" templateName:"description"`
}

func (cmd *CopySnapshot) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("source-id"), params.Key("source-region"),
		params.Opt("description", "encrypted"),
	))
}

func (cmd *CopySnapshot) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*ec2.CopySnapshotOutput).SnapshotId)
}
