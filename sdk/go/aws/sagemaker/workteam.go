// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Sagemaker Workteam resource.
//
// ## Example Usage
// ### Cognito Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/sagemaker"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sagemaker.NewWorkteam(ctx, "example", &sagemaker.WorkteamArgs{
// 			WorkteamName:  pulumi.String("example"),
// 			WorkforceName: pulumi.Any(aws_sagemaker_workforce.Example.Id),
// 			Description:   pulumi.String("example"),
// 			MemberDefinitions: sagemaker.WorkteamMemberDefinitionArray{
// 				&sagemaker.WorkteamMemberDefinitionArgs{
// 					CognitoMemberDefinition: &sagemaker.WorkteamMemberDefinitionCognitoMemberDefinitionArgs{
// 						ClientId:  pulumi.Any(aws_cognito_user_pool_client.Example.Id),
// 						UserPool:  pulumi.Any(aws_cognito_user_pool_domain.Example.User_pool_id),
// 						UserGroup: pulumi.Any(aws_cognito_user_group.Example.Id),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Oidc Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/sagemaker"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sagemaker.NewWorkteam(ctx, "example", &sagemaker.WorkteamArgs{
// 			WorkteamName:  pulumi.String("example"),
// 			WorkforceName: pulumi.Any(aws_sagemaker_workforce.Example.Id),
// 			Description:   pulumi.String("example"),
// 			MemberDefinitions: sagemaker.WorkteamMemberDefinitionArray{
// 				&sagemaker.WorkteamMemberDefinitionArgs{
// 					OidcMemberDefinition: &sagemaker.WorkteamMemberDefinitionOidcMemberDefinitionArgs{
// 						Groups: pulumi.StringArray{
// 							pulumi.String("example"),
// 						},
// 					},
// 				},
// 			},
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
// Sagemaker Workteams can be imported using the `workteam_name`, e.g.
//
// ```sh
//  $ pulumi import aws:sagemaker/workteam:Workteam example example
// ```
type Workteam struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) assigned by AWS to this Workteam.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A description of the work team.
	Description pulumi.StringOutput `pulumi:"description"`
	// A list of Member Definitions that contains objects that identify the workers that make up the work team. Workforces can be created using Amazon Cognito or your own OIDC Identity Provider (IdP). For private workforces created using Amazon Cognito use `cognitoMemberDefinition`. For workforces created using your own OIDC identity provider (IdP) use `oidcMemberDefinition`. Do not provide input for both of these parameters in a single request. see Member Definition details below.
	MemberDefinitions WorkteamMemberDefinitionArrayOutput `pulumi:"memberDefinitions"`
	// Configures notification of workers regarding available or expiring work items. see Notification Configuration details below.
	NotificationConfiguration WorkteamNotificationConfigurationPtrOutput `pulumi:"notificationConfiguration"`
	// The subdomain for your OIDC Identity Provider.
	Subdomain pulumi.StringOutput `pulumi:"subdomain"`
	// A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The name of the Workteam (must be unique).
	WorkforceName pulumi.StringOutput `pulumi:"workforceName"`
	// The name of the workforce.
	WorkteamName pulumi.StringOutput `pulumi:"workteamName"`
}

