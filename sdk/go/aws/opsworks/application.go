// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an OpsWorks application resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"os"
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/opsworks"
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
//			_, err := opsworks.NewApplication(ctx, "foo-app", &opsworks.ApplicationArgs{
//				ShortName:   pulumi.String("foobar"),
//				StackId:     pulumi.Any(aws_opsworks_stack.Main.Id),
//				Type:        pulumi.String("rails"),
//				Description: pulumi.String("This is a Rails application"),
//				Domains: pulumi.StringArray{
//					pulumi.String("example.com"),
//					pulumi.String("sub.example.com"),
//				},
//				Environments: opsworks.ApplicationEnvironmentArray{
//					&opsworks.ApplicationEnvironmentArgs{
//						Key:    pulumi.String("key"),
//						Value:  pulumi.String("value"),
//						Secure: pulumi.Bool(false),
//					},
//				},
//				AppSources: opsworks.ApplicationAppSourceArray{
//					&opsworks.ApplicationAppSourceArgs{
//						Type:     pulumi.String("git"),
//						Revision: pulumi.String("master"),
//						Url:      pulumi.String("https://github.com/example.git"),
//					},
//				},
//				EnableSsl: pulumi.Bool(true),
//				SslConfigurations: opsworks.ApplicationSslConfigurationArray{
//					&opsworks.ApplicationSslConfigurationArgs{
//						PrivateKey:  readFileOrPanic("./foobar.key"),
//						Certificate: readFileOrPanic("./foobar.crt"),
//					},
//				},
//				DocumentRoot:       pulumi.String("public"),
//				AutoBundleOnDeploy: pulumi.String("true"),
//				RailsEnv:           pulumi.String("staging"),
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
// Opsworks Application can be imported using the `id`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:opsworks/application:Application test <id>
//
// ```
type Application struct {
	pulumi.CustomResourceState

	// SCM configuration of the app as described below.
	AppSources ApplicationAppSourceArrayOutput `pulumi:"appSources"`
	// Run bundle install when deploying for application of type `rails`.
	AutoBundleOnDeploy pulumi.StringPtrOutput `pulumi:"autoBundleOnDeploy"`
	// Specify activity and workflow workers for your app using the aws-flow gem.
	AwsFlowRubySettings pulumi.StringPtrOutput `pulumi:"awsFlowRubySettings"`
	// The data source's ARN.
	DataSourceArn pulumi.StringPtrOutput `pulumi:"dataSourceArn"`
	// The database name.
	DataSourceDatabaseName pulumi.StringPtrOutput `pulumi:"dataSourceDatabaseName"`
	// The data source's type one of `AutoSelectOpsworksMysqlInstance`, `OpsworksMysqlInstance`, or `RdsDbInstance`.
	DataSourceType pulumi.StringPtrOutput `pulumi:"dataSourceType"`
	// A description of the app.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Subfolder for the document root for application of type `rails`.
	DocumentRoot pulumi.StringPtrOutput `pulumi:"documentRoot"`
	// A list of virtual host alias.
	Domains pulumi.StringArrayOutput `pulumi:"domains"`
	// Whether to enable SSL for the app. This must be set in order to let `ssl_configuration.private_key`, `ssl_configuration.certificate` and `ssl_configuration.chain` take effect.
	EnableSsl pulumi.BoolPtrOutput `pulumi:"enableSsl"`
	// Object to define environment variables.  Object is described below.
	Environments ApplicationEnvironmentArrayOutput `pulumi:"environments"`
	// A human-readable name for the application.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Rails environment for application of type `rails`.
	RailsEnv pulumi.StringPtrOutput `pulumi:"railsEnv"`
	// A short, machine-readable name for the application. This can only be defined on resource creation and ignored on resource update.
	ShortName pulumi.StringOutput `pulumi:"shortName"`
	// The SSL configuration of the app. Object is described below.
	SslConfigurations ApplicationSslConfigurationArrayOutput `pulumi:"sslConfigurations"`
	// ID of the stack the application will belong to.
	StackId pulumi.StringOutput `pulumi:"stackId"`
	// Opsworks application type. One of `aws-flow-ruby`, `java`, `rails`, `php`, `nodejs`, `static` or `other`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StackId == nil {
		return nil, errors.New("invalid value for required argument 'StackId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource Application
	err := ctx.RegisterResource("aws:opsworks/application:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("aws:opsworks/application:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
	// SCM configuration of the app as described below.
	AppSources []ApplicationAppSource `pulumi:"appSources"`
	// Run bundle install when deploying for application of type `rails`.
	AutoBundleOnDeploy *string `pulumi:"autoBundleOnDeploy"`
	// Specify activity and workflow workers for your app using the aws-flow gem.
	AwsFlowRubySettings *string `pulumi:"awsFlowRubySettings"`
	// The data source's ARN.
	DataSourceArn *string `pulumi:"dataSourceArn"`
	// The database name.
	DataSourceDatabaseName *string `pulumi:"dataSourceDatabaseName"`
	// The data source's type one of `AutoSelectOpsworksMysqlInstance`, `OpsworksMysqlInstance`, or `RdsDbInstance`.
	DataSourceType *string `pulumi:"dataSourceType"`
	// A description of the app.
	Description *string `pulumi:"description"`
	// Subfolder for the document root for application of type `rails`.
	DocumentRoot *string `pulumi:"documentRoot"`
	// A list of virtual host alias.
	Domains []string `pulumi:"domains"`
	// Whether to enable SSL for the app. This must be set in order to let `ssl_configuration.private_key`, `ssl_configuration.certificate` and `ssl_configuration.chain` take effect.
	EnableSsl *bool `pulumi:"enableSsl"`
	// Object to define environment variables.  Object is described below.
	Environments []ApplicationEnvironment `pulumi:"environments"`
	// A human-readable name for the application.
	Name *string `pulumi:"name"`
	// The name of the Rails environment for application of type `rails`.
	RailsEnv *string `pulumi:"railsEnv"`
	// A short, machine-readable name for the application. This can only be defined on resource creation and ignored on resource update.
	ShortName *string `pulumi:"shortName"`
	// The SSL configuration of the app. Object is described below.
	SslConfigurations []ApplicationSslConfiguration `pulumi:"sslConfigurations"`
	// ID of the stack the application will belong to.
	StackId *string `pulumi:"stackId"`
	// Opsworks application type. One of `aws-flow-ruby`, `java`, `rails`, `php`, `nodejs`, `static` or `other`.
	Type *string `pulumi:"type"`
}

type ApplicationState struct {
	// SCM configuration of the app as described below.
	AppSources ApplicationAppSourceArrayInput
	// Run bundle install when deploying for application of type `rails`.
	AutoBundleOnDeploy pulumi.StringPtrInput
	// Specify activity and workflow workers for your app using the aws-flow gem.
	AwsFlowRubySettings pulumi.StringPtrInput
	// The data source's ARN.
	DataSourceArn pulumi.StringPtrInput
	// The database name.
	DataSourceDatabaseName pulumi.StringPtrInput
	// The data source's type one of `AutoSelectOpsworksMysqlInstance`, `OpsworksMysqlInstance`, or `RdsDbInstance`.
	DataSourceType pulumi.StringPtrInput
	// A description of the app.
	Description pulumi.StringPtrInput
	// Subfolder for the document root for application of type `rails`.
	DocumentRoot pulumi.StringPtrInput
	// A list of virtual host alias.
	Domains pulumi.StringArrayInput
	// Whether to enable SSL for the app. This must be set in order to let `ssl_configuration.private_key`, `ssl_configuration.certificate` and `ssl_configuration.chain` take effect.
	EnableSsl pulumi.BoolPtrInput
	// Object to define environment variables.  Object is described below.
	Environments ApplicationEnvironmentArrayInput
	// A human-readable name for the application.
	Name pulumi.StringPtrInput
	// The name of the Rails environment for application of type `rails`.
	RailsEnv pulumi.StringPtrInput
	// A short, machine-readable name for the application. This can only be defined on resource creation and ignored on resource update.
	ShortName pulumi.StringPtrInput
	// The SSL configuration of the app. Object is described below.
	SslConfigurations ApplicationSslConfigurationArrayInput
	// ID of the stack the application will belong to.
	StackId pulumi.StringPtrInput
	// Opsworks application type. One of `aws-flow-ruby`, `java`, `rails`, `php`, `nodejs`, `static` or `other`.
	Type pulumi.StringPtrInput
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// SCM configuration of the app as described below.
	AppSources []ApplicationAppSource `pulumi:"appSources"`
	// Run bundle install when deploying for application of type `rails`.
	AutoBundleOnDeploy *string `pulumi:"autoBundleOnDeploy"`
	// Specify activity and workflow workers for your app using the aws-flow gem.
	AwsFlowRubySettings *string `pulumi:"awsFlowRubySettings"`
	// The data source's ARN.
	DataSourceArn *string `pulumi:"dataSourceArn"`
	// The database name.
	DataSourceDatabaseName *string `pulumi:"dataSourceDatabaseName"`
	// The data source's type one of `AutoSelectOpsworksMysqlInstance`, `OpsworksMysqlInstance`, or `RdsDbInstance`.
	DataSourceType *string `pulumi:"dataSourceType"`
	// A description of the app.
	Description *string `pulumi:"description"`
	// Subfolder for the document root for application of type `rails`.
	DocumentRoot *string `pulumi:"documentRoot"`
	// A list of virtual host alias.
	Domains []string `pulumi:"domains"`
	// Whether to enable SSL for the app. This must be set in order to let `ssl_configuration.private_key`, `ssl_configuration.certificate` and `ssl_configuration.chain` take effect.
	EnableSsl *bool `pulumi:"enableSsl"`
	// Object to define environment variables.  Object is described below.
	Environments []ApplicationEnvironment `pulumi:"environments"`
	// A human-readable name for the application.
	Name *string `pulumi:"name"`
	// The name of the Rails environment for application of type `rails`.
	RailsEnv *string `pulumi:"railsEnv"`
	// A short, machine-readable name for the application. This can only be defined on resource creation and ignored on resource update.
	ShortName *string `pulumi:"shortName"`
	// The SSL configuration of the app. Object is described below.
	SslConfigurations []ApplicationSslConfiguration `pulumi:"sslConfigurations"`
	// ID of the stack the application will belong to.
	StackId string `pulumi:"stackId"`
	// Opsworks application type. One of `aws-flow-ruby`, `java`, `rails`, `php`, `nodejs`, `static` or `other`.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// SCM configuration of the app as described below.
	AppSources ApplicationAppSourceArrayInput
	// Run bundle install when deploying for application of type `rails`.
	AutoBundleOnDeploy pulumi.StringPtrInput
	// Specify activity and workflow workers for your app using the aws-flow gem.
	AwsFlowRubySettings pulumi.StringPtrInput
	// The data source's ARN.
	DataSourceArn pulumi.StringPtrInput
	// The database name.
	DataSourceDatabaseName pulumi.StringPtrInput
	// The data source's type one of `AutoSelectOpsworksMysqlInstance`, `OpsworksMysqlInstance`, or `RdsDbInstance`.
	DataSourceType pulumi.StringPtrInput
	// A description of the app.
	Description pulumi.StringPtrInput
	// Subfolder for the document root for application of type `rails`.
	DocumentRoot pulumi.StringPtrInput
	// A list of virtual host alias.
	Domains pulumi.StringArrayInput
	// Whether to enable SSL for the app. This must be set in order to let `ssl_configuration.private_key`, `ssl_configuration.certificate` and `ssl_configuration.chain` take effect.
	EnableSsl pulumi.BoolPtrInput
	// Object to define environment variables.  Object is described below.
	Environments ApplicationEnvironmentArrayInput
	// A human-readable name for the application.
	Name pulumi.StringPtrInput
	// The name of the Rails environment for application of type `rails`.
	RailsEnv pulumi.StringPtrInput
	// A short, machine-readable name for the application. This can only be defined on resource creation and ignored on resource update.
	ShortName pulumi.StringPtrInput
	// The SSL configuration of the app. Object is described below.
	SslConfigurations ApplicationSslConfigurationArrayInput
	// ID of the stack the application will belong to.
	StackId pulumi.StringInput
	// Opsworks application type. One of `aws-flow-ruby`, `java`, `rails`, `php`, `nodejs`, `static` or `other`.
	Type pulumi.StringInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

type ApplicationInput interface {
	pulumi.Input

	ToApplicationOutput() ApplicationOutput
	ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput
}

func (*Application) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (i *Application) ToApplicationOutput() ApplicationOutput {
	return i.ToApplicationOutputWithContext(context.Background())
}

func (i *Application) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutput)
}

// ApplicationArrayInput is an input type that accepts ApplicationArray and ApplicationArrayOutput values.
// You can construct a concrete instance of `ApplicationArrayInput` via:
//
//	ApplicationArray{ ApplicationArgs{...} }
type ApplicationArrayInput interface {
	pulumi.Input

	ToApplicationArrayOutput() ApplicationArrayOutput
	ToApplicationArrayOutputWithContext(context.Context) ApplicationArrayOutput
}

type ApplicationArray []ApplicationInput

func (ApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Application)(nil)).Elem()
}

