// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package connect

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Amazon Connect Phone Number resource. For more information see
// [Amazon Connect: Getting Started](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-get-started.html)
//
// ## Example Usage
// ### Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/connect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := connect.NewPhoneNumber(ctx, "example", &connect.PhoneNumberArgs{
//				TargetArn:   pulumi.Any(aws_connect_instance.Example.Arn),
//				CountryCode: pulumi.String("US"),
//				Type:        pulumi.String("DID"),
//				Tags: pulumi.StringMap{
//					"hello": pulumi.String("world"),
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
// ### Description
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/connect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := connect.NewPhoneNumber(ctx, "example", &connect.PhoneNumberArgs{
//				TargetArn:   pulumi.Any(aws_connect_instance.Example.Arn),
//				CountryCode: pulumi.String("US"),
//				Type:        pulumi.String("DID"),
//				Description: pulumi.String("example description"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Prefix to filter phone numbers
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/connect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := connect.NewPhoneNumber(ctx, "example", &connect.PhoneNumberArgs{
//				TargetArn:   pulumi.Any(aws_connect_instance.Example.Arn),
//				CountryCode: pulumi.String("US"),
//				Type:        pulumi.String("DID"),
//				Prefix:      pulumi.String("+18005"),
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
// Amazon Connect Phone Numbers can be imported using its `id` e.g.,
//
// ```sh
//
//	$ pulumi import aws:connect/phoneNumber:PhoneNumber example 12345678-abcd-1234-efgh-9876543210ab
//
// ```
type PhoneNumber struct {
	pulumi.CustomResourceState

	// The ARN of the phone number.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ISO country code. For a list of Valid values, refer to [PhoneNumberCountryCode](https://docs.aws.amazon.com/connect/latest/APIReference/API_SearchAvailablePhoneNumbers.html#connect-SearchAvailablePhoneNumbers-request-PhoneNumberCountryCode).
	CountryCode pulumi.StringOutput `pulumi:"countryCode"`
	// The description of the phone number.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The phone number. Phone numbers are formatted `[+] [country code] [subscriber number including area code]`.
	PhoneNumber pulumi.StringOutput `pulumi:"phoneNumber"`
	// The prefix of the phone number that is used to filter available phone numbers. If provided, it must contain `+` as part of the country code. Do not specify this argument when importing the resource.
	Prefix pulumi.StringPtrOutput `pulumi:"prefix"`
	// The status of the phone number. Valid Values: `CLAIMED` | `IN_PROGRESS` | `FAILED`.
	Statuses PhoneNumberStatusArrayOutput `pulumi:"statuses"`
	// Tags to apply to the Phone Number. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The Amazon Resource Name (ARN) for Amazon Connect instances that phone numbers are claimed to.
	TargetArn pulumi.StringOutput `pulumi:"targetArn"`
	// The type of phone number. Valid Values: `TOLL_FREE` | `DID`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPhoneNumber registers a new resource with the given unique name, arguments, and options.
func NewPhoneNumber(ctx *pulumi.Context,
	name string, args *PhoneNumberArgs, opts ...pulumi.ResourceOption) (*PhoneNumber, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CountryCode == nil {
		return nil, errors.New("invalid value for required argument 'CountryCode'")
	}
	if args.TargetArn == nil {
		return nil, errors.New("invalid value for required argument 'TargetArn'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource PhoneNumber
	err := ctx.RegisterResource("aws:connect/phoneNumber:PhoneNumber", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPhoneNumber gets an existing PhoneNumber resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPhoneNumber(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PhoneNumberState, opts ...pulumi.ResourceOption) (*PhoneNumber, error) {
	var resource PhoneNumber
	err := ctx.ReadResource("aws:connect/phoneNumber:PhoneNumber", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PhoneNumber resources.
type phoneNumberState struct {
	// The ARN of the phone number.
	Arn *string `pulumi:"arn"`
	// The ISO country code. For a list of Valid values, refer to [PhoneNumberCountryCode](https://docs.aws.amazon.com/connect/latest/APIReference/API_SearchAvailablePhoneNumbers.html#connect-SearchAvailablePhoneNumbers-request-PhoneNumberCountryCode).
	CountryCode *string `pulumi:"countryCode"`
	// The description of the phone number.
	Description *string `pulumi:"description"`
	// The phone number. Phone numbers are formatted `[+] [country code] [subscriber number including area code]`.
	PhoneNumber *string `pulumi:"phoneNumber"`
	// The prefix of the phone number that is used to filter available phone numbers. If provided, it must contain `+` as part of the country code. Do not specify this argument when importing the resource.
	Prefix *string `pulumi:"prefix"`
	// The status of the phone number. Valid Values: `CLAIMED` | `IN_PROGRESS` | `FAILED`.
	Statuses []PhoneNumberStatus `pulumi:"statuses"`
	// Tags to apply to the Phone Number. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The Amazon Resource Name (ARN) for Amazon Connect instances that phone numbers are claimed to.
	TargetArn *string `pulumi:"targetArn"`
	// The type of phone number. Valid Values: `TOLL_FREE` | `DID`.
	Type *string `pulumi:"type"`
}

type PhoneNumberState struct {
	// The ARN of the phone number.
	Arn pulumi.StringPtrInput
	// The ISO country code. For a list of Valid values, refer to [PhoneNumberCountryCode](https://docs.aws.amazon.com/connect/latest/APIReference/API_SearchAvailablePhoneNumbers.html#connect-SearchAvailablePhoneNumbers-request-PhoneNumberCountryCode).
	CountryCode pulumi.StringPtrInput
	// The description of the phone number.
	Description pulumi.StringPtrInput
	// The phone number. Phone numbers are formatted `[+] [country code] [subscriber number including area code]`.
	PhoneNumber pulumi.StringPtrInput
	// The prefix of the phone number that is used to filter available phone numbers. If provided, it must contain `+` as part of the country code. Do not specify this argument when importing the resource.
	Prefix pulumi.StringPtrInput
	// The status of the phone number. Valid Values: `CLAIMED` | `IN_PROGRESS` | `FAILED`.
	Statuses PhoneNumberStatusArrayInput
	// Tags to apply to the Phone Number. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// The Amazon Resource Name (ARN) for Amazon Connect instances that phone numbers are claimed to.
	TargetArn pulumi.StringPtrInput
	// The type of phone number. Valid Values: `TOLL_FREE` | `DID`.
	Type pulumi.StringPtrInput
}

func (PhoneNumberState) ElementType() reflect.Type {
	return reflect.TypeOf((*phoneNumberState)(nil)).Elem()
}

type phoneNumberArgs struct {
	// The ISO country code. For a list of Valid values, refer to [PhoneNumberCountryCode](https://docs.aws.amazon.com/connect/latest/APIReference/API_SearchAvailablePhoneNumbers.html#connect-SearchAvailablePhoneNumbers-request-PhoneNumberCountryCode).
	CountryCode string `pulumi:"countryCode"`
	// The description of the phone number.
	Description *string `pulumi:"description"`
	// The prefix of the phone number that is used to filter available phone numbers. If provided, it must contain `+` as part of the country code. Do not specify this argument when importing the resource.
	Prefix *string `pulumi:"prefix"`
	// Tags to apply to the Phone Number. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The Amazon Resource Name (ARN) for Amazon Connect instances that phone numbers are claimed to.
	TargetArn string `pulumi:"targetArn"`
	// The type of phone number. Valid Values: `TOLL_FREE` | `DID`.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a PhoneNumber resource.
type PhoneNumberArgs struct {
	// The ISO country code. For a list of Valid values, refer to [PhoneNumberCountryCode](https://docs.aws.amazon.com/connect/latest/APIReference/API_SearchAvailablePhoneNumbers.html#connect-SearchAvailablePhoneNumbers-request-PhoneNumberCountryCode).
	CountryCode pulumi.StringInput
	// The description of the phone number.
	Description pulumi.StringPtrInput
	// The prefix of the phone number that is used to filter available phone numbers. If provided, it must contain `+` as part of the country code. Do not specify this argument when importing the resource.
	Prefix pulumi.StringPtrInput
	// Tags to apply to the Phone Number. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The Amazon Resource Name (ARN) for Amazon Connect instances that phone numbers are claimed to.
	TargetArn pulumi.StringInput
	// The type of phone number. Valid Values: `TOLL_FREE` | `DID`.
	Type pulumi.StringInput
}

func (PhoneNumberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*phoneNumberArgs)(nil)).Elem()
}

type PhoneNumberInput interface {
	pulumi.Input

	ToPhoneNumberOutput() PhoneNumberOutput
	ToPhoneNumberOutputWithContext(ctx context.Context) PhoneNumberOutput
}

func (*PhoneNumber) ElementType() reflect.Type {
	return reflect.TypeOf((**PhoneNumber)(nil)).Elem()
}

func (i *PhoneNumber) ToPhoneNumberOutput() PhoneNumberOutput {
	return i.ToPhoneNumberOutputWithContext(context.Background())
}

func (i *PhoneNumber) ToPhoneNumberOutputWithContext(ctx context.Context) PhoneNumberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PhoneNumberOutput)
}

// PhoneNumberArrayInput is an input type that accepts PhoneNumberArray and PhoneNumberArrayOutput values.
// You can construct a concrete instance of `PhoneNumberArrayInput` via:
//
//	PhoneNumberArray{ PhoneNumberArgs{...} }
type PhoneNumberArrayInput interface {
	pulumi.Input

	ToPhoneNumberArrayOutput() PhoneNumberArrayOutput
	ToPhoneNumberArrayOutputWithContext(context.Context) PhoneNumberArrayOutput
}

type PhoneNumberArray []PhoneNumberInput

func (PhoneNumberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PhoneNumber)(nil)).Elem()
}

func (i PhoneNumberArray) ToPhoneNumberArrayOutput() PhoneNumberArrayOutput {
	return i.ToPhoneNumberArrayOutputWithContext(context.Background())
}

func (i PhoneNumberArray) ToPhoneNumberArrayOutputWithContext(ctx context.Context) PhoneNumberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PhoneNumberArrayOutput)
}

// PhoneNumberMapInput is an input type that accepts PhoneNumberMap and PhoneNumberMapOutput values.
// You can construct a concrete instance of `PhoneNumberMapInput` via:
//
//	PhoneNumberMap{ "key": PhoneNumberArgs{...} }
type PhoneNumberMapInput interface {
	pulumi.Input

	ToPhoneNumberMapOutput() PhoneNumberMapOutput
	ToPhoneNumberMapOutputWithContext(context.Context) PhoneNumberMapOutput
}

type PhoneNumberMap map[string]PhoneNumberInput

func (PhoneNumberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PhoneNumber)(nil)).Elem()
}

func (i PhoneNumberMap) ToPhoneNumberMapOutput() PhoneNumberMapOutput {
	return i.ToPhoneNumberMapOutputWithContext(context.Background())
}

func (i PhoneNumberMap) ToPhoneNumberMapOutputWithContext(ctx context.Context) PhoneNumberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PhoneNumberMapOutput)
}

