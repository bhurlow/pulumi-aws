// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appsync

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GraphQLApi struct {
	pulumi.CustomResourceState

	// One or more additional authentication providers for the GraphqlApi. Defined below.
	AdditionalAuthenticationProviders GraphQLApiAdditionalAuthenticationProviderArrayOutput `pulumi:"additionalAuthenticationProviders"`
	// ARN
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`, `AWS_LAMBDA`
	AuthenticationType pulumi.StringOutput `pulumi:"authenticationType"`
	// Sets the value of the GraphQL API to enable (`ENABLED`) or disable (`DISABLED`) introspection. If no value is provided, the introspection configuration will be set to ENABLED by default. This field will produce an error if the operation attempts to use the introspection feature while this field is disabled. For more information about introspection, see [GraphQL introspection](https://graphql.org/learn/introspection/).
	IntrospectionConfig pulumi.StringPtrOutput `pulumi:"introspectionConfig"`
	// Nested argument containing Lambda authorizer configuration. Defined below.
	LambdaAuthorizerConfig GraphQLApiLambdaAuthorizerConfigPtrOutput `pulumi:"lambdaAuthorizerConfig"`
	// Nested argument containing logging configuration. Defined below.
	LogConfig GraphQLApiLogConfigPtrOutput `pulumi:"logConfig"`
	// User-supplied name for the GraphqlApi.
	Name pulumi.StringOutput `pulumi:"name"`
	// Nested argument containing OpenID Connect configuration. Defined below.
	OpenidConnectConfig GraphQLApiOpenidConnectConfigPtrOutput `pulumi:"openidConnectConfig"`
	// The maximum depth a query can have in a single request. Depth refers to the amount of nested levels allowed in the body of query. The default value is `0` (or unspecified), which indicates there's no depth limit. If you set a limit, it can be between `1` and `75` nested levels. This field will produce a limit error if the operation falls out of bounds.
	//
	// Note that fields can still be set to nullable or non-nullable. If a non-nullable field produces an error, the error will be thrown upwards to the first nullable field available.
	QueryDepthLimit pulumi.IntPtrOutput `pulumi:"queryDepthLimit"`
	// The maximum number of resolvers that can be invoked in a single request. The default value is `0` (or unspecified), which will set the limit to `10000`. When specified, the limit value can be between `1` and `10000`. This field will produce a limit error if the operation falls out of bounds.
	ResolverCountLimit pulumi.IntPtrOutput `pulumi:"resolverCountLimit"`
	// Schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
	Schema pulumi.StringPtrOutput `pulumi:"schema"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Map of URIs associated with the APIE.g., `uris["GRAPHQL"] = https://ID.appsync-api.REGION.amazonaws.com/graphql`
	Uris pulumi.StringMapOutput `pulumi:"uris"`
	// Amazon Cognito User Pool configuration. Defined below.
	UserPoolConfig GraphQLApiUserPoolConfigPtrOutput `pulumi:"userPoolConfig"`
	// Sets the value of the GraphQL API to public (`GLOBAL`) or private (`PRIVATE`). If no value is provided, the visibility will be set to `GLOBAL` by default. This value cannot be changed once the API has been created.
	Visibility pulumi.StringPtrOutput `pulumi:"visibility"`
	// Whether tracing with X-ray is enabled. Defaults to false.
	XrayEnabled pulumi.BoolPtrOutput `pulumi:"xrayEnabled"`
}

