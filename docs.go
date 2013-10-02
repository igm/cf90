/*

Command Line

Application
  app.create    name=APPNAME instances=1 mem=256
  app.list
  app.push      name=APPNAME path=.
  app.start     name=APPNAME
  app.stop      name=APPNAME
  app.delete    name=APPNAME

  app.detail    name=APPNAME
  app.scale     name=APPNAME instances=10
  app.map       [NAME] [HOST] [DOMAIN]
  app.unmap     [NAME] [URL]
  app.logs      [NAME]
  app.stats     [NAME]
  app.crashlogs name=APPNAME

  app.restart   name=APPNAME
  app.rename    current=APPNAME new=APPNAME
  ?app.instance  name=APPNAME
  app.events    [NAME]
  app.file      [NAME] [PATH]
  app.files     [NAME] [PATH]

Service
  service.list

  service.info   [SERVICE]
  service.bind   [SERVICE] [APPNAME]
  service.unbind [SERVICE] [APPNAME]
  service.rename [OLD] [NEW]

Organization
  org.list
  org.use

  org.create
  org.delete
  org.rename

Space
  space.use

  space.list
  space.create
  space.delete
  space.rename

Route
  route.list
  route.delete

Domain
  domain.list

  domain.map
  domain.unmap

Target
  target.list
  target.use

  $ cf90 app.list
*/
package main
