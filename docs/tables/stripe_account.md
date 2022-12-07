# Table: stripe_account

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| individual | json | X | √ | Information about the person represented by the account. This field is null unless business_type is set to individual. | 
| type | string | X | √ | The Stripe account type. Can be standard, express, or custom. | 
| id | string | X | √ | Unique identifier for the account. | 
| email | string | X | √ | An email address associated with the account. You can treat this as metadata: it is not used for authentication or messaging account holders. | 
| charges_enabled | bool | X | √ | Whether the account can create live charges. | 
| created | timestamp | X | √ | Time at which the account was created. | 
| deleted | bool | X | √ | True if the customer is marked as deleted. | 
| business_profile | json | X | √ | Business information about the account. | 
| company | json | X | √ | Information about the company or business. This field is available for any business_type. | 
| default_currency | string | X | √ | Three-letter ISO currency code representing the default currency for the account. | 
| metadata | json | X | √ | Set of key-value pairs that you can attach to an account. This can be useful for storing additional information about the account in a structured format. | 
| settings | json | X | √ | Options for customizing how the account functions within Stripe. | 
| requirements | json | X | √ | Information about the requirements for the account, including what information needs to be collected, and by when. | 
| business_type | string | X | √ | The business type. | 
| capabilities | json | X | √ | A hash containing the set of capabilities that was requested for this account and their associated states. Keys are names of capabilities. You can see the full list here. Values may be active, inactive, or pending. | 
| country | string | X | √ | The account’s country. | 
| external_accounts | json | X | √ | External accounts (bank accounts and debit cards) currently attached to this account. | 
| payouts_enabled | bool | X | √ | Whether Stripe can send payouts to this account. | 
| details_submitted | bool | X | √ | Whether account details have been submitted. Standard accounts cannot receive payouts before this is true. | 
| tos_acceptance | json | X | √ | Details on the acceptance of the Stripe Services Agreement. | 


