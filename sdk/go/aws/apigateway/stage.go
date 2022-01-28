// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an API Gateway Stage. A stage is a named reference to a deployment, which can be done via the `apigateway.Deployment` resource. Stages can be optionally managed further with the `apigateway.BasePathMapping` resource, `apigateway.DomainName` resource, and `awsApiMethodSettings` resource. For more information, see the [API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-stages.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"crypto/sha1"
// 	"encoding/json"
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/apigateway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func sha1Hash(input string) string {
// 	hash := sha1.Sum([]byte(input))
// 	return hex.EncodeToString(hash[:])
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		tmpJSON0, err := json.Marshal(map[string]interface{}{
// 			"openapi": "3.0.1",
// 			"info": map[string]interface{}{
// 				"title":   "example",
// 				"version": "1.0",
// 			},
// 			"paths": map[string]interface{}{
// 				"/path1": map[string]interface{}{
// 					"get": map[string]interface{}{
// 						"x-amazon-apigateway-integration": map[string]interface{}{
// 							"httpMethod":           "GET",
// 							"payloadFormatVersion": "1.0",
// 							"type":                 "HTTP_PROXY",
// 							"uri":                  "https://ip-ranges.amazonaws.com/ip-ranges.json",
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		json0 := string(tmpJSON0)
// 		exampleRestApi, err := apigateway.NewRestApi(ctx, "exampleRestApi", &apigateway.RestApiArgs{
// 			Body: pulumi.String(json0),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleDeployment, err := apigateway.NewDeployment(ctx, "exampleDeployment", &apigateway.DeploymentArgs{
// 			RestApi: exampleRestApi.ID(),
// 			Triggers: pulumi.StringMap{
// 				"redeployment": exampleRestApi.Body.ApplyT(func(body string) (pulumi.String, error) {
// 					var _zero pulumi.String
// 					tmpJSON1, err := json.Marshal(body)
// 					if err != nil {
// 						return _zero, err
// 					}
// 					json1 := string(tmpJSON1)
// 					return json1, nil
// 				}).(pulumi.StringOutput).ApplyT(func(toJSON string) (pulumi.String, error) {
// 					return sha1Hash(toJSON), nil
// 				}).(pulumi.StringOutput),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleStage, err := apigateway.NewStage(ctx, "exampleStage", &apigateway.StageArgs{
// 			Deployment: exampleDeployment.ID(),
// 			RestApi:    exampleRestApi.ID(),
// 			StageName:  pulumi.String("example"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apigateway.NewMethodSettings(ctx, "exampleMethodSettings", &apigateway.MethodSettingsArgs{
// 			RestApi:    exampleRestApi.ID(),
// 			StageName:  exampleStage.StageName,
// 			MethodPath: pulumi.String("*/*"),
// 			Settings: &apigateway.MethodSettingsSettingsArgs{
// 				MetricsEnabled: pulumi.Bool(true),
// 				LoggingLevel:   pulumi.String("INFO"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Managing the API Logging CloudWatch Log Group
//
// API Gateway provides the ability to [enable CloudWatch API logging](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html). To manage the CloudWatch Log Group when this feature is enabled, the `cloudwatch.LogGroup` resource can be used where the name matches the API Gateway naming convention. If the CloudWatch Log Group previously exists, the `cloudwatch.LogGroup` resource can be imported as a one time operation and recreation of the environment can occur without import.
//
// > The below configuration uses [`dependsOn`](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson) to prevent ordering issues with API Gateway automatically creating the log group first and a variable for naming consistency. Other ordering and naming methodologies may be more appropriate for your environment.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/apigateway"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		stageName := "example"
// 		if param := cfg.Get("stageName"); param != "" {
// 			stageName = param
// 		}
// 		_, err := apigateway.NewRestApi(ctx, "exampleRestApi", nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleLogGroup, err := cloudwatch.NewLogGroup(ctx, "exampleLogGroup", &cloudwatch.LogGroupArgs{
// 			RetentionInDays: pulumi.Int(7),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apigateway.NewStage(ctx, "exampleStage", &apigateway.StageArgs{
// 			StageName: pulumi.String(stageName),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleLogGroup,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// `aws_api_gateway_stage` can be imported using `REST-API-ID/STAGE-NAME`, e.g.,
//
// ```sh
//  $ pulumi import aws:apigateway/stage:Stage example 12345abcde/example
// ```
type Stage struct {
	pulumi.CustomResourceState

	// Enables access logs for the API stage. Detailed below.
	AccessLogSettings StageAccessLogSettingsPtrOutput `pulumi:"accessLogSettings"`
	// Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Specifies whether a cache cluster is enabled for the stage
	CacheClusterEnabled pulumi.BoolPtrOutput `pulumi:"cacheClusterEnabled"`
	// The size of the cache cluster for the stage, if enabled. Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
	CacheClusterSize pulumi.StringPtrOutput `pulumi:"cacheClusterSize"`
	// The identifier of a client certificate for the stage.
	ClientCertificateId pulumi.StringPtrOutput `pulumi:"clientCertificateId"`
	// The ID of the deployment that the stage points to
	Deployment pulumi.StringOutput `pulumi:"deployment"`
	// The description of the stage
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The version of the associated API documentation
	DocumentationVersion pulumi.StringPtrOutput `pulumi:"documentationVersion"`
	// The execution ARN to be used in `lambdaPermission`'s `sourceArn`
	// when allowing API Gateway to invoke a Lambda function,
	// e.g., `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod`
	ExecutionArn pulumi.StringOutput `pulumi:"executionArn"`
	// The URL to invoke the API pointing to the stage,
	// e.g., `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod`
	InvokeUrl pulumi.StringOutput `pulumi:"invokeUrl"`
	// The ID of the associated REST API
	RestApi pulumi.StringOutput `pulumi:"restApi"`
	// The name of the stage
	StageName pulumi.StringOutput `pulumi:"stageName"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// A map that defines the stage variables
	Variables pulumi.StringMapOutput `pulumi:"variables"`
	// The ARN of the WebAcl associated with the Stage.
	WebAclArn pulumi.StringOutput `pulumi:"webAclArn"`
	// Whether active tracing with X-ray is enabled. Defaults to `false`.
	XrayTracingEnabled pulumi.BoolPtrOutput `pulumi:"xrayTracingEnabled"`
}

// NewStage registers a new resource with the given unique name, arguments, and options.
func NewStage(ctx *pulumi.Context,
	name string, args *StageArgs, opts ...pulumi.ResourceOption) (*Stage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Deployment == nil {
		return nil, errors.New("invalid value for required argument 'Deployment'")
	}
	if args.RestApi == nil {
		return nil, errors.New("invalid value for required argument 'RestApi'")
	}
	if args.StageName == nil {
		return nil, errors.New("invalid value for required argument 'StageName'")
	}
	var resource Stage
	err := ctx.RegisterResource("aws:apigateway/stage:Stage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStage gets an existing Stage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StageState, opts ...pulumi.ResourceOption) (*Stage, error) {
	var resource Stage
	err := ctx.ReadResource("aws:apigateway/stage:Stage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Stage resources.
type stageState struct {
	// Enables access logs for the API stage. Detailed below.
	AccessLogSettings *StageAccessLogSettings `pulumi:"accessLogSettings"`
	// Amazon Resource Name (ARN)
	Arn *string `pulumi:"arn"`
	// Specifies whether a cache cluster is enabled for the stage
	CacheClusterEnabled *bool `pulumi:"cacheClusterEnabled"`
	// The size of the cache cluster for the stage, if enabled. Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
	CacheClusterSize *string `pulumi:"cacheClusterSize"`
	// The identifier of a client certificate for the stage.
	ClientCertificateId *string `pulumi:"clientCertificateId"`
	// The ID of the deployment that the stage points to
	Deployment interface{} `pulumi:"deployment"`
	// The description of the stage
	Description *string `pulumi:"description"`
	// The version of the associated API documentation
	DocumentationVersion *string `pulumi:"documentationVersion"`
	// The execution ARN to be used in `lambdaPermission`'s `sourceArn`
	// when allowing API Gateway to invoke a Lambda function,
	// e.g., `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod`
	ExecutionArn *string `pulumi:"executionArn"`
	// The URL to invoke the API pointing to the stage,
	// e.g., `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod`
	InvokeUrl *string `pulumi:"invokeUrl"`
	// The ID of the associated REST API
	RestApi interface{} `pulumi:"restApi"`
	// The name of the stage
	StageName *string `pulumi:"stageName"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll map[string]string `pulumi:"tagsAll"`
	// A map that defines the stage variables
	Variables map[string]string `pulumi:"variables"`
	// The ARN of the WebAcl associated with the Stage.
	WebAclArn *string `pulumi:"webAclArn"`
	// Whether active tracing with X-ray is enabled. Defaults to `false`.
	XrayTracingEnabled *bool `pulumi:"xrayTracingEnabled"`
}

type StageState struct {
	// Enables access logs for the API stage. Detailed below.
	AccessLogSettings StageAccessLogSettingsPtrInput
	// Amazon Resource Name (ARN)
	Arn pulumi.StringPtrInput
	// Specifies whether a cache cluster is enabled for the stage
	CacheClusterEnabled pulumi.BoolPtrInput
	// The size of the cache cluster for the stage, if enabled. Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
	CacheClusterSize pulumi.StringPtrInput
	// The identifier of a client certificate for the stage.
	ClientCertificateId pulumi.StringPtrInput
	// The ID of the deployment that the stage points to
	Deployment pulumi.Input
	// The description of the stage
	Description pulumi.StringPtrInput
	// The version of the associated API documentation
	DocumentationVersion pulumi.StringPtrInput
	// The execution ARN to be used in `lambdaPermission`'s `sourceArn`
	// when allowing API Gateway to invoke a Lambda function,
	// e.g., `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod`
	ExecutionArn pulumi.StringPtrInput
	// The URL to invoke the API pointing to the stage,
	// e.g., `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod`
	InvokeUrl pulumi.StringPtrInput
	// The ID of the associated REST API
	RestApi pulumi.Input
	// The name of the stage
	StageName pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapInput
	// A map that defines the stage variables
	Variables pulumi.StringMapInput
	// The ARN of the WebAcl associated with the Stage.
	WebAclArn pulumi.StringPtrInput
	// Whether active tracing with X-ray is enabled. Defaults to `false`.
	XrayTracingEnabled pulumi.BoolPtrInput
}

func (StageState) ElementType() reflect.Type {
	return reflect.TypeOf((*stageState)(nil)).Elem()
}

type stageArgs struct {
	// Enables access logs for the API stage. Detailed below.
	AccessLogSettings *StageAccessLogSettings `pulumi:"accessLogSettings"`
	// Specifies whether a cache cluster is enabled for the stage
	CacheClusterEnabled *bool `pulumi:"cacheClusterEnabled"`
	// The size of the cache cluster for the stage, if enabled. Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
	CacheClusterSize *string `pulumi:"cacheClusterSize"`
	// The identifier of a client certificate for the stage.
	ClientCertificateId *string `pulumi:"clientCertificateId"`
	// The ID of the deployment that the stage points to
	Deployment interface{} `pulumi:"deployment"`
	// The description of the stage
	Description *string `pulumi:"description"`
	// The version of the associated API documentation
	DocumentationVersion *string `pulumi:"documentationVersion"`
	// The ID of the associated REST API
	RestApi interface{} `pulumi:"restApi"`
	// The name of the stage
	StageName string `pulumi:"stageName"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map that defines the stage variables
	Variables map[string]string `pulumi:"variables"`
	// Whether active tracing with X-ray is enabled. Defaults to `false`.
	XrayTracingEnabled *bool `pulumi:"xrayTracingEnabled"`
}

// The set of arguments for constructing a Stage resource.
type StageArgs struct {
	// Enables access logs for the API stage. Detailed below.
	AccessLogSettings StageAccessLogSettingsPtrInput
	// Specifies whether a cache cluster is enabled for the stage
	CacheClusterEnabled pulumi.BoolPtrInput
	// The size of the cache cluster for the stage, if enabled. Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
	CacheClusterSize pulumi.StringPtrInput
	// The identifier of a client certificate for the stage.
	ClientCertificateId pulumi.StringPtrInput
	// The ID of the deployment that the stage points to
	Deployment pulumi.Input
	// The description of the stage
	Description pulumi.StringPtrInput
	// The version of the associated API documentation
	DocumentationVersion pulumi.StringPtrInput
	// The ID of the associated REST API
	RestApi pulumi.Input
	// The name of the stage
	StageName pulumi.StringInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map that defines the stage variables
	Variables pulumi.StringMapInput
	// Whether active tracing with X-ray is enabled. Defaults to `false`.
	XrayTracingEnabled pulumi.BoolPtrInput
}

func (StageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stageArgs)(nil)).Elem()
}

type StageInput interface {
	pulumi.Input

	ToStageOutput() StageOutput
	ToStageOutputWithContext(ctx context.Context) StageOutput
}

func (*Stage) ElementType() reflect.Type {
	return reflect.TypeOf((*Stage)(nil))
}

func (i *Stage) ToStageOutput() StageOutput {
	return i.ToStageOutputWithContext(context.Background())
}

func (i *Stage) ToStageOutputWithContext(ctx context.Context) StageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageOutput)
}

func (i *Stage) ToStagePtrOutput() StagePtrOutput {
	return i.ToStagePtrOutputWithContext(context.Background())
}

func (i *Stage) ToStagePtrOutputWithContext(ctx context.Context) StagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StagePtrOutput)
}

type StagePtrInput interface {
	pulumi.Input

	ToStagePtrOutput() StagePtrOutput
	ToStagePtrOutputWithContext(ctx context.Context) StagePtrOutput
}

type stagePtrType StageArgs

func (*stagePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Stage)(nil))
}

func (i *stagePtrType) ToStagePtrOutput() StagePtrOutput {
	return i.ToStagePtrOutputWithContext(context.Background())
}

func (i *stagePtrType) ToStagePtrOutputWithContext(ctx context.Context) StagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StagePtrOutput)
}

// StageArrayInput is an input type that accepts StageArray and StageArrayOutput values.
// You can construct a concrete instance of `StageArrayInput` via:
//
//          StageArray{ StageArgs{...} }
type StageArrayInput interface {
	pulumi.Input

	ToStageArrayOutput() StageArrayOutput
	ToStageArrayOutputWithContext(context.Context) StageArrayOutput
}

type StageArray []StageInput

func (StageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Stage)(nil)).Elem()
}

func (i StageArray) ToStageArrayOutput() StageArrayOutput {
	return i.ToStageArrayOutputWithContext(context.Background())
}

func (i StageArray) ToStageArrayOutputWithContext(ctx context.Context) StageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageArrayOutput)
}

// StageMapInput is an input type that accepts StageMap and StageMapOutput values.
// You can construct a concrete instance of `StageMapInput` via:
//
//          StageMap{ "key": StageArgs{...} }
type StageMapInput interface {
	pulumi.Input

	ToStageMapOutput() StageMapOutput
	ToStageMapOutputWithContext(context.Context) StageMapOutput
}

type StageMap map[string]StageInput

func (StageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Stage)(nil)).Elem()
}

func (i StageMap) ToStageMapOutput() StageMapOutput {
	return i.ToStageMapOutputWithContext(context.Background())
}

func (i StageMap) ToStageMapOutputWithContext(ctx context.Context) StageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageMapOutput)
}

type StageOutput struct{ *pulumi.OutputState }

func (StageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Stage)(nil))
}

func (o StageOutput) ToStageOutput() StageOutput {
	return o
}

func (o StageOutput) ToStageOutputWithContext(ctx context.Context) StageOutput {
	return o
}

func (o StageOutput) ToStagePtrOutput() StagePtrOutput {
	return o.ToStagePtrOutputWithContext(context.Background())
}

func (o StageOutput) ToStagePtrOutputWithContext(ctx context.Context) StagePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Stage) *Stage {
		return &v
	}).(StagePtrOutput)
}

type StagePtrOutput struct{ *pulumi.OutputState }

func (StagePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Stage)(nil))
}

func (o StagePtrOutput) ToStagePtrOutput() StagePtrOutput {
	return o
}

func (o StagePtrOutput) ToStagePtrOutputWithContext(ctx context.Context) StagePtrOutput {
	return o
}

func (o StagePtrOutput) Elem() StageOutput {
	return o.ApplyT(func(v *Stage) Stage {
		if v != nil {
			return *v
		}
		var ret Stage
		return ret
	}).(StageOutput)
}

type StageArrayOutput struct{ *pulumi.OutputState }

func (StageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Stage)(nil))
}

func (o StageArrayOutput) ToStageArrayOutput() StageArrayOutput {
	return o
}

func (o StageArrayOutput) ToStageArrayOutputWithContext(ctx context.Context) StageArrayOutput {
	return o
}

func (o StageArrayOutput) Index(i pulumi.IntInput) StageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Stage {
		return vs[0].([]Stage)[vs[1].(int)]
	}).(StageOutput)
}

type StageMapOutput struct{ *pulumi.OutputState }

func (StageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Stage)(nil))
}

func (o StageMapOutput) ToStageMapOutput() StageMapOutput {
	return o
}

func (o StageMapOutput) ToStageMapOutputWithContext(ctx context.Context) StageMapOutput {
	return o
}

func (o StageMapOutput) MapIndex(k pulumi.StringInput) StageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Stage {
		return vs[0].(map[string]Stage)[vs[1].(string)]
	}).(StageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StageInput)(nil)).Elem(), &Stage{})
	pulumi.RegisterInputType(reflect.TypeOf((*StagePtrInput)(nil)).Elem(), &Stage{})
	pulumi.RegisterInputType(reflect.TypeOf((*StageArrayInput)(nil)).Elem(), StageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StageMapInput)(nil)).Elem(), StageMap{})
	pulumi.RegisterOutputType(StageOutput{})
	pulumi.RegisterOutputType(StagePtrOutput{})
	pulumi.RegisterOutputType(StageArrayOutput{})
	pulumi.RegisterOutputType(StageMapOutput{})
}