// NewGraphQLApi registers a new resource with the given unique name, arguments, and options.
func NewGraphQLApi(ctx *pulumi.Context,
	name string, args *GraphQLApiArgs, opts ...pulumi.ResourceOption) (*GraphQLApi, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthenticationType == nil {
		return nil, errors.New("invalid value for required argument 'AuthenticationType'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GraphQLApi
	err := ctx.RegisterResource("aws:appsync/graphQLApi:GraphQLApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGraphQLApi gets an existing GraphQLApi resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGraphQLApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GraphQLApiState, opts ...pulumi.ResourceOption) (*GraphQLApi, error) {
	var resource GraphQLApi
	err := ctx.ReadResource("aws:appsync/graphQLApi:GraphQLApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GraphQLApi resources.
type graphQLApiState struct {
	// One or more additional authentication providers for the GraphqlApi. Defined below.
	AdditionalAuthenticationProviders []GraphQLApiAdditionalAuthenticationProvider `pulumi:"additionalAuthenticationProviders"`
	// ARN
	Arn *string `pulumi:"arn"`
	// Authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`, `AWS_LAMBDA`
	AuthenticationType *string `pulumi:"authenticationType"`
	// Sets the value of the GraphQL API to enable (`ENABLED`) or disable (`DISABLED`) introspection. If no value is provided, the introspection configuration will be set to ENABLED by default. This field will produce an error if the operation attempts to use the introspection feature while this field is disabled. For more information about introspection, see [GraphQL introspection](https://graphql.org/learn/introspection/).
	IntrospectionConfig *string `pulumi:"introspectionConfig"`
	// Nested argument containing Lambda authorizer configuration. Defined below.
	LambdaAuthorizerConfig *GraphQLApiLambdaAuthorizerConfig `pulumi:"lambdaAuthorizerConfig"`
	// Nested argument containing logging configuration. Defined below.
	LogConfig *GraphQLApiLogConfig `pulumi:"logConfig"`
	// User-supplied name for the GraphqlApi.
	Name *string `pulumi:"name"`
	// Nested argument containing OpenID Connect configuration. Defined below.
	OpenidConnectConfig *GraphQLApiOpenidConnectConfig `pulumi:"openidConnectConfig"`
	// The maximum depth a query can have in a single request. Depth refers to the amount of nested levels allowed in the body of query. The default value is `0` (or unspecified), which indicates there's no depth limit. If you set a limit, it can be between `1` and `75` nested levels. This field will produce a limit error if the operation falls out of bounds.
	//
	// Note that fields can still be set to nullable or non-nullable. If a non-nullable field produces an error, the error will be thrown upwards to the first nullable field available.
	QueryDepthLimit *int `pulumi:"queryDepthLimit"`
	// The maximum number of resolvers that can be invoked in a single request. The default value is `0` (or unspecified), which will set the limit to `10000`. When specified, the limit value can be between `1` and `10000`. This field will produce a limit error if the operation falls out of bounds.
	ResolverCountLimit *int `pulumi:"resolverCountLimit"`
	// Schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
	Schema *string `pulumi:"schema"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Map of URIs associated with the APIE.g., `uris["GRAPHQL"] = https://ID.appsync-api.REGION.amazonaws.com/graphql`
	Uris map[string]string `pulumi:"uris"`
	// Amazon Cognito User Pool configuration. Defined below.
	UserPoolConfig *GraphQLApiUserPoolConfig `pulumi:"userPoolConfig"`
	// Sets the value of the GraphQL API to public (`GLOBAL`) or private (`PRIVATE`). If no value is provided, the visibility will be set to `GLOBAL` by default. This value cannot be changed once the API has been created.
	Visibility *string `pulumi:"visibility"`
	// Whether tracing with X-ray is enabled. Defaults to false.
	XrayEnabled *bool `pulumi:"xrayEnabled"`
}

type GraphQLApiState struct {
	// One or more additional authentication providers for the GraphqlApi. Defined below.
	AdditionalAuthenticationProviders GraphQLApiAdditionalAuthenticationProviderArrayInput
	// ARN
	Arn pulumi.StringPtrInput
	// Authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`, `AWS_LAMBDA`
	AuthenticationType pulumi.StringPtrInput
	// Sets the value of the GraphQL API to enable (`ENABLED`) or disable (`DISABLED`) introspection. If no value is provided, the introspection configuration will be set to ENABLED by default. This field will produce an error if the operation attempts to use the introspection feature while this field is disabled. For more information about introspection, see [GraphQL introspection](https://graphql.org/learn/introspection/).
	IntrospectionConfig pulumi.StringPtrInput
	// Nested argument containing Lambda authorizer configuration. Defined below.
	LambdaAuthorizerConfig GraphQLApiLambdaAuthorizerConfigPtrInput
	// Nested argument containing logging configuration. Defined below.
	LogConfig GraphQLApiLogConfigPtrInput
	// User-supplied name for the GraphqlApi.
	Name pulumi.StringPtrInput
	// Nested argument containing OpenID Connect configuration. Defined below.
	OpenidConnectConfig GraphQLApiOpenidConnectConfigPtrInput
	// The maximum depth a query can have in a single request. Depth refers to the amount of nested levels allowed in the body of query. The default value is `0` (or unspecified), which indicates there's no depth limit. If you set a limit, it can be between `1` and `75` nested levels. This field will produce a limit error if the operation falls out of bounds.
	//
	// Note that fields can still be set to nullable or non-nullable. If a non-nullable field produces an error, the error will be thrown upwards to the first nullable field available.
	QueryDepthLimit pulumi.IntPtrInput
	// The maximum number of resolvers that can be invoked in a single request. The default value is `0` (or unspecified), which will set the limit to `10000`. When specified, the limit value can be between `1` and `10000`. This field will produce a limit error if the operation falls out of bounds.
	ResolverCountLimit pulumi.IntPtrInput
	// Schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
	Schema pulumi.StringPtrInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// Map of URIs associated with the APIE.g., `uris["GRAPHQL"] = https://ID.appsync-api.REGION.amazonaws.com/graphql`
	Uris pulumi.StringMapInput
	// Amazon Cognito User Pool configuration. Defined below.
	UserPoolConfig GraphQLApiUserPoolConfigPtrInput
	// Sets the value of the GraphQL API to public (`GLOBAL`) or private (`PRIVATE`). If no value is provided, the visibility will be set to `GLOBAL` by default. This value cannot be changed once the API has been created.
	Visibility pulumi.StringPtrInput
	// Whether tracing with X-ray is enabled. Defaults to false.
	XrayEnabled pulumi.BoolPtrInput
}

func (GraphQLApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*graphQLApiState)(nil)).Elem()
}

type graphQLApiArgs struct {
	// One or more additional authentication providers for the GraphqlApi. Defined below.
	AdditionalAuthenticationProviders []GraphQLApiAdditionalAuthenticationProvider `pulumi:"additionalAuthenticationProviders"`
	// Authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`, `AWS_LAMBDA`
	AuthenticationType string `pulumi:"authenticationType"`
	// Sets the value of the GraphQL API to enable (`ENABLED`) or disable (`DISABLED`) introspection. If no value is provided, the introspection configuration will be set to ENABLED by default. This field will produce an error if the operation attempts to use the introspection feature while this field is disabled. For more information about introspection, see [GraphQL introspection](https://graphql.org/learn/introspection/).
	IntrospectionConfig *string `pulumi:"introspectionConfig"`
	// Nested argument containing Lambda authorizer configuration. Defined below.
	LambdaAuthorizerConfig *GraphQLApiLambdaAuthorizerConfig `pulumi:"lambdaAuthorizerConfig"`
	// Nested argument containing logging configuration. Defined below.
	LogConfig *GraphQLApiLogConfig `pulumi:"logConfig"`
	// User-supplied name for the GraphqlApi.
	Name *string `pulumi:"name"`
	// Nested argument containing OpenID Connect configuration. Defined below.
	OpenidConnectConfig *GraphQLApiOpenidConnectConfig `pulumi:"openidConnectConfig"`
	// The maximum depth a query can have in a single request. Depth refers to the amount of nested levels allowed in the body of query. The default value is `0` (or unspecified), which indicates there's no depth limit. If you set a limit, it can be between `1` and `75` nested levels. This field will produce a limit error if the operation falls out of bounds.
	//
	// Note that fields can still be set to nullable or non-nullable. If a non-nullable field produces an error, the error will be thrown upwards to the first nullable field available.
	QueryDepthLimit *int `pulumi:"queryDepthLimit"`
	// The maximum number of resolvers that can be invoked in a single request. The default value is `0` (or unspecified), which will set the limit to `10000`. When specified, the limit value can be between `1` and `10000`. This field will produce a limit error if the operation falls out of bounds.
	ResolverCountLimit *int `pulumi:"resolverCountLimit"`
	// Schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
	Schema *string `pulumi:"schema"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Amazon Cognito User Pool configuration. Defined below.
	UserPoolConfig *GraphQLApiUserPoolConfig `pulumi:"userPoolConfig"`
	// Sets the value of the GraphQL API to public (`GLOBAL`) or private (`PRIVATE`). If no value is provided, the visibility will be set to `GLOBAL` by default. This value cannot be changed once the API has been created.
	Visibility *string `pulumi:"visibility"`
	// Whether tracing with X-ray is enabled. Defaults to false.
	XrayEnabled *bool `pulumi:"xrayEnabled"`
}

// The set of arguments for constructing a GraphQLApi resource.
type GraphQLApiArgs struct {
	// One or more additional authentication providers for the GraphqlApi. Defined below.
	AdditionalAuthenticationProviders GraphQLApiAdditionalAuthenticationProviderArrayInput
	// Authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`, `AWS_LAMBDA`
	AuthenticationType pulumi.StringInput
	// Sets the value of the GraphQL API to enable (`ENABLED`) or disable (`DISABLED`) introspection. If no value is provided, the introspection configuration will be set to ENABLED by default. This field will produce an error if the operation attempts to use the introspection feature while this field is disabled. For more information about introspection, see [GraphQL introspection](https://graphql.org/learn/introspection/).
	IntrospectionConfig pulumi.StringPtrInput
	// Nested argument containing Lambda authorizer configuration. Defined below.
	LambdaAuthorizerConfig GraphQLApiLambdaAuthorizerConfigPtrInput
	// Nested argument containing logging configuration. Defined below.
	LogConfig GraphQLApiLogConfigPtrInput
	// User-supplied name for the GraphqlApi.
	Name pulumi.StringPtrInput
	// Nested argument containing OpenID Connect configuration. Defined below.
	OpenidConnectConfig GraphQLApiOpenidConnectConfigPtrInput
	// The maximum depth a query can have in a single request. Depth refers to the amount of nested levels allowed in the body of query. The default value is `0` (or unspecified), which indicates there's no depth limit. If you set a limit, it can be between `1` and `75` nested levels. This field will produce a limit error if the operation falls out of bounds.
	//
	// Note that fields can still be set to nullable or non-nullable. If a non-nullable field produces an error, the error will be thrown upwards to the first nullable field available.
	QueryDepthLimit pulumi.IntPtrInput
	// The maximum number of resolvers that can be invoked in a single request. The default value is `0` (or unspecified), which will set the limit to `10000`. When specified, the limit value can be between `1` and `10000`. This field will produce a limit error if the operation falls out of bounds.
	ResolverCountLimit pulumi.IntPtrInput
	// Schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
	Schema pulumi.StringPtrInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Amazon Cognito User Pool configuration. Defined below.
	UserPoolConfig GraphQLApiUserPoolConfigPtrInput
	// Sets the value of the GraphQL API to public (`GLOBAL`) or private (`PRIVATE`). If no value is provided, the visibility will be set to `GLOBAL` by default. This value cannot be changed once the API has been created.
	Visibility pulumi.StringPtrInput
	// Whether tracing with X-ray is enabled. Defaults to false.
	XrayEnabled pulumi.BoolPtrInput
}

func (GraphQLApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*graphQLApiArgs)(nil)).Elem()
}

type GraphQLApiInput interface {
	pulumi.Input

	ToGraphQLApiOutput() GraphQLApiOutput
	ToGraphQLApiOutputWithContext(ctx context.Context) GraphQLApiOutput
}

func (*GraphQLApi) ElementType() reflect.Type {
	return reflect.TypeOf((**GraphQLApi)(nil)).Elem()
}

func (i *GraphQLApi) ToGraphQLApiOutput() GraphQLApiOutput {
	return i.ToGraphQLApiOutputWithContext(context.Background())
}

func (i *GraphQLApi) ToGraphQLApiOutputWithContext(ctx context.Context) GraphQLApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphQLApiOutput)
}

