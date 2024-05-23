// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides an Attack Challenge Mode resource.
 *
 * Attack Challenge Mode prevent malicious traffic by showing a verification challenge for every visitor.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vercel from "@pulumiverse/vercel";
 *
 * const exampleProject = new vercel.Project("exampleProject", {});
 * const exampleAttackChallengeMode = new vercel.AttackChallengeMode("exampleAttackChallengeMode", {
 *     projectId: exampleProject.id,
 *     enabled: true,
 * });
 * ```
 *
 * ## Import
 *
 * You can import via the team_id and project_id.
 *
 * - team_id can be found in the team `settings` tab in the Vercel UI.
 *
 * - project_id can be found in the project `settings` tab in the Vercel UI.
 *
 * ```sh
 * $ pulumi import vercel:index/attackChallengeMode:AttackChallengeMode example team_xxxxxxxxxxxxxxxxxxxxxxxx/prj_xxxxxxxxxxxxxxxxxxxxxxxxxxxx
 * ```
 */
export class AttackChallengeMode extends pulumi.CustomResource {
    /**
     * Get an existing AttackChallengeMode resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AttackChallengeModeState, opts?: pulumi.CustomResourceOptions): AttackChallengeMode {
        return new AttackChallengeMode(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vercel:index/attackChallengeMode:AttackChallengeMode';

    /**
     * Returns true if the given object is an instance of AttackChallengeMode.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AttackChallengeMode {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AttackChallengeMode.__pulumiType;
    }

    /**
     * Whether Attack Challenge Mode is enabled or not.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * The ID of the Project to adjust the CPU for.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The ID of the team the Project exists under. Required when configuring a team resource if a default team has not been set in the provider.
     */
    public readonly teamId!: pulumi.Output<string>;

    /**
     * Create a AttackChallengeMode resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AttackChallengeModeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AttackChallengeModeArgs | AttackChallengeModeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AttackChallengeModeState | undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["teamId"] = state ? state.teamId : undefined;
        } else {
            const args = argsOrState as AttackChallengeModeArgs | undefined;
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["teamId"] = args ? args.teamId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AttackChallengeMode.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AttackChallengeMode resources.
 */
export interface AttackChallengeModeState {
    /**
     * Whether Attack Challenge Mode is enabled or not.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The ID of the Project to adjust the CPU for.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The ID of the team the Project exists under. Required when configuring a team resource if a default team has not been set in the provider.
     */
    teamId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AttackChallengeMode resource.
 */
export interface AttackChallengeModeArgs {
    /**
     * Whether Attack Challenge Mode is enabled or not.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * The ID of the Project to adjust the CPU for.
     */
    projectId: pulumi.Input<string>;
    /**
     * The ID of the team the Project exists under. Required when configuring a team resource if a default team has not been set in the provider.
     */
    teamId?: pulumi.Input<string>;
}