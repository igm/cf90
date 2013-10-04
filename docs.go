/*

Command Line tool to interact with CloudFoundry instance.

Application
  app.create    name=APPNAME instances=1 mem=256
  app.list
  app.push      name=APPNAME path=.
  app.start     name=APPNAME
  app.stop      name=APPNAME
  app.delete    name=APPNAME
  app.map       name=APPNAM host=HOST domain=DOMAIN
  app.unmap
  app.detail

  TODO:
  app.scale
  app.logs
  app.stats
  app.crashlogs
  app.restart
  app.rename
  ?app.instance
  app.events
  app.file
  app.files

Service
  service.list

  TODO:
  service.info
  service.bind
  service.unbind
  service.rename

Organization
  org.list
  org.use   org=ORGNAME

  TODO:
  org.create
  org.delete
  org.rename

Space
  space.use  space=SPACENAME
  space.list

  TODO:
  space.create
  space.delete
  space.rename

Route
  route.list
  route.create  host=HOST domain=DOMAIN
  route.delete  host=HOST domain=DOMAIN

Domain
  domain.list

  TODO:
  domain.map
  domain.unmap

Target
  target.add
  target.info
  target.list
  target.login
  target.logout
  target.rm
  target.use


*/
package main
