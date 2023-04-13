# API Usage Guide

This API provides endpoints to perform CRUD operations for user accounts and to perform transactions between accounts.

## User Accounts

### Create a new user
To create a new user, send a POST request to `/users` with the following fields in the request body:
`
{
"id": "03eb9399-e526-431f-812f-2fda01659000",
"name": "Test User",
"balance": 1000
}
`
The `id` field must be a string, while the `name` field must also be a string, and `balance` must be a float.

### Get a user
To retrieve a specific user, send a GET request to `/users/{id}` where `{id}` is the user ID.

### Update a user
To update a user, send a PUT request to `/users/{id}` with the fields you want to update in the request body.
`
{
"name": "Updated User",
"balance": 2000
}
`
### Delete a user
To delete a user, send a DELETE request to `/users/{id}` where `{id}` is the user ID.

## Transactions

### Make a transaction
To make a transaction between two users, send a POST request to `/transactions` with the following fields in the request body:
`
{
"from_id": "03eb9399-e526-431f-812f-2fda01659000",
"to_id": "b84f33d6-7849-4f9c-9f81-6d1ba6aa2b6c",
"amount": 500
}
`
The `from_id` field is the ID of the user making the transfer, while the `to_id` field is the ID of the user receiving the transfer. The `amount` field is the amount of the transfer.
