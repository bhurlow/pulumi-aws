// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package codecommit

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CodeCommit Repository Resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/codecommit"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := codecommit.NewRepository(ctx, "test", &codecommit.RepositoryArgs{
//				Description:    pulumi.String("This is the Sample App Repository"),
//				RepositoryName: pulumi.String("MyTestRepository"),
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
// Codecommit repository can be imported using repository name, e.g.,
//
// ```sh
//
//	$ pulumi import aws:codecommit/repository:Repository imported ExistingRepo
//
// ```
type Repository struct {
	pulumi.CustomResourceState

	// The ARN of the repository
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The URL to use for cloning the repository over HTTPS.
	CloneUrlHttp pulumi.StringOutput `pulumi:"cloneUrlHttp"`
	// The URL to use for cloning the repository over SSH.
	CloneUrlSsh pulumi.StringOutput `pulumi:"cloneUrlSsh"`
	// The default branch of the repository. The branch specified here needs to exist.
	DefaultBranch pulumi.StringPtrOutput `pulumi:"defaultBranch"`
	// The description of the repository. This needs to be less than 1000 characters
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the repository
	RepositoryId pulumi.StringOutput `pulumi:"repositoryId"`
	// The name for the repository. This needs to be less than 100 characters.
	RepositoryName pulumi.StringOutput `pulumi:"repositoryName"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewRepository registers a new resource with the given unique name, arguments, and options.
func NewRepository(ctx *pulumi.Context,
	name string, args *RepositoryArgs, opts ...pulumi.ResourceOption) (*Repository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RepositoryName == nil {
		return nil, errors.New("invalid value for required argument 'RepositoryName'")
	}
	var resource Repository
	err := ctx.RegisterResource("aws:codecommit/repository:Repository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepository gets an existing Repository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryState, opts ...pulumi.ResourceOption) (*Repository, error) {
	var resource Repository
	err := ctx.ReadResource("aws:codecommit/repository:Repository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Repository resources.
type repositoryState struct {
	// The ARN of the repository
	Arn *string `pulumi:"arn"`
	// The URL to use for cloning the repository over HTTPS.
	CloneUrlHttp *string `pulumi:"cloneUrlHttp"`
	// The URL to use for cloning the repository over SSH.
	CloneUrlSsh *string `pulumi:"cloneUrlSsh"`
	// The default branch of the repository. The branch specified here needs to exist.
	DefaultBranch *string `pulumi:"defaultBranch"`
	// The description of the repository. This needs to be less than 1000 characters
	Description *string `pulumi:"description"`
	// The ID of the repository
	RepositoryId *string `pulumi:"repositoryId"`
	// The name for the repository. This needs to be less than 100 characters.
	RepositoryName *string `pulumi:"repositoryName"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type RepositoryState struct {
	// The ARN of the repository
	Arn pulumi.StringPtrInput
	// The URL to use for cloning the repository over HTTPS.
	CloneUrlHttp pulumi.StringPtrInput
	// The URL to use for cloning the repository over SSH.
	CloneUrlSsh pulumi.StringPtrInput
	// The default branch of the repository. The branch specified here needs to exist.
	DefaultBranch pulumi.StringPtrInput
	// The description of the repository. This needs to be less than 1000 characters
	Description pulumi.StringPtrInput
	// The ID of the repository
	RepositoryId pulumi.StringPtrInput
	// The name for the repository. This needs to be less than 100 characters.
	RepositoryName pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (RepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryState)(nil)).Elem()
}

