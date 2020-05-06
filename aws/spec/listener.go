package awsspec

import (
	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/elbv2/elbv2iface"
	"github.com/hbbio/awless/cloud"
	"github.com/hbbio/awless/logger"
	"github.com/hbbio/awless/template/params"
)

type CreateListener struct {
	_            string `action:"create" entity:"listener" awsAPI:"elbv2" awsCall:"CreateListener" awsInput:"elbv2.CreateListenerInput" awsOutput:"elbv2.CreateListenerOutput"`
	logger       *logger.Logger
	graph        cloud.GraphAPI
	api          elbv2iface.ELBV2API
	Actiontype   *string `awsName:"DefaultActions[0]Type" awsType:"awsslicestruct" templateName:"actiontype"`
	Targetgroup  *string `awsName:"DefaultActions[0]TargetGroupArn" awsType:"awsslicestruct" templateName:"targetgroup"`
	Loadbalancer *string `awsName:"LoadBalancerArn" awsType:"awsstr" templateName:"loadbalancer"`
	Port         *int64  `awsName:"Port" awsType:"awsint64" templateName:"port"`
	Protocol     *string `awsName:"Protocol" awsType:"awsstr" templateName:"protocol"`
	Certificate  *string `awsName:"Certificates[0]CertificateArn" awsType:"awsslicestruct" templateName:"certificate"`
	Sslpolicy    *string `awsName:"SslPolicy" awsType:"awsstr" templateName:"sslpolicy"`
}

func (cmd *CreateListener) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("actiontype"), params.Key("loadbalancer"), params.Key("port"), params.Key("protocol"), params.Key("targetgroup"),
		params.Opt("certificate", "sslpolicy"),
	))
}

func (cmd *CreateListener) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*elbv2.CreateListenerOutput).Listeners[0].ListenerArn)
}

type AttachListener struct {
	_           string `action:"attach" entity:"listener" awsAPI:"elbv2" awsCall:"AddListenerCertificates" awsInput:"elbv2.AddListenerCertificatesInput" awsOutput:"elbv2.AddListenerCertificatesOutput"`
	logger      *logger.Logger
	graph       cloud.GraphAPI
	api         elbv2iface.ELBV2API
	Id          *string `awsName:"ListenerArn" awsType:"awsstr" templateName:"id"`
	Certificate *string `awsName:"Certificates[0]CertificateArn" awsType:"awsslicestruct" templateName:"certificate"`
}

func (cmd *AttachListener) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("id"), params.Key("certificate")))
}

type DeleteListener struct {
	_      string `action:"delete" entity:"listener" awsAPI:"elbv2" awsCall:"DeleteListener" awsInput:"elbv2.DeleteListenerInput" awsOutput:"elbv2.DeleteListenerOutput"`
	logger *logger.Logger
	graph  cloud.GraphAPI
	api    elbv2iface.ELBV2API
	Id     *string `awsName:"ListenerArn" awsType:"awsstr" templateName:"id"`
}

func (cmd *DeleteListener) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("id")))
}
