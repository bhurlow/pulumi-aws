// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appconfig

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Associates an AppConfig Extension with a Resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appconfig"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			testTopic, err := sns.NewTopic(ctx, "testTopic", nil)
//			if err != nil {
//				return err
//			}
//			testPolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
//				Statements: []iam.GetPolicyDocumentStatement{
//					{
//						Actions: []string{
//							"sts:AssumeRole",
//						},
//						Principals: []iam.GetPolicyDocumentStatementPrincipal{
//							{
//								Type: "Service",
//								Identifiers: []string{
//									"appconfig.amazonaws.com",
//								},
//							},
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			testRole, err := iam.NewRole(ctx, "testRole", &iam.RoleArgs{
//				AssumeRolePolicy: *pulumi.String(testPolicyDocument.Json),
//			})
//			if err != nil {
//				return err
//			}
//			testExtension, err := appconfig.NewExtension(ctx, "testExtension", &appconfig.ExtensionArgs{
//				Description: pulumi.String("test description"),
//				ActionPoints: appconfig.ExtensionActionPointArray{
//					&appconfig.ExtensionActionPointArgs{
//						Point: pulumi.String("ON_DEPLOYMENT_COMPLETE"),
//						Actions: appconfig.ExtensionActionPointActionArray{
//							&appconfig.ExtensionActionPointActionArgs{
//								Name:    pulumi.String("test"),
//								RoleArn: testRole.Arn,
//								Uri:     testTopic.Arn,
//							},
//						},
//					},
//				},
//				Tags: pulumi.StringMap{
//					"Type": pulumi.String("AppConfig Extension"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			testApplication, err := appconfig.NewApplication(ctx, "testApplication", nil)
//			if err != nil {
//				return err
//			}
//			_, err = appconfig.NewExtensionAssociation(ctx, "testExtensionAssociation", &appconfig.ExtensionAssociationArgs{
//				ExtensionArn: testExtension.Arn,
//				ResourceArn:  testApplication.Arn,
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
//	to = aws_appconfig_extension_association.example
//
//	id = "71rxuzt" } Using `pulumi import`, import AppConfig Extension Associations using their extension association ID. For exampleconsole % pulumi import aws_appconfig_extension_association.example 71rxuzt
type ExtensionAssociation struct {
	pulumi.CustomResourceState

	// ARN of the AppConfig Extension Association.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ARN of the extension defined in the association.
	ExtensionArn pulumi.StringOutput `pulumi:"extensionArn"`
	// The version number for the extension defined in the association.
	ExtensionVersion pulumi.IntOutput `pulumi:"extensionVersion"`
	// The parameter names and values defined for the association.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The ARN of the application, configuration profile, or environment to associate with the extension.
	ResourceArn pulumi.StringOutput `pulumi:"resourceArn"`
}

// NewExtensionAssociation registers a new resource with the given unique name, arguments, and options.
func NewExtensionAssociation(ctx *pulumi.Context,
	name string, args *ExtensionAssociationArgs, opts ...pulumi.ResourceOption) (*ExtensionAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExtensionArn == nil {
		return nil, errors.New("invalid value for required argument 'ExtensionArn'")
	}
	if args.ResourceArn == nil {
		return nil, errors.New("invalid value for required argument 'ResourceArn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ExtensionAssociation
	err := ctx.RegisterResource("aws:appconfig/extensionAssociation:ExtensionAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExtensionAssociation gets an existing ExtensionAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExtensionAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExtensionAssociationState, opts ...pulumi.ResourceOption) (*ExtensionAssociation, error) {
	var resource ExtensionAssociation
	err := ctx.ReadResource("aws:appconfig/extensionAssociation:ExtensionAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExtensionAssociation resources.
type extensionAssociationState struct {
	// ARN of the AppConfig Extension Association.
	Arn *string `pulumi:"arn"`
	// The ARN of the extension defined in the association.
	ExtensionArn *string `pulumi:"extensionArn"`
	// The version number for the extension defined in the association.
	ExtensionVersion *int `pulumi:"extensionVersion"`
	// The parameter names and values defined for the association.
	Parameters map[string]string `pulumi:"parameters"`
	// The ARN of the application, configuration profile, or environment to associate with the extension.
	ResourceArn *string `pulumi:"resourceArn"`
}

type ExtensionAssociationState struct {
	// ARN of the AppConfig Extension Association.
	Arn pulumi.StringPtrInput
	// The ARN of the extension defined in the association.
	ExtensionArn pulumi.StringPtrInput
	// The version number for the extension defined in the association.
	ExtensionVersion pulumi.IntPtrInput
	// The parameter names and values defined for the association.
	Parameters pulumi.StringMapInput
	// The ARN of the application, configuration profile, or environment to associate with the extension.
	ResourceArn pulumi.StringPtrInput
}

func (ExtensionAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionAssociationState)(nil)).Elem()
}

type extensionAssociationArgs struct {
	// The ARN of the extension defined in the association.
	ExtensionArn string `pulumi:"extensionArn"`
	// The parameter names and values defined for the association.
	Parameters map[string]string `pulumi:"parameters"`
	// The ARN of the application, configuration profile, or environment to associate with the extension.
	ResourceArn string `pulumi:"resourceArn"`
}

// The set of arguments for constructing a ExtensionAssociation resource.
type ExtensionAssociationArgs struct {
	// The ARN of the extension defined in the association.
	ExtensionArn pulumi.StringInput
	// The parameter names and values defined for the association.
	Parameters pulumi.StringMapInput
	// The ARN of the application, configuration profile, or environment to associate with the extension.
	ResourceArn pulumi.StringInput
}

func (ExtensionAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionAssociationArgs)(nil)).Elem()
}

type ExtensionAssociationInput interface {
	pulumi.Input

	ToExtensionAssociationOutput() ExtensionAssociationOutput
	ToExtensionAssociationOutputWithContext(ctx context.Context) ExtensionAssociationOutput
}

func (*ExtensionAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtensionAssociation)(nil)).Elem()
}

func (i *ExtensionAssociation) ToExtensionAssociationOutput() ExtensionAssociationOutput {
	return i.ToExtensionAssociationOutputWithContext(context.Background())
}

func (i *ExtensionAssociation) ToExtensionAssociationOutputWithContext(ctx context.Context) ExtensionAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionAssociationOutput)
}

