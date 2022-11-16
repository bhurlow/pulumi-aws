// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an SSM Patch Baseline resource.
//
// > **NOTE on Patch Baselines:** The `approvedPatches` and `approvalRule` are
// both marked as optional fields, but the Patch Baseline requires that at least one
// of them is specified.
//
// ## Example Usage
// ### Basic Usage
//
// Using `approvedPatches` only.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ssm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ssm.NewPatchBaseline(ctx, "production", &ssm.PatchBaselineArgs{
//				ApprovedPatches: pulumi.StringArray{
//					pulumi.String("KB123456"),
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
// ### Advanced Usage, specifying patch filters
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ssm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ssm.NewPatchBaseline(ctx, "production", &ssm.PatchBaselineArgs{
//				ApprovalRules: ssm.PatchBaselineApprovalRuleArray{
//					&ssm.PatchBaselineApprovalRuleArgs{
//						ApproveAfterDays: pulumi.Int(7),
//						ComplianceLevel:  pulumi.String("HIGH"),
//						PatchFilters: ssm.PatchBaselineApprovalRulePatchFilterArray{
//							&ssm.PatchBaselineApprovalRulePatchFilterArgs{
//								Key: pulumi.String("PRODUCT"),
//								Values: pulumi.StringArray{
//									pulumi.String("WindowsServer2016"),
//								},
//							},
//							&ssm.PatchBaselineApprovalRulePatchFilterArgs{
//								Key: pulumi.String("CLASSIFICATION"),
//								Values: pulumi.StringArray{
//									pulumi.String("CriticalUpdates"),
//									pulumi.String("SecurityUpdates"),
//									pulumi.String("Updates"),
//								},
//							},
//							&ssm.PatchBaselineApprovalRulePatchFilterArgs{
//								Key: pulumi.String("MSRC_SEVERITY"),
//								Values: pulumi.StringArray{
//									pulumi.String("Critical"),
//									pulumi.String("Important"),
//									pulumi.String("Moderate"),
//								},
//							},
//						},
//					},
//					&ssm.PatchBaselineApprovalRuleArgs{
//						ApproveAfterDays: pulumi.Int(7),
//						PatchFilters: ssm.PatchBaselineApprovalRulePatchFilterArray{
//							&ssm.PatchBaselineApprovalRulePatchFilterArgs{
//								Key: pulumi.String("PRODUCT"),
//								Values: pulumi.StringArray{
//									pulumi.String("WindowsServer2012"),
//								},
//							},
//						},
//					},
//				},
//				ApprovedPatches: pulumi.StringArray{
//					pulumi.String("KB123456"),
//					pulumi.String("KB456789"),
//				},
//				Description: pulumi.String("Patch Baseline Description"),
//				GlobalFilters: ssm.PatchBaselineGlobalFilterArray{
//					&ssm.PatchBaselineGlobalFilterArgs{
//						Key: pulumi.String("PRODUCT"),
//						Values: pulumi.StringArray{
//							pulumi.String("WindowsServer2008"),
//						},
//					},
//					&ssm.PatchBaselineGlobalFilterArgs{
//						Key: pulumi.String("CLASSIFICATION"),
//						Values: pulumi.StringArray{
//							pulumi.String("ServicePacks"),
//						},
//					},
//					&ssm.PatchBaselineGlobalFilterArgs{
//						Key: pulumi.String("MSRC_SEVERITY"),
//						Values: pulumi.StringArray{
//							pulumi.String("Low"),
//						},
//					},
//				},
//				RejectedPatches: pulumi.StringArray{
//					pulumi.String("KB987654"),
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
// ### Advanced usage, specifying Microsoft application and Windows patch rules
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ssm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ssm.NewPatchBaseline(ctx, "windowsOsApps", &ssm.PatchBaselineArgs{
//				ApprovalRules: ssm.PatchBaselineApprovalRuleArray{
//					&ssm.PatchBaselineApprovalRuleArgs{
//						ApproveAfterDays: pulumi.Int(7),
//						PatchFilters: ssm.PatchBaselineApprovalRulePatchFilterArray{
//							&ssm.PatchBaselineApprovalRulePatchFilterArgs{
//								Key: pulumi.String("CLASSIFICATION"),
//								Values: pulumi.StringArray{
//									pulumi.String("CriticalUpdates"),
//									pulumi.String("SecurityUpdates"),
//								},
//							},
//							&ssm.PatchBaselineApprovalRulePatchFilterArgs{
//								Key: pulumi.String("MSRC_SEVERITY"),
//								Values: pulumi.StringArray{
//									pulumi.String("Critical"),
//									pulumi.String("Important"),
//								},
//							},
//						},
//					},
//					&ssm.PatchBaselineApprovalRuleArgs{
//						ApproveAfterDays: pulumi.Int(7),
//						PatchFilters: ssm.PatchBaselineApprovalRulePatchFilterArray{
//							&ssm.PatchBaselineApprovalRulePatchFilterArgs{
//								Key: pulumi.String("PATCH_SET"),
//								Values: pulumi.StringArray{
//									pulumi.String("APPLICATION"),
//								},
//							},
//							&ssm.PatchBaselineApprovalRulePatchFilterArgs{
//								Key: pulumi.String("PRODUCT"),
//								Values: pulumi.StringArray{
//									pulumi.String("Office 2013"),
//									pulumi.String("Office 2016"),
//								},
//							},
//						},
//					},
//				},
//				Description:     pulumi.String("Patch both Windows and Microsoft apps"),
//				OperatingSystem: pulumi.String("WINDOWS"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Advanced usage, specifying alternate patch source repository
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ssm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ssm.NewPatchBaseline(ctx, "al201709", &ssm.PatchBaselineArgs{
//				ApprovalRules: ssm.PatchBaselineApprovalRuleArray{
//					nil,
//				},
//				Description:     pulumi.String("My patch repository for Amazon Linux 2017.09"),
//				OperatingSystem: pulumi.String("AMAZON_LINUX"),
//				Sources: ssm.PatchBaselineSourceArray{
//					&ssm.PatchBaselineSourceArgs{
//						Configuration: pulumi.String(fmt.Sprintf(`[amzn-main]
//
// name=amzn-main-Base
// mirrorlist=http://repo./$awsregion./$awsdomain//$releasever/main/mirror.list
// mirrorlist_expire=300
// metadata_expire=300
// priority=10
// failovermethod=priority
// fastestmirror_enabled=0
// gpgcheck=1
// gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-amazon-ga
// enabled=1
// retries=3
// timeout=5
// report_instanceid=yes
//
// `)),
//
//						Name: pulumi.String("My-AL2017.09"),
//						Products: pulumi.StringArray{
//							pulumi.String("AmazonLinux2017.09"),
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
// SSM Patch Baselines can be imported by their baseline ID, e.g.,
//
// ```sh
//
//	$ pulumi import aws:ssm/patchBaseline:PatchBaseline example pb-12345678
//
// ```
type PatchBaseline struct {
	pulumi.CustomResourceState

	// A set of rules used to include patches in the baseline.
	// Up to 10 approval rules can be specified.
	// See `approvalRule` below.
	ApprovalRules PatchBaselineApprovalRuleArrayOutput `pulumi:"approvalRules"`
	// A list of explicitly approved patches for the baseline.
	// Cannot be specified with `approvalRule`.
	ApprovedPatches pulumi.StringArrayOutput `pulumi:"approvedPatches"`
	// The compliance level for approved patches.
	// This means that if an approved patch is reported as missing, this is the severity of the compliance violation.
	// Valid values are `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, `INFORMATIONAL`, `UNSPECIFIED`.
	// The default value is `UNSPECIFIED`.
	ApprovedPatchesComplianceLevel pulumi.StringPtrOutput `pulumi:"approvedPatchesComplianceLevel"`
	// Indicates whether the list of approved patches includes non-security updates that should be applied to the instances.
	// Applies to Linux instances only.
	ApprovedPatchesEnableNonSecurity pulumi.BoolPtrOutput `pulumi:"approvedPatchesEnableNonSecurity"`
	// The ARN of the patch baseline.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the patch baseline.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A set of global filters used to exclude patches from the baseline.
	// Up to 4 global filters can be specified using Key/Value pairs.
	// Valid Keys are `PRODUCT`, `CLASSIFICATION`, `MSRC_SEVERITY`, and `PATCH_ID`.
	GlobalFilters PatchBaselineGlobalFilterArrayOutput `pulumi:"globalFilters"`
	// The name specified to identify the patch source.
	Name pulumi.StringOutput `pulumi:"name"`
	// The operating system the patch baseline applies to.
	// Valid values are
	// `AMAZON_LINUX`,
	// `AMAZON_LINUX_2`,
	// `AMAZON_LINUX_2022`,
	// `CENTOS`,
	// `DEBIAN`,
	// `MACOS`,
	// `ORACLE_LINUX`,
	// `RASPBIAN`,
	// `REDHAT_ENTERPRISE_LINUX`,
	// `ROCKY_LINUX`,
	// `SUSE`,
	// `UBUNTU`, and
	// `WINDOWS`.
	// The default value is `WINDOWS`.
	OperatingSystem pulumi.StringPtrOutput `pulumi:"operatingSystem"`
	// A list of rejected patches.
	RejectedPatches pulumi.StringArrayOutput `pulumi:"rejectedPatches"`
	// The action for Patch Manager to take on patches included in the `rejectedPatches` list.
	// Valid values are `ALLOW_AS_DEPENDENCY` and `BLOCK`.
	RejectedPatchesAction pulumi.StringOutput `pulumi:"rejectedPatchesAction"`
	// Configuration block with alternate sources for patches.
	// Applies to Linux instances only.
	// See `source` below.
	Sources PatchBaselineSourceArrayOutput `pulumi:"sources"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewPatchBaseline registers a new resource with the given unique name, arguments, and options.
func NewPatchBaseline(ctx *pulumi.Context,
	name string, args *PatchBaselineArgs, opts ...pulumi.ResourceOption) (*PatchBaseline, error) {
	if args == nil {
		args = &PatchBaselineArgs{}
	}

	var resource PatchBaseline
	err := ctx.RegisterResource("aws:ssm/patchBaseline:PatchBaseline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPatchBaseline gets an existing PatchBaseline resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPatchBaseline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PatchBaselineState, opts ...pulumi.ResourceOption) (*PatchBaseline, error) {
	var resource PatchBaseline
	err := ctx.ReadResource("aws:ssm/patchBaseline:PatchBaseline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PatchBaseline resources.
type patchBaselineState struct {
	// A set of rules used to include patches in the baseline.
	// Up to 10 approval rules can be specified.
	// See `approvalRule` below.
	ApprovalRules []PatchBaselineApprovalRule `pulumi:"approvalRules"`
	// A list of explicitly approved patches for the baseline.
	// Cannot be specified with `approvalRule`.
	ApprovedPatches []string `pulumi:"approvedPatches"`
	// The compliance level for approved patches.
	// This means that if an approved patch is reported as missing, this is the severity of the compliance violation.
	// Valid values are `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, `INFORMATIONAL`, `UNSPECIFIED`.
	// The default value is `UNSPECIFIED`.
	ApprovedPatchesComplianceLevel *string `pulumi:"approvedPatchesComplianceLevel"`
	// Indicates whether the list of approved patches includes non-security updates that should be applied to the instances.
	// Applies to Linux instances only.
	ApprovedPatchesEnableNonSecurity *bool `pulumi:"approvedPatchesEnableNonSecurity"`
	// The ARN of the patch baseline.
	Arn *string `pulumi:"arn"`
	// The description of the patch baseline.
	Description *string `pulumi:"description"`
	// A set of global filters used to exclude patches from the baseline.
	// Up to 4 global filters can be specified using Key/Value pairs.
	// Valid Keys are `PRODUCT`, `CLASSIFICATION`, `MSRC_SEVERITY`, and `PATCH_ID`.
	GlobalFilters []PatchBaselineGlobalFilter `pulumi:"globalFilters"`
	// The name specified to identify the patch source.
	Name *string `pulumi:"name"`
	// The operating system the patch baseline applies to.
	// Valid values are
	// `AMAZON_LINUX`,
	// `AMAZON_LINUX_2`,
	// `AMAZON_LINUX_2022`,
	// `CENTOS`,
	// `DEBIAN`,
	// `MACOS`,
	// `ORACLE_LINUX`,
	// `RASPBIAN`,
	// `REDHAT_ENTERPRISE_LINUX`,
	// `ROCKY_LINUX`,
	// `SUSE`,
	// `UBUNTU`, and
	// `WINDOWS`.
	// The default value is `WINDOWS`.
	OperatingSystem *string `pulumi:"operatingSystem"`
	// A list of rejected patches.
	RejectedPatches []string `pulumi:"rejectedPatches"`
	// The action for Patch Manager to take on patches included in the `rejectedPatches` list.
	// Valid values are `ALLOW_AS_DEPENDENCY` and `BLOCK`.
	RejectedPatchesAction *string `pulumi:"rejectedPatchesAction"`
	// Configuration block with alternate sources for patches.
	// Applies to Linux instances only.
	// See `source` below.
	Sources []PatchBaselineSource `pulumi:"sources"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type PatchBaselineState struct {
	// A set of rules used to include patches in the baseline.
	// Up to 10 approval rules can be specified.
	// See `approvalRule` below.
	ApprovalRules PatchBaselineApprovalRuleArrayInput
	// A list of explicitly approved patches for the baseline.
	// Cannot be specified with `approvalRule`.
	ApprovedPatches pulumi.StringArrayInput
	// The compliance level for approved patches.
	// This means that if an approved patch is reported as missing, this is the severity of the compliance violation.
	// Valid values are `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, `INFORMATIONAL`, `UNSPECIFIED`.
	// The default value is `UNSPECIFIED`.
	ApprovedPatchesComplianceLevel pulumi.StringPtrInput
	// Indicates whether the list of approved patches includes non-security updates that should be applied to the instances.
	// Applies to Linux instances only.
	ApprovedPatchesEnableNonSecurity pulumi.BoolPtrInput
	// The ARN of the patch baseline.
	Arn pulumi.StringPtrInput
	// The description of the patch baseline.
	Description pulumi.StringPtrInput
	// A set of global filters used to exclude patches from the baseline.
	// Up to 4 global filters can be specified using Key/Value pairs.
	// Valid Keys are `PRODUCT`, `CLASSIFICATION`, `MSRC_SEVERITY`, and `PATCH_ID`.
	GlobalFilters PatchBaselineGlobalFilterArrayInput
	// The name specified to identify the patch source.
	Name pulumi.StringPtrInput
	// The operating system the patch baseline applies to.
	// Valid values are
	// `AMAZON_LINUX`,
	// `AMAZON_LINUX_2`,
	// `AMAZON_LINUX_2022`,
	// `CENTOS`,
	// `DEBIAN`,
	// `MACOS`,
	// `ORACLE_LINUX`,
	// `RASPBIAN`,
	// `REDHAT_ENTERPRISE_LINUX`,
	// `ROCKY_LINUX`,
	// `SUSE`,
	// `UBUNTU`, and
	// `WINDOWS`.
	// The default value is `WINDOWS`.
	OperatingSystem pulumi.StringPtrInput
	// A list of rejected patches.
	RejectedPatches pulumi.StringArrayInput
	// The action for Patch Manager to take on patches included in the `rejectedPatches` list.
	// Valid values are `ALLOW_AS_DEPENDENCY` and `BLOCK`.
	RejectedPatchesAction pulumi.StringPtrInput
	// Configuration block with alternate sources for patches.
	// Applies to Linux instances only.
	// See `source` below.
	Sources PatchBaselineSourceArrayInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (PatchBaselineState) ElementType() reflect.Type {
	return reflect.TypeOf((*patchBaselineState)(nil)).Elem()
}

type patchBaselineArgs struct {
	// A set of rules used to include patches in the baseline.
	// Up to 10 approval rules can be specified.
	// See `approvalRule` below.
	ApprovalRules []PatchBaselineApprovalRule `pulumi:"approvalRules"`
	// A list of explicitly approved patches for the baseline.
	// Cannot be specified with `approvalRule`.
	ApprovedPatches []string `pulumi:"approvedPatches"`
	// The compliance level for approved patches.
	// This means that if an approved patch is reported as missing, this is the severity of the compliance violation.
	// Valid values are `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, `INFORMATIONAL`, `UNSPECIFIED`.
	// The default value is `UNSPECIFIED`.
	ApprovedPatchesComplianceLevel *string `pulumi:"approvedPatchesComplianceLevel"`
	// Indicates whether the list of approved patches includes non-security updates that should be applied to the instances.
	// Applies to Linux instances only.
	ApprovedPatchesEnableNonSecurity *bool `pulumi:"approvedPatchesEnableNonSecurity"`
	// The description of the patch baseline.
	Description *string `pulumi:"description"`
	// A set of global filters used to exclude patches from the baseline.
	// Up to 4 global filters can be specified using Key/Value pairs.
	// Valid Keys are `PRODUCT`, `CLASSIFICATION`, `MSRC_SEVERITY`, and `PATCH_ID`.
	GlobalFilters []PatchBaselineGlobalFilter `pulumi:"globalFilters"`
	// The name specified to identify the patch source.
	Name *string `pulumi:"name"`
	// The operating system the patch baseline applies to.
	// Valid values are
	// `AMAZON_LINUX`,
	// `AMAZON_LINUX_2`,
	// `AMAZON_LINUX_2022`,
	// `CENTOS`,
	// `DEBIAN`,
	// `MACOS`,
	// `ORACLE_LINUX`,
	// `RASPBIAN`,
	// `REDHAT_ENTERPRISE_LINUX`,
	// `ROCKY_LINUX`,
	// `SUSE`,
	// `UBUNTU`, and
	// `WINDOWS`.
	// The default value is `WINDOWS`.
	OperatingSystem *string `pulumi:"operatingSystem"`
	// A list of rejected patches.
	RejectedPatches []string `pulumi:"rejectedPatches"`
	// The action for Patch Manager to take on patches included in the `rejectedPatches` list.
	// Valid values are `ALLOW_AS_DEPENDENCY` and `BLOCK`.
	RejectedPatchesAction *string `pulumi:"rejectedPatchesAction"`
	// Configuration block with alternate sources for patches.
	// Applies to Linux instances only.
	// See `source` below.
	Sources []PatchBaselineSource `pulumi:"sources"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PatchBaseline resource.
type PatchBaselineArgs struct {
	// A set of rules used to include patches in the baseline.
	// Up to 10 approval rules can be specified.
	// See `approvalRule` below.
	ApprovalRules PatchBaselineApprovalRuleArrayInput
	// A list of explicitly approved patches for the baseline.
	// Cannot be specified with `approvalRule`.
	ApprovedPatches pulumi.StringArrayInput
	// The compliance level for approved patches.
	// This means that if an approved patch is reported as missing, this is the severity of the compliance violation.
	// Valid values are `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, `INFORMATIONAL`, `UNSPECIFIED`.
	// The default value is `UNSPECIFIED`.
	ApprovedPatchesComplianceLevel pulumi.StringPtrInput
	// Indicates whether the list of approved patches includes non-security updates that should be applied to the instances.
	// Applies to Linux instances only.
	ApprovedPatchesEnableNonSecurity pulumi.BoolPtrInput
	// The description of the patch baseline.
	Description pulumi.StringPtrInput
	// A set of global filters used to exclude patches from the baseline.
	// Up to 4 global filters can be specified using Key/Value pairs.
	// Valid Keys are `PRODUCT`, `CLASSIFICATION`, `MSRC_SEVERITY`, and `PATCH_ID`.
	GlobalFilters PatchBaselineGlobalFilterArrayInput
	// The name specified to identify the patch source.
	Name pulumi.StringPtrInput
	// The operating system the patch baseline applies to.
	// Valid values are
	// `AMAZON_LINUX`,
	// `AMAZON_LINUX_2`,
	// `AMAZON_LINUX_2022`,
	// `CENTOS`,
	// `DEBIAN`,
	// `MACOS`,
	// `ORACLE_LINUX`,
	// `RASPBIAN`,
	// `REDHAT_ENTERPRISE_LINUX`,
	// `ROCKY_LINUX`,
	// `SUSE`,
	// `UBUNTU`, and
	// `WINDOWS`.
	// The default value is `WINDOWS`.
	OperatingSystem pulumi.StringPtrInput
	// A list of rejected patches.
	RejectedPatches pulumi.StringArrayInput
	// The action for Patch Manager to take on patches included in the `rejectedPatches` list.
	// Valid values are `ALLOW_AS_DEPENDENCY` and `BLOCK`.
	RejectedPatchesAction pulumi.StringPtrInput
	// Configuration block with alternate sources for patches.
	// Applies to Linux instances only.
	// See `source` below.
	Sources PatchBaselineSourceArrayInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (PatchBaselineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*patchBaselineArgs)(nil)).Elem()
}

type PatchBaselineInput interface {
	pulumi.Input

	ToPatchBaselineOutput() PatchBaselineOutput
	ToPatchBaselineOutputWithContext(ctx context.Context) PatchBaselineOutput
}

func (*PatchBaseline) ElementType() reflect.Type {
	return reflect.TypeOf((**PatchBaseline)(nil)).Elem()
}

func (i *PatchBaseline) ToPatchBaselineOutput() PatchBaselineOutput {
	return i.ToPatchBaselineOutputWithContext(context.Background())
}

func (i *PatchBaseline) ToPatchBaselineOutputWithContext(ctx context.Context) PatchBaselineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PatchBaselineOutput)
}

// PatchBaselineArrayInput is an input type that accepts PatchBaselineArray and PatchBaselineArrayOutput values.
// You can construct a concrete instance of `PatchBaselineArrayInput` via:
//
//	PatchBaselineArray{ PatchBaselineArgs{...} }
type PatchBaselineArrayInput interface {
	pulumi.Input

	ToPatchBaselineArrayOutput() PatchBaselineArrayOutput
	ToPatchBaselineArrayOutputWithContext(context.Context) PatchBaselineArrayOutput
}

type PatchBaselineArray []PatchBaselineInput

func (PatchBaselineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PatchBaseline)(nil)).Elem()
}

func (i PatchBaselineArray) ToPatchBaselineArrayOutput() PatchBaselineArrayOutput {
	return i.ToPatchBaselineArrayOutputWithContext(context.Background())
}

func (i PatchBaselineArray) ToPatchBaselineArrayOutputWithContext(ctx context.Context) PatchBaselineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PatchBaselineArrayOutput)
}

