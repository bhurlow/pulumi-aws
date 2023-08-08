// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type WorkspaceNetworkAccessControl struct {
	// An array of prefix list IDs.
	PrefixListIds []string `pulumi:"prefixListIds"`
	// An array of Amazon VPC endpoint IDs for the workspace. The only VPC endpoints that can be specified here are interface VPC endpoints for Grafana workspaces (using the com.amazonaws.[region].grafana-workspace service endpoint). Other VPC endpoints will be ignored.
	VpceIds []string `pulumi:"vpceIds"`
}

// WorkspaceNetworkAccessControlInput is an input type that accepts WorkspaceNetworkAccessControlArgs and WorkspaceNetworkAccessControlOutput values.
// You can construct a concrete instance of `WorkspaceNetworkAccessControlInput` via:
//
//	WorkspaceNetworkAccessControlArgs{...}
type WorkspaceNetworkAccessControlInput interface {
	pulumi.Input

	ToWorkspaceNetworkAccessControlOutput() WorkspaceNetworkAccessControlOutput
	ToWorkspaceNetworkAccessControlOutputWithContext(context.Context) WorkspaceNetworkAccessControlOutput
}

type WorkspaceNetworkAccessControlArgs struct {
	// An array of prefix list IDs.
	PrefixListIds pulumi.StringArrayInput `pulumi:"prefixListIds"`
	// An array of Amazon VPC endpoint IDs for the workspace. The only VPC endpoints that can be specified here are interface VPC endpoints for Grafana workspaces (using the com.amazonaws.[region].grafana-workspace service endpoint). Other VPC endpoints will be ignored.
	VpceIds pulumi.StringArrayInput `pulumi:"vpceIds"`
}

func (WorkspaceNetworkAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceNetworkAccessControl)(nil)).Elem()
}

func (i WorkspaceNetworkAccessControlArgs) ToWorkspaceNetworkAccessControlOutput() WorkspaceNetworkAccessControlOutput {
	return i.ToWorkspaceNetworkAccessControlOutputWithContext(context.Background())
}

func (i WorkspaceNetworkAccessControlArgs) ToWorkspaceNetworkAccessControlOutputWithContext(ctx context.Context) WorkspaceNetworkAccessControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceNetworkAccessControlOutput)
}

func (i WorkspaceNetworkAccessControlArgs) ToWorkspaceNetworkAccessControlPtrOutput() WorkspaceNetworkAccessControlPtrOutput {
	return i.ToWorkspaceNetworkAccessControlPtrOutputWithContext(context.Background())
}

func (i WorkspaceNetworkAccessControlArgs) ToWorkspaceNetworkAccessControlPtrOutputWithContext(ctx context.Context) WorkspaceNetworkAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceNetworkAccessControlOutput).ToWorkspaceNetworkAccessControlPtrOutputWithContext(ctx)
}

// WorkspaceNetworkAccessControlPtrInput is an input type that accepts WorkspaceNetworkAccessControlArgs, WorkspaceNetworkAccessControlPtr and WorkspaceNetworkAccessControlPtrOutput values.
// You can construct a concrete instance of `WorkspaceNetworkAccessControlPtrInput` via:
//
//	        WorkspaceNetworkAccessControlArgs{...}
//
//	or:
//
//	        nil
type WorkspaceNetworkAccessControlPtrInput interface {
	pulumi.Input

	ToWorkspaceNetworkAccessControlPtrOutput() WorkspaceNetworkAccessControlPtrOutput
	ToWorkspaceNetworkAccessControlPtrOutputWithContext(context.Context) WorkspaceNetworkAccessControlPtrOutput
}

type workspaceNetworkAccessControlPtrType WorkspaceNetworkAccessControlArgs

func WorkspaceNetworkAccessControlPtr(v *WorkspaceNetworkAccessControlArgs) WorkspaceNetworkAccessControlPtrInput {
	return (*workspaceNetworkAccessControlPtrType)(v)
}

func (*workspaceNetworkAccessControlPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceNetworkAccessControl)(nil)).Elem()
}

