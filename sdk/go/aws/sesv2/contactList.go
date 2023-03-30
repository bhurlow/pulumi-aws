// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sesv2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for managing an AWS SESv2 (Simple Email V2) Contact List.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/sesv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sesv2.NewContactList(ctx, "example", &sesv2.ContactListArgs{
//				ContactListName: pulumi.String("example"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Extended Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/sesv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sesv2.NewContactList(ctx, "example", &sesv2.ContactListArgs{
//				ContactListName: pulumi.String("example"),
//				Description:     pulumi.String("description"),
//				Topics: sesv2.ContactListTopicArray{
//					&sesv2.ContactListTopicArgs{
//						DefaultSubscriptionStatus: pulumi.String("OPT_IN"),
//						Description:               pulumi.String("topic description"),
//						DisplayName:               pulumi.String("Example Topic"),
//						TopicName:                 pulumi.String("example-topic"),
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
// SESv2 (Simple Email V2) Contact List can be imported using the `example_id_arg`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:sesv2/contactList:ContactList example example
//
// ```
type ContactList struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the contact list.
	ContactListName pulumi.StringOutput `pulumi:"contactListName"`
	// A timestamp noting when the contact list was created in ISO 8601 format.
	CreatedTimestamp pulumi.StringOutput `pulumi:"createdTimestamp"`
	// A description of what the contact list is about.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A timestamp noting the last time the contact list was updated in ISO 8601 format.
	LastUpdatedTimestamp pulumi.StringOutput `pulumi:"lastUpdatedTimestamp"`
	// Key-value map of resource tags for the contact list. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags    pulumi.StringMapOutput `pulumi:"tags"`
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Configuration block(s) with topic for the contact list. Detailed below.
	Topics ContactListTopicArrayOutput `pulumi:"topics"`
}

// NewContactList registers a new resource with the given unique name, arguments, and options.
func NewContactList(ctx *pulumi.Context,
	name string, args *ContactListArgs, opts ...pulumi.ResourceOption) (*ContactList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContactListName == nil {
		return nil, errors.New("invalid value for required argument 'ContactListName'")
	}
	var resource ContactList
	err := ctx.RegisterResource("aws:sesv2/contactList:ContactList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContactList gets an existing ContactList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContactList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContactListState, opts ...pulumi.ResourceOption) (*ContactList, error) {
	var resource ContactList
	err := ctx.ReadResource("aws:sesv2/contactList:ContactList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContactList resources.
type contactListState struct {
	Arn *string `pulumi:"arn"`
	// The name of the contact list.
	ContactListName *string `pulumi:"contactListName"`
	// A timestamp noting when the contact list was created in ISO 8601 format.
	CreatedTimestamp *string `pulumi:"createdTimestamp"`
	// A description of what the contact list is about.
	Description *string `pulumi:"description"`
	// A timestamp noting the last time the contact list was updated in ISO 8601 format.
	LastUpdatedTimestamp *string `pulumi:"lastUpdatedTimestamp"`
	// Key-value map of resource tags for the contact list. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags    map[string]string `pulumi:"tags"`
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Configuration block(s) with topic for the contact list. Detailed below.
	Topics []ContactListTopic `pulumi:"topics"`
}

type ContactListState struct {
	Arn pulumi.StringPtrInput
	// The name of the contact list.
	ContactListName pulumi.StringPtrInput
	// A timestamp noting when the contact list was created in ISO 8601 format.
	CreatedTimestamp pulumi.StringPtrInput
	// A description of what the contact list is about.
	Description pulumi.StringPtrInput
	// A timestamp noting the last time the contact list was updated in ISO 8601 format.
	LastUpdatedTimestamp pulumi.StringPtrInput
	// Key-value map of resource tags for the contact list. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags    pulumi.StringMapInput
	TagsAll pulumi.StringMapInput
	// Configuration block(s) with topic for the contact list. Detailed below.
	Topics ContactListTopicArrayInput
}

func (ContactListState) ElementType() reflect.Type {
	return reflect.TypeOf((*contactListState)(nil)).Elem()
}

type contactListArgs struct {
	// The name of the contact list.
	ContactListName string `pulumi:"contactListName"`
	// A description of what the contact list is about.
	Description *string `pulumi:"description"`
	// Key-value map of resource tags for the contact list. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Configuration block(s) with topic for the contact list. Detailed below.
	Topics []ContactListTopic `pulumi:"topics"`
}

// The set of arguments for constructing a ContactList resource.
type ContactListArgs struct {
	// The name of the contact list.
	ContactListName pulumi.StringInput
	// A description of what the contact list is about.
	Description pulumi.StringPtrInput
	// Key-value map of resource tags for the contact list. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Configuration block(s) with topic for the contact list. Detailed below.
	Topics ContactListTopicArrayInput
}

func (ContactListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contactListArgs)(nil)).Elem()
}

type ContactListInput interface {
	pulumi.Input

	ToContactListOutput() ContactListOutput
	ToContactListOutputWithContext(ctx context.Context) ContactListOutput
}

func (*ContactList) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactList)(nil)).Elem()
}

