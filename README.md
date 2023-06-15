# Users

* User object
```
{
  id: integer
  username: string
  email: string
  created_at: datetime(iso 8601)
  created_by: string
  updated_at: datetime(iso 8601)
  updated_by: string
}
```
**POST /login**
----
Login as user
* **Headers**
  - Content-Type: application/json
* **URL Params**  
  None
* **Data Params**  
  None
* **Success Response:**
* **Code:** 200  
  **Content:**
```
{
  username: string
  password: string
}
```
* **Error Response**
* **Error Response:**
  * **Code:** 404  
    **Content:**
    `{ message : "Wrong username/password" }`

**GET /logout**
----
Logout user
* **Headers**
  - Authorization: Bearer `<JWT Token>`
* **URL Params**  
  None
* **Data Params**  
  None
* **Success Response:**
* **Code:** 200

**GET /users**
----
Returns all users in the system.
* **Headers**
  - Content-Type: application/json
  - Authorization: Bearer `<JWT Token>`
* **URL Params**  
    None
* **Data Params**  
    username: string
    role: string
* **Success Response:**
* **Code:** 200  
  **Content:**
```
{
  message: string
  users: [
           {<user_object>},
           {<user_object>},
           {<user_object>}
         ]
}
```

**GET /users/:username**
----
Returns the specified user.
* **Headers**  
  - Content-Type: application/json  
  - Authorization: Bearer `<JWT Token>`
* **URL Params**  
  *Required:* `username=[string]`
* **Data Params**  
  None
* **Success Response:**
* **Code:** 200  
  **Content:**  `{ <user_object> }`
* **Error Response:**
    * **Code:** 404  
      **Content:** 
      `{ message : "User doesn't exist" }`  
    OR
    * **Code:** 401  
      **Content:** `{ message : "You are unauthorized to make this request." }`

**POST /users**
----
Creates a new User and returns the new object.
* **Headers**  
  - Content-Type: application/json
  - Authorization: Bearer `<JWT Token>`
* **URL Params**  
  None
* **Data Params**
```
  {
    name: string, //optional
    username: string, //required
    password: string, //required
  }
```
* **Success Response:**
  * **Code:** 200  
    **Content:**  `{ <user_object> }`
  * **Code:** 400
    **Content:** `{ message : "Username already taken" }`

**PATCH /users/:username**
----
Updates fields on the specified user and returns the updated object.
* **Headers**
  - Authorization: Bearer `<JWT Token>`
* **URL Params**  
    *Required:* `username=[string]`
* **Data Params**
```
  {
    name: string,
  }
```
* **Headers**  
  Content-Type: application/json  
  Cookie: Bearer `<OAuth Token>`
* **Success Response:**
* **Code:** 200  
  **Content:**  `{ <user_object> }`
* **Error Response:**
    * **Code:** 404  
      **Content:** `{ message : "User doesn't exist" }`  
      OR
    * **Code:** 401  
      **Content:** `{ message : "You are unauthorized to make this request." }`

**PATCH /me**
----
Updates fields on the specified user and returns the updated object.
* **Headers**
  - Authorization: Bearer `<JWT Token>`
* **URL Params**  
  *Required:* `username=[string]`
* **Data Params**
```
  {
    old_password: string //required
    password: string //required
  }
```
* **Headers**  
  Content-Type: application/json  
  Cookie: Bearer `<OAuth Token>`
* **Success Response:**
* **Code:** 200  
  **Content:**  `{ <user_object> }`
* **Error Response:**
  * **Code:** 404  
    **Content:** `{ error : "User doesn't exist" }`

**DELETE /users/:username**
----
Deletes the specified user.
* **URL Params**  
  *Required:* `username=[string]`
* **Data Params**  
  None
* **Headers**  
  Content-Type: application/json  
  Authorization: Bearer `<OAuth Token>`
* **Success Response:**
    * **Code:** 204
* **Error Response:**
    * **Code:** 404  
      **Content:** `{ error : "User doesn't exist" }`  
      OR
    * **Code:** 401  
      **Content:** `{ error : "You are unauthorized to make this request." }`

# Audit

* Audit Object
```
{
  id: integer,
  date_time: datetime(iso 8601),
  username: string,
  action: string,
  entity: string,
  entity_id: string,
  data_before: <JSON Object>,
  data_after: <JSON Object>
}
```

**GET /audits**
----
* **URL Params**  
  None
* **Data Params**
```
  {
    username: string
    object: string
    object_id: string
    date_start: datetime
    date_end: datetime
   }
``` 
* **Headers**  
  Content-Type: application/json  
  Authorization: Bearer `<OAuth Token>`
* **Success Response:**
  * **Code:** 204
    **Content:** 
    ```
      {
        message: "Success",
        data: <Audit Object>
      }
    ```
* **Error Response:**
  * **Code:** 404  
    **Content:** `{ error : "User doesn't exist" }`  
    OR
  * **Code:** 401  
    **Content:** `{ error : "You are unauthorized to make this request." }`


**POST /audits**
----
* **URL Params**  
  None
* **Data Params**
```
  {
    action: string
  }
```
* **Headers**  
  Content-Type: application/json  
  Authorization: Bearer `<OAuth Token>`
* **Success Response:**
  * **Code:** 204
    **Content:**
    ```
      {
        message: "Success",
        data: <Audit Object>
      }
    ```
* **Error Response:**
  * **Code:** 404  
    **Content:** `{ error : "User doesn't exist" }`  
    OR
  * **Code:** 401  
    **Content:** `{ error : "You are unauthorized to make this request." }`