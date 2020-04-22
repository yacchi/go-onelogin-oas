add
| .paths["/1/users"].get.responses["200"].content["application/json"].schema.required = ["data"]
| .paths["/1/users/{id}"].get.responses["200"].content["application/json"].schema.properties.data |= {"type": "array", "items": .}
| .paths["/1/users/{id}"].get.responses["200"].content["application/json"].schema.required = ["data"]
| .paths["/1/users/{id}/roles"].get.responses["200"].content["application/json"].schema.properties.data |= {"type": "array", "items": .}
| .paths["/1/users/{id}/roles"].get.responses["200"].content["application/json"].schema.required = ["data"]
| .paths["/1/users/custom_attributes"].get.responses["200"].content["application/json"].schema.properties.data |= {"type": "array", "items": .}
| .paths["/1/users/custom_attributes"].get.responses["200"].content["application/json"].schema.required = ["data"]
| .paths["/1/users/{id}/apps"].get.responses["200"].content["application/json"].schema.required = ["data"]
| .paths["/1/roles"].get.responses["200"].content["application/json"].schema.required = ["data"]
| .paths["/1/roles/{id}"].get.responses["200"].content["application/json"].schema.required = ["data"]
| .paths["/1/groups/{id}"].get.responses["200"].content["application/json"].schema.required = ["data"]
| .components.schemas.User.required |= .+ ["id"]
| .components.schemas.Role.required |= .+ ["id", "name"]
| .components.schemas.Group.required |= .+ ["id", "name"]
