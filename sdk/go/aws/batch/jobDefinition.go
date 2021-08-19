// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package batch

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Batch Job Definition resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/batch"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := batch.NewJobDefinition(ctx, "test", &batch.JobDefinitionArgs{
// 			ContainerProperties: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "	\"command\": [\"ls\", \"-la\"],\n", "	\"image\": \"busybox\",\n", "	\"memory\": 1024,\n", "	\"vcpus\": 1,\n", "	\"volumes\": [\n", "      {\n", "        \"host\": {\n", "          \"sourcePath\": \"/tmp\"\n", "        },\n", "        \"name\": \"tmp\"\n", "      }\n", "    ],\n", "	\"environment\": [\n", "		{\"name\": \"VARNAME\", \"value\": \"VARVAL\"}\n", "	],\n", "	\"mountPoints\": [\n", "		{\n", "          \"sourceVolume\": \"tmp\",\n", "          \"containerPath\": \"/tmp\",\n", "          \"readOnly\": false\n", "        }\n", "	],\n", "    \"ulimits\": [\n", "      {\n", "        \"hardLimit\": 1024,\n", "        \"name\": \"nofile\",\n", "        \"softLimit\": 1024\n", "      }\n", "    ]\n", "}\n", "\n")),
// 			Type: pulumi.String("container"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Fargate Platform Capability
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/batch"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		assumeRolePolicy, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
// 			Statements: []iam.GetPolicyDocumentStatement{
// 				iam.GetPolicyDocumentStatement{
// 					Actions: []string{
// 						"sts:AssumeRole",
// 					},
// 					Principals: []iam.GetPolicyDocumentStatementPrincipal{
// 						iam.GetPolicyDocumentStatementPrincipal{
// 							Type: "Service",
// 							Identifiers: []string{
// 								"ecs-tasks.amazonaws.com",
// 							},
// 						},
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ecsTaskExecutionRole, err := iam.NewRole(ctx, "ecsTaskExecutionRole", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(assumeRolePolicy.Json),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicyAttachment(ctx, "ecsTaskExecutionRolePolicy", &iam.RolePolicyAttachmentArgs{
// 			Role:      ecsTaskExecutionRole.Name,
// 			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = batch.NewJobDefinition(ctx, "test", &batch.JobDefinitionArgs{
// 			Type: pulumi.String("container"),
// 			PlatformCapabilities: pulumi.StringArray{
// 				pulumi.String("FARGATE"),
// 			},
// 			ContainerProperties: ecsTaskExecutionRole.Arn.ApplyT(func(arn string) (string, error) {
// 				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"command\": [\"echo\", \"test\"],\n", "  \"image\": \"busybox\",\n", "  \"fargatePlatformConfiguration\": {\n", "    \"platformVersion\": \"LATEST\"\n", "  },\n", "  \"resourceRequirements\": [\n", "    {\"type\": \"VCPU\", \"value\": \"0.25\"},\n", "    {\"type\": \"MEMORY\", \"value\": \"512\"}\n", "  ],\n", "  \"executionRoleArn\": \"", arn, "\"\n", "}\n"), nil
// 			}).(pulumi.StringOutput),
// 		})
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
// Batch Job Definition can be imported using the `arn`, e.g.
//
// ```sh
//  $ pulumi import aws:batch/jobDefinition:JobDefinition test arn:aws:batch:us-east-1:123456789012:job-definition/sample
// ```
type JobDefinition struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name of the job definition.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A valid [container properties](http://docs.aws.amazon.com/batch/latest/APIReference/API_RegisterJobDefinition.html)
	// provided as a single valid JSON document. This parameter is required if the `type` parameter is `container`.
	ContainerProperties pulumi.StringPtrOutput `pulumi:"containerProperties"`
	// Specifies the name of the job definition.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the parameter substitution placeholders to set in the job definition.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The platform capabilities required by the job definition. If no value is specified, it defaults to `EC2`. To run the job on Fargate resources, specify `FARGATE`.
	PlatformCapabilities pulumi.StringArrayOutput `pulumi:"platformCapabilities"`
	// Specifies whether to propagate the tags from the job definition to the corresponding Amazon ECS task. Default is `false`.
	PropagateTags pulumi.BoolPtrOutput `pulumi:"propagateTags"`
	// Specifies the retry strategy to use for failed jobs that are submitted with this job definition.
	// Maximum number of `retryStrategy` is `1`.  Defined below.
	RetryStrategy JobDefinitionRetryStrategyPtrOutput `pulumi:"retryStrategy"`
	// The revision of the job definition.
	Revision pulumi.IntOutput `pulumi:"revision"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of `timeout` is `1`. Defined below.
	Timeout JobDefinitionTimeoutPtrOutput `pulumi:"timeout"`
	// The type of job definition.  Must be `container`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewJobDefinition registers a new resource with the given unique name, arguments, and options.
func NewJobDefinition(ctx *pulumi.Context,
	name string, args *JobDefinitionArgs, opts ...pulumi.ResourceOption) (*JobDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource JobDefinition
	err := ctx.RegisterResource("aws:batch/jobDefinition:JobDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobDefinition gets an existing JobDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobDefinitionState, opts ...pulumi.ResourceOption) (*JobDefinition, error) {
	var resource JobDefinition
	err := ctx.ReadResource("aws:batch/jobDefinition:JobDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobDefinition resources.
type jobDefinitionState struct {
	// The Amazon Resource Name of the job definition.
	Arn *string `pulumi:"arn"`
	// A valid [container properties](http://docs.aws.amazon.com/batch/latest/APIReference/API_RegisterJobDefinition.html)
	// provided as a single valid JSON document. This parameter is required if the `type` parameter is `container`.
	ContainerProperties *string `pulumi:"containerProperties"`
	// Specifies the name of the job definition.
	Name *string `pulumi:"name"`
	// Specifies the parameter substitution placeholders to set in the job definition.
	Parameters map[string]string `pulumi:"parameters"`
	// The platform capabilities required by the job definition. If no value is specified, it defaults to `EC2`. To run the job on Fargate resources, specify `FARGATE`.
	PlatformCapabilities []string `pulumi:"platformCapabilities"`
	// Specifies whether to propagate the tags from the job definition to the corresponding Amazon ECS task. Default is `false`.
	PropagateTags *bool `pulumi:"propagateTags"`
	// Specifies the retry strategy to use for failed jobs that are submitted with this job definition.
	// Maximum number of `retryStrategy` is `1`.  Defined below.
	RetryStrategy *JobDefinitionRetryStrategy `pulumi:"retryStrategy"`
	// The revision of the job definition.
	Revision *int `pulumi:"revision"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of `timeout` is `1`. Defined below.
	Timeout *JobDefinitionTimeout `pulumi:"timeout"`
	// The type of job definition.  Must be `container`.
	Type *string `pulumi:"type"`
}

type JobDefinitionState struct {
	// The Amazon Resource Name of the job definition.
	Arn pulumi.StringPtrInput
	// A valid [container properties](http://docs.aws.amazon.com/batch/latest/APIReference/API_RegisterJobDefinition.html)
	// provided as a single valid JSON document. This parameter is required if the `type` parameter is `container`.
	ContainerProperties pulumi.StringPtrInput
	// Specifies the name of the job definition.
	Name pulumi.StringPtrInput
	// Specifies the parameter substitution placeholders to set in the job definition.
	Parameters pulumi.StringMapInput
	// The platform capabilities required by the job definition. If no value is specified, it defaults to `EC2`. To run the job on Fargate resources, specify `FARGATE`.
	PlatformCapabilities pulumi.StringArrayInput
	// Specifies whether to propagate the tags from the job definition to the corresponding Amazon ECS task. Default is `false`.
	PropagateTags pulumi.BoolPtrInput
	// Specifies the retry strategy to use for failed jobs that are submitted with this job definition.
	// Maximum number of `retryStrategy` is `1`.  Defined below.
	RetryStrategy JobDefinitionRetryStrategyPtrInput
	// The revision of the job definition.
	Revision pulumi.IntPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of `timeout` is `1`. Defined below.
	Timeout JobDefinitionTimeoutPtrInput
	// The type of job definition.  Must be `container`.
	Type pulumi.StringPtrInput
}

func (JobDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobDefinitionState)(nil)).Elem()
}

type jobDefinitionArgs struct {
	// A valid [container properties](http://docs.aws.amazon.com/batch/latest/APIReference/API_RegisterJobDefinition.html)
	// provided as a single valid JSON document. This parameter is required if the `type` parameter is `container`.
	ContainerProperties *string `pulumi:"containerProperties"`
	// Specifies the name of the job definition.
	Name *string `pulumi:"name"`
	// Specifies the parameter substitution placeholders to set in the job definition.
	Parameters map[string]string `pulumi:"parameters"`
	// The platform capabilities required by the job definition. If no value is specified, it defaults to `EC2`. To run the job on Fargate resources, specify `FARGATE`.
	PlatformCapabilities []string `pulumi:"platformCapabilities"`
	// Specifies whether to propagate the tags from the job definition to the corresponding Amazon ECS task. Default is `false`.
	PropagateTags *bool `pulumi:"propagateTags"`
	// Specifies the retry strategy to use for failed jobs that are submitted with this job definition.
	// Maximum number of `retryStrategy` is `1`.  Defined below.
	RetryStrategy *JobDefinitionRetryStrategy `pulumi:"retryStrategy"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of `timeout` is `1`. Defined below.
	Timeout *JobDefinitionTimeout `pulumi:"timeout"`
	// The type of job definition.  Must be `container`.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a JobDefinition resource.
type JobDefinitionArgs struct {
	// A valid [container properties](http://docs.aws.amazon.com/batch/latest/APIReference/API_RegisterJobDefinition.html)
	// provided as a single valid JSON document. This parameter is required if the `type` parameter is `container`.
	ContainerProperties pulumi.StringPtrInput
	// Specifies the name of the job definition.
	Name pulumi.StringPtrInput
	// Specifies the parameter substitution placeholders to set in the job definition.
	Parameters pulumi.StringMapInput
	// The platform capabilities required by the job definition. If no value is specified, it defaults to `EC2`. To run the job on Fargate resources, specify `FARGATE`.
	PlatformCapabilities pulumi.StringArrayInput
	// Specifies whether to propagate the tags from the job definition to the corresponding Amazon ECS task. Default is `false`.
	PropagateTags pulumi.BoolPtrInput
	// Specifies the retry strategy to use for failed jobs that are submitted with this job definition.
	// Maximum number of `retryStrategy` is `1`.  Defined below.
	RetryStrategy JobDefinitionRetryStrategyPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of `timeout` is `1`. Defined below.
	Timeout JobDefinitionTimeoutPtrInput
	// The type of job definition.  Must be `container`.
	Type pulumi.StringInput
}

func (JobDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobDefinitionArgs)(nil)).Elem()
}

type JobDefinitionInput interface {
	pulumi.Input

	ToJobDefinitionOutput() JobDefinitionOutput
	ToJobDefinitionOutputWithContext(ctx context.Context) JobDefinitionOutput
}

func (*JobDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((*JobDefinition)(nil))
}

func (i *JobDefinition) ToJobDefinitionOutput() JobDefinitionOutput {
	return i.ToJobDefinitionOutputWithContext(context.Background())
}

func (i *JobDefinition) ToJobDefinitionOutputWithContext(ctx context.Context) JobDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDefinitionOutput)
}

func (i *JobDefinition) ToJobDefinitionPtrOutput() JobDefinitionPtrOutput {
	return i.ToJobDefinitionPtrOutputWithContext(context.Background())
}

func (i *JobDefinition) ToJobDefinitionPtrOutputWithContext(ctx context.Context) JobDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDefinitionPtrOutput)
}

type JobDefinitionPtrInput interface {
	pulumi.Input

	ToJobDefinitionPtrOutput() JobDefinitionPtrOutput
	ToJobDefinitionPtrOutputWithContext(ctx context.Context) JobDefinitionPtrOutput
}

type jobDefinitionPtrType JobDefinitionArgs

func (*jobDefinitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDefinition)(nil))
}

func (i *jobDefinitionPtrType) ToJobDefinitionPtrOutput() JobDefinitionPtrOutput {
	return i.ToJobDefinitionPtrOutputWithContext(context.Background())
}

func (i *jobDefinitionPtrType) ToJobDefinitionPtrOutputWithContext(ctx context.Context) JobDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDefinitionPtrOutput)
}

// JobDefinitionArrayInput is an input type that accepts JobDefinitionArray and JobDefinitionArrayOutput values.
// You can construct a concrete instance of `JobDefinitionArrayInput` via:
//
//          JobDefinitionArray{ JobDefinitionArgs{...} }
type JobDefinitionArrayInput interface {
	pulumi.Input

	ToJobDefinitionArrayOutput() JobDefinitionArrayOutput
	ToJobDefinitionArrayOutputWithContext(context.Context) JobDefinitionArrayOutput
}

type JobDefinitionArray []JobDefinitionInput

func (JobDefinitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*JobDefinition)(nil)).Elem()
}

func (i JobDefinitionArray) ToJobDefinitionArrayOutput() JobDefinitionArrayOutput {
	return i.ToJobDefinitionArrayOutputWithContext(context.Background())
}

func (i JobDefinitionArray) ToJobDefinitionArrayOutputWithContext(ctx context.Context) JobDefinitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDefinitionArrayOutput)
}

// JobDefinitionMapInput is an input type that accepts JobDefinitionMap and JobDefinitionMapOutput values.
// You can construct a concrete instance of `JobDefinitionMapInput` via:
//
//          JobDefinitionMap{ "key": JobDefinitionArgs{...} }
type JobDefinitionMapInput interface {
	pulumi.Input

	ToJobDefinitionMapOutput() JobDefinitionMapOutput
	ToJobDefinitionMapOutputWithContext(context.Context) JobDefinitionMapOutput
}

type JobDefinitionMap map[string]JobDefinitionInput

func (JobDefinitionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*JobDefinition)(nil)).Elem()
}

func (i JobDefinitionMap) ToJobDefinitionMapOutput() JobDefinitionMapOutput {
	return i.ToJobDefinitionMapOutputWithContext(context.Background())
}

func (i JobDefinitionMap) ToJobDefinitionMapOutputWithContext(ctx context.Context) JobDefinitionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDefinitionMapOutput)
}

type JobDefinitionOutput struct{ *pulumi.OutputState }

func (JobDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobDefinition)(nil))
}

func (o JobDefinitionOutput) ToJobDefinitionOutput() JobDefinitionOutput {
	return o
}

func (o JobDefinitionOutput) ToJobDefinitionOutputWithContext(ctx context.Context) JobDefinitionOutput {
	return o
}

func (o JobDefinitionOutput) ToJobDefinitionPtrOutput() JobDefinitionPtrOutput {
	return o.ToJobDefinitionPtrOutputWithContext(context.Background())
}

func (o JobDefinitionOutput) ToJobDefinitionPtrOutputWithContext(ctx context.Context) JobDefinitionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobDefinition) *JobDefinition {
		return &v
	}).(JobDefinitionPtrOutput)
}

type JobDefinitionPtrOutput struct{ *pulumi.OutputState }

func (JobDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDefinition)(nil))
}

func (o JobDefinitionPtrOutput) ToJobDefinitionPtrOutput() JobDefinitionPtrOutput {
	return o
}

func (o JobDefinitionPtrOutput) ToJobDefinitionPtrOutputWithContext(ctx context.Context) JobDefinitionPtrOutput {
	return o
}

func (o JobDefinitionPtrOutput) Elem() JobDefinitionOutput {
	return o.ApplyT(func(v *JobDefinition) JobDefinition {
		if v != nil {
			return *v
		}
		var ret JobDefinition
		return ret
	}).(JobDefinitionOutput)
}

type JobDefinitionArrayOutput struct{ *pulumi.OutputState }

func (JobDefinitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobDefinition)(nil))
}

func (o JobDefinitionArrayOutput) ToJobDefinitionArrayOutput() JobDefinitionArrayOutput {
	return o
}

func (o JobDefinitionArrayOutput) ToJobDefinitionArrayOutputWithContext(ctx context.Context) JobDefinitionArrayOutput {
	return o
}

func (o JobDefinitionArrayOutput) Index(i pulumi.IntInput) JobDefinitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JobDefinition {
		return vs[0].([]JobDefinition)[vs[1].(int)]
	}).(JobDefinitionOutput)
}

type JobDefinitionMapOutput struct{ *pulumi.OutputState }

func (JobDefinitionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]JobDefinition)(nil))
}

func (o JobDefinitionMapOutput) ToJobDefinitionMapOutput() JobDefinitionMapOutput {
	return o
}

func (o JobDefinitionMapOutput) ToJobDefinitionMapOutputWithContext(ctx context.Context) JobDefinitionMapOutput {
	return o
}

func (o JobDefinitionMapOutput) MapIndex(k pulumi.StringInput) JobDefinitionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) JobDefinition {
		return vs[0].(map[string]JobDefinition)[vs[1].(string)]
	}).(JobDefinitionOutput)
}

func init() {
	pulumi.RegisterOutputType(JobDefinitionOutput{})
	pulumi.RegisterOutputType(JobDefinitionPtrOutput{})
	pulumi.RegisterOutputType(JobDefinitionArrayOutput{})
	pulumi.RegisterOutputType(JobDefinitionMapOutput{})
}
