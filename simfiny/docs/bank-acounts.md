## Bank Accounts

```http title="HTTP"
BASE URI - /api/v1/financial-service/bank-account
```

### Get All Bank Accounts

List all bank accounts for the current user

```http title="HTTP"
GET /api/v1/financial-service/bank-accounts/user/{userId}
```

### Request Path

| Attribute         | Type     | Required   | Description                                                              |
| ----------------- | -------- | ---------- | ------------------------------------------------------------------------ |
| `userId`   | number   | yes        | The ID of the user whose bank accounts we want to get

### Update Bank Account

Updates a bank account (only manually linked accounts)

```http title="HTTP"
PUT /api/v1/financial-service/bank-account/{bankAccountId}
```

### Request Path

| Attribute         | Type     | Required   | Description                                                              |
| ----------------- | -------- | ---------- | ------------------------------------------------------------------------ |
| `bankAccountId`   | number   | yes        | The bank account we want to update

### Get Bank Account Balance

get the bank account balances

```http title="HTTP"
GET /api/v1/financial-service/bank-account/{bankAccountId}/balance
```

### Request Path

| Attribute         | Type     | Required   | Description                                                              |
| ----------------- | -------- | ---------- | ------------------------------------------------------------------------ |
| `bankAccountId`   | number   | yes        | The bank account we want to update

### Create Bank Account (Manual Link)

Create a bank account for
the provided link. Note: Bank accounts can only be
created this way for manual links. Attempting to create a bank
account for a link that is associated with Plaid will result in an error.

```http title="HTTP"
POST /api/v1/financial-service/bank-account/user/{userId}
```

### Request Path

| Attribute         | Type     | Required   | Description                                                              |
| ----------------- | -------- | ---------- | ------------------------------------------------------------------------ |
| `userId`   | number   | yes        | The userId to whom this account should be tied to
