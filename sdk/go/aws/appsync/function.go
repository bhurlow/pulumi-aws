// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appsync

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an AppSync Function.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/appsync"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleGraphQLApi, err := appsync.NewGraphQLApi(ctx, "exampleGraphQLApi", &appsync.GraphQLApiArgs{
//				AuthenticationType: pulumi.String("API_KEY"),
//				Schema: pulumi.String(fmt.Sprintf(`type Mutation {
//	  putPost(id: ID!, title: String!): Post
//	}
//
//	type Post {
//	  id: ID!
//	  title: String!
//	}
//
//	type Query {
//	  singlePost(id: ID!): Post
//	}
//
//	schema {
//	  query: Query
//	  mutation: Mutation
//	}
//
// `)),
//
//			})
//			if err != nil {
//				return err
//			}
//			exampleDataSource, err := appsync.NewDataSource(ctx, "exampleDataSource", &appsync.DataSourceArgs{
//				ApiId: exampleGraphQLApi.ID(),
//				Name:  pulumi.String("example"),
//				Type:  pulumi.String("HTTP"),
//				HttpConfig: &appsync.DataSourceHttpConfigArgs{
//					Endpoint: pulumi.String("http://example.com"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = appsync.NewFunction(ctx, "exampleFunction", &appsync.FunctionArgs{
//				ApiId:      exampleGraphQLApi.ID(),
//				DataSource: exampleDataSource.Name,
//				Name:       pulumi.String("example"),
//				RequestMappingTemplate: pulumi.String(fmt.Sprintf(`{
//	    "version": "2018-05-29",
//	    "method": "GET",
//	    "resourcePath": "/",
//	    "params":{
//	        "headers": $utils.http.copyheaders($ctx.request.headers)
//	    }
//	}
//
// `)),
//
//				ResponseMappingTemplate: pulumi.String(fmt.Sprintf("#if($ctx.result.statusCode == 200)\n    $ctx.result.body\n#else\n    $utils.appendError($ctx.result.body, $ctx.result.statusCode)\n#end\n")),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### With Code
//
// ```go
// package main
//
// import (
//
//	"io/ioutil"
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/appsync"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := ioutil.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := appsync.NewFunction(ctx, "example", &appsync.FunctionArgs{
//				ApiId:      pulumi.Any(aws_appsync_graphql_api.Example.Id),
//				DataSource: pulumi.Any(aws_appsync_datasource.Example.Name),
//				Name:       pulumi.String("example"),
//				Code:       readFileOrPanic("some-code-dir"),
//				Runtime: &appsync.FunctionRuntimeArgs{
//					Name:           pulumi.String("APPSYNC_JS"),
//					RuntimeVersion: pulumi.String("1.0.0"),
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
// `aws_appsync_function` can be imported using the AppSync API ID and Function ID separated by `-`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:appsync/function:Function example xxxxx-yyyyy
//
// ```
type Function struct {
	pulumi.CustomResourceState

	// ID of the associated AppSync API.
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// ARN of the Function object.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The function code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
	Code pulumi.StringPtrOutput `pulumi:"code"`
	// Function data source name.
	DataSource pulumi.StringOutput `pulumi:"dataSource"`
	// Function description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Unique ID representing the Function object.
	FunctionId pulumi.StringOutput `pulumi:"functionId"`
	// Version of the request mapping template. Currently the supported value is `2018-05-29`. Does not apply when specifying `code`.
	FunctionVersion pulumi.StringOutput `pulumi:"functionVersion"`
	// Maximum batching size for a resolver. Valid values are between `0` and `2000`.
	MaxBatchSize pulumi.IntPtrOutput `pulumi:"maxBatchSize"`
	// The name of the runtime to use. Currently, the only allowed value is `APPSYNC_JS`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
	RequestMappingTemplate pulumi.StringPtrOutput `pulumi:"requestMappingTemplate"`
	// Function response mapping template.
	ResponseMappingTemplate pulumi.StringPtrOutput `pulumi:"responseMappingTemplate"`
	// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified. See Runtime.
	Runtime FunctionRuntimePtrOutput `pulumi:"runtime"`
	// Describes a Sync configuration for a resolver. See Sync Config.
	SyncConfig FunctionSyncConfigPtrOutput `pulumi:"syncConfig"`
}

