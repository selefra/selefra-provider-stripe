package tables

import (
	"context"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-stripe/stripe_client"
	"github.com/stripe/stripe-go"
)

type TableStripeCouponGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableStripeCouponGenerator{}

func (x *TableStripeCouponGenerator) GetTableName() string {
	return "stripe_coupon"
}

func (x *TableStripeCouponGenerator) GetTableDescription() string {
	return ""
}

func (x *TableStripeCouponGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableStripeCouponGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableStripeCouponGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := stripe_client.Connect(ctx, taskClient.(*stripe_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			params := &stripe.CouponListParams{
				ListParams: stripe.ListParams{
					Context: ctx,
					Limit:   stripe.Int64(100),
				},
			}

			var count int64
			i := conn.Coupons.List(params)
			for i.Next() {
				resultChannel <- i.Coupon()
				count++

			}
			if err := i.Err(); err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableStripeCouponGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableStripeCouponGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("amount_off").ColumnType(schema.ColumnTypeInt).Description("Amount (in the currency specified) that will be taken off the subtotal of any invoices for this customer.").
			Extractor(column_value_extractor.StructSelector("AmountOff")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeJSON).Description("Set of key-value pairs that you can attach to an coupon. This can be useful for storing additional information about the coupon in a structured format.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("percent_off").ColumnType(schema.ColumnTypeFloat).Description("Percent that will be taken off the subtotal of any invoices for this customer for the duration of the coupon. For example, a coupon with percent_off of 50 will make a $100 invoice $50 instead.").
			Extractor(column_value_extractor.StructSelector("PercentOff")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("redeem_by").ColumnType(schema.ColumnTypeTimestamp).Description("Date after which the coupon can no longer be redeemed.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("times_redeemed").ColumnType(schema.ColumnTypeInt).Description("Number of times this coupon has been applied to a customer.").
			Extractor(column_value_extractor.StructSelector("TimesRedeemed")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique identifier for the coupon.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The coupon’s full name or business name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("duration_in_months").ColumnType(schema.ColumnTypeInt).Description("If duration is repeating, the number of months the coupon applies. Null if coupon duration is forever or once.").
			Extractor(column_value_extractor.StructSelector("DurationInMonths")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("livemode").ColumnType(schema.ColumnTypeBool).Description("Has the value true if the coupon exists in live mode or the value false if the coupon exists in test mode.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("valid").ColumnType(schema.ColumnTypeBool).Description("Taking account of the above properties, whether this coupon can still be applied to a customer.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("Time at which the coupon was created.").
			Extractor(column_value_extractor.StructSelector("Created")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("currency").ColumnType(schema.ColumnTypeString).Description("If amount_off has been set, the three-letter ISO code for the currency of the amount to take off.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deleted").ColumnType(schema.ColumnTypeBool).Description("True if the customer is marked as deleted.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("duration").ColumnType(schema.ColumnTypeString).Description("One of forever, once, and repeating. Describes how long a customer who applies this coupon will get the discount.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_redemptions").ColumnType(schema.ColumnTypeInt).Description("Maximum number of times this coupon can be redeemed, in total, across all customers, before it is no longer valid.").
			Extractor(column_value_extractor.StructSelector("MaxRedemptions")).Build(),
	}
}

func (x *TableStripeCouponGenerator) GetSubTables() []*schema.Table {
	return nil
}
