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

// Provides an AppSync API Key.
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
//			exampleGraphQLApi, err := appsync.NewGraphQLApi(ctx, "exampleGraphQLApi", &appsync.GraphQLApiArgs{
//				AuthenticationType: pulumi.String("API_KEY"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = appsync.NewApiKey(ctx, "exampleApiKey", &appsync.ApiKeyArgs{
//				ApiId:   exampleGraphQLApi.ID(),
//				Expires: pulumi.String("2018-05-03T04:00:00Z"),
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
//	to = aws_appsync_api_key.example
//
//	id = "xxxxx:yyyyy" } Using `pulumi import`, import `aws_appsync_api_key` using the AppSync API ID and key separated by `:`. For exampleconsole % pulumi import aws_appsync_api_key.example xxxxx:yyyyy
type ApiKey struct {
	pulumi.CustomResourceState

	// ID of the associated AppSync API
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// API key description. Defaults to "Managed by Pulumi".
	Description pulumi.StringOutput `pulumi:"description"`
	// RFC3339 string representation of the expiry date. Rounded down to nearest hour. By default, it is 7 days from the date of creation.
	Expires pulumi.StringPtrOutput `pulumi:"expires"`
	// API key
	Key pulumi.StringOutput `pulumi:"key"`
}

// NewApiKey registers a new resource with the given unique name, arguments, and options.
func NewApiKey(ctx *pulumi.Context,
	name string, args *ApiKeyArgs, opts ...pulumi.ResourceOption) (*ApiKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.Description == nil {
		args.Description = pulumi.StringPtr("Managed by Pulumi")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"key",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApiKey
	err := ctx.RegisterResource("aws:appsync/apiKey:ApiKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiKey gets an existing ApiKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiKeyState, opts ...pulumi.ResourceOption) (*ApiKey, error) {
	var resource ApiKey
	err := ctx.ReadResource("aws:appsync/apiKey:ApiKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiKey resources.
type apiKeyState struct {
	// ID of the associated AppSync API
	ApiId *string `pulumi:"apiId"`
	// API key description. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// RFC3339 string representation of the expiry date. Rounded down to nearest hour. By default, it is 7 days from the date of creation.
	Expires *string `pulumi:"expires"`
	// API key
	Key *string `pulumi:"key"`
}

type ApiKeyState struct {
	// ID of the associated AppSync API
	ApiId pulumi.StringPtrInput
	// API key description. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// RFC3339 string representation of the expiry date. Rounded down to nearest hour. By default, it is 7 days from the date of creation.
	Expires pulumi.StringPtrInput
	// API key
	Key pulumi.StringPtrInput
}

func (ApiKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiKeyState)(nil)).Elem()
}

type apiKeyArgs struct {
	// ID of the associated AppSync API
	ApiId string `pulumi:"apiId"`
	// API key description. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// RFC3339 string representation of the expiry date. Rounded down to nearest hour. By default, it is 7 days from the date of creation.
	Expires *string `pulumi:"expires"`
}

// The set of arguments for constructing a ApiKey resource.
type ApiKeyArgs struct {
	// ID of the associated AppSync API
	ApiId pulumi.StringInput
	// API key description. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// RFC3339 string representation of the expiry date. Rounded down to nearest hour. By default, it is 7 days from the date of creation.
	Expires pulumi.StringPtrInput
}

func (ApiKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiKeyArgs)(nil)).Elem()
}

type ApiKeyInput interface {
	pulumi.Input

	ToApiKeyOutput() ApiKeyOutput
	ToApiKeyOutputWithContext(ctx context.Context) ApiKeyOutput
}

func (*ApiKey) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiKey)(nil)).Elem()
}

func (i *ApiKey) ToApiKeyOutput() ApiKeyOutput {
	return i.ToApiKeyOutputWithContext(context.Background())
}

func (i *ApiKey) ToApiKeyOutputWithContext(ctx context.Context) ApiKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiKeyOutput)
}