// PatchBaselineMapInput is an input type that accepts PatchBaselineMap and PatchBaselineMapOutput values.
// You can construct a concrete instance of `PatchBaselineMapInput` via:
//
//	PatchBaselineMap{ "key": PatchBaselineArgs{...} }
type PatchBaselineMapInput interface {
	pulumi.Input

	ToPatchBaselineMapOutput() PatchBaselineMapOutput
	ToPatchBaselineMapOutputWithContext(context.Context) PatchBaselineMapOutput
}

type PatchBaselineMap map[string]PatchBaselineInput

func (PatchBaselineMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PatchBaseline)(nil)).Elem()
}

func (i PatchBaselineMap) ToPatchBaselineMapOutput() PatchBaselineMapOutput {
	return i.ToPatchBaselineMapOutputWithContext(context.Background())
}

func (i PatchBaselineMap) ToPatchBaselineMapOutputWithContext(ctx context.Context) PatchBaselineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PatchBaselineMapOutput)
}

type PatchBaselineOutput struct{ *pulumi.OutputState }

func (PatchBaselineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PatchBaseline)(nil)).Elem()
}

func (o PatchBaselineOutput) ToPatchBaselineOutput() PatchBaselineOutput {
	return o
}

func (o PatchBaselineOutput) ToPatchBaselineOutputWithContext(ctx context.Context) PatchBaselineOutput {
	return o
}

