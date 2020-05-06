package awsspec

import (
	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/route53/route53iface"
	"github.com/hbbio/awless/cloud"
	"github.com/hbbio/awless/logger"
	"github.com/hbbio/awless/template/params"
)

type CreateZone struct {
	_               string `action:"create" entity:"zone" awsAPI:"route53" awsCall:"CreateHostedZone" awsInput:"route53.CreateHostedZoneInput" awsOutput:"route53.CreateHostedZoneOutput"`
	logger          *logger.Logger
	graph           cloud.GraphAPI
	api             route53iface.Route53API
	Callerreference *string `awsName:"CallerReference" awsType:"awsstr" templateName:"callerreference"`
	Name            *string `awsName:"Name" awsType:"awsstr" templateName:"name"`
	Delegationsetid *string `awsName:"DelegationSetId" awsType:"awsstr" templateName:"delegationsetid"`
	Comment         *string `awsName:"HostedZoneConfig.Comment" awsType:"awsstr" templateName:"comment"`
	Isprivate       *bool   `awsName:"HostedZoneConfig.PrivateZone" awsType:"awsbool" templateName:"isprivate"`
	Vpcid           *string `awsName:"VPC.VPCId" awsType:"awsstr" templateName:"vpcid"`
	Vpcregion       *string `awsName:"VPC.VPCRegion" awsType:"awsstr" templateName:"vpcregion"`
}

func (cmd *CreateZone) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("callerreference"), params.Key("name"),
		params.Opt("comment", "delegationsetid", "isprivate", "vpcid", "vpcregion"),
	))
}

func (cmd *CreateZone) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*route53.CreateHostedZoneOutput).HostedZone.Id)
}

type DeleteZone struct {
	_      string `action:"delete" entity:"zone" awsAPI:"route53" awsCall:"DeleteHostedZone" awsInput:"route53.DeleteHostedZoneInput" awsOutput:"route53.DeleteHostedZoneOutput"`
	logger *logger.Logger
	graph  cloud.GraphAPI
	api    route53iface.Route53API
	Id     *string `awsName:"Id" awsType:"awsstr" templateName:"id"`
}

func (cmd *DeleteZone) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("id")))
}
