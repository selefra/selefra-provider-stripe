# Table: stripe_product

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ | An arbitrary string attached to the product. Often useful for displaying to users. | 
| unit_label | string | X | √ | A label that represents units of this product in Stripe and on customers’ receipts and invoices. When set, this will be included in associated invoice line item descriptions. | 
| livemode | bool | X | √ | Has the value true if the product exists in live mode or the value false if the product exists in test mode. | 
| metadata | json | X | √ | Set of key-value pairs that you can attach to an product. This can be useful for storing additional information about the product in a structured format. | 
| statement_descriptor | string | X | √ | Extra information about a product which will appear on your customer’s credit card statement. In the case that multiple products are billed at once, the first statement descriptor will be used. | 
| url | string | X | √ | A URL of a publicly-accessible webpage for this product. | 
| active | bool | X | √ | Whether the product is currently available for purchase. | 
| created | timestamp | X | √ | Time at which the product was created. | 
| package_dimensions | json | X | √ | The dimensions of this product for shipping purposes. | 
| id | string | X | √ | Unique identifier for the product. | 
| name | string | X | √ | The product’s full name or business name. | 
| type | string | X | √ | The product type. | 
| images | json | X | √ | A list of up to 8 URLs of images for this product, meant to be displayable to the customer. | 
| shippable | bool | X | √ | Whether this product is shipped (i.e., physical goods). | 
| updated | timestamp | X | √ | Time at which the product was updated. | 


