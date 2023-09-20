// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manages an Amazon API Gateway Version 2 API.
//
// > **Note:** Amazon API Gateway Version 2 resources are used for creating and deploying WebSocket and HTTP APIs. To create and deploy REST APIs, use Amazon API Gateway Version 1 resources.
//
// ## Example Usage
// ### Basic WebSocket API
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigatewayv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := apigatewayv2.NewApi(ctx, "example", &apigatewayv2.ApiArgs{
//				ProtocolType:             pulumi.String("WEBSOCKET"),
//				RouteSelectionExpression: pulumi.String("$request.body.action"),
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
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigatewayv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := apigatewayv2.NewApi(ctx, "example", &apigatewayv2.ApiArgs{
//				ProtocolType: pulumi.String("HTTP"),
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
// Using `pulumi import`, import `aws_apigatewayv2_api` using the API identifier. For example:
//
// ```sh
//
//	$ pulumi import aws:apigatewayv2/api:Api example aabbccddee
//
// ```
type Api struct {
	pulumi.CustomResourceState

	// URI of the API, of the form `https://{api-id}.execute-api.{region}.amazonaws.com` for HTTP APIs and `wss://{api-id}.execute-api.{region}.amazonaws.com` for WebSocket APIs.
	ApiEndpoint pulumi.StringOutput `pulumi:"apiEndpoint"`
	// An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
	// Valid values: `$context.authorizer.usageIdentifierKey`, `$request.header.x-api-key`. Defaults to `$request.header.x-api-key`.
	// Applicable for WebSocket APIs.
	ApiKeySelectionExpression pulumi.StringPtrOutput `pulumi:"apiKeySelectionExpression"`
	// ARN of the API.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// An OpenAPI specification that defines the set of routes and integrations to create as part of the HTTP APIs. Supported only for HTTP APIs.
	Body pulumi.StringPtrOutput `pulumi:"body"`
	// Cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html). Applicable for HTTP APIs.
	CorsConfiguration ApiCorsConfigurationPtrOutput `pulumi:"corsConfiguration"`
	// Part of _quick create_. Specifies any credentials required for the integration. Applicable for HTTP APIs.
	CredentialsArn pulumi.StringPtrOutput `pulumi:"credentialsArn"`
	// Description of the API. Must be less than or equal to 1024 characters in length.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether clients can invoke the API by using the default `execute-api` endpoint.
	// By default, clients can invoke the API with the default `{api_id}.execute-api.{region}.amazonaws.com endpoint`.
	// To require that clients use a custom domain name to invoke the API, disable the default endpoint.
	DisableExecuteApiEndpoint pulumi.BoolPtrOutput `pulumi:"disableExecuteApiEndpoint"`
	// ARN prefix to be used in an `lambda.Permission`'s `sourceArn` attribute
	// or in an `iam.Policy` to authorize access to the [`@connections` API](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html).
	// See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-control-access-iam.html) for details.
	ExecutionArn pulumi.StringOutput `pulumi:"executionArn"`
	// Whether warnings should return an error while API Gateway is creating or updating the resource using an OpenAPI specification. Defaults to `false`. Applicable for HTTP APIs.
	FailOnWarnings pulumi.BoolPtrOutput `pulumi:"failOnWarnings"`
	// Name of the API. Must be less than or equal to 128 characters in length.
	Name pulumi.StringOutput `pulumi:"name"`
	// API protocol. Valid values: `HTTP`, `WEBSOCKET`.
	ProtocolType pulumi.StringOutput `pulumi:"protocolType"`
	// Part of _quick create_. Specifies any [route key](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html). Applicable for HTTP APIs.
	RouteKey pulumi.StringPtrOutput `pulumi:"routeKey"`
	// The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
	// Defaults to `$request.method $request.path`.
	RouteSelectionExpression pulumi.StringPtrOutput `pulumi:"routeSelectionExpression"`
	// Map of tags to assign to the API. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Part of _quick create_. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
	// For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
	// The type of the integration will be `HTTP_PROXY` or `AWS_PROXY`, respectively. Applicable for HTTP APIs.
	Target pulumi.StringPtrOutput `pulumi:"target"`
	// Version identifier for the API. Must be between 1 and 64 characters in length.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewApi registers a new resource with the given unique name, arguments, and options.
func NewApi(ctx *pulumi.Context,
	name string, args *ApiArgs, opts ...pulumi.ResourceOption) (*Api, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProtocolType == nil {
		return nil, errors.New("invalid value for required argument 'ProtocolType'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Api
	err := ctx.RegisterResource("aws:apigatewayv2/api:Api", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApi gets an existing Api resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiState, opts ...pulumi.ResourceOption) (*Api, error) {
	var resource Api
	err := ctx.ReadResource("aws:apigatewayv2/api:Api", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Api resources.
type apiState struct {
	// URI of the API, of the form `https://{api-id}.execute-api.{region}.amazonaws.com` for HTTP APIs and `wss://{api-id}.execute-api.{region}.amazonaws.com` for WebSocket APIs.
	ApiEndpoint *string `pulumi:"apiEndpoint"`
	// An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
	// Valid values: `$context.authorizer.usageIdentifierKey`, `$request.header.x-api-key`. Defaults to `$request.header.x-api-key`.
	// Applicable for WebSocket APIs.
	ApiKeySelectionExpression *string `pulumi:"apiKeySelectionExpression"`
	// ARN of the API.
	Arn *string `pulumi:"arn"`
	// An OpenAPI specification that defines the set of routes and integrations to create as part of the HTTP APIs. Supported only for HTTP APIs.
	Body *string `pulumi:"body"`
	// Cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html). Applicable for HTTP APIs.
	CorsConfiguration *ApiCorsConfiguration `pulumi:"corsConfiguration"`
	// Part of _quick create_. Specifies any credentials required for the integration. Applicable for HTTP APIs.
	CredentialsArn *string `pulumi:"credentialsArn"`
	// Description of the API. Must be less than or equal to 1024 characters in length.
	Description *string `pulumi:"description"`
	// Whether clients can invoke the API by using the default `execute-api` endpoint.
	// By default, clients can invoke the API with the default `{api_id}.execute-api.{region}.amazonaws.com endpoint`.
	// To require that clients use a custom domain name to invoke the API, disable the default endpoint.
	DisableExecuteApiEndpoint *bool `pulumi:"disableExecuteApiEndpoint"`
	// ARN prefix to be used in an `lambda.Permission`'s `sourceArn` attribute
	// or in an `iam.Policy` to authorize access to the [`@connections` API](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html).
	// See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-control-access-iam.html) for details.
	ExecutionArn *string `pulumi:"executionArn"`
	// Whether warnings should return an error while API Gateway is creating or updating the resource using an OpenAPI specification. Defaults to `false`. Applicable for HTTP APIs.
	FailOnWarnings *bool `pulumi:"failOnWarnings"`
	// Name of the API. Must be less than or equal to 128 characters in length.
	Name *string `pulumi:"name"`
	// API protocol. Valid values: `HTTP`, `WEBSOCKET`.
	ProtocolType *string `pulumi:"protocolType"`
	// Part of _quick create_. Specifies any [route key](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html). Applicable for HTTP APIs.
	RouteKey *string `pulumi:"routeKey"`
	// The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
	// Defaults to `$request.method $request.path`.
	RouteSelectionExpression *string `pulumi:"routeSelectionExpression"`
	// Map of tags to assign to the API. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Part of _quick create_. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
	// For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
	// The type of the integration will be `HTTP_PROXY` or `AWS_PROXY`, respectively. Applicable for HTTP APIs.
	Target *string `pulumi:"target"`
	// Version identifier for the API. Must be between 1 and 64 characters in length.
	Version *string `pulumi:"version"`
}

type ApiState struct {
	// URI of the API, of the form `https://{api-id}.execute-api.{region}.amazonaws.com` for HTTP APIs and `wss://{api-id}.execute-api.{region}.amazonaws.com` for WebSocket APIs.
	ApiEndpoint pulumi.StringPtrInput
	// An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
	// Valid values: `$context.authorizer.usageIdentifierKey`, `$request.header.x-api-key`. Defaults to `$request.header.x-api-key`.
	// Applicable for WebSocket APIs.
	ApiKeySelectionExpression pulumi.StringPtrInput
	// ARN of the API.
	Arn pulumi.StringPtrInput
	// An OpenAPI specification that defines the set of routes and integrations to create as part of the HTTP APIs. Supported only for HTTP APIs.
	Body pulumi.StringPtrInput
	// Cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html). Applicable for HTTP APIs.
	CorsConfiguration ApiCorsConfigurationPtrInput
	// Part of _quick create_. Specifies any credentials required for the integration. Applicable for HTTP APIs.
	CredentialsArn pulumi.StringPtrInput
	// Description of the API. Must be less than or equal to 1024 characters in length.
	Description pulumi.StringPtrInput
	// Whether clients can invoke the API by using the default `execute-api` endpoint.
	// By default, clients can invoke the API with the default `{api_id}.execute-api.{region}.amazonaws.com endpoint`.
	// To require that clients use a custom domain name to invoke the API, disable the default endpoint.
	DisableExecuteApiEndpoint pulumi.BoolPtrInput
	// ARN prefix to be used in an `lambda.Permission`'s `sourceArn` attribute
	// or in an `iam.Policy` to authorize access to the [`@connections` API](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html).
	// See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-control-access-iam.html) for details.
	ExecutionArn pulumi.StringPtrInput
	// Whether warnings should return an error while API Gateway is creating or updating the resource using an OpenAPI specification. Defaults to `false`. Applicable for HTTP APIs.
	FailOnWarnings pulumi.BoolPtrInput
	// Name of the API. Must be less than or equal to 128 characters in length.
	Name pulumi.StringPtrInput
	// API protocol. Valid values: `HTTP`, `WEBSOCKET`.
	ProtocolType pulumi.StringPtrInput
	// Part of _quick create_. Specifies any [route key](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html). Applicable for HTTP APIs.
	RouteKey pulumi.StringPtrInput
	// The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
	// Defaults to `$request.method $request.path`.
	RouteSelectionExpression pulumi.StringPtrInput
	// Map of tags to assign to the API. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// Part of _quick create_. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
	// For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
	// The type of the integration will be `HTTP_PROXY` or `AWS_PROXY`, respectively. Applicable for HTTP APIs.
	Target pulumi.StringPtrInput
	// Version identifier for the API. Must be between 1 and 64 characters in length.
	Version pulumi.StringPtrInput
}

func (ApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiState)(nil)).Elem()
}

type apiArgs struct {
	// An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
	// Valid values: `$context.authorizer.usageIdentifierKey`, `$request.header.x-api-key`. Defaults to `$request.header.x-api-key`.
	// Applicable for WebSocket APIs.
	ApiKeySelectionExpression *string `pulumi:"apiKeySelectionExpression"`
	// An OpenAPI specification that defines the set of routes and integrations to create as part of the HTTP APIs. Supported only for HTTP APIs.
	Body *string `pulumi:"body"`
	// Cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html). Applicable for HTTP APIs.
	CorsConfiguration *ApiCorsConfiguration `pulumi:"corsConfiguration"`
	// Part of _quick create_. Specifies any credentials required for the integration. Applicable for HTTP APIs.
	CredentialsArn *string `pulumi:"credentialsArn"`
	// Description of the API. Must be less than or equal to 1024 characters in length.
	Description *string `pulumi:"description"`
	// Whether clients can invoke the API by using the default `execute-api` endpoint.
	// By default, clients can invoke the API with the default `{api_id}.execute-api.{region}.amazonaws.com endpoint`.
	// To require that clients use a custom domain name to invoke the API, disable the default endpoint.
	DisableExecuteApiEndpoint *bool `pulumi:"disableExecuteApiEndpoint"`
	// Whether warnings should return an error while API Gateway is creating or updating the resource using an OpenAPI specification. Defaults to `false`. Applicable for HTTP APIs.
	FailOnWarnings *bool `pulumi:"failOnWarnings"`
	// Name of the API. Must be less than or equal to 128 characters in length.
	Name *string `pulumi:"name"`
	// API protocol. Valid values: `HTTP`, `WEBSOCKET`.
	ProtocolType string `pulumi:"protocolType"`
	// Part of _quick create_. Specifies any [route key](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html). Applicable for HTTP APIs.
	RouteKey *string `pulumi:"routeKey"`
	// The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
	// Defaults to `$request.method $request.path`.
	RouteSelectionExpression *string `pulumi:"routeSelectionExpression"`
	// Map of tags to assign to the API. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Part of _quick create_. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
	// For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
	// The type of the integration will be `HTTP_PROXY` or `AWS_PROXY`, respectively. Applicable for HTTP APIs.
	Target *string `pulumi:"target"`
	// Version identifier for the API. Must be between 1 and 64 characters in length.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Api resource.
type ApiArgs struct {
	// An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
	// Valid values: `$context.authorizer.usageIdentifierKey`, `$request.header.x-api-key`. Defaults to `$request.header.x-api-key`.
	// Applicable for WebSocket APIs.
	ApiKeySelectionExpression pulumi.StringPtrInput
	// An OpenAPI specification that defines the set of routes and integrations to create as part of the HTTP APIs. Supported only for HTTP APIs.
	Body pulumi.StringPtrInput
	// Cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html). Applicable for HTTP APIs.
	CorsConfiguration ApiCorsConfigurationPtrInput
	// Part of _quick create_. Specifies any credentials required for the integration. Applicable for HTTP APIs.
	CredentialsArn pulumi.StringPtrInput
	// Description of the API. Must be less than or equal to 1024 characters in length.
	Description pulumi.StringPtrInput
	// Whether clients can invoke the API by using the default `execute-api` endpoint.
	// By default, clients can invoke the API with the default `{api_id}.execute-api.{region}.amazonaws.com endpoint`.
	// To require that clients use a custom domain name to invoke the API, disable the default endpoint.
	DisableExecuteApiEndpoint pulumi.BoolPtrInput
	// Whether warnings should return an error while API Gateway is creating or updating the resource using an OpenAPI specification. Defaults to `false`. Applicable for HTTP APIs.
	FailOnWarnings pulumi.BoolPtrInput
	// Name of the API. Must be less than or equal to 128 characters in length.
	Name pulumi.StringPtrInput
	// API protocol. Valid values: `HTTP`, `WEBSOCKET`.
	ProtocolType pulumi.StringInput
	// Part of _quick create_. Specifies any [route key](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html). Applicable for HTTP APIs.
	RouteKey pulumi.StringPtrInput
	// The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
	// Defaults to `$request.method $request.path`.
	RouteSelectionExpression pulumi.StringPtrInput
	// Map of tags to assign to the API. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Part of _quick create_. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
	// For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
	// The type of the integration will be `HTTP_PROXY` or `AWS_PROXY`, respectively. Applicable for HTTP APIs.
	Target pulumi.StringPtrInput
	// Version identifier for the API. Must be between 1 and 64 characters in length.
	Version pulumi.StringPtrInput
}

func (ApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiArgs)(nil)).Elem()
}

type ApiInput interface {
	pulumi.Input

	ToApiOutput() ApiOutput
	ToApiOutputWithContext(ctx context.Context) ApiOutput
}

func (*Api) ElementType() reflect.Type {
	return reflect.TypeOf((**Api)(nil)).Elem()
}

func (i *Api) ToApiOutput() ApiOutput {
	return i.ToApiOutputWithContext(context.Background())
}

func (i *Api) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOutput)
}

func (i *Api) ToOutput(ctx context.Context) pulumix.Output[*Api] {
	return pulumix.Output[*Api]{
		OutputState: i.ToApiOutputWithContext(ctx).OutputState,
	}
}

// ApiArrayInput is an input type that accepts ApiArray and ApiArrayOutput values.
// You can construct a concrete instance of `ApiArrayInput` via:
//
//	ApiArray{ ApiArgs{...} }
type ApiArrayInput interface {
	pulumi.Input

	ToApiArrayOutput() ApiArrayOutput
	ToApiArrayOutputWithContext(context.Context) ApiArrayOutput
}

type ApiArray []ApiInput

func (ApiArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Api)(nil)).Elem()
}

func (i ApiArray) ToApiArrayOutput() ApiArrayOutput {
	return i.ToApiArrayOutputWithContext(context.Background())
}

func (i ApiArray) ToApiArrayOutputWithContext(ctx context.Context) ApiArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiArrayOutput)
}

func (i ApiArray) ToOutput(ctx context.Context) pulumix.Output[[]*Api] {
	return pulumix.Output[[]*Api]{
		OutputState: i.ToApiArrayOutputWithContext(ctx).OutputState,
	}
}

// ApiMapInput is an input type that accepts ApiMap and ApiMapOutput values.
// You can construct a concrete instance of `ApiMapInput` via:
//
//	ApiMap{ "key": ApiArgs{...} }
type ApiMapInput interface {
	pulumi.Input

	ToApiMapOutput() ApiMapOutput
	ToApiMapOutputWithContext(context.Context) ApiMapOutput
}

type ApiMap map[string]ApiInput

func (ApiMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Api)(nil)).Elem()
}

