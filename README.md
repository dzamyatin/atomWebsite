
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
- vue js admin
- nginx for static
- begin transaction should return "metric" tx
- cqrs implementation
  - event bus
  - command dispatcher
- <span style="color:yellow">IN PROGRESS</span> bus (command handler pattern) database implementation
- http server
- grpc client
- payment 
- crypt payment method
- metric prometheus (see: https://habr.com/ru/companies/otus/articles/769806/)
  - prometeus
  - grafana
    - dashboard
      0) GC metric (mem, cpu, GC time, heap size)
      1) all request timing
      2) average request time
      3) all db query timing (should looks like bukets .01, .1, ... 10) To IDatabase
      4) average query time
- check http to grpc proxy (see: https://habr.com/ru/articles/658769/)

## backlog
- code generation of decorators (find a package) to add metrics
- kubernetes
- helm
