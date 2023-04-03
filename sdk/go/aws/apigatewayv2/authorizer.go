// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Amazon API Gateway Version 2 authorizer.
// More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api.html).
//
// ## Example Usage
// ### Basic WebSocket API
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/apigatewayv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := apigatewayv2.NewAuthorizer(ctx, "example", &apigatewayv2.AuthorizerArgs{
//				ApiId:          pulumi.Any(aws_apigatewayv2_api.Example.Id),
//				AuthorizerType: pulumi.String("REQUEST"),
//				AuthorizerUri:  pulumi.Any(aws_lambda_function.Example.Invoke_arn),
//				IdentitySources: pulumi.StringArray{
//					pulumi.String("route.request.header.Auth"),
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
// ### Basic HTTP API
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/apigatewayv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := apigatewayv2.NewAuthorizer(ctx, "example", &apigatewayv2.AuthorizerArgs{
//				ApiId:          pulumi.Any(aws_apigatewayv2_api.Example.Id),
//				AuthorizerType: pulumi.String("REQUEST"),
//				AuthorizerUri:  pulumi.Any(aws_lambda_function.Example.Invoke_arn),
//				IdentitySources: pulumi.StringArray{
//					pulumi.String("$request.header.Authorization"),
//				},
//				AuthorizerPayloadFormatVersion: pulumi.String("2.0"),
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
// `aws_apigatewayv2_authorizer` can be imported by using the API identifier and authorizer identifier, e.g.,
//
// ```sh
//
//	$ pulumi import aws:apigatewayv2/authorizer:Authorizer example aabbccddee/1122334
//
// ```
type Authorizer struct {
	pulumi.CustomResourceState

	// API identifier.
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// Required credentials as an IAM role for API Gateway to invoke the authorizer.
	// Supported only for `REQUEST` authorizers.
	AuthorizerCredentialsArn pulumi.StringPtrOutput `pulumi:"authorizerCredentialsArn"`
	// Format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers.
	// Valid values: `1.0`, `2.0`.
	AuthorizerPayloadFormatVersion pulumi.StringPtrOutput `pulumi:"authorizerPayloadFormatVersion"`
	// Time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled.
	// If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Defaults to `300`.
	// Supported only for HTTP API Lambda authorizers.
	AuthorizerResultTtlInSeconds pulumi.IntOutput `pulumi:"authorizerResultTtlInSeconds"`
	// Authorizer type. Valid values: `JWT`, `REQUEST`.
	// Specify `REQUEST` for a Lambda function using incoming request parameters.
	// For HTTP APIs, specify `JWT` to use JSON Web Tokens.
	AuthorizerType pulumi.StringOutput `pulumi:"authorizerType"`
	// Authorizer's Uniform Resource Identifier (URI).
	// For `REQUEST` authorizers this must be a well-formed Lambda function URI, such as the `invokeArn` attribute of the `lambda.Function` resource.
	// Supported only for `REQUEST` authorizers. Must be between 1 and 2048 characters in length.
	AuthorizerUri pulumi.StringPtrOutput `pulumi:"authorizerUri"`
	// Whether a Lambda authorizer returns a response in a simple format. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy.
	// Supported only for HTTP APIs.
	EnableSimpleResponses pulumi.BoolPtrOutput `pulumi:"enableSimpleResponses"`
	// Identity sources for which authorization is requested.
	// For `REQUEST` authorizers the value is a list of one or more mapping expressions of the specified request parameters.
	// For `JWT` authorizers the single entry specifies where to extract the JSON Web Token (JWT) from inbound requests.
	IdentitySources pulumi.StringArrayOutput `pulumi:"identitySources"`
	// Configuration of a JWT authorizer. Required for the `JWT` authorizer type.
	// Supported only for HTTP APIs.
	JwtConfiguration AuthorizerJwtConfigurationPtrOutput `pulumi:"jwtConfiguration"`
	// Name of the authorizer. Must be between 1 and 128 characters in length.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewAuthorizer registers a new resource with the given unique name, arguments, and options.
func NewAuthorizer(ctx *pulumi.Context,
	name string, args *AuthorizerArgs, opts ...pulumi.ResourceOption) (*Authorizer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.AuthorizerType == nil {
		return nil, errors.New("invalid value for required argument 'AuthorizerType'")
	}
	var resource Authorizer
	err := ctx.RegisterResource("aws:apigatewayv2/authorizer:Authorizer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthorizer gets an existing Authorizer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthorizer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthorizerState, opts ...pulumi.ResourceOption) (*Authorizer, error) {
	var resource Authorizer
	err := ctx.ReadResource("aws:apigatewayv2/authorizer:Authorizer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Authorizer resources.
type authorizerState struct {
	// API identifier.
	ApiId *string `pulumi:"apiId"`
	// Required credentials as an IAM role for API Gateway to invoke the authorizer.
	// Supported only for `REQUEST` authorizers.
	AuthorizerCredentialsArn *string `pulumi:"authorizerCredentialsArn"`
	// Format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers.
	// Valid values: `1.0`, `2.0`.
	AuthorizerPayloadFormatVersion *string `pulumi:"authorizerPayloadFormatVersion"`
	// Time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled.
	// If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Defaults to `300`.
	// Supported only for HTTP API Lambda authorizers.
	AuthorizerResultTtlInSeconds *int `pulumi:"authorizerResultTtlInSeconds"`
	// Authorizer type. Valid values: `JWT`, `REQUEST`.
	// Specify `REQUEST` for a Lambda function using incoming request parameters.
	// For HTTP APIs, specify `JWT` to use JSON Web Tokens.
	AuthorizerType *string `pulumi:"authorizerType"`
	// Authorizer's Uniform Resource Identifier (URI).
	// For `REQUEST` authorizers this must be a well-formed Lambda function URI, such as the `invokeArn` attribute of the `lambda.Function` resource.
	// Supported only for `REQUEST` authorizers. Must be between 1 and 2048 characters in length.
	AuthorizerUri *string `pulumi:"authorizerUri"`
	// Whether a Lambda authorizer returns a response in a simple format. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy.
	// Supported only for HTTP APIs.
	EnableSimpleResponses *bool `pulumi:"enableSimpleResponses"`
	// Identity sources for which authorization is requested.
	// For `REQUEST` authorizers the value is a list of one or more mapping expressions of the specified request parameters.
	// For `JWT` authorizers the single entry specifies where to extract the JSON Web Token (JWT) from inbound requests.
	IdentitySources []string `pulumi:"identitySources"`
	// Configuration of a JWT authorizer. Required for the `JWT` authorizer type.
	// Supported only for HTTP APIs.
	JwtConfiguration *AuthorizerJwtConfiguration `pulumi:"jwtConfiguration"`
	// Name of the authorizer. Must be between 1 and 128 characters in length.
	Name *string `pulumi:"name"`
}

type AuthorizerState struct {
	// API identifier.
	ApiId pulumi.StringPtrInput
	// Required credentials as an IAM role for API Gateway to invoke the authorizer.
	// Supported only for `REQUEST` authorizers.
	AuthorizerCredentialsArn pulumi.StringPtrInput
	// Format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers.
	// Valid values: `1.0`, `2.0`.
	AuthorizerPayloadFormatVersion pulumi.StringPtrInput
	// Time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled.
	// If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Defaults to `300`.
	// Supported only for HTTP API Lambda authorizers.
	AuthorizerResultTtlInSeconds pulumi.IntPtrInput
	// Authorizer type. Valid values: `JWT`, `REQUEST`.
	// Specify `REQUEST` for a Lambda function using incoming request parameters.
	// For HTTP APIs, specify `JWT` to use JSON Web Tokens.
	AuthorizerType pulumi.StringPtrInput
	// Authorizer's Uniform Resource Identifier (URI).
	// For `REQUEST` authorizers this must be a well-formed Lambda function URI, such as the `invokeArn` attribute of the `lambda.Function` resource.
	// Supported only for `REQUEST` authorizers. Must be between 1 and 2048 characters in length.
	AuthorizerUri pulumi.StringPtrInput
	// Whether a Lambda authorizer returns a response in a simple format. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy.
	// Supported only for HTTP APIs.
	EnableSimpleResponses pulumi.BoolPtrInput
	// Identity sources for which authorization is requested.
	// For `REQUEST` authorizers the value is a list of one or more mapping expressions of the specified request parameters.
	// For `JWT` authorizers the single entry specifies where to extract the JSON Web Token (JWT) from inbound requests.
	IdentitySources pulumi.StringArrayInput
	// Configuration of a JWT authorizer. Required for the `JWT` authorizer type.
	// Supported only for HTTP APIs.
	JwtConfiguration AuthorizerJwtConfigurationPtrInput
	// Name of the authorizer. Must be between 1 and 128 characters in length.
	Name pulumi.StringPtrInput
}

func (AuthorizerState) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizerState)(nil)).Elem()
}

type authorizerArgs struct {
	// API identifier.
	ApiId string `pulumi:"apiId"`
	// Required credentials as an IAM role for API Gateway to invoke the authorizer.
	// Supported only for `REQUEST` authorizers.
	AuthorizerCredentialsArn *string `pulumi:"authorizerCredentialsArn"`
	// Format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers.
	// Valid values: `1.0`, `2.0`.
	AuthorizerPayloadFormatVersion *string `pulumi:"authorizerPayloadFormatVersion"`
	// Time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled.
	// If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Defaults to `300`.
	// Supported only for HTTP API Lambda authorizers.
	AuthorizerResultTtlInSeconds *int `pulumi:"authorizerResultTtlInSeconds"`
	// Authorizer type. Valid values: `JWT`, `REQUEST`.
	// Specify `REQUEST` for a Lambda function using incoming request parameters.
	// For HTTP APIs, specify `JWT` to use JSON Web Tokens.
	AuthorizerType string `pulumi:"authorizerType"`
	// Authorizer's Uniform Resource Identifier (URI).
	// For `REQUEST` authorizers this must be a well-formed Lambda function URI, such as the `invokeArn` attribute of the `lambda.Function` resource.
	// Supported only for `REQUEST` authorizers. Must be between 1 and 2048 characters in length.
	AuthorizerUri *string `pulumi:"authorizerUri"`
	// Whether a Lambda authorizer returns a response in a simple format. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy.
	// Supported only for HTTP APIs.
	EnableSimpleResponses *bool `pulumi:"enableSimpleResponses"`
	// Identity sources for which authorization is requested.
	// For `REQUEST` authorizers the value is a list of one or more mapping expressions of the specified request parameters.
	// For `JWT` authorizers the single entry specifies where to extract the JSON Web Token (JWT) from inbound requests.
	IdentitySources []string `pulumi:"identitySources"`
	// Configuration of a JWT authorizer. Required for the `JWT` authorizer type.
	// Supported only for HTTP APIs.
	JwtConfiguration *AuthorizerJwtConfiguration `pulumi:"jwtConfiguration"`
	// Name of the authorizer. Must be between 1 and 128 characters in length.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Authorizer resource.
type AuthorizerArgs struct {
	// API identifier.
	ApiId pulumi.StringInput
	// Required credentials as an IAM role for API Gateway to invoke the authorizer.
	// Supported only for `REQUEST` authorizers.
	AuthorizerCredentialsArn pulumi.StringPtrInput
	// Format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers.
	// Valid values: `1.0`, `2.0`.
	AuthorizerPayloadFormatVersion pulumi.StringPtrInput
	// Time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled.
	// If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Defaults to `300`.
	// Supported only for HTTP API Lambda authorizers.
	AuthorizerResultTtlInSeconds pulumi.IntPtrInput
	// Authorizer type. Valid values: `JWT`, `REQUEST`.
	// Specify `REQUEST` for a Lambda function using incoming request parameters.
	// For HTTP APIs, specify `JWT` to use JSON Web Tokens.
	AuthorizerType pulumi.StringInput
	// Authorizer's Uniform Resource Identifier (URI).
	// For `REQUEST` authorizers this must be a well-formed Lambda function URI, such as the `invokeArn` attribute of the `lambda.Function` resource.
	// Supported only for `REQUEST` authorizers. Must be between 1 and 2048 characters in length.
	AuthorizerUri pulumi.StringPtrInput
	// Whether a Lambda authorizer returns a response in a simple format. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy.
	// Supported only for HTTP APIs.
	EnableSimpleResponses pulumi.BoolPtrInput
	// Identity sources for which authorization is requested.
	// For `REQUEST` authorizers the value is a list of one or more mapping expressions of the specified request parameters.
	// For `JWT` authorizers the single entry specifies where to extract the JSON Web Token (JWT) from inbound requests.
	IdentitySources pulumi.StringArrayInput
	// Configuration of a JWT authorizer. Required for the `JWT` authorizer type.
	// Supported only for HTTP APIs.
	JwtConfiguration AuthorizerJwtConfigurationPtrInput
	// Name of the authorizer. Must be between 1 and 128 characters in length.
	Name pulumi.StringPtrInput
}

func (AuthorizerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizerArgs)(nil)).Elem()
}

type AuthorizerInput interface {
	pulumi.Input

	ToAuthorizerOutput() AuthorizerOutput
	ToAuthorizerOutputWithContext(ctx context.Context) AuthorizerOutput
}

func (*Authorizer) ElementType() reflect.Type {
	return reflect.TypeOf((**Authorizer)(nil)).Elem()
}

func (i *Authorizer) ToAuthorizerOutput() AuthorizerOutput {
	return i.ToAuthorizerOutputWithContext(context.Background())
}

func (i *Authorizer) ToAuthorizerOutputWithContext(ctx context.Context) AuthorizerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizerOutput)
}

// AuthorizerArrayInput is an input type that accepts AuthorizerArray and AuthorizerArrayOutput values.
// You can construct a concrete instance of `AuthorizerArrayInput` via:
//
//	AuthorizerArray{ AuthorizerArgs{...} }
type AuthorizerArrayInput interface {
	pulumi.Input

	ToAuthorizerArrayOutput() AuthorizerArrayOutput
	ToAuthorizerArrayOutputWithContext(context.Context) AuthorizerArrayOutput
}

type AuthorizerArray []AuthorizerInput

func (AuthorizerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Authorizer)(nil)).Elem()
}