func (i ApiMap) ToApiMapOutput() ApiMapOutput {
	return i.ToApiMapOutputWithContext(context.Background())
}

func (i ApiMap) ToApiMapOutputWithContext(ctx context.Context) ApiMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiMapOutput)
}

func (i ApiMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Api] {
	return pulumix.Output[map[string]*Api]{
		OutputState: i.ToApiMapOutputWithContext(ctx).OutputState,
	}
}

type ApiOutput struct{ *pulumi.OutputState }

func (ApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Api)(nil)).Elem()
}

func (o ApiOutput) ToApiOutput() ApiOutput {
	return o
}

func (o ApiOutput) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return o
}

func (o ApiOutput) ToOutput(ctx context.Context) pulumix.Output[*Api] {
	return pulumix.Output[*Api]{
		OutputState: o.OutputState,
	}
}

// URI of the API, of the form `https://{api-id}.execute-api.{region}.amazonaws.com` for HTTP APIs and `wss://{api-id}.execute-api.{region}.amazonaws.com` for WebSocket APIs.
func (o ApiOutput) ApiEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.ApiEndpoint }).(pulumi.StringOutput)
}

// An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
// Valid values: `$context.authorizer.usageIdentifierKey`, `$request.header.x-api-key`. Defaults to `$request.header.x-api-key`.
// Applicable for WebSocket APIs.
func (o ApiOutput) ApiKeySelectionExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ApiKeySelectionExpression }).(pulumi.StringPtrOutput)
}

