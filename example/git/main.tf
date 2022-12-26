resource "autobot_git_integration" "git" {
    alias            = "MyGit"
    groups=["test"]
    host="aseweewewwew"
    token="cdscdscsc"
}

output "git"{
    value=autobot_git_integration.git
}