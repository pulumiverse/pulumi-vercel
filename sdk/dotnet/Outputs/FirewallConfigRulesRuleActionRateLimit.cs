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
    public sealed class FirewallConfigRulesRuleActionRateLimit
    {
        /// <summary>
        /// Action to take when rate limit is exceeded
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// Rate limiting algorithm
        /// </summary>
        public readonly string Algo;
        /// <summary>
        /// Keys used to bucket an individual client
        /// </summary>
        public readonly ImmutableArray<string> Keys;
        /// <summary>
        /// number of requests allowed in the window
        /// </summary>
        public readonly int Limit;
        /// <summary>
        /// Time window in seconds
        /// </summary>
        public readonly int Window;

        [OutputConstructor]
        private FirewallConfigRulesRuleActionRateLimit(
            string action,

            string algo,

            ImmutableArray<string> keys,

            int limit,

            int window)
        {
            Action = action;
            Algo = algo;
            Keys = keys;
            Limit = limit;
            Window = window;
        }
    }
}
