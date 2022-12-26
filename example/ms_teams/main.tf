resource "autobot_ms_teams_integration" "ms_teams" {
    alias            = "MyTeams"
    groups=["test"]
    webhook="webhook"
}

output "ms_teams"{
    value=autobot_ms_teams_integration.ms_teams
}