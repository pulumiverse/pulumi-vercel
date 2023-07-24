// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vercel.Outputs
{

    [OutputType]
    public sealed class DnsRecordSrv
    {
        /// <summary>
        /// The TCP or UDP port on which the service is to be found.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The priority of the target host, lower value means more preferred.
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// The canonical hostname of the machine providing the service, ending in a dot.
        /// </summary>
        public readonly string Target;
        /// <summary>
        /// A relative weight for records with the same priority, higher value means higher chance of getting picked.
        /// </summary>
        public readonly int Weight;

        [OutputConstructor]
        private DnsRecordSrv(
            int port,

            int priority,

            string target,

            int weight)
        {
            Port = port;
            Priority = priority;
            Target = target;
            Weight = weight;
        }
    }
}
