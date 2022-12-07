package tables

import (
	"context"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-stripe/stripe_client"
	"github.com/stripe/stripe-go"
)

type TableStripePlanGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableStripePlanGenerator{}

func (x *TableStripePlanGenerator) GetTableName() string {
	return "stripe_plan"
}

func (x *TableStripePlanGenerator) GetTableDescription() string {
	return ""
}

func (x *TableStripePlanGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableStripePlanGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableStripePlanGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := stripe_client.Connect(ctx, taskClient.(*stripe_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			params := &stripe.PlanListParams{
				ListParams: stripe.ListParams{
					Context: ctx,
					Limit:   stripe.Int64(100),
				},
			}

			var count int64
			i := conn.Plans.List(params)
			for i.Next() {
				resultChannel <- i.Plan()
				count++

			}
			if err := i.Err(); err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableStripePlanGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableStripePlanGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("billing_scheme").ColumnType(schema.ColumnTypeString).Description("Describes how to compute the price per period. Either per_unit or tiered.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("Time at which the plan was created.").
			Extractor(column_value_extractor.StructSelector("Created")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("currency").ColumnType(schema.ColumnTypeString).Description("Three-letter ISO currency code, in lowercase. Must be a supported currency.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("interval_count").ColumnType(schema.ColumnTypeInt).Description("The number of intervals (specified in the interval attribute) between subscription billings. For example, interval=month and interval_count=3 bills every 3 months.").
			Extractor(column_value_extractor.StructSelector("IntervalCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("product_id").ColumnType(schema.ColumnTypeString).Description("ID of the product whose pricing this plan determines.").
			Extractor(column_value_extractor.StructSelector("Product.ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("nickname").ColumnType(schema.ColumnTypeString).Description("A brief description of the plan, hidden from customers.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("amount").ColumnType(schema.ColumnTypeInt).Description("The unit amount in cents to be charged, represented as a whole integer if possible. Only set if billing_scheme=per_unit.").
			Extractor(column_value_extractor.StructSelector("Amount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeJSON).Description("Set of key-value pairs that you can attach to an plan. This can be useful for storing additional information about the plan in a structured format.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("interval").ColumnType(schema.ColumnTypeString).Description("The frequency at which a subscription is billed. One of day, week, month or year.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("livemode").ColumnType(schema.ColumnTypeBool).Description("Has the value true if the plan exists in live mode or the value false if the plan exists in test mode.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aggregate_usage").ColumnType(schema.ColumnTypeString).Description("Specifies a usage aggregation strategy for plans of usage_type=metered. Allowed values are sum for summing up all usage during a period, last_during_period for using the last usage record reported within a period, last_ever for using the last usage record ever (across period bounds) or max which uses the usage record with the maximum reported usage during a period. Defaults to sum.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("amount_decimal").ColumnType(schema.ColumnTypeFloat).Description("The unit amount in cents to be charged, represented as a decimal string with at most 12 decimal places. Only set if billing_scheme=per_unit.").
			Extractor(column_value_extractor.StructSelector("AmountDecimal")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tiers").ColumnType(schema.ColumnTypeJSON).Description("Each element represents a pricing tier. This parameter requires billing_scheme to be set to tiered.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tiers_mode").ColumnType(schema.ColumnTypeString).Description("Defines if the tiering price should be graduated or volume based. In volume-based tiering, the maximum quantity within a period determines the per unit price. In graduated tiering, pricing can change as the quantity grows.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("transform_usage").ColumnType(schema.ColumnTypeJSON).Description("Apply a transformation to the reported usage or set quantity before computing the amount billed.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique identifier for the plan.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("active").ColumnType(schema.ColumnTypeBool).Description("Whether the plan is currently available for purchase.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("usage_type").ColumnType(schema.ColumnTypeString).Description("Configures how the quantity per period should be determined. Can be either metered or licensed. licensed automatically bills the quantity set when adding it to a subscription. metered aggregates the total usage based on usage records. Defaults to licensed.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deleted").ColumnType(schema.ColumnTypeBool).Description("True if the plan is marked as deleted.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("trial_period_days").ColumnType(schema.ColumnTypeInt).Description("Default number of trial days when subscribing a customer to this plan using trial_from_plan=true.").
			Extractor(column_value_extractor.StructSelector("TrialPeriodDays")).Build(),
	}
}

func (x *TableStripePlanGenerator) GetSubTables() []*schema.Table {
	return nil
}