func (i ApplicationArray) ToApplicationArrayOutput() ApplicationArrayOutput {
	return i.ToApplicationArrayOutputWithContext(context.Background())
}

func (i ApplicationArray) ToApplicationArrayOutputWithContext(ctx context.Context) ApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationArrayOutput)
}

// ApplicationMapInput is an input type that accepts ApplicationMap and ApplicationMapOutput values.
// You can construct a concrete instance of `ApplicationMapInput` via:
//
//	ApplicationMap{ "key": ApplicationArgs{...} }
type ApplicationMapInput interface {
	pulumi.Input

	ToApplicationMapOutput() ApplicationMapOutput
	ToApplicationMapOutputWithContext(context.Context) ApplicationMapOutput
}

type ApplicationMap map[string]ApplicationInput

func (ApplicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Application)(nil)).Elem()
}

func (i ApplicationMap) ToApplicationMapOutput() ApplicationMapOutput {
	return i.ToApplicationMapOutputWithContext(context.Background())
}

func (i ApplicationMap) ToApplicationMapOutputWithContext(ctx context.Context) ApplicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationMapOutput)
}

type ApplicationOutput struct{ *pulumi.OutputState }

func (ApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (o ApplicationOutput) ToApplicationOutput() ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return o
}

// SCM configuration of the app as described below.
func (o ApplicationOutput) AppSources() ApplicationAppSourceArrayOutput {
	return o.ApplyT(func(v *Application) ApplicationAppSourceArrayOutput { return v.AppSources }).(ApplicationAppSourceArrayOutput)
}

