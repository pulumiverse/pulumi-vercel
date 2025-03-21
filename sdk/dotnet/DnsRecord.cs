// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Vercel
{
    [VercelResourceType("vercel:index/dnsRecord:DnsRecord")]
    public partial class DnsRecord : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A comment explaining what the DNS record is for.
        /// </summary>
        [Output("comment")]
        public Output<string> Comment { get; private set; } = null!;

        /// <summary>
        /// The domain name, or zone, that the DNS record should be created beneath.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// The priority of the MX record. The priority specifies the sequence that an email server receives emails. A smaller value
        /// indicates a higher priority.
        /// </summary>
        [Output("mxPriority")]
        public Output<int?> MxPriority { get; private set; } = null!;

        /// <summary>
        /// The subdomain name of the record. This should be an empty string if the rercord is for the root domain.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Settings for an SRV record.
        /// </summary>
        [Output("srv")]
        public Output<Outputs.DnsRecordSrv?> Srv { get; private set; } = null!;

        /// <summary>
        /// The team ID that the domain and DNS records belong to. Required when configuring a team resource if a default team has
        /// not been set in the provider.
        /// </summary>
        [Output("teamId")]
        public Output<string> TeamId { get; private set; } = null!;

        /// <summary>
        /// The TTL value in seconds. Must be a number between 60 and 2147483647. If unspecified, it will default to 60 seconds.
        /// </summary>
        [Output("ttl")]
        public Output<int> Ttl { get; private set; } = null!;

        /// <summary>
        /// The type of DNS record. Available types: `A`, `AAAA`, `ALIAS`, `CAA`, `CNAME`, `MX`, `NS`, `SRV`, `TXT`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The value of the DNS record. The format depends on the 'type' property. For an 'A' record, this should be a valid IPv4
        /// address. For an 'AAAA' record, this should be an IPv6 address. For 'ALIAS' records, this should be a hostname. For 'CAA'
        /// records, this should specify specify which Certificate Authorities (CAs) are allowed to issue certificates for the
        /// domain. For 'CNAME' records, this should be a different domain name. For 'MX' records, this should specify the mail
        /// server responsible for accepting messages on behalf of the domain name. For 'TXT' records, this can contain arbitrary
        /// text.
        /// </summary>
        [Output("value")]
        public Output<string?> Value { get; private set; } = null!;


        /// <summary>
        /// Create a DnsRecord resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DnsRecord(string name, DnsRecordArgs args, CustomResourceOptions? options = null)
            : base("vercel:index/dnsRecord:DnsRecord", name, args ?? new DnsRecordArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DnsRecord(string name, Input<string> id, DnsRecordState? state = null, CustomResourceOptions? options = null)
            : base("vercel:index/dnsRecord:DnsRecord", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DnsRecord resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DnsRecord Get(string name, Input<string> id, DnsRecordState? state = null, CustomResourceOptions? options = null)
        {
            return new DnsRecord(name, id, state, options);
        }
    }

    public sealed class DnsRecordArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A comment explaining what the DNS record is for.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// The domain name, or zone, that the DNS record should be created beneath.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// The priority of the MX record. The priority specifies the sequence that an email server receives emails. A smaller value
        /// indicates a higher priority.
        /// </summary>
        [Input("mxPriority")]
        public Input<int>? MxPriority { get; set; }

        /// <summary>
        /// The subdomain name of the record. This should be an empty string if the rercord is for the root domain.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Settings for an SRV record.
        /// </summary>
        [Input("srv")]
        public Input<Inputs.DnsRecordSrvArgs>? Srv { get; set; }

        /// <summary>
        /// The team ID that the domain and DNS records belong to. Required when configuring a team resource if a default team has
        /// not been set in the provider.
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        /// <summary>
        /// The TTL value in seconds. Must be a number between 60 and 2147483647. If unspecified, it will default to 60 seconds.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// The type of DNS record. Available types: `A`, `AAAA`, `ALIAS`, `CAA`, `CNAME`, `MX`, `NS`, `SRV`, `TXT`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The value of the DNS record. The format depends on the 'type' property. For an 'A' record, this should be a valid IPv4
        /// address. For an 'AAAA' record, this should be an IPv6 address. For 'ALIAS' records, this should be a hostname. For 'CAA'
        /// records, this should specify specify which Certificate Authorities (CAs) are allowed to issue certificates for the
        /// domain. For 'CNAME' records, this should be a different domain name. For 'MX' records, this should specify the mail
        /// server responsible for accepting messages on behalf of the domain name. For 'TXT' records, this can contain arbitrary
        /// text.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public DnsRecordArgs()
        {
        }
        public static new DnsRecordArgs Empty => new DnsRecordArgs();
    }

    public sealed class DnsRecordState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A comment explaining what the DNS record is for.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// The domain name, or zone, that the DNS record should be created beneath.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The priority of the MX record. The priority specifies the sequence that an email server receives emails. A smaller value
        /// indicates a higher priority.
        /// </summary>
        [Input("mxPriority")]
        public Input<int>? MxPriority { get; set; }

        /// <summary>
        /// The subdomain name of the record. This should be an empty string if the rercord is for the root domain.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Settings for an SRV record.
        /// </summary>
        [Input("srv")]
        public Input<Inputs.DnsRecordSrvGetArgs>? Srv { get; set; }

        /// <summary>
        /// The team ID that the domain and DNS records belong to. Required when configuring a team resource if a default team has
        /// not been set in the provider.
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        /// <summary>
        /// The TTL value in seconds. Must be a number between 60 and 2147483647. If unspecified, it will default to 60 seconds.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// The type of DNS record. Available types: `A`, `AAAA`, `ALIAS`, `CAA`, `CNAME`, `MX`, `NS`, `SRV`, `TXT`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The value of the DNS record. The format depends on the 'type' property. For an 'A' record, this should be a valid IPv4
        /// address. For an 'AAAA' record, this should be an IPv6 address. For 'ALIAS' records, this should be a hostname. For 'CAA'
        /// records, this should specify specify which Certificate Authorities (CAs) are allowed to issue certificates for the
        /// domain. For 'CNAME' records, this should be a different domain name. For 'MX' records, this should specify the mail
        /// server responsible for accepting messages on behalf of the domain name. For 'TXT' records, this can contain arbitrary
        /// text.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public DnsRecordState()
        {
        }
        public static new DnsRecordState Empty => new DnsRecordState();
    }
}
