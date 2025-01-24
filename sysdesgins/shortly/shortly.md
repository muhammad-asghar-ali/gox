# Design a URL Shortener - shortly

1.  **Requirements:**

    1. **Functional Requirements:**

       - Generate the short url for given long url.
       - Every url must be unique.
       - Redriect user to orignal long url when click on short url.
       - User can customize the url (optional).
       - Set the short user expriation time.
       - Provide analytics to link usage.
       - May save the user information (optional).

    2. **Non Functional Requirements:**
       - High availability (the service should up like 99.9% time).
       - Low latency (redirect to url should heppen in ms).
       - Scalability (the system handle 1M records per day).
       - Security to prevent malicious use, such as phishing.

2.  **Capacity Estimation:**

    1. **Assumptions:**
       - **Daily requests per day to short urls** ~ 1000000.
       - **Read and Write ratio:** 100:1 (for every URL creation, we expect 100 redirects).
       - **Peak Traffic:** 10x of the average load.
       - **Orignal Url length:** 100 characters.
    2. **Throughput Requirements:**
       - **Average write per second (WPS):** (1,000,000 requests / 24 \* 60 \* 60 seconds) ~ 12
       - **Peak WPS:** 12 \* 10 = 120
       - **Average read per second (RPS):** 12 \* 100 = 1200
       - **Peal RPS:** 10 \* 1200 = 12000
    3. **Storage Estimation:**
       - We need the following informations for each URL.
         - **Short URL:** 7 characters
         - **Long URL:** 100 characters
         - **CreationDate:** 8 bytes (timestamp)
         - **ExpirationDate:** 8 bytes (timestamp)
         - **ClickCount:** 4 bytes (integer)
         - **UserID:** 8 bytes.
       - Total stroage:
         - **Storage per URL:** 7 + 100 + 8 + 8 + 4 + 8 = 135 bytes
         - Stroage for one year:
           - **Total URLs per Year:** 1,000,000 × 365 = 365,000,000
           - **Total Storage per Year:** 365,000,000 × 135 bytes ~ 48 GB
    4. **Bandwidth Estimation:** (optional)
       Assuming the HTTP 301 redirect response size is about 500 bytes (includes headers and the short URL).
       - **Total Read Bandwidth per Day:** 10000000 \* 100 \* 500 bytes = 50 GB / day
       - **Peak Bandwidth:** 500 bytes × 12,000 RPS = 6 MB/s (the peak bandwidth could be as high average).
    5. **Caching Estimation:**
       - The system is ead heavy so using cache can reduce the latency for read requests.
       - Can cache hot URLs, can identify the URLs where 20% of the URLs generate 80% of the read traffic.
       - 1 million writes per day, and cache only 20%, so the formula will be:
         - 1M \* 0.2 \* 135 Bytes ~ 26M
         - Cache hit ratio: 90:10
    6. **Infrastructure Sizing:**
       - **API Servers:** start with 1-2 instances each capabile of 200 to 300 RPS.
       - **Database:** single database node to handle both storage and high read/write throughput.
       - **Cache Layer:** single node, depending on the load and cache hit ratio.

3.  **High Level Design:**

    On a high level, we would need following components in our design:

    - **Load Balancer:** Distributes incoming requests across multiple application servers.
    - **Application Servers:** Handles incoming requests for shortening URLs and redirecting users and analytics (optional).
    - **Database:** Stores mappings between short URLs and long URLs.
    - **Cache:** Stores frequently accessed URL mappings for faster retrieval.

    **NOTE:** we can split the services to write and read sperate services.

    ![Design](shortly.png)

4.  **Database Design:**

    1. **SQL vs NoSQL:**
       To choose right database we need to understand our need. Let consider some factors:
       - We need to store billion records.
       - Read queries are much higher then the write.
       - We don't need joins.
       - Highly scalable and available.

    Given these points, a NoSQL database like MongoDB, Cassandra are better option due to their ability to efficiently handle billions.

    2. **Schema Design:**
       In inital stage we need only two tables which are store data. One is to store the user related database and the other table store the information about url.

       1. users
       1. url_store

       ![Database Design](database_design.png)

5.  **System API Design:**
    Design RESTful APIs that are efficient and scalable. We need following API implement the basic CRUD on urls and also provide the user registration and login endpoints. Here are the LIST of APIs needed to achieve the core functionality of a system.

    1.  **URL Shortening:**
        **Endpoint: POST ->** `/api/v1/shorten`
        We can make access the user with and without onboarding to the system.

        - Sample Request body:

          ```json
          {
            "long_url": "https://www.example.com/some/very/long/url",
            "custom_alias": "optional_custom_alias",
            "expiration_date": "optional_expiration_date"
          }
          ```

        - Sample Response body:

          ```json
          {
            "short_url": "https://www.example.com/abc123"
          }
          ```

    2.  **URL Redirection API:**
        **Endpoint: GET ->** `/api/v1/{short_url_key}`

        ```json
           HTTP 302 Redirect to the original long URL
        ```

    3.  **User Registration:** (optional)
        **Endpoint: POST ->** `/api/v1/register`
        Onboard the user to the system

        - Sample Request body:

          ```json
          {
            "name": "name",
            "email": "email@email.com",
            "password": "password"
          }
          ```

        - Sample Response body: (OK case) and otherwise return error message

          ```json
          {
            "message": "user register successfully"
          }
          ```

    4.  **User Login:** (optional)
        **Endpoint: POST ->** `/api/v1/login`
        Give access user to the system.

        - Sample Request body:

          ```json
          {
            "email": "email@email.com",
            "password": "password"
          }
          ```

        - Sample Response body: (OK case) and otherwise return error message

          ```json
          {
            "access_token": "access_token",
            "refresh_token": "refresh_token"
          }
          ```