func (i AuthorizerArray) ToAuthorizerArrayOutput() AuthorizerArrayOutput {
	return i.ToAuthorizerArrayOutputWithContext(context.Background())
}

func (i AuthorizerArray) ToAuthorizerArrayOutputWithContext(ctx context.Context) AuthorizerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizerArrayOutput)
}

// AuthorizerMapInput is an input type that accepts AuthorizerMap and AuthorizerMapOutput values.
// You can construct a concrete instance of `AuthorizerMapInput` via:
//
//	AuthorizerMap{ "key": AuthorizerArgs{...} }
type AuthorizerMapInput interface {
	pulumi.Input

	ToAuthorizerMapOutput() AuthorizerMapOutput
	ToAuthorizerMapOutputWithContext(context.Context) AuthorizerMapOutput
}

type AuthorizerMap map[string]AuthorizerInput

func (AuthorizerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Authorizer)(nil)).Elem()
}

func (i AuthorizerMap) ToAuthorizerMapOutput() AuthorizerMapOutput {
	return i.ToAuthorizerMapOutputWithContext(context.Background())
}

func (i AuthorizerMap) ToAuthorizerMapOutputWithContext(ctx context.Context) AuthorizerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizerMapOutput)
}

type AuthorizerOutput struct{ *pulumi.OutputState }

