# Users

* User object
```
{
  username: string
  email: string
  created_at: datetime(iso 8601)
  created_by: string
  updated_at: datetime(iso 8601)
  updated_by: string
  new_user: bool //flag if the account changed password or not
}
```
**POST /login**
----
Login
* **Headers**
    - Content-Type: application/json
* **URL Params**  
  None
* **Data Params**
```
{
  username: string
  password: string
}
```
* **Success Response:**\
    * **Code:** 200\
      **Header:** \
      ```Set-Cookies: SESSION_ID=<session-code>; domain=<ip-server>; httpOnly``` \
      **Content:**
      ```
      {
        code: 200
        message: "Authenticated"
        data: {
          username: string //username account
          role: string //role account
          new_user: bool //flag if the account changed password or not
        }
      }
      ```
* **Error Response**
* **Error Response:**
    * **Code:** 400\
      **Content:**
      ```
      {
        code: 400
        message: "Invalid Username/Password"
      }
      ```

**GET /logout**
----
Logout user
* **Headers**
    - Content-Type: application/json
    - Cookies: `Session_ID=<session-code>`
* **URL Params**  
  None
* **Data Params** \
  None
* **Success Response:**
    * **Code:** 200 \
      **Header:** \
      ```Set-Cookies: SESSION_ID=<session-code>; Max-age=0; domain=<ip-server>; httpOnly```

**GET /users**
----
Get users in the system.
* **Headers**
    - Content-Type: application/json
    - Cookies: `Session_ID=<session-code>`
* **URL Params**
    - perpage
    - page
    - username
    - role (val: user, admin)
* **Data Params**  
  None
* **Success Response:**
    * **Code:** 200\
      **Content:**
      ```
      {
        code: 200
        message: "Success retrieve audit"
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
    - Cookies: `Session_ID=<session-code>`
* **URL Params**  
  None
* **Data Params**  
  None
* **Success Response:**
    * **Code:** 200  
      **Content:**
      ```
      {
        code: 200
        message: Success retrieve user
        data: <user_object>
      }
      ```
* **Error Response:**
    * **Code:** 404  
      **Content:**
      ```
        {
          code: 404
          message : "User not found" 
        }
      ```
    * **Code:** 401  
      **Content:**
      ```
        { 
          code: 401
          message : "Unauthorized" 
        }
      ```

**POST /users**
----
Creates a new User and returns the new object.
* **Headers**
    - Content-Type: application/json
    - Cookies: `Session_ID=<session-code>`
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
      **Content:**
      ```
      {
        code: 200
        message: "Success create user"
        data: <user_object>
      }
      ```
* **Error Response:**
    * **Code:** 400\
      **Content:**
      ```
      {
        code: 400
        message : "Username already taken" 
      }
      ```

**PATCH /users/:username**
----
Updates fields on the specified user and returns the updated object.
* **Headers**
    - Content-Type: application/json
    - Cookies: `Session_ID=<session-code>`
* **URL Params**  
  None
* **Data Params**
  ```
  {
    name: string
    old_password: string \\required if field password is not null
    password: string
  }
  ```
* **Success Response:**
    * **Code:** 200  
      **Content:**
      ```
      {
        code: 200
        message: "Success update user" 
        data: <user_object>
      }
      ```
* **Error Response:**
    * **Code:** 404  
      **Content:**
      ```
      {
      code: 400
      message : "Username already taken"
      }
      ```
    * **Code:** 401  
      **Content:**
      ```
      {
      code: 401
      message : "Unauthorized"
      }
      ```

**GET /me**
----
Get profile for authenticated user.
* **Headers**
    - Content-Type: application/json
    - Cookies: `Session_ID=<session-code>`
* **URL Params**\
  None
* **Data Params**\
  None
* **Success Response:**
    * **Code:** 200  
      **Content:**
      ```
      {
        "code": 200
        "message": "Success retrieve data"
        "data": {
          "<user-object>"
        }
      }
      ```
* **Error Response:**
    * **Code:** 401  
      **Content:**
      ```json
      {
        "code": 200,
        "message": "Unauthorized"
      }
      ```      

**PATCH /me**
----
Update profile for authenticated user.
* **Headers**
    - Content-Type: application/json
    - Cookies: `Session_ID=<session-code>`
* **URL Params**  
  None
* **Data Params**
  ```
    {
      name: string //required
    }
  ```
* **Success Response:**
    * **Code:** 200  
      **Content:**
      ```json
      {
        "code": 200,
        "message": "Success update profile"
      }
      ```
* **Error Response:**
    * **Code:** 401  
      **Content:**
      ```json
      {
        "code": 200,
        "message": "Unauthorized"
      }
      ```    

**PATCH /me/password**
----
Change password for authenticated user.
* **Headers**
    - Content-Type: application/json
    - Cookies: `Session_ID=<session-code>`
* **URL Params**  
  None
* **Data Params**
  ```
    {
      old_password: string //required
      password: string //required
    }
  ```
* **Success Response:**
    * **Code:** 200  
      **Content:**
      ```json
      {
        "code": 200,
        "message": "Success update password"
      }
      ```
* **Error Response:**
    * **Code:** 401  
      **Content:**
      ```json
      {
        "code": 200,
        "message": "Unauthorized"
      }
      ```

**DELETE /users/:username**
----
Deletes the specified user.
* **Headers**
    - Content-Type: application/json
    - Cookies: `Session_ID=<session-code>`
* **URL Params**  
  None
* **Data Params**  
  None
* **Success Response:**
    * **Code:** 204
* **Error Response:**
    * **Code:** 401  
      **Content:**
      ```
      {
       code: 401
       message: "Unauthorized"
      }
      ```

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
Get data audits
* **Headers**
    - Content-Type: application/json
    - Cookies: `Session_ID=<session-code>`
* **URL Params**
    - page
    - perpage
    - username
    - object
    - object_id
    - from (YYYY-MM-DD)
    - to (YYYY-MM-DD)
* **Data Params** \
  None
* **Success Response:**
    * **Code:** 200
      **Content:**
      ```
      {
        code: 200
        message: "Success retrieve audit",
        data: [
          <Audit Object>,
          <Audit Object>,
          <Audit Object>
        ]
      }
      ```
* **Error Response:**
    * **Code:** 401  
      **Content:**
      ```json
      {
        "code": 401,
        "message": "Unauthorized"
      }
      ```

**POST /audits**
----
Create audit data with specified action
* **Headers**
    - Content-Type: application/json
    - Cookies: `Session_ID=<session-code>`
* **URL Params**  
  None
* **Data Params**
```
  {
    action: string
  }
