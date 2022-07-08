package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ tfsdk.DataSourceType = envoyLbDataSourceType{}
var _ tfsdk.DataSource = envoyLbDataSource{}

type envoyLbDataSourceType struct{}

func (t envoyLbDataSourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Envoy LB data source",

		Attributes: map[string]tfsdk.Attribute{
			"configurable_attribute": {
				MarkdownDescription: "Envoy LB configurable attribute",
				Optional:            true,
				Type:                types.StringType,
			},
			"id": {
				MarkdownDescription: "Envoy LB identifier",
				Type:                types.StringType,
				Computed:            true,
			},
		},
	}, nil
}

func (t envoyLbDataSourceType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return envoyLbDataSource{
		provider: provider,
	}, diags
}

type envoyLbDataSourceData struct {
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
	Id                    types.String `tfsdk:"id"`
}

type envoyLbDataSource struct {
	provider provider
}

func (d envoyLbDataSource) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var data envoyLbDataSourceData

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	log.Printf("got here")

	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("got here")

	data.Id = types.String{Value: "example-id"}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}
