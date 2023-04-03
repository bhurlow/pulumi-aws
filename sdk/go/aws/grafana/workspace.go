// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Amazon Managed Grafana workspace resource.
//
// ## Example Usage
// ### Basic configuration
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/grafana"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"Version": "2012-10-17",
//				"Statement": []map[string]interface{}{
//					map[string]interface{}{
//						"Action": "sts:AssumeRole",
//						"Effect": "Allow",
//						"Sid":    "",
//						"Principal": map[string]interface{}{
//							"Service": "grafana.amazonaws.com",
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			assume, err := iam.NewRole(ctx, "assume", &iam.RoleArgs{
//				AssumeRolePolicy: pulumi.String(json0),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = grafana.NewWorkspace(ctx, "example", &grafana.WorkspaceArgs{
//				AccountAccessType: pulumi.String("CURRENT_ACCOUNT"),
//				AuthenticationProviders: pulumi.StringArray{
//					pulumi.String("SAML"),
//				},
//				PermissionType: pulumi.String("SERVICE_MANAGED"),
//				RoleArn:        assume.Arn,
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
// Grafana Workspace can be imported using the workspace's `id`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:grafana/workspace:Workspace example g-2054c75a02
//
// ```
type Workspace struct {
	pulumi.CustomResourceState

	// The type of account access for the workspace. Valid values are `CURRENT_ACCOUNT` and `ORGANIZATION`. If `ORGANIZATION` is specified, then `organizationalUnits` must also be present.
	AccountAccessType pulumi.StringOutput `pulumi:"accountAccessType"`
	// The Amazon Resource Name (ARN) of the Grafana workspace.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The authentication providers for the workspace. Valid values are `AWS_SSO`, `SAML`, or both.
	AuthenticationProviders pulumi.StringArrayOutput `pulumi:"authenticationProviders"`
	// The configuration string for the workspace that you create. For more information about the format and configuration options available, see [Working in your Grafana workspace](https://docs.aws.amazon.com/grafana/latest/userguide/AMG-configure-workspace.html).
	Configuration pulumi.StringOutput `pulumi:"configuration"`
	// The data sources for the workspace. Valid values are `AMAZON_OPENSEARCH_SERVICE`, `ATHENA`, `CLOUDWATCH`, `PROMETHEUS`, `REDSHIFT`, `SITEWISE`, `TIMESTREAM`, `XRAY`
	DataSources pulumi.StringArrayOutput `pulumi:"dataSources"`
	// The workspace description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The endpoint of the Grafana workspace.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// The version of Grafana running on the workspace.
	GrafanaVersion pulumi.StringOutput `pulumi:"grafanaVersion"`
	// The Grafana workspace name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Configuration for network access to your workspace.See Network Access Control below.
	NetworkAccessControl WorkspaceNetworkAccessControlPtrOutput `pulumi:"networkAccessControl"`
	// The notification destinations. If a data source is specified here, Amazon Managed Grafana will create IAM roles and permissions needed to use these destinations. Must be set to `SNS`.
	NotificationDestinations pulumi.StringArrayOutput `pulumi:"notificationDestinations"`
	// The role name that the workspace uses to access resources through Amazon Organizations.
	OrganizationRoleName pulumi.StringPtrOutput `pulumi:"organizationRoleName"`
	// The Amazon Organizations organizational units that the workspace is authorized to use data sources from.
	OrganizationalUnits pulumi.StringArrayOutput `pulumi:"organizationalUnits"`
	// The permission type of the workspace. If `SERVICE_MANAGED` is specified, the IAM roles and IAM policy attachments are generated automatically. If `CUSTOMER_MANAGED` is specified, the IAM roles and IAM policy attachments will not be created.
	PermissionType pulumi.StringOutput `pulumi:"permissionType"`
	// The IAM role ARN that the workspace assumes.
	RoleArn                 pulumi.StringPtrOutput `pulumi:"roleArn"`
	SamlConfigurationStatus pulumi.StringOutput    `pulumi:"samlConfigurationStatus"`
	// The AWS CloudFormation stack set name that provisions IAM roles to be used by the workspace.
	StackSetName pulumi.StringPtrOutput `pulumi:"stackSetName"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The configuration settings for an Amazon VPC that contains data sources for your Grafana workspace to connect to. See VPC Configuration below.
	VpcConfiguration WorkspaceVpcConfigurationPtrOutput `pulumi:"vpcConfiguration"`
}

// NewWorkspace registers a new resource with the given unique name, arguments, and options.
func NewWorkspace(ctx *pulumi.Context,
	name string, args *WorkspaceArgs, opts ...pulumi.ResourceOption) (*Workspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountAccessType == nil {
		return nil, errors.New("invalid value for required argument 'AccountAccessType'")
	}
	if args.AuthenticationProviders == nil {
		return nil, errors.New("invalid value for required argument 'AuthenticationProviders'")
	}
	if args.PermissionType == nil {
		return nil, errors.New("invalid value for required argument 'PermissionType'")
	}
	var resource Workspace
	err := ctx.RegisterResource("aws:grafana/workspace:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspace gets an existing Workspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("aws:grafana/workspace:Workspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workspace resources.
type workspaceState struct {
	// The type of account access for the workspace. Valid values are `CURRENT_ACCOUNT` and `ORGANIZATION`. If `ORGANIZATION` is specified, then `organizationalUnits` must also be present.
	AccountAccessType *string `pulumi:"accountAccessType"`
	// The Amazon Resource Name (ARN) of the Grafana workspace.
	Arn *string `pulumi:"arn"`
	// The authentication providers for the workspace. Valid values are `AWS_SSO`, `SAML`, or both.
	AuthenticationProviders []string `pulumi:"authenticationProviders"`
	// The configuration string for the workspace that you create. For more information about the format and configuration options available, see [Working in your Grafana workspace](https://docs.aws.amazon.com/grafana/latest/userguide/AMG-configure-workspace.html).
	Configuration *string `pulumi:"configuration"`
	// The data sources for the workspace. Valid values are `AMAZON_OPENSEARCH_SERVICE`, `ATHENA`, `CLOUDWATCH`, `PROMETHEUS`, `REDSHIFT`, `SITEWISE`, `TIMESTREAM`, `XRAY`
	DataSources []string `pulumi:"dataSources"`
	// The workspace description.
	Description *string `pulumi:"description"`
	// The endpoint of the Grafana workspace.
	Endpoint *string `pulumi:"endpoint"`
	// The version of Grafana running on the workspace.
	GrafanaVersion *string `pulumi:"grafanaVersion"`
	// The Grafana workspace name.
	Name *string `pulumi:"name"`
	// Configuration for network access to your workspace.See Network Access Control below.
	NetworkAccessControl *WorkspaceNetworkAccessControl `pulumi:"networkAccessControl"`
	// The notification destinations. If a data source is specified here, Amazon Managed Grafana will create IAM roles and permissions needed to use these destinations. Must be set to `SNS`.
	NotificationDestinations []string `pulumi:"notificationDestinations"`
	// The role name that the workspace uses to access resources through Amazon Organizations.
	OrganizationRoleName *string `pulumi:"organizationRoleName"`
	// The Amazon Organizations organizational units that the workspace is authorized to use data sources from.
	OrganizationalUnits []string `pulumi:"organizationalUnits"`
	// The permission type of the workspace. If `SERVICE_MANAGED` is specified, the IAM roles and IAM policy attachments are generated automatically. If `CUSTOMER_MANAGED` is specified, the IAM roles and IAM policy attachments will not be created.
	PermissionType *string `pulumi:"permissionType"`
	// The IAM role ARN that the workspace assumes.
	RoleArn                 *string `pulumi:"roleArn"`
	SamlConfigurationStatus *string `pulumi:"samlConfigurationStatus"`
	// The AWS CloudFormation stack set name that provisions IAM roles to be used by the workspace.
	StackSetName *string `pulumi:"stackSetName"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The configuration settings for an Amazon VPC that contains data sources for your Grafana workspace to connect to. See VPC Configuration below.
	VpcConfiguration *WorkspaceVpcConfiguration `pulumi:"vpcConfiguration"`
}

type WorkspaceState struct {
	// The type of account access for the workspace. Valid values are `CURRENT_ACCOUNT` and `ORGANIZATION`. If `ORGANIZATION` is specified, then `organizationalUnits` must also be present.
	AccountAccessType pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the Grafana workspace.
	Arn pulumi.StringPtrInput
	// The authentication providers for the workspace. Valid values are `AWS_SSO`, `SAML`, or both.
	AuthenticationProviders pulumi.StringArrayInput
	// The configuration string for the workspace that you create. For more information about the format and configuration options available, see [Working in your Grafana workspace](https://docs.aws.amazon.com/grafana/latest/userguide/AMG-configure-workspace.html).
	Configuration pulumi.StringPtrInput
	// The data sources for the workspace. Valid values are `AMAZON_OPENSEARCH_SERVICE`, `ATHENA`, `CLOUDWATCH`, `PROMETHEUS`, `REDSHIFT`, `SITEWISE`, `TIMESTREAM`, `XRAY`
	DataSources pulumi.StringArrayInput
	// The workspace description.
	Description pulumi.StringPtrInput
	// The endpoint of the Grafana workspace.
	Endpoint pulumi.StringPtrInput
	// The version of Grafana running on the workspace.
	GrafanaVersion pulumi.StringPtrInput
	// The Grafana workspace name.
	Name pulumi.StringPtrInput
	// Configuration for network access to your workspace.See Network Access Control below.
	NetworkAccessControl WorkspaceNetworkAccessControlPtrInput
	// The notification destinations. If a data source is specified here, Amazon Managed Grafana will create IAM roles and permissions needed to use these destinations. Must be set to `SNS`.
	NotificationDestinations pulumi.StringArrayInput
	// The role name that the workspace uses to access resources through Amazon Organizations.
	OrganizationRoleName pulumi.StringPtrInput
	// The Amazon Organizations organizational units that the workspace is authorized to use data sources from.
	OrganizationalUnits pulumi.StringArrayInput
	// The permission type of the workspace. If `SERVICE_MANAGED` is specified, the IAM roles and IAM policy attachments are generated automatically. If `CUSTOMER_MANAGED` is specified, the IAM roles and IAM policy attachments will not be created.
	PermissionType pulumi.StringPtrInput
	// The IAM role ARN that the workspace assumes.
	RoleArn                 pulumi.StringPtrInput
	SamlConfigurationStatus pulumi.StringPtrInput
	// The AWS CloudFormation stack set name that provisions IAM roles to be used by the workspace.
	StackSetName pulumi.StringPtrInput
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// The configuration settings for an Amazon VPC that contains data sources for your Grafana workspace to connect to. See VPC Configuration below.
	VpcConfiguration WorkspaceVpcConfigurationPtrInput
}

func (WorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceState)(nil)).Elem()
}

