// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a WAFv2 Web ACL Association.
//
// > **NOTE on associating a WAFv2 Web ACL with a Cloudfront distribution:** Do not use this resource to associate a WAFv2 Web ACL with a Cloudfront Distribution. The [AWS API call backing this resource](https://docs.aws.amazon.com/waf/latest/APIReference/API_AssociateWebACL.html) notes that you should use the `webAclId` property on the `cloudfrontDistribution` instead.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"crypto/sha1"
//	"encoding/json"
//	"fmt"
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/apigateway"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/wafv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func sha1Hash(input string) string {
//		hash := sha1.Sum([]byte(input))
//		return hex.EncodeToString(hash[:])
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"openapi": "3.0.1",
//				"info": map[string]interface{}{
//					"title":   "example",
//					"version": "1.0",
//				},
//				"paths": map[string]interface{}{
//					"/path1": map[string]interface{}{
//						"get": map[string]interface{}{
//							"x-amazon-apigateway-integration": map[string]interface{}{
//								"httpMethod":           "GET",
//								"payloadFormatVersion": "1.0",
//								"type":                 "HTTP_PROXY",
//								"uri":                  "https://ip-ranges.amazonaws.com/ip-ranges.json",
//							},
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			exampleRestApi, err := apigateway.NewRestApi(ctx, "exampleRestApi", &apigateway.RestApiArgs{
//				Body: pulumi.String(json0),
//			})
//			if err != nil {
//				return err
//			}
//			exampleDeployment, err := apigateway.NewDeployment(ctx, "exampleDeployment", &apigateway.DeploymentArgs{
//				RestApi: exampleRestApi.ID(),
//				Triggers: pulumi.StringMap{
//					"redeployment": exampleRestApi.Body.ApplyT(func(body string) (pulumi.String, error) {
//						var _zero pulumi.String
//						tmpJSON1, err := json.Marshal(body)
//						if err != nil {
//							return _zero, err
//						}
//						json1 := string(tmpJSON1)
//						return json1, nil
//					}).(pulumi.StringOutput).ApplyT(func(toJSON string) (pulumi.String, error) {
//						return sha1Hash(toJSON), nil
//					}).(pulumi.StringOutput),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleStage, err := apigateway.NewStage(ctx, "exampleStage", &apigateway.StageArgs{
//				Deployment: exampleDeployment.ID(),
//				RestApi:    exampleRestApi.ID(),
//				StageName:  pulumi.String("example"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleWebAcl, err := wafv2.NewWebAcl(ctx, "exampleWebAcl", &wafv2.WebAclArgs{
//				Scope: pulumi.String("REGIONAL"),
//				DefaultAction: &wafv2.WebAclDefaultActionArgs{
//					Allow: nil,
//				},
//				VisibilityConfig: &wafv2.WebAclVisibilityConfigArgs{
//					CloudwatchMetricsEnabled: pulumi.Bool(false),
//					MetricName:               pulumi.String("friendly-metric-name"),
//					SampledRequestsEnabled:   pulumi.Bool(false),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = wafv2.NewWebAclAssociation(ctx, "exampleWebAclAssociation", &wafv2.WebAclAssociationArgs{
//				ResourceArn: exampleStage.Arn,
//				WebAclArn:   exampleWebAcl.Arn,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// WAFv2 Web ACL Association can be imported using `WEB_ACL_ARN,RESOURCE_ARN` e.g.,
//
// ```sh
//
//	$ pulumi import aws:wafv2/webAclAssociation:WebAclAssociation example arn:aws:wafv2:...7ce849ea,arn:aws:apigateway:...ages/name
//
// ```
//
// [1]: https://docs.aws.amazon.com/waf/latest/APIReference/API_AssociateWebACL.html
type WebAclAssociation struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the resource to associate with the web ACL. This must be an ARN of an Application Load Balancer, an Amazon API Gateway stage, or an Amazon Cognito User Pool.
	ResourceArn pulumi.StringOutput `pulumi:"resourceArn"`
	// The Amazon Resource Name (ARN) of the Web ACL that you want to associate with the resource.
	WebAclArn pulumi.StringOutput `pulumi:"webAclArn"`
}

// NewWebAclAssociation registers a new resource with the given unique name, arguments, and options.
func NewWebAclAssociation(ctx *pulumi.Context,
	name string, args *WebAclAssociationArgs, opts ...pulumi.ResourceOption) (*WebAclAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceArn == nil {
		return nil, errors.New("invalid value for required argument 'ResourceArn'")
	}
	if args.WebAclArn == nil {
		return nil, errors.New("invalid value for required argument 'WebAclArn'")
	}
	var resource WebAclAssociation
	err := ctx.RegisterResource("aws:wafv2/webAclAssociation:WebAclAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAclAssociation gets an existing WebAclAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAclAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAclAssociationState, opts ...pulumi.ResourceOption) (*WebAclAssociation, error) {
	var resource WebAclAssociation
	err := ctx.ReadResource("aws:wafv2/webAclAssociation:WebAclAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAclAssociation resources.
type webAclAssociationState struct {
	// The Amazon Resource Name (ARN) of the resource to associate with the web ACL. This must be an ARN of an Application Load Balancer, an Amazon API Gateway stage, or an Amazon Cognito User Pool.
	ResourceArn *string `pulumi:"resourceArn"`
	// The Amazon Resource Name (ARN) of the Web ACL that you want to associate with the resource.
	WebAclArn *string `pulumi:"webAclArn"`
}

type WebAclAssociationState struct {
	// The Amazon Resource Name (ARN) of the resource to associate with the web ACL. This must be an ARN of an Application Load Balancer, an Amazon API Gateway stage, or an Amazon Cognito User Pool.
	ResourceArn pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the Web ACL that you want to associate with the resource.
	WebAclArn pulumi.StringPtrInput
}

func (WebAclAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAclAssociationState)(nil)).Elem()
}

type webAclAssociationArgs struct {
	// The Amazon Resource Name (ARN) of the resource to associate with the web ACL. This must be an ARN of an Application Load Balancer, an Amazon API Gateway stage, or an Amazon Cognito User Pool.
	ResourceArn string `pulumi:"resourceArn"`
	// The Amazon Resource Name (ARN) of the Web ACL that you want to associate with the resource.
	WebAclArn string `pulumi:"webAclArn"`
}

// The set of arguments for constructing a WebAclAssociation resource.
type WebAclAssociationArgs struct {
	// The Amazon Resource Name (ARN) of the resource to associate with the web ACL. This must be an ARN of an Application Load Balancer, an Amazon API Gateway stage, or an Amazon Cognito User Pool.
	ResourceArn pulumi.StringInput
	// The Amazon Resource Name (ARN) of the Web ACL that you want to associate with the resource.
	WebAclArn pulumi.StringInput
}

func (WebAclAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAclAssociationArgs)(nil)).Elem()
}

