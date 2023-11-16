// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package directoryservice

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a trust relationship between two Active Directory Directories.
//
// The directories may either be both AWS Managed Microsoft AD domains or an AWS Managed Microsoft AD domain and a self-managed Active Directory Domain.
//
// The Trust relationship must be configured on both sides of the relationship.
// If a Trust has only been created on one side, it will be in the state `VerifyFailed`.
// Once the second Trust is created, the first will update to the correct state.
//
// ## Example Usage
// ### Two-Way Trust
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directoryservice"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			oneDirectory, err := directoryservice.NewDirectory(ctx, "oneDirectory", &directoryservice.DirectoryArgs{
//				Name: pulumi.String("one.example.com"),
//				Type: pulumi.String("MicrosoftAD"),
//			})
//			if err != nil {
//				return err
//			}
//			twoDirectory, err := directoryservice.NewDirectory(ctx, "twoDirectory", &directoryservice.DirectoryArgs{
//				Name: pulumi.String("two.example.com"),
//				Type: pulumi.String("MicrosoftAD"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = directoryservice.NewTrust(ctx, "oneTrust", &directoryservice.TrustArgs{
//				DirectoryId:                 oneDirectory.ID(),
//				RemoteDomainName:            twoDirectory.Name,
//				TrustDirection:              pulumi.String("Two-Way"),
//				TrustPassword:               pulumi.String("Some0therPassword"),
//				ConditionalForwarderIpAddrs: twoDirectory.DnsIpAddresses,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = directoryservice.NewTrust(ctx, "twoTrust", &directoryservice.TrustArgs{
//				DirectoryId:                 twoDirectory.ID(),
//				RemoteDomainName:            oneDirectory.Name,
//				TrustDirection:              pulumi.String("Two-Way"),
//				TrustPassword:               pulumi.String("Some0therPassword"),
//				ConditionalForwarderIpAddrs: oneDirectory.DnsIpAddresses,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### One-Way Trust
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directoryservice"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			oneDirectory, err := directoryservice.NewDirectory(ctx, "oneDirectory", &directoryservice.DirectoryArgs{
//				Name: pulumi.String("one.example.com"),
//				Type: pulumi.String("MicrosoftAD"),
//			})
//			if err != nil {
//				return err
//			}
//			twoDirectory, err := directoryservice.NewDirectory(ctx, "twoDirectory", &directoryservice.DirectoryArgs{
//				Name: pulumi.String("two.example.com"),
//				Type: pulumi.String("MicrosoftAD"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = directoryservice.NewTrust(ctx, "oneTrust", &directoryservice.TrustArgs{
//				DirectoryId:                 oneDirectory.ID(),
//				RemoteDomainName:            twoDirectory.Name,
//				TrustDirection:              pulumi.String("One-Way: Incoming"),
//				TrustPassword:               pulumi.String("Some0therPassword"),
//				ConditionalForwarderIpAddrs: twoDirectory.DnsIpAddresses,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = directoryservice.NewTrust(ctx, "twoTrust", &directoryservice.TrustArgs{
//				DirectoryId:                 twoDirectory.ID(),
//				RemoteDomainName:            oneDirectory.Name,
//				TrustDirection:              pulumi.String("One-Way: Outgoing"),
//				TrustPassword:               pulumi.String("Some0therPassword"),
//				ConditionalForwarderIpAddrs: oneDirectory.DnsIpAddresses,
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
// Using `pulumi import`, import the Trust relationship using the directory ID and remote domain name, separated by a `/`. For example:
//
// ```sh
//
//	$ pulumi import aws:directoryservice/trust:Trust example d-926724cf57/directory.example.com
//
// ```
type Trust struct {
	pulumi.CustomResourceState

	// Set of IPv4 addresses for the DNS server associated with the remote Directory.
	// Can contain between 1 and 4 values.
	ConditionalForwarderIpAddrs pulumi.StringArrayOutput `pulumi:"conditionalForwarderIpAddrs"`
	// Date and time when the Trust was created.
	CreatedDateTime pulumi.StringOutput `pulumi:"createdDateTime"`
	// Whether to delete the conditional forwarder when deleting the Trust relationship.
	DeleteAssociatedConditionalForwarder pulumi.BoolOutput `pulumi:"deleteAssociatedConditionalForwarder"`
	// ID of the Directory.
	DirectoryId pulumi.StringOutput `pulumi:"directoryId"`
	// Date and time when the Trust was last updated.
	LastUpdatedDateTime pulumi.StringOutput `pulumi:"lastUpdatedDateTime"`
	// Fully qualified domain name of the remote Directory.
	RemoteDomainName pulumi.StringOutput `pulumi:"remoteDomainName"`
	// Whether to enable selective authentication.
	// Valid values are `Enabled` and `Disabled`.
	// Default value is `Disabled`.
	SelectiveAuth pulumi.StringOutput `pulumi:"selectiveAuth"`
	// Date and time when the Trust state in `trustState` was last updated.
	StateLastUpdatedDateTime pulumi.StringOutput `pulumi:"stateLastUpdatedDateTime"`
	// The direction of the Trust relationship.
	// Valid values are `One-Way: Outgoing`, `One-Way: Incoming`, and `Two-Way`.
	TrustDirection pulumi.StringOutput `pulumi:"trustDirection"`
	// Password for the Trust.
	// Does not need to match the passwords for either Directory.
	// Can contain upper- and lower-case letters, numbers, and punctuation characters.
	// May be up to 128 characters long.
	TrustPassword pulumi.StringOutput `pulumi:"trustPassword"`
	// State of the Trust relationship.
	// One of `Created`, `VerifyFailed`,`Verified`, `UpdateFailed`,`Updated`,`Deleted`, or `Failed`.
	TrustState pulumi.StringOutput `pulumi:"trustState"`
	// Reason for the Trust state set in `trustState`.
	TrustStateReason pulumi.StringOutput `pulumi:"trustStateReason"`
	// Type of the Trust relationship.
	// Valid values are `Forest` and `External`.
	// Default value is `Forest`.
	TrustType pulumi.StringOutput `pulumi:"trustType"`
}