func (i *workspaceNetworkAccessControlPtrType) ToWorkspaceNetworkAccessControlPtrOutput() WorkspaceNetworkAccessControlPtrOutput {
	return i.ToWorkspaceNetworkAccessControlPtrOutputWithContext(context.Background())
}

func (i *workspaceNetworkAccessControlPtrType) ToWorkspaceNetworkAccessControlPtrOutputWithContext(ctx context.Context) WorkspaceNetworkAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceNetworkAccessControlPtrOutput)
}

type WorkspaceNetworkAccessControlOutput struct{ *pulumi.OutputState }

func (WorkspaceNetworkAccessControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceNetworkAccessControl)(nil)).Elem()
}

func (o WorkspaceNetworkAccessControlOutput) ToWorkspaceNetworkAccessControlOutput() WorkspaceNetworkAccessControlOutput {
	return o
}

func (o WorkspaceNetworkAccessControlOutput) ToWorkspaceNetworkAccessControlOutputWithContext(ctx context.Context) WorkspaceNetworkAccessControlOutput {
	return o
}

func (o WorkspaceNetworkAccessControlOutput) ToWorkspaceNetworkAccessControlPtrOutput() WorkspaceNetworkAccessControlPtrOutput {
	return o.ToWorkspaceNetworkAccessControlPtrOutputWithContext(context.Background())
}

func (o WorkspaceNetworkAccessControlOutput) ToWorkspaceNetworkAccessControlPtrOutputWithContext(ctx context.Context) WorkspaceNetworkAccessControlPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceNetworkAccessControl) *WorkspaceNetworkAccessControl {
		return &v
	}).(WorkspaceNetworkAccessControlPtrOutput)
}

// An array of prefix list IDs.
func (o WorkspaceNetworkAccessControlOutput) PrefixListIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WorkspaceNetworkAccessControl) []string { return v.PrefixListIds }).(pulumi.StringArrayOutput)
}

// An array of Amazon VPC endpoint IDs for the workspace. The only VPC endpoints that can be specified here are interface VPC endpoints for Grafana workspaces (using the com.amazonaws.[region].grafana-workspace service endpoint). Other VPC endpoints will be ignored.
func (o WorkspaceNetworkAccessControlOutput) VpceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WorkspaceNetworkAccessControl) []string { return v.VpceIds }).(pulumi.StringArrayOutput)
}

type WorkspaceNetworkAccessControlPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceNetworkAccessControlPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceNetworkAccessControl)(nil)).Elem()
}

func (o WorkspaceNetworkAccessControlPtrOutput) ToWorkspaceNetworkAccessControlPtrOutput() WorkspaceNetworkAccessControlPtrOutput {
	return o
}

func (o WorkspaceNetworkAccessControlPtrOutput) ToWorkspaceNetworkAccessControlPtrOutputWithContext(ctx context.Context) WorkspaceNetworkAccessControlPtrOutput {
	return o
}

func (o WorkspaceNetworkAccessControlPtrOutput) Elem() WorkspaceNetworkAccessControlOutput {
	return o.ApplyT(func(v *WorkspaceNetworkAccessControl) WorkspaceNetworkAccessControl {
		if v != nil {
			return *v
		}
		var ret WorkspaceNetworkAccessControl
		return ret
	}).(WorkspaceNetworkAccessControlOutput)
}

// An array of prefix list IDs.
func (o WorkspaceNetworkAccessControlPtrOutput) PrefixListIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkspaceNetworkAccessControl) []string {
		if v == nil {
			return nil
		}
		return v.PrefixListIds
	}).(pulumi.StringArrayOutput)
}

// An array of Amazon VPC endpoint IDs for the workspace. The only VPC endpoints that can be specified here are interface VPC endpoints for Grafana workspaces (using the com.amazonaws.[region].grafana-workspace service endpoint). Other VPC endpoints will be ignored.
func (o WorkspaceNetworkAccessControlPtrOutput) VpceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkspaceNetworkAccessControl) []string {
		if v == nil {
			return nil
		}
		return v.VpceIds
	}).(pulumi.StringArrayOutput)
}

