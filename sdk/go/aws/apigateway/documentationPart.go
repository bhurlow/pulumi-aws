// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a settings of an API Gateway Documentation Part.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleRestApi, err := apigateway.NewRestApi(ctx, "exampleRestApi", nil)
//			if err != nil {
//				return err
//			}
//			_, err = apigateway.NewDocumentationPart(ctx, "exampleDocumentationPart", &apigateway.DocumentationPartArgs{
//				Location: &apigateway.DocumentationPartLocationArgs{
//					Type:   pulumi.String("METHOD"),
//					Method: pulumi.String("GET"),
//					Path:   pulumi.String("/example"),
//				},
//				Properties: pulumi.String("{\"description\":\"Example description\"}"),
//				RestApiId:  exampleRestApi.ID(),
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
// Using `pulumi import`, import API Gateway documentation_parts using `REST-API-ID/DOC-PART-ID`. For example:
//
// ```sh
//
//	$ pulumi import aws:apigateway/documentationPart:DocumentationPart example 5i4e1ko720/3oyy3t
//
// ```
type DocumentationPart struct {
	pulumi.CustomResourceState

	// Location of the targeted API entity of the to-be-created documentation part. See below.
	Location DocumentationPartLocationOutput `pulumi:"location"`
	// Content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
	Properties pulumi.StringOutput `pulumi:"properties"`
	// ID of the associated Rest API
	RestApiId pulumi.StringOutput `pulumi:"restApiId"`
}

// NewDocumentationPart registers a new resource with the given unique name, arguments, and options.
func NewDocumentationPart(ctx *pulumi.Context,
	name string, args *DocumentationPartArgs, opts ...pulumi.ResourceOption) (*DocumentationPart, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.RestApiId == nil {
		return nil, errors.New("invalid value for required argument 'RestApiId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DocumentationPart
	err := ctx.RegisterResource("aws:apigateway/documentationPart:DocumentationPart", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDocumentationPart gets an existing DocumentationPart resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDocumentationPart(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DocumentationPartState, opts ...pulumi.ResourceOption) (*DocumentationPart, error) {
	var resource DocumentationPart
	err := ctx.ReadResource("aws:apigateway/documentationPart:DocumentationPart", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DocumentationPart resources.
type documentationPartState struct {
	// Location of the targeted API entity of the to-be-created documentation part. See below.
	Location *DocumentationPartLocation `pulumi:"location"`
	// Content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
	Properties *string `pulumi:"properties"`
	// ID of the associated Rest API
	RestApiId *string `pulumi:"restApiId"`
}

type DocumentationPartState struct {
	// Location of the targeted API entity of the to-be-created documentation part. See below.
	Location DocumentationPartLocationPtrInput
	// Content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
	Properties pulumi.StringPtrInput
	// ID of the associated Rest API
	RestApiId pulumi.StringPtrInput
}

func (DocumentationPartState) ElementType() reflect.Type {
	return reflect.TypeOf((*documentationPartState)(nil)).Elem()
}

type documentationPartArgs struct {
	// Location of the targeted API entity of the to-be-created documentation part. See below.
	Location DocumentationPartLocation `pulumi:"location"`
	// Content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
	Properties string `pulumi:"properties"`
	// ID of the associated Rest API
	RestApiId string `pulumi:"restApiId"`
}

// The set of arguments for constructing a DocumentationPart resource.
type DocumentationPartArgs struct {
	// Location of the targeted API entity of the to-be-created documentation part. See below.
	Location DocumentationPartLocationInput
	// Content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
	Properties pulumi.StringInput
	// ID of the associated Rest API
	RestApiId pulumi.StringInput
}

func (DocumentationPartArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*documentationPartArgs)(nil)).Elem()
}

type DocumentationPartInput interface {
	pulumi.Input

	ToDocumentationPartOutput() DocumentationPartOutput
	ToDocumentationPartOutputWithContext(ctx context.Context) DocumentationPartOutput
}

func (*DocumentationPart) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentationPart)(nil)).Elem()
}

func (i *DocumentationPart) ToDocumentationPartOutput() DocumentationPartOutput {
	return i.ToDocumentationPartOutputWithContext(context.Background())
}

func (i *DocumentationPart) ToDocumentationPartOutputWithContext(ctx context.Context) DocumentationPartOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentationPartOutput)
}