// NewTrust registers a new resource with the given unique name, arguments, and options.
func NewTrust(ctx *pulumi.Context,
	name string, args *TrustArgs, opts ...pulumi.ResourceOption) (*Trust, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DirectoryId == nil {
		return nil, errors.New("invalid value for required argument 'DirectoryId'")
	}
	if args.RemoteDomainName == nil {
		return nil, errors.New("invalid value for required argument 'RemoteDomainName'")
	}
	if args.TrustDirection == nil {
		return nil, errors.New("invalid value for required argument 'TrustDirection'")
	}
	if args.TrustPassword == nil {
		return nil, errors.New("invalid value for required argument 'TrustPassword'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Trust
	err := ctx.RegisterResource("aws:directoryservice/trust:Trust", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrust gets an existing Trust resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrust(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrustState, opts ...pulumi.ResourceOption) (*Trust, error) {
	var resource Trust
	err := ctx.ReadResource("aws:directoryservice/trust:Trust", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Trust resources.
type trustState struct {
	// Set of IPv4 addresses for the DNS server associated with the remote Directory.
	// Can contain between 1 and 4 values.
	ConditionalForwarderIpAddrs []string `pulumi:"conditionalForwarderIpAddrs"`
	// Date and time when the Trust was created.
	CreatedDateTime *string `pulumi:"createdDateTime"`
	// Whether to delete the conditional forwarder when deleting the Trust relationship.
	DeleteAssociatedConditionalForwarder *bool `pulumi:"deleteAssociatedConditionalForwarder"`
	// ID of the Directory.
	DirectoryId *string `pulumi:"directoryId"`
	// Date and time when the Trust was last updated.
	LastUpdatedDateTime *string `pulumi:"lastUpdatedDateTime"`
	// Fully qualified domain name of the remote Directory.
	RemoteDomainName *string `pulumi:"remoteDomainName"`
	// Whether to enable selective authentication.
	// Valid values are `Enabled` and `Disabled`.
	// Default value is `Disabled`.
	SelectiveAuth *string `pulumi:"selectiveAuth"`
	// Date and time when the Trust state in `trustState` was last updated.
	StateLastUpdatedDateTime *string `pulumi:"stateLastUpdatedDateTime"`
	// The direction of the Trust relationship.
	// Valid values are `One-Way: Outgoing`, `One-Way: Incoming`, and `Two-Way`.
	TrustDirection *string `pulumi:"trustDirection"`
	// Password for the Trust.
	// Does not need to match the passwords for either Directory.
	// Can contain upper- and lower-case letters, numbers, and punctuation characters.
	// May be up to 128 characters long.
	TrustPassword *string `pulumi:"trustPassword"`
	// State of the Trust relationship.
	// One of `Created`, `VerifyFailed`,`Verified`, `UpdateFailed`,`Updated`,`Deleted`, or `Failed`.
	TrustState *string `pulumi:"trustState"`
	// Reason for the Trust state set in `trustState`.
	TrustStateReason *string `pulumi:"trustStateReason"`
	// Type of the Trust relationship.
	// Valid values are `Forest` and `External`.
	// Default value is `Forest`.
	TrustType *string `pulumi:"trustType"`
}

type TrustState struct {
	// Set of IPv4 addresses for the DNS server associated with the remote Directory.
	// Can contain between 1 and 4 values.
	ConditionalForwarderIpAddrs pulumi.StringArrayInput
	// Date and time when the Trust was created.
	CreatedDateTime pulumi.StringPtrInput
	// Whether to delete the conditional forwarder when deleting the Trust relationship.
	DeleteAssociatedConditionalForwarder pulumi.BoolPtrInput
	// ID of the Directory.
	DirectoryId pulumi.StringPtrInput
	// Date and time when the Trust was last updated.
	LastUpdatedDateTime pulumi.StringPtrInput
	// Fully qualified domain name of the remote Directory.
	RemoteDomainName pulumi.StringPtrInput
	// Whether to enable selective authentication.
	// Valid values are `Enabled` and `Disabled`.
	// Default value is `Disabled`.
	SelectiveAuth pulumi.StringPtrInput
	// Date and time when the Trust state in `trustState` was last updated.
	StateLastUpdatedDateTime pulumi.StringPtrInput
	// The direction of the Trust relationship.
	// Valid values are `One-Way: Outgoing`, `One-Way: Incoming`, and `Two-Way`.
	TrustDirection pulumi.StringPtrInput
	// Password for the Trust.
	// Does not need to match the passwords for either Directory.
	// Can contain upper- and lower-case letters, numbers, and punctuation characters.
	// May be up to 128 characters long.
	TrustPassword pulumi.StringPtrInput
	// State of the Trust relationship.
	// One of `Created`, `VerifyFailed`,`Verified`, `UpdateFailed`,`Updated`,`Deleted`, or `Failed`.
	TrustState pulumi.StringPtrInput
	// Reason for the Trust state set in `trustState`.
	TrustStateReason pulumi.StringPtrInput
	// Type of the Trust relationship.
	// Valid values are `Forest` and `External`.
	// Default value is `Forest`.
	TrustType pulumi.StringPtrInput
}

func (TrustState) ElementType() reflect.Type {
	return reflect.TypeOf((*trustState)(nil)).Elem()
}

type trustArgs struct {
	// Set of IPv4 addresses for the DNS server associated with the remote Directory.
	// Can contain between 1 and 4 values.
	ConditionalForwarderIpAddrs []string `pulumi:"conditionalForwarderIpAddrs"`
	// Whether to delete the conditional forwarder when deleting the Trust relationship.
	DeleteAssociatedConditionalForwarder *bool `pulumi:"deleteAssociatedConditionalForwarder"`
	// ID of the Directory.
	DirectoryId string `pulumi:"directoryId"`
	// Fully qualified domain name of the remote Directory.
	RemoteDomainName string `pulumi:"remoteDomainName"`
	// Whether to enable selective authentication.
	// Valid values are `Enabled` and `Disabled`.
	// Default value is `Disabled`.
	SelectiveAuth *string `pulumi:"selectiveAuth"`
	// The direction of the Trust relationship.
	// Valid values are `One-Way: Outgoing`, `One-Way: Incoming`, and `Two-Way`.
	TrustDirection string `pulumi:"trustDirection"`
	// Password for the Trust.
	// Does not need to match the passwords for either Directory.
	// Can contain upper- and lower-case letters, numbers, and punctuation characters.
	// May be up to 128 characters long.
	TrustPassword string `pulumi:"trustPassword"`
	// Type of the Trust relationship.
	// Valid values are `Forest` and `External`.
	// Default value is `Forest`.
	TrustType *string `pulumi:"trustType"`
}

// The set of arguments for constructing a Trust resource.
type TrustArgs struct {
	// Set of IPv4 addresses for the DNS server associated with the remote Directory.
	// Can contain between 1 and 4 values.
	ConditionalForwarderIpAddrs pulumi.StringArrayInput
	// Whether to delete the conditional forwarder when deleting the Trust relationship.
	DeleteAssociatedConditionalForwarder pulumi.BoolPtrInput
	// ID of the Directory.
	DirectoryId pulumi.StringInput
	// Fully qualified domain name of the remote Directory.
	RemoteDomainName pulumi.StringInput
	// Whether to enable selective authentication.
	// Valid values are `Enabled` and `Disabled`.
	// Default value is `Disabled`.
	SelectiveAuth pulumi.StringPtrInput
	// The direction of the Trust relationship.
	// Valid values are `One-Way: Outgoing`, `One-Way: Incoming`, and `Two-Way`.
	TrustDirection pulumi.StringInput
	// Password for the Trust.
	// Does not need to match the passwords for either Directory.
	// Can contain upper- and lower-case letters, numbers, and punctuation characters.
	// May be up to 128 characters long.
	TrustPassword pulumi.StringInput
	// Type of the Trust relationship.
	// Valid values are `Forest` and `External`.
	// Default value is `Forest`.
	TrustType pulumi.StringPtrInput
}

func (TrustArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trustArgs)(nil)).Elem()
}

type TrustInput interface {
	pulumi.Input

	ToTrustOutput() TrustOutput
	ToTrustOutputWithContext(ctx context.Context) TrustOutput
}

func (*Trust) ElementType() reflect.Type {
	return reflect.TypeOf((**Trust)(nil)).Elem()
}

func (i *Trust) ToTrustOutput() TrustOutput {
	return i.ToTrustOutputWithContext(context.Background())
}

func (i *Trust) ToTrustOutputWithContext(ctx context.Context) TrustOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustOutput)
}

// TrustArrayInput is an input type that accepts TrustArray and TrustArrayOutput values.
// You can construct a concrete instance of `TrustArrayInput` via:
//
//	TrustArray{ TrustArgs{...} }
type TrustArrayInput interface {
	pulumi.Input

	ToTrustArrayOutput() TrustArrayOutput
	ToTrustArrayOutputWithContext(context.Context) TrustArrayOutput
}

type TrustArray []TrustInput

func (TrustArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Trust)(nil)).Elem()
}

