// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sesv2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for managing an AWS SESv2 (Simple Email V2) Email Identity.
//
// ## Example Usage
//
// ### Basic Usage
// ### Email Address Identity
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sesv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sesv2.NewEmailIdentity(ctx, "example", &sesv2.EmailIdentityArgs{
//				EmailIdentity: pulumi.String("testing@example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Domain Identity
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sesv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sesv2.NewEmailIdentity(ctx, "example", &sesv2.EmailIdentityArgs{
//				EmailIdentity: pulumi.String("example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Configuration Set
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sesv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleConfigurationSet, err := sesv2.NewConfigurationSet(ctx, "exampleConfigurationSet", &sesv2.ConfigurationSetArgs{
//				ConfigurationSetName: pulumi.String("example"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = sesv2.NewEmailIdentity(ctx, "exampleEmailIdentity", &sesv2.EmailIdentityArgs{
//				EmailIdentity:        pulumi.String("example.com"),
//				ConfigurationSetName: exampleConfigurationSet.ConfigurationSetName,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### DKIM Signing Attributes (BYODKIM)
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sesv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sesv2.NewEmailIdentity(ctx, "example", &sesv2.EmailIdentityArgs{
//				DkimSigningAttributes: &sesv2.EmailIdentityDkimSigningAttributesArgs{
//					DomainSigningPrivateKey: pulumi.String("MIIJKAIBAAKCAgEA2Se7p8zvnI4yh+Gh9j2rG5e2aRXjg03Y8saiupLnadPH9xvM..."),
//					DomainSigningSelector:   pulumi.String("example"),
//				},
//				EmailIdentity: pulumi.String("example.com"),
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
// terraform import {
//
//	to = aws_sesv2_email_identity.example
//
//	id = "example.com" } Using `pulumi import`, import SESv2 (Simple Email V2) Email Identity using the `email_identity`. For exampleconsole % pulumi import aws_sesv2_email_identity.example example.com
type EmailIdentity struct {
	pulumi.CustomResourceState

	// ARN of the Email Identity.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The configuration set to use by default when sending from this identity. Note that any configuration set defined in the email sending request takes precedence.
	ConfigurationSetName pulumi.StringPtrOutput `pulumi:"configurationSetName"`
	// The configuration of the DKIM authentication settings for an email domain identity.
	DkimSigningAttributes EmailIdentityDkimSigningAttributesOutput `pulumi:"dkimSigningAttributes"`
	// The email address or domain to verify.
	//
	// The following arguments are optional:
	EmailIdentity pulumi.StringOutput `pulumi:"emailIdentity"`
	// The email identity type. Valid values: `EMAIL_ADDRESS`, `DOMAIN`.
	IdentityType pulumi.StringOutput `pulumi:"identityType"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Specifies whether or not the identity is verified.
	VerifiedForSendingStatus pulumi.BoolOutput `pulumi:"verifiedForSendingStatus"`
}

// NewEmailIdentity registers a new resource with the given unique name, arguments, and options.
func NewEmailIdentity(ctx *pulumi.Context,
	name string, args *EmailIdentityArgs, opts ...pulumi.ResourceOption) (*EmailIdentity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EmailIdentity == nil {
		return nil, errors.New("invalid value for required argument 'EmailIdentity'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EmailIdentity
	err := ctx.RegisterResource("aws:sesv2/emailIdentity:EmailIdentity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEmailIdentity gets an existing EmailIdentity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEmailIdentity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmailIdentityState, opts ...pulumi.ResourceOption) (*EmailIdentity, error) {
	var resource EmailIdentity
	err := ctx.ReadResource("aws:sesv2/emailIdentity:EmailIdentity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EmailIdentity resources.
type emailIdentityState struct {
	// ARN of the Email Identity.
	Arn *string `pulumi:"arn"`
	// The configuration set to use by default when sending from this identity. Note that any configuration set defined in the email sending request takes precedence.
	ConfigurationSetName *string `pulumi:"configurationSetName"`
	// The configuration of the DKIM authentication settings for an email domain identity.
	DkimSigningAttributes *EmailIdentityDkimSigningAttributes `pulumi:"dkimSigningAttributes"`
	// The email address or domain to verify.
	//
	// The following arguments are optional:
	EmailIdentity *string `pulumi:"emailIdentity"`
	// The email identity type. Valid values: `EMAIL_ADDRESS`, `DOMAIN`.
	IdentityType *string `pulumi:"identityType"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Specifies whether or not the identity is verified.
	VerifiedForSendingStatus *bool `pulumi:"verifiedForSendingStatus"`
}

type EmailIdentityState struct {
	// ARN of the Email Identity.
	Arn pulumi.StringPtrInput
	// The configuration set to use by default when sending from this identity. Note that any configuration set defined in the email sending request takes precedence.
	ConfigurationSetName pulumi.StringPtrInput
	// The configuration of the DKIM authentication settings for an email domain identity.
	DkimSigningAttributes EmailIdentityDkimSigningAttributesPtrInput
	// The email address or domain to verify.
	//
	// The following arguments are optional:
	EmailIdentity pulumi.StringPtrInput
	// The email identity type. Valid values: `EMAIL_ADDRESS`, `DOMAIN`.
	IdentityType pulumi.StringPtrInput
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// Specifies whether or not the identity is verified.
	VerifiedForSendingStatus pulumi.BoolPtrInput
}

func (EmailIdentityState) ElementType() reflect.Type {
	return reflect.TypeOf((*emailIdentityState)(nil)).Elem()
}

type emailIdentityArgs struct {
	// The configuration set to use by default when sending from this identity. Note that any configuration set defined in the email sending request takes precedence.
	ConfigurationSetName *string `pulumi:"configurationSetName"`
	// The configuration of the DKIM authentication settings for an email domain identity.
	DkimSigningAttributes *EmailIdentityDkimSigningAttributes `pulumi:"dkimSigningAttributes"`
	// The email address or domain to verify.
	//
	// The following arguments are optional:
	EmailIdentity string `pulumi:"emailIdentity"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a EmailIdentity resource.
type EmailIdentityArgs struct {
	// The configuration set to use by default when sending from this identity. Note that any configuration set defined in the email sending request takes precedence.
	ConfigurationSetName pulumi.StringPtrInput
	// The configuration of the DKIM authentication settings for an email domain identity.
	DkimSigningAttributes EmailIdentityDkimSigningAttributesPtrInput
	// The email address or domain to verify.
	//
	// The following arguments are optional:
	EmailIdentity pulumi.StringInput
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (EmailIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*emailIdentityArgs)(nil)).Elem()
}

type EmailIdentityInput interface {
	pulumi.Input

	ToEmailIdentityOutput() EmailIdentityOutput
	ToEmailIdentityOutputWithContext(ctx context.Context) EmailIdentityOutput
}

func (*EmailIdentity) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailIdentity)(nil)).Elem()
}

func (i *EmailIdentity) ToEmailIdentityOutput() EmailIdentityOutput {
	return i.ToEmailIdentityOutputWithContext(context.Background())
}

func (i *EmailIdentity) ToEmailIdentityOutputWithContext(ctx context.Context) EmailIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailIdentityOutput)
}

// EmailIdentityArrayInput is an input type that accepts EmailIdentityArray and EmailIdentityArrayOutput values.
// You can construct a concrete instance of `EmailIdentityArrayInput` via:
//
//	EmailIdentityArray{ EmailIdentityArgs{...} }
type EmailIdentityArrayInput interface {
	pulumi.Input

	ToEmailIdentityArrayOutput() EmailIdentityArrayOutput
	ToEmailIdentityArrayOutputWithContext(context.Context) EmailIdentityArrayOutput
}

type EmailIdentityArray []EmailIdentityInput

func (EmailIdentityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EmailIdentity)(nil)).Elem()
}

func (i EmailIdentityArray) ToEmailIdentityArrayOutput() EmailIdentityArrayOutput {
	return i.ToEmailIdentityArrayOutputWithContext(context.Background())
}

func (i EmailIdentityArray) ToEmailIdentityArrayOutputWithContext(ctx context.Context) EmailIdentityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailIdentityArrayOutput)
}

// EmailIdentityMapInput is an input type that accepts EmailIdentityMap and EmailIdentityMapOutput values.
// You can construct a concrete instance of `EmailIdentityMapInput` via:
//
//	EmailIdentityMap{ "key": EmailIdentityArgs{...} }
type EmailIdentityMapInput interface {
	pulumi.Input

	ToEmailIdentityMapOutput() EmailIdentityMapOutput
	ToEmailIdentityMapOutputWithContext(context.Context) EmailIdentityMapOutput
}

type EmailIdentityMap map[string]EmailIdentityInput

func (EmailIdentityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EmailIdentity)(nil)).Elem()
}

func (i EmailIdentityMap) ToEmailIdentityMapOutput() EmailIdentityMapOutput {
	return i.ToEmailIdentityMapOutputWithContext(context.Background())
}

func (i EmailIdentityMap) ToEmailIdentityMapOutputWithContext(ctx context.Context) EmailIdentityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailIdentityMapOutput)
}

type EmailIdentityOutput struct{ *pulumi.OutputState }

func (EmailIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailIdentity)(nil)).Elem()
}

func (o EmailIdentityOutput) ToEmailIdentityOutput() EmailIdentityOutput {
	return o
}

func (o EmailIdentityOutput) ToEmailIdentityOutputWithContext(ctx context.Context) EmailIdentityOutput {
	return o
}

// ARN of the Email Identity.
func (o EmailIdentityOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailIdentity) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The configuration set to use by default when sending from this identity. Note that any configuration set defined in the email sending request takes precedence.
func (o EmailIdentityOutput) ConfigurationSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmailIdentity) pulumi.StringPtrOutput { return v.ConfigurationSetName }).(pulumi.StringPtrOutput)
}

// The configuration of the DKIM authentication settings for an email domain identity.
func (o EmailIdentityOutput) DkimSigningAttributes() EmailIdentityDkimSigningAttributesOutput {
	return o.ApplyT(func(v *EmailIdentity) EmailIdentityDkimSigningAttributesOutput { return v.DkimSigningAttributes }).(EmailIdentityDkimSigningAttributesOutput)
}

// The email address or domain to verify.
//
// The following arguments are optional:
func (o EmailIdentityOutput) EmailIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailIdentity) pulumi.StringOutput { return v.EmailIdentity }).(pulumi.StringOutput)
}

// The email identity type. Valid values: `EMAIL_ADDRESS`, `DOMAIN`.
func (o EmailIdentityOutput) IdentityType() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailIdentity) pulumi.StringOutput { return v.IdentityType }).(pulumi.StringOutput)
}

// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o EmailIdentityOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EmailIdentity) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o EmailIdentityOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EmailIdentity) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Specifies whether or not the identity is verified.
func (o EmailIdentityOutput) VerifiedForSendingStatus() pulumi.BoolOutput {
	return o.ApplyT(func(v *EmailIdentity) pulumi.BoolOutput { return v.VerifiedForSendingStatus }).(pulumi.BoolOutput)
}

type EmailIdentityArrayOutput struct{ *pulumi.OutputState }

func (EmailIdentityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EmailIdentity)(nil)).Elem()
}

func (o EmailIdentityArrayOutput) ToEmailIdentityArrayOutput() EmailIdentityArrayOutput {
	return o
}

func (o EmailIdentityArrayOutput) ToEmailIdentityArrayOutputWithContext(ctx context.Context) EmailIdentityArrayOutput {
	return o
}

func (o EmailIdentityArrayOutput) Index(i pulumi.IntInput) EmailIdentityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EmailIdentity {
		return vs[0].([]*EmailIdentity)[vs[1].(int)]
	}).(EmailIdentityOutput)
}

type EmailIdentityMapOutput struct{ *pulumi.OutputState }

func (EmailIdentityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EmailIdentity)(nil)).Elem()
}

func (o EmailIdentityMapOutput) ToEmailIdentityMapOutput() EmailIdentityMapOutput {
	return o
}

func (o EmailIdentityMapOutput) ToEmailIdentityMapOutputWithContext(ctx context.Context) EmailIdentityMapOutput {
	return o
}

func (o EmailIdentityMapOutput) MapIndex(k pulumi.StringInput) EmailIdentityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EmailIdentity {
		return vs[0].(map[string]*EmailIdentity)[vs[1].(string)]
	}).(EmailIdentityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EmailIdentityInput)(nil)).Elem(), &EmailIdentity{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailIdentityArrayInput)(nil)).Elem(), EmailIdentityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailIdentityMapInput)(nil)).Elem(), EmailIdentityMap{})
	pulumi.RegisterOutputType(EmailIdentityOutput{})
	pulumi.RegisterOutputType(EmailIdentityArrayOutput{})
	pulumi.RegisterOutputType(EmailIdentityMapOutput{})
}