// NewFunction registers a new resource with the given unique name, arguments, and options.
func NewFunction(ctx *pulumi.Context,
	name string, args *FunctionArgs, opts ...pulumi.ResourceOption) (*Function, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.DataSource == nil {
		return nil, errors.New("invalid value for required argument 'DataSource'")
	}
	var resource Function
	err := ctx.RegisterResource("aws:appsync/function:Function", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunction gets an existing Function resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionState, opts ...pulumi.ResourceOption) (*Function, error) {
	var resource Function
	err := ctx.ReadResource("aws:appsync/function:Function", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Function resources.
type functionState struct {
	// ID of the associated AppSync API.
	ApiId *string `pulumi:"apiId"`
	// ARN of the Function object.
	Arn *string `pulumi:"arn"`
	// The function code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
	Code *string `pulumi:"code"`
	// Function data source name.
	DataSource *string `pulumi:"dataSource"`
	// Function description.
	Description *string `pulumi:"description"`
	// Unique ID representing the Function object.
	FunctionId *string `pulumi:"functionId"`
	// Version of the request mapping template. Currently the supported value is `2018-05-29`. Does not apply when specifying `code`.
	FunctionVersion *string `pulumi:"functionVersion"`
	// Maximum batching size for a resolver. Valid values are between `0` and `2000`.
	MaxBatchSize *int `pulumi:"maxBatchSize"`
	// The name of the runtime to use. Currently, the only allowed value is `APPSYNC_JS`.
	Name *string `pulumi:"name"`
	// Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
	RequestMappingTemplate *string `pulumi:"requestMappingTemplate"`
	// Function response mapping template.
	ResponseMappingTemplate *string `pulumi:"responseMappingTemplate"`
	// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified. See Runtime.
	Runtime *FunctionRuntime `pulumi:"runtime"`
	// Describes a Sync configuration for a resolver. See Sync Config.
	SyncConfig *FunctionSyncConfig `pulumi:"syncConfig"`
}

type FunctionState struct {
	// ID of the associated AppSync API.
	ApiId pulumi.StringPtrInput
	// ARN of the Function object.
	Arn pulumi.StringPtrInput
	// The function code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
	Code pulumi.StringPtrInput
	// Function data source name.
	DataSource pulumi.StringPtrInput
	// Function description.
	Description pulumi.StringPtrInput
	// Unique ID representing the Function object.
	FunctionId pulumi.StringPtrInput
	// Version of the request mapping template. Currently the supported value is `2018-05-29`. Does not apply when specifying `code`.
	FunctionVersion pulumi.StringPtrInput
	// Maximum batching size for a resolver. Valid values are between `0` and `2000`.
	MaxBatchSize pulumi.IntPtrInput
	// The name of the runtime to use. Currently, the only allowed value is `APPSYNC_JS`.
	Name pulumi.StringPtrInput
	// Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
	RequestMappingTemplate pulumi.StringPtrInput
	// Function response mapping template.
	ResponseMappingTemplate pulumi.StringPtrInput
	// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified. See Runtime.
	Runtime FunctionRuntimePtrInput
	// Describes a Sync configuration for a resolver. See Sync Config.
	SyncConfig FunctionSyncConfigPtrInput
}

func (FunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionState)(nil)).Elem()
}

type functionArgs struct {
	// ID of the associated AppSync API.
	ApiId string `pulumi:"apiId"`
	// The function code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
	Code *string `pulumi:"code"`
	// Function data source name.
	DataSource string `pulumi:"dataSource"`
	// Function description.
	Description *string `pulumi:"description"`
	// Version of the request mapping template. Currently the supported value is `2018-05-29`. Does not apply when specifying `code`.
	FunctionVersion *string `pulumi:"functionVersion"`
	// Maximum batching size for a resolver. Valid values are between `0` and `2000`.
	MaxBatchSize *int `pulumi:"maxBatchSize"`
	// The name of the runtime to use. Currently, the only allowed value is `APPSYNC_JS`.
	Name *string `pulumi:"name"`
	// Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
	RequestMappingTemplate *string `pulumi:"requestMappingTemplate"`
	// Function response mapping template.
	ResponseMappingTemplate *string `pulumi:"responseMappingTemplate"`
	// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified. See Runtime.
	Runtime *FunctionRuntime `pulumi:"runtime"`
	// Describes a Sync configuration for a resolver. See Sync Config.
	SyncConfig *FunctionSyncConfig `pulumi:"syncConfig"`
}

// The set of arguments for constructing a Function resource.
type FunctionArgs struct {
	// ID of the associated AppSync API.
	ApiId pulumi.StringInput
	// The function code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
	Code pulumi.StringPtrInput
	// Function data source name.
	DataSource pulumi.StringInput
	// Function description.
	Description pulumi.StringPtrInput
	// Version of the request mapping template. Currently the supported value is `2018-05-29`. Does not apply when specifying `code`.
	FunctionVersion pulumi.StringPtrInput
	// Maximum batching size for a resolver. Valid values are between `0` and `2000`.
	MaxBatchSize pulumi.IntPtrInput
	// The name of the runtime to use. Currently, the only allowed value is `APPSYNC_JS`.
	Name pulumi.StringPtrInput
	// Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
	RequestMappingTemplate pulumi.StringPtrInput
	// Function response mapping template.
	ResponseMappingTemplate pulumi.StringPtrInput
	// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified. See Runtime.
	Runtime FunctionRuntimePtrInput
	// Describes a Sync configuration for a resolver. See Sync Config.
	SyncConfig FunctionSyncConfigPtrInput
}

func (FunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionArgs)(nil)).Elem()
}

