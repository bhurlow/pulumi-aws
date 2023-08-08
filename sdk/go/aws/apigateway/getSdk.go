// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
//			_, err := apigateway.GetSdk(ctx, &apigateway.GetSdkArgs{
//				RestApiId: aws_api_gateway_stage.Example.Rest_api_id,
//				StageName: aws_api_gateway_stage.Example.Stage_name,
//				SdkType:   "android",
//				Parameters: map[string]interface{}{
//					"groupId":         "example",
//					"artifactId":      "example",
//					"artifactVersion": "example",
//					"invokerPackage":  "example",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetSdk(ctx *pulumi.Context, args *GetSdkArgs, opts ...pulumi.InvokeOption) (*GetSdkResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSdkResult
	err := ctx.Invoke("aws:apigateway/getSdk:getSdk", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSdk.
type GetSdkArgs struct {
	// Key-value map of query string parameters `sdkType` properties of the SDK. For SDK Type of `objectivec` or `swift`, a parameter named `classPrefix` is required. For SDK Type of `android`, parameters named `groupId`, `artifactId`, `artifactVersion`, and `invokerPackage` are required. For SDK Type of `java`, parameters named `serviceName` and `javaPackageName` are required.
	Parameters map[string]string `pulumi:"parameters"`
	// Identifier of the associated REST API.
	RestApiId string `pulumi:"restApiId"`
	// Language for the generated SDK. Currently `java`, `javascript`, `android`, `objectivec` (for iOS), `swift` (for iOS), and `ruby` are supported.
	SdkType string `pulumi:"sdkType"`
	// Name of the Stage that will be exported.
	StageName string `pulumi:"stageName"`
}

// A collection of values returned by getSdk.
type GetSdkResult struct {
	// SDK as a string.
	Body string `pulumi:"body"`
	// Content-disposition header value in the HTTP response.
	ContentDisposition string `pulumi:"contentDisposition"`
	// Content-type header value in the HTTP response.
	ContentType string `pulumi:"contentType"`
	// The provider-assigned unique ID for this managed resource.
	Id         string            `pulumi:"id"`
	Parameters map[string]string `pulumi:"parameters"`
	RestApiId  string            `pulumi:"restApiId"`
	SdkType    string            `pulumi:"sdkType"`
	StageName  string            `pulumi:"stageName"`
}

func GetSdkOutput(ctx *pulumi.Context, args GetSdkOutputArgs, opts ...pulumi.InvokeOption) GetSdkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSdkResult, error) {
			args := v.(GetSdkArgs)
			r, err := GetSdk(ctx, &args, opts...)
			var s GetSdkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSdkResultOutput)
}

// A collection of arguments for invoking getSdk.
type GetSdkOutputArgs struct {
	// Key-value map of query string parameters `sdkType` properties of the SDK. For SDK Type of `objectivec` or `swift`, a parameter named `classPrefix` is required. For SDK Type of `android`, parameters named `groupId`, `artifactId`, `artifactVersion`, and `invokerPackage` are required. For SDK Type of `java`, parameters named `serviceName` and `javaPackageName` are required.
	Parameters pulumi.StringMapInput `pulumi:"parameters"`
	// Identifier of the associated REST API.
	RestApiId pulumi.StringInput `pulumi:"restApiId"`
	// Language for the generated SDK. Currently `java`, `javascript`, `android`, `objectivec` (for iOS), `swift` (for iOS), and `ruby` are supported.
	SdkType pulumi.StringInput `pulumi:"sdkType"`
	// Name of the Stage that will be exported.
	StageName pulumi.StringInput `pulumi:"stageName"`
}

func (GetSdkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSdkArgs)(nil)).Elem()
}

// A collection of values returned by getSdk.
type GetSdkResultOutput struct{ *pulumi.OutputState }

func (GetSdkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSdkResult)(nil)).Elem()
}

func (o GetSdkResultOutput) ToGetSdkResultOutput() GetSdkResultOutput {
	return o
}

func (o GetSdkResultOutput) ToGetSdkResultOutputWithContext(ctx context.Context) GetSdkResultOutput {
	return o
}

// SDK as a string.
func (o GetSdkResultOutput) Body() pulumi.StringOutput {
	return o.ApplyT(func(v GetSdkResult) string { return v.Body }).(pulumi.StringOutput)
}

// Content-disposition header value in the HTTP response.
func (o GetSdkResultOutput) ContentDisposition() pulumi.StringOutput {
	return o.ApplyT(func(v GetSdkResult) string { return v.ContentDisposition }).(pulumi.StringOutput)
}

// Content-type header value in the HTTP response.
func (o GetSdkResultOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v GetSdkResult) string { return v.ContentType }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSdkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSdkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSdkResultOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetSdkResult) map[string]string { return v.Parameters }).(pulumi.StringMapOutput)
}

func (o GetSdkResultOutput) RestApiId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSdkResult) string { return v.RestApiId }).(pulumi.StringOutput)
}

func (o GetSdkResultOutput) SdkType() pulumi.StringOutput {
	return o.ApplyT(func(v GetSdkResult) string { return v.SdkType }).(pulumi.StringOutput)
}

func (o GetSdkResultOutput) StageName() pulumi.StringOutput {
	return o.ApplyT(func(v GetSdkResult) string { return v.StageName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSdkResultOutput{})
}
