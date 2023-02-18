# User And Accounts
```http title="HTTP"
BASE URI - /api/v1/financial-service/user-account
```

## Create Financial User + Account
Creates a new login, user and account. Logins are used for authentication (not needed for our use case), users tie authentication to an account, and accounts hold budgeting data and stripe.
```http title="HTTP"
POST - /api/v1/financial-service/user-account
```

## Delete Financial User + Account
Deletes a created account
```http title="HTTP"
DELETE - /api/v1/financial-service/user-account/{userID}
```
