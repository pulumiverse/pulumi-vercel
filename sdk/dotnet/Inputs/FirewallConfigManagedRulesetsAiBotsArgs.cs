// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Vercel.Inputs
{

    public sealed class FirewallConfigManagedRulesetsAiBotsArgs : global::Pulumi.ResourceArgs
    {
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("active")]
        public Input<bool>? Active { get; set; }

        public FirewallConfigManagedRulesetsAiBotsArgs()
        {
        }
        public static new FirewallConfigManagedRulesetsAiBotsArgs Empty => new FirewallConfigManagedRulesetsAiBotsArgs();
    }
}