```
* **Success Response:**
    * **Code:** 204
      **Content:**
      ```json
        {
          "code": 204,
          "message": "Success"
        }
      ```
* **Error Response:**
    * **Code:** 401  
      **Content:**
      ```json
      {
        "code": 401,
        "message": "Unauthorized"
      }
      ```

**GET /audits/export**
----
Download audits data as csv
* **Headers**
    - Content-Type: application/json
    - Cookies: `Session_ID=<session-code>`
* **URL Params**
    - username
    - object
    - object_id
    - from (YYYY-MM-DD)
    - to (YYYY-MM-DD)
* **Data Params**\
  None
* **Success Response:**
    * **Code:** 200
      **Header:**
        - Content-Disposition: attachment; filename=`filename`
* **Error Response:**
    * **Code:** 401  
      **Content:**
      ```json
      {
        "code": 401,
        "message": "Unauthorized"
      }
      ```

# BRIVA

* BRIVA Object
```
{
  Brivano: string
  CorpName: string
  IsActive: string
}
```

**GET /briva/:brivano**
----
Get data briva with brivano
* **Headers**
    - Content-Type: application/json
    - Cookies: `Session_ID=<session-code>`
* **URL Params**  \
  None
* **Data Params** \
  None
* **Success Response:**
    * **Code:** 200
      **Content**
      ```
      {
        code: 200
        message: "Brivano has been found"
        data: <BRIVA object>
      }
      ```
* **Error Response:**
    * **Code:** 401  
      **Content:**
      ```json
      {
        "code": 401,
        "message": "Unauthorized"
      }
      ```

**POST /briva/:brivano**
----
Change BRIVA's active state
* **Headers**
    - Content-Type: application/json
    - Cookies: `Session_ID=<session-code>`
* **URL Params**  
  None
* **Data Params**
  ```
  {
    active: string (val: "1" or "0")
  }
  ```
* **Success Response:**
    * **Code:** 200
      **Content**
      ```
      {
        code: 200
        message: "Success update briva"
      }
      ```
* **Error Response:**
    * **Code:** 401  
      **Content:**
      ```json
      {
        "code": 401,
        "message": "Unauthorized"
      }
      ```

# SPAN

* SPAN Object
```
{
  DocumentNumber: string
  DocumentDate: string
  BeneficiaryBankCode: string
  StatusCode: string
  EmailAddress: string
  BeneficiaryAccount: string
  Amount: string
  BeneficiaryBank: string
  is_patched: bool //status if span already patched or not
}
```

**GET /span/:documentNumber**
----
Get SPAN data with document number
* **Headers**
    - Content-Type: application/json
    - Cookies: `Session_ID=<session-code>`
* **URL Params**  
  None
* **Data Params**\
  None
* **Success Response:**
    * **Code:** 200
      **Content**
      ```
      {
        code: 200
        message: "Success retrieve span"
        data: <SPAN object>
      }
      ```
* **Error Response:**
    * **Code:** 401  
      **Content:**
      ```json
      {
        "code": 401,
        "message": "Unauthorized"
      }
      ```

**POST /SPAN/:documentNumber**
----
Patch data span for specific document number
* **Headers**
    - Content-Type: application/json
    - Cookies: `Session_ID=<session-code>`
* **URL Params**  
  None
* **Data Params**
  None
* **Success Response:**
    * **Code:** 200
      **Content**
      ```
      {
        code: 200
        message: "Success update SPAN"
      }
      ```
* **Error Response:**
    * **Code:** 401  
      **Content:**
      ```json
      {
        "code": 401,
        "message": "Unauthorized"
      }
      ```

# Server Utilization

* Server Utilization Object
```
{
  hostname: string
  cpu_percentage: string //float number (percentage)
  memory_usage: string //float number (percentage)
  system_uptime: string //float number (percentage)
  disks: [
    <name-disk>: string //float number (percentage),
    <name-disk>: string //float number (percentage)
  ]
}
```

**GET /server-utilization/latest-data**
----
Get latest data for all server monitoring
* **Headers**
    - Content-Type: application/json
    - Cookies: `Session_ID=<session-code>`
* **URL Params**  
  None
* **Data Params**
  None
* **Success Response:**
    * **Code:** 200
      **Content**
      ```
      {
        code: 200
        message: "Success get latest data"
        data: {
          safe: [
            <Server Utilization Object>,
            <Server Utilization Object>,
            <Server Utilization Object>
          ],
          threshold: [
            <Server Utilization Object>,
            <Server Utilization Object>,
            <Server Utilization Object>
          ]
        }
      }
      ```
* **Error Response:**
    * **Code:** 401  
      **Content:**
      ```json
      {
        "code": 401,
        "message": "Unauthorized"
      }
      ```

**GET /server-utilization/update-host**
----
Update server monitoring list
* **Headers**
    - Content-Type: application/json
    - Cookies: `Session_ID=<session-code>`
* **URL Params**  
  None
* **Data Params**
  None
* **Success Response:**
    * **Code:** 200
      **Content**
      ```
      {
        code: 200
        message: "Success update host list"
      }
      ```
* **Error Response:**
    * **Code:** 401  
      **Content:**
      ```json
      {
        "code": 401,
        "message": "Unauthorized"
      }
      ```

# Configuration

**POST /config/session-duration**
----
Change session duration IDLE
* **Headers**
    - Content-Type: application/json
    - Cookies: `Session_ID=<session-code>`
* **URL Params**  
  None
* **Data Params**
  ```
  {
    duration: integer //in seconds
  }
  ```
* **Success Response:**
    * **Code:** 200
      **Content**
      ```
      {
        code: 200
        message: "Success update configuration"
      }
      ```
* **Error Response:**
    * **Code:** 401  
      **Content:**
      ```json
      {
        "code": 401,
        "message": "Unauthorized"
      }
      ```