func (i *ContactList) ToContactListOutput() ContactListOutput {
	return i.ToContactListOutputWithContext(context.Background())
}

func (i *ContactList) ToContactListOutputWithContext(ctx context.Context) ContactListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactListOutput)
}

// ContactListArrayInput is an input type that accepts ContactListArray and ContactListArrayOutput values.
// You can construct a concrete instance of `ContactListArrayInput` via:
//
//	ContactListArray{ ContactListArgs{...} }
type ContactListArrayInput interface {
	pulumi.Input

	ToContactListArrayOutput() ContactListArrayOutput
	ToContactListArrayOutputWithContext(context.Context) ContactListArrayOutput
}

type ContactListArray []ContactListInput

func (ContactListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContactList)(nil)).Elem()
}

func (i ContactListArray) ToContactListArrayOutput() ContactListArrayOutput {
	return i.ToContactListArrayOutputWithContext(context.Background())
}

func (i ContactListArray) ToContactListArrayOutputWithContext(ctx context.Context) ContactListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactListArrayOutput)
}

// ContactListMapInput is an input type that accepts ContactListMap and ContactListMapOutput values.
// You can construct a concrete instance of `ContactListMapInput` via:
//
//	ContactListMap{ "key": ContactListArgs{...} }
type ContactListMapInput interface {
	pulumi.Input

	ToContactListMapOutput() ContactListMapOutput
	ToContactListMapOutputWithContext(context.Context) ContactListMapOutput
}

type ContactListMap map[string]ContactListInput

func (ContactListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContactList)(nil)).Elem()
}

func (i ContactListMap) ToContactListMapOutput() ContactListMapOutput {
	return i.ToContactListMapOutputWithContext(context.Background())
}

func (i ContactListMap) ToContactListMapOutputWithContext(ctx context.Context) ContactListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactListMapOutput)
}

type ContactListOutput struct{ *pulumi.OutputState }

func (ContactListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactList)(nil)).Elem()
}

func (o ContactListOutput) ToContactListOutput() ContactListOutput {
	return o
}

func (o ContactListOutput) ToContactListOutputWithContext(ctx context.Context) ContactListOutput {
	return o
}

func (o ContactListOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ContactList) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The name of the contact list.
func (o ContactListOutput) ContactListName() pulumi.StringOutput {
	return o.ApplyT(func(v *ContactList) pulumi.StringOutput { return v.ContactListName }).(pulumi.StringOutput)
}

// A timestamp noting when the contact list was created in ISO 8601 format.
func (o ContactListOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *ContactList) pulumi.StringOutput { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

// A description of what the contact list is about.
func (o ContactListOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactList) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A timestamp noting the last time the contact list was updated in ISO 8601 format.
func (o ContactListOutput) LastUpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *ContactList) pulumi.StringOutput { return v.LastUpdatedTimestamp }).(pulumi.StringOutput)
}

// Key-value map of resource tags for the contact list. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ContactListOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ContactList) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ContactListOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ContactList) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Configuration block(s) with topic for the contact list. Detailed below.
func (o ContactListOutput) Topics() ContactListTopicArrayOutput {
	return o.ApplyT(func(v *ContactList) ContactListTopicArrayOutput { return v.Topics }).(ContactListTopicArrayOutput)
}

type ContactListArrayOutput struct{ *pulumi.OutputState }

func (ContactListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContactList)(nil)).Elem()
}

func (o ContactListArrayOutput) ToContactListArrayOutput() ContactListArrayOutput {
	return o
}

func (o ContactListArrayOutput) ToContactListArrayOutputWithContext(ctx context.Context) ContactListArrayOutput {
	return o
}

func (o ContactListArrayOutput) Index(i pulumi.IntInput) ContactListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContactList {
		return vs[0].([]*ContactList)[vs[1].(int)]
	}).(ContactListOutput)
}

type ContactListMapOutput struct{ *pulumi.OutputState }

func (ContactListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContactList)(nil)).Elem()
}

func (o ContactListMapOutput) ToContactListMapOutput() ContactListMapOutput {
	return o
}

func (o ContactListMapOutput) ToContactListMapOutputWithContext(ctx context.Context) ContactListMapOutput {
	return o
}

func (o ContactListMapOutput) MapIndex(k pulumi.StringInput) ContactListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContactList {
		return vs[0].(map[string]*ContactList)[vs[1].(string)]
	}).(ContactListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContactListInput)(nil)).Elem(), &ContactList{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContactListArrayInput)(nil)).Elem(), ContactListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContactListMapInput)(nil)).Elem(), ContactListMap{})
	pulumi.RegisterOutputType(ContactListOutput{})
	pulumi.RegisterOutputType(ContactListArrayOutput{})
	pulumi.RegisterOutputType(ContactListMapOutput{})
}
