resource "autobotai_inventory_schedule" "inventory" {
        integration_id="qwedsawqdwq_sdcsdcdsdsc"
        integration_type="azure"
        cron_expression="0 0 * * *"
}
output "inventoryRead"{
    value=autobotai_inventory_schedule.inventory
}

