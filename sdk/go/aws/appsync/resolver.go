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

// Provides an AppSync Resolver.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appsync"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			testGraphQLApi, err := appsync.NewGraphQLApi(ctx, "testGraphQLApi", &appsync.GraphQLApiArgs{
//				AuthenticationType: pulumi.String("API_KEY"),
//				Schema: pulumi.String(`type Mutation {
//		putPost(id: ID!, title: String!): Post
//	}
//
//	type Post {
//		id: ID!
//		title: String!
//	}
//
//	type Query {
//		singlePost(id: ID!): Post
//	}
//
//	schema {
//		query: Query
//		mutation: Mutation
//	}
//
// `),
//
//			})
//			if err != nil {
//				return err
//			}
//			testDataSource, err := appsync.NewDataSource(ctx, "testDataSource", &appsync.DataSourceArgs{
//				ApiId: testGraphQLApi.ID(),
//				Name:  pulumi.String("my_example"),
//				Type:  pulumi.String("HTTP"),
//				HttpConfig: &appsync.DataSourceHttpConfigArgs{
//					Endpoint: pulumi.String("http://example.com"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = appsync.NewResolver(ctx, "testResolver", &appsync.ResolverArgs{
//				ApiId:      testGraphQLApi.ID(),
//				Field:      pulumi.String("singlePost"),
//				Type:       pulumi.String("Query"),
//				DataSource: testDataSource.Name,
//				RequestTemplate: pulumi.String(`{
//	    "version": "2018-05-29",
//	    "method": "GET",
//	    "resourcePath": "/",
//	    "params":{
//	        "headers": $utils.http.copyheaders($ctx.request.headers)
//	    }
//	}
//
// `),
//
//				ResponseTemplate: pulumi.String(`#if($ctx.result.statusCode == 200)
//	    $ctx.result.body
//
// #else
//
//	$utils.appendError($ctx.result.body, $ctx.result.statusCode)
//
// #end
// `),
//
//				CachingConfig: &appsync.ResolverCachingConfigArgs{
//					CachingKeys: pulumi.StringArray{
//						pulumi.String("$context.identity.sub"),
//						pulumi.String("$context.arguments.id"),
//					},
//					Ttl: pulumi.Int(60),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = appsync.NewResolver(ctx, "mutationPipelineTest", &appsync.ResolverArgs{
//				Type:             pulumi.String("Mutation"),
//				ApiId:            testGraphQLApi.ID(),
//				Field:            pulumi.String("pipelineTest"),
//				RequestTemplate:  pulumi.String("{}"),
//				ResponseTemplate: pulumi.String("$util.toJson($ctx.result)"),
//				Kind:             pulumi.String("PIPELINE"),
//				PipelineConfig: &appsync.ResolverPipelineConfigArgs{
//					Functions: pulumi.StringArray{
//						aws_appsync_function.Test1.Function_id,
//						aws_appsync_function.Test2.Function_id,
//						aws_appsync_function.Test3.Function_id,
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
// ### JS
//
// ```go
// package main
//
// import (
//
//	"os"
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appsync"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := os.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := appsync.NewResolver(ctx, "example", &appsync.ResolverArgs{
//				Type:  pulumi.String("Query"),
//				ApiId: pulumi.Any(aws_appsync_graphql_api.Test.Id),
//				Field: pulumi.String("pipelineTest"),
//				Kind:  pulumi.String("PIPELINE"),
//				Code:  readFileOrPanic("some-code-dir"),
//				Runtime: &appsync.ResolverRuntimeArgs{
//					Name:           pulumi.String("APPSYNC_JS"),
//					RuntimeVersion: pulumi.String("1.0.0"),
//				},
//				PipelineConfig: &appsync.ResolverPipelineConfigArgs{
//					Functions: pulumi.StringArray{
//						aws_appsync_function.Test.Function_id,
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
//	to = aws_appsync_resolver.example
//
//	id = "abcdef123456-exampleType-exampleField" } Using `pulumi import`, import `aws_appsync_resolver` using the `api_id`, a hyphen, `type`, a hypen and `field`. For exampleconsole % pulumi import aws_appsync_resolver.example abcdef123456-exampleType-exampleField
type Resolver struct {
	pulumi.CustomResourceState

	// API ID for the GraphQL API.
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// ARN
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The Caching Config. See Caching Config.
	CachingConfig ResolverCachingConfigPtrOutput `pulumi:"cachingConfig"`
	// The function code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
	Code pulumi.StringPtrOutput `pulumi:"code"`
	// Data source name.
	DataSource pulumi.StringPtrOutput `pulumi:"dataSource"`
	// Field name from the schema defined in the GraphQL API.
	Field pulumi.StringOutput `pulumi:"field"`
	// Resolver type. Valid values are `UNIT` and `PIPELINE`.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Maximum batching size for a resolver. Valid values are between `0` and `2000`.
	MaxBatchSize pulumi.IntPtrOutput `pulumi:"maxBatchSize"`
	// The caching configuration for the resolver. See Pipeline Config.
	PipelineConfig ResolverPipelineConfigPtrOutput `pulumi:"pipelineConfig"`
	// Request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
	RequestTemplate pulumi.StringPtrOutput `pulumi:"requestTemplate"`
	// Response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
	ResponseTemplate pulumi.StringPtrOutput `pulumi:"responseTemplate"`
	// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified. See Runtime.
	Runtime ResolverRuntimePtrOutput `pulumi:"runtime"`
	// Describes a Sync configuration for a resolver. See Sync Config.
	SyncConfig ResolverSyncConfigPtrOutput `pulumi:"syncConfig"`
	// Type name from the schema defined in the GraphQL API.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewResolver registers a new resource with the given unique name, arguments, and options.
func NewResolver(ctx *pulumi.Context,
	name string, args *ResolverArgs, opts ...pulumi.ResourceOption) (*Resolver, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.Field == nil {
		return nil, errors.New("invalid value for required argument 'Field'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Resolver
	err := ctx.RegisterResource("aws:appsync/resolver:Resolver", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolver gets an existing Resolver resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolver(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverState, opts ...pulumi.ResourceOption) (*Resolver, error) {
	var resource Resolver
	err := ctx.ReadResource("aws:appsync/resolver:Resolver", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Resolver resources.
type resolverState struct {
	// API ID for the GraphQL API.
	ApiId *string `pulumi:"apiId"`
	// ARN
	Arn *string `pulumi:"arn"`
	// The Caching Config. See Caching Config.
	CachingConfig *ResolverCachingConfig `pulumi:"cachingConfig"`
	// The function code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
	Code *string `pulumi:"code"`
	// Data source name.
	DataSource *string `pulumi:"dataSource"`
	// Field name from the schema defined in the GraphQL API.
	Field *string `pulumi:"field"`
	// Resolver type. Valid values are `UNIT` and `PIPELINE`.
	Kind *string `pulumi:"kind"`
	// Maximum batching size for a resolver. Valid values are between `0` and `2000`.
	MaxBatchSize *int `pulumi:"maxBatchSize"`
	// The caching configuration for the resolver. See Pipeline Config.
	PipelineConfig *ResolverPipelineConfig `pulumi:"pipelineConfig"`
	// Request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
	RequestTemplate *string `pulumi:"requestTemplate"`
	// Response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
	ResponseTemplate *string `pulumi:"responseTemplate"`
	// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified. See Runtime.
	Runtime *ResolverRuntime `pulumi:"runtime"`
	// Describes a Sync configuration for a resolver. See Sync Config.
	SyncConfig *ResolverSyncConfig `pulumi:"syncConfig"`
	// Type name from the schema defined in the GraphQL API.
	Type *string `pulumi:"type"`
}

type ResolverState struct {
	// API ID for the GraphQL API.
	ApiId pulumi.StringPtrInput
	// ARN
	Arn pulumi.StringPtrInput
	// The Caching Config. See Caching Config.
	CachingConfig ResolverCachingConfigPtrInput
	// The function code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
	Code pulumi.StringPtrInput
	// Data source name.
	DataSource pulumi.StringPtrInput
	// Field name from the schema defined in the GraphQL API.
	Field pulumi.StringPtrInput
	// Resolver type. Valid values are `UNIT` and `PIPELINE`.
	Kind pulumi.StringPtrInput
	// Maximum batching size for a resolver. Valid values are between `0` and `2000`.
	MaxBatchSize pulumi.IntPtrInput
	// The caching configuration for the resolver. See Pipeline Config.
	PipelineConfig ResolverPipelineConfigPtrInput
	// Request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
	RequestTemplate pulumi.StringPtrInput
	// Response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
	ResponseTemplate pulumi.StringPtrInput
	// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified. See Runtime.
	Runtime ResolverRuntimePtrInput
	// Describes a Sync configuration for a resolver. See Sync Config.
	SyncConfig ResolverSyncConfigPtrInput
	// Type name from the schema defined in the GraphQL API.
	Type pulumi.StringPtrInput
}

func (ResolverState) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverState)(nil)).Elem()
}

type resolverArgs struct {
	// API ID for the GraphQL API.
	ApiId string `pulumi:"apiId"`
	// The Caching Config. See Caching Config.
	CachingConfig *ResolverCachingConfig `pulumi:"cachingConfig"`
	// The function code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
	Code *string `pulumi:"code"`
	// Data source name.
	DataSource *string `pulumi:"dataSource"`
	// Field name from the schema defined in the GraphQL API.
	Field string `pulumi:"field"`
	// Resolver type. Valid values are `UNIT` and `PIPELINE`.
	Kind *string `pulumi:"kind"`
	// Maximum batching size for a resolver. Valid values are between `0` and `2000`.
	MaxBatchSize *int `pulumi:"maxBatchSize"`
	// The caching configuration for the resolver. See Pipeline Config.
	PipelineConfig *ResolverPipelineConfig `pulumi:"pipelineConfig"`
	// Request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
	RequestTemplate *string `pulumi:"requestTemplate"`
	// Response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
	ResponseTemplate *string `pulumi:"responseTemplate"`
	// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified. See Runtime.
	Runtime *ResolverRuntime `pulumi:"runtime"`
	// Describes a Sync configuration for a resolver. See Sync Config.
	SyncConfig *ResolverSyncConfig `pulumi:"syncConfig"`
	// Type name from the schema defined in the GraphQL API.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Resolver resource.
type ResolverArgs struct {
	// API ID for the GraphQL API.
	ApiId pulumi.StringInput
	// The Caching Config. See Caching Config.
	CachingConfig ResolverCachingConfigPtrInput
	// The function code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
	Code pulumi.StringPtrInput
	// Data source name.
	DataSource pulumi.StringPtrInput
	// Field name from the schema defined in the GraphQL API.
	Field pulumi.StringInput
	// Resolver type. Valid values are `UNIT` and `PIPELINE`.
	Kind pulumi.StringPtrInput
	// Maximum batching size for a resolver. Valid values are between `0` and `2000`.
	MaxBatchSize pulumi.IntPtrInput
	// The caching configuration for the resolver. See Pipeline Config.
	PipelineConfig ResolverPipelineConfigPtrInput
	// Request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
	RequestTemplate pulumi.StringPtrInput
	// Response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
	ResponseTemplate pulumi.StringPtrInput
	// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified. See Runtime.
	Runtime ResolverRuntimePtrInput
	// Describes a Sync configuration for a resolver. See Sync Config.
	SyncConfig ResolverSyncConfigPtrInput
	// Type name from the schema defined in the GraphQL API.
	Type pulumi.StringInput
}

func (ResolverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverArgs)(nil)).Elem()
}

type ResolverInput interface {
	pulumi.Input

	ToResolverOutput() ResolverOutput
	ToResolverOutputWithContext(ctx context.Context) ResolverOutput
}

func (*Resolver) ElementType() reflect.Type {
	return reflect.TypeOf((**Resolver)(nil)).Elem()
}

func (i *Resolver) ToResolverOutput() ResolverOutput {
	return i.ToResolverOutputWithContext(context.Background())
}

func (i *Resolver) ToResolverOutputWithContext(ctx context.Context) ResolverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverOutput)
}

// ResolverArrayInput is an input type that accepts ResolverArray and ResolverArrayOutput values.
// You can construct a concrete instance of `ResolverArrayInput` via:
//
//	ResolverArray{ ResolverArgs{...} }
type ResolverArrayInput interface {
	pulumi.Input

	ToResolverArrayOutput() ResolverArrayOutput
	ToResolverArrayOutputWithContext(context.Context) ResolverArrayOutput
}

type ResolverArray []ResolverInput

func (ResolverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Resolver)(nil)).Elem()
}

