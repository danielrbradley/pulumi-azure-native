// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.StreamAnalytics.Inputs
{

    /// <summary>
    /// The properties that are associated with a streaming job.
    /// </summary>
    public sealed class StreamingJobPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Controls certain runtime behaviors of the streaming job.
        /// </summary>
        [Input("compatibilityLevel")]
        public Input<string>? CompatibilityLevel { get; set; }

        /// <summary>
        /// The data locale of the stream analytics job. Value should be the name of a supported .NET Culture from the set https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx. Defaults to 'en-US' if none specified.
        /// </summary>
        [Input("dataLocale")]
        public Input<string>? DataLocale { get; set; }

        /// <summary>
        /// The maximum tolerable delay in seconds where events arriving late could be included.  Supported range is -1 to 1814399 (20.23:59:59 days) and -1 is used to specify wait indefinitely. If the property is absent, it is interpreted to have a value of -1.
        /// </summary>
        [Input("eventsLateArrivalMaxDelayInSeconds")]
        public Input<int>? EventsLateArrivalMaxDelayInSeconds { get; set; }

        /// <summary>
        /// The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.
        /// </summary>
        [Input("eventsOutOfOrderMaxDelayInSeconds")]
        public Input<int>? EventsOutOfOrderMaxDelayInSeconds { get; set; }

        /// <summary>
        /// Indicates the policy to apply to events that arrive out of order in the input event stream.
        /// </summary>
        [Input("eventsOutOfOrderPolicy")]
        public Input<string>? EventsOutOfOrderPolicy { get; set; }

        [Input("functions")]
        private InputList<Inputs.FunctionArgs>? _functions;

        /// <summary>
        /// A list of one or more functions for the streaming job. The name property for each function is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
        /// </summary>
        public InputList<Inputs.FunctionArgs> Functions
        {
            get => _functions ?? (_functions = new InputList<Inputs.FunctionArgs>());
            set => _functions = value;
        }

        [Input("inputs")]
        private InputList<Inputs.InputArgs>? _inputs;

        /// <summary>
        /// A list of one or more inputs to the streaming job. The name property for each input is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual input.
        /// </summary>
        public InputList<Inputs.InputArgs> Inputs
        {
            get => _inputs ?? (_inputs = new InputList<Inputs.InputArgs>());
            set => _inputs = value;
        }

        /// <summary>
        /// Indicates the policy to apply to events that arrive at the output and cannot be written to the external storage due to being malformed (missing column values, column values of wrong type or size).
        /// </summary>
        [Input("outputErrorPolicy")]
        public Input<string>? OutputErrorPolicy { get; set; }

        /// <summary>
        /// This property should only be utilized when it is desired that the job be started immediately upon creation. Value may be JobStartTime, CustomTime, or LastOutputEventTime to indicate whether the starting point of the output event stream should start whenever the job is started, start at a custom user time stamp specified via the outputStartTime property, or start from the last event output time.
        /// </summary>
        [Input("outputStartMode")]
        public Input<string>? OutputStartMode { get; set; }

        /// <summary>
        /// Value is either an ISO-8601 formatted time stamp that indicates the starting point of the output event stream, or null to indicate that the output event stream will start whenever the streaming job is started. This property must have a value if outputStartMode is set to CustomTime.
        /// </summary>
        [Input("outputStartTime")]
        public Input<string>? OutputStartTime { get; set; }

        [Input("outputs")]
        private InputList<Inputs.OutputArgs>? _outputs;

        /// <summary>
        /// A list of one or more outputs for the streaming job. The name property for each output is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual output.
        /// </summary>
        public InputList<Inputs.OutputArgs> Outputs
        {
            get => _outputs ?? (_outputs = new InputList<Inputs.OutputArgs>());
            set => _outputs = value;
        }

        /// <summary>
        /// Describes the SKU of the streaming job. Required on PUT (CreateOrReplace) requests.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.SkuArgs>? Sku { get; set; }

        /// <summary>
        /// Indicates the query and the number of streaming units to use for the streaming job. The name property of the transformation is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
        /// </summary>
        [Input("transformation")]
        public Input<Inputs.TransformationArgs>? Transformation { get; set; }

        public StreamingJobPropertiesArgs()
        {
        }
    }
}