func (AuthorizerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Authorizer)(nil)).Elem()
}

func (o AuthorizerOutput) ToAuthorizerOutput() AuthorizerOutput {
	return o
}

func (o AuthorizerOutput) ToAuthorizerOutputWithContext(ctx context.Context) AuthorizerOutput {
	return o
}

// API identifier.
func (o AuthorizerOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *Authorizer) pulumi.StringOutput { return v.ApiId }).(pulumi.StringOutput)
}

// Required credentials as an IAM role for API Gateway to invoke the authorizer.
// Supported only for `REQUEST` authorizers.
func (o AuthorizerOutput) AuthorizerCredentialsArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Authorizer) pulumi.StringPtrOutput { return v.AuthorizerCredentialsArn }).(pulumi.StringPtrOutput)
}

// Format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers.
// Valid values: `1.0`, `2.0`.
func (o AuthorizerOutput) AuthorizerPayloadFormatVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Authorizer) pulumi.StringPtrOutput { return v.AuthorizerPayloadFormatVersion }).(pulumi.StringPtrOutput)
}

// Time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled.
// If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Defaults to `300`.
// Supported only for HTTP API Lambda authorizers.
func (o AuthorizerOutput) AuthorizerResultTtlInSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v *Authorizer) pulumi.IntOutput { return v.AuthorizerResultTtlInSeconds }).(pulumi.IntOutput)
}