type repositoryArgs struct {
	// The default branch of the repository. The branch specified here needs to exist.
	DefaultBranch *string `pulumi:"defaultBranch"`
	// The description of the repository. This needs to be less than 1000 characters
	Description *string `pulumi:"description"`
	// The name for the repository. This needs to be less than 100 characters.
	RepositoryName string `pulumi:"repositoryName"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Repository resource.
type RepositoryArgs struct {
	// The default branch of the repository. The branch specified here needs to exist.
	DefaultBranch pulumi.StringPtrInput
	// The description of the repository. This needs to be less than 1000 characters
	Description pulumi.StringPtrInput
	// The name for the repository. This needs to be less than 100 characters.
	RepositoryName pulumi.StringInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (RepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryArgs)(nil)).Elem()
}

type RepositoryInput interface {
	pulumi.Input

	ToRepositoryOutput() RepositoryOutput
	ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput
}

func (*Repository) ElementType() reflect.Type {
	return reflect.TypeOf((**Repository)(nil)).Elem()
}

func (i *Repository) ToRepositoryOutput() RepositoryOutput {
	return i.ToRepositoryOutputWithContext(context.Background())
}

func (i *Repository) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryOutput)
}

// RepositoryArrayInput is an input type that accepts RepositoryArray and RepositoryArrayOutput values.
// You can construct a concrete instance of `RepositoryArrayInput` via:
//
//	RepositoryArray{ RepositoryArgs{...} }
type RepositoryArrayInput interface {
	pulumi.Input

	ToRepositoryArrayOutput() RepositoryArrayOutput
	ToRepositoryArrayOutputWithContext(context.Context) RepositoryArrayOutput
}

type RepositoryArray []RepositoryInput

func (RepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Repository)(nil)).Elem()
}

func (i RepositoryArray) ToRepositoryArrayOutput() RepositoryArrayOutput {
	return i.ToRepositoryArrayOutputWithContext(context.Background())
}

func (i RepositoryArray) ToRepositoryArrayOutputWithContext(ctx context.Context) RepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryArrayOutput)
}

// RepositoryMapInput is an input type that accepts RepositoryMap and RepositoryMapOutput values.
// You can construct a concrete instance of `RepositoryMapInput` via:
//
//	RepositoryMap{ "key": RepositoryArgs{...} }
type RepositoryMapInput interface {
	pulumi.Input

	ToRepositoryMapOutput() RepositoryMapOutput
	ToRepositoryMapOutputWithContext(context.Context) RepositoryMapOutput
}

type RepositoryMap map[string]RepositoryInput

func (RepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Repository)(nil)).Elem()
}

func (i RepositoryMap) ToRepositoryMapOutput() RepositoryMapOutput {
	return i.ToRepositoryMapOutputWithContext(context.Background())
}

func (i RepositoryMap) ToRepositoryMapOutputWithContext(ctx context.Context) RepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryMapOutput)
}

type RepositoryOutput struct{ *pulumi.OutputState }

func (RepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Repository)(nil)).Elem()
}

func (o RepositoryOutput) ToRepositoryOutput() RepositoryOutput {
	return o
}

func (o RepositoryOutput) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return o
}

// The ARN of the repository
func (o RepositoryOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The URL to use for cloning the repository over HTTPS.
func (o RepositoryOutput) CloneUrlHttp() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.CloneUrlHttp }).(pulumi.StringOutput)
}

// The URL to use for cloning the repository over SSH.
func (o RepositoryOutput) CloneUrlSsh() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.CloneUrlSsh }).(pulumi.StringOutput)
}

// The default branch of the repository. The branch specified here needs to exist.
func (o RepositoryOutput) DefaultBranch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringPtrOutput { return v.DefaultBranch }).(pulumi.StringPtrOutput)
}

// The description of the repository. This needs to be less than 1000 characters
func (o RepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the repository
func (o RepositoryOutput) RepositoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.RepositoryId }).(pulumi.StringOutput)
}

// The name for the repository. This needs to be less than 100 characters.
func (o RepositoryOutput) RepositoryName() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.RepositoryName }).(pulumi.StringOutput)
}

// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o RepositoryOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o RepositoryOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type RepositoryArrayOutput struct{ *pulumi.OutputState }

func (RepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Repository)(nil)).Elem()
}

func (o RepositoryArrayOutput) ToRepositoryArrayOutput() RepositoryArrayOutput {
	return o
}

func (o RepositoryArrayOutput) ToRepositoryArrayOutputWithContext(ctx context.Context) RepositoryArrayOutput {
	return o
}

func (o RepositoryArrayOutput) Index(i pulumi.IntInput) RepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Repository {
		return vs[0].([]*Repository)[vs[1].(int)]
	}).(RepositoryOutput)
}

type RepositoryMapOutput struct{ *pulumi.OutputState }

func (RepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Repository)(nil)).Elem()
}

func (o RepositoryMapOutput) ToRepositoryMapOutput() RepositoryMapOutput {
	return o
}

func (o RepositoryMapOutput) ToRepositoryMapOutputWithContext(ctx context.Context) RepositoryMapOutput {
	return o
}

func (o RepositoryMapOutput) MapIndex(k pulumi.StringInput) RepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Repository {
		return vs[0].(map[string]*Repository)[vs[1].(string)]
	}).(RepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryInput)(nil)).Elem(), &Repository{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryArrayInput)(nil)).Elem(), RepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryMapInput)(nil)).Elem(), RepositoryMap{})
	pulumi.RegisterOutputType(RepositoryOutput{})
	pulumi.RegisterOutputType(RepositoryArrayOutput{})
	pulumi.RegisterOutputType(RepositoryMapOutput{})
}