// NewWorkteam registers a new resource with the given unique name, arguments, and options.
func NewWorkteam(ctx *pulumi.Context,
	name string, args *WorkteamArgs, opts ...pulumi.ResourceOption) (*Workteam, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.MemberDefinitions == nil {
		return nil, errors.New("invalid value for required argument 'MemberDefinitions'")
	}
	if args.WorkforceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkforceName'")
	}
	if args.WorkteamName == nil {
		return nil, errors.New("invalid value for required argument 'WorkteamName'")
	}
	var resource Workteam
	err := ctx.RegisterResource("aws:sagemaker/workteam:Workteam", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkteam gets an existing Workteam resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkteam(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkteamState, opts ...pulumi.ResourceOption) (*Workteam, error) {
	var resource Workteam
	err := ctx.ReadResource("aws:sagemaker/workteam:Workteam", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workteam resources.
type workteamState struct {
	// The Amazon Resource Name (ARN) assigned by AWS to this Workteam.
	Arn *string `pulumi:"arn"`
	// A description of the work team.
	Description *string `pulumi:"description"`
	// A list of Member Definitions that contains objects that identify the workers that make up the work team. Workforces can be created using Amazon Cognito or your own OIDC Identity Provider (IdP). For private workforces created using Amazon Cognito use `cognitoMemberDefinition`. For workforces created using your own OIDC identity provider (IdP) use `oidcMemberDefinition`. Do not provide input for both of these parameters in a single request. see Member Definition details below.
	MemberDefinitions []WorkteamMemberDefinition `pulumi:"memberDefinitions"`
	// Configures notification of workers regarding available or expiring work items. see Notification Configuration details below.
	NotificationConfiguration *WorkteamNotificationConfiguration `pulumi:"notificationConfiguration"`
	// The subdomain for your OIDC Identity Provider.
	Subdomain *string `pulumi:"subdomain"`
	// A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The name of the Workteam (must be unique).
	WorkforceName *string `pulumi:"workforceName"`
	// The name of the workforce.
	WorkteamName *string `pulumi:"workteamName"`
}

type WorkteamState struct {
	// The Amazon Resource Name (ARN) assigned by AWS to this Workteam.
	Arn pulumi.StringPtrInput
	// A description of the work team.
	Description pulumi.StringPtrInput
	// A list of Member Definitions that contains objects that identify the workers that make up the work team. Workforces can be created using Amazon Cognito or your own OIDC Identity Provider (IdP). For private workforces created using Amazon Cognito use `cognitoMemberDefinition`. For workforces created using your own OIDC identity provider (IdP) use `oidcMemberDefinition`. Do not provide input for both of these parameters in a single request. see Member Definition details below.
	MemberDefinitions WorkteamMemberDefinitionArrayInput
	// Configures notification of workers regarding available or expiring work items. see Notification Configuration details below.
	NotificationConfiguration WorkteamNotificationConfigurationPtrInput
	// The subdomain for your OIDC Identity Provider.
	Subdomain pulumi.StringPtrInput
	// A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapInput
	// The name of the Workteam (must be unique).
	WorkforceName pulumi.StringPtrInput
	// The name of the workforce.
	WorkteamName pulumi.StringPtrInput
}

func (WorkteamState) ElementType() reflect.Type {
	return reflect.TypeOf((*workteamState)(nil)).Elem()
}

type workteamArgs struct {
	// A description of the work team.
	Description string `pulumi:"description"`
	// A list of Member Definitions that contains objects that identify the workers that make up the work team. Workforces can be created using Amazon Cognito or your own OIDC Identity Provider (IdP). For private workforces created using Amazon Cognito use `cognitoMemberDefinition`. For workforces created using your own OIDC identity provider (IdP) use `oidcMemberDefinition`. Do not provide input for both of these parameters in a single request. see Member Definition details below.
	MemberDefinitions []WorkteamMemberDefinition `pulumi:"memberDefinitions"`
	// Configures notification of workers regarding available or expiring work items. see Notification Configuration details below.
	NotificationConfiguration *WorkteamNotificationConfiguration `pulumi:"notificationConfiguration"`
	// A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The name of the Workteam (must be unique).
	WorkforceName string `pulumi:"workforceName"`
	// The name of the workforce.
	WorkteamName string `pulumi:"workteamName"`
}

// The set of arguments for constructing a Workteam resource.
type WorkteamArgs struct {
	// A description of the work team.
	Description pulumi.StringInput
	// A list of Member Definitions that contains objects that identify the workers that make up the work team. Workforces can be created using Amazon Cognito or your own OIDC Identity Provider (IdP). For private workforces created using Amazon Cognito use `cognitoMemberDefinition`. For workforces created using your own OIDC identity provider (IdP) use `oidcMemberDefinition`. Do not provide input for both of these parameters in a single request. see Member Definition details below.
	MemberDefinitions WorkteamMemberDefinitionArrayInput
	// Configures notification of workers regarding available or expiring work items. see Notification Configuration details below.
	NotificationConfiguration WorkteamNotificationConfigurationPtrInput
	// A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The name of the Workteam (must be unique).
	WorkforceName pulumi.StringInput
	// The name of the workforce.
	WorkteamName pulumi.StringInput
}

func (WorkteamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workteamArgs)(nil)).Elem()
}

type WorkteamInput interface {
	pulumi.Input

	ToWorkteamOutput() WorkteamOutput
	ToWorkteamOutputWithContext(ctx context.Context) WorkteamOutput
}

func (*Workteam) ElementType() reflect.Type {
	return reflect.TypeOf((*Workteam)(nil))
}

func (i *Workteam) ToWorkteamOutput() WorkteamOutput {
	return i.ToWorkteamOutputWithContext(context.Background())
}

func (i *Workteam) ToWorkteamOutputWithContext(ctx context.Context) WorkteamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkteamOutput)
}

func (i *Workteam) ToWorkteamPtrOutput() WorkteamPtrOutput {
	return i.ToWorkteamPtrOutputWithContext(context.Background())
}

func (i *Workteam) ToWorkteamPtrOutputWithContext(ctx context.Context) WorkteamPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkteamPtrOutput)
}

type WorkteamPtrInput interface {
	pulumi.Input

	ToWorkteamPtrOutput() WorkteamPtrOutput
	ToWorkteamPtrOutputWithContext(ctx context.Context) WorkteamPtrOutput
}

type workteamPtrType WorkteamArgs

func (*workteamPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Workteam)(nil))
}