// ExtensionAssociationArrayInput is an input type that accepts ExtensionAssociationArray and ExtensionAssociationArrayOutput values.
// You can construct a concrete instance of `ExtensionAssociationArrayInput` via:
//
//	ExtensionAssociationArray{ ExtensionAssociationArgs{...} }
type ExtensionAssociationArrayInput interface {
	pulumi.Input

	ToExtensionAssociationArrayOutput() ExtensionAssociationArrayOutput
	ToExtensionAssociationArrayOutputWithContext(context.Context) ExtensionAssociationArrayOutput
}

type ExtensionAssociationArray []ExtensionAssociationInput

func (ExtensionAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExtensionAssociation)(nil)).Elem()
}

func (i ExtensionAssociationArray) ToExtensionAssociationArrayOutput() ExtensionAssociationArrayOutput {
	return i.ToExtensionAssociationArrayOutputWithContext(context.Background())
}

func (i ExtensionAssociationArray) ToExtensionAssociationArrayOutputWithContext(ctx context.Context) ExtensionAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionAssociationArrayOutput)
}

// ExtensionAssociationMapInput is an input type that accepts ExtensionAssociationMap and ExtensionAssociationMapOutput values.
// You can construct a concrete instance of `ExtensionAssociationMapInput` via:
//
//	ExtensionAssociationMap{ "key": ExtensionAssociationArgs{...} }
type ExtensionAssociationMapInput interface {
	pulumi.Input

	ToExtensionAssociationMapOutput() ExtensionAssociationMapOutput
	ToExtensionAssociationMapOutputWithContext(context.Context) ExtensionAssociationMapOutput
}

