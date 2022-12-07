package provider

import (
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-stripe/tables"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&tables.TableStripeCustomerGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableStripeSubscriptionGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableStripeInvoiceGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableStripePlanGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableStripeProductGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableStripeAccountGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableStripeCouponGenerator{}),
	}
}
