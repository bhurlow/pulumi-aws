// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Glue User Defined Function Resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleCatalogDatabase, err := glue.NewCatalogDatabase(ctx, "exampleCatalogDatabase", &glue.CatalogDatabaseArgs{
//				Name: pulumi.String("my_database"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = glue.NewUserDefinedFunction(ctx, "exampleUserDefinedFunction", &glue.UserDefinedFunctionArgs{
//				CatalogId:    exampleCatalogDatabase.CatalogId,
//				DatabaseName: exampleCatalogDatabase.Name,
//				ClassName:    pulumi.String("class"),
//				OwnerName:    pulumi.String("owner"),
//				OwnerType:    pulumi.String("GROUP"),
//				ResourceUris: glue.UserDefinedFunctionResourceUriArray{
//					&glue.UserDefinedFunctionResourceUriArgs{
//						ResourceType: pulumi.String("ARCHIVE"),
//						Uri:          pulumi.String("uri"),
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
// terraform import {
//
//	to = aws_glue_user_defined_function.func
//
//	id = "123456789012:my_database:my_func" } Using `pulumi import`, import Glue User Defined Functions using the `catalog_id:database_name:function_name`. If you have not set a Catalog ID specify the AWS Account ID that the database is in. For exampleconsole % pulumi import aws_glue_user_defined_function.func 123456789012:my_database:my_func
type UserDefinedFunction struct {
	pulumi.CustomResourceState

	// The ARN of the Glue User Defined Function.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// ID of the Glue Catalog to create the function in. If omitted, this defaults to the AWS Account ID.
	CatalogId pulumi.StringPtrOutput `pulumi:"catalogId"`
	// The Java class that contains the function code.
	ClassName pulumi.StringOutput `pulumi:"className"`
	// The time at which the function was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The name of the Database to create the Function.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// The name of the function.
	Name pulumi.StringOutput `pulumi:"name"`
	// The owner of the function.
	OwnerName pulumi.StringOutput `pulumi:"ownerName"`
	// The owner type. can be one of `USER`, `ROLE`, and `GROUP`.
	OwnerType pulumi.StringOutput `pulumi:"ownerType"`
	// The configuration block for Resource URIs. See resource uris below for more details.
	ResourceUris UserDefinedFunctionResourceUriArrayOutput `pulumi:"resourceUris"`
}