func (i TrustArray) ToTrustArrayOutput() TrustArrayOutput {
	return i.ToTrustArrayOutputWithContext(context.Background())
}

func (i TrustArray) ToTrustArrayOutputWithContext(ctx context.Context) TrustArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustArrayOutput)
}

// TrustMapInput is an input type that accepts TrustMap and TrustMapOutput values.
// You can construct a concrete instance of `TrustMapInput` via:
//
//	TrustMap{ "key": TrustArgs{...} }
type TrustMapInput interface {
	pulumi.Input

	ToTrustMapOutput() TrustMapOutput
	ToTrustMapOutputWithContext(context.Context) TrustMapOutput
}

type TrustMap map[string]TrustInput

func (TrustMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Trust)(nil)).Elem()
}

func (i TrustMap) ToTrustMapOutput() TrustMapOutput {
	return i.ToTrustMapOutputWithContext(context.Background())
}

func (i TrustMap) ToTrustMapOutputWithContext(ctx context.Context) TrustMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustMapOutput)
}

type TrustOutput struct{ *pulumi.OutputState }

func (TrustOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Trust)(nil)).Elem()
}

func (o TrustOutput) ToTrustOutput() TrustOutput {
	return o
}

func (o TrustOutput) ToTrustOutputWithContext(ctx context.Context) TrustOutput {
	return o
}

