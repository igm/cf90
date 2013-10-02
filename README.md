# CloudFoundry CLI - cf90

`cf90` is a command line tool to interact with cloud foundry instance. It is entirelly written in GO language.
The tool is still under development.

## Installation
`$ go get github.com/igm/cf90`

## Supported commands


```
Commands:
  help         - Shows this help message, use [COMMAND] for command parameters
  license      - show license information
Application
  app.addroute - Add route to application
  app.create   - Create new application
  app.delete   - Delete application
  app.list     - Show a list of apps
  app.push     - Push application
  app.start    - Start application
  app.stop     - Stop application
Domain
  domain.list  - Show a list of domains
Organization
  org.list     - Show all organizations
  org.use      - Set default organization
Route
  route.create - Create a route in current space
  route.delete - Delete route
  route.list   - Show all routes
Service
  service.list - Show a list of services
Space
  space.list   - Show all spaces in organization
  space.use    - Set default space
Target
  target.list  - Show known targets
  target.use   - Set current target
```