type ExtensionAssociationMap map[string]ExtensionAssociationInput

func (ExtensionAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExtensionAssociation)(nil)).Elem()
}

func (i ExtensionAssociationMap) ToExtensionAssociationMapOutput() ExtensionAssociationMapOutput {
	return i.ToExtensionAssociationMapOutputWithContext(context.Background())
}

func (i ExtensionAssociationMap) ToExtensionAssociationMapOutputWithContext(ctx context.Context) ExtensionAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionAssociationMapOutput)
}

type ExtensionAssociationOutput struct{ *pulumi.OutputState }

func (ExtensionAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtensionAssociation)(nil)).Elem()
}

func (o ExtensionAssociationOutput) ToExtensionAssociationOutput() ExtensionAssociationOutput {
	return o
}

func (o ExtensionAssociationOutput) ToExtensionAssociationOutputWithContext(ctx context.Context) ExtensionAssociationOutput {
	return o
}

// ARN of the AppConfig Extension Association.
func (o ExtensionAssociationOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensionAssociation) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The ARN of the extension defined in the association.
func (o ExtensionAssociationOutput) ExtensionArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensionAssociation) pulumi.StringOutput { return v.ExtensionArn }).(pulumi.StringOutput)
}

// The version number for the extension defined in the association.
func (o ExtensionAssociationOutput) ExtensionVersion() pulumi.IntOutput {
	return o.ApplyT(func(v *ExtensionAssociation) pulumi.IntOutput { return v.ExtensionVersion }).(pulumi.IntOutput)
}

// The parameter names and values defined for the association.
func (o ExtensionAssociationOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ExtensionAssociation) pulumi.StringMapOutput { return v.Parameters }).(pulumi.StringMapOutput)
}

// The ARN of the application, configuration profile, or environment to associate with the extension.
func (o ExtensionAssociationOutput) ResourceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensionAssociation) pulumi.StringOutput { return v.ResourceArn }).(pulumi.StringOutput)
}

type ExtensionAssociationArrayOutput struct{ *pulumi.OutputState }

func (ExtensionAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExtensionAssociation)(nil)).Elem()
}

func (o ExtensionAssociationArrayOutput) ToExtensionAssociationArrayOutput() ExtensionAssociationArrayOutput {
	return o
}

func (o ExtensionAssociationArrayOutput) ToExtensionAssociationArrayOutputWithContext(ctx context.Context) ExtensionAssociationArrayOutput {
	return o
}

func (o ExtensionAssociationArrayOutput) Index(i pulumi.IntInput) ExtensionAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ExtensionAssociation {
		return vs[0].([]*ExtensionAssociation)[vs[1].(int)]
	}).(ExtensionAssociationOutput)
}

type ExtensionAssociationMapOutput struct{ *pulumi.OutputState }

func (ExtensionAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExtensionAssociation)(nil)).Elem()
}

func (o ExtensionAssociationMapOutput) ToExtensionAssociationMapOutput() ExtensionAssociationMapOutput {
	return o
}

func (o ExtensionAssociationMapOutput) ToExtensionAssociationMapOutputWithContext(ctx context.Context) ExtensionAssociationMapOutput {
	return o
}

func (o ExtensionAssociationMapOutput) MapIndex(k pulumi.StringInput) ExtensionAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ExtensionAssociation {
		return vs[0].(map[string]*ExtensionAssociation)[vs[1].(string)]
	}).(ExtensionAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExtensionAssociationInput)(nil)).Elem(), &ExtensionAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtensionAssociationArrayInput)(nil)).Elem(), ExtensionAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtensionAssociationMapInput)(nil)).Elem(), ExtensionAssociationMap{})
	pulumi.RegisterOutputType(ExtensionAssociationOutput{})
	pulumi.RegisterOutputType(ExtensionAssociationArrayOutput{})
	pulumi.RegisterOutputType(ExtensionAssociationMapOutput{})
}