type PhoneNumberOutput struct{ *pulumi.OutputState }

func (PhoneNumberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PhoneNumber)(nil)).Elem()
}

func (o PhoneNumberOutput) ToPhoneNumberOutput() PhoneNumberOutput {
	return o
}

func (o PhoneNumberOutput) ToPhoneNumberOutputWithContext(ctx context.Context) PhoneNumberOutput {
	return o
}

// The ARN of the phone number.
func (o PhoneNumberOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *PhoneNumber) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The ISO country code. For a list of Valid values, refer to [PhoneNumberCountryCode](https://docs.aws.amazon.com/connect/latest/APIReference/API_SearchAvailablePhoneNumbers.html#connect-SearchAvailablePhoneNumbers-request-PhoneNumberCountryCode).
func (o PhoneNumberOutput) CountryCode() pulumi.StringOutput {
	return o.ApplyT(func(v *PhoneNumber) pulumi.StringOutput { return v.CountryCode }).(pulumi.StringOutput)
}

// The description of the phone number.
func (o PhoneNumberOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PhoneNumber) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The phone number. Phone numbers are formatted `[+] [country code] [subscriber number including area code]`.
func (o PhoneNumberOutput) PhoneNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *PhoneNumber) pulumi.StringOutput { return v.PhoneNumber }).(pulumi.StringOutput)
}

