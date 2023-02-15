// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package evidently

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CloudWatch Evidently Launch resource.
//
// ## Example Usage
// ### Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/evidently"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := evidently.NewLaunch(ctx, "example", &evidently.LaunchArgs{
//				Project: pulumi.Any(aws_evidently_project.Example.Name),
//				Groups: evidently.LaunchGroupArray{
//					&evidently.LaunchGroupArgs{
//						Feature:   pulumi.Any(aws_evidently_feature.Example.Name),
//						Name:      pulumi.String("Variation1"),
//						Variation: pulumi.String("Variation1"),
//					},
//				},
//				ScheduledSplitsConfig: &evidently.LaunchScheduledSplitsConfigArgs{
//					Steps: evidently.LaunchScheduledSplitsConfigStepArray{
//						&evidently.LaunchScheduledSplitsConfigStepArgs{
//							GroupWeights: pulumi.IntMap{
//								"Variation1": pulumi.Int(0),
//							},
//							StartTime: pulumi.String("2024-01-07 01:43:59+00:00"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### With description
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/evidently"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := evidently.NewLaunch(ctx, "example", &evidently.LaunchArgs{
//				Project:     pulumi.Any(aws_evidently_project.Example.Name),
//				Description: pulumi.String("example description"),
//				Groups: evidently.LaunchGroupArray{
//					&evidently.LaunchGroupArgs{
//						Feature:   pulumi.Any(aws_evidently_feature.Example.Name),
//						Name:      pulumi.String("Variation1"),
//						Variation: pulumi.String("Variation1"),
//					},
//				},
//				ScheduledSplitsConfig: &evidently.LaunchScheduledSplitsConfigArgs{
//					Steps: evidently.LaunchScheduledSplitsConfigStepArray{
//						&evidently.LaunchScheduledSplitsConfigStepArgs{
//							GroupWeights: pulumi.IntMap{
//								"Variation1": pulumi.Int(0),
//							},
//							StartTime: pulumi.String("2024-01-07 01:43:59+00:00"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### With multiple groups
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/evidently"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := evidently.NewLaunch(ctx, "example", &evidently.LaunchArgs{
//				Project: pulumi.Any(aws_evidently_project.Example.Name),
//				Groups: evidently.LaunchGroupArray{
//					&evidently.LaunchGroupArgs{
//						Feature:     pulumi.Any(aws_evidently_feature.Example.Name),
//						Name:        pulumi.String("Variation1"),
//						Variation:   pulumi.String("Variation1"),
//						Description: pulumi.String("first-group"),
//					},
//					&evidently.LaunchGroupArgs{
//						Feature:     pulumi.Any(aws_evidently_feature.Example.Name),
//						Name:        pulumi.String("Variation2"),
//						Variation:   pulumi.String("Variation2"),
//						Description: pulumi.String("second-group"),
//					},
//				},
//				ScheduledSplitsConfig: &evidently.LaunchScheduledSplitsConfigArgs{
//					Steps: evidently.LaunchScheduledSplitsConfigStepArray{
//						&evidently.LaunchScheduledSplitsConfigStepArgs{
//							GroupWeights: pulumi.IntMap{
//								"Variation1": pulumi.Int(0),
//								"Variation2": pulumi.Int(0),
//							},
//							StartTime: pulumi.String("2024-01-07 01:43:59+00:00"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### With metricMonitors
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/evidently"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := evidently.NewLaunch(ctx, "example", &evidently.LaunchArgs{
//				Project: pulumi.Any(aws_evidently_project.Example.Name),
//				Groups: evidently.LaunchGroupArray{
//					&evidently.LaunchGroupArgs{
//						Feature:   pulumi.Any(aws_evidently_feature.Example.Name),
//						Name:      pulumi.String("Variation1"),
//						Variation: pulumi.String("Variation1"),
//					},
//				},
//				MetricMonitors: evidently.LaunchMetricMonitorArray{
//					&evidently.LaunchMetricMonitorArgs{
//						MetricDefinition: &evidently.LaunchMetricMonitorMetricDefinitionArgs{
//							EntityIdKey:  pulumi.String("entity_id_key1"),
//							EventPattern: pulumi.String("{\"Price\":[{\"numeric\":[\">\",11,\"<=\",22]}]}"),
//							Name:         pulumi.String("name1"),
//							UnitLabel:    pulumi.String("unit_label1"),
//							ValueKey:     pulumi.String("value_key1"),
//						},
//					},
//					&evidently.LaunchMetricMonitorArgs{
//						MetricDefinition: &evidently.LaunchMetricMonitorMetricDefinitionArgs{
//							EntityIdKey:  pulumi.String("entity_id_key2"),
//							EventPattern: pulumi.String("{\"Price\":[{\"numeric\":[\">\",9,\"<=\",19]}]}"),
//							Name:         pulumi.String("name2"),
//							UnitLabel:    pulumi.String("unit_label2"),
//							ValueKey:     pulumi.String("value_key2"),
//						},
//					},
//				},
//				ScheduledSplitsConfig: &evidently.LaunchScheduledSplitsConfigArgs{
//					Steps: evidently.LaunchScheduledSplitsConfigStepArray{
//						&evidently.LaunchScheduledSplitsConfigStepArgs{
//							GroupWeights: pulumi.IntMap{
//								"Variation1": pulumi.Int(0),
//							},
//							StartTime: pulumi.String("2024-01-07 01:43:59+00:00"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### With randomizationSalt
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/evidently"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := evidently.NewLaunch(ctx, "example", &evidently.LaunchArgs{
//				Project:           pulumi.Any(aws_evidently_project.Example.Name),
//				RandomizationSalt: pulumi.String("example randomization salt"),
//				Groups: evidently.LaunchGroupArray{
//					&evidently.LaunchGroupArgs{
//						Feature:   pulumi.Any(aws_evidently_feature.Example.Name),
//						Name:      pulumi.String("Variation1"),
//						Variation: pulumi.String("Variation1"),
//					},
//				},
//				ScheduledSplitsConfig: &evidently.LaunchScheduledSplitsConfigArgs{
//					Steps: evidently.LaunchScheduledSplitsConfigStepArray{
//						&evidently.LaunchScheduledSplitsConfigStepArgs{
//							GroupWeights: pulumi.IntMap{
//								"Variation1": pulumi.Int(0),
//							},
//							StartTime: pulumi.String("2024-01-07 01:43:59+00:00"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### With multiple steps
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/evidently"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := evidently.NewLaunch(ctx, "example", &evidently.LaunchArgs{
//				Project: pulumi.Any(aws_evidently_project.Example.Name),
//				Groups: evidently.LaunchGroupArray{
//					&evidently.LaunchGroupArgs{
//						Feature:   pulumi.Any(aws_evidently_feature.Example.Name),
//						Name:      pulumi.String("Variation1"),
//						Variation: pulumi.String("Variation1"),
//					},
//					&evidently.LaunchGroupArgs{
//						Feature:   pulumi.Any(aws_evidently_feature.Example.Name),
//						Name:      pulumi.String("Variation2"),
//						Variation: pulumi.String("Variation2"),
//					},
//				},
//				ScheduledSplitsConfig: &evidently.LaunchScheduledSplitsConfigArgs{
//					Steps: evidently.LaunchScheduledSplitsConfigStepArray{
//						&evidently.LaunchScheduledSplitsConfigStepArgs{
//							GroupWeights: pulumi.IntMap{
//								"Variation1": pulumi.Int(15),
//								"Variation2": pulumi.Int(10),
//							},
//							StartTime: pulumi.String("2024-01-07 01:43:59+00:00"),
//						},
//						&evidently.LaunchScheduledSplitsConfigStepArgs{
//							GroupWeights: pulumi.IntMap{
//								"Variation1": pulumi.Int(20),
//								"Variation2": pulumi.Int(25),
//							},
//							StartTime: pulumi.String("2024-01-08 01:43:59+00:00"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### With segment overrides
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/evidently"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := evidently.NewLaunch(ctx, "example", &evidently.LaunchArgs{
//				Project: pulumi.Any(aws_evidently_project.Example.Name),
//				Groups: evidently.LaunchGroupArray{
//					&evidently.LaunchGroupArgs{
//						Feature:   pulumi.Any(aws_evidently_feature.Example.Name),
//						Name:      pulumi.String("Variation1"),
//						Variation: pulumi.String("Variation1"),
//					},
//					&evidently.LaunchGroupArgs{
//						Feature:   pulumi.Any(aws_evidently_feature.Example.Name),
//						Name:      pulumi.String("Variation2"),
//						Variation: pulumi.String("Variation2"),
//					},
//				},
//				ScheduledSplitsConfig: &evidently.LaunchScheduledSplitsConfigArgs{
//					Steps: evidently.LaunchScheduledSplitsConfigStepArray{
//						&evidently.LaunchScheduledSplitsConfigStepArgs{
//							GroupWeights: pulumi.IntMap{
//								"Variation1": pulumi.Int(0),
//								"Variation2": pulumi.Int(0),
//							},
//							SegmentOverrides: evidently.LaunchScheduledSplitsConfigStepSegmentOverrideArray{
//								&evidently.LaunchScheduledSplitsConfigStepSegmentOverrideArgs{
//									EvaluationOrder: pulumi.Int(1),
//									Segment:         pulumi.Any(aws_evidently_segment.Example.Name),
//									Weights: pulumi.IntMap{
//										"Variation2": pulumi.Int(10000),
//									},
//								},
//								&evidently.LaunchScheduledSplitsConfigStepSegmentOverrideArgs{
//									EvaluationOrder: pulumi.Int(2),
//									Segment:         pulumi.Any(aws_evidently_segment.Example.Name),
//									Weights: pulumi.IntMap{
//										"Variation1": pulumi.Int(40000),
//										"Variation2": pulumi.Int(30000),
//									},
//								},
//							},
//							StartTime: pulumi.String("2024-01-08 01:43:59+00:00"),
//						},
//					},
//				},
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
// CloudWatch Evidently Launch can be imported using the `name` of the launch and `name` or `arn` of the hosting CloudWatch Evidently Project separated by a `:`, e.g. with the `name` of the launch and `arn` of the project,
//
// ```sh
//
//	$ pulumi import aws:evidently/launch:Launch example exampleLaunchName:arn:aws:evidently:us-east-1:123456789012:project/exampleProjectName
//
// ```
//
//	e.g. with the `name` of the launch and `name` of the project,
//
// ```sh
//
//	$ pulumi import aws:evidently/launch:Launch example exampleLaunchName:exampleProjectName
//
// ```
type Launch struct {
	pulumi.CustomResourceState

	// The ARN of the launch.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The date and time that the launch is created.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// Specifies the description of the launch.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A block that contains information about the start and end times of the launch. Detailed below
	Executions LaunchExecutionArrayOutput `pulumi:"executions"`
	// One or up to five blocks that contain the feature and variations that are to be used for the launch. Detailed below.
	Groups LaunchGroupArrayOutput `pulumi:"groups"`
	// The date and time that the launch was most recently updated.
	LastUpdatedTime pulumi.StringOutput `pulumi:"lastUpdatedTime"`
	// One or up to three blocks that define the metrics that will be used to monitor the launch performance. Detailed below.
	MetricMonitors LaunchMetricMonitorArrayOutput `pulumi:"metricMonitors"`
	// The name for the new launch. Minimum length of `1`. Maximum length of `127`.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name or ARN of the project that is to contain the new launch.
	Project pulumi.StringOutput `pulumi:"project"`
	// When Evidently assigns a particular user session to a launch, it must use a randomization ID to determine which variation the user session is served. This randomization ID is a combination of the entity ID and randomizationSalt. If you omit randomizationSalt, Evidently uses the launch name as the randomizationSalt.
	RandomizationSalt pulumi.StringPtrOutput `pulumi:"randomizationSalt"`
	// A block that defines the traffic allocation percentages among the feature variations during each step of the launch. Detailed below.
	ScheduledSplitsConfig LaunchScheduledSplitsConfigPtrOutput `pulumi:"scheduledSplitsConfig"`
	// The current state of the launch. Valid values are `CREATED`, `UPDATING`, `RUNNING`, `COMPLETED`, and `CANCELLED`.
	Status pulumi.StringOutput `pulumi:"status"`
	// If the launch was stopped, this is the string that was entered by the person who stopped the launch, to explain why it was stopped.
	StatusReason pulumi.StringOutput `pulumi:"statusReason"`
	// Tags to apply to the launch. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The type of launch.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLaunch registers a new resource with the given unique name, arguments, and options.
func NewLaunch(ctx *pulumi.Context,
	name string, args *LaunchArgs, opts ...pulumi.ResourceOption) (*Launch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Groups == nil {
		return nil, errors.New("invalid value for required argument 'Groups'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Launch
	err := ctx.RegisterResource("aws:evidently/launch:Launch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLaunch gets an existing Launch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLaunch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LaunchState, opts ...pulumi.ResourceOption) (*Launch, error) {
	var resource Launch
	err := ctx.ReadResource("aws:evidently/launch:Launch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Launch resources.
type launchState struct {
	// The ARN of the launch.
	Arn *string `pulumi:"arn"`
	// The date and time that the launch is created.
	CreatedTime *string `pulumi:"createdTime"`
	// Specifies the description of the launch.
	Description *string `pulumi:"description"`
	// A block that contains information about the start and end times of the launch. Detailed below
	Executions []LaunchExecution `pulumi:"executions"`
	// One or up to five blocks that contain the feature and variations that are to be used for the launch. Detailed below.
	Groups []LaunchGroup `pulumi:"groups"`
	// The date and time that the launch was most recently updated.
	LastUpdatedTime *string `pulumi:"lastUpdatedTime"`
	// One or up to three blocks that define the metrics that will be used to monitor the launch performance. Detailed below.
	MetricMonitors []LaunchMetricMonitor `pulumi:"metricMonitors"`
	// The name for the new launch. Minimum length of `1`. Maximum length of `127`.
	Name *string `pulumi:"name"`
	// The name or ARN of the project that is to contain the new launch.
	Project *string `pulumi:"project"`
	// When Evidently assigns a particular user session to a launch, it must use a randomization ID to determine which variation the user session is served. This randomization ID is a combination of the entity ID and randomizationSalt. If you omit randomizationSalt, Evidently uses the launch name as the randomizationSalt.
	RandomizationSalt *string `pulumi:"randomizationSalt"`
	// A block that defines the traffic allocation percentages among the feature variations during each step of the launch. Detailed below.
	ScheduledSplitsConfig *LaunchScheduledSplitsConfig `pulumi:"scheduledSplitsConfig"`
	// The current state of the launch. Valid values are `CREATED`, `UPDATING`, `RUNNING`, `COMPLETED`, and `CANCELLED`.
	Status *string `pulumi:"status"`
	// If the launch was stopped, this is the string that was entered by the person who stopped the launch, to explain why it was stopped.
	StatusReason *string `pulumi:"statusReason"`
	// Tags to apply to the launch. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The type of launch.
	Type *string `pulumi:"type"`
}

type LaunchState struct {
	// The ARN of the launch.
	Arn pulumi.StringPtrInput
	// The date and time that the launch is created.
	CreatedTime pulumi.StringPtrInput
	// Specifies the description of the launch.
	Description pulumi.StringPtrInput
	// A block that contains information about the start and end times of the launch. Detailed below
	Executions LaunchExecutionArrayInput
	// One or up to five blocks that contain the feature and variations that are to be used for the launch. Detailed below.
	Groups LaunchGroupArrayInput
	// The date and time that the launch was most recently updated.
	LastUpdatedTime pulumi.StringPtrInput
	// One or up to three blocks that define the metrics that will be used to monitor the launch performance. Detailed below.
	MetricMonitors LaunchMetricMonitorArrayInput
	// The name for the new launch. Minimum length of `1`. Maximum length of `127`.
	Name pulumi.StringPtrInput
	// The name or ARN of the project that is to contain the new launch.
	Project pulumi.StringPtrInput
	// When Evidently assigns a particular user session to a launch, it must use a randomization ID to determine which variation the user session is served. This randomization ID is a combination of the entity ID and randomizationSalt. If you omit randomizationSalt, Evidently uses the launch name as the randomizationSalt.
	RandomizationSalt pulumi.StringPtrInput
	// A block that defines the traffic allocation percentages among the feature variations during each step of the launch. Detailed below.
	ScheduledSplitsConfig LaunchScheduledSplitsConfigPtrInput
	// The current state of the launch. Valid values are `CREATED`, `UPDATING`, `RUNNING`, `COMPLETED`, and `CANCELLED`.
	Status pulumi.StringPtrInput
	// If the launch was stopped, this is the string that was entered by the person who stopped the launch, to explain why it was stopped.
	StatusReason pulumi.StringPtrInput
	// Tags to apply to the launch. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// The type of launch.
	Type pulumi.StringPtrInput
}

func (LaunchState) ElementType() reflect.Type {
	return reflect.TypeOf((*launchState)(nil)).Elem()
}

type launchArgs struct {
	// Specifies the description of the launch.
	Description *string `pulumi:"description"`
	// One or up to five blocks that contain the feature and variations that are to be used for the launch. Detailed below.
	Groups []LaunchGroup `pulumi:"groups"`
	// One or up to three blocks that define the metrics that will be used to monitor the launch performance. Detailed below.
	MetricMonitors []LaunchMetricMonitor `pulumi:"metricMonitors"`
	// The name for the new launch. Minimum length of `1`. Maximum length of `127`.
	Name *string `pulumi:"name"`
	// The name or ARN of the project that is to contain the new launch.
	Project string `pulumi:"project"`
	// When Evidently assigns a particular user session to a launch, it must use a randomization ID to determine which variation the user session is served. This randomization ID is a combination of the entity ID and randomizationSalt. If you omit randomizationSalt, Evidently uses the launch name as the randomizationSalt.
	RandomizationSalt *string `pulumi:"randomizationSalt"`
	// A block that defines the traffic allocation percentages among the feature variations during each step of the launch. Detailed below.
	ScheduledSplitsConfig *LaunchScheduledSplitsConfig `pulumi:"scheduledSplitsConfig"`
	// Tags to apply to the launch. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Launch resource.
type LaunchArgs struct {
	// Specifies the description of the launch.
	Description pulumi.StringPtrInput
	// One or up to five blocks that contain the feature and variations that are to be used for the launch. Detailed below.
	Groups LaunchGroupArrayInput
	// One or up to three blocks that define the metrics that will be used to monitor the launch performance. Detailed below.
	MetricMonitors LaunchMetricMonitorArrayInput
	// The name for the new launch. Minimum length of `1`. Maximum length of `127`.
	Name pulumi.StringPtrInput
	// The name or ARN of the project that is to contain the new launch.
	Project pulumi.StringInput
	// When Evidently assigns a particular user session to a launch, it must use a randomization ID to determine which variation the user session is served. This randomization ID is a combination of the entity ID and randomizationSalt. If you omit randomizationSalt, Evidently uses the launch name as the randomizationSalt.
	RandomizationSalt pulumi.StringPtrInput
	// A block that defines the traffic allocation percentages among the feature variations during each step of the launch. Detailed below.
	ScheduledSplitsConfig LaunchScheduledSplitsConfigPtrInput
	// Tags to apply to the launch. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (LaunchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*launchArgs)(nil)).Elem()
}

type LaunchInput interface {
	pulumi.Input

	ToLaunchOutput() LaunchOutput
	ToLaunchOutputWithContext(ctx context.Context) LaunchOutput
}

func (*Launch) ElementType() reflect.Type {
	return reflect.TypeOf((**Launch)(nil)).Elem()
}

func (i *Launch) ToLaunchOutput() LaunchOutput {
	return i.ToLaunchOutputWithContext(context.Background())
}

func (i *Launch) ToLaunchOutputWithContext(ctx context.Context) LaunchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchOutput)
}

// LaunchArrayInput is an input type that accepts LaunchArray and LaunchArrayOutput values.
// You can construct a concrete instance of `LaunchArrayInput` via:
//
//	LaunchArray{ LaunchArgs{...} }
type LaunchArrayInput interface {
	pulumi.Input

	ToLaunchArrayOutput() LaunchArrayOutput
	ToLaunchArrayOutputWithContext(context.Context) LaunchArrayOutput
}

type LaunchArray []LaunchInput

func (LaunchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Launch)(nil)).Elem()
}

func (i LaunchArray) ToLaunchArrayOutput() LaunchArrayOutput {
	return i.ToLaunchArrayOutputWithContext(context.Background())
}

func (i LaunchArray) ToLaunchArrayOutputWithContext(ctx context.Context) LaunchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchArrayOutput)
}

// LaunchMapInput is an input type that accepts LaunchMap and LaunchMapOutput values.
// You can construct a concrete instance of `LaunchMapInput` via:
//
//	LaunchMap{ "key": LaunchArgs{...} }
type LaunchMapInput interface {
	pulumi.Input

	ToLaunchMapOutput() LaunchMapOutput
	ToLaunchMapOutputWithContext(context.Context) LaunchMapOutput
}

type LaunchMap map[string]LaunchInput

func (LaunchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Launch)(nil)).Elem()
}

