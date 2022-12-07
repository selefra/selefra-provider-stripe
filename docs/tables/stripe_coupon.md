# Table: stripe_coupon

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| amount_off | int | X | √ | Amount (in the currency specified) that will be taken off the subtotal of any invoices for this customer. | 
| metadata | json | X | √ | Set of key-value pairs that you can attach to an coupon. This can be useful for storing additional information about the coupon in a structured format. | 
| percent_off | float | X | √ | Percent that will be taken off the subtotal of any invoices for this customer for the duration of the coupon. For example, a coupon with percent_off of 50 will make a $100 invoice $50 instead. | 
| redeem_by | timestamp | X | √ | Date after which the coupon can no longer be redeemed. | 
| times_redeemed | int | X | √ | Number of times this coupon has been applied to a customer. | 
| id | string | X | √ | Unique identifier for the coupon. | 
| name | string | X | √ | The coupon’s full name or business name. | 
| duration_in_months | int | X | √ | If duration is repeating, the number of months the coupon applies. Null if coupon duration is forever or once. | 
| livemode | bool | X | √ | Has the value true if the coupon exists in live mode or the value false if the coupon exists in test mode. | 
| valid | bool | X | √ | Taking account of the above properties, whether this coupon can still be applied to a customer. | 
| created | timestamp | X | √ | Time at which the coupon was created. | 
| currency | string | X | √ | If amount_off has been set, the three-letter ISO code for the currency of the amount to take off. | 
| deleted | bool | X | √ | True if the customer is marked as deleted. | 
| duration | string | X | √ | One of forever, once, and repeating. Describes how long a customer who applies this coupon will get the discount. | 
| max_redemptions | int | X | √ | Maximum number of times this coupon can be redeemed, in total, across all customers, before it is no longer valid. | 