func (i ResolverArray) ToResolverArrayOutput() ResolverArrayOutput {
	return i.ToResolverArrayOutputWithContext(context.Background())
}

func (i ResolverArray) ToResolverArrayOutputWithContext(ctx context.Context) ResolverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverArrayOutput)
}

// ResolverMapInput is an input type that accepts ResolverMap and ResolverMapOutput values.
// You can construct a concrete instance of `ResolverMapInput` via:
//
//	ResolverMap{ "key": ResolverArgs{...} }
type ResolverMapInput interface {
	pulumi.Input

	ToResolverMapOutput() ResolverMapOutput
	ToResolverMapOutputWithContext(context.Context) ResolverMapOutput
}

type ResolverMap map[string]ResolverInput

func (ResolverMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Resolver)(nil)).Elem()
}

func (i ResolverMap) ToResolverMapOutput() ResolverMapOutput {
	return i.ToResolverMapOutputWithContext(context.Background())
}

func (i ResolverMap) ToResolverMapOutputWithContext(ctx context.Context) ResolverMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverMapOutput)
}

type ResolverOutput struct{ *pulumi.OutputState }

func (ResolverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Resolver)(nil)).Elem()
}

func (o ResolverOutput) ToResolverOutput() ResolverOutput {
	return o
}

func (o ResolverOutput) ToResolverOutputWithContext(ctx context.Context) ResolverOutput {
	return o
}

