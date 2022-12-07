package tables

import (
	"context"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-stripe/stripe_client"
	"github.com/stripe/stripe-go"
)

type TableStripeInvoiceGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableStripeInvoiceGenerator{}

func (x *TableStripeInvoiceGenerator) GetTableName() string {
	return "stripe_invoice"
}

func (x *TableStripeInvoiceGenerator) GetTableDescription() string {
	return ""
}

func (x *TableStripeInvoiceGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableStripeInvoiceGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableStripeInvoiceGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := stripe_client.Connect(ctx, taskClient.(*stripe_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			params := &stripe.InvoiceListParams{
				ListParams: stripe.ListParams{
					Context: ctx,
					Limit:   stripe.Int64(100),
					Expand:  stripe.StringSlice([]string{"data.default_payment_method", "data.default_source", "data.subscription"}),
				},
			}

			var count int64
			i := conn.Invoices.List(params)
			for i.Next() {
				resultChannel <- i.Invoice()
				count++

			}
			if err := i.Err(); err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableStripeInvoiceGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableStripeInvoiceGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).Description("ID of the subscription that this invoice was prepared for, if any.").
			Extractor(column_value_extractor.StructSelector("Subscription.ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_country").ColumnType(schema.ColumnTypeString).Description("The country of the business associated with this invoice, most often the business creating the invoice.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_advance").ColumnType(schema.ColumnTypeBool).Description("Controls whether Stripe will perform automatic collection of the invoice. When false, the invoice’s state will not automatically advance without an explicit action.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeJSON).Description("Set of key-value pairs that you can attach to an invoice. This can be useful for storing additional information about the invoice in a structured format.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("paid").ColumnType(schema.ColumnTypeBool).Description("Whether payment was successfully collected for this invoice. An invoice can be paid (most commonly) with a charge or with credit from the customer’s account balance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("period_start").ColumnType(schema.ColumnTypeTimestamp).Description("Start of the usage period during which invoice items were added to this invoice.").
			Extractor(column_value_extractor.StructSelector("PeriodStart")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("amount_paid").ColumnType(schema.ColumnTypeInt).Description("The amount, in cents, that was paid.").
			Extractor(column_value_extractor.StructSelector("AmountPaid")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("customer_tax_exempt").ColumnType(schema.ColumnTypeString).Description("The customer’s tax exempt status. Until the invoice is finalized, this field will equal customer.tax_exempt. Once the invoice is finalized, this field will no longer be updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_tax_rates").ColumnType(schema.ColumnTypeJSON).Description("The tax rates applied to this invoice, if any.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("due_date").ColumnType(schema.ColumnTypeTimestamp).Description("The date on which payment for this invoice is due. This value will be null for invoices where collection_method=charge_automatically.").
			Extractor(column_value_extractor.StructSelector("DueDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_transitions").ColumnType(schema.ColumnTypeJSON).Description("The timestamps at which the invoice status was updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("currency").ColumnType(schema.ColumnTypeString).Description("Three-letter ISO currency code, in lowercase. Must be a supported currency.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lines").ColumnType(schema.ColumnTypeJSON).Description("The individual line items that make up the invoice. lines is sorted as follows: invoice items in reverse chronological order, followed by the subscription, if any.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("transfer_data").ColumnType(schema.ColumnTypeJSON).Description("The account (if any) the payment will be attributed to for tax reporting, and where funds from the payment will be transferred to for the invoice.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pre_payment_credit_notes_amount").ColumnType(schema.ColumnTypeInt).Description("Total amount of all pre-payment credit notes issued for this invoice.").
			Extractor(column_value_extractor.StructSelector("PrePaymentCreditNotesAmount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("starting_balance").ColumnType(schema.ColumnTypeInt).Description("Starting customer balance before the invoice is finalized. If the invoice has not been finalized yet, this will be the current customer balance.").
			Extractor(column_value_extractor.StructSelector("StartingBalance")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("customer").ColumnType(schema.ColumnTypeJSON).Description("The ID of the customer who will be billed.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("customer_phone").ColumnType(schema.ColumnTypeString).Description("The customer’s phone number. Until the invoice is finalized, this field will equal customer.phone. Once the invoice is finalized, this field will no longer be updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("livemode").ColumnType(schema.ColumnTypeBool).Description("Has the value true if the invoice exists in live mode or the value false if the invoice exists in test mode.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("collection_method").ColumnType(schema.ColumnTypeString).Description("Either charge_automatically, or send_invoice. When charging automatically, Stripe will attempt to pay this invoice using the default source attached to the customer. When sending an invoice, Stripe will email this invoice to the customer with payment instructions.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_source").ColumnType(schema.ColumnTypeString).Description("ID of the default payment source for the invoice. It must belong to the customer associated with the invoice and be in a chargeable state. If not set, defaults to the subscription’s default source, if any, or to the customer’s default source.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ending_balance").ColumnType(schema.ColumnTypeInt).Description("Ending customer balance after the invoice is finalized. Invoices are finalized approximately an hour after successful webhook delivery or when payment collection is attempted for the invoice. If the invoice has not been finalized yet, this will be null.").
			Extractor(column_value_extractor.StructSelector("EndingBalance")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_proration_date").ColumnType(schema.ColumnTypeTimestamp).Description("Only set for upcoming invoices that preview prorations. The time used to calculate prorations.").
			Extractor(column_value_extractor.StructSelector("SubscriptionProrationDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("amount_remaining").ColumnType(schema.ColumnTypeInt).Description("The amount remaining, in cents, that is due.").
			Extractor(column_value_extractor.StructSelector("AmountRemaining")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attempt_count").ColumnType(schema.ColumnTypeInt).Description("Number of payment attempts made for this invoice, from the perspective of the payment retry schedule. Any payment attempt counts as the first attempt, and subsequently only automatic retries increment the attempt count. In other words, manual payment attempts after the first attempt do not affect the retry schedule.").
			Extractor(column_value_extractor.StructSelector("AttemptCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("customer_tax_ids").ColumnType(schema.ColumnTypeJSON).Description("The customer’s tax IDs. Until the invoice is finalized, this field will contain the same tax IDs as customer.tax_ids. Once the invoice is finalized, this field will no longer be updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_payment_method").ColumnType(schema.ColumnTypeString).Description("ID of the default payment method for the invoice. It must belong to the customer associated with the invoice. If not set, defaults to the subscription’s default payment method, if any, or to the default payment method in the customer’s invoice settings.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("threshold_reason").ColumnType(schema.ColumnTypeJSON).Description("If billing_reason is set to subscription_threshold this returns more information on which threshold rules triggered the invoice.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("charge").ColumnType(schema.ColumnTypeJSON).Description("ID of the latest charge generated for this invoice, if any.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("customer_email").ColumnType(schema.ColumnTypeString).Description("The customer’s email. Until the invoice is finalized, this field will equal customer.email. Once the invoice is finalized, this field will no longer be updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("customer_name").ColumnType(schema.ColumnTypeString).Description("The customer’s name. Until the invoice is finalized, this field will equal customer.name. Once the invoice is finalized, this field will no longer be updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("next_payment_attempt").ColumnType(schema.ColumnTypeTimestamp).Description("The time at which payment will next be attempted. This value will be null for invoices where collection_method=send_invoice.").
			Extractor(column_value_extractor.StructSelector("NextPaymentAttempt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_name").ColumnType(schema.ColumnTypeString).Description("The public name of the business associated with this invoice, most often the business creating the invoice.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("customer_address").ColumnType(schema.ColumnTypeJSON).Description("The customer’s address. Until the invoice is finalized, this field will equal customer.address. Once the invoice is finalized, this field will no longer be updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("invoice_pdf").ColumnType(schema.ColumnTypeString).Description("The link to download the PDF for the invoice. If the invoice has not been finalized yet, this will be null.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("statement_descriptor").ColumnType(schema.ColumnTypeString).Description("Extra information about an invoice for the customer’s credit card statement.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("total").ColumnType(schema.ColumnTypeInt).Description("Total after discounts and taxes.").
			Extractor(column_value_extractor.StructSelector("Total")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("webhooks_delivered_at").ColumnType(schema.ColumnTypeTimestamp).Description("Invoices are automatically paid or sent 1 hour after webhooks are delivered, or until all webhook delivery attempts have been exhausted. This field tracks the time when webhooks for this invoice were successfully delivered. If the invoice had no webhooks to deliver, this will be set while the invoice is being created.").
			Extractor(column_value_extractor.StructSelector("WebhooksDeliveredAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("An arbitrary string attached to the object. Often useful for displaying to users. Referenced as ‘memo’ in the Dashboard.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("footer").ColumnType(schema.ColumnTypeString).Description("Footer displayed on the invoice.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hosted_invoice_url").ColumnType(schema.ColumnTypeString).Description("The URL for the hosted invoice page, which allows customers to view and pay an invoice. If the invoice has not been finalized yet, this will be null.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subtotal").ColumnType(schema.ColumnTypeInt).Description("Total of all subscriptions, invoice items, and prorations on the invoice before any invoice level discount or tax is applied. Item discounts are already incorporated").
			Extractor(column_value_extractor.StructSelector("Subtotal")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The status of the invoice, one of draft, open, paid, uncollectible, or void.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("application_fee_amount").ColumnType(schema.ColumnTypeInt).Description("The fee in cents that will be applied to the invoice and transferred to the application owner’s Stripe account when the invoice is paid.").
			Extractor(column_value_extractor.StructSelector("ApplicationFeeAmount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("discount").ColumnType(schema.ColumnTypeJSON).Description("Describes the current discount applied to this invoice, if there is one. Not populated if there are multiple discounts.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("payment_intent").ColumnType(schema.ColumnTypeJSON).Description("The PaymentIntent associated with this invoice. The PaymentIntent is generated when the invoice is finalized, and can then be used to pay the invoice. Note that voiding an invoice will cancel the PaymentIntent.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("total_tax_amounts").ColumnType(schema.ColumnTypeJSON).Description("The aggregate amounts calculated per tax rate for all line items.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attempted").ColumnType(schema.ColumnTypeBool).Description("Whether an attempt has been made to pay the invoice. An invoice is not attempted until 1 hour after the invoice.created webhook, for example, so you might not want to display that invoice as unpaid to your users.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("custom_fields").ColumnType(schema.ColumnTypeJSON).Description("Custom fields displayed on the invoice.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("customer_shipping").ColumnType(schema.ColumnTypeJSON).Description("The customer’s shipping information. Until the invoice is finalized, this field will equal customer.shipping. Once the invoice is finalized, this field will no longer be updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("receipt_number").ColumnType(schema.ColumnTypeString).Description("This is the transaction number that appears on email receipts sent for this invoice.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tax").ColumnType(schema.ColumnTypeInt).Description("The amount of tax on this invoice. This is the sum of all the tax amounts on this invoice.").
			Extractor(column_value_extractor.StructSelector("Tax")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("number").ColumnType(schema.ColumnTypeString).Description("A unique, identifying string that appears on emails sent to the customer for this invoice. This starts with the customer’s unique invoice_prefix if it is specified.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("amount_due").ColumnType(schema.ColumnTypeInt).Description("Final amount due at this time for this invoice. If the invoice’s total is smaller than the minimum charge amount, for example, or if there is account credit that can be applied to the invoice, the amount_due may be 0. If there is a positive starting_balance for the invoice (the customer owes money), the amount_due will also take that into account. The charge that gets generated for the invoice will be for the amount specified in amount_due.").
			Extractor(column_value_extractor.StructSelector("AmountDue")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("Time at which the invoice was created.").
			Extractor(column_value_extractor.StructSelector("Created")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("post_payment_credit_notes_amount").ColumnType(schema.ColumnTypeInt).Description("Total amount of all post-payment credit notes issued for this invoice.").
			Extractor(column_value_extractor.StructSelector("PostPaymentCreditNotesAmount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique identifier for the invoice.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("billing_reason").ColumnType(schema.ColumnTypeString).Description("Indicates the reason why the invoice was created. subscription_cycle indicates an invoice created by a subscription advancing into a new period. subscription_create indicates an invoice created due to creating a subscription. subscription_update indicates an invoice created due to updating a subscription. subscription is set for all old invoices to indicate either a change to a subscription or a period advancement. manual is set for all invoices unrelated to a subscription (for example: created via the invoice editor). The upcoming value is reserved for simulated invoices per the upcoming invoice endpoint. subscription_threshold indicates an invoice created due to a billing threshold being reached.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("period_end").ColumnType(schema.ColumnTypeTimestamp).Description("End of the usage period during which invoice items were added to this invoice.").
			Extractor(column_value_extractor.StructSelector("PeriodEnd")).Build(),
	}
}

func (x *TableStripeInvoiceGenerator) GetSubTables() []*schema.Table {
	return nil
}
