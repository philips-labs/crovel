# Crovel
Crovel is not [Shovel](https://www.rabbitmq.com/shovel.html) but does almost the same, namely to forward messages from one RabbitMQ exchange to another. Only, it does this with a [crovel](https://www.urbandictionary.com/define.php?term=Crovel) instead of a shovel.

# Usage
Set the following environment variables to wire things up

| Environment | Description | Required | Default |
|-------------|-------------|----------|---------| 
| SRC_EXCHANGE | Name of source exchange | yes |  |
| SRC_ROUTING_KEY | Source routing key | no | `#` |
| SRC_EXCHANGE_TYPE | Source exchange type | no | `topic` |
| DEST_EXCHANGE | Destination exchange | yes | |
| DEST_EXCHANGE_TYPE | Destination exchange type | no | `topic` |

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
    SRC_EXCHANGE: foo
    DEST_EXCHANGE: bar
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
    SRC_EXCHANGE  = "foo"
    DEST_EXCHANGE = "bar"
  }
}

```

# Contact / Getting help
andy.lo-a-foe@philips.com

# License
License is MIT
