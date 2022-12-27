resource "autobotai_ms_teams_integration" "ms_teams" {
    alias            = "MyTeams"
    groups=["test"]
    webhook="webhook"
}

output "ms_teams"{
    value=autobotai_ms_teams_integration.ms_teams
}