// Run bundle install when deploying for application of type `rails`.
func (o ApplicationOutput) AutoBundleOnDeploy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.AutoBundleOnDeploy }).(pulumi.StringPtrOutput)
}

// Specify activity and workflow workers for your app using the aws-flow gem.
func (o ApplicationOutput) AwsFlowRubySettings() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.AwsFlowRubySettings }).(pulumi.StringPtrOutput)
}

// The data source's ARN.
func (o ApplicationOutput) DataSourceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.DataSourceArn }).(pulumi.StringPtrOutput)
}

// The database name.
func (o ApplicationOutput) DataSourceDatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.DataSourceDatabaseName }).(pulumi.StringPtrOutput)
}

// The data source's type one of `AutoSelectOpsworksMysqlInstance`, `OpsworksMysqlInstance`, or `RdsDbInstance`.
func (o ApplicationOutput) DataSourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.DataSourceType }).(pulumi.StringPtrOutput)
}

// A description of the app.
func (o ApplicationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Subfolder for the document root for application of type `rails`.
func (o ApplicationOutput) DocumentRoot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.DocumentRoot }).(pulumi.StringPtrOutput)
}

// A list of virtual host alias.
func (o ApplicationOutput) Domains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Application) pulumi.StringArrayOutput { return v.Domains }).(pulumi.StringArrayOutput)
}

