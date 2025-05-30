// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides an Edge Config Item.
 *
 * An Edge Config is a global data store that enables experimentation with feature flags, A/B testing, critical redirects, and more.
 *
 * An Edge Config Item is a value within an Edge Config.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vercel from "@pulumiverse/vercel";
 *
 * const example = new vercel.EdgeConfig("example", {name: "example"});
 * const exampleEdgeConfigItem = new vercel.EdgeConfigItem("example", {
 *     edgeConfigId: example.id,
 *     key: "foobar",
 *     value: "baz",
 * });
 * ```
 *
 * ## Import
 *
 * If importing into a personal account, or with a team configured on
 *
 * the provider, simply use the edge config id and the key of the item to import.
 *
 * - edge_config_id can be found by navigating to the Edge Config in the Vercel UI. It should begin with `ecfg_`.
 *
 * - key is the key of teh item to import.
 *
 * ```sh
 * $ pulumi import vercel:index/edgeConfigItem:EdgeConfigItem example ecfg_xxxxxxxxxxxxxxxxxxxxxxxxxxxx/example_key
 * ```
 *
 * Alternatively, you can import via the team_id, edge_config_id and the key of the item to import.
 *
 * - team_id can be found in the team `settings` tab in the Vercel UI.
 *
 * - edge_config_id can be found by navigating to the Edge Config in the Vercel UI. It should begin with `ecfg_`.
 *
 * - key is the key of the item to import.
 *
 * ```sh
 * $ pulumi import vercel:index/edgeConfigItem:EdgeConfigItem example team_xxxxxxxxxxxxxxxxxxxxxxxx/ecfg_xxxxxxxxxxxxxxxxxxxxxxxxxxxx/example_key
 * ```
 */
export class EdgeConfigItem extends pulumi.CustomResource {
    /**
     * Get an existing EdgeConfigItem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EdgeConfigItemState, opts?: pulumi.CustomResourceOptions): EdgeConfigItem {
        return new EdgeConfigItem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vercel:index/edgeConfigItem:EdgeConfigItem';

    /**
     * Returns true if the given object is an instance of EdgeConfigItem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EdgeConfigItem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EdgeConfigItem.__pulumiType;
    }

    /**
     * The ID of the Edge Config store.
     */
    public readonly edgeConfigId!: pulumi.Output<string>;
    /**
     * The name of the key you want to add to or update within your Edge Config.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * The ID of the team the Edge Config should exist under. Required when configuring a team resource if a default team has not been set in the provider.
     */
    public readonly teamId!: pulumi.Output<string>;
    /**
     * The value you want to assign to the key.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a EdgeConfigItem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EdgeConfigItemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EdgeConfigItemArgs | EdgeConfigItemState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EdgeConfigItemState | undefined;
            resourceInputs["edgeConfigId"] = state ? state.edgeConfigId : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["teamId"] = state ? state.teamId : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as EdgeConfigItemArgs | undefined;
            if ((!args || args.edgeConfigId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'edgeConfigId'");
            }
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            resourceInputs["edgeConfigId"] = args ? args.edgeConfigId : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["teamId"] = args ? args.teamId : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EdgeConfigItem.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EdgeConfigItem resources.
 */
export interface EdgeConfigItemState {
    /**
     * The ID of the Edge Config store.
     */
    edgeConfigId?: pulumi.Input<string>;
    /**
     * The name of the key you want to add to or update within your Edge Config.
     */
    key?: pulumi.Input<string>;
    /**
     * The ID of the team the Edge Config should exist under. Required when configuring a team resource if a default team has not been set in the provider.
     */
    teamId?: pulumi.Input<string>;
    /**
     * The value you want to assign to the key.
     */
    value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EdgeConfigItem resource.
 */
export interface EdgeConfigItemArgs {
    /**
     * The ID of the Edge Config store.
     */
    edgeConfigId: pulumi.Input<string>;
    /**
     * The name of the key you want to add to or update within your Edge Config.
     */
    key: pulumi.Input<string>;
    /**
     * The ID of the team the Edge Config should exist under. Required when configuring a team resource if a default team has not been set in the provider.
     */
    teamId?: pulumi.Input<string>;
    /**
     * The value you want to assign to the key.
     */
    value: pulumi.Input<string>;
}
