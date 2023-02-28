resource "autobotai_google_chat_integration" "google_chat" {
    alias            = "ProjectTesting"
    groups=["test"]
    webhook="webhook_url"
}

output "google_chat"{
    value=autobotai_google_chat_integration.google_chat
}