// Whether to enable SSL for the app. This must be set in order to let `ssl_configuration.private_key`, `ssl_configuration.certificate` and `ssl_configuration.chain` take effect.
func (o ApplicationOutput) EnableSsl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.BoolPtrOutput { return v.EnableSsl }).(pulumi.BoolPtrOutput)
}

// Object to define environment variables.  Object is described below.
func (o ApplicationOutput) Environments() ApplicationEnvironmentArrayOutput {
	return o.ApplyT(func(v *Application) ApplicationEnvironmentArrayOutput { return v.Environments }).(ApplicationEnvironmentArrayOutput)
}

// A human-readable name for the application.
func (o ApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the Rails environment for application of type `rails`.
func (o ApplicationOutput) RailsEnv() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.RailsEnv }).(pulumi.StringPtrOutput)
}

// A short, machine-readable name for the application. This can only be defined on resource creation and ignored on resource update.
func (o ApplicationOutput) ShortName() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.ShortName }).(pulumi.StringOutput)
}

// The SSL configuration of the app. Object is described below.
func (o ApplicationOutput) SslConfigurations() ApplicationSslConfigurationArrayOutput {
	return o.ApplyT(func(v *Application) ApplicationSslConfigurationArrayOutput { return v.SslConfigurations }).(ApplicationSslConfigurationArrayOutput)
}

