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

// Provides an HTTP Method Response for an API Gateway Resource.
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
//			myDemoAPI, err := apigateway.NewRestApi(ctx, "myDemoAPI", &apigateway.RestApiArgs{
//				Description: pulumi.String("This is my API for demonstration purposes"),
//			})
//			if err != nil {
//				return err
//			}
//			myDemoResource, err := apigateway.NewResource(ctx, "myDemoResource", &apigateway.ResourceArgs{
//				RestApi:  myDemoAPI.ID(),
//				ParentId: myDemoAPI.RootResourceId,
//				PathPart: pulumi.String("mydemoresource"),
//			})
//			if err != nil {
//				return err
//			}
//			myDemoMethod, err := apigateway.NewMethod(ctx, "myDemoMethod", &apigateway.MethodArgs{
//				RestApi:       myDemoAPI.ID(),
//				ResourceId:    myDemoResource.ID(),
//				HttpMethod:    pulumi.String("GET"),
//				Authorization: pulumi.String("NONE"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = apigateway.NewIntegration(ctx, "myDemoIntegration", &apigateway.IntegrationArgs{
//				RestApi:    myDemoAPI.ID(),
//				ResourceId: myDemoResource.ID(),
//				HttpMethod: myDemoMethod.HttpMethod,
//				Type:       pulumi.String("MOCK"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = apigateway.NewMethodResponse(ctx, "response200", &apigateway.MethodResponseArgs{
//				RestApi:    myDemoAPI.ID(),
//				ResourceId: myDemoResource.ID(),
//				HttpMethod: myDemoMethod.HttpMethod,
//				StatusCode: pulumi.String("200"),
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
//	to = aws_api_gateway_method_response.example
//
//	id = "12345abcde/67890fghij/GET/200" } Using `pulumi import`, import `aws_api_gateway_method_response` using `REST-API-ID/RESOURCE-ID/HTTP-METHOD/STATUS-CODE`. For exampleconsole % pulumi import aws_api_gateway_method_response.example 12345abcde/67890fghij/GET/200
type MethodResponse struct {
	pulumi.CustomResourceState

	// HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
	HttpMethod pulumi.StringOutput `pulumi:"httpMethod"`
	// API resource ID
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// Map of the API models used for the response's content type
	ResponseModels pulumi.StringMapOutput `pulumi:"responseModels"`
	// Map of response parameters that can be sent to the caller.
	// For example: `responseParameters = { "method.response.header.X-Some-Header" = true }`
	// would define that the header `X-Some-Header` can be provided on the response.
	ResponseParameters pulumi.BoolMapOutput `pulumi:"responseParameters"`
	// ID of the associated REST API
	RestApi pulumi.StringOutput `pulumi:"restApi"`
	// HTTP status code
	StatusCode pulumi.StringOutput `pulumi:"statusCode"`
}