// API ID for the GraphQL API.
func (o ResolverOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *Resolver) pulumi.StringOutput { return v.ApiId }).(pulumi.StringOutput)
}

// ARN
func (o ResolverOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Resolver) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The Caching Config. See Caching Config.
func (o ResolverOutput) CachingConfig() ResolverCachingConfigPtrOutput {
	return o.ApplyT(func(v *Resolver) ResolverCachingConfigPtrOutput { return v.CachingConfig }).(ResolverCachingConfigPtrOutput)
}

// The function code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
func (o ResolverOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Resolver) pulumi.StringPtrOutput { return v.Code }).(pulumi.StringPtrOutput)
}

// Data source name.
func (o ResolverOutput) DataSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Resolver) pulumi.StringPtrOutput { return v.DataSource }).(pulumi.StringPtrOutput)
}

// Field name from the schema defined in the GraphQL API.
func (o ResolverOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v *Resolver) pulumi.StringOutput { return v.Field }).(pulumi.StringOutput)
}

// Resolver type. Valid values are `UNIT` and `PIPELINE`.
func (o ResolverOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Resolver) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Maximum batching size for a resolver. Valid values are between `0` and `2000`.
func (o ResolverOutput) MaxBatchSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Resolver) pulumi.IntPtrOutput { return v.MaxBatchSize }).(pulumi.IntPtrOutput)
}

