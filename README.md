
# TODO

## MVP

- <span style="color:green">DONE</span> grpc server
- <span style="color:green">DONE</span> wire (DI) by Gooogl
- <span style="color:green">DONE</span> process manager to gracefully shutdown app
- <span style="color:green">DONE</span> viper config
- <span style="color:green">DONE</span> add database
- <span style="color:green">DONE</span> main.go common entrypoint of apps
- <span style="color:green">DONE</span> command migration
- <span style="color:green">DONE</span> metric health
- <span style="color:green">DONE</span> sqlx
- <span style="color:green">DONE</span> registration endpoint
- <span style="color:green">DONE</span> login endpoint
- <span style="color:green">DONE</span> bus (command handler pattern, cqrs implementation) 
  - <span style="color:green">DONE</span> event bus
  - <span style="color:green">DONE</span> command dispatcher
  - <span style="color:green">DONE</span> memory provider
  - <span style="color:green">DONE</span> postgress provider
    - <span style="color:green">DONE</span> pg dispatch
    - <span style="color:green">DONE</span> pg handler
- vue js admin (https://demos.creative-tim.com/vue-argon-design-system/documentation/?_ga=2.135621303.708551717.1744408739-559679860.1744408739)
  - login
  - payment
  - personal cabinet
  - download app
- nginx for static
- begin transaction should return "metric" tx
- registration
  - email verification
    - smtp server
- grpc client
- reserch which payment method are better (yandexpay)
- payment implementation
- reserch which crypt method are better
- crypt payment method
- kubernetes deploy
  - statefulset for postgress
  - deploy for app
  - volume

## backlog
- test cover
- kafka instead of postgress
- code generation of decorators (find a package) to add metrics
- helm
- phone verifiction
  - telegram verify
  - whatsup|vk / rollback messager
- metric prometheus (see: https://habr.com/ru/companies/otus/articles/769806/)
  - prometeus
  - grafana
    - dashboard
      0) GC metric (mem, cpu, GC time, heap size)
      1) all request timing
      2) average request time
      3) all db query timing (should looks like bukets .01, .1, ... 10) To IDatabase
      4) average query time

## interesting
- check http to grpc proxy (see: https://habr.com/ru/articles/658769/)
- http server (which point?)