// NewMethodResponse registers a new resource with the given unique name, arguments, and options.
func NewMethodResponse(ctx *pulumi.Context,
	name string, args *MethodResponseArgs, opts ...pulumi.ResourceOption) (*MethodResponse, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HttpMethod == nil {
		return nil, errors.New("invalid value for required argument 'HttpMethod'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.RestApi == nil {
		return nil, errors.New("invalid value for required argument 'RestApi'")
	}
	if args.StatusCode == nil {
		return nil, errors.New("invalid value for required argument 'StatusCode'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MethodResponse
	err := ctx.RegisterResource("aws:apigateway/methodResponse:MethodResponse", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMethodResponse gets an existing MethodResponse resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMethodResponse(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MethodResponseState, opts ...pulumi.ResourceOption) (*MethodResponse, error) {
	var resource MethodResponse
	err := ctx.ReadResource("aws:apigateway/methodResponse:MethodResponse", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MethodResponse resources.
type methodResponseState struct {
	// HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
	HttpMethod *string `pulumi:"httpMethod"`
	// API resource ID
	ResourceId *string `pulumi:"resourceId"`
	// Map of the API models used for the response's content type
	ResponseModels map[string]string `pulumi:"responseModels"`
	// Map of response parameters that can be sent to the caller.
	// For example: `responseParameters = { "method.response.header.X-Some-Header" = true }`
	// would define that the header `X-Some-Header` can be provided on the response.
	ResponseParameters map[string]bool `pulumi:"responseParameters"`
	// ID of the associated REST API
	RestApi interface{} `pulumi:"restApi"`
	// HTTP status code
	StatusCode *string `pulumi:"statusCode"`
}

type MethodResponseState struct {
	// HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
	HttpMethod pulumi.StringPtrInput
	// API resource ID
	ResourceId pulumi.StringPtrInput
	// Map of the API models used for the response's content type
	ResponseModels pulumi.StringMapInput
	// Map of response parameters that can be sent to the caller.
	// For example: `responseParameters = { "method.response.header.X-Some-Header" = true }`
	// would define that the header `X-Some-Header` can be provided on the response.
	ResponseParameters pulumi.BoolMapInput
	// ID of the associated REST API
	RestApi pulumi.Input
	// HTTP status code
	StatusCode pulumi.StringPtrInput
}

func (MethodResponseState) ElementType() reflect.Type {
	return reflect.TypeOf((*methodResponseState)(nil)).Elem()
}

type methodResponseArgs struct {
	// HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
	HttpMethod string `pulumi:"httpMethod"`
	// API resource ID
	ResourceId string `pulumi:"resourceId"`
	// Map of the API models used for the response's content type
	ResponseModels map[string]string `pulumi:"responseModels"`
	// Map of response parameters that can be sent to the caller.
	// For example: `responseParameters = { "method.response.header.X-Some-Header" = true }`
	// would define that the header `X-Some-Header` can be provided on the response.
	ResponseParameters map[string]bool `pulumi:"responseParameters"`
	// ID of the associated REST API
	RestApi interface{} `pulumi:"restApi"`
	// HTTP status code
	StatusCode string `pulumi:"statusCode"`
}

// The set of arguments for constructing a MethodResponse resource.
type MethodResponseArgs struct {
	// HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
	HttpMethod pulumi.StringInput
	// API resource ID
	ResourceId pulumi.StringInput
	// Map of the API models used for the response's content type
	ResponseModels pulumi.StringMapInput
	// Map of response parameters that can be sent to the caller.
	// For example: `responseParameters = { "method.response.header.X-Some-Header" = true }`
	// would define that the header `X-Some-Header` can be provided on the response.
	ResponseParameters pulumi.BoolMapInput
	// ID of the associated REST API
	RestApi pulumi.Input
	// HTTP status code
	StatusCode pulumi.StringInput
}

func (MethodResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*methodResponseArgs)(nil)).Elem()
}

type MethodResponseInput interface {
	pulumi.Input

	ToMethodResponseOutput() MethodResponseOutput
	ToMethodResponseOutputWithContext(ctx context.Context) MethodResponseOutput
}

func (*MethodResponse) ElementType() reflect.Type {
	return reflect.TypeOf((**MethodResponse)(nil)).Elem()
}

func (i *MethodResponse) ToMethodResponseOutput() MethodResponseOutput {
	return i.ToMethodResponseOutputWithContext(context.Background())
}

func (i *MethodResponse) ToMethodResponseOutputWithContext(ctx context.Context) MethodResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MethodResponseOutput)
}

// MethodResponseArrayInput is an input type that accepts MethodResponseArray and MethodResponseArrayOutput values.
// You can construct a concrete instance of `MethodResponseArrayInput` via:
//
//	MethodResponseArray{ MethodResponseArgs{...} }
type MethodResponseArrayInput interface {
	pulumi.Input

	ToMethodResponseArrayOutput() MethodResponseArrayOutput
	ToMethodResponseArrayOutputWithContext(context.Context) MethodResponseArrayOutput
}

type MethodResponseArray []MethodResponseInput

func (MethodResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MethodResponse)(nil)).Elem()
}

func (i MethodResponseArray) ToMethodResponseArrayOutput() MethodResponseArrayOutput {
	return i.ToMethodResponseArrayOutputWithContext(context.Background())
}

func (i MethodResponseArray) ToMethodResponseArrayOutputWithContext(ctx context.Context) MethodResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MethodResponseArrayOutput)
}