type WorkspaceVpcConfiguration struct {
	// The list of Amazon EC2 security group IDs attached to the Amazon VPC for your Grafana workspace to connect.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The list of Amazon EC2 subnet IDs created in the Amazon VPC for your Grafana workspace to connect.
	SubnetIds []string `pulumi:"subnetIds"`
}

// WorkspaceVpcConfigurationInput is an input type that accepts WorkspaceVpcConfigurationArgs and WorkspaceVpcConfigurationOutput values.
// You can construct a concrete instance of `WorkspaceVpcConfigurationInput` via:
//
//	WorkspaceVpcConfigurationArgs{...}
type WorkspaceVpcConfigurationInput interface {
	pulumi.Input

	ToWorkspaceVpcConfigurationOutput() WorkspaceVpcConfigurationOutput
	ToWorkspaceVpcConfigurationOutputWithContext(context.Context) WorkspaceVpcConfigurationOutput
}

type WorkspaceVpcConfigurationArgs struct {
	// The list of Amazon EC2 security group IDs attached to the Amazon VPC for your Grafana workspace to connect.
	SecurityGroupIds pulumi.StringArrayInput `pulumi:"securityGroupIds"`
	// The list of Amazon EC2 subnet IDs created in the Amazon VPC for your Grafana workspace to connect.
	SubnetIds pulumi.StringArrayInput `pulumi:"subnetIds"`
}

func (WorkspaceVpcConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceVpcConfiguration)(nil)).Elem()
}

func (i WorkspaceVpcConfigurationArgs) ToWorkspaceVpcConfigurationOutput() WorkspaceVpcConfigurationOutput {
	return i.ToWorkspaceVpcConfigurationOutputWithContext(context.Background())
}

func (i WorkspaceVpcConfigurationArgs) ToWorkspaceVpcConfigurationOutputWithContext(ctx context.Context) WorkspaceVpcConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceVpcConfigurationOutput)
}

func (i WorkspaceVpcConfigurationArgs) ToWorkspaceVpcConfigurationPtrOutput() WorkspaceVpcConfigurationPtrOutput {
	return i.ToWorkspaceVpcConfigurationPtrOutputWithContext(context.Background())
}

func (i WorkspaceVpcConfigurationArgs) ToWorkspaceVpcConfigurationPtrOutputWithContext(ctx context.Context) WorkspaceVpcConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceVpcConfigurationOutput).ToWorkspaceVpcConfigurationPtrOutputWithContext(ctx)
}

// WorkspaceVpcConfigurationPtrInput is an input type that accepts WorkspaceVpcConfigurationArgs, WorkspaceVpcConfigurationPtr and WorkspaceVpcConfigurationPtrOutput values.
// You can construct a concrete instance of `WorkspaceVpcConfigurationPtrInput` via:
//
//	        WorkspaceVpcConfigurationArgs{...}
//
//	or:
//
//	        nil
type WorkspaceVpcConfigurationPtrInput interface {
	pulumi.Input

	ToWorkspaceVpcConfigurationPtrOutput() WorkspaceVpcConfigurationPtrOutput
	ToWorkspaceVpcConfigurationPtrOutputWithContext(context.Context) WorkspaceVpcConfigurationPtrOutput
}

type workspaceVpcConfigurationPtrType WorkspaceVpcConfigurationArgs

func WorkspaceVpcConfigurationPtr(v *WorkspaceVpcConfigurationArgs) WorkspaceVpcConfigurationPtrInput {
	return (*workspaceVpcConfigurationPtrType)(v)
}

func (*workspaceVpcConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceVpcConfiguration)(nil)).Elem()
}

func (i *workspaceVpcConfigurationPtrType) ToWorkspaceVpcConfigurationPtrOutput() WorkspaceVpcConfigurationPtrOutput {
	return i.ToWorkspaceVpcConfigurationPtrOutputWithContext(context.Background())
}

