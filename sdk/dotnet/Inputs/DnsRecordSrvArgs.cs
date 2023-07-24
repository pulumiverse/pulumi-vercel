// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vercel.Inputs
{

    public sealed class DnsRecordSrvArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The TCP or UDP port on which the service is to be found.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// The priority of the target host, lower value means more preferred.
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        /// <summary>
        /// The canonical hostname of the machine providing the service, ending in a dot.
        /// </summary>
        [Input("target", required: true)]
        public Input<string> Target { get; set; } = null!;

        /// <summary>
        /// A relative weight for records with the same priority, higher value means higher chance of getting picked.
        /// </summary>
        [Input("weight", required: true)]
        public Input<int> Weight { get; set; } = null!;

        public DnsRecordSrvArgs()
        {
        }
        public static new DnsRecordSrvArgs Empty => new DnsRecordSrvArgs();
    }
}
