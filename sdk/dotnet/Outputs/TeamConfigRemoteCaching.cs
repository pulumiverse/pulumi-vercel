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
    public sealed class TeamConfigRemoteCaching
    {
        /// <summary>
        /// Indicates if Remote Caching is enabled.
        /// </summary>
        public readonly bool? Enabled;

        [OutputConstructor]
        private TeamConfigRemoteCaching(bool? enabled)
        {
            Enabled = enabled;
        }
    }
}
