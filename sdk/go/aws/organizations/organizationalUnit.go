// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create an organizational unit.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := organizations.NewOrganizationalUnit(ctx, "example", &organizations.OrganizationalUnitArgs{
//				ParentId: pulumi.Any(aws_organizations_organization.Example.Roots[0].Id),
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
// AWS Organizations Organizational Units can be imported by using the `id`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:organizations/organizationalUnit:OrganizationalUnit example ou-1234567
//
// ```
type OrganizationalUnit struct {
	pulumi.CustomResourceState

	// List of child accounts for this Organizational Unit. Does not return account information for child Organizational Units. All elements have these attributes:
	Accounts OrganizationalUnitAccountArrayOutput `pulumi:"accounts"`
	// ARN of the organizational unit
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name for the organizational unit
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the parent organizational unit, which may be the root
	ParentId pulumi.StringOutput `pulumi:"parentId"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewOrganizationalUnit registers a new resource with the given unique name, arguments, and options.
func NewOrganizationalUnit(ctx *pulumi.Context,
	name string, args *OrganizationalUnitArgs, opts ...pulumi.ResourceOption) (*OrganizationalUnit, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ParentId == nil {
		return nil, errors.New("invalid value for required argument 'ParentId'")
	}
	var resource OrganizationalUnit
	err := ctx.RegisterResource("aws:organizations/organizationalUnit:OrganizationalUnit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationalUnit gets an existing OrganizationalUnit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationalUnit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationalUnitState, opts ...pulumi.ResourceOption) (*OrganizationalUnit, error) {
	var resource OrganizationalUnit
	err := ctx.ReadResource("aws:organizations/organizationalUnit:OrganizationalUnit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationalUnit resources.
type organizationalUnitState struct {
	// List of child accounts for this Organizational Unit. Does not return account information for child Organizational Units. All elements have these attributes:
	Accounts []OrganizationalUnitAccount `pulumi:"accounts"`
	// ARN of the organizational unit
	Arn *string `pulumi:"arn"`
	// The name for the organizational unit
	Name *string `pulumi:"name"`
	// ID of the parent organizational unit, which may be the root
	ParentId *string `pulumi:"parentId"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type OrganizationalUnitState struct {
	// List of child accounts for this Organizational Unit. Does not return account information for child Organizational Units. All elements have these attributes:
	Accounts OrganizationalUnitAccountArrayInput
	// ARN of the organizational unit
	Arn pulumi.StringPtrInput
	// The name for the organizational unit
	Name pulumi.StringPtrInput
	// ID of the parent organizational unit, which may be the root
	ParentId pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (OrganizationalUnitState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationalUnitState)(nil)).Elem()
}