// GraphQLApiArrayInput is an input type that accepts GraphQLApiArray and GraphQLApiArrayOutput values.
// You can construct a concrete instance of `GraphQLApiArrayInput` via:
//
//	GraphQLApiArray{ GraphQLApiArgs{...} }
type GraphQLApiArrayInput interface {
	pulumi.Input

	ToGraphQLApiArrayOutput() GraphQLApiArrayOutput
	ToGraphQLApiArrayOutputWithContext(context.Context) GraphQLApiArrayOutput
}

type GraphQLApiArray []GraphQLApiInput

func (GraphQLApiArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GraphQLApi)(nil)).Elem()
}

func (i GraphQLApiArray) ToGraphQLApiArrayOutput() GraphQLApiArrayOutput {
	return i.ToGraphQLApiArrayOutputWithContext(context.Background())
}

func (i GraphQLApiArray) ToGraphQLApiArrayOutputWithContext(ctx context.Context) GraphQLApiArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphQLApiArrayOutput)
}

// GraphQLApiMapInput is an input type that accepts GraphQLApiMap and GraphQLApiMapOutput values.
// You can construct a concrete instance of `GraphQLApiMapInput` via:
//
//	GraphQLApiMap{ "key": GraphQLApiArgs{...} }
type GraphQLApiMapInput interface {
	pulumi.Input

	ToGraphQLApiMapOutput() GraphQLApiMapOutput
	ToGraphQLApiMapOutputWithContext(context.Context) GraphQLApiMapOutput
}

