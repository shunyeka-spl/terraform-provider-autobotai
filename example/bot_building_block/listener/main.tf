resource "autobotai_listener" "listener" {
        description = "Testing"
        name = "testing"
        webhook_source_integration_type = "conformity"
}

output "listenerRead"{
        value=autobotai_listener.listener
}
