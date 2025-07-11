// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a Shared Environment Variable resource.
 *
 * A Shared Environment Variable resource defines an Environment Variable that can be shared between multiple Vercel Projects.
 *
 * For more detailed information, please see the [Vercel documentation](https://vercel.com/docs/concepts/projects/environment-variables/shared-environment-variables).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vercel from "@pulumiverse/vercel";
 *
 * const example = new vercel.Project("example", {
 *     name: "example",
 *     gitRepository: {
 *         type: "github",
 *         repo: "vercel/some-repo",
 *     },
 * });
 * // A shared environment variable that will be created
 * // and associated with the "example" project.
 * const exampleSharedEnvironmentVariable = new vercel.SharedEnvironmentVariable("example", {
 *     key: "EXAMPLE",
 *     value: "some_value",
 *     targets: ["production"],
 *     comment: "an example shared variable",
 *     projectIds: [example.id],
 * });
 * ```
 *
 * ## Import
 *
 * You can import via the team_id and environment variable id.
 *
 * - team_id can be found in the team `settings` tab in the Vercel UI.
 *
 * - environment variable id can be taken from the network tab inside developer tools, while you are on the project page.
 *
 * # 
 *
 * Note also, that the value field for sensitive environment variables will be imported as `null`.
 *
 * ```sh
 * $ pulumi import vercel:index/sharedEnvironmentVariable:SharedEnvironmentVariable example team_xxxxxxxxxxxxxxxxxxxxxxxx/env_yyyyyyyyyyyyy
 * ```
 */
export class SharedEnvironmentVariable extends pulumi.CustomResource {
    /**
     * Get an existing SharedEnvironmentVariable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SharedEnvironmentVariableState, opts?: pulumi.CustomResourceOptions): SharedEnvironmentVariable {
        return new SharedEnvironmentVariable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vercel:index/sharedEnvironmentVariable:SharedEnvironmentVariable';

    /**
     * Returns true if the given object is an instance of SharedEnvironmentVariable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SharedEnvironmentVariable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SharedEnvironmentVariable.__pulumiType;
    }

    /**
     * Whether the shared environment variable should be applied to all custom environments in the linked projects.
     */
    public readonly applyToAllCustomEnvironments!: pulumi.Output<boolean>;
    /**
     * A comment explaining what the environment variable is for.
     */
    public readonly comment!: pulumi.Output<string>;
    /**
     * The name of the Environment Variable.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * The ID of the Vercel project.
     */
    public readonly projectIds!: pulumi.Output<string[]>;
    /**
     * Whether the Environment Variable is sensitive or not. (May be affected by a [team-wide environment variable policy](https://vercel.com/docs/projects/environment-variables/sensitive-environment-variables#environment-variables-policy))
     */
    public readonly sensitive!: pulumi.Output<boolean>;
    /**
     * The environments that the Environment Variable should be present on. Valid targets are either `production`, `preview`, or `development`.
     */
    public readonly targets!: pulumi.Output<string[]>;
    /**
     * The ID of the Vercel team. Shared environment variables require a team.
     */
    public readonly teamId!: pulumi.Output<string>;
    /**
     * The value of the Environment Variable.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a SharedEnvironmentVariable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SharedEnvironmentVariableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SharedEnvironmentVariableArgs | SharedEnvironmentVariableState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SharedEnvironmentVariableState | undefined;
            resourceInputs["applyToAllCustomEnvironments"] = state ? state.applyToAllCustomEnvironments : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["projectIds"] = state ? state.projectIds : undefined;
            resourceInputs["sensitive"] = state ? state.sensitive : undefined;
            resourceInputs["targets"] = state ? state.targets : undefined;
            resourceInputs["teamId"] = state ? state.teamId : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as SharedEnvironmentVariableArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.projectIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectIds'");
            }
            if ((!args || args.targets === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targets'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            resourceInputs["applyToAllCustomEnvironments"] = args ? args.applyToAllCustomEnvironments : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["projectIds"] = args ? args.projectIds : undefined;
            resourceInputs["sensitive"] = args ? args.sensitive : undefined;
            resourceInputs["targets"] = args ? args.targets : undefined;
            resourceInputs["teamId"] = args ? args.teamId : undefined;
            resourceInputs["value"] = args?.value ? pulumi.secret(args.value) : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["value"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SharedEnvironmentVariable.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SharedEnvironmentVariable resources.
 */
export interface SharedEnvironmentVariableState {
    /**
     * Whether the shared environment variable should be applied to all custom environments in the linked projects.
     */
    applyToAllCustomEnvironments?: pulumi.Input<boolean>;
    /**
     * A comment explaining what the environment variable is for.
     */
    comment?: pulumi.Input<string>;
    /**
     * The name of the Environment Variable.
     */
    key?: pulumi.Input<string>;
    /**
     * The ID of the Vercel project.
     */
    projectIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether the Environment Variable is sensitive or not. (May be affected by a [team-wide environment variable policy](https://vercel.com/docs/projects/environment-variables/sensitive-environment-variables#environment-variables-policy))
     */
    sensitive?: pulumi.Input<boolean>;
    /**
     * The environments that the Environment Variable should be present on. Valid targets are either `production`, `preview`, or `development`.
     */
    targets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the Vercel team. Shared environment variables require a team.
     */
    teamId?: pulumi.Input<string>;
    /**
     * The value of the Environment Variable.
     */
    value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SharedEnvironmentVariable resource.
 */
export interface SharedEnvironmentVariableArgs {
    /**
     * Whether the shared environment variable should be applied to all custom environments in the linked projects.
     */
    applyToAllCustomEnvironments?: pulumi.Input<boolean>;
    /**
     * A comment explaining what the environment variable is for.
     */
    comment?: pulumi.Input<string>;
    /**
     * The name of the Environment Variable.
     */
    key: pulumi.Input<string>;
    /**
     * The ID of the Vercel project.
     */
    projectIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether the Environment Variable is sensitive or not. (May be affected by a [team-wide environment variable policy](https://vercel.com/docs/projects/environment-variables/sensitive-environment-variables#environment-variables-policy))
     */
    sensitive?: pulumi.Input<boolean>;
    /**
     * The environments that the Environment Variable should be present on. Valid targets are either `production`, `preview`, or `development`.
     */
    targets: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the Vercel team. Shared environment variables require a team.
     */
    teamId?: pulumi.Input<string>;
    /**
     * The value of the Environment Variable.
     */
    value: pulumi.Input<string>;
}