// A set of rules used to include patches in the baseline.
// Up to 10 approval rules can be specified.
// See `approvalRule` below.
func (o PatchBaselineOutput) ApprovalRules() PatchBaselineApprovalRuleArrayOutput {
	return o.ApplyT(func(v *PatchBaseline) PatchBaselineApprovalRuleArrayOutput { return v.ApprovalRules }).(PatchBaselineApprovalRuleArrayOutput)
}

// A list of explicitly approved patches for the baseline.
// Cannot be specified with `approvalRule`.
func (o PatchBaselineOutput) ApprovedPatches() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PatchBaseline) pulumi.StringArrayOutput { return v.ApprovedPatches }).(pulumi.StringArrayOutput)
}

// The compliance level for approved patches.
// This means that if an approved patch is reported as missing, this is the severity of the compliance violation.
// Valid values are `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, `INFORMATIONAL`, `UNSPECIFIED`.
// The default value is `UNSPECIFIED`.
func (o PatchBaselineOutput) ApprovedPatchesComplianceLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PatchBaseline) pulumi.StringPtrOutput { return v.ApprovedPatchesComplianceLevel }).(pulumi.StringPtrOutput)
}

// Indicates whether the list of approved patches includes non-security updates that should be applied to the instances.
// Applies to Linux instances only.
func (o PatchBaselineOutput) ApprovedPatchesEnableNonSecurity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PatchBaseline) pulumi.BoolPtrOutput { return v.ApprovedPatchesEnableNonSecurity }).(pulumi.BoolPtrOutput)
}