type FunctionInput interface {
	pulumi.Input

	ToFunctionOutput() FunctionOutput
	ToFunctionOutputWithContext(ctx context.Context) FunctionOutput
}

func (*Function) ElementType() reflect.Type {
	return reflect.TypeOf((**Function)(nil)).Elem()
}

func (i *Function) ToFunctionOutput() FunctionOutput {
	return i.ToFunctionOutputWithContext(context.Background())
}

func (i *Function) ToFunctionOutputWithContext(ctx context.Context) FunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionOutput)
}

// FunctionArrayInput is an input type that accepts FunctionArray and FunctionArrayOutput values.
// You can construct a concrete instance of `FunctionArrayInput` via:
//
//	FunctionArray{ FunctionArgs{...} }
type FunctionArrayInput interface {
	pulumi.Input

	ToFunctionArrayOutput() FunctionArrayOutput
	ToFunctionArrayOutputWithContext(context.Context) FunctionArrayOutput
}

type FunctionArray []FunctionInput

func (FunctionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Function)(nil)).Elem()
}

func (i FunctionArray) ToFunctionArrayOutput() FunctionArrayOutput {
	return i.ToFunctionArrayOutputWithContext(context.Background())
}

func (i FunctionArray) ToFunctionArrayOutputWithContext(ctx context.Context) FunctionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionArrayOutput)
}

// FunctionMapInput is an input type that accepts FunctionMap and FunctionMapOutput values.
// You can construct a concrete instance of `FunctionMapInput` via:
//
//	FunctionMap{ "key": FunctionArgs{...} }
type FunctionMapInput interface {
	pulumi.Input

	ToFunctionMapOutput() FunctionMapOutput
	ToFunctionMapOutputWithContext(context.Context) FunctionMapOutput
}

type FunctionMap map[string]FunctionInput

func (FunctionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Function)(nil)).Elem()
}

func (i FunctionMap) ToFunctionMapOutput() FunctionMapOutput {
	return i.ToFunctionMapOutputWithContext(context.Background())
}