func (i *workspaceVpcConfigurationPtrType) ToWorkspaceVpcConfigurationPtrOutputWithContext(ctx context.Context) WorkspaceVpcConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceVpcConfigurationPtrOutput)
}

type WorkspaceVpcConfigurationOutput struct{ *pulumi.OutputState }

func (WorkspaceVpcConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceVpcConfiguration)(nil)).Elem()
}

func (o WorkspaceVpcConfigurationOutput) ToWorkspaceVpcConfigurationOutput() WorkspaceVpcConfigurationOutput {
	return o
}

func (o WorkspaceVpcConfigurationOutput) ToWorkspaceVpcConfigurationOutputWithContext(ctx context.Context) WorkspaceVpcConfigurationOutput {
	return o
}

func (o WorkspaceVpcConfigurationOutput) ToWorkspaceVpcConfigurationPtrOutput() WorkspaceVpcConfigurationPtrOutput {
	return o.ToWorkspaceVpcConfigurationPtrOutputWithContext(context.Background())
}

func (o WorkspaceVpcConfigurationOutput) ToWorkspaceVpcConfigurationPtrOutputWithContext(ctx context.Context) WorkspaceVpcConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceVpcConfiguration) *WorkspaceVpcConfiguration {
		return &v
	}).(WorkspaceVpcConfigurationPtrOutput)
}

// The list of Amazon EC2 security group IDs attached to the Amazon VPC for your Grafana workspace to connect.
func (o WorkspaceVpcConfigurationOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WorkspaceVpcConfiguration) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The list of Amazon EC2 subnet IDs created in the Amazon VPC for your Grafana workspace to connect.
func (o WorkspaceVpcConfigurationOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WorkspaceVpcConfiguration) []string { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

type WorkspaceVpcConfigurationPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceVpcConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceVpcConfiguration)(nil)).Elem()
}

func (o WorkspaceVpcConfigurationPtrOutput) ToWorkspaceVpcConfigurationPtrOutput() WorkspaceVpcConfigurationPtrOutput {
	return o
}

func (o WorkspaceVpcConfigurationPtrOutput) ToWorkspaceVpcConfigurationPtrOutputWithContext(ctx context.Context) WorkspaceVpcConfigurationPtrOutput {
	return o
}

func (o WorkspaceVpcConfigurationPtrOutput) Elem() WorkspaceVpcConfigurationOutput {
	return o.ApplyT(func(v *WorkspaceVpcConfiguration) WorkspaceVpcConfiguration {
		if v != nil {
			return *v
		}
		var ret WorkspaceVpcConfiguration
		return ret
	}).(WorkspaceVpcConfigurationOutput)
}

// The list of Amazon EC2 security group IDs attached to the Amazon VPC for your Grafana workspace to connect.
func (o WorkspaceVpcConfigurationPtrOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkspaceVpcConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.SecurityGroupIds
	}).(pulumi.StringArrayOutput)
}

// The list of Amazon EC2 subnet IDs created in the Amazon VPC for your Grafana workspace to connect.
func (o WorkspaceVpcConfigurationPtrOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkspaceVpcConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.SubnetIds
	}).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceNetworkAccessControlInput)(nil)).Elem(), WorkspaceNetworkAccessControlArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceNetworkAccessControlPtrInput)(nil)).Elem(), WorkspaceNetworkAccessControlArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceVpcConfigurationInput)(nil)).Elem(), WorkspaceVpcConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceVpcConfigurationPtrInput)(nil)).Elem(), WorkspaceVpcConfigurationArgs{})
	pulumi.RegisterOutputType(WorkspaceNetworkAccessControlOutput{})
	pulumi.RegisterOutputType(WorkspaceNetworkAccessControlPtrOutput{})
	pulumi.RegisterOutputType(WorkspaceVpcConfigurationOutput{})
	pulumi.RegisterOutputType(WorkspaceVpcConfigurationPtrOutput{})
}