// NewUserDefinedFunction registers a new resource with the given unique name, arguments, and options.
func NewUserDefinedFunction(ctx *pulumi.Context,
	name string, args *UserDefinedFunctionArgs, opts ...pulumi.ResourceOption) (*UserDefinedFunction, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClassName == nil {
		return nil, errors.New("invalid value for required argument 'ClassName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.OwnerName == nil {
		return nil, errors.New("invalid value for required argument 'OwnerName'")
	}
	if args.OwnerType == nil {
		return nil, errors.New("invalid value for required argument 'OwnerType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserDefinedFunction
	err := ctx.RegisterResource("aws:glue/userDefinedFunction:UserDefinedFunction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserDefinedFunction gets an existing UserDefinedFunction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserDefinedFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserDefinedFunctionState, opts ...pulumi.ResourceOption) (*UserDefinedFunction, error) {
	var resource UserDefinedFunction
	err := ctx.ReadResource("aws:glue/userDefinedFunction:UserDefinedFunction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserDefinedFunction resources.
type userDefinedFunctionState struct {
	// The ARN of the Glue User Defined Function.
	Arn *string `pulumi:"arn"`
	// ID of the Glue Catalog to create the function in. If omitted, this defaults to the AWS Account ID.
	CatalogId *string `pulumi:"catalogId"`
	// The Java class that contains the function code.
	ClassName *string `pulumi:"className"`
	// The time at which the function was created.
	CreateTime *string `pulumi:"createTime"`
	// The name of the Database to create the Function.
	DatabaseName *string `pulumi:"databaseName"`
	// The name of the function.
	Name *string `pulumi:"name"`
	// The owner of the function.
	OwnerName *string `pulumi:"ownerName"`
	// The owner type. can be one of `USER`, `ROLE`, and `GROUP`.
	OwnerType *string `pulumi:"ownerType"`
	// The configuration block for Resource URIs. See resource uris below for more details.
	ResourceUris []UserDefinedFunctionResourceUri `pulumi:"resourceUris"`
}

type UserDefinedFunctionState struct {
	// The ARN of the Glue User Defined Function.
	Arn pulumi.StringPtrInput
	// ID of the Glue Catalog to create the function in. If omitted, this defaults to the AWS Account ID.
	CatalogId pulumi.StringPtrInput
	// The Java class that contains the function code.
	ClassName pulumi.StringPtrInput
	// The time at which the function was created.
	CreateTime pulumi.StringPtrInput
	// The name of the Database to create the Function.
	DatabaseName pulumi.StringPtrInput
	// The name of the function.
	Name pulumi.StringPtrInput
	// The owner of the function.
	OwnerName pulumi.StringPtrInput
	// The owner type. can be one of `USER`, `ROLE`, and `GROUP`.
	OwnerType pulumi.StringPtrInput
	// The configuration block for Resource URIs. See resource uris below for more details.
	ResourceUris UserDefinedFunctionResourceUriArrayInput
}

func (UserDefinedFunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*userDefinedFunctionState)(nil)).Elem()
}

type userDefinedFunctionArgs struct {
	// ID of the Glue Catalog to create the function in. If omitted, this defaults to the AWS Account ID.
	CatalogId *string `pulumi:"catalogId"`
	// The Java class that contains the function code.
	ClassName string `pulumi:"className"`
	// The name of the Database to create the Function.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the function.
	Name *string `pulumi:"name"`
	// The owner of the function.
	OwnerName string `pulumi:"ownerName"`
	// The owner type. can be one of `USER`, `ROLE`, and `GROUP`.
	OwnerType string `pulumi:"ownerType"`
	// The configuration block for Resource URIs. See resource uris below for more details.
	ResourceUris []UserDefinedFunctionResourceUri `pulumi:"resourceUris"`
}

// The set of arguments for constructing a UserDefinedFunction resource.
type UserDefinedFunctionArgs struct {
	// ID of the Glue Catalog to create the function in. If omitted, this defaults to the AWS Account ID.
	CatalogId pulumi.StringPtrInput
	// The Java class that contains the function code.
	ClassName pulumi.StringInput
	// The name of the Database to create the Function.
	DatabaseName pulumi.StringInput
	// The name of the function.
	Name pulumi.StringPtrInput
	// The owner of the function.
	OwnerName pulumi.StringInput
	// The owner type. can be one of `USER`, `ROLE`, and `GROUP`.
	OwnerType pulumi.StringInput
	// The configuration block for Resource URIs. See resource uris below for more details.
	ResourceUris UserDefinedFunctionResourceUriArrayInput
}

func (UserDefinedFunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userDefinedFunctionArgs)(nil)).Elem()
}

type UserDefinedFunctionInput interface {
	pulumi.Input

	ToUserDefinedFunctionOutput() UserDefinedFunctionOutput
	ToUserDefinedFunctionOutputWithContext(ctx context.Context) UserDefinedFunctionOutput
}

func (*UserDefinedFunction) ElementType() reflect.Type {
	return reflect.TypeOf((**UserDefinedFunction)(nil)).Elem()
}

func (i *UserDefinedFunction) ToUserDefinedFunctionOutput() UserDefinedFunctionOutput {
	return i.ToUserDefinedFunctionOutputWithContext(context.Background())
}

func (i *UserDefinedFunction) ToUserDefinedFunctionOutputWithContext(ctx context.Context) UserDefinedFunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDefinedFunctionOutput)
}

// UserDefinedFunctionArrayInput is an input type that accepts UserDefinedFunctionArray and UserDefinedFunctionArrayOutput values.
// You can construct a concrete instance of `UserDefinedFunctionArrayInput` via:
//
//	UserDefinedFunctionArray{ UserDefinedFunctionArgs{...} }
type UserDefinedFunctionArrayInput interface {
	pulumi.Input

	ToUserDefinedFunctionArrayOutput() UserDefinedFunctionArrayOutput
	ToUserDefinedFunctionArrayOutputWithContext(context.Context) UserDefinedFunctionArrayOutput
}

type UserDefinedFunctionArray []UserDefinedFunctionInput

func (UserDefinedFunctionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserDefinedFunction)(nil)).Elem()
}

func (i UserDefinedFunctionArray) ToUserDefinedFunctionArrayOutput() UserDefinedFunctionArrayOutput {
	return i.ToUserDefinedFunctionArrayOutputWithContext(context.Background())
}

func (i UserDefinedFunctionArray) ToUserDefinedFunctionArrayOutputWithContext(ctx context.Context) UserDefinedFunctionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDefinedFunctionArrayOutput)
}

// UserDefinedFunctionMapInput is an input type that accepts UserDefinedFunctionMap and UserDefinedFunctionMapOutput values.
// You can construct a concrete instance of `UserDefinedFunctionMapInput` via:
//
//	UserDefinedFunctionMap{ "key": UserDefinedFunctionArgs{...} }
type UserDefinedFunctionMapInput interface {
	pulumi.Input

	ToUserDefinedFunctionMapOutput() UserDefinedFunctionMapOutput
	ToUserDefinedFunctionMapOutputWithContext(context.Context) UserDefinedFunctionMapOutput
}

type UserDefinedFunctionMap map[string]UserDefinedFunctionInput

