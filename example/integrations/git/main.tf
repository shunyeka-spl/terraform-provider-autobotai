resource "autobotai_git_integration" "git" {
    alias            = "MyGit"
    groups=["test"]
    host="aseweewewwew"
    token="cdscdscsc"
}

output "git"{
    value=autobotai_git_integration.git
}