func (i FunctionMap) ToFunctionMapOutputWithContext(ctx context.Context) FunctionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionMapOutput)
}

type FunctionOutput struct{ *pulumi.OutputState }

func (FunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Function)(nil)).Elem()
}

func (o FunctionOutput) ToFunctionOutput() FunctionOutput {
	return o
}

func (o FunctionOutput) ToFunctionOutputWithContext(ctx context.Context) FunctionOutput {
	return o
}

// ID of the associated AppSync API.
func (o FunctionOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.ApiId }).(pulumi.StringOutput)
}

// ARN of the Function object.
func (o FunctionOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The function code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
func (o FunctionOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.StringPtrOutput { return v.Code }).(pulumi.StringPtrOutput)
}

// Function data source name.
func (o FunctionOutput) DataSource() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.DataSource }).(pulumi.StringOutput)
}

// Function description.
func (o FunctionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Unique ID representing the Function object.
func (o FunctionOutput) FunctionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.FunctionId }).(pulumi.StringOutput)
}

// Version of the request mapping template. Currently the supported value is `2018-05-29`. Does not apply when specifying `code`.
func (o FunctionOutput) FunctionVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.FunctionVersion }).(pulumi.StringOutput)
}

// Maximum batching size for a resolver. Valid values are between `0` and `2000`.
func (o FunctionOutput) MaxBatchSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.IntPtrOutput { return v.MaxBatchSize }).(pulumi.IntPtrOutput)
}

// The name of the runtime to use. Currently, the only allowed value is `APPSYNC_JS`.
func (o FunctionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
func (o FunctionOutput) RequestMappingTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.StringPtrOutput { return v.RequestMappingTemplate }).(pulumi.StringPtrOutput)
}

// Function response mapping template.
func (o FunctionOutput) ResponseMappingTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.StringPtrOutput { return v.ResponseMappingTemplate }).(pulumi.StringPtrOutput)
}

// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified. See Runtime.
func (o FunctionOutput) Runtime() FunctionRuntimePtrOutput {
	return o.ApplyT(func(v *Function) FunctionRuntimePtrOutput { return v.Runtime }).(FunctionRuntimePtrOutput)
}

// Describes a Sync configuration for a resolver. See Sync Config.
func (o FunctionOutput) SyncConfig() FunctionSyncConfigPtrOutput {
	return o.ApplyT(func(v *Function) FunctionSyncConfigPtrOutput { return v.SyncConfig }).(FunctionSyncConfigPtrOutput)
}

type FunctionArrayOutput struct{ *pulumi.OutputState }

func (FunctionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Function)(nil)).Elem()
}

func (o FunctionArrayOutput) ToFunctionArrayOutput() FunctionArrayOutput {
	return o
}

func (o FunctionArrayOutput) ToFunctionArrayOutputWithContext(ctx context.Context) FunctionArrayOutput {
	return o
}

func (o FunctionArrayOutput) Index(i pulumi.IntInput) FunctionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Function {
		return vs[0].([]*Function)[vs[1].(int)]
	}).(FunctionOutput)
}

type FunctionMapOutput struct{ *pulumi.OutputState }

func (FunctionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Function)(nil)).Elem()
}

func (o FunctionMapOutput) ToFunctionMapOutput() FunctionMapOutput {
	return o
}

func (o FunctionMapOutput) ToFunctionMapOutputWithContext(ctx context.Context) FunctionMapOutput {
	return o
}

func (o FunctionMapOutput) MapIndex(k pulumi.StringInput) FunctionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Function {
		return vs[0].(map[string]*Function)[vs[1].(string)]
	}).(FunctionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionInput)(nil)).Elem(), &Function{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionArrayInput)(nil)).Elem(), FunctionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionMapInput)(nil)).Elem(), FunctionMap{})
	pulumi.RegisterOutputType(FunctionOutput{})
	pulumi.RegisterOutputType(FunctionArrayOutput{})
	pulumi.RegisterOutputType(FunctionMapOutput{})
}