// ID of the stack the application will belong to.
func (o ApplicationOutput) StackId() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.StackId }).(pulumi.StringOutput)
}

// Opsworks application type. One of `aws-flow-ruby`, `java`, `rails`, `php`, `nodejs`, `static` or `other`.
func (o ApplicationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type ApplicationArrayOutput struct{ *pulumi.OutputState }

func (ApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Application)(nil)).Elem()
}

func (o ApplicationArrayOutput) ToApplicationArrayOutput() ApplicationArrayOutput {
	return o
}

func (o ApplicationArrayOutput) ToApplicationArrayOutputWithContext(ctx context.Context) ApplicationArrayOutput {
	return o
}

func (o ApplicationArrayOutput) Index(i pulumi.IntInput) ApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Application {
		return vs[0].([]*Application)[vs[1].(int)]
	}).(ApplicationOutput)
}

type ApplicationMapOutput struct{ *pulumi.OutputState }

func (ApplicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Application)(nil)).Elem()
}

func (o ApplicationMapOutput) ToApplicationMapOutput() ApplicationMapOutput {
	return o
}

func (o ApplicationMapOutput) ToApplicationMapOutputWithContext(ctx context.Context) ApplicationMapOutput {
	return o
}

func (o ApplicationMapOutput) MapIndex(k pulumi.StringInput) ApplicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Application {
		return vs[0].(map[string]*Application)[vs[1].(string)]
	}).(ApplicationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationInput)(nil)).Elem(), &Application{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationArrayInput)(nil)).Elem(), ApplicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationMapInput)(nil)).Elem(), ApplicationMap{})
	pulumi.RegisterOutputType(ApplicationOutput{})
	pulumi.RegisterOutputType(ApplicationArrayOutput{})
	pulumi.RegisterOutputType(ApplicationMapOutput{})
}
