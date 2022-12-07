package tables

import (
	"context"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-stripe/stripe_client"
)

type TableStripeAccountGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableStripeAccountGenerator{}

func (x *TableStripeAccountGenerator) GetTableName() string {
	return "stripe_account"
}

func (x *TableStripeAccountGenerator) GetTableDescription() string {
	return ""
}

func (x *TableStripeAccountGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableStripeAccountGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableStripeAccountGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := stripe_client.Connect(ctx, taskClient.(*stripe_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			item, err := conn.Account.Get()
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			resultChannel <- item
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableStripeAccountGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableStripeAccountGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("individual").ColumnType(schema.ColumnTypeJSON).Description("Information about the person represented by the account. This field is null unless business_type is set to individual.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("The Stripe account type. Can be standard, express, or custom.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique identifier for the account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email").ColumnType(schema.ColumnTypeString).Description("An email address associated with the account. You can treat this as metadata: it is not used for authentication or messaging account holders.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("charges_enabled").ColumnType(schema.ColumnTypeBool).Description("Whether the account can create live charges.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("Time at which the account was created.").
			Extractor(column_value_extractor.StructSelector("Created")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deleted").ColumnType(schema.ColumnTypeBool).Description("True if the customer is marked as deleted.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("business_profile").ColumnType(schema.ColumnTypeJSON).Description("Business information about the account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("company").ColumnType(schema.ColumnTypeJSON).Description("Information about the company or business. This field is available for any business_type.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_currency").ColumnType(schema.ColumnTypeString).Description("Three-letter ISO currency code representing the default currency for the account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeJSON).Description("Set of key-value pairs that you can attach to an account. This can be useful for storing additional information about the account in a structured format.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("settings").ColumnType(schema.ColumnTypeJSON).Description("Options for customizing how the account functions within Stripe.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("requirements").ColumnType(schema.ColumnTypeJSON).Description("Information about the requirements for the account, including what information needs to be collected, and by when.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("business_type").ColumnType(schema.ColumnTypeString).Description("The business type.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("capabilities").ColumnType(schema.ColumnTypeJSON).Description("A hash containing the set of capabilities that was requested for this account and their associated states. Keys are names of capabilities. You can see the full list here. Values may be active, inactive, or pending.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("country").ColumnType(schema.ColumnTypeString).Description("The account’s country.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("external_accounts").ColumnType(schema.ColumnTypeJSON).Description("External accounts (bank accounts and debit cards) currently attached to this account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("payouts_enabled").ColumnType(schema.ColumnTypeBool).Description("Whether Stripe can send payouts to this account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("details_submitted").ColumnType(schema.ColumnTypeBool).Description("Whether account details have been submitted. Standard accounts cannot receive payouts before this is true.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tos_acceptance").ColumnType(schema.ColumnTypeJSON).Description("Details on the acceptance of the Stripe Services Agreement.").
			Extractor(column_value_extractor.StructSelector("TOSAcceptance")).Build(),
	}
}

func (x *TableStripeAccountGenerator) GetSubTables() []*schema.Table {
	return nil
}