func (UserDefinedFunctionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserDefinedFunction)(nil)).Elem()
}

func (i UserDefinedFunctionMap) ToUserDefinedFunctionMapOutput() UserDefinedFunctionMapOutput {
	return i.ToUserDefinedFunctionMapOutputWithContext(context.Background())
}

func (i UserDefinedFunctionMap) ToUserDefinedFunctionMapOutputWithContext(ctx context.Context) UserDefinedFunctionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDefinedFunctionMapOutput)
}

type UserDefinedFunctionOutput struct{ *pulumi.OutputState }

func (UserDefinedFunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserDefinedFunction)(nil)).Elem()
}

func (o UserDefinedFunctionOutput) ToUserDefinedFunctionOutput() UserDefinedFunctionOutput {
	return o
}

func (o UserDefinedFunctionOutput) ToUserDefinedFunctionOutputWithContext(ctx context.Context) UserDefinedFunctionOutput {
	return o
}

// The ARN of the Glue User Defined Function.
func (o UserDefinedFunctionOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDefinedFunction) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// ID of the Glue Catalog to create the function in. If omitted, this defaults to the AWS Account ID.
func (o UserDefinedFunctionOutput) CatalogId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserDefinedFunction) pulumi.StringPtrOutput { return v.CatalogId }).(pulumi.StringPtrOutput)
}

// The Java class that contains the function code.
func (o UserDefinedFunctionOutput) ClassName() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDefinedFunction) pulumi.StringOutput { return v.ClassName }).(pulumi.StringOutput)
}

// The time at which the function was created.
func (o UserDefinedFunctionOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDefinedFunction) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The name of the Database to create the Function.
func (o UserDefinedFunctionOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDefinedFunction) pulumi.StringOutput { return v.DatabaseName }).(pulumi.StringOutput)
}

// The name of the function.
func (o UserDefinedFunctionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDefinedFunction) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The owner of the function.
func (o UserDefinedFunctionOutput) OwnerName() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDefinedFunction) pulumi.StringOutput { return v.OwnerName }).(pulumi.StringOutput)
}

// The owner type. can be one of `USER`, `ROLE`, and `GROUP`.
func (o UserDefinedFunctionOutput) OwnerType() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDefinedFunction) pulumi.StringOutput { return v.OwnerType }).(pulumi.StringOutput)
}

// The configuration block for Resource URIs. See resource uris below for more details.
func (o UserDefinedFunctionOutput) ResourceUris() UserDefinedFunctionResourceUriArrayOutput {
	return o.ApplyT(func(v *UserDefinedFunction) UserDefinedFunctionResourceUriArrayOutput { return v.ResourceUris }).(UserDefinedFunctionResourceUriArrayOutput)
}

type UserDefinedFunctionArrayOutput struct{ *pulumi.OutputState }

func (UserDefinedFunctionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserDefinedFunction)(nil)).Elem()
}

func (o UserDefinedFunctionArrayOutput) ToUserDefinedFunctionArrayOutput() UserDefinedFunctionArrayOutput {
	return o
}

func (o UserDefinedFunctionArrayOutput) ToUserDefinedFunctionArrayOutputWithContext(ctx context.Context) UserDefinedFunctionArrayOutput {
	return o
}

func (o UserDefinedFunctionArrayOutput) Index(i pulumi.IntInput) UserDefinedFunctionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserDefinedFunction {
		return vs[0].([]*UserDefinedFunction)[vs[1].(int)]
	}).(UserDefinedFunctionOutput)
}

type UserDefinedFunctionMapOutput struct{ *pulumi.OutputState }

func (UserDefinedFunctionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserDefinedFunction)(nil)).Elem()
}

func (o UserDefinedFunctionMapOutput) ToUserDefinedFunctionMapOutput() UserDefinedFunctionMapOutput {
	return o
}

func (o UserDefinedFunctionMapOutput) ToUserDefinedFunctionMapOutputWithContext(ctx context.Context) UserDefinedFunctionMapOutput {
	return o
}

func (o UserDefinedFunctionMapOutput) MapIndex(k pulumi.StringInput) UserDefinedFunctionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserDefinedFunction {
		return vs[0].(map[string]*UserDefinedFunction)[vs[1].(string)]
	}).(UserDefinedFunctionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserDefinedFunctionInput)(nil)).Elem(), &UserDefinedFunction{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserDefinedFunctionArrayInput)(nil)).Elem(), UserDefinedFunctionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserDefinedFunctionMapInput)(nil)).Elem(), UserDefinedFunctionMap{})
	pulumi.RegisterOutputType(UserDefinedFunctionOutput{})
	pulumi.RegisterOutputType(UserDefinedFunctionArrayOutput{})
	pulumi.RegisterOutputType(UserDefinedFunctionMapOutput{})
}