// Set of IPv4 addresses for the DNS server associated with the remote Directory.
// Can contain between 1 and 4 values.
func (o TrustOutput) ConditionalForwarderIpAddrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Trust) pulumi.StringArrayOutput { return v.ConditionalForwarderIpAddrs }).(pulumi.StringArrayOutput)
}

// Date and time when the Trust was created.
func (o TrustOutput) CreatedDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Trust) pulumi.StringOutput { return v.CreatedDateTime }).(pulumi.StringOutput)
}

// Whether to delete the conditional forwarder when deleting the Trust relationship.
func (o TrustOutput) DeleteAssociatedConditionalForwarder() pulumi.BoolOutput {
	return o.ApplyT(func(v *Trust) pulumi.BoolOutput { return v.DeleteAssociatedConditionalForwarder }).(pulumi.BoolOutput)
}

// ID of the Directory.
func (o TrustOutput) DirectoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *Trust) pulumi.StringOutput { return v.DirectoryId }).(pulumi.StringOutput)
}

// Date and time when the Trust was last updated.
func (o TrustOutput) LastUpdatedDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Trust) pulumi.StringOutput { return v.LastUpdatedDateTime }).(pulumi.StringOutput)
}

// Fully qualified domain name of the remote Directory.
func (o TrustOutput) RemoteDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Trust) pulumi.StringOutput { return v.RemoteDomainName }).(pulumi.StringOutput)
}