// The caching configuration for the resolver. See Pipeline Config.
func (o ResolverOutput) PipelineConfig() ResolverPipelineConfigPtrOutput {
	return o.ApplyT(func(v *Resolver) ResolverPipelineConfigPtrOutput { return v.PipelineConfig }).(ResolverPipelineConfigPtrOutput)
}

// Request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
func (o ResolverOutput) RequestTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Resolver) pulumi.StringPtrOutput { return v.RequestTemplate }).(pulumi.StringPtrOutput)
}

// Response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver. Required for non-Lambda resolvers.
func (o ResolverOutput) ResponseTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Resolver) pulumi.StringPtrOutput { return v.ResponseTemplate }).(pulumi.StringPtrOutput)
}

// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified. See Runtime.
func (o ResolverOutput) Runtime() ResolverRuntimePtrOutput {
	return o.ApplyT(func(v *Resolver) ResolverRuntimePtrOutput { return v.Runtime }).(ResolverRuntimePtrOutput)
}

// Describes a Sync configuration for a resolver. See Sync Config.
func (o ResolverOutput) SyncConfig() ResolverSyncConfigPtrOutput {
	return o.ApplyT(func(v *Resolver) ResolverSyncConfigPtrOutput { return v.SyncConfig }).(ResolverSyncConfigPtrOutput)
}

// Type name from the schema defined in the GraphQL API.
func (o ResolverOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Resolver) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type ResolverArrayOutput struct{ *pulumi.OutputState }

func (ResolverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Resolver)(nil)).Elem()
}

func (o ResolverArrayOutput) ToResolverArrayOutput() ResolverArrayOutput {
	return o
}

func (o ResolverArrayOutput) ToResolverArrayOutputWithContext(ctx context.Context) ResolverArrayOutput {
	return o
}

func (o ResolverArrayOutput) Index(i pulumi.IntInput) ResolverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Resolver {
		return vs[0].([]*Resolver)[vs[1].(int)]
	}).(ResolverOutput)
}

type ResolverMapOutput struct{ *pulumi.OutputState }

func (ResolverMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Resolver)(nil)).Elem()
}

func (o ResolverMapOutput) ToResolverMapOutput() ResolverMapOutput {
	return o
}

func (o ResolverMapOutput) ToResolverMapOutputWithContext(ctx context.Context) ResolverMapOutput {
	return o
}

func (o ResolverMapOutput) MapIndex(k pulumi.StringInput) ResolverOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Resolver {
		return vs[0].(map[string]*Resolver)[vs[1].(string)]
	}).(ResolverOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverInput)(nil)).Elem(), &Resolver{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverArrayInput)(nil)).Elem(), ResolverArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverMapInput)(nil)).Elem(), ResolverMap{})
	pulumi.RegisterOutputType(ResolverOutput{})
	pulumi.RegisterOutputType(ResolverArrayOutput{})
	pulumi.RegisterOutputType(ResolverMapOutput{})
}