// DocumentationPartArrayInput is an input type that accepts DocumentationPartArray and DocumentationPartArrayOutput values.
// You can construct a concrete instance of `DocumentationPartArrayInput` via:
//
//	DocumentationPartArray{ DocumentationPartArgs{...} }
type DocumentationPartArrayInput interface {
	pulumi.Input

	ToDocumentationPartArrayOutput() DocumentationPartArrayOutput
	ToDocumentationPartArrayOutputWithContext(context.Context) DocumentationPartArrayOutput
}

type DocumentationPartArray []DocumentationPartInput

func (DocumentationPartArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentationPart)(nil)).Elem()
}

func (i DocumentationPartArray) ToDocumentationPartArrayOutput() DocumentationPartArrayOutput {
	return i.ToDocumentationPartArrayOutputWithContext(context.Background())
}

func (i DocumentationPartArray) ToDocumentationPartArrayOutputWithContext(ctx context.Context) DocumentationPartArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentationPartArrayOutput)
}

// DocumentationPartMapInput is an input type that accepts DocumentationPartMap and DocumentationPartMapOutput values.
// You can construct a concrete instance of `DocumentationPartMapInput` via:
//
//	DocumentationPartMap{ "key": DocumentationPartArgs{...} }
type DocumentationPartMapInput interface {
	pulumi.Input

	ToDocumentationPartMapOutput() DocumentationPartMapOutput
	ToDocumentationPartMapOutputWithContext(context.Context) DocumentationPartMapOutput
}

type DocumentationPartMap map[string]DocumentationPartInput

func (DocumentationPartMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentationPart)(nil)).Elem()
}

func (i DocumentationPartMap) ToDocumentationPartMapOutput() DocumentationPartMapOutput {
	return i.ToDocumentationPartMapOutputWithContext(context.Background())
}

func (i DocumentationPartMap) ToDocumentationPartMapOutputWithContext(ctx context.Context) DocumentationPartMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentationPartMapOutput)
}

type DocumentationPartOutput struct{ *pulumi.OutputState }

func (DocumentationPartOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentationPart)(nil)).Elem()
}

func (o DocumentationPartOutput) ToDocumentationPartOutput() DocumentationPartOutput {
	return o
}

func (o DocumentationPartOutput) ToDocumentationPartOutputWithContext(ctx context.Context) DocumentationPartOutput {
	return o
}

// Location of the targeted API entity of the to-be-created documentation part. See below.
func (o DocumentationPartOutput) Location() DocumentationPartLocationOutput {
	return o.ApplyT(func(v *DocumentationPart) DocumentationPartLocationOutput { return v.Location }).(DocumentationPartLocationOutput)
}

// Content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
func (o DocumentationPartOutput) Properties() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentationPart) pulumi.StringOutput { return v.Properties }).(pulumi.StringOutput)
}

// ID of the associated Rest API
func (o DocumentationPartOutput) RestApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentationPart) pulumi.StringOutput { return v.RestApiId }).(pulumi.StringOutput)
}

type DocumentationPartArrayOutput struct{ *pulumi.OutputState }

func (DocumentationPartArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentationPart)(nil)).Elem()
}

func (o DocumentationPartArrayOutput) ToDocumentationPartArrayOutput() DocumentationPartArrayOutput {
	return o
}

func (o DocumentationPartArrayOutput) ToDocumentationPartArrayOutputWithContext(ctx context.Context) DocumentationPartArrayOutput {
	return o
}

func (o DocumentationPartArrayOutput) Index(i pulumi.IntInput) DocumentationPartOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DocumentationPart {
		return vs[0].([]*DocumentationPart)[vs[1].(int)]
	}).(DocumentationPartOutput)
}

type DocumentationPartMapOutput struct{ *pulumi.OutputState }

func (DocumentationPartMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentationPart)(nil)).Elem()
}

func (o DocumentationPartMapOutput) ToDocumentationPartMapOutput() DocumentationPartMapOutput {
	return o
}

func (o DocumentationPartMapOutput) ToDocumentationPartMapOutputWithContext(ctx context.Context) DocumentationPartMapOutput {
	return o
}

func (o DocumentationPartMapOutput) MapIndex(k pulumi.StringInput) DocumentationPartOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DocumentationPart {
		return vs[0].(map[string]*DocumentationPart)[vs[1].(string)]
	}).(DocumentationPartOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentationPartInput)(nil)).Elem(), &DocumentationPart{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentationPartArrayInput)(nil)).Elem(), DocumentationPartArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentationPartMapInput)(nil)).Elem(), DocumentationPartMap{})
	pulumi.RegisterOutputType(DocumentationPartOutput{})
	pulumi.RegisterOutputType(DocumentationPartArrayOutput{})
	pulumi.RegisterOutputType(DocumentationPartMapOutput{})
}