// The ARN of the patch baseline.
func (o PatchBaselineOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *PatchBaseline) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The description of the patch baseline.
func (o PatchBaselineOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PatchBaseline) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A set of global filters used to exclude patches from the baseline.
// Up to 4 global filters can be specified using Key/Value pairs.
// Valid Keys are `PRODUCT`, `CLASSIFICATION`, `MSRC_SEVERITY`, and `PATCH_ID`.
func (o PatchBaselineOutput) GlobalFilters() PatchBaselineGlobalFilterArrayOutput {
	return o.ApplyT(func(v *PatchBaseline) PatchBaselineGlobalFilterArrayOutput { return v.GlobalFilters }).(PatchBaselineGlobalFilterArrayOutput)
}

// The name specified to identify the patch source.
func (o PatchBaselineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PatchBaseline) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The operating system the patch baseline applies to.
// Valid values are
// `AMAZON_LINUX`,
// `AMAZON_LINUX_2`,
// `AMAZON_LINUX_2022`,
// `CENTOS`,
// `DEBIAN`,
// `MACOS`,
// `ORACLE_LINUX`,
// `RASPBIAN`,
// `REDHAT_ENTERPRISE_LINUX`,
// `ROCKY_LINUX`,
// `SUSE`,
// `UBUNTU`, and
// `WINDOWS`.
// The default value is `WINDOWS`.
func (o PatchBaselineOutput) OperatingSystem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PatchBaseline) pulumi.StringPtrOutput { return v.OperatingSystem }).(pulumi.StringPtrOutput)
}