type WebAclAssociationInput interface {
	pulumi.Input

	ToWebAclAssociationOutput() WebAclAssociationOutput
	ToWebAclAssociationOutputWithContext(ctx context.Context) WebAclAssociationOutput
}

func (*WebAclAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAclAssociation)(nil)).Elem()
}

func (i *WebAclAssociation) ToWebAclAssociationOutput() WebAclAssociationOutput {
	return i.ToWebAclAssociationOutputWithContext(context.Background())
}

func (i *WebAclAssociation) ToWebAclAssociationOutputWithContext(ctx context.Context) WebAclAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAclAssociationOutput)
}

// WebAclAssociationArrayInput is an input type that accepts WebAclAssociationArray and WebAclAssociationArrayOutput values.
// You can construct a concrete instance of `WebAclAssociationArrayInput` via:
//
//	WebAclAssociationArray{ WebAclAssociationArgs{...} }
type WebAclAssociationArrayInput interface {
	pulumi.Input

	ToWebAclAssociationArrayOutput() WebAclAssociationArrayOutput
	ToWebAclAssociationArrayOutputWithContext(context.Context) WebAclAssociationArrayOutput
}

type WebAclAssociationArray []WebAclAssociationInput

func (WebAclAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebAclAssociation)(nil)).Elem()
}