// The prefix of the phone number that is used to filter available phone numbers. If provided, it must contain `+` as part of the country code. Do not specify this argument when importing the resource.
func (o PhoneNumberOutput) Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PhoneNumber) pulumi.StringPtrOutput { return v.Prefix }).(pulumi.StringPtrOutput)
}

// The status of the phone number. Valid Values: `CLAIMED` | `IN_PROGRESS` | `FAILED`.
func (o PhoneNumberOutput) Statuses() PhoneNumberStatusArrayOutput {
	return o.ApplyT(func(v *PhoneNumber) PhoneNumberStatusArrayOutput { return v.Statuses }).(PhoneNumberStatusArrayOutput)
}

// Tags to apply to the Phone Number. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o PhoneNumberOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PhoneNumber) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o PhoneNumberOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PhoneNumber) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The Amazon Resource Name (ARN) for Amazon Connect instances that phone numbers are claimed to.
func (o PhoneNumberOutput) TargetArn() pulumi.StringOutput {
	return o.ApplyT(func(v *PhoneNumber) pulumi.StringOutput { return v.TargetArn }).(pulumi.StringOutput)
}

// The type of phone number. Valid Values: `TOLL_FREE` | `DID`.
func (o PhoneNumberOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PhoneNumber) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type PhoneNumberArrayOutput struct{ *pulumi.OutputState }