// Whether to enable selective authentication.
// Valid values are `Enabled` and `Disabled`.
// Default value is `Disabled`.
func (o TrustOutput) SelectiveAuth() pulumi.StringOutput {
	return o.ApplyT(func(v *Trust) pulumi.StringOutput { return v.SelectiveAuth }).(pulumi.StringOutput)
}

// Date and time when the Trust state in `trustState` was last updated.
func (o TrustOutput) StateLastUpdatedDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Trust) pulumi.StringOutput { return v.StateLastUpdatedDateTime }).(pulumi.StringOutput)
}

// The direction of the Trust relationship.
// Valid values are `One-Way: Outgoing`, `One-Way: Incoming`, and `Two-Way`.
func (o TrustOutput) TrustDirection() pulumi.StringOutput {
	return o.ApplyT(func(v *Trust) pulumi.StringOutput { return v.TrustDirection }).(pulumi.StringOutput)
}

// Password for the Trust.
// Does not need to match the passwords for either Directory.
// Can contain upper- and lower-case letters, numbers, and punctuation characters.
// May be up to 128 characters long.
func (o TrustOutput) TrustPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *Trust) pulumi.StringOutput { return v.TrustPassword }).(pulumi.StringOutput)
}

// State of the Trust relationship.
// One of `Created`, `VerifyFailed`,`Verified`, `UpdateFailed`,`Updated`,`Deleted`, or `Failed`.
func (o TrustOutput) TrustState() pulumi.StringOutput {
	return o.ApplyT(func(v *Trust) pulumi.StringOutput { return v.TrustState }).(pulumi.StringOutput)
}

// Reason for the Trust state set in `trustState`.
func (o TrustOutput) TrustStateReason() pulumi.StringOutput {
	return o.ApplyT(func(v *Trust) pulumi.StringOutput { return v.TrustStateReason }).(pulumi.StringOutput)
}

// Type of the Trust relationship.
// Valid values are `Forest` and `External`.
// Default value is `Forest`.
func (o TrustOutput) TrustType() pulumi.StringOutput {
	return o.ApplyT(func(v *Trust) pulumi.StringOutput { return v.TrustType }).(pulumi.StringOutput)
}

type TrustArrayOutput struct{ *pulumi.OutputState }

func (TrustArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Trust)(nil)).Elem()
}

func (o TrustArrayOutput) ToTrustArrayOutput() TrustArrayOutput {
	return o
}

func (o TrustArrayOutput) ToTrustArrayOutputWithContext(ctx context.Context) TrustArrayOutput {
	return o
}

func (o TrustArrayOutput) Index(i pulumi.IntInput) TrustOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Trust {
		return vs[0].([]*Trust)[vs[1].(int)]
	}).(TrustOutput)
}

type TrustMapOutput struct{ *pulumi.OutputState }

func (TrustMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Trust)(nil)).Elem()
}

func (o TrustMapOutput) ToTrustMapOutput() TrustMapOutput {
	return o
}

func (o TrustMapOutput) ToTrustMapOutputWithContext(ctx context.Context) TrustMapOutput {
	return o
}

func (o TrustMapOutput) MapIndex(k pulumi.StringInput) TrustOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Trust {
		return vs[0].(map[string]*Trust)[vs[1].(string)]
	}).(TrustOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrustInput)(nil)).Elem(), &Trust{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrustArrayInput)(nil)).Elem(), TrustArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrustMapInput)(nil)).Elem(), TrustMap{})
	pulumi.RegisterOutputType(TrustOutput{})
	pulumi.RegisterOutputType(TrustArrayOutput{})
	pulumi.RegisterOutputType(TrustMapOutput{})
}