// ARN of the API.
func (o ApiOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// An OpenAPI specification that defines the set of routes and integrations to create as part of the HTTP APIs. Supported only for HTTP APIs.
func (o ApiOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.Body }).(pulumi.StringPtrOutput)
}

// Cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html). Applicable for HTTP APIs.
func (o ApiOutput) CorsConfiguration() ApiCorsConfigurationPtrOutput {
	return o.ApplyT(func(v *Api) ApiCorsConfigurationPtrOutput { return v.CorsConfiguration }).(ApiCorsConfigurationPtrOutput)
}

// Part of _quick create_. Specifies any credentials required for the integration. Applicable for HTTP APIs.
func (o ApiOutput) CredentialsArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.CredentialsArn }).(pulumi.StringPtrOutput)
}

// Description of the API. Must be less than or equal to 1024 characters in length.
func (o ApiOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether clients can invoke the API by using the default `execute-api` endpoint.
// By default, clients can invoke the API with the default `{api_id}.execute-api.{region}.amazonaws.com endpoint`.
// To require that clients use a custom domain name to invoke the API, disable the default endpoint.
func (o ApiOutput) DisableExecuteApiEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.BoolPtrOutput { return v.DisableExecuteApiEndpoint }).(pulumi.BoolPtrOutput)
}