// MethodResponseMapInput is an input type that accepts MethodResponseMap and MethodResponseMapOutput values.
// You can construct a concrete instance of `MethodResponseMapInput` via:
//
//	MethodResponseMap{ "key": MethodResponseArgs{...} }
type MethodResponseMapInput interface {
	pulumi.Input

	ToMethodResponseMapOutput() MethodResponseMapOutput
	ToMethodResponseMapOutputWithContext(context.Context) MethodResponseMapOutput
}

type MethodResponseMap map[string]MethodResponseInput

func (MethodResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MethodResponse)(nil)).Elem()
}

func (i MethodResponseMap) ToMethodResponseMapOutput() MethodResponseMapOutput {
	return i.ToMethodResponseMapOutputWithContext(context.Background())
}

func (i MethodResponseMap) ToMethodResponseMapOutputWithContext(ctx context.Context) MethodResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MethodResponseMapOutput)
}

type MethodResponseOutput struct{ *pulumi.OutputState }

func (MethodResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MethodResponse)(nil)).Elem()
}

func (o MethodResponseOutput) ToMethodResponseOutput() MethodResponseOutput {
	return o
}

func (o MethodResponseOutput) ToMethodResponseOutputWithContext(ctx context.Context) MethodResponseOutput {
	return o
}

// HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
func (o MethodResponseOutput) HttpMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *MethodResponse) pulumi.StringOutput { return v.HttpMethod }).(pulumi.StringOutput)
}

// API resource ID
func (o MethodResponseOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *MethodResponse) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// Map of the API models used for the response's content type
func (o MethodResponseOutput) ResponseModels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MethodResponse) pulumi.StringMapOutput { return v.ResponseModels }).(pulumi.StringMapOutput)
}

// Map of response parameters that can be sent to the caller.
// For example: `responseParameters = { "method.response.header.X-Some-Header" = true }`
// would define that the header `X-Some-Header` can be provided on the response.
func (o MethodResponseOutput) ResponseParameters() pulumi.BoolMapOutput {
	return o.ApplyT(func(v *MethodResponse) pulumi.BoolMapOutput { return v.ResponseParameters }).(pulumi.BoolMapOutput)
}

// ID of the associated REST API
func (o MethodResponseOutput) RestApi() pulumi.StringOutput {
	return o.ApplyT(func(v *MethodResponse) pulumi.StringOutput { return v.RestApi }).(pulumi.StringOutput)
}

// HTTP status code
func (o MethodResponseOutput) StatusCode() pulumi.StringOutput {
	return o.ApplyT(func(v *MethodResponse) pulumi.StringOutput { return v.StatusCode }).(pulumi.StringOutput)
}

type MethodResponseArrayOutput struct{ *pulumi.OutputState }

func (MethodResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MethodResponse)(nil)).Elem()
}

func (o MethodResponseArrayOutput) ToMethodResponseArrayOutput() MethodResponseArrayOutput {
	return o
}

func (o MethodResponseArrayOutput) ToMethodResponseArrayOutputWithContext(ctx context.Context) MethodResponseArrayOutput {
	return o
}

func (o MethodResponseArrayOutput) Index(i pulumi.IntInput) MethodResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MethodResponse {
		return vs[0].([]*MethodResponse)[vs[1].(int)]
	}).(MethodResponseOutput)
}

type MethodResponseMapOutput struct{ *pulumi.OutputState }

func (MethodResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MethodResponse)(nil)).Elem()
}

func (o MethodResponseMapOutput) ToMethodResponseMapOutput() MethodResponseMapOutput {
	return o
}

func (o MethodResponseMapOutput) ToMethodResponseMapOutputWithContext(ctx context.Context) MethodResponseMapOutput {
	return o
}

func (o MethodResponseMapOutput) MapIndex(k pulumi.StringInput) MethodResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MethodResponse {
		return vs[0].(map[string]*MethodResponse)[vs[1].(string)]
	}).(MethodResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MethodResponseInput)(nil)).Elem(), &MethodResponse{})
	pulumi.RegisterInputType(reflect.TypeOf((*MethodResponseArrayInput)(nil)).Elem(), MethodResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MethodResponseMapInput)(nil)).Elem(), MethodResponseMap{})
	pulumi.RegisterOutputType(MethodResponseOutput{})
	pulumi.RegisterOutputType(MethodResponseArrayOutput{})
	pulumi.RegisterOutputType(MethodResponseMapOutput{})
}