type GraphQLApiMap map[string]GraphQLApiInput

func (GraphQLApiMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GraphQLApi)(nil)).Elem()
}

func (i GraphQLApiMap) ToGraphQLApiMapOutput() GraphQLApiMapOutput {
	return i.ToGraphQLApiMapOutputWithContext(context.Background())
}

func (i GraphQLApiMap) ToGraphQLApiMapOutputWithContext(ctx context.Context) GraphQLApiMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphQLApiMapOutput)
}

type GraphQLApiOutput struct{ *pulumi.OutputState }

func (GraphQLApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GraphQLApi)(nil)).Elem()
}

func (o GraphQLApiOutput) ToGraphQLApiOutput() GraphQLApiOutput {
	return o
}

func (o GraphQLApiOutput) ToGraphQLApiOutputWithContext(ctx context.Context) GraphQLApiOutput {
	return o
}

// One or more additional authentication providers for the GraphqlApi. Defined below.
func (o GraphQLApiOutput) AdditionalAuthenticationProviders() GraphQLApiAdditionalAuthenticationProviderArrayOutput {
	return o.ApplyT(func(v *GraphQLApi) GraphQLApiAdditionalAuthenticationProviderArrayOutput {
		return v.AdditionalAuthenticationProviders
	}).(GraphQLApiAdditionalAuthenticationProviderArrayOutput)
}

