# REST API

This document describes specification for REST API endpoint used by the system.

List of available endpoints:

- [Register](#register)
- [Login](#login)
- [List Users](#list-users)
- [Request Upload URL](#request-upload-url)
- [Upload Project Image](#upload-project-image)
- [Submit Project](#submit-project)
- [Review Project by Admin](#review-project-by-admin)
- [List Projects](#list-projects)
- [Get Project](#get-project)
- [Donate to Project](#donate-to-project)
- [List Project Donations](#list-project-donations)

Beside these endpoints, client need to handle properly [Common Errors](./common_errors.md) that may occurs throughout all specified REST API endpoints in the system.

## Register 

POST: `/user/register`

This endpoint is used to register a new user to the system. It's provided for both donor and requester.

**Example Request:**

```json
POST /register
Content-Type: application/json

{
  "username": "alfonso",
  "email": "alfonso@gmail.com",
  "password": "123456",
  "role": "donor"
}
```

> Note:
> - `role` value must be either `donor` or `requester`.

**Success Response:**

```json
HTTP/1.1 200 OK
Content-Type: application/json

{
  "ok": true,
  "data": {
    "id": 1,
    "username": "alfonso",
    "email": "alfonso@gmail.com"
  },
  "ts": 1704954526
}
```

**Error Response:**

No specific error response.

[Back to Top](#rest-api)

---

## Login

POST: `/user/login`

This endpoint is used to login to the system. It returns a JWT token that can be used to access other endpoints named `access_token`.

The decoded access token will contains at least the following information:

```json
{
  "id": 1,
  "username": "alfonso",
  "email": "alfonso@gmail.com"
}
```

The JWT secret will be `donation-hub`, and the access token will be valid for 1 year.

**Example Request:**

```json
POST /session
Content-Type: application/json

{
  "username": "alfonso",
  "password": "123456"
}
```
**Success Response:**

```json
HTTP/1.1 200 OK
Content-Type: application/json

{
  "ok": true,
  "data": {
    "id": 1,
    "username": "alfonso",
    "email": "alfonso@gmail.com",
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJhbGZvbnNvIiwiaWF0IjoxNzM2NTY4MDAwfQ.rWcjBC_4Xg7-P37cdQ6JemVWSFgTQVKl4L_Yt7NMJu4",
  },
  "ts": 1704954526
}
```

**Error Response:**

- Invalid Credentials

  ```json
  HTTP/1.1 401 Unauthorized
  Content-Type: application/json

  {
    "ok": false,
    "err": "ERR_INVALID_CREDS",
    "msg": "Invalid username or password",
    "ts": 1704954526
  }

[Back to Top](#rest-api)

---

## List Users

GET: `/users`

This endpoint is used to list all users available in the system.

**Query Params:**

- `limit` => Limit the number of users returned. The value is integer.
- `page` => The page number. The value is integer. This is used for pagination.
- `role` => Filter users based on role. The value is `donor`, `requester`. If not provided, it will return all users with role `donor` and `requester`.

**Example Request:**

```json

GET /users?limit=10&page=1&role=donor

```

**Success Response:**

  ```json
  HTTP/1.1 200 OK
  Content-Type: application/json

  {
    "ok": true,
    "data": {
        "users": [
            {
                "id": 1,
                "username": "johndoe",
                "email": "john@gmail.com",
                "roles": ["donor"]
            },
            {
                "id": 2,
                "username": "janedoe",
                "email": "jane@gmail.com",
                "roles": ["donor, requester"]
            }
        ],
        "page": 1,
        "total_pages": 10
    },
    "ts": 1704954526
  }
  ```

**Error Response:**

No specific error response.

[Back to Top](#rest-api)

---

## Request Upload URL

GET: `/projects/upload`

This endpoint is used to request an upload URL for a new image. The image will be used to upload a new project.

**Headers:**

- `Authorization` => The value is `Bearer {access_token}` with user role `donor`.

**Query Params:**

- `mime_type` (String): mime type of file to be uploaded, valid values are `image/jpeg`, `image/png`.
- `file_size` (Integer): size of to be uploaded file in bytes, the maximum accepted value are `1 MB` or `1048576` for image.

**Example Request:**

```json

GET /projects/upload?mime_type=image/jpeg&file_size=151
Authorization: Bearer {access_token}

```

**Success Response:**

```json
HTTP/1.1 200 OK
Content-Type: application/json

{
    "status": 200,
    "data": {
        "mime_type": "image/jpeg",
        "file_size": 15125,
        "url": "https://media-donation-hub.ap-southeast-3.amazonaws.com/images/MTcxNDk2Mzg4MTE4OTU4MTIzMzYyNA.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIA3SPNTG6NFOYVY3M6%2F20200614%2Fap-southeast-3%2Fs3%2Faws4_request&X-Amz-Date=20200614T124720Z&X-Amz-Expires=60&X-Amz-SignedHeaders=content-type%3Bhost%3Bx-amz-acl&X-Amz-Signature=57ee52ad6e091419cf5b442caff07d4198a79ab146acd6fb675da59c81d01126",
        "expires_at": 1592139327
    },
    "ts": "2024-06-14T12:45:39.783Z"
}
```

**Error Response:**

- File too large:

    ```json
    HTTP/1.1 200 OK
    Content-Type: application/json

    {
        "status": 413,
        "err": "ERR_FILE_TOO_LARGE",
        "ts": "2018-02-02T08:30:39.765Z"
    }
    ```

    Client will get this error if the value of `file_size` exceeding allowed values.

- Unsupported `type`:

    ```json
    {
        "status": 415,
        "err": "ERR_UNSUPPORTED_TYPE",
        "ts": "2018-02-02T08:30:39.765Z"
    }
    ```

    Client will get this error if the value of `type` is not supported.

[Back to Top](#rest-api)

---

## Upload Project Image

PUT: `/{url_from_request_upload}`

This endpoint is used to upload images for a new project directly to the S3 bucket from the client side. Please check [this](https://aws.amazon.com/blogs/compute/uploading-to-amazon-s3-directly-from-a-web-or-mobile-application/) as a reference.

**Example Request:**
```binary
{binary_data}
```

**Success Response:**

```
HTTP/1.1 200 OK
```

**Error Response:**

- Forbidden:

    ```xml
    HTTP/1.1 403 Forbidden
    Content-Type: application/xml

    <?xml version="1.0" encoding="UTF-8"?>
    <Error>
        <Code>SignatureDoesNotMatch</Code>
        <Message>The request signature we calculated does not match the signature you provided. Check your key and signing method.</Message>
        <AWSAccessKeyId>ASIA3SPNTG6NLIRKRS5Q</AWSAccessKeyId>
        <StringToSign>AWS4-HMAC-SHA256
    20240506T025121Z
    20240506/ap-southeast-3/s3/aws4_request
    e4d640228354a7f6f868b2c14d5f23412a1a65f742deeb2d7261c22347be7c7a</StringToSign>
        <SignatureProvided>612bb8ba847a5d244c2e35e559828ddf5a7172511748c8727f5ea7ce622c58cf</SignatureProvided>
        <StringToSignBytes>41 57 53 34 2d 48 4d 41 43 2d 53 48 41 32 35 36 0a 32 30 32 34 30 35 30 36 54 30 32 35 31 32 31 5a 0a 32 30 32 34 30 35 30 36 2f 65 75 2d 77 65 73 74 2d 31 2f 73 33 2f 61 77 73 34 5f 72 65 71 75 65 73 74 0a 65 34 64 36 34 30 32 32 38 33 35 34 61 37 66 36 66 38 36 38 62 32 63 31 34 64 35 66 32 33 34 31 32 61 31 61 36 35 66 37 34 32 64 65 65 62 32 64 37 32 36 31 63 32 32 33 34 37 62 65 37 63 37 61</StringToSignBytes>
        <CanonicalRequest>PUT
    /images/MTcxNDk2Mzg4MTE4OTU4MTIzMzYyNA.jpg
    X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=ASIA3SPNTG6NLIRKRS5Q%2F20240506%2Fap-southeast-3%2Fs3%2Faws4_request&amp;X-Amz-Date=20240506T025121Z&amp;X-Amz-Expires=600&amp;X-Amz-Security-Token=IQoJb3JpZ2luX2VjEFsaCWV1LXdlc3QtMSJIMEYCIQCjmmavrluycraKTaJZEyGCCIXe88tmK4cE2FwdcIPa%2FwIhANbFr2yQ9i6B68JsLA7cgovrPgCnMVtuT%2FAeU4SLlpUeKo4ECLT%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FwEQBBoMNzk1NjA0MTA1MTE0IgwsvwJUeiv8Q8yzBfwq4gOjVDZZunHb1NPxMPRWMBHEtrBkmOfSloHSYHHyxAWC3GU68LZHHnT4vMelEFLkceCibvrEYK5P7qseQCgKvz4dPpSWI6bUCIq0Y58XivTppYry2l%2FnO7zztk1E0fa%2FUXg%2BHY1jrpkGJnndp9CxFvCsUG9sJp4%2FQpG3ODC7VD9b13uAlgiX8ZyLdKNxlOWh2upJiZhqN06O%2FpiC%2Fd8U9Ni%2F9t28Q5H%2BhReOT0jZ2ur%2FRVQ3FL%2Bc7%2F%2F8wZLGfYUFALSzapVbA62tKpLivGJXsFBhFcB2m5pSg8uWsfPMqE3yhUCtJxBjWUVZuDfo%2BoDu8QroSHKJN5kPDV3%2FF700TIGFXd5AcN6Lv3OZX7sQ2wFshvyGN81X2CBJc%2Fh5%2BsvKUFSNzpaFqX3ZcYd4bbTxGtsqq%2BO3l8z2F7vauBueOGvBMXale6aTS9kpfJoc%2BQwshUzJerqB26TdO6loOuytiC6WbyME3Cuh1yC8BOzFrEIGDhIQJiuYkz1uRLb7XPbKP6xnZ2J3rPpsZz7V1w%2BJ31qxmbtJdwSbm1pJYLH3%2FqMh6VsuxWLO376VR5MTxwoMKDnxeV6AclMUlgUocjMywQpawjikdl9XkSs7GO6Cs8rRroAEs7WNKT5psjS90xjGjJ1bFTCy%2F%2BCxBjqkARMJ%2FS%2Bw4vUe7foGaNBfdkdxqOobfj9iRPqEeR5hdMq2HpfPwLNoSfdiMH%2Bi4fhN2MJXqQd60CEvx6qaebz8RG9pWIv0rZK%2B%2FhfYB3IpB7VFP1JdlQMN%2BtmeawvmsRqK0LN3MoJfKGvLUlMbpGEJw8ykxmYB3Gng3Bq7dDQA9Gt5pIPTmajZoupQa6SbvV0Y8Z8HtqkDrofpRgFM0AudMKGitLqF&amp;X-Amz-SignedHeaders=content-length%3Bcontent-type%3Bhost
    content-length:192
    content-type:application/octet-stream
    host:media-donation-hub.s3.ap-southeast-3.amazonaws.com

    content-length;content-type;host
    UNSIGNED-PAYLOAD</CanonicalRequest>
        <CanonicalRequestBytes>50 55 54 0a 2f 69 6d 61 67 65 73 2f 4d 54 63 78 4e 44 6b 32 4d 7a 67 34 4d 54 45 34 4f 54 55 34 4d 54 49 7a 4d 7a 59 79 4e 41 2e 6a 70 67 0a 58 2d 41 6d 7a 2d 41 6c 67 6f 72 69 74 68 6d 3d 41 57 53 34 2d 48 4d 41 43 2d 53 48 41 32 35 36 26 58 2d 41 6d 7a 2d 43 72 65 64 65 6e 74 69 61 6c 3d 41 53 49 41 33 53 50 4e 54 47 36 4e 4c 49 52 4b 52 53 35 51 25 32 46 32 30 32 34 30 35 30 36 25 32 46 65 75 2d 77 65 73 74 2d 31 25 32 46 73 33 25 32 46 61 77 73 34 5f 72 65 71 75 65 73 74 26 58 2d 41 6d 7a 2d 44 61 74 65 3d 32 30 32 34 30 35 30 36 54 30 32 35 31 32 31 5a 26 58 2d 41 6d 7a 2d 45 78 70 69 72 65 73 3d 36 30 30 26 58 2d 41 6d 7a 2d 53 65 63 75 72 69 74 79 2d 54 6f 6b 65 6e 3d 49 51 6f 4a 62 33 4a 70 5a 32 6c 75 58 32 56 6a 45 46 73 61 43 57 56 31 4c 58 64 6c 63 33 51 74 4d 53 4a 49 4d 45 59 43 49 51 43 6a 6d 6d 61 76 72 6c 75 79 63 72 61 4b 54 61 4a 5a 45 79 47 43 43 49 58 65 38 38 74 6d 4b 34 63 45 32 46 77 64 63 49 50 61 25 32 46 77 49 68 41 4e 62 46 72 32 79 51 39 69 36 42 36 38 4a 73 4c 41 37 63 67 6f 76 72 50 67 43 6e 4d 56 74 75 54 25 32 46 41 65 55 34 53 4c 6c 70 55 65 4b 6f 34 45 43 4c 54 25 32 46 25 32 46 25 32 46 25 32 46 25 32 46 25 32 46 25 32 46 25 32 46 25 32 46 25 32 46 77 45 51 42 42 6f 4d 4e 7a 6b 31 4e 6a 41 30 4d 54 41 31 4d 54 45 30 49 67 77 73 76 77 4a 55 65 69 76 38 51 38 79 7a 42 66 77 71 34 67 4f 6a 56 44 5a 5a 75 6e 48 62 31 4e 50 78 4d 50 52 57 4d 42 48 45 74 72 42 6b 6d 4f 66 53 6c 6f 48 53 59 48 48 79 78 41 57 43 33 47 55 36 38 4c 5a 48 48 6e 54 34 76 4d 65 6c 45 46 4c 6b 63 65 43 69 62 76 72 45 59 4b 35 50 37 71 73 65 51 43 67 4b 76 7a 34 64 50 70 53 57 49 36 62 55 43 49 71 30 59 35 38 58 69 76 54 70 70 59 72 79 32 6c 25 32 46 6e 4f 37 7a 7a 74 6b 31 45 30 66 61 25 32 46 55 58 67 25 32 42 48 59 31 6a 72 70 6b 47 4a 6e 6e 64 70 39 43 78 46 76 43 73 55 47 39 73 4a 70 34 25 32 46 51 70 47 33 4f 44 43 37 56 44 39 62 31 33 75 41 6c 67 69 58 38 5a 79 4c 64 4b 4e 78 6c 4f 57 68 32 75 70 4a 69 5a 68 71 4e 30 36 4f 25 32 46 70 69 43 25 32 46 64 38 55 39 4e 69 25 32 46 39 74 32 38 51 35 48 25 32 42 68 52 65 4f 54 30 6a 5a 32 75 72 25 32 46 52 56 51 33 46 4c 25 32 42 63 37 25 32 46 25 32 46 38 77 5a 4c 47 66 59 55 46 41 4c 53 7a 61 70 56 62 41 36 32 74 4b 70 4c 69 76 47 4a 58 73 46 42 68 46 63 42 32 6d 35 70 53 67 38 75 57 73 66 50 4d 71 45 33 79 68 55 43 74 4a 78 42 6a 57 55 56 5a 75 44 66 6f 25 32 42 6f 44 75 38 51 72 6f 53 48 4b 4a 4e 35 6b 50 44 56 33 25 32 46 46 37 30 30 54 49 47 46 58 64 35 41 63 4e 36 4c 76 33 4f 5a 58 37 73 51 32 77 46 73 68 76 79 47 4e 38 31 58 32 43 42 4a 63 25 32 46 68 35 25 32 42 73 76 4b 55 46 53 4e 7a 70 61 46 71 58 33 5a 63 59 64 34 62 62 54 78 47 74 73 71 71 25 32 42 4f 33 6c 38 7a 32 46 37 76 61 75 42 75 65 4f 47 76 42 4d 58 61 6c 65 36 61 54 53 39 6b 70 66 4a 6f 63 25 32 42 51 77 73 68 55 7a 4a 65 72 71 42 32 36 54 64 4f 36 6c 6f 4f 75 79 74 69 43 36 57 62 79 4d 45 33 43 75 68 31 79 43 38 42 4f 7a 46 72 45 49 47 44 68 49 51 4a 69 75 59 6b 7a 31 75 52 4c 62 37 58 50 62 4b 50 36 78 6e 5a 32 4a 33 72 50 70 73 5a 7a 37 56 31 77 25 32 42 4a 33 31 71 78 6d 62 74 4a 64 77 53 62 6d 31 70 4a 59 4c 48 33 25 32 46 71 4d 68 36 56 73 75 78 57 4c 4f 33 37 36 56 52 35 4d 54 78 77 6f 4d 4b 44 6e 78 65 56 36 41 63 6c 4d 55 6c 67 55 6f 63 6a 4d 79 77 51 70 61 77 6a 69 6b 64 6c 39 58 6b 53 73 37 47 4f 36 43 73 38 72 52 72 6f 41 45 73 37 57 4e 4b 54 35 70 73 6a 53 39 30 78 6a 47 6a 4a 31 62 46 54 43 79 25 32 46 25 32 42 43 78 42 6a 71 6b 41 52 4d 4a 25 32 46 53 25 32 42 77 34 76 55 65 37 66 6f 47 61 4e 42 66 64 6b 64 78 71 4f 6f 62 66 6a 39 69 52 50 71 45 65 52 35 68 64 4d 71 32 48 70 66 50 77 4c 4e 6f 53 66 64 69 4d 48 25 32 42 69 34 66 68 4e 32 4d 4a 58 71 51 64 36 30 43 45 76 78 36 71 61 65 62 7a 38 52 47 39 70 57 49 76 30 72 5a 4b 25 32 42 25 32 46 68 66 59 42 33 49 70 42 37 56 46 50 31 4a 64 6c 51 4d 4e 25 32 42 74 6d 65 61 77 76 6d 73 52 71 4b 30 4c 4e 33 4d 6f 4a 66 4b 47 76 4c 55 6c 4d 62 70 47 45 4a 77 38 79 6b 78 6d 59 42 33 47 6e 67 33 42 71 37 64 44 51 41 39 47 74 35 70 49 50 54 6d 61 6a 5a 6f 75 70 51 61 36 53 62 76 56 30 59 38 5a 38 48 74 71 6b 44 72 6f 66 70 52 67 46 4d 30 41 75 64 4d 4b 47 69 74 4c 71 46 26 58 2d 41 6d 7a 2d 53 69 67 6e 65 64 48 65 61 64 65 72 73 3d 63 6f 6e 74 65 6e 74 2d 6c 65 6e 67 74 68 25 33 42 63 6f 6e 74 65 6e 74 2d 74 79 70 65 25 33 42 68 6f 73 74 0a 63 6f 6e 74 65 6e 74 2d 6c 65 6e 67 74 68 3a 31 39 32 0a 63 6f 6e 74 65 6e 74 2d 74 79 70 65 3a 61 70 70 6c 69 63 61 74 69 6f 6e 2f 6f 63 74 65 74 2d 73 74 72 65 61 6d 0a 68 6f 73 74 3a 64 65 76 63 68 61 74 2d 6d 65 64 69 61 2d 62 75 63 6b 65 74 2e 73 33 2e 65 75 2d 77 65 73 74 2d 31 2e 61 6d 61 7a 6f 6e 61 77 73 2e 63 6f 6d 0a 0a 63 6f 6e 74 65 6e 74 2d 6c 65 6e 67 74 68 3b 63 6f 6e 74 65 6e 74 2d 74 79 70 65 3b 68 6f 73 74 0a 55 4e 53 49 47 4e 45 44 2d 50 41 59 4c 4f 41 44</CanonicalRequestBytes>
        <RequestId>5MRXCKS5RDMP9HQN</RequestId>
        <HostId>NE68YefU5DDp73PZYPAgMAiu72W0jmfBXI5PcusdMkDDoAhhUsUuF51gZdEbz1dfDOUJUJtI3yI=</HostId>
    </Error>
    ```

    This error returned by AWS S3 if the request signature does not match the signature provided.

[Back to Top](#rest-api)

---

## Submit Project

POST: `/projects`

This endpoint is used to submit a new project to the system.

**Headers:**

- `Authorization` => The value is `Bearer {access_token}` with user role `donor`.

**Example Request:**

```json
POST /projects
Authorization: Bearer {access_token}

{
  "title": "Project Title",
  "description": "Project Description",
  "image_urls": [
    "https://media-donation-hub.ap-southeast-3.amazonaws.com/images/MTcxNDk2Mzg4MTE4OTU4MTIzMzYyNA.jpg",
    "https://media-donation-hub.ap-southeast-3.amazonaws.com/images/MTcxNDk2Mzg4MT38dsjndjnsdjs3.jpg"
  ]
  "due_at": 1704945600,
  "target_amount": 1000000,
  "currency": "USD"
}

```

**Success Response:**

  ```json
  HTTP/1.1 200 OK
  Content-Type: application/json

  {
    "ok": true,
    "data": {
      "id": 1,
      "title": "Project Title",
      "description": "Project Description",
      "image_urls": [
        "https://media-donation-hub.ap-southeast-3.amazonaws.com/images/MTcxNDk2Mzg4MTE4OTU4MTIzMzYyNA.jpg",
        "https://media-donation-hub.ap-southeast-3.amazonaws.com/images/MTcxNDk2Mzg4MT38dsjndjnsdjs3.jpg"
      ],
      "due_at": 1704945600,
      "target_amount": 1000000,
      "currency": "USD"
    },
    "ts": 1704954526
  }
  ```

**Error Response:**

- Invalid Due Time - Bad Request
  
    ```json
    HTTP/1.1 400 Bad Request
    Content-Type: application/json
  
    {
      "ok": false,
      "err": "ERR_BAD_REQUEST",
      "msg": "Due time must be at least 1 week from now",
      "ts": 1704954526
    }
    ```

    This error returned if the due time is less than 1 week from project submission time.

- Invalid Image URLs - Bad Request

    ```json
    HTTP/1.1 400 Bad Request
    Content-Type: application/json
  
    {
      "ok": false,
      "err": "ERR_BAD_REQUEST",
      "msg": "Image URLs must be between 1 and 4 URLs",
      "ts": 1704954526
    }
    ```

    This error returned if the image URLs is invalid.

[Back to Top](#rest-api)

---

## Review Project by Admin

PUT: `/projects/{project_id}/review`

This endpoint is used to review a project by admin. Admin can approve or reject a project.

**Headers:**

- `Authorization` => The value is `Bearer {access_token}` with user role `admin`.

**Example Request:**

```json

PUT /projects/1/review
Authorization: Bearer {access_token}

{
  "status": "approved"
}

```

**Success Response:**

- Approved:

  ```json
  HTTP/1.1 200 OK
  Content-Type: application/json

  {
    "ok": true,
    "ts": 1704954526
  }
  ```

- Rejected:

  ```json
  HTTP/1.1 200 OK
  Content-Type: application/json

  {
    "ok": true,
    "ts": 1704954526
  }
  ```

**Error Response:**

No specific error response.

[Back to Top](#rest-api)

---

## List Projects

GET: `/projects`

This endpoint is used to list all projects available in the system.

**Headers:**

- `Authorization`, __OPTIONAL__ => The value is `Bearer {access_token}`.

**Query Params:**

- `status` => Filter projects based on status. The value is `approved`, `rejected`, `need_review`, `completed`. If not provided, it will return projects with all status.
  - For `need_review` status, only admin can see the projects.
  
  > Note:
  > - `need_review`: initial status of the project after submission.
  > - `approved`: status of the project after admin review and approved.
  > - `rejected`: status of the project after admin review and rejected.
  > - `completed`: status of the project once the target amount is reached or the due time is reached.
- `start_ts` => Filter projects based on start timestamp. The value is integer.
- `end_ts` => Filter projects based on end timestamp. The value is integer.
- `limit` => Limit the number of projects returned. The value is integer. Default is 10.
- `last_key` => The last key of the previous request. The value is string. This is used for pagination.
  - If `last_key` is provided, the response will return the next projects after the last key.
  - If `last_key` is not provided, the response will return the first projects.

**Example Request:**

```json
GET /projects?status=approved&limit=10&start_ts=1704945600&end_ts=1704955600&last_key=MTcxNDk2Mzg4MT38dsjndjnsdjs3
```

**Success Response:**

  ```json
  HTTP/1.1 200 OK
  Content-Type: application/json

  {
    "ok": true,
    "data": {
        "projects": [
            {
                "id": 1,
                "title": "Project Title",
                "description": "Project Description",
                "image_urls": [
                    "https://media-donation-hub.ap-southeast-3.amazonaws.com/images/MTcxNDk2Mzg4MTE4OTU4MTIzMzYyNA.jpg",
                    "https://media-donation-hub.ap-southeast-3.amazonaws.com/images/MTcxNDk2Mzg4MT38dsjndjnsdjs3.jpg"
                ],
                "due_at": 1704945600,
                "target_amount": 1000000,
                "currency": "USD",
                "status": "approved",
                "requester": {
                    "id": 1,
                    "username": "johndoe",
                    "email": "johndoe@gmail.com"
                }
            },
            {
                "id": 2,
                "title": "Project Title 2",
                "description": "Project Description 2",
                "image_urls": [
                    "https://media-donation-hub.ap-southeast-3.amazonaws.com/images/MTcxNDk2Mzg4MTE4OTU4MTIzMzYyNA.jpg",
                    "https://media-donation-hub.ap-southeast-3.amazonaws.com/images/MTcxNDk2Mzg4MT38dsjndjnsdjs3.jpg"
                ],
                "due_at": 1704945600,
                "target_amount": 1000000,
                "currency": "USD",
                "status": "approved",
                "requester": {
                    "id": 2,
                    "username": "janedoe",
                    "email": "janedoe@gmail.com"
                }
            }
        ],
        "last_key": "MTcxNDk2Mzg4MT38dsjndjnsdjs3"
    },
    "ts": 1704954526
  }
  ```

**Error Response:**

No specific error response.

[Back to Top](#rest-api)

---

## Get Project

GET: `/projects/{project_id}`

This endpoint is used to get a project by ID.

**Example Request:**

```json
GET /projects/1
```

**Success Response:**

  ```json
  HTTP/1.1 200 OK
  Content-Type: application/json

  {
    "ok": true,
    "data": {
        "id": 1,
        "title": "Project Title",
        "description": "Project Description",
        "image_urls": [
            "https://media-donation-hub.ap-southeast-3.amazonaws.com/images/MTcxNDk2Mzg4MTE4OTU4MTIzMzYyNA.jpg",
            "https://media-donation-hub.ap-southeast-3.amazonaws.com/images/MTcxNDk2Mzg4MT38dsjndjnsdjs3.jpg"
        ],
        "due_at": 1704945600,
        "target_amount": 1000000,
        "collection_amount": 500000,
        "currency": "USD",
        "status": "approved",
        "requester": {
            "id": 1,
            "username": "johndoe",
            "email": "johndoe@gmail.com"
        }
    },
    "ts": 1704954526
  }
  ```

**Error Response:**

No specific error response.

[Back to Top](#rest-api)

---

## Donate to Project

POST: `/projects/{project_id}/donations`

This endpoint is used to donate to a project.

**Headers:**

- `Authorization` => The value is `Bearer {access_token}` with user role `donor`.

**Example Request:**

```json

POST /projects/1/donations
Authorization: Bearer {access_token}

{
  "amount": 100000,
  "currency": "USD", 
  "message": "Wish you all the best"
}

```

**Success Response:**

  ```json
  HTTP/1.1 200 OK
  Content-Type: application/json

  {
    "ok": true,
    "ts": 1704954526
  }
  ```

**Error Response:**

- Too much donation:

    ```json
    HTTP/1.1 409 Conflict
    Content-Type: application/json
  
    {
      "ok": false,
      "err": "ERR_TOO_MUCH_DONATION",
      "msg": "Donation amount must be less than target amount",
      "ts": 1704954526
    }
    ```

    This error returned if the donation collection amount is more than the target amount.


[Back to Top](#rest-api)

---

## List Project Donations

GET: `/projects/{project_id}/donations`

This endpoint is used to list all donations for a project.

**Query Params:**

- `limit` => Limit the number of donations returned. The value is integer.
- `last_key` => The last key of the previous request. The value is string. This is used for pagination.
  - If `last_key` is provided, the response will return the next donations after the last key.
  - If `last_key` is not provided, the response will return the first donations.

**Example Request:**

```json

GET /projects/1/donations?limit=10&last_key=MTcxNDk2Mzg

```

**Success Response:**

  ```json
  HTTP/1.1 200 OK
  Content-Type: application/json

  {
    "ok": true,
    "data": {
        "donations": [
            {
                "id": 1,
                "amount": 100000,
                "currency": "USD",
                "message": "Wish you all the best",
                "donor": {
                    "id": 1,
                    "username": "johndoe",
                },  
                "created_at": 1704954526
            },
            {
                "id": 2,
                "amount": 100000,
                "currency": "USD",
                "message": "Wish you all the best",
                "donor": {
                    "id": 2,
                    "username": "janedoe",
                },
                "created_at": 1704954526
            }
        ],
        "last_key": "MTcxNDk2Mzg"
    },
    "ts": 1704954526
  }
  ```

**Error Response:**

No specific error response.

[Back to Top](#rest-api)

---