func (i LaunchMap) ToLaunchMapOutput() LaunchMapOutput {
	return i.ToLaunchMapOutputWithContext(context.Background())
}

func (i LaunchMap) ToLaunchMapOutputWithContext(ctx context.Context) LaunchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchMapOutput)
}

type LaunchOutput struct{ *pulumi.OutputState }

func (LaunchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Launch)(nil)).Elem()
}

func (o LaunchOutput) ToLaunchOutput() LaunchOutput {
	return o
}

func (o LaunchOutput) ToLaunchOutputWithContext(ctx context.Context) LaunchOutput {
	return o
}

// The ARN of the launch.
func (o LaunchOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Launch) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The date and time that the launch is created.
func (o LaunchOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Launch) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// Specifies the description of the launch.
func (o LaunchOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Launch) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A block that contains information about the start and end times of the launch. Detailed below
func (o LaunchOutput) Executions() LaunchExecutionArrayOutput {
	return o.ApplyT(func(v *Launch) LaunchExecutionArrayOutput { return v.Executions }).(LaunchExecutionArrayOutput)
}

// One or up to five blocks that contain the feature and variations that are to be used for the launch. Detailed below.
func (o LaunchOutput) Groups() LaunchGroupArrayOutput {
	return o.ApplyT(func(v *Launch) LaunchGroupArrayOutput { return v.Groups }).(LaunchGroupArrayOutput)
}

