// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticBeanstalk
{
    /// <summary>
    /// Provides an Elastic Beanstalk Environment Resource. Elastic Beanstalk allows
    /// you to deploy and manage applications in the AWS cloud without worrying about
    /// the infrastructure that runs those applications.
    /// 
    /// Environments are often things such as `development`, `integration`, or
    /// `production`.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var tftest = new Aws.ElasticBeanstalk.Application("tftest", new Aws.ElasticBeanstalk.ApplicationArgs
    ///         {
    ///             Description = "tf-test-desc",
    ///         });
    ///         var tfenvtest = new Aws.ElasticBeanstalk.Environment("tfenvtest", new Aws.ElasticBeanstalk.EnvironmentArgs
    ///         {
    ///             Application = tftest.Name,
    ///             SolutionStackName = "64bit Amazon Linux 2015.03 v2.0.3 running Go 1.4",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Option Settings
    /// 
    /// Some options can be stack-specific, check [AWS Docs](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options-general.html)
    /// for supported options and examples.
    /// 
    /// The `setting` and `all_settings` mappings support the following format:
    /// 
    /// * `namespace` - unique namespace identifying the option's associated AWS resource
    /// * `name` - name of the configuration option
    /// * `value` - value for the configuration option
    /// * `resource` - (Optional) resource name for [scheduled action](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options-general.html#command-options-general-autoscalingscheduledaction)
    /// 
    /// ### Example With Options
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var tftest = new Aws.ElasticBeanstalk.Application("tftest", new Aws.ElasticBeanstalk.ApplicationArgs
    ///         {
    ///             Description = "tf-test-desc",
    ///         });
    ///         var tfenvtest = new Aws.ElasticBeanstalk.Environment("tfenvtest", new Aws.ElasticBeanstalk.EnvironmentArgs
    ///         {
    ///             Application = tftest.Name,
    ///             SolutionStackName = "64bit Amazon Linux 2015.03 v2.0.3 running Go 1.4",
    ///             Settings = 
    ///             {
    ///                 new Aws.ElasticBeanstalk.Inputs.EnvironmentSettingArgs
    ///                 {
    ///                     Namespace = "aws:ec2:vpc",
    ///                     Name = "VPCId",
    ///                     Value = "vpc-xxxxxxxx",
    ///                 },
    ///                 new Aws.ElasticBeanstalk.Inputs.EnvironmentSettingArgs
    ///                 {
    ///                     Namespace = "aws:ec2:vpc",
    ///                     Name = "Subnets",
    ///                     Value = "subnet-xxxxxxxx",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Elastic Beanstalk Environments can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:elasticbeanstalk/environment:Environment prodenv e-rpqsewtp2j
    /// ```
    /// </summary>
    [AwsResourceType("aws:elasticbeanstalk/environment:Environment")]
    public partial class Environment : Pulumi.CustomResource
    {
        /// <summary>
        /// List of all option settings configured in this Environment. These
        /// are a combination of default settings and their overrides from `setting` in
        /// the configuration.
        /// </summary>
        [Output("allSettings")]
        public Output<ImmutableArray<Outputs.EnvironmentAllSetting>> AllSettings { get; private set; } = null!;

        /// <summary>
        /// Name of the application that contains the version
        /// to be deployed
        /// </summary>
        [Output("application")]
        public Output<string> Application { get; private set; } = null!;

        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The autoscaling groups used by this Environment.
        /// </summary>
        [Output("autoscalingGroups")]
        public Output<ImmutableArray<string>> AutoscalingGroups { get; private set; } = null!;

        /// <summary>
        /// Fully qualified DNS name for this Environment.
        /// </summary>
        [Output("cname")]
        public Output<string> Cname { get; private set; } = null!;

        /// <summary>
        /// Prefix to use for the fully qualified DNS name of
        /// the Environment.
        /// </summary>
        [Output("cnamePrefix")]
        public Output<string> CnamePrefix { get; private set; } = null!;

        /// <summary>
        /// Short description of the Environment
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The URL to the Load Balancer for this Environment
        /// </summary>
        [Output("endpointUrl")]
        public Output<string> EndpointUrl { get; private set; } = null!;

        /// <summary>
        /// Instances used by this Environment.
        /// </summary>
        [Output("instances")]
        public Output<ImmutableArray<string>> Instances { get; private set; } = null!;

        /// <summary>
        /// Launch configurations in use by this Environment.
        /// </summary>
        [Output("launchConfigurations")]
        public Output<ImmutableArray<string>> LaunchConfigurations { get; private set; } = null!;

        /// <summary>
        /// Elastic load balancers in use by this Environment.
        /// </summary>
        [Output("loadBalancers")]
        public Output<ImmutableArray<string>> LoadBalancers { get; private set; } = null!;

        /// <summary>
        /// A unique name for this Environment. This name is used
        /// in the application URL
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the Elastic Beanstalk [Platform](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-platformarn)
        /// to use in deployment
        /// </summary>
        [Output("platformArn")]
        public Output<string> PlatformArn { get; private set; } = null!;

        /// <summary>
        /// The time between polling the AWS API to
        /// check if changes have been applied. Use this to adjust the rate of API calls
        /// for any `create` or `update` action. Minimum `10s`, maximum `180s`. Omit this to
        /// use the default behavior, which is an exponential backoff
        /// </summary>
        [Output("pollInterval")]
        public Output<string?> PollInterval { get; private set; } = null!;

        /// <summary>
        /// SQS queues in use by this Environment.
        /// </summary>
        [Output("queues")]
        public Output<ImmutableArray<string>> Queues { get; private set; } = null!;

        /// <summary>
        /// Option settings to configure the new Environment. These
        /// override specific values that are set as defaults. The format is detailed
        /// below in Option Settings
        /// </summary>
        [Output("settings")]
        public Output<ImmutableArray<Outputs.EnvironmentSetting>> Settings { get; private set; } = null!;

        /// <summary>
        /// A solution stack to base your environment
        /// off of. Example stacks can be found in the [Amazon API documentation](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html)
        /// </summary>
        [Output("solutionStackName")]
        public Output<string> SolutionStackName { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The name of the Elastic Beanstalk Configuration
        /// template to use in deployment
        /// </summary>
        [Output("templateName")]
        public Output<string?> TemplateName { get; private set; } = null!;

        /// <summary>
        /// Elastic Beanstalk Environment tier. Valid values are `Worker`
        /// or `WebServer`. If tier is left blank `WebServer` will be used.
        /// </summary>
        [Output("tier")]
        public Output<string?> Tier { get; private set; } = null!;

        /// <summary>
        /// Autoscaling triggers in use by this Environment.
        /// </summary>
        [Output("triggers")]
        public Output<ImmutableArray<string>> Triggers { get; private set; } = null!;

        /// <summary>
        /// The name of the Elastic Beanstalk Application Version
        /// to use in deployment.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// The maximum
        /// [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
        /// wait for an Elastic Beanstalk Environment to be in a ready state before timing
        /// out.
        /// </summary>
        [Output("waitForReadyTimeout")]
        public Output<string?> WaitForReadyTimeout { get; private set; } = null!;


        /// <summary>
        /// Create a Environment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Environment(string name, EnvironmentArgs args, CustomResourceOptions? options = null)
            : base("aws:elasticbeanstalk/environment:Environment", name, args ?? new EnvironmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Environment(string name, Input<string> id, EnvironmentState? state = null, CustomResourceOptions? options = null)
            : base("aws:elasticbeanstalk/environment:Environment", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Environment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Environment Get(string name, Input<string> id, EnvironmentState? state = null, CustomResourceOptions? options = null)
        {
            return new Environment(name, id, state, options);
        }
    }

    public sealed class EnvironmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the application that contains the version
        /// to be deployed
        /// </summary>
        [Input("application", required: true)]
        public Input<string> Application { get; set; } = null!;

        /// <summary>
        /// Prefix to use for the fully qualified DNS name of
        /// the Environment.
        /// </summary>
        [Input("cnamePrefix")]
        public Input<string>? CnamePrefix { get; set; }

        /// <summary>
        /// Short description of the Environment
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A unique name for this Environment. This name is used
        /// in the application URL
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the Elastic Beanstalk [Platform](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-platformarn)
        /// to use in deployment
        /// </summary>
        [Input("platformArn")]
        public Input<string>? PlatformArn { get; set; }

        /// <summary>
        /// The time between polling the AWS API to
        /// check if changes have been applied. Use this to adjust the rate of API calls
        /// for any `create` or `update` action. Minimum `10s`, maximum `180s`. Omit this to
        /// use the default behavior, which is an exponential backoff
        /// </summary>
        [Input("pollInterval")]
        public Input<string>? PollInterval { get; set; }

        [Input("settings")]
        private InputList<Inputs.EnvironmentSettingArgs>? _settings;

        /// <summary>
        /// Option settings to configure the new Environment. These
        /// override specific values that are set as defaults. The format is detailed
        /// below in Option Settings
        /// </summary>
        public InputList<Inputs.EnvironmentSettingArgs> Settings
        {
            get => _settings ?? (_settings = new InputList<Inputs.EnvironmentSettingArgs>());
            set => _settings = value;
        }

        /// <summary>
        /// A solution stack to base your environment
        /// off of. Example stacks can be found in the [Amazon API documentation](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html)
        /// </summary>
        [Input("solutionStackName")]
        public Input<string>? SolutionStackName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The name of the Elastic Beanstalk Configuration
        /// template to use in deployment
        /// </summary>
        [Input("templateName")]
        public Input<string>? TemplateName { get; set; }

        /// <summary>
        /// Elastic Beanstalk Environment tier. Valid values are `Worker`
        /// or `WebServer`. If tier is left blank `WebServer` will be used.
        /// </summary>
        [Input("tier")]
        public Input<string>? Tier { get; set; }

        /// <summary>
        /// The name of the Elastic Beanstalk Application Version
        /// to use in deployment.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The maximum
        /// [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
        /// wait for an Elastic Beanstalk Environment to be in a ready state before timing
        /// out.
        /// </summary>
        [Input("waitForReadyTimeout")]
        public Input<string>? WaitForReadyTimeout { get; set; }

        public EnvironmentArgs()
        {
        }
    }

    public sealed class EnvironmentState : Pulumi.ResourceArgs
    {
        [Input("allSettings")]
        private InputList<Inputs.EnvironmentAllSettingGetArgs>? _allSettings;

        /// <summary>
        /// List of all option settings configured in this Environment. These
        /// are a combination of default settings and their overrides from `setting` in
        /// the configuration.
        /// </summary>
        public InputList<Inputs.EnvironmentAllSettingGetArgs> AllSettings
        {
            get => _allSettings ?? (_allSettings = new InputList<Inputs.EnvironmentAllSettingGetArgs>());
            set => _allSettings = value;
        }

        /// <summary>
        /// Name of the application that contains the version
        /// to be deployed
        /// </summary>
        [Input("application")]
        public Input<string>? Application { get; set; }

        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("autoscalingGroups")]
        private InputList<string>? _autoscalingGroups;

        /// <summary>
        /// The autoscaling groups used by this Environment.
        /// </summary>
        public InputList<string> AutoscalingGroups
        {
            get => _autoscalingGroups ?? (_autoscalingGroups = new InputList<string>());
            set => _autoscalingGroups = value;
        }

        /// <summary>
        /// Fully qualified DNS name for this Environment.
        /// </summary>
        [Input("cname")]
        public Input<string>? Cname { get; set; }

        /// <summary>
        /// Prefix to use for the fully qualified DNS name of
        /// the Environment.
        /// </summary>
        [Input("cnamePrefix")]
        public Input<string>? CnamePrefix { get; set; }

        /// <summary>
        /// Short description of the Environment
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The URL to the Load Balancer for this Environment
        /// </summary>
        [Input("endpointUrl")]
        public Input<string>? EndpointUrl { get; set; }

        [Input("instances")]
        private InputList<string>? _instances;

        /// <summary>
        /// Instances used by this Environment.
        /// </summary>
        public InputList<string> Instances
        {
            get => _instances ?? (_instances = new InputList<string>());
            set => _instances = value;
        }

        [Input("launchConfigurations")]
        private InputList<string>? _launchConfigurations;

        /// <summary>
        /// Launch configurations in use by this Environment.
        /// </summary>
        public InputList<string> LaunchConfigurations
        {
            get => _launchConfigurations ?? (_launchConfigurations = new InputList<string>());
            set => _launchConfigurations = value;
        }

        [Input("loadBalancers")]
        private InputList<string>? _loadBalancers;

        /// <summary>
        /// Elastic load balancers in use by this Environment.
        /// </summary>
        public InputList<string> LoadBalancers
        {
            get => _loadBalancers ?? (_loadBalancers = new InputList<string>());
            set => _loadBalancers = value;
        }

        /// <summary>
        /// A unique name for this Environment. This name is used
        /// in the application URL
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the Elastic Beanstalk [Platform](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-platformarn)
        /// to use in deployment
        /// </summary>
        [Input("platformArn")]
        public Input<string>? PlatformArn { get; set; }

        /// <summary>
        /// The time between polling the AWS API to
        /// check if changes have been applied. Use this to adjust the rate of API calls
        /// for any `create` or `update` action. Minimum `10s`, maximum `180s`. Omit this to
        /// use the default behavior, which is an exponential backoff
        /// </summary>
        [Input("pollInterval")]
        public Input<string>? PollInterval { get; set; }

        [Input("queues")]
        private InputList<string>? _queues;

        /// <summary>
        /// SQS queues in use by this Environment.
        /// </summary>
        public InputList<string> Queues
        {
            get => _queues ?? (_queues = new InputList<string>());
            set => _queues = value;
        }

        [Input("settings")]
        private InputList<Inputs.EnvironmentSettingGetArgs>? _settings;

        /// <summary>
        /// Option settings to configure the new Environment. These
        /// override specific values that are set as defaults. The format is detailed
        /// below in Option Settings
        /// </summary>
        public InputList<Inputs.EnvironmentSettingGetArgs> Settings
        {
            get => _settings ?? (_settings = new InputList<Inputs.EnvironmentSettingGetArgs>());
            set => _settings = value;
        }

        /// <summary>
        /// A solution stack to base your environment
        /// off of. Example stacks can be found in the [Amazon API documentation](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html)
        /// </summary>
        [Input("solutionStackName")]
        public Input<string>? SolutionStackName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// The name of the Elastic Beanstalk Configuration
        /// template to use in deployment
        /// </summary>
        [Input("templateName")]
        public Input<string>? TemplateName { get; set; }

        /// <summary>
        /// Elastic Beanstalk Environment tier. Valid values are `Worker`
        /// or `WebServer`. If tier is left blank `WebServer` will be used.
        /// </summary>
        [Input("tier")]
        public Input<string>? Tier { get; set; }

        [Input("triggers")]
        private InputList<string>? _triggers;

        /// <summary>
        /// Autoscaling triggers in use by this Environment.
        /// </summary>
        public InputList<string> Triggers
        {
            get => _triggers ?? (_triggers = new InputList<string>());
            set => _triggers = value;
        }

        /// <summary>
        /// The name of the Elastic Beanstalk Application Version
        /// to use in deployment.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The maximum
        /// [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
        /// wait for an Elastic Beanstalk Environment to be in a ready state before timing
        /// out.
        /// </summary>
        [Input("waitForReadyTimeout")]
        public Input<string>? WaitForReadyTimeout { get; set; }

        public EnvironmentState()
        {
        }
    }
}