// ARN prefix to be used in an `lambda.Permission`'s `sourceArn` attribute
// or in an `iam.Policy` to authorize access to the [`@connections` API](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html).
// See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-control-access-iam.html) for details.
func (o ApiOutput) ExecutionArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.ExecutionArn }).(pulumi.StringOutput)
}

// Whether warnings should return an error while API Gateway is creating or updating the resource using an OpenAPI specification. Defaults to `false`. Applicable for HTTP APIs.
func (o ApiOutput) FailOnWarnings() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.BoolPtrOutput { return v.FailOnWarnings }).(pulumi.BoolPtrOutput)
}

// Name of the API. Must be less than or equal to 128 characters in length.
func (o ApiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// API protocol. Valid values: `HTTP`, `WEBSOCKET`.
func (o ApiOutput) ProtocolType() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.ProtocolType }).(pulumi.StringOutput)
}

// Part of _quick create_. Specifies any [route key](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html). Applicable for HTTP APIs.
func (o ApiOutput) RouteKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.RouteKey }).(pulumi.StringPtrOutput)
}

// The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
// Defaults to `$request.method $request.path`.
func (o ApiOutput) RouteSelectionExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.RouteSelectionExpression }).(pulumi.StringPtrOutput)
}

// Map of tags to assign to the API. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ApiOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Api) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o ApiOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Api) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Part of _quick create_. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
// For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
// The type of the integration will be `HTTP_PROXY` or `AWS_PROXY`, respectively. Applicable for HTTP APIs.
func (o ApiOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.Target }).(pulumi.StringPtrOutput)
}