func (i *workteamPtrType) ToWorkteamPtrOutput() WorkteamPtrOutput {
	return i.ToWorkteamPtrOutputWithContext(context.Background())
}

func (i *workteamPtrType) ToWorkteamPtrOutputWithContext(ctx context.Context) WorkteamPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkteamPtrOutput)
}

// WorkteamArrayInput is an input type that accepts WorkteamArray and WorkteamArrayOutput values.
// You can construct a concrete instance of `WorkteamArrayInput` via:
//
//          WorkteamArray{ WorkteamArgs{...} }
type WorkteamArrayInput interface {
	pulumi.Input

	ToWorkteamArrayOutput() WorkteamArrayOutput
	ToWorkteamArrayOutputWithContext(context.Context) WorkteamArrayOutput
}

type WorkteamArray []WorkteamInput

func (WorkteamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Workteam)(nil)).Elem()
}

func (i WorkteamArray) ToWorkteamArrayOutput() WorkteamArrayOutput {
	return i.ToWorkteamArrayOutputWithContext(context.Background())
}

func (i WorkteamArray) ToWorkteamArrayOutputWithContext(ctx context.Context) WorkteamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkteamArrayOutput)
}

// WorkteamMapInput is an input type that accepts WorkteamMap and WorkteamMapOutput values.
// You can construct a concrete instance of `WorkteamMapInput` via:
//
//          WorkteamMap{ "key": WorkteamArgs{...} }
type WorkteamMapInput interface {
	pulumi.Input

	ToWorkteamMapOutput() WorkteamMapOutput
	ToWorkteamMapOutputWithContext(context.Context) WorkteamMapOutput
}

type WorkteamMap map[string]WorkteamInput

func (WorkteamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Workteam)(nil)).Elem()
}

func (i WorkteamMap) ToWorkteamMapOutput() WorkteamMapOutput {
	return i.ToWorkteamMapOutputWithContext(context.Background())
}

func (i WorkteamMap) ToWorkteamMapOutputWithContext(ctx context.Context) WorkteamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkteamMapOutput)
}

type WorkteamOutput struct{ *pulumi.OutputState }

func (WorkteamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Workteam)(nil))
}

func (o WorkteamOutput) ToWorkteamOutput() WorkteamOutput {
	return o
}

func (o WorkteamOutput) ToWorkteamOutputWithContext(ctx context.Context) WorkteamOutput {
	return o
}

func (o WorkteamOutput) ToWorkteamPtrOutput() WorkteamPtrOutput {
	return o.ToWorkteamPtrOutputWithContext(context.Background())
}

func (o WorkteamOutput) ToWorkteamPtrOutputWithContext(ctx context.Context) WorkteamPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Workteam) *Workteam {
		return &v
	}).(WorkteamPtrOutput)
}

type WorkteamPtrOutput struct{ *pulumi.OutputState }

func (WorkteamPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workteam)(nil))
}

func (o WorkteamPtrOutput) ToWorkteamPtrOutput() WorkteamPtrOutput {
	return o
}

func (o WorkteamPtrOutput) ToWorkteamPtrOutputWithContext(ctx context.Context) WorkteamPtrOutput {
	return o
}

func (o WorkteamPtrOutput) Elem() WorkteamOutput {
	return o.ApplyT(func(v *Workteam) Workteam {
		if v != nil {
			return *v
		}
		var ret Workteam
		return ret
	}).(WorkteamOutput)
}

type WorkteamArrayOutput struct{ *pulumi.OutputState }

func (WorkteamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Workteam)(nil))
}

func (o WorkteamArrayOutput) ToWorkteamArrayOutput() WorkteamArrayOutput {
	return o
}

func (o WorkteamArrayOutput) ToWorkteamArrayOutputWithContext(ctx context.Context) WorkteamArrayOutput {
	return o
}

func (o WorkteamArrayOutput) Index(i pulumi.IntInput) WorkteamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Workteam {
		return vs[0].([]Workteam)[vs[1].(int)]
	}).(WorkteamOutput)
}

type WorkteamMapOutput struct{ *pulumi.OutputState }

func (WorkteamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Workteam)(nil))
}

func (o WorkteamMapOutput) ToWorkteamMapOutput() WorkteamMapOutput {
	return o
}

func (o WorkteamMapOutput) ToWorkteamMapOutputWithContext(ctx context.Context) WorkteamMapOutput {
	return o
}

func (o WorkteamMapOutput) MapIndex(k pulumi.StringInput) WorkteamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Workteam {
		return vs[0].(map[string]Workteam)[vs[1].(string)]
	}).(WorkteamOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkteamOutput{})
	pulumi.RegisterOutputType(WorkteamPtrOutput{})
	pulumi.RegisterOutputType(WorkteamArrayOutput{})
	pulumi.RegisterOutputType(WorkteamMapOutput{})
}
