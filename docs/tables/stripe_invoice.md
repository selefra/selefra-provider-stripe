# Table: stripe_invoice

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| subscription_id | string | X | √ | ID of the subscription that this invoice was prepared for, if any. | 
| account_country | string | X | √ | The country of the business associated with this invoice, most often the business creating the invoice. | 
| auto_advance | bool | X | √ | Controls whether Stripe will perform automatic collection of the invoice. When false, the invoice’s state will not automatically advance without an explicit action. | 
| metadata | json | X | √ | Set of key-value pairs that you can attach to an invoice. This can be useful for storing additional information about the invoice in a structured format. | 
| paid | bool | X | √ | Whether payment was successfully collected for this invoice. An invoice can be paid (most commonly) with a charge or with credit from the customer’s account balance. | 
| period_start | timestamp | X | √ | Start of the usage period during which invoice items were added to this invoice. | 
| amount_paid | int | X | √ | The amount, in cents, that was paid. | 
| customer_tax_exempt | string | X | √ | The customer’s tax exempt status. Until the invoice is finalized, this field will equal customer.tax_exempt. Once the invoice is finalized, this field will no longer be updated. | 
| default_tax_rates | json | X | √ | The tax rates applied to this invoice, if any. | 
| due_date | timestamp | X | √ | The date on which payment for this invoice is due. This value will be null for invoices where collection_method=charge_automatically. | 
| status_transitions | json | X | √ | The timestamps at which the invoice status was updated. | 
| currency | string | X | √ | Three-letter ISO currency code, in lowercase. Must be a supported currency. | 
| lines | json | X | √ | The individual line items that make up the invoice. lines is sorted as follows: invoice items in reverse chronological order, followed by the subscription, if any. | 
| transfer_data | json | X | √ | The account (if any) the payment will be attributed to for tax reporting, and where funds from the payment will be transferred to for the invoice. | 
| pre_payment_credit_notes_amount | int | X | √ | Total amount of all pre-payment credit notes issued for this invoice. | 
| starting_balance | int | X | √ | Starting customer balance before the invoice is finalized. If the invoice has not been finalized yet, this will be the current customer balance. | 
| customer | json | X | √ | The ID of the customer who will be billed. | 
| customer_phone | string | X | √ | The customer’s phone number. Until the invoice is finalized, this field will equal customer.phone. Once the invoice is finalized, this field will no longer be updated. | 
| livemode | bool | X | √ | Has the value true if the invoice exists in live mode or the value false if the invoice exists in test mode. | 
| collection_method | string | X | √ | Either charge_automatically, or send_invoice. When charging automatically, Stripe will attempt to pay this invoice using the default source attached to the customer. When sending an invoice, Stripe will email this invoice to the customer with payment instructions. | 
| default_source | string | X | √ | ID of the default payment source for the invoice. It must belong to the customer associated with the invoice and be in a chargeable state. If not set, defaults to the subscription’s default source, if any, or to the customer’s default source. | 
| ending_balance | int | X | √ | Ending customer balance after the invoice is finalized. Invoices are finalized approximately an hour after successful webhook delivery or when payment collection is attempted for the invoice. If the invoice has not been finalized yet, this will be null. | 
| subscription_proration_date | timestamp | X | √ | Only set for upcoming invoices that preview prorations. The time used to calculate prorations. | 
| amount_remaining | int | X | √ | The amount remaining, in cents, that is due. | 
| attempt_count | int | X | √ | Number of payment attempts made for this invoice, from the perspective of the payment retry schedule. Any payment attempt counts as the first attempt, and subsequently only automatic retries increment the attempt count. In other words, manual payment attempts after the first attempt do not affect the retry schedule. | 
| customer_tax_ids | json | X | √ | The customer’s tax IDs. Until the invoice is finalized, this field will contain the same tax IDs as customer.tax_ids. Once the invoice is finalized, this field will no longer be updated. | 
| default_payment_method | string | X | √ | ID of the default payment method for the invoice. It must belong to the customer associated with the invoice. If not set, defaults to the subscription’s default payment method, if any, or to the default payment method in the customer’s invoice settings. | 
| threshold_reason | json | X | √ | If billing_reason is set to subscription_threshold this returns more information on which threshold rules triggered the invoice. | 
| charge | json | X | √ | ID of the latest charge generated for this invoice, if any. | 
| customer_email | string | X | √ | The customer’s email. Until the invoice is finalized, this field will equal customer.email. Once the invoice is finalized, this field will no longer be updated. | 
| customer_name | string | X | √ | The customer’s name. Until the invoice is finalized, this field will equal customer.name. Once the invoice is finalized, this field will no longer be updated. | 
| next_payment_attempt | timestamp | X | √ | The time at which payment will next be attempted. This value will be null for invoices where collection_method=send_invoice. | 
| account_name | string | X | √ | The public name of the business associated with this invoice, most often the business creating the invoice. | 
| customer_address | json | X | √ | The customer’s address. Until the invoice is finalized, this field will equal customer.address. Once the invoice is finalized, this field will no longer be updated. | 
| invoice_pdf | string | X | √ | The link to download the PDF for the invoice. If the invoice has not been finalized yet, this will be null. | 
| statement_descriptor | string | X | √ | Extra information about an invoice for the customer’s credit card statement. | 
| total | int | X | √ | Total after discounts and taxes. | 
| webhooks_delivered_at | timestamp | X | √ | Invoices are automatically paid or sent 1 hour after webhooks are delivered, or until all webhook delivery attempts have been exhausted. This field tracks the time when webhooks for this invoice were successfully delivered. If the invoice had no webhooks to deliver, this will be set while the invoice is being created. | 
| description | string | X | √ | An arbitrary string attached to the object. Often useful for displaying to users. Referenced as ‘memo’ in the Dashboard. | 
| footer | string | X | √ | Footer displayed on the invoice. | 
| hosted_invoice_url | string | X | √ | The URL for the hosted invoice page, which allows customers to view and pay an invoice. If the invoice has not been finalized yet, this will be null. | 
| subtotal | int | X | √ | Total of all subscriptions, invoice items, and prorations on the invoice before any invoice level discount or tax is applied. Item discounts are already incorporated | 
| status | string | X | √ | The status of the invoice, one of draft, open, paid, uncollectible, or void. | 
| application_fee_amount | int | X | √ | The fee in cents that will be applied to the invoice and transferred to the application owner’s Stripe account when the invoice is paid. | 
| discount | json | X | √ | Describes the current discount applied to this invoice, if there is one. Not populated if there are multiple discounts. | 
| payment_intent | json | X | √ | The PaymentIntent associated with this invoice. The PaymentIntent is generated when the invoice is finalized, and can then be used to pay the invoice. Note that voiding an invoice will cancel the PaymentIntent. | 
| total_tax_amounts | json | X | √ | The aggregate amounts calculated per tax rate for all line items. | 
| attempted | bool | X | √ | Whether an attempt has been made to pay the invoice. An invoice is not attempted until 1 hour after the invoice.created webhook, for example, so you might not want to display that invoice as unpaid to your users. | 
| custom_fields | json | X | √ | Custom fields displayed on the invoice. | 
| customer_shipping | json | X | √ | The customer’s shipping information. Until the invoice is finalized, this field will equal customer.shipping. Once the invoice is finalized, this field will no longer be updated. | 
| receipt_number | string | X | √ | This is the transaction number that appears on email receipts sent for this invoice. | 
| tax | int | X | √ | The amount of tax on this invoice. This is the sum of all the tax amounts on this invoice. | 
| number | string | X | √ | A unique, identifying string that appears on emails sent to the customer for this invoice. This starts with the customer’s unique invoice_prefix if it is specified. | 
| amount_due | int | X | √ | Final amount due at this time for this invoice. If the invoice’s total is smaller than the minimum charge amount, for example, or if there is account credit that can be applied to the invoice, the amount_due may be 0. If there is a positive starting_balance for the invoice (the customer owes money), the amount_due will also take that into account. The charge that gets generated for the invoice will be for the amount specified in amount_due. | 
| created | timestamp | X | √ | Time at which the invoice was created. | 
| post_payment_credit_notes_amount | int | X | √ | Total amount of all post-payment credit notes issued for this invoice. | 
| id | string | X | √ | Unique identifier for the invoice. | 
| billing_reason | string | X | √ | Indicates the reason why the invoice was created. subscription_cycle indicates an invoice created by a subscription advancing into a new period. subscription_create indicates an invoice created due to creating a subscription. subscription_update indicates an invoice created due to updating a subscription. subscription is set for all old invoices to indicate either a change to a subscription or a period advancement. manual is set for all invoices unrelated to a subscription (for example: created via the invoice editor). The upcoming value is reserved for simulated invoices per the upcoming invoice endpoint. subscription_threshold indicates an invoice created due to a billing threshold being reached. | 
| period_end | timestamp | X | √ | End of the usage period during which invoice items were added to this invoice. | 


