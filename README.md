# Alexandria Identity API
![Go](https://github.com/alexandria-oss/identity-api/workflows/Go/badge.svg)

The Identity API is responsible of reading and some few writing operations related to users, special writing operations such as register, sign in and many more must be done using the AWS Cognito client API/SDK directly.

It uses gRPC and HTTP communication protocols to expose its API.

The Identity API v1 is using Amazon Web Service's Cognito as user pool and identity federation manager.

Alexandria is currently licensed under the MIT license.

## Endpoints
| Method              |     HTTP Mapping                        |  HTTP Request body  |  HTTP Response body |
|---------------------|:---------------------------------------:|:-------------------:|:-------------------:|
| **List**            |  GET /admin/user                        |   N/A               |   User* list        |
| **Get**             |  GET /user/{author-id}                  |   N/A               |   User*             |
| **Delete**          |  DELETE /admin/user/{user-id}           |   N/A               |   protobuf.empty/{} |
| **Restore/Active**  |  PATCH /admin/user/{user-id}            |   N/A               |   protobuf.empty/{} |
| **HardDelete**      |  DELETE /admin/user/{user-id}           |   N/A               |   protobuf.empty/{} |

### Accepted Queries
The list method accepts multiple queries to make data fetching easier for everyone.

The following fields are accepted by the service.
- page_token = string
- page_size = int32 (min. 1, max. 100)
- name = string
- email = string
- middle_name = string
- family_name = string
- locale = string
- show_disabled = boolean


## Contribution
Alexandria is an open-source project, that means everyone’s help is appreciated.

If you'd like to contribute, please look at the [Go Contribution Guidelines](https://github.com/alexandria-oss/alexandria/tree/master/docs/GO_CONTRIBUTION.md).

[Click here](https://github.com/alexandria-oss/alexandria/tree/master/docs) if you're looking for our docs about engineering, Alexandria API, etc.

## Maintenance
- Main maintainer: [maestre3d](https://github.com/maestre3d)