func (PhoneNumberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PhoneNumber)(nil)).Elem()
}

func (o PhoneNumberArrayOutput) ToPhoneNumberArrayOutput() PhoneNumberArrayOutput {
	return o
}

func (o PhoneNumberArrayOutput) ToPhoneNumberArrayOutputWithContext(ctx context.Context) PhoneNumberArrayOutput {
	return o
}

func (o PhoneNumberArrayOutput) Index(i pulumi.IntInput) PhoneNumberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PhoneNumber {
		return vs[0].([]*PhoneNumber)[vs[1].(int)]
	}).(PhoneNumberOutput)
}

type PhoneNumberMapOutput struct{ *pulumi.OutputState }

func (PhoneNumberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PhoneNumber)(nil)).Elem()
}

func (o PhoneNumberMapOutput) ToPhoneNumberMapOutput() PhoneNumberMapOutput {
	return o
}

func (o PhoneNumberMapOutput) ToPhoneNumberMapOutputWithContext(ctx context.Context) PhoneNumberMapOutput {
	return o
}

func (o PhoneNumberMapOutput) MapIndex(k pulumi.StringInput) PhoneNumberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PhoneNumber {
		return vs[0].(map[string]*PhoneNumber)[vs[1].(string)]
	}).(PhoneNumberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PhoneNumberInput)(nil)).Elem(), &PhoneNumber{})
	pulumi.RegisterInputType(reflect.TypeOf((*PhoneNumberArrayInput)(nil)).Elem(), PhoneNumberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PhoneNumberMapInput)(nil)).Elem(), PhoneNumberMap{})
	pulumi.RegisterOutputType(PhoneNumberOutput{})
	pulumi.RegisterOutputType(PhoneNumberArrayOutput{})
	pulumi.RegisterOutputType(PhoneNumberMapOutput{})
}
