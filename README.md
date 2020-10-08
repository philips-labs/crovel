# Crovel
Crovel is not [Shovel](https://www.rabbitmq.com/shovel.html) but does almost the same, namely to forward messages from one RabbitMQ exchange to another. Only, it does this with a [crovel](https://www.urbandictionary.com/define.php?term=Crovel) instead of a shovel.

# Usage
Set the following environment variables to wire things up

| Environment | Description | Required | Default |
|-------------|-------------|----------|---------| 
| CROVEL_SRC_EXCHANGE | Name of source exchange | yes |  |
| CROVEL_SRC_ROUTING_KEY | Source routing key | no | `#` |
| CROVEL_SRC_EXCHANGE_TYPE | Source exchange type | no | `topic` |
| CROVEL_DEST_EXCHANGE | Destination exchange | yes | |
| CROVEL_DEST_EXCHANGE_TYPE | Destination exchange type | no | `topic` |

# Deploy
Deploy to Cloud Foundry and bind the `RabbitMQ` service to the app

## Example manifest
```yaml
---
applications:
- name: crovel
  docker:
    image: loafoe/crovel:latest
  instances: 1
  memory: 64M
  disk_quota: 128M
  health-check-type: process
  env:
    CROVEL_SRC_EXCHANGE: foo
    CROVEL_DEST_EXCHANGE: bar
  services:
  - rabbitmq
```

## Example Terraform
```hcl
resource "cloudfoundry_app" "crovel" {
  name         = "crovel"
  space        = data.cloudfoundry_space.space.id
  memory       = 64
  disk_quota   = 128
  docker_image = "loafoe/crovel:latest"
  health_check_type = "process"
  environment = {
    CROVEL_SRC_EXCHANGE  = "foo"
    CROVEL_DEST_EXCHANGE = "bar"
  }
}

```

# Ideas
- Support forwarding to different RabbitMQ cluster
- Support mulitple exchange forwards per instance
- Configurable durability

# Contact / Getting help
andy.lo-a-foe@philips.com

# License
License is MIT