// ARN
func (o GraphQLApiOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *GraphQLApi) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`, `AWS_LAMBDA`
func (o GraphQLApiOutput) AuthenticationType() pulumi.StringOutput {
	return o.ApplyT(func(v *GraphQLApi) pulumi.StringOutput { return v.AuthenticationType }).(pulumi.StringOutput)
}

// Sets the value of the GraphQL API to enable (`ENABLED`) or disable (`DISABLED`) introspection. If no value is provided, the introspection configuration will be set to ENABLED by default. This field will produce an error if the operation attempts to use the introspection feature while this field is disabled. For more information about introspection, see [GraphQL introspection](https://graphql.org/learn/introspection/).
func (o GraphQLApiOutput) IntrospectionConfig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GraphQLApi) pulumi.StringPtrOutput { return v.IntrospectionConfig }).(pulumi.StringPtrOutput)
}

// Nested argument containing Lambda authorizer configuration. Defined below.
func (o GraphQLApiOutput) LambdaAuthorizerConfig() GraphQLApiLambdaAuthorizerConfigPtrOutput {
	return o.ApplyT(func(v *GraphQLApi) GraphQLApiLambdaAuthorizerConfigPtrOutput { return v.LambdaAuthorizerConfig }).(GraphQLApiLambdaAuthorizerConfigPtrOutput)
}

// Nested argument containing logging configuration. Defined below.
func (o GraphQLApiOutput) LogConfig() GraphQLApiLogConfigPtrOutput {
	return o.ApplyT(func(v *GraphQLApi) GraphQLApiLogConfigPtrOutput { return v.LogConfig }).(GraphQLApiLogConfigPtrOutput)
}

// User-supplied name for the GraphqlApi.
func (o GraphQLApiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GraphQLApi) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Nested argument containing OpenID Connect configuration. Defined below.
func (o GraphQLApiOutput) OpenidConnectConfig() GraphQLApiOpenidConnectConfigPtrOutput {
	return o.ApplyT(func(v *GraphQLApi) GraphQLApiOpenidConnectConfigPtrOutput { return v.OpenidConnectConfig }).(GraphQLApiOpenidConnectConfigPtrOutput)
}

// The maximum depth a query can have in a single request. Depth refers to the amount of nested levels allowed in the body of query. The default value is `0` (or unspecified), which indicates there's no depth limit. If you set a limit, it can be between `1` and `75` nested levels. This field will produce a limit error if the operation falls out of bounds.
//
// Note that fields can still be set to nullable or non-nullable. If a non-nullable field produces an error, the error will be thrown upwards to the first nullable field available.
func (o GraphQLApiOutput) QueryDepthLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GraphQLApi) pulumi.IntPtrOutput { return v.QueryDepthLimit }).(pulumi.IntPtrOutput)
}

// The maximum number of resolvers that can be invoked in a single request. The default value is `0` (or unspecified), which will set the limit to `10000`. When specified, the limit value can be between `1` and `10000`. This field will produce a limit error if the operation falls out of bounds.
func (o GraphQLApiOutput) ResolverCountLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GraphQLApi) pulumi.IntPtrOutput { return v.ResolverCountLimit }).(pulumi.IntPtrOutput)
}

// Schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
func (o GraphQLApiOutput) Schema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GraphQLApi) pulumi.StringPtrOutput { return v.Schema }).(pulumi.StringPtrOutput)
}

// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o GraphQLApiOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GraphQLApi) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o GraphQLApiOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GraphQLApi) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Map of URIs associated with the APIE.g., `uris["GRAPHQL"] = https://ID.appsync-api.REGION.amazonaws.com/graphql`
func (o GraphQLApiOutput) Uris() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GraphQLApi) pulumi.StringMapOutput { return v.Uris }).(pulumi.StringMapOutput)
}

