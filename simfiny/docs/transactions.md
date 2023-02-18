## Transactions

```http title="HTTP"
BASE URI - /api/v1/financial-service/bank-account/transactions
```

### Get Transactions

This endpoint enabled users to list transactions for a specified bank account id. Supported accounts include manual and non-manual links.
Transactions are returned sorted by the date they were authorized (descending) and then by their numeric Id (descending). This means that transactions that were first seen later will be higher in the
list than they may have actually occurred.

```http title="HTTP"
GET /api/v1/financial-service/bank-account/{bankAccountId}/transactions
```

### Request Path

| Attribute         | Type     | Required   | Description                                                              |
| ----------------- | -------- | ---------- | ------------------------------------------------------------------------ |
| `bankAccountId`   | number   | yes        | The ID of the bank account the spending objects belong to.               |

## Update Transaction

This endpoint updates a given transaction for a specified connected account

```http title="HTTP"
PUT /api/v1/financial-service/bank-account/{bankAccountId}/transaction/{transactionId}
```

### Request Path

| Attribute         | Type     | Required   | Description                                                              |
| ----------------- | -------- | ---------- | ------------------------------------------------------------------------ |
| `bankAccountId`   | number   | yes        | The ID of the bank account which the transaction is tied to.
| `transactionId`   | number   | yes        | The ID of the transaction to be updated.

## Delete Transaction

This endpoint deletes a given transaction for a specified connected account. Note this only works for manually
linked bank account

```http title="HTTP"
DELETE /api/v1/financial-service/bank-account/{bankAccountId}/transaction/{transactionId}
```

### Request Path

| Attribute         | Type     | Required   | Description                                                              |
| ----------------- | -------- | ---------- | ------------------------------------------------------------------------ |
| `bankAccountId`   | number   | yes        | The ID of the bank account which the transaction is tied to.
| `transactionId`   | number   | yes        | The ID of the transaction to be deleted.

## Create Transaction

This endpoint enabled us to create a given transaction for a specified connected account. Note this only works for manually linked bank account

```http title="HTTP"
POST /api/v1/financial-service/bank-account/{bankAccountId}/transaction
```

### Request Path

| Attribute         | Type     | Required   | Description                                                              |
| ----------------- | -------- | ---------- | ------------------------------------------------------------------------ |
| `bankAccountId`   | number   | yes        | The ID of the bank account which the transaction is tied to.
| `transactionId`   | number   | yes        | The ID of the transaction to be deleted.

## Get Transactions For Spending

Spending objects are used to represent how the user wants to
spend their money and how frequently they want to or need to spend it.
Spending objects create an
["earmark"](https://www.merriam-webster.com/dictionary/earmark) against the
            available balance of the bank account they are associated with.
This is then used to allow monetr to calculate an amount that is safe for
the user to spend at any given time; while making sure they still have funds
for their defined financial obligations.

```http title="HTTP"
POST /api/v1/financial-service/bank-account/transactions/spending/{spendingId}
```

### Request Path

| Attribute         | Type     | Required   | Description                                                              |
| ----------------- | -------- | ---------- | ------------------------------------------------------------------------ |
| `spendingId`   | number   | yes        | The ID of the spending object the transaction is tied to.