// A list of rejected patches.
func (o PatchBaselineOutput) RejectedPatches() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PatchBaseline) pulumi.StringArrayOutput { return v.RejectedPatches }).(pulumi.StringArrayOutput)
}

// The action for Patch Manager to take on patches included in the `rejectedPatches` list.
// Valid values are `ALLOW_AS_DEPENDENCY` and `BLOCK`.
func (o PatchBaselineOutput) RejectedPatchesAction() pulumi.StringOutput {
	return o.ApplyT(func(v *PatchBaseline) pulumi.StringOutput { return v.RejectedPatchesAction }).(pulumi.StringOutput)
}

// Configuration block with alternate sources for patches.
// Applies to Linux instances only.
// See `source` below.
func (o PatchBaselineOutput) Sources() PatchBaselineSourceArrayOutput {
	return o.ApplyT(func(v *PatchBaseline) PatchBaselineSourceArrayOutput { return v.Sources }).(PatchBaselineSourceArrayOutput)
}

// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o PatchBaselineOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PatchBaseline) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o PatchBaselineOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PatchBaseline) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type PatchBaselineArrayOutput struct{ *pulumi.OutputState }

func (PatchBaselineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PatchBaseline)(nil)).Elem()
}

func (o PatchBaselineArrayOutput) ToPatchBaselineArrayOutput() PatchBaselineArrayOutput {
	return o
}

func (o PatchBaselineArrayOutput) ToPatchBaselineArrayOutputWithContext(ctx context.Context) PatchBaselineArrayOutput {
	return o
}

func (o PatchBaselineArrayOutput) Index(i pulumi.IntInput) PatchBaselineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PatchBaseline {
		return vs[0].([]*PatchBaseline)[vs[1].(int)]
	}).(PatchBaselineOutput)
}

type PatchBaselineMapOutput struct{ *pulumi.OutputState }

func (PatchBaselineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PatchBaseline)(nil)).Elem()
}

func (o PatchBaselineMapOutput) ToPatchBaselineMapOutput() PatchBaselineMapOutput {
	return o
}

func (o PatchBaselineMapOutput) ToPatchBaselineMapOutputWithContext(ctx context.Context) PatchBaselineMapOutput {
	return o
}

func (o PatchBaselineMapOutput) MapIndex(k pulumi.StringInput) PatchBaselineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PatchBaseline {
		return vs[0].(map[string]*PatchBaseline)[vs[1].(string)]
	}).(PatchBaselineOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PatchBaselineInput)(nil)).Elem(), &PatchBaseline{})
	pulumi.RegisterInputType(reflect.TypeOf((*PatchBaselineArrayInput)(nil)).Elem(), PatchBaselineArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PatchBaselineMapInput)(nil)).Elem(), PatchBaselineMap{})
	pulumi.RegisterOutputType(PatchBaselineOutput{})
	pulumi.RegisterOutputType(PatchBaselineArrayOutput{})
	pulumi.RegisterOutputType(PatchBaselineMapOutput{})
}
