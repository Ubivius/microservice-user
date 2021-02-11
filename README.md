# microservice-user
Microservice to manage users.

Beginning of code was based off of Nic Jackson's building microservices in Go : https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_


**Get Users**
----
  Returns json data about every user.

* **URL**

  /users

* **Method:**

  `GET`
  
*  **URL Params**

  None

* **Data Params**

  None


**Get User By Id**
----
  Returns json data about a specific user.

* **URL**

  /users/:id

* **Method:**

  `GET`
  
*  **URL Params**

   **Required:**
 
   `id=[integer]`

* **Data Params**

  None


**Add new user**
----
  Add new user with specific data

* **URL**

  /users

* **Method:**

  `POST`
  
*  **URL Params**
 
  None

* **Data Params**

  **Required:**

  `username=[string]`
  `email=[string]`
  `dateofbirth=[string]`


**Update user**
----
  Update user data

* **URL**

  /users

* **Method:**

  `PUT`
  
*  **URL Params**
 
  None

* **Data Params**

  **Required:**

  `id=[integer]`


**Delete user**
----
  Delete user

* **URL**

  /users/:id

* **Method:**

  `DELETE`
  
*  **URL Params**
 
  **Required:**
 
   `id=[integer]`

* **Data Params**

  None