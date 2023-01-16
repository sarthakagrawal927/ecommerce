# Design

## Backend

### Datastores
- PostgreSQL - Primary database
  - Tables: Product, User, Company, Feedback, Inbox, Transactions, FAQs, Bookmarks/Cart, Orders, ProductCat, ProductTag, Address, BillingInfo
- MongoDB - For storing large documents of products & users (product detail & user detail in future) - Very less used
  - Documents: ProductDetails  - later stage
- KeyValue Store & meilisearch/elasticSearch - In future.

## Frontend

### Pages List
- General
  - Landing
  - About Company
  - User Page
  - Single Product
  - Products Page

- Admin
  - CRUD Products
  - Edit Company