type workspaceArgs struct {
	// The type of account access for the workspace. Valid values are `CURRENT_ACCOUNT` and `ORGANIZATION`. If `ORGANIZATION` is specified, then `organizationalUnits` must also be present.
	AccountAccessType string `pulumi:"accountAccessType"`
	// The authentication providers for the workspace. Valid values are `AWS_SSO`, `SAML`, or both.
	AuthenticationProviders []string `pulumi:"authenticationProviders"`
	// The configuration string for the workspace that you create. For more information about the format and configuration options available, see [Working in your Grafana workspace](https://docs.aws.amazon.com/grafana/latest/userguide/AMG-configure-workspace.html).
	Configuration *string `pulumi:"configuration"`
	// The data sources for the workspace. Valid values are `AMAZON_OPENSEARCH_SERVICE`, `ATHENA`, `CLOUDWATCH`, `PROMETHEUS`, `REDSHIFT`, `SITEWISE`, `TIMESTREAM`, `XRAY`
	DataSources []string `pulumi:"dataSources"`
	// The workspace description.
	Description *string `pulumi:"description"`
	// The Grafana workspace name.
	Name *string `pulumi:"name"`
	// Configuration for network access to your workspace.See Network Access Control below.
	NetworkAccessControl *WorkspaceNetworkAccessControl `pulumi:"networkAccessControl"`
	// The notification destinations. If a data source is specified here, Amazon Managed Grafana will create IAM roles and permissions needed to use these destinations. Must be set to `SNS`.
	NotificationDestinations []string `pulumi:"notificationDestinations"`
	// The role name that the workspace uses to access resources through Amazon Organizations.
	OrganizationRoleName *string `pulumi:"organizationRoleName"`
	// The Amazon Organizations organizational units that the workspace is authorized to use data sources from.
	OrganizationalUnits []string `pulumi:"organizationalUnits"`
	// The permission type of the workspace. If `SERVICE_MANAGED` is specified, the IAM roles and IAM policy attachments are generated automatically. If `CUSTOMER_MANAGED` is specified, the IAM roles and IAM policy attachments will not be created.
	PermissionType string `pulumi:"permissionType"`
	// The IAM role ARN that the workspace assumes.
	RoleArn *string `pulumi:"roleArn"`
	// The AWS CloudFormation stack set name that provisions IAM roles to be used by the workspace.
	StackSetName *string `pulumi:"stackSetName"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
	Tags map[string]string `pulumi:"tags"`
	// The configuration settings for an Amazon VPC that contains data sources for your Grafana workspace to connect to. See VPC Configuration below.
	VpcConfiguration *WorkspaceVpcConfiguration `pulumi:"vpcConfiguration"`
}

// The set of arguments for constructing a Workspace resource.
type WorkspaceArgs struct {
	// The type of account access for the workspace. Valid values are `CURRENT_ACCOUNT` and `ORGANIZATION`. If `ORGANIZATION` is specified, then `organizationalUnits` must also be present.
	AccountAccessType pulumi.StringInput
	// The authentication providers for the workspace. Valid values are `AWS_SSO`, `SAML`, or both.
	AuthenticationProviders pulumi.StringArrayInput
	// The configuration string for the workspace that you create. For more information about the format and configuration options available, see [Working in your Grafana workspace](https://docs.aws.amazon.com/grafana/latest/userguide/AMG-configure-workspace.html).
	Configuration pulumi.StringPtrInput
	// The data sources for the workspace. Valid values are `AMAZON_OPENSEARCH_SERVICE`, `ATHENA`, `CLOUDWATCH`, `PROMETHEUS`, `REDSHIFT`, `SITEWISE`, `TIMESTREAM`, `XRAY`
	DataSources pulumi.StringArrayInput
	// The workspace description.
	Description pulumi.StringPtrInput
	// The Grafana workspace name.
	Name pulumi.StringPtrInput
	// Configuration for network access to your workspace.See Network Access Control below.
	NetworkAccessControl WorkspaceNetworkAccessControlPtrInput
	// The notification destinations. If a data source is specified here, Amazon Managed Grafana will create IAM roles and permissions needed to use these destinations. Must be set to `SNS`.
	NotificationDestinations pulumi.StringArrayInput
	// The role name that the workspace uses to access resources through Amazon Organizations.
	OrganizationRoleName pulumi.StringPtrInput
	// The Amazon Organizations organizational units that the workspace is authorized to use data sources from.
	OrganizationalUnits pulumi.StringArrayInput
	// The permission type of the workspace. If `SERVICE_MANAGED` is specified, the IAM roles and IAM policy attachments are generated automatically. If `CUSTOMER_MANAGED` is specified, the IAM roles and IAM policy attachments will not be created.
	PermissionType pulumi.StringInput
	// The IAM role ARN that the workspace assumes.
	RoleArn pulumi.StringPtrInput
	// The AWS CloudFormation stack set name that provisions IAM roles to be used by the workspace.
	StackSetName pulumi.StringPtrInput
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
	Tags pulumi.StringMapInput
	// The configuration settings for an Amazon VPC that contains data sources for your Grafana workspace to connect to. See VPC Configuration below.
	VpcConfiguration WorkspaceVpcConfigurationPtrInput
}

func (WorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceArgs)(nil)).Elem()
}

type WorkspaceInput interface {
	pulumi.Input

	ToWorkspaceOutput() WorkspaceOutput
	ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput
}

func (*Workspace) ElementType() reflect.Type {
	return reflect.TypeOf((**Workspace)(nil)).Elem()
}

func (i *Workspace) ToWorkspaceOutput() WorkspaceOutput {
	return i.ToWorkspaceOutputWithContext(context.Background())
}

func (i *Workspace) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceOutput)
}

// WorkspaceArrayInput is an input type that accepts WorkspaceArray and WorkspaceArrayOutput values.
// You can construct a concrete instance of `WorkspaceArrayInput` via:
//
//	WorkspaceArray{ WorkspaceArgs{...} }
type WorkspaceArrayInput interface {
	pulumi.Input

	ToWorkspaceArrayOutput() WorkspaceArrayOutput
	ToWorkspaceArrayOutputWithContext(context.Context) WorkspaceArrayOutput
}

type WorkspaceArray []WorkspaceInput

func (WorkspaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Workspace)(nil)).Elem()
}

func (i WorkspaceArray) ToWorkspaceArrayOutput() WorkspaceArrayOutput {
	return i.ToWorkspaceArrayOutputWithContext(context.Background())
}

func (i WorkspaceArray) ToWorkspaceArrayOutputWithContext(ctx context.Context) WorkspaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceArrayOutput)
}

// WorkspaceMapInput is an input type that accepts WorkspaceMap and WorkspaceMapOutput values.
// You can construct a concrete instance of `WorkspaceMapInput` via:
//
//	WorkspaceMap{ "key": WorkspaceArgs{...} }
type WorkspaceMapInput interface {
	pulumi.Input

	ToWorkspaceMapOutput() WorkspaceMapOutput
	ToWorkspaceMapOutputWithContext(context.Context) WorkspaceMapOutput
}

type WorkspaceMap map[string]WorkspaceInput

func (WorkspaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Workspace)(nil)).Elem()
}

func (i WorkspaceMap) ToWorkspaceMapOutput() WorkspaceMapOutput {
	return i.ToWorkspaceMapOutputWithContext(context.Background())
}

func (i WorkspaceMap) ToWorkspaceMapOutputWithContext(ctx context.Context) WorkspaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceMapOutput)
}

type WorkspaceOutput struct{ *pulumi.OutputState }

func (WorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workspace)(nil)).Elem()
}

func (o WorkspaceOutput) ToWorkspaceOutput() WorkspaceOutput {
	return o
}

func (o WorkspaceOutput) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return o
}

// The type of account access for the workspace. Valid values are `CURRENT_ACCOUNT` and `ORGANIZATION`. If `ORGANIZATION` is specified, then `organizationalUnits` must also be present.
func (o WorkspaceOutput) AccountAccessType() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.AccountAccessType }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the Grafana workspace.
func (o WorkspaceOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The authentication providers for the workspace. Valid values are `AWS_SSO`, `SAML`, or both.
func (o WorkspaceOutput) AuthenticationProviders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringArrayOutput { return v.AuthenticationProviders }).(pulumi.StringArrayOutput)
}

// The configuration string for the workspace that you create. For more information about the format and configuration options available, see [Working in your Grafana workspace](https://docs.aws.amazon.com/grafana/latest/userguide/AMG-configure-workspace.html).
func (o WorkspaceOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Configuration }).(pulumi.StringOutput)
}

// The data sources for the workspace. Valid values are `AMAZON_OPENSEARCH_SERVICE`, `ATHENA`, `CLOUDWATCH`, `PROMETHEUS`, `REDSHIFT`, `SITEWISE`, `TIMESTREAM`, `XRAY`
func (o WorkspaceOutput) DataSources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringArrayOutput { return v.DataSources }).(pulumi.StringArrayOutput)
}

// The workspace description.
func (o WorkspaceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The endpoint of the Grafana workspace.
func (o WorkspaceOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// The version of Grafana running on the workspace.
func (o WorkspaceOutput) GrafanaVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.GrafanaVersion }).(pulumi.StringOutput)
}

// The Grafana workspace name.
func (o WorkspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Configuration for network access to your workspace.See Network Access Control below.
func (o WorkspaceOutput) NetworkAccessControl() WorkspaceNetworkAccessControlPtrOutput {
	return o.ApplyT(func(v *Workspace) WorkspaceNetworkAccessControlPtrOutput { return v.NetworkAccessControl }).(WorkspaceNetworkAccessControlPtrOutput)
}

// The notification destinations. If a data source is specified here, Amazon Managed Grafana will create IAM roles and permissions needed to use these destinations. Must be set to `SNS`.
func (o WorkspaceOutput) NotificationDestinations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringArrayOutput { return v.NotificationDestinations }).(pulumi.StringArrayOutput)
}

// The role name that the workspace uses to access resources through Amazon Organizations.
func (o WorkspaceOutput) OrganizationRoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.OrganizationRoleName }).(pulumi.StringPtrOutput)
}

// The Amazon Organizations organizational units that the workspace is authorized to use data sources from.
func (o WorkspaceOutput) OrganizationalUnits() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringArrayOutput { return v.OrganizationalUnits }).(pulumi.StringArrayOutput)
}

// The permission type of the workspace. If `SERVICE_MANAGED` is specified, the IAM roles and IAM policy attachments are generated automatically. If `CUSTOMER_MANAGED` is specified, the IAM roles and IAM policy attachments will not be created.
func (o WorkspaceOutput) PermissionType() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.PermissionType }).(pulumi.StringOutput)
}

// The IAM role ARN that the workspace assumes.
func (o WorkspaceOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.RoleArn }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) SamlConfigurationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.SamlConfigurationStatus }).(pulumi.StringOutput)
}

// The AWS CloudFormation stack set name that provisions IAM roles to be used by the workspace.
func (o WorkspaceOutput) StackSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.StackSetName }).(pulumi.StringPtrOutput)
}

// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
func (o WorkspaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o WorkspaceOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The configuration settings for an Amazon VPC that contains data sources for your Grafana workspace to connect to. See VPC Configuration below.
func (o WorkspaceOutput) VpcConfiguration() WorkspaceVpcConfigurationPtrOutput {
	return o.ApplyT(func(v *Workspace) WorkspaceVpcConfigurationPtrOutput { return v.VpcConfiguration }).(WorkspaceVpcConfigurationPtrOutput)
}

type WorkspaceArrayOutput struct{ *pulumi.OutputState }

func (WorkspaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Workspace)(nil)).Elem()
}

func (o WorkspaceArrayOutput) ToWorkspaceArrayOutput() WorkspaceArrayOutput {
	return o
}

func (o WorkspaceArrayOutput) ToWorkspaceArrayOutputWithContext(ctx context.Context) WorkspaceArrayOutput {
	return o
}

func (o WorkspaceArrayOutput) Index(i pulumi.IntInput) WorkspaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Workspace {
		return vs[0].([]*Workspace)[vs[1].(int)]
	}).(WorkspaceOutput)
}

type WorkspaceMapOutput struct{ *pulumi.OutputState }

func (WorkspaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Workspace)(nil)).Elem()
}

func (o WorkspaceMapOutput) ToWorkspaceMapOutput() WorkspaceMapOutput {
	return o
}

func (o WorkspaceMapOutput) ToWorkspaceMapOutputWithContext(ctx context.Context) WorkspaceMapOutput {
	return o
}

func (o WorkspaceMapOutput) MapIndex(k pulumi.StringInput) WorkspaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Workspace {
		return vs[0].(map[string]*Workspace)[vs[1].(string)]
	}).(WorkspaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceInput)(nil)).Elem(), &Workspace{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceArrayInput)(nil)).Elem(), WorkspaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceMapInput)(nil)).Elem(), WorkspaceMap{})
	pulumi.RegisterOutputType(WorkspaceOutput{})
	pulumi.RegisterOutputType(WorkspaceArrayOutput{})
	pulumi.RegisterOutputType(WorkspaceMapOutput{})
}
