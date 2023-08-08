// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkmanager

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a site in a global network.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkmanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleGlobalNetwork, err := networkmanager.NewGlobalNetwork(ctx, "exampleGlobalNetwork", nil)
//			if err != nil {
//				return err
//			}
//			_, err = networkmanager.NewSite(ctx, "exampleSite", &networkmanager.SiteArgs{
//				GlobalNetworkId: exampleGlobalNetwork.ID(),
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
//	to = aws_networkmanager_site.example
//
//	id = "arn:aws:networkmanager::123456789012:site/global-network-0d47f6t230mz46dy4/site-444555aaabbb11223" } Using `pulumi import`, import `aws_networkmanager_site` using the site ARN. For exampleconsole % pulumi import aws_networkmanager_site.example arn:aws:networkmanager::123456789012:site/global-network-0d47f6t230mz46dy4/site-444555aaabbb11223
type Site struct {
	pulumi.CustomResourceState

	// Site Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description of the Site.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the Global Network to create the site in.
	GlobalNetworkId pulumi.StringOutput `pulumi:"globalNetworkId"`
	// The site location as documented below.
	Location SiteLocationPtrOutput `pulumi:"location"`
	// Key-value tags for the Site. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewSite registers a new resource with the given unique name, arguments, and options.
func NewSite(ctx *pulumi.Context,
	name string, args *SiteArgs, opts ...pulumi.ResourceOption) (*Site, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GlobalNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'GlobalNetworkId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Site
	err := ctx.RegisterResource("aws:networkmanager/site:Site", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSite gets an existing Site resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteState, opts ...pulumi.ResourceOption) (*Site, error) {
	var resource Site
	err := ctx.ReadResource("aws:networkmanager/site:Site", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Site resources.
type siteState struct {
	// Site Amazon Resource Name (ARN)
	Arn *string `pulumi:"arn"`
	// Description of the Site.
	Description *string `pulumi:"description"`
	// The ID of the Global Network to create the site in.
	GlobalNetworkId *string `pulumi:"globalNetworkId"`
	// The site location as documented below.
	Location *SiteLocation `pulumi:"location"`
	// Key-value tags for the Site. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type SiteState struct {
	// Site Amazon Resource Name (ARN)
	Arn pulumi.StringPtrInput
	// Description of the Site.
	Description pulumi.StringPtrInput
	// The ID of the Global Network to create the site in.
	GlobalNetworkId pulumi.StringPtrInput
	// The site location as documented below.
	Location SiteLocationPtrInput
	// Key-value tags for the Site. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (SiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteState)(nil)).Elem()
}

type siteArgs struct {
	// Description of the Site.
	Description *string `pulumi:"description"`
	// The ID of the Global Network to create the site in.
	GlobalNetworkId string `pulumi:"globalNetworkId"`
	// The site location as documented below.
	Location *SiteLocation `pulumi:"location"`
	// Key-value tags for the Site. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Site resource.
type SiteArgs struct {
	// Description of the Site.
	Description pulumi.StringPtrInput
	// The ID of the Global Network to create the site in.
	GlobalNetworkId pulumi.StringInput
	// The site location as documented below.
	Location SiteLocationPtrInput
	// Key-value tags for the Site. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (SiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteArgs)(nil)).Elem()
}

type SiteInput interface {
	pulumi.Input

	ToSiteOutput() SiteOutput
	ToSiteOutputWithContext(ctx context.Context) SiteOutput
}

func (*Site) ElementType() reflect.Type {
	return reflect.TypeOf((**Site)(nil)).Elem()
}

func (i *Site) ToSiteOutput() SiteOutput {
	return i.ToSiteOutputWithContext(context.Background())
}

func (i *Site) ToSiteOutputWithContext(ctx context.Context) SiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteOutput)
}

// SiteArrayInput is an input type that accepts SiteArray and SiteArrayOutput values.
// You can construct a concrete instance of `SiteArrayInput` via:
//
//	SiteArray{ SiteArgs{...} }
type SiteArrayInput interface {
	pulumi.Input

	ToSiteArrayOutput() SiteArrayOutput
	ToSiteArrayOutputWithContext(context.Context) SiteArrayOutput
}

type SiteArray []SiteInput

func (SiteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Site)(nil)).Elem()
}

func (i SiteArray) ToSiteArrayOutput() SiteArrayOutput {
	return i.ToSiteArrayOutputWithContext(context.Background())
}

func (i SiteArray) ToSiteArrayOutputWithContext(ctx context.Context) SiteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteArrayOutput)
}

// SiteMapInput is an input type that accepts SiteMap and SiteMapOutput values.
// You can construct a concrete instance of `SiteMapInput` via:
//
//	SiteMap{ "key": SiteArgs{...} }
type SiteMapInput interface {
	pulumi.Input

	ToSiteMapOutput() SiteMapOutput
	ToSiteMapOutputWithContext(context.Context) SiteMapOutput
}

type SiteMap map[string]SiteInput

func (SiteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Site)(nil)).Elem()
}

func (i SiteMap) ToSiteMapOutput() SiteMapOutput {
	return i.ToSiteMapOutputWithContext(context.Background())
}

func (i SiteMap) ToSiteMapOutputWithContext(ctx context.Context) SiteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteMapOutput)
}

type SiteOutput struct{ *pulumi.OutputState }

func (SiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Site)(nil)).Elem()
}

func (o SiteOutput) ToSiteOutput() SiteOutput {
	return o
}

func (o SiteOutput) ToSiteOutputWithContext(ctx context.Context) SiteOutput {
	return o
}

// Site Amazon Resource Name (ARN)
func (o SiteOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Site) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Description of the Site.
func (o SiteOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Site) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the Global Network to create the site in.
func (o SiteOutput) GlobalNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *Site) pulumi.StringOutput { return v.GlobalNetworkId }).(pulumi.StringOutput)
}

// The site location as documented below.
func (o SiteOutput) Location() SiteLocationPtrOutput {
	return o.ApplyT(func(v *Site) SiteLocationPtrOutput { return v.Location }).(SiteLocationPtrOutput)
}

// Key-value tags for the Site. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o SiteOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Site) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o SiteOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Site) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type SiteArrayOutput struct{ *pulumi.OutputState }

func (SiteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Site)(nil)).Elem()
}

func (o SiteArrayOutput) ToSiteArrayOutput() SiteArrayOutput {
	return o
}

func (o SiteArrayOutput) ToSiteArrayOutputWithContext(ctx context.Context) SiteArrayOutput {
	return o
}

func (o SiteArrayOutput) Index(i pulumi.IntInput) SiteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Site {
		return vs[0].([]*Site)[vs[1].(int)]
	}).(SiteOutput)
}

type SiteMapOutput struct{ *pulumi.OutputState }

func (SiteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Site)(nil)).Elem()
}

func (o SiteMapOutput) ToSiteMapOutput() SiteMapOutput {
	return o
}

func (o SiteMapOutput) ToSiteMapOutputWithContext(ctx context.Context) SiteMapOutput {
	return o
}

func (o SiteMapOutput) MapIndex(k pulumi.StringInput) SiteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Site {
		return vs[0].(map[string]*Site)[vs[1].(string)]
	}).(SiteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SiteInput)(nil)).Elem(), &Site{})
	pulumi.RegisterInputType(reflect.TypeOf((*SiteArrayInput)(nil)).Elem(), SiteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SiteMapInput)(nil)).Elem(), SiteMap{})
	pulumi.RegisterOutputType(SiteOutput{})
	pulumi.RegisterOutputType(SiteArrayOutput{})
	pulumi.RegisterOutputType(SiteMapOutput{})
}