// ApiKeyArrayInput is an input type that accepts ApiKeyArray and ApiKeyArrayOutput values.
// You can construct a concrete instance of `ApiKeyArrayInput` via:
//
//	ApiKeyArray{ ApiKeyArgs{...} }
type ApiKeyArrayInput interface {
	pulumi.Input

	ToApiKeyArrayOutput() ApiKeyArrayOutput
	ToApiKeyArrayOutputWithContext(context.Context) ApiKeyArrayOutput
}

type ApiKeyArray []ApiKeyInput

func (ApiKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiKey)(nil)).Elem()
}

func (i ApiKeyArray) ToApiKeyArrayOutput() ApiKeyArrayOutput {
	return i.ToApiKeyArrayOutputWithContext(context.Background())
}

func (i ApiKeyArray) ToApiKeyArrayOutputWithContext(ctx context.Context) ApiKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiKeyArrayOutput)
}

// ApiKeyMapInput is an input type that accepts ApiKeyMap and ApiKeyMapOutput values.
// You can construct a concrete instance of `ApiKeyMapInput` via:
//
//	ApiKeyMap{ "key": ApiKeyArgs{...} }
type ApiKeyMapInput interface {
	pulumi.Input

	ToApiKeyMapOutput() ApiKeyMapOutput
	ToApiKeyMapOutputWithContext(context.Context) ApiKeyMapOutput
}

type ApiKeyMap map[string]ApiKeyInput

func (ApiKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiKey)(nil)).Elem()
}

func (i ApiKeyMap) ToApiKeyMapOutput() ApiKeyMapOutput {
	return i.ToApiKeyMapOutputWithContext(context.Background())
}

func (i ApiKeyMap) ToApiKeyMapOutputWithContext(ctx context.Context) ApiKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiKeyMapOutput)
}

type ApiKeyOutput struct{ *pulumi.OutputState }

func (ApiKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiKey)(nil)).Elem()
}

func (o ApiKeyOutput) ToApiKeyOutput() ApiKeyOutput {
	return o
}

func (o ApiKeyOutput) ToApiKeyOutputWithContext(ctx context.Context) ApiKeyOutput {
	return o
}

// ID of the associated AppSync API
func (o ApiKeyOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.StringOutput { return v.ApiId }).(pulumi.StringOutput)
}

// API key description. Defaults to "Managed by Pulumi".
func (o ApiKeyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// RFC3339 string representation of the expiry date. Rounded down to nearest hour. By default, it is 7 days from the date of creation.
func (o ApiKeyOutput) Expires() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.StringPtrOutput { return v.Expires }).(pulumi.StringPtrOutput)
}

// API key
func (o ApiKeyOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

type ApiKeyArrayOutput struct{ *pulumi.OutputState }

func (ApiKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiKey)(nil)).Elem()
}

func (o ApiKeyArrayOutput) ToApiKeyArrayOutput() ApiKeyArrayOutput {
	return o
}

func (o ApiKeyArrayOutput) ToApiKeyArrayOutputWithContext(ctx context.Context) ApiKeyArrayOutput {
	return o
}

func (o ApiKeyArrayOutput) Index(i pulumi.IntInput) ApiKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApiKey {
		return vs[0].([]*ApiKey)[vs[1].(int)]
	}).(ApiKeyOutput)
}

type ApiKeyMapOutput struct{ *pulumi.OutputState }

func (ApiKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiKey)(nil)).Elem()
}

func (o ApiKeyMapOutput) ToApiKeyMapOutput() ApiKeyMapOutput {
	return o
}

func (o ApiKeyMapOutput) ToApiKeyMapOutputWithContext(ctx context.Context) ApiKeyMapOutput {
	return o
}

func (o ApiKeyMapOutput) MapIndex(k pulumi.StringInput) ApiKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApiKey {
		return vs[0].(map[string]*ApiKey)[vs[1].(string)]
	}).(ApiKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiKeyInput)(nil)).Elem(), &ApiKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiKeyArrayInput)(nil)).Elem(), ApiKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiKeyMapInput)(nil)).Elem(), ApiKeyMap{})
	pulumi.RegisterOutputType(ApiKeyOutput{})
	pulumi.RegisterOutputType(ApiKeyArrayOutput{})
	pulumi.RegisterOutputType(ApiKeyMapOutput{})
}
