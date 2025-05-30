// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides a DNS Record resource.
 *
 * DNS records are instructions that live in authoritative DNS servers and provide information about a domain.
 *
 * > The `value` field must be specified on all DNS record types except `SRV`. When using `SRV` DNS records, the `srv` field must be specified.
 *
 * For more detailed information, please see the [Vercel documentation](https://vercel.com/docs/concepts/projects/custom-domains#dns-records)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vercel from "@pulumiverse/vercel";
 *
 * const a = new vercel.DnsRecord("a", {
 *     domain: "example.com",
 *     name: "subdomain",
 *     type: "A",
 *     ttl: 60,
 *     value: "192.168.0.1",
 * });
 * const aaaa = new vercel.DnsRecord("aaaa", {
 *     domain: "example.com",
 *     name: "subdomain",
 *     type: "AAAA",
 *     ttl: 60,
 *     value: "::0",
 * });
 * const alias = new vercel.DnsRecord("alias", {
 *     domain: "example.com",
 *     name: "subdomain",
 *     type: "ALIAS",
 *     ttl: 60,
 *     value: "example2.com.",
 * });
 * const caa = new vercel.DnsRecord("caa", {
 *     domain: "example.com",
 *     name: "subdomain",
 *     type: "CAA",
 *     ttl: 60,
 *     value: "1 issue \"letsencrypt.org\"",
 * });
 * const cname = new vercel.DnsRecord("cname", {
 *     domain: "example.com",
 *     name: "subdomain",
 *     type: "CNAME",
 *     ttl: 60,
 *     value: "example2.com.",
 * });
 * const mx = new vercel.DnsRecord("mx", {
 *     domain: "example.com",
 *     name: "subdomain",
 *     type: "MX",
 *     ttl: 60,
 *     mxPriority: 333,
 *     value: "example2.com.",
 * });
 * const srv = new vercel.DnsRecord("srv", {
 *     domain: "example.com",
 *     name: "subdomain",
 *     type: "SRV",
 *     ttl: 60,
 *     srv: {
 *         port: 6000,
 *         weight: 60,
 *         priority: 127,
 *         target: "example2.com.",
 *     },
 * });
 * const txt = new vercel.DnsRecord("txt", {
 *     domain: "example.com",
 *     name: "subdomain",
 *     type: "TXT",
 *     ttl: 60,
 *     value: "some text value",
 * });
 * ```
 *
 * ## Import
 *
 * If importing into a personal account, or with a team configured on
 *
 * the provider, simply use the record id.
 *
 * - record_id can be taken from the network tab inside developer tools, while you are on the domains page,
 *
 * or can be queried from the Vercel API directly (https://vercel.com/docs/rest-api/endpoints/dns#list-existing-dns-records).
 *
 * ```sh
 * $ pulumi import vercel:index/dnsRecord:DnsRecord example rec_xxxxxxxxxxxxxxxxxxxxxxxxxxxx
 * ```
 *
 * Alternatively, you can import via the team_id and record_id.
 *
 * - team_id can be found in the team `settings` tab in the Vercel UI.
 *
 * - record_id can be taken from the network tab inside developer tools, while you are on the domains page,
 *
 * or can be queried from the Vercel API directly (https://vercel.com/docs/rest-api/endpoints/dns#list-existing-dns-records).
 *
 * ```sh
 * $ pulumi import vercel:index/dnsRecord:DnsRecord example team_xxxxxxxxxxxxxxxxxxxxxxxx/rec_xxxxxxxxxxxxxxxxxxxxxxxxxxxx
 * ```
 */
export class DnsRecord extends pulumi.CustomResource {
    /**
     * Get an existing DnsRecord resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DnsRecordState, opts?: pulumi.CustomResourceOptions): DnsRecord {
        return new DnsRecord(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vercel:index/dnsRecord:DnsRecord';

    /**
     * Returns true if the given object is an instance of DnsRecord.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DnsRecord {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DnsRecord.__pulumiType;
    }

    /**
     * A comment explaining what the DNS record is for.
     */
    public readonly comment!: pulumi.Output<string>;
    /**
     * The domain name, or zone, that the DNS record should be created beneath.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * The priority of the MX record. The priority specifies the sequence that an email server receives emails. A smaller value indicates a higher priority.
     */
    public readonly mxPriority!: pulumi.Output<number | undefined>;
    /**
     * The subdomain name of the record. This should be an empty string if the rercord is for the root domain.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Settings for an SRV record.
     */
    public readonly srv!: pulumi.Output<outputs.DnsRecordSrv | undefined>;
    /**
     * The team ID that the domain and DNS records belong to. Required when configuring a team resource if a default team has not been set in the provider.
     */
    public readonly teamId!: pulumi.Output<string>;
    /**
     * The TTL value in seconds. Must be a number between 60 and 2147483647. If unspecified, it will default to 60 seconds.
     */
    public readonly ttl!: pulumi.Output<number>;
    /**
     * The type of DNS record. Available types: `A`, `AAAA`, `ALIAS`, `CAA`, `CNAME`, `MX`, `NS`, `SRV`, `TXT`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The value of the DNS record. The format depends on the 'type' property.
     * For an 'A' record, this should be a valid IPv4 address.
     * For an 'AAAA' record, this should be an IPv6 address.
     * For 'ALIAS' records, this should be a hostname.
     * For 'CAA' records, this should specify specify which Certificate Authorities (CAs) are allowed to issue certificates for the domain.
     * For 'CNAME' records, this should be a different domain name.
     * For 'MX' records, this should specify the mail server responsible for accepting messages on behalf of the domain name.
     * For 'TXT' records, this can contain arbitrary text.
     */
    public readonly value!: pulumi.Output<string | undefined>;

    /**
     * Create a DnsRecord resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DnsRecordArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DnsRecordArgs | DnsRecordState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DnsRecordState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["mxPriority"] = state ? state.mxPriority : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["srv"] = state ? state.srv : undefined;
            resourceInputs["teamId"] = state ? state.teamId : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as DnsRecordArgs | undefined;
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["mxPriority"] = args ? args.mxPriority : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["srv"] = args ? args.srv : undefined;
            resourceInputs["teamId"] = args ? args.teamId : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DnsRecord.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DnsRecord resources.
 */
export interface DnsRecordState {
    /**
     * A comment explaining what the DNS record is for.
     */
    comment?: pulumi.Input<string>;
    /**
     * The domain name, or zone, that the DNS record should be created beneath.
     */
    domain?: pulumi.Input<string>;
    /**
     * The priority of the MX record. The priority specifies the sequence that an email server receives emails. A smaller value indicates a higher priority.
     */
    mxPriority?: pulumi.Input<number>;
    /**
     * The subdomain name of the record. This should be an empty string if the rercord is for the root domain.
     */
    name?: pulumi.Input<string>;
    /**
     * Settings for an SRV record.
     */
    srv?: pulumi.Input<inputs.DnsRecordSrv>;
    /**
     * The team ID that the domain and DNS records belong to. Required when configuring a team resource if a default team has not been set in the provider.
     */
    teamId?: pulumi.Input<string>;
    /**
     * The TTL value in seconds. Must be a number between 60 and 2147483647. If unspecified, it will default to 60 seconds.
     */
    ttl?: pulumi.Input<number>;
    /**
     * The type of DNS record. Available types: `A`, `AAAA`, `ALIAS`, `CAA`, `CNAME`, `MX`, `NS`, `SRV`, `TXT`.
     */
    type?: pulumi.Input<string>;
    /**
     * The value of the DNS record. The format depends on the 'type' property.
     * For an 'A' record, this should be a valid IPv4 address.
     * For an 'AAAA' record, this should be an IPv6 address.
     * For 'ALIAS' records, this should be a hostname.
     * For 'CAA' records, this should specify specify which Certificate Authorities (CAs) are allowed to issue certificates for the domain.
     * For 'CNAME' records, this should be a different domain name.
     * For 'MX' records, this should specify the mail server responsible for accepting messages on behalf of the domain name.
     * For 'TXT' records, this can contain arbitrary text.
     */
    value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DnsRecord resource.
 */
export interface DnsRecordArgs {
    /**
     * A comment explaining what the DNS record is for.
     */
    comment?: pulumi.Input<string>;
    /**
     * The domain name, or zone, that the DNS record should be created beneath.
     */
    domain: pulumi.Input<string>;
    /**
     * The priority of the MX record. The priority specifies the sequence that an email server receives emails. A smaller value indicates a higher priority.
     */
    mxPriority?: pulumi.Input<number>;
    /**
     * The subdomain name of the record. This should be an empty string if the rercord is for the root domain.
     */
    name?: pulumi.Input<string>;
    /**
     * Settings for an SRV record.
     */
    srv?: pulumi.Input<inputs.DnsRecordSrv>;
    /**
     * The team ID that the domain and DNS records belong to. Required when configuring a team resource if a default team has not been set in the provider.
     */
    teamId?: pulumi.Input<string>;
    /**
     * The TTL value in seconds. Must be a number between 60 and 2147483647. If unspecified, it will default to 60 seconds.
     */
    ttl?: pulumi.Input<number>;
    /**
     * The type of DNS record. Available types: `A`, `AAAA`, `ALIAS`, `CAA`, `CNAME`, `MX`, `NS`, `SRV`, `TXT`.
     */
    type: pulumi.Input<string>;
    /**
     * The value of the DNS record. The format depends on the 'type' property.
     * For an 'A' record, this should be a valid IPv4 address.
     * For an 'AAAA' record, this should be an IPv6 address.
     * For 'ALIAS' records, this should be a hostname.
     * For 'CAA' records, this should specify specify which Certificate Authorities (CAs) are allowed to issue certificates for the domain.
     * For 'CNAME' records, this should be a different domain name.
     * For 'MX' records, this should specify the mail server responsible for accepting messages on behalf of the domain name.
     * For 'TXT' records, this can contain arbitrary text.
     */
    value?: pulumi.Input<string>;
}