// Authorizer type. Valid values: `JWT`, `REQUEST`.
// Specify `REQUEST` for a Lambda function using incoming request parameters.
// For HTTP APIs, specify `JWT` to use JSON Web Tokens.
func (o AuthorizerOutput) AuthorizerType() pulumi.StringOutput {
	return o.ApplyT(func(v *Authorizer) pulumi.StringOutput { return v.AuthorizerType }).(pulumi.StringOutput)
}

// Authorizer's Uniform Resource Identifier (URI).
// For `REQUEST` authorizers this must be a well-formed Lambda function URI, such as the `invokeArn` attribute of the `lambda.Function` resource.
// Supported only for `REQUEST` authorizers. Must be between 1 and 2048 characters in length.
func (o AuthorizerOutput) AuthorizerUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Authorizer) pulumi.StringPtrOutput { return v.AuthorizerUri }).(pulumi.StringPtrOutput)
}

// Whether a Lambda authorizer returns a response in a simple format. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy.
// Supported only for HTTP APIs.
func (o AuthorizerOutput) EnableSimpleResponses() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Authorizer) pulumi.BoolPtrOutput { return v.EnableSimpleResponses }).(pulumi.BoolPtrOutput)
}

// Identity sources for which authorization is requested.
// For `REQUEST` authorizers the value is a list of one or more mapping expressions of the specified request parameters.
// For `JWT` authorizers the single entry specifies where to extract the JSON Web Token (JWT) from inbound requests.
func (o AuthorizerOutput) IdentitySources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Authorizer) pulumi.StringArrayOutput { return v.IdentitySources }).(pulumi.StringArrayOutput)
}

