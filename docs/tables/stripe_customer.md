# Table: stripe_customer

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| default_source | json | X | √ | ID of the default payment source for the customer. | 
| phone | string | X | √ | The customer’s phone number. | 
| created | timestamp | X | √ | Time at which the object was created. | 
| delinquent | bool | X | √ | When the customer’s latest invoice is billed by charging automatically, delinquent is true if the invoice’s latest charge failed. When the customer’s latest invoice is billed by sending an invoice, delinquent is true if the invoice isn’t paid by its due date. If an invoice is marked uncollectible by dunning, delinquent doesn’t get reset to false. | 
| discount | json | X | √ | Describes the current discount active on the customer, if there is one. | 
| invoice_settings | json | X | √ | The customer’s default invoice settings. | 
| livemode | bool | X | √ | Has the value true if the object exists in live mode or the value false if the object exists in test mode. | 
| description | string | X | √ | An arbitrary string attached to the object. Often useful for displaying to users. | 
| invoice_prefix | string | X | √ | The prefix for the customer used to generate unique invoice numbers. | 
| next_invoice_sequence | int | X | √ | The suffix of the customer’s next invoice number, e.g., 0001. | 
| id | string | X | √ | Unique identifier for the customer. | 
| email | string | X | √ | The customer’s email address. | 
| balance | int | X | √ | Current balance, if any, being stored on the customer. If negative, the customer has credit to apply to their next invoice. If positive, the customer has an amount owed that will be added to their next invoice. The balance does not refer to any unpaid invoices; it solely takes into account amounts that have yet to be successfully applied to any invoice. This balance is only taken into account as invoices are finalized. | 
| currency | string | X | √ | Three-letter ISO code for the currency the customer can be charged in for recurring billing purposes. | 
| deleted | bool | X | √ | True if the customer is marked as deleted. | 
| shipping | json | X | √ | Mailing and shipping address for the customer. Appears on invoices emailed to this customer. | 
| tax_exempt | string | X | √ | Describes the customer’s tax exemption status. One of none, exempt, or reverse. | 
| tax_ids | json | X | √ | The customer’s tax IDs. | 
| address | json | X | √ | The customer’s address. | 
| metadata | json | X | √ | Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format. | 
| name | string | X | √ | The customer’s full name or business name. | 
| preferred_locales | json | X | √ | The customer’s preferred locales (languages), ordered by preference. | 


