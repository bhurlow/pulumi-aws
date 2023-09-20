// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package amplify

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides an Amplify App resource, a fullstack serverless app hosted on the [AWS Amplify Console](https://docs.aws.amazon.com/amplify/latest/userguide/welcome.html).
//
// > **Note:** When you create/update an Amplify App from the provider, you may end up with the error "BadRequestException: You should at least provide one valid token" because of authentication issues. See the section "Repository with Tokens" below.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/amplify"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := amplify.NewApp(ctx, "example", &amplify.AppArgs{
//				BuildSpec: pulumi.String(`  version: 0.1
//	  frontend:
//	    phases:
//	      preBuild:
//	        commands:
//	          - yarn install
//	      build:
//	        commands:
//	          - yarn run build
//	    artifacts:
//	      baseDirectory: build
//	      files:
//	        - '**/*'
//	    cache:
//	      paths:
//	        - node_modules/**/*
//
// `),
//
//				CustomRules: amplify.AppCustomRuleArray{
//					&amplify.AppCustomRuleArgs{
//						Source: pulumi.String("/<*>"),
//						Status: pulumi.String("404"),
//						Target: pulumi.String("/index.html"),
//					},
//				},
//				EnvironmentVariables: pulumi.StringMap{
//					"ENV": pulumi.String("test"),
//				},
//				Repository: pulumi.String("https://github.com/example/app"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Repository with Tokens
//
// If you create a new Amplify App with the `repository` argument, you also need to set `oauthToken` or `accessToken` for authentication. For GitHub, get a [personal access token](https://help.github.com/en/github/authenticating-to-github/creating-a-personal-access-token-for-the-command-line) and set `accessToken` as follows:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/amplify"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := amplify.NewApp(ctx, "example", &amplify.AppArgs{
//				AccessToken: pulumi.String("..."),
//				Repository:  pulumi.String("https://github.com/example/app"),
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
// You can omit `accessToken` if you import an existing Amplify App created by the Amplify Console (using OAuth for authentication).
// ### Auto Branch Creation
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/amplify"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := amplify.NewApp(ctx, "example", &amplify.AppArgs{
//				AutoBranchCreationConfig: &amplify.AppAutoBranchCreationConfigArgs{
//					EnableAutoBuild: pulumi.Bool(true),
//				},
//				AutoBranchCreationPatterns: pulumi.StringArray{
//					pulumi.String("*"),
//					pulumi.String("*/**"),
//				},
//				EnableAutoBranchCreation: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Rewrites and Redirects
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/amplify"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := amplify.NewApp(ctx, "example", &amplify.AppArgs{
//				CustomRules: amplify.AppCustomRuleArray{
//					&amplify.AppCustomRuleArgs{
//						Source: pulumi.String("/api/<*>"),
//						Status: pulumi.String("200"),
//						Target: pulumi.String("https://api.example.com/api/<*>"),
//					},
//					&amplify.AppCustomRuleArgs{
//						Source: pulumi.String("</^[^.]+$|\\.(?!(css|gif|ico|jpg|js|png|txt|svg|woff|ttf|map|json)$)([^.]+$)/>"),
//						Status: pulumi.String("200"),
//						Target: pulumi.String("/index.html"),
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
// ### Custom Image
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/amplify"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := amplify.NewApp(ctx, "example", &amplify.AppArgs{
//				EnvironmentVariables: pulumi.StringMap{
//					"_CUSTOM_IMAGE": pulumi.String("node:16"),
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
// Using `pulumi import`, import Amplify App using Amplify App ID (appId). For example:
//
// ```sh
//
//	$ pulumi import aws:amplify/app:App example d2ypk4k47z8u6
//
// ```
//
//	App ID can be obtained from App ARN (e.g., `arn:aws:amplify:us-east-1:12345678:apps/d2ypk4k47z8u6`).
type App struct {
	pulumi.CustomResourceState

	// Personal access token for a third-party source control system for an Amplify app. The personal access token is used to create a webhook and a read-only deploy key. The token is not stored.
	AccessToken pulumi.StringPtrOutput `pulumi:"accessToken"`
	// ARN of the Amplify app.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Automated branch creation configuration for an Amplify app. An `autoBranchCreationConfig` block is documented below.
	AutoBranchCreationConfig AppAutoBranchCreationConfigOutput `pulumi:"autoBranchCreationConfig"`
	// Automated branch creation glob patterns for an Amplify app.
	AutoBranchCreationPatterns pulumi.StringArrayOutput `pulumi:"autoBranchCreationPatterns"`
	// Credentials for basic authorization for an Amplify app.
	BasicAuthCredentials pulumi.StringPtrOutput `pulumi:"basicAuthCredentials"`
	// The [build specification](https://docs.aws.amazon.com/amplify/latest/userguide/build-settings.html) (build spec) for an Amplify app.
	BuildSpec pulumi.StringOutput `pulumi:"buildSpec"`
	// Custom rewrite and redirect rules for an Amplify app. A `customRule` block is documented below.
	CustomRules AppCustomRuleArrayOutput `pulumi:"customRules"`
	// Default domain for the Amplify app.
	DefaultDomain pulumi.StringOutput `pulumi:"defaultDomain"`
	// Description for an Amplify app.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Enables automated branch creation for an Amplify app.
	EnableAutoBranchCreation pulumi.BoolPtrOutput `pulumi:"enableAutoBranchCreation"`
	// Enables basic authorization for an Amplify app. This will apply to all branches that are part of this app.
	EnableBasicAuth pulumi.BoolPtrOutput `pulumi:"enableBasicAuth"`
	// Enables auto-building of branches for the Amplify App.
	EnableBranchAutoBuild pulumi.BoolPtrOutput `pulumi:"enableBranchAutoBuild"`
	// Automatically disconnects a branch in the Amplify Console when you delete a branch from your Git repository.
	EnableBranchAutoDeletion pulumi.BoolPtrOutput `pulumi:"enableBranchAutoDeletion"`
	// Environment variables map for an Amplify app.
	EnvironmentVariables pulumi.StringMapOutput `pulumi:"environmentVariables"`
	// AWS Identity and Access Management (IAM) service role for an Amplify app.
	IamServiceRoleArn pulumi.StringPtrOutput `pulumi:"iamServiceRoleArn"`
	// Name for an Amplify app.
	Name pulumi.StringOutput `pulumi:"name"`
	// OAuth token for a third-party source control system for an Amplify app. The OAuth token is used to create a webhook and a read-only deploy key. The OAuth token is not stored.
	OauthToken pulumi.StringPtrOutput `pulumi:"oauthToken"`
	// Platform or framework for an Amplify app. Valid values: `WEB`, `WEB_COMPUTE`. Default value: `WEB`.
	Platform pulumi.StringPtrOutput `pulumi:"platform"`
	// Describes the information about a production branch for an Amplify app. A `productionBranch` block is documented below.
	ProductionBranches AppProductionBranchArrayOutput `pulumi:"productionBranches"`
	// Repository for an Amplify app.
	Repository pulumi.StringPtrOutput `pulumi:"repository"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewApp registers a new resource with the given unique name, arguments, and options.
func NewApp(ctx *pulumi.Context,
	name string, args *AppArgs, opts ...pulumi.ResourceOption) (*App, error) {
	if args == nil {
		args = &AppArgs{}
	}

	if args.AccessToken != nil {
		args.AccessToken = pulumi.ToSecret(args.AccessToken).(pulumi.StringPtrInput)
	}
	if args.BasicAuthCredentials != nil {
		args.BasicAuthCredentials = pulumi.ToSecret(args.BasicAuthCredentials).(pulumi.StringPtrInput)
	}
	if args.OauthToken != nil {
		args.OauthToken = pulumi.ToSecret(args.OauthToken).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"accessToken",
		"basicAuthCredentials",
		"oauthToken",
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource App
	err := ctx.RegisterResource("aws:amplify/app:App", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApp gets an existing App resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppState, opts ...pulumi.ResourceOption) (*App, error) {
	var resource App
	err := ctx.ReadResource("aws:amplify/app:App", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering App resources.
type appState struct {
	// Personal access token for a third-party source control system for an Amplify app. The personal access token is used to create a webhook and a read-only deploy key. The token is not stored.
	AccessToken *string `pulumi:"accessToken"`
	// ARN of the Amplify app.
	Arn *string `pulumi:"arn"`
	// Automated branch creation configuration for an Amplify app. An `autoBranchCreationConfig` block is documented below.
	AutoBranchCreationConfig *AppAutoBranchCreationConfig `pulumi:"autoBranchCreationConfig"`
	// Automated branch creation glob patterns for an Amplify app.
	AutoBranchCreationPatterns []string `pulumi:"autoBranchCreationPatterns"`
	// Credentials for basic authorization for an Amplify app.
	BasicAuthCredentials *string `pulumi:"basicAuthCredentials"`
	// The [build specification](https://docs.aws.amazon.com/amplify/latest/userguide/build-settings.html) (build spec) for an Amplify app.
	BuildSpec *string `pulumi:"buildSpec"`
	// Custom rewrite and redirect rules for an Amplify app. A `customRule` block is documented below.
	CustomRules []AppCustomRule `pulumi:"customRules"`
	// Default domain for the Amplify app.
	DefaultDomain *string `pulumi:"defaultDomain"`
	// Description for an Amplify app.
	Description *string `pulumi:"description"`
	// Enables automated branch creation for an Amplify app.
	EnableAutoBranchCreation *bool `pulumi:"enableAutoBranchCreation"`
	// Enables basic authorization for an Amplify app. This will apply to all branches that are part of this app.
	EnableBasicAuth *bool `pulumi:"enableBasicAuth"`
	// Enables auto-building of branches for the Amplify App.
	EnableBranchAutoBuild *bool `pulumi:"enableBranchAutoBuild"`
	// Automatically disconnects a branch in the Amplify Console when you delete a branch from your Git repository.
	EnableBranchAutoDeletion *bool `pulumi:"enableBranchAutoDeletion"`
	// Environment variables map for an Amplify app.
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	// AWS Identity and Access Management (IAM) service role for an Amplify app.
	IamServiceRoleArn *string `pulumi:"iamServiceRoleArn"`
	// Name for an Amplify app.
	Name *string `pulumi:"name"`
	// OAuth token for a third-party source control system for an Amplify app. The OAuth token is used to create a webhook and a read-only deploy key. The OAuth token is not stored.
	OauthToken *string `pulumi:"oauthToken"`
	// Platform or framework for an Amplify app. Valid values: `WEB`, `WEB_COMPUTE`. Default value: `WEB`.
	Platform *string `pulumi:"platform"`
	// Describes the information about a production branch for an Amplify app. A `productionBranch` block is documented below.
	ProductionBranches []AppProductionBranch `pulumi:"productionBranches"`
	// Repository for an Amplify app.
	Repository *string `pulumi:"repository"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type AppState struct {
	// Personal access token for a third-party source control system for an Amplify app. The personal access token is used to create a webhook and a read-only deploy key. The token is not stored.
	AccessToken pulumi.StringPtrInput
	// ARN of the Amplify app.
	Arn pulumi.StringPtrInput
	// Automated branch creation configuration for an Amplify app. An `autoBranchCreationConfig` block is documented below.
	AutoBranchCreationConfig AppAutoBranchCreationConfigPtrInput
	// Automated branch creation glob patterns for an Amplify app.
	AutoBranchCreationPatterns pulumi.StringArrayInput
	// Credentials for basic authorization for an Amplify app.
	BasicAuthCredentials pulumi.StringPtrInput
	// The [build specification](https://docs.aws.amazon.com/amplify/latest/userguide/build-settings.html) (build spec) for an Amplify app.
	BuildSpec pulumi.StringPtrInput
	// Custom rewrite and redirect rules for an Amplify app. A `customRule` block is documented below.
	CustomRules AppCustomRuleArrayInput
	// Default domain for the Amplify app.
	DefaultDomain pulumi.StringPtrInput
	// Description for an Amplify app.
	Description pulumi.StringPtrInput
	// Enables automated branch creation for an Amplify app.
	EnableAutoBranchCreation pulumi.BoolPtrInput
	// Enables basic authorization for an Amplify app. This will apply to all branches that are part of this app.
	EnableBasicAuth pulumi.BoolPtrInput
	// Enables auto-building of branches for the Amplify App.
	EnableBranchAutoBuild pulumi.BoolPtrInput
	// Automatically disconnects a branch in the Amplify Console when you delete a branch from your Git repository.
	EnableBranchAutoDeletion pulumi.BoolPtrInput
	// Environment variables map for an Amplify app.
	EnvironmentVariables pulumi.StringMapInput
	// AWS Identity and Access Management (IAM) service role for an Amplify app.
	IamServiceRoleArn pulumi.StringPtrInput
	// Name for an Amplify app.
	Name pulumi.StringPtrInput
	// OAuth token for a third-party source control system for an Amplify app. The OAuth token is used to create a webhook and a read-only deploy key. The OAuth token is not stored.
	OauthToken pulumi.StringPtrInput
	// Platform or framework for an Amplify app. Valid values: `WEB`, `WEB_COMPUTE`. Default value: `WEB`.
	Platform pulumi.StringPtrInput
	// Describes the information about a production branch for an Amplify app. A `productionBranch` block is documented below.
	ProductionBranches AppProductionBranchArrayInput
	// Repository for an Amplify app.
	Repository pulumi.StringPtrInput
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (AppState) ElementType() reflect.Type {
	return reflect.TypeOf((*appState)(nil)).Elem()
}

type appArgs struct {
	// Personal access token for a third-party source control system for an Amplify app. The personal access token is used to create a webhook and a read-only deploy key. The token is not stored.
	AccessToken *string `pulumi:"accessToken"`
	// Automated branch creation configuration for an Amplify app. An `autoBranchCreationConfig` block is documented below.
	AutoBranchCreationConfig *AppAutoBranchCreationConfig `pulumi:"autoBranchCreationConfig"`
	// Automated branch creation glob patterns for an Amplify app.
	AutoBranchCreationPatterns []string `pulumi:"autoBranchCreationPatterns"`
	// Credentials for basic authorization for an Amplify app.
	BasicAuthCredentials *string `pulumi:"basicAuthCredentials"`
	// The [build specification](https://docs.aws.amazon.com/amplify/latest/userguide/build-settings.html) (build spec) for an Amplify app.
	BuildSpec *string `pulumi:"buildSpec"`
	// Custom rewrite and redirect rules for an Amplify app. A `customRule` block is documented below.
	CustomRules []AppCustomRule `pulumi:"customRules"`
	// Description for an Amplify app.
	Description *string `pulumi:"description"`
	// Enables automated branch creation for an Amplify app.
	EnableAutoBranchCreation *bool `pulumi:"enableAutoBranchCreation"`
	// Enables basic authorization for an Amplify app. This will apply to all branches that are part of this app.
	EnableBasicAuth *bool `pulumi:"enableBasicAuth"`
	// Enables auto-building of branches for the Amplify App.
	EnableBranchAutoBuild *bool `pulumi:"enableBranchAutoBuild"`
	// Automatically disconnects a branch in the Amplify Console when you delete a branch from your Git repository.
	EnableBranchAutoDeletion *bool `pulumi:"enableBranchAutoDeletion"`
	// Environment variables map for an Amplify app.
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	// AWS Identity and Access Management (IAM) service role for an Amplify app.
	IamServiceRoleArn *string `pulumi:"iamServiceRoleArn"`
	// Name for an Amplify app.
	Name *string `pulumi:"name"`
	// OAuth token for a third-party source control system for an Amplify app. The OAuth token is used to create a webhook and a read-only deploy key. The OAuth token is not stored.
	OauthToken *string `pulumi:"oauthToken"`
	// Platform or framework for an Amplify app. Valid values: `WEB`, `WEB_COMPUTE`. Default value: `WEB`.
	Platform *string `pulumi:"platform"`
	// Repository for an Amplify app.
	Repository *string `pulumi:"repository"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a App resource.
type AppArgs struct {
	// Personal access token for a third-party source control system for an Amplify app. The personal access token is used to create a webhook and a read-only deploy key. The token is not stored.
	AccessToken pulumi.StringPtrInput
	// Automated branch creation configuration for an Amplify app. An `autoBranchCreationConfig` block is documented below.
	AutoBranchCreationConfig AppAutoBranchCreationConfigPtrInput
	// Automated branch creation glob patterns for an Amplify app.
	AutoBranchCreationPatterns pulumi.StringArrayInput
	// Credentials for basic authorization for an Amplify app.
	BasicAuthCredentials pulumi.StringPtrInput
	// The [build specification](https://docs.aws.amazon.com/amplify/latest/userguide/build-settings.html) (build spec) for an Amplify app.
	BuildSpec pulumi.StringPtrInput
	// Custom rewrite and redirect rules for an Amplify app. A `customRule` block is documented below.
	CustomRules AppCustomRuleArrayInput
	// Description for an Amplify app.
	Description pulumi.StringPtrInput
	// Enables automated branch creation for an Amplify app.
	EnableAutoBranchCreation pulumi.BoolPtrInput
	// Enables basic authorization for an Amplify app. This will apply to all branches that are part of this app.
	EnableBasicAuth pulumi.BoolPtrInput
	// Enables auto-building of branches for the Amplify App.
	EnableBranchAutoBuild pulumi.BoolPtrInput
	// Automatically disconnects a branch in the Amplify Console when you delete a branch from your Git repository.
	EnableBranchAutoDeletion pulumi.BoolPtrInput
	// Environment variables map for an Amplify app.
	EnvironmentVariables pulumi.StringMapInput
	// AWS Identity and Access Management (IAM) service role for an Amplify app.
	IamServiceRoleArn pulumi.StringPtrInput
	// Name for an Amplify app.
	Name pulumi.StringPtrInput
	// OAuth token for a third-party source control system for an Amplify app. The OAuth token is used to create a webhook and a read-only deploy key. The OAuth token is not stored.
	OauthToken pulumi.StringPtrInput
	// Platform or framework for an Amplify app. Valid values: `WEB`, `WEB_COMPUTE`. Default value: `WEB`.
	Platform pulumi.StringPtrInput
	// Repository for an Amplify app.
	Repository pulumi.StringPtrInput
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (AppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appArgs)(nil)).Elem()
}

type AppInput interface {
	pulumi.Input

	ToAppOutput() AppOutput
	ToAppOutputWithContext(ctx context.Context) AppOutput
}

func (*App) ElementType() reflect.Type {
	return reflect.TypeOf((**App)(nil)).Elem()
}

func (i *App) ToAppOutput() AppOutput {
	return i.ToAppOutputWithContext(context.Background())
}

func (i *App) ToAppOutputWithContext(ctx context.Context) AppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppOutput)
}

func (i *App) ToOutput(ctx context.Context) pulumix.Output[*App] {
	return pulumix.Output[*App]{
		OutputState: i.ToAppOutputWithContext(ctx).OutputState,
	}
}

// AppArrayInput is an input type that accepts AppArray and AppArrayOutput values.
// You can construct a concrete instance of `AppArrayInput` via:
//
//	AppArray{ AppArgs{...} }
type AppArrayInput interface {
	pulumi.Input

	ToAppArrayOutput() AppArrayOutput
	ToAppArrayOutputWithContext(context.Context) AppArrayOutput
}

type AppArray []AppInput

func (AppArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*App)(nil)).Elem()
}

func (i AppArray) ToAppArrayOutput() AppArrayOutput {
	return i.ToAppArrayOutputWithContext(context.Background())
}

func (i AppArray) ToAppArrayOutputWithContext(ctx context.Context) AppArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppArrayOutput)
}

func (i AppArray) ToOutput(ctx context.Context) pulumix.Output[[]*App] {
	return pulumix.Output[[]*App]{
		OutputState: i.ToAppArrayOutputWithContext(ctx).OutputState,
	}
}

// AppMapInput is an input type that accepts AppMap and AppMapOutput values.
// You can construct a concrete instance of `AppMapInput` via:
//
//	AppMap{ "key": AppArgs{...} }
type AppMapInput interface {
	pulumi.Input

	ToAppMapOutput() AppMapOutput
	ToAppMapOutputWithContext(context.Context) AppMapOutput
}

type AppMap map[string]AppInput

func (AppMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*App)(nil)).Elem()
}

func (i AppMap) ToAppMapOutput() AppMapOutput {
	return i.ToAppMapOutputWithContext(context.Background())
}

func (i AppMap) ToAppMapOutputWithContext(ctx context.Context) AppMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppMapOutput)
}

func (i AppMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*App] {
	return pulumix.Output[map[string]*App]{
		OutputState: i.ToAppMapOutputWithContext(ctx).OutputState,
	}
}

type AppOutput struct{ *pulumi.OutputState }

func (AppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**App)(nil)).Elem()
}

func (o AppOutput) ToAppOutput() AppOutput {
	return o
}

func (o AppOutput) ToAppOutputWithContext(ctx context.Context) AppOutput {
	return o
}

func (o AppOutput) ToOutput(ctx context.Context) pulumix.Output[*App] {
	return pulumix.Output[*App]{
		OutputState: o.OutputState,
	}
}

// Personal access token for a third-party source control system for an Amplify app. The personal access token is used to create a webhook and a read-only deploy key. The token is not stored.
func (o AppOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.AccessToken }).(pulumi.StringPtrOutput)
}

// ARN of the Amplify app.
func (o AppOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Automated branch creation configuration for an Amplify app. An `autoBranchCreationConfig` block is documented below.
func (o AppOutput) AutoBranchCreationConfig() AppAutoBranchCreationConfigOutput {
	return o.ApplyT(func(v *App) AppAutoBranchCreationConfigOutput { return v.AutoBranchCreationConfig }).(AppAutoBranchCreationConfigOutput)
}

// Automated branch creation glob patterns for an Amplify app.
func (o AppOutput) AutoBranchCreationPatterns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *App) pulumi.StringArrayOutput { return v.AutoBranchCreationPatterns }).(pulumi.StringArrayOutput)
}

// Credentials for basic authorization for an Amplify app.
func (o AppOutput) BasicAuthCredentials() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.BasicAuthCredentials }).(pulumi.StringPtrOutput)
}

// The [build specification](https://docs.aws.amazon.com/amplify/latest/userguide/build-settings.html) (build spec) for an Amplify app.
func (o AppOutput) BuildSpec() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.BuildSpec }).(pulumi.StringOutput)
}

// Custom rewrite and redirect rules for an Amplify app. A `customRule` block is documented below.
func (o AppOutput) CustomRules() AppCustomRuleArrayOutput {
	return o.ApplyT(func(v *App) AppCustomRuleArrayOutput { return v.CustomRules }).(AppCustomRuleArrayOutput)
}

// Default domain for the Amplify app.
func (o AppOutput) DefaultDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.DefaultDomain }).(pulumi.StringOutput)
}

// Description for an Amplify app.
func (o AppOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Enables automated branch creation for an Amplify app.
func (o AppOutput) EnableAutoBranchCreation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *App) pulumi.BoolPtrOutput { return v.EnableAutoBranchCreation }).(pulumi.BoolPtrOutput)
}

// Enables basic authorization for an Amplify app. This will apply to all branches that are part of this app.
func (o AppOutput) EnableBasicAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *App) pulumi.BoolPtrOutput { return v.EnableBasicAuth }).(pulumi.BoolPtrOutput)
}

// Enables auto-building of branches for the Amplify App.
func (o AppOutput) EnableBranchAutoBuild() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *App) pulumi.BoolPtrOutput { return v.EnableBranchAutoBuild }).(pulumi.BoolPtrOutput)
}

// Automatically disconnects a branch in the Amplify Console when you delete a branch from your Git repository.
func (o AppOutput) EnableBranchAutoDeletion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *App) pulumi.BoolPtrOutput { return v.EnableBranchAutoDeletion }).(pulumi.BoolPtrOutput)
}

// Environment variables map for an Amplify app.
func (o AppOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v *App) pulumi.StringMapOutput { return v.EnvironmentVariables }).(pulumi.StringMapOutput)
}

// AWS Identity and Access Management (IAM) service role for an Amplify app.
func (o AppOutput) IamServiceRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.IamServiceRoleArn }).(pulumi.StringPtrOutput)
}

// Name for an Amplify app.
func (o AppOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// OAuth token for a third-party source control system for an Amplify app. The OAuth token is used to create a webhook and a read-only deploy key. The OAuth token is not stored.
func (o AppOutput) OauthToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.OauthToken }).(pulumi.StringPtrOutput)
}

// Platform or framework for an Amplify app. Valid values: `WEB`, `WEB_COMPUTE`. Default value: `WEB`.
func (o AppOutput) Platform() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.Platform }).(pulumi.StringPtrOutput)
}

// Describes the information about a production branch for an Amplify app. A `productionBranch` block is documented below.
func (o AppOutput) ProductionBranches() AppProductionBranchArrayOutput {
	return o.ApplyT(func(v *App) AppProductionBranchArrayOutput { return v.ProductionBranches }).(AppProductionBranchArrayOutput)
}

// Repository for an Amplify app.
func (o AppOutput) Repository() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.Repository }).(pulumi.StringPtrOutput)
}

// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o AppOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *App) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o AppOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *App) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type AppArrayOutput struct{ *pulumi.OutputState }

func (AppArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*App)(nil)).Elem()
}

func (o AppArrayOutput) ToAppArrayOutput() AppArrayOutput {
	return o
}

func (o AppArrayOutput) ToAppArrayOutputWithContext(ctx context.Context) AppArrayOutput {
	return o
}

func (o AppArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*App] {
	return pulumix.Output[[]*App]{
		OutputState: o.OutputState,
	}
}

func (o AppArrayOutput) Index(i pulumi.IntInput) AppOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *App {
		return vs[0].([]*App)[vs[1].(int)]
	}).(AppOutput)
}

type AppMapOutput struct{ *pulumi.OutputState }

func (AppMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*App)(nil)).Elem()
}

func (o AppMapOutput) ToAppMapOutput() AppMapOutput {
	return o
}

func (o AppMapOutput) ToAppMapOutputWithContext(ctx context.Context) AppMapOutput {
	return o
}

func (o AppMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*App] {
	return pulumix.Output[map[string]*App]{
		OutputState: o.OutputState,
	}
}

func (o AppMapOutput) MapIndex(k pulumi.StringInput) AppOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *App {
		return vs[0].(map[string]*App)[vs[1].(string)]
	}).(AppOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppInput)(nil)).Elem(), &App{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppArrayInput)(nil)).Elem(), AppArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppMapInput)(nil)).Elem(), AppMap{})
	pulumi.RegisterOutputType(AppOutput{})
	pulumi.RegisterOutputType(AppArrayOutput{})
	pulumi.RegisterOutputType(AppMapOutput{})
}
