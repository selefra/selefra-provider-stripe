# Table: stripe_plan

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| billing_scheme | string | X | √ | Describes how to compute the price per period. Either per_unit or tiered. | 
| created | timestamp | X | √ | Time at which the plan was created. | 
| currency | string | X | √ | Three-letter ISO currency code, in lowercase. Must be a supported currency. | 
| interval_count | int | X | √ | The number of intervals (specified in the interval attribute) between subscription billings. For example, interval=month and interval_count=3 bills every 3 months. | 
| product_id | string | X | √ | ID of the product whose pricing this plan determines. | 
| nickname | string | X | √ | A brief description of the plan, hidden from customers. | 
| amount | int | X | √ | The unit amount in cents to be charged, represented as a whole integer if possible. Only set if billing_scheme=per_unit. | 
| metadata | json | X | √ | Set of key-value pairs that you can attach to an plan. This can be useful for storing additional information about the plan in a structured format. | 
| interval | string | X | √ | The frequency at which a subscription is billed. One of day, week, month or year. | 
| livemode | bool | X | √ | Has the value true if the plan exists in live mode or the value false if the plan exists in test mode. | 
| aggregate_usage | string | X | √ | Specifies a usage aggregation strategy for plans of usage_type=metered. Allowed values are sum for summing up all usage during a period, last_during_period for using the last usage record reported within a period, last_ever for using the last usage record ever (across period bounds) or max which uses the usage record with the maximum reported usage during a period. Defaults to sum. | 
| amount_decimal | float | X | √ | The unit amount in cents to be charged, represented as a decimal string with at most 12 decimal places. Only set if billing_scheme=per_unit. | 
| tiers | json | X | √ | Each element represents a pricing tier. This parameter requires billing_scheme to be set to tiered. | 
| tiers_mode | string | X | √ | Defines if the tiering price should be graduated or volume based. In volume-based tiering, the maximum quantity within a period determines the per unit price. In graduated tiering, pricing can change as the quantity grows. | 
| transform_usage | json | X | √ | Apply a transformation to the reported usage or set quantity before computing the amount billed. | 
| id | string | X | √ | Unique identifier for the plan. | 
| active | bool | X | √ | Whether the plan is currently available for purchase. | 
| usage_type | string | X | √ | Configures how the quantity per period should be determined. Can be either metered or licensed. licensed automatically bills the quantity set when adding it to a subscription. metered aggregates the total usage based on usage records. Defaults to licensed. | 
| deleted | bool | X | √ | True if the plan is marked as deleted. | 
| trial_period_days | int | X | √ | Default number of trial days when subscribing a customer to this plan using trial_from_plan=true. | 