// Amazon Cognito User Pool configuration. Defined below.
func (o GraphQLApiOutput) UserPoolConfig() GraphQLApiUserPoolConfigPtrOutput {
	return o.ApplyT(func(v *GraphQLApi) GraphQLApiUserPoolConfigPtrOutput { return v.UserPoolConfig }).(GraphQLApiUserPoolConfigPtrOutput)
}

// Sets the value of the GraphQL API to public (`GLOBAL`) or private (`PRIVATE`). If no value is provided, the visibility will be set to `GLOBAL` by default. This value cannot be changed once the API has been created.
func (o GraphQLApiOutput) Visibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GraphQLApi) pulumi.StringPtrOutput { return v.Visibility }).(pulumi.StringPtrOutput)
}

// Whether tracing with X-ray is enabled. Defaults to false.
func (o GraphQLApiOutput) XrayEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GraphQLApi) pulumi.BoolPtrOutput { return v.XrayEnabled }).(pulumi.BoolPtrOutput)
}

type GraphQLApiArrayOutput struct{ *pulumi.OutputState }

func (GraphQLApiArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GraphQLApi)(nil)).Elem()
}

func (o GraphQLApiArrayOutput) ToGraphQLApiArrayOutput() GraphQLApiArrayOutput {
	return o
}

func (o GraphQLApiArrayOutput) ToGraphQLApiArrayOutputWithContext(ctx context.Context) GraphQLApiArrayOutput {
	return o
}

func (o GraphQLApiArrayOutput) Index(i pulumi.IntInput) GraphQLApiOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GraphQLApi {
		return vs[0].([]*GraphQLApi)[vs[1].(int)]
	}).(GraphQLApiOutput)
}

type GraphQLApiMapOutput struct{ *pulumi.OutputState }

func (GraphQLApiMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GraphQLApi)(nil)).Elem()
}

func (o GraphQLApiMapOutput) ToGraphQLApiMapOutput() GraphQLApiMapOutput {
	return o
}

func (o GraphQLApiMapOutput) ToGraphQLApiMapOutputWithContext(ctx context.Context) GraphQLApiMapOutput {
	return o
}

func (o GraphQLApiMapOutput) MapIndex(k pulumi.StringInput) GraphQLApiOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GraphQLApi {
		return vs[0].(map[string]*GraphQLApi)[vs[1].(string)]
	}).(GraphQLApiOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GraphQLApiInput)(nil)).Elem(), &GraphQLApi{})
	pulumi.RegisterInputType(reflect.TypeOf((*GraphQLApiArrayInput)(nil)).Elem(), GraphQLApiArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GraphQLApiMapInput)(nil)).Elem(), GraphQLApiMap{})
	pulumi.RegisterOutputType(GraphQLApiOutput{})
	pulumi.RegisterOutputType(GraphQLApiArrayOutput{})
	pulumi.RegisterOutputType(GraphQLApiMapOutput{})
}
