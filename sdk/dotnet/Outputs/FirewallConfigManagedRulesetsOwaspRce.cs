// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Vercel.Outputs
{

    [OutputType]
    public sealed class FirewallConfigManagedRulesetsOwaspRce
    {
        public readonly string Action;
        public readonly bool? Active;

        [OutputConstructor]
        private FirewallConfigManagedRulesetsOwaspRce(
            string action,

            bool? active)
        {
            Action = action;
            Active = active;
        }
    }
}
