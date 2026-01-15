// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vercel

import (
	"context"
	"fmt"
	"path/filepath"

	_ "embed" // nolint: golint

	"github.com/vercel/terraform-provider-vercel/v4/vercel"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	pfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"

	"github.com/pulumiverse/pulumi-vercel/provider/v4/pkg/version"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "vercel"
	// modules:
	mainMod   = "index" // the vercel module
	publisher = "pulumiverse"
)

//go:embed cmd/pulumi-resource-vercel/bridge-metadata.json
var bridgeMetadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := pfbridge.ShimProvider(vercel.New())

	caser := cases.Title(language.English)

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                       p,
		Name:                    mainPkg,
		DisplayName:             caser.String(mainPkg),
		Publisher:               caser.String(publisher),
		Version:                 version.Version,
		PluginDownloadURL:       "github://api.github.com/pulumiverse",
		Keywords:                []string{"pulumi", "vercel", "category/cloud"},
		License:                 "Apache-2.0",
		GitHubOrg:               "vercel",
		TFProviderModuleVersion: "v4",
		Repository:              "https://github.com/pulumiverse/pulumi-vercel",
		Resources: map[string]*tfbridge.ResourceInfo{
			"vercel_alias": {
				Fields: map[string]*tfbridge.SchemaInfo{
					"alias": {
						CSharpName: "DeploymentAlias",
					},
				},
			},
			"vercel_edge_config_schema": {
				ComputeID: func(_ context.Context, state resource.PropertyMap) (resource.ID, error) {
					return resource.ID(state["id"].StringValue()), nil
				},
			},
			"vercel_team_config": {
				ComputeID: func(_ context.Context, state resource.PropertyMap) (resource.ID, error) {
					return resource.ID(state["id"].StringValue()), nil
				},
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: fmt.Sprintf("@%s/%s", publisher, mainPkg),

			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			PackageName:          fmt.Sprintf("%s_%s", publisher, mainPkg),
			PyProject:            struct{ Enabled bool }{true},
			RespectSchemaVersion: true,
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/%s/pulumi-%s/sdk/", publisher, mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
			RespectSchemaVersion:           true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: caser.String(publisher),
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			RespectSchemaVersion: true,
		},
		Java: &tfbridge.JavaInfo{
			BasePackage: "com.pulumiverse",
		},
		MetadataInfo: tfbridge.NewProviderMetadata(bridgeMetadata),
	}

	// These are new API's that you may opt to use to automatically compute resource tokens,
	// and apply auto aliasing for full backwards compatibility.
	// For more information, please reference
	// https://pkg.go.dev/github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge#ProviderInfo.ComputeTokens
	prov.MustComputeTokens(tokens.SingleModule("vercel_", mainMod,
		tokens.MakeStandard(mainPkg)))
	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
