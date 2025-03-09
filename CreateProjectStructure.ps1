# Set the root directory name
$rootDir = "project-root"

# Create project root directory
New-Item -ItemType Directory -Force -Path $rootDir

# Create subdirectories
New-Item -ItemType Directory -Force -Path "$rootDir\cmd"
New-Item -ItemType Directory -Force -Path "$rootDir\pkg\bot"
New-Item -ItemType Directory -Force -Path "$rootDir\pkg\api"
New-Item -ItemType Directory -Force -Path "$rootDir\pkg\database"
New-Item -ItemType Directory -Force -Path "$rootDir\pkg\vnstat"
New-Item -ItemType Directory -Force -Path "$rootDir\pkg\utils"
New-Item -ItemType Directory -Force -Path "$rootDir\scripts"
New-Item -ItemType Directory -Force -Path "$rootDir\docs"

# Create files in cmd
New-Item -ItemType File -Force -Path "$rootDir\cmd\main.go"
New-Item -ItemType File -Force -Path "$rootDir\cmd\agent.go"

# Create files in pkg/bot
New-Item -ItemType File -Force -Path "$rootDir\pkg\bot\bot.go"
New-Item -ItemType File -Force -Path "$rootDir\pkg\bot\commands.go"

# Create files in pkg/api
New-Item -ItemType File -Force -Path "$rootDir\pkg\api\server.go"
New-Item -ItemType File -Force -Path "$rootDir\pkg\api\auth.go"

# Create files in pkg/database
New-Item -ItemType File -Force -Path "$rootDir\pkg\database\db.go"
New-Item -ItemType File -Force -Path "$rootDir\pkg\database\models.go"
New-Item -ItemType File -Force -Path "$rootDir\pkg\database\vps.go"

# Create file in pkg/vnstat
New-Item -ItemType File -Force -Path "$rootDir\pkg\vnstat\vnstat.go"

# Create files in pkg/utils
New-Item -ItemType File -Force -Path "$rootDir\pkg\utils\config.go"
New-Item -ItemType File -Force -Path "$rootDir\pkg\utils\logger.go"

# Create agent installation script
New-Item -ItemType File -Force -Path "$rootDir\scripts\install_agent.sh"

# Create documentation file
New-Item -ItemType File -Force -Path "$rootDir\docs\README.md"

# Create Go module files
New-Item -ItemType File -Force -Path "$rootDir\go.mod"
New-Item -ItemType File -Force -Path "$rootDir\go.sum"

Write-Output "Project structure created successfully in the '$rootDir' directory."