// The date and time that the launch was most recently updated.
func (o LaunchOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Launch) pulumi.StringOutput { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

// One or up to three blocks that define the metrics that will be used to monitor the launch performance. Detailed below.
func (o LaunchOutput) MetricMonitors() LaunchMetricMonitorArrayOutput {
	return o.ApplyT(func(v *Launch) LaunchMetricMonitorArrayOutput { return v.MetricMonitors }).(LaunchMetricMonitorArrayOutput)
}

// The name for the new launch. Minimum length of `1`. Maximum length of `127`.
func (o LaunchOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Launch) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name or ARN of the project that is to contain the new launch.
func (o LaunchOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Launch) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// When Evidently assigns a particular user session to a launch, it must use a randomization ID to determine which variation the user session is served. This randomization ID is a combination of the entity ID and randomizationSalt. If you omit randomizationSalt, Evidently uses the launch name as the randomizationSalt.
func (o LaunchOutput) RandomizationSalt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Launch) pulumi.StringPtrOutput { return v.RandomizationSalt }).(pulumi.StringPtrOutput)
}

// A block that defines the traffic allocation percentages among the feature variations during each step of the launch. Detailed below.
func (o LaunchOutput) ScheduledSplitsConfig() LaunchScheduledSplitsConfigPtrOutput {
	return o.ApplyT(func(v *Launch) LaunchScheduledSplitsConfigPtrOutput { return v.ScheduledSplitsConfig }).(LaunchScheduledSplitsConfigPtrOutput)
}