// Configuration of a JWT authorizer. Required for the `JWT` authorizer type.
// Supported only for HTTP APIs.
func (o AuthorizerOutput) JwtConfiguration() AuthorizerJwtConfigurationPtrOutput {
	return o.ApplyT(func(v *Authorizer) AuthorizerJwtConfigurationPtrOutput { return v.JwtConfiguration }).(AuthorizerJwtConfigurationPtrOutput)
}

// Name of the authorizer. Must be between 1 and 128 characters in length.
func (o AuthorizerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Authorizer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type AuthorizerArrayOutput struct{ *pulumi.OutputState }

func (AuthorizerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Authorizer)(nil)).Elem()
}

func (o AuthorizerArrayOutput) ToAuthorizerArrayOutput() AuthorizerArrayOutput {
	return o
}

func (o AuthorizerArrayOutput) ToAuthorizerArrayOutputWithContext(ctx context.Context) AuthorizerArrayOutput {
	return o
}

func (o AuthorizerArrayOutput) Index(i pulumi.IntInput) AuthorizerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Authorizer {
		return vs[0].([]*Authorizer)[vs[1].(int)]
	}).(AuthorizerOutput)
}

type AuthorizerMapOutput struct{ *pulumi.OutputState }

func (AuthorizerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Authorizer)(nil)).Elem()
}

func (o AuthorizerMapOutput) ToAuthorizerMapOutput() AuthorizerMapOutput {
	return o
}

func (o AuthorizerMapOutput) ToAuthorizerMapOutputWithContext(ctx context.Context) AuthorizerMapOutput {
	return o
}

func (o AuthorizerMapOutput) MapIndex(k pulumi.StringInput) AuthorizerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Authorizer {
		return vs[0].(map[string]*Authorizer)[vs[1].(string)]
	}).(AuthorizerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthorizerInput)(nil)).Elem(), &Authorizer{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthorizerArrayInput)(nil)).Elem(), AuthorizerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthorizerMapInput)(nil)).Elem(), AuthorizerMap{})
	pulumi.RegisterOutputType(AuthorizerOutput{})
	pulumi.RegisterOutputType(AuthorizerArrayOutput{})
	pulumi.RegisterOutputType(AuthorizerMapOutput{})
}
