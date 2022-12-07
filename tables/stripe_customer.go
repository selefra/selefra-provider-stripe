package tables

import (
	"context"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-stripe/stripe_client"
	"github.com/stripe/stripe-go"
)

type TableStripeCustomerGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableStripeCustomerGenerator{}

func (x *TableStripeCustomerGenerator) GetTableName() string {
	return "stripe_customer"
}

func (x *TableStripeCustomerGenerator) GetTableDescription() string {
	return ""
}

func (x *TableStripeCustomerGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableStripeCustomerGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableStripeCustomerGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := stripe_client.Connect(ctx, taskClient.(*stripe_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			params := &stripe.CustomerListParams{
				ListParams: stripe.ListParams{
					Context: ctx,
					Limit:   stripe.Int64(100),
				},
			}

			var count int64
			i := conn.Customers.List(params)
			for i.Next() {
				resultChannel <- i.Customer()
				count++

			}
			if err := i.Err(); err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableStripeCustomerGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableStripeCustomerGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("default_source").ColumnType(schema.ColumnTypeJSON).Description("ID of the default payment source for the customer.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("phone").ColumnType(schema.ColumnTypeString).Description("The customer’s phone number.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("Time at which the object was created.").
			Extractor(column_value_extractor.StructSelector("Created")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delinquent").ColumnType(schema.ColumnTypeBool).Description("When the customer’s latest invoice is billed by charging automatically, delinquent is true if the invoice’s latest charge failed. When the customer’s latest invoice is billed by sending an invoice, delinquent is true if the invoice isn’t paid by its due date. If an invoice is marked uncollectible by dunning, delinquent doesn’t get reset to false.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("discount").ColumnType(schema.ColumnTypeJSON).Description("Describes the current discount active on the customer, if there is one.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("invoice_settings").ColumnType(schema.ColumnTypeJSON).Description("The customer’s default invoice settings.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("livemode").ColumnType(schema.ColumnTypeBool).Description("Has the value true if the object exists in live mode or the value false if the object exists in test mode.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("An arbitrary string attached to the object. Often useful for displaying to users.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("invoice_prefix").ColumnType(schema.ColumnTypeString).Description("The prefix for the customer used to generate unique invoice numbers.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("next_invoice_sequence").ColumnType(schema.ColumnTypeInt).Description("The suffix of the customer’s next invoice number, e.g., 0001.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique identifier for the customer.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email").ColumnType(schema.ColumnTypeString).Description("The customer’s email address.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("balance").ColumnType(schema.ColumnTypeInt).Description("Current balance, if any, being stored on the customer. If negative, the customer has credit to apply to their next invoice. If positive, the customer has an amount owed that will be added to their next invoice. The balance does not refer to any unpaid invoices; it solely takes into account amounts that have yet to be successfully applied to any invoice. This balance is only taken into account as invoices are finalized.").
			Extractor(column_value_extractor.StructSelector("Balance")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("currency").ColumnType(schema.ColumnTypeString).Description("Three-letter ISO code for the currency the customer can be charged in for recurring billing purposes.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deleted").ColumnType(schema.ColumnTypeBool).Description("True if the customer is marked as deleted.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shipping").ColumnType(schema.ColumnTypeJSON).Description("Mailing and shipping address for the customer. Appears on invoices emailed to this customer.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tax_exempt").ColumnType(schema.ColumnTypeString).Description("Describes the customer’s tax exemption status. One of none, exempt, or reverse.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tax_ids").ColumnType(schema.ColumnTypeJSON).Description("The customer’s tax IDs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("address").ColumnType(schema.ColumnTypeJSON).Description("The customer’s address.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeJSON).Description("Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The customer’s full name or business name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("preferred_locales").ColumnType(schema.ColumnTypeJSON).Description("The customer’s preferred locales (languages), ordered by preference.").Build(),
	}
}

func (x *TableStripeCustomerGenerator) GetSubTables() []*schema.Table {
	return nil
}