type organizationalUnitArgs struct {
	// The name for the organizational unit
	Name *string `pulumi:"name"`
	// ID of the parent organizational unit, which may be the root
	ParentId string `pulumi:"parentId"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a OrganizationalUnit resource.
type OrganizationalUnitArgs struct {
	// The name for the organizational unit
	Name pulumi.StringPtrInput
	// ID of the parent organizational unit, which may be the root
	ParentId pulumi.StringInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (OrganizationalUnitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationalUnitArgs)(nil)).Elem()
}

type OrganizationalUnitInput interface {
	pulumi.Input

	ToOrganizationalUnitOutput() OrganizationalUnitOutput
	ToOrganizationalUnitOutputWithContext(ctx context.Context) OrganizationalUnitOutput
}

func (*OrganizationalUnit) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationalUnit)(nil)).Elem()
}

func (i *OrganizationalUnit) ToOrganizationalUnitOutput() OrganizationalUnitOutput {
	return i.ToOrganizationalUnitOutputWithContext(context.Background())
}

func (i *OrganizationalUnit) ToOrganizationalUnitOutputWithContext(ctx context.Context) OrganizationalUnitOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationalUnitOutput)
}

// OrganizationalUnitArrayInput is an input type that accepts OrganizationalUnitArray and OrganizationalUnitArrayOutput values.
// You can construct a concrete instance of `OrganizationalUnitArrayInput` via:
//
//	OrganizationalUnitArray{ OrganizationalUnitArgs{...} }
type OrganizationalUnitArrayInput interface {
	pulumi.Input

	ToOrganizationalUnitArrayOutput() OrganizationalUnitArrayOutput
	ToOrganizationalUnitArrayOutputWithContext(context.Context) OrganizationalUnitArrayOutput
}

type OrganizationalUnitArray []OrganizationalUnitInput

func (OrganizationalUnitArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationalUnit)(nil)).Elem()
}

func (i OrganizationalUnitArray) ToOrganizationalUnitArrayOutput() OrganizationalUnitArrayOutput {
	return i.ToOrganizationalUnitArrayOutputWithContext(context.Background())
}

func (i OrganizationalUnitArray) ToOrganizationalUnitArrayOutputWithContext(ctx context.Context) OrganizationalUnitArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationalUnitArrayOutput)
}

// OrganizationalUnitMapInput is an input type that accepts OrganizationalUnitMap and OrganizationalUnitMapOutput values.
// You can construct a concrete instance of `OrganizationalUnitMapInput` via:
//
//	OrganizationalUnitMap{ "key": OrganizationalUnitArgs{...} }
type OrganizationalUnitMapInput interface {
	pulumi.Input

	ToOrganizationalUnitMapOutput() OrganizationalUnitMapOutput
	ToOrganizationalUnitMapOutputWithContext(context.Context) OrganizationalUnitMapOutput
}

type OrganizationalUnitMap map[string]OrganizationalUnitInput

func (OrganizationalUnitMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationalUnit)(nil)).Elem()
}

func (i OrganizationalUnitMap) ToOrganizationalUnitMapOutput() OrganizationalUnitMapOutput {
	return i.ToOrganizationalUnitMapOutputWithContext(context.Background())
}

func (i OrganizationalUnitMap) ToOrganizationalUnitMapOutputWithContext(ctx context.Context) OrganizationalUnitMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationalUnitMapOutput)
}

type OrganizationalUnitOutput struct{ *pulumi.OutputState }

func (OrganizationalUnitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationalUnit)(nil)).Elem()
}

func (o OrganizationalUnitOutput) ToOrganizationalUnitOutput() OrganizationalUnitOutput {
	return o
}

func (o OrganizationalUnitOutput) ToOrganizationalUnitOutputWithContext(ctx context.Context) OrganizationalUnitOutput {
	return o
}

// List of child accounts for this Organizational Unit. Does not return account information for child Organizational Units. All elements have these attributes:
func (o OrganizationalUnitOutput) Accounts() OrganizationalUnitAccountArrayOutput {
	return o.ApplyT(func(v *OrganizationalUnit) OrganizationalUnitAccountArrayOutput { return v.Accounts }).(OrganizationalUnitAccountArrayOutput)
}

// ARN of the organizational unit
func (o OrganizationalUnitOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationalUnit) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The name for the organizational unit
func (o OrganizationalUnitOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationalUnit) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of the parent organizational unit, which may be the root
func (o OrganizationalUnitOutput) ParentId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationalUnit) pulumi.StringOutput { return v.ParentId }).(pulumi.StringOutput)
}

// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o OrganizationalUnitOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OrganizationalUnit) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o OrganizationalUnitOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OrganizationalUnit) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type OrganizationalUnitArrayOutput struct{ *pulumi.OutputState }

func (OrganizationalUnitArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationalUnit)(nil)).Elem()
}

func (o OrganizationalUnitArrayOutput) ToOrganizationalUnitArrayOutput() OrganizationalUnitArrayOutput {
	return o
}

func (o OrganizationalUnitArrayOutput) ToOrganizationalUnitArrayOutputWithContext(ctx context.Context) OrganizationalUnitArrayOutput {
	return o
}

func (o OrganizationalUnitArrayOutput) Index(i pulumi.IntInput) OrganizationalUnitOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrganizationalUnit {
		return vs[0].([]*OrganizationalUnit)[vs[1].(int)]
	}).(OrganizationalUnitOutput)
}

type OrganizationalUnitMapOutput struct{ *pulumi.OutputState }

func (OrganizationalUnitMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationalUnit)(nil)).Elem()
}

func (o OrganizationalUnitMapOutput) ToOrganizationalUnitMapOutput() OrganizationalUnitMapOutput {
	return o
}

func (o OrganizationalUnitMapOutput) ToOrganizationalUnitMapOutputWithContext(ctx context.Context) OrganizationalUnitMapOutput {
	return o
}

func (o OrganizationalUnitMapOutput) MapIndex(k pulumi.StringInput) OrganizationalUnitOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrganizationalUnit {
		return vs[0].(map[string]*OrganizationalUnit)[vs[1].(string)]
	}).(OrganizationalUnitOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationalUnitInput)(nil)).Elem(), &OrganizationalUnit{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationalUnitArrayInput)(nil)).Elem(), OrganizationalUnitArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationalUnitMapInput)(nil)).Elem(), OrganizationalUnitMap{})
	pulumi.RegisterOutputType(OrganizationalUnitOutput{})
	pulumi.RegisterOutputType(OrganizationalUnitArrayOutput{})
	pulumi.RegisterOutputType(OrganizationalUnitMapOutput{})
}
