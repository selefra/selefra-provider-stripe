package tables

import (
	"context"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-stripe/stripe_client"
	"github.com/stripe/stripe-go"
)

type TableStripeProductGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableStripeProductGenerator{}

func (x *TableStripeProductGenerator) GetTableName() string {
	return "stripe_product"
}

func (x *TableStripeProductGenerator) GetTableDescription() string {
	return ""
}

func (x *TableStripeProductGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableStripeProductGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableStripeProductGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := stripe_client.Connect(ctx, taskClient.(*stripe_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			params := &stripe.ProductListParams{
				ListParams: stripe.ListParams{
					Context: ctx,
					Limit:   stripe.Int64(100),
				},
			}

			var count int64
			i := conn.Products.List(params)
			for i.Next() {
				resultChannel <- i.Product()
				count++

			}
			if err := i.Err(); err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableStripeProductGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableStripeProductGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("An arbitrary string attached to the product. Often useful for displaying to users.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("unit_label").ColumnType(schema.ColumnTypeString).Description("A label that represents units of this product in Stripe and on customers’ receipts and invoices. When set, this will be included in associated invoice line item descriptions.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("livemode").ColumnType(schema.ColumnTypeBool).Description("Has the value true if the product exists in live mode or the value false if the product exists in test mode.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeJSON).Description("Set of key-value pairs that you can attach to an product. This can be useful for storing additional information about the product in a structured format.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("statement_descriptor").ColumnType(schema.ColumnTypeString).Description("Extra information about a product which will appear on your customer’s credit card statement. In the case that multiple products are billed at once, the first statement descriptor will be used.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("url").ColumnType(schema.ColumnTypeString).Description("A URL of a publicly-accessible webpage for this product.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("active").ColumnType(schema.ColumnTypeBool).Description("Whether the product is currently available for purchase.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("Time at which the product was created.").
			Extractor(column_value_extractor.StructSelector("Created")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("package_dimensions").ColumnType(schema.ColumnTypeJSON).Description("The dimensions of this product for shipping purposes.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique identifier for the product.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The product’s full name or business name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("The product type.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("images").ColumnType(schema.ColumnTypeJSON).Description("A list of up to 8 URLs of images for this product, meant to be displayable to the customer.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shippable").ColumnType(schema.ColumnTypeBool).Description("Whether this product is shipped (i.e., physical goods).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated").ColumnType(schema.ColumnTypeTimestamp).Description("Time at which the product was updated.").
			Extractor(column_value_extractor.StructSelector("Updated")).Build(),
	}
}

func (x *TableStripeProductGenerator) GetSubTables() []*schema.Table {
	return nil
}