func (i WebAclAssociationArray) ToWebAclAssociationArrayOutput() WebAclAssociationArrayOutput {
	return i.ToWebAclAssociationArrayOutputWithContext(context.Background())
}

func (i WebAclAssociationArray) ToWebAclAssociationArrayOutputWithContext(ctx context.Context) WebAclAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAclAssociationArrayOutput)
}

// WebAclAssociationMapInput is an input type that accepts WebAclAssociationMap and WebAclAssociationMapOutput values.
// You can construct a concrete instance of `WebAclAssociationMapInput` via:
//
//	WebAclAssociationMap{ "key": WebAclAssociationArgs{...} }
type WebAclAssociationMapInput interface {
	pulumi.Input

	ToWebAclAssociationMapOutput() WebAclAssociationMapOutput
	ToWebAclAssociationMapOutputWithContext(context.Context) WebAclAssociationMapOutput
}

type WebAclAssociationMap map[string]WebAclAssociationInput

func (WebAclAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebAclAssociation)(nil)).Elem()
}

func (i WebAclAssociationMap) ToWebAclAssociationMapOutput() WebAclAssociationMapOutput {
	return i.ToWebAclAssociationMapOutputWithContext(context.Background())
}

func (i WebAclAssociationMap) ToWebAclAssociationMapOutputWithContext(ctx context.Context) WebAclAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAclAssociationMapOutput)
}

type WebAclAssociationOutput struct{ *pulumi.OutputState }

func (WebAclAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAclAssociation)(nil)).Elem()
}

func (o WebAclAssociationOutput) ToWebAclAssociationOutput() WebAclAssociationOutput {
	return o
}

func (o WebAclAssociationOutput) ToWebAclAssociationOutputWithContext(ctx context.Context) WebAclAssociationOutput {
	return o
}

// The Amazon Resource Name (ARN) of the resource to associate with the web ACL. This must be an ARN of an Application Load Balancer, an Amazon API Gateway stage, or an Amazon Cognito User Pool.
func (o WebAclAssociationOutput) ResourceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAclAssociation) pulumi.StringOutput { return v.ResourceArn }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the Web ACL that you want to associate with the resource.
func (o WebAclAssociationOutput) WebAclArn() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAclAssociation) pulumi.StringOutput { return v.WebAclArn }).(pulumi.StringOutput)
}

type WebAclAssociationArrayOutput struct{ *pulumi.OutputState }

func (WebAclAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebAclAssociation)(nil)).Elem()
}

func (o WebAclAssociationArrayOutput) ToWebAclAssociationArrayOutput() WebAclAssociationArrayOutput {
	return o
}

func (o WebAclAssociationArrayOutput) ToWebAclAssociationArrayOutputWithContext(ctx context.Context) WebAclAssociationArrayOutput {
	return o
}

func (o WebAclAssociationArrayOutput) Index(i pulumi.IntInput) WebAclAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WebAclAssociation {
		return vs[0].([]*WebAclAssociation)[vs[1].(int)]
	}).(WebAclAssociationOutput)
}

type WebAclAssociationMapOutput struct{ *pulumi.OutputState }

func (WebAclAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebAclAssociation)(nil)).Elem()
}

func (o WebAclAssociationMapOutput) ToWebAclAssociationMapOutput() WebAclAssociationMapOutput {
	return o
}

func (o WebAclAssociationMapOutput) ToWebAclAssociationMapOutputWithContext(ctx context.Context) WebAclAssociationMapOutput {
	return o
}

func (o WebAclAssociationMapOutput) MapIndex(k pulumi.StringInput) WebAclAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WebAclAssociation {
		return vs[0].(map[string]*WebAclAssociation)[vs[1].(string)]
	}).(WebAclAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebAclAssociationInput)(nil)).Elem(), &WebAclAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebAclAssociationArrayInput)(nil)).Elem(), WebAclAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebAclAssociationMapInput)(nil)).Elem(), WebAclAssociationMap{})
	pulumi.RegisterOutputType(WebAclAssociationOutput{})
	pulumi.RegisterOutputType(WebAclAssociationArrayOutput{})
	pulumi.RegisterOutputType(WebAclAssociationMapOutput{})
}