// The current state of the launch. Valid values are `CREATED`, `UPDATING`, `RUNNING`, `COMPLETED`, and `CANCELLED`.
func (o LaunchOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Launch) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// If the launch was stopped, this is the string that was entered by the person who stopped the launch, to explain why it was stopped.
func (o LaunchOutput) StatusReason() pulumi.StringOutput {
	return o.ApplyT(func(v *Launch) pulumi.StringOutput { return v.StatusReason }).(pulumi.StringOutput)
}

// Tags to apply to the launch. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o LaunchOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Launch) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o LaunchOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Launch) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The type of launch.
func (o LaunchOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Launch) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type LaunchArrayOutput struct{ *pulumi.OutputState }

func (LaunchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Launch)(nil)).Elem()
}

func (o LaunchArrayOutput) ToLaunchArrayOutput() LaunchArrayOutput {
	return o
}

func (o LaunchArrayOutput) ToLaunchArrayOutputWithContext(ctx context.Context) LaunchArrayOutput {
	return o
}

func (o LaunchArrayOutput) Index(i pulumi.IntInput) LaunchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Launch {
		return vs[0].([]*Launch)[vs[1].(int)]
	}).(LaunchOutput)
}

type LaunchMapOutput struct{ *pulumi.OutputState }

func (LaunchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Launch)(nil)).Elem()
}

func (o LaunchMapOutput) ToLaunchMapOutput() LaunchMapOutput {
	return o
}

func (o LaunchMapOutput) ToLaunchMapOutputWithContext(ctx context.Context) LaunchMapOutput {
	return o
}

func (o LaunchMapOutput) MapIndex(k pulumi.StringInput) LaunchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Launch {
		return vs[0].(map[string]*Launch)[vs[1].(string)]
	}).(LaunchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LaunchInput)(nil)).Elem(), &Launch{})
	pulumi.RegisterInputType(reflect.TypeOf((*LaunchArrayInput)(nil)).Elem(), LaunchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LaunchMapInput)(nil)).Elem(), LaunchMap{})
	pulumi.RegisterOutputType(LaunchOutput{})
	pulumi.RegisterOutputType(LaunchArrayOutput{})
	pulumi.RegisterOutputType(LaunchMapOutput{})
}
