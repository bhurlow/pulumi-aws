// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Quicksight.Inputs
{

    public sealed class DashboardParametersArgs : global::Pulumi.ResourceArgs
    {
        [Input("dateTimeParameters")]
        private InputList<Inputs.DashboardParametersDateTimeParameterArgs>? _dateTimeParameters;

        /// <summary>
        /// A list of parameters that have a data type of date-time. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DateTimeParameter.html).
        /// </summary>
        public InputList<Inputs.DashboardParametersDateTimeParameterArgs> DateTimeParameters
        {
            get => _dateTimeParameters ?? (_dateTimeParameters = new InputList<Inputs.DashboardParametersDateTimeParameterArgs>());
            set => _dateTimeParameters = value;
        }

        [Input("decimalParameters")]
        private InputList<Inputs.DashboardParametersDecimalParameterArgs>? _decimalParameters;

        /// <summary>
        /// A list of parameters that have a data type of decimal. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DecimalParameter.html).
        /// </summary>
        public InputList<Inputs.DashboardParametersDecimalParameterArgs> DecimalParameters
        {
            get => _decimalParameters ?? (_decimalParameters = new InputList<Inputs.DashboardParametersDecimalParameterArgs>());
            set => _decimalParameters = value;
        }

        [Input("integerParameters")]
        private InputList<Inputs.DashboardParametersIntegerParameterArgs>? _integerParameters;

        /// <summary>
        /// A list of parameters that have a data type of integer. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_IntegerParameter.html).
        /// </summary>
        public InputList<Inputs.DashboardParametersIntegerParameterArgs> IntegerParameters
        {
            get => _integerParameters ?? (_integerParameters = new InputList<Inputs.DashboardParametersIntegerParameterArgs>());
            set => _integerParameters = value;
        }

        [Input("stringParameters")]
        private InputList<Inputs.DashboardParametersStringParameterArgs>? _stringParameters;

        /// <summary>
        /// A list of parameters that have a data type of string. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_StringParameter.html).
        /// </summary>
        public InputList<Inputs.DashboardParametersStringParameterArgs> StringParameters
        {
            get => _stringParameters ?? (_stringParameters = new InputList<Inputs.DashboardParametersStringParameterArgs>());
            set => _stringParameters = value;
        }

        public DashboardParametersArgs()
        {
        }
        public static new DashboardParametersArgs Empty => new DashboardParametersArgs();
    }
}