// Version identifier for the API. Must be between 1 and 64 characters in length.
func (o ApiOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

type ApiArrayOutput struct{ *pulumi.OutputState }

func (ApiArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Api)(nil)).Elem()
}

func (o ApiArrayOutput) ToApiArrayOutput() ApiArrayOutput {
	return o
}

func (o ApiArrayOutput) ToApiArrayOutputWithContext(ctx context.Context) ApiArrayOutput {
	return o
}

func (o ApiArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Api] {
	return pulumix.Output[[]*Api]{
		OutputState: o.OutputState,
	}
}

func (o ApiArrayOutput) Index(i pulumi.IntInput) ApiOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Api {
		return vs[0].([]*Api)[vs[1].(int)]
	}).(ApiOutput)
}

type ApiMapOutput struct{ *pulumi.OutputState }

func (ApiMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Api)(nil)).Elem()
}

func (o ApiMapOutput) ToApiMapOutput() ApiMapOutput {
	return o
}

func (o ApiMapOutput) ToApiMapOutputWithContext(ctx context.Context) ApiMapOutput {
	return o
}

func (o ApiMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Api] {
	return pulumix.Output[map[string]*Api]{
		OutputState: o.OutputState,
	}
}

func (o ApiMapOutput) MapIndex(k pulumi.StringInput) ApiOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Api {
		return vs[0].(map[string]*Api)[vs[1].(string)]
	}).(ApiOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiInput)(nil)).Elem(), &Api{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiArrayInput)(nil)).Elem(), ApiArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiMapInput)(nil)).Elem(), ApiMap{})
	pulumi.RegisterOutputType(ApiOutput{})
	pulumi.RegisterOutputType(ApiArrayOutput{})
	pulumi.RegisterOutputType(ApiMapOutput{})
}
