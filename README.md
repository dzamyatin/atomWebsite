
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
- <span style="color:green">DONE</span> registration
- <span style="color:green">DONE</span> login
- <span style="color:lightgreen">MVP</span> email-confirm
  - <span style="color:green">DONE</span> yandex smtp
  - google smtp
  - mail smtp
- <span style="color:yellow">IN PROGRESS</span> phone-confirm
  - usecasees: confirmphone, sendphoneconfirmation
  - front pages to handle confirmations
  - <span style="color:green">DONE</span> telegram
    - <span style="color:green">DONE</span> sender to store (chat lint, phone) to send msg
      - <span style="color:green">DONE</span>to store
      - <span style="color:green">DONE</span> to send
        - <span style="color:green">DONE</span> counter
        - <span style="color:green">DONE</span> sequentially send one to another accross all senders by couner
  - whatsapp
  - vk
- remember password
  - <span style="color:green">DONE</span> for email
  - <span style="color:orange">TEST REQ AFTER EMAIL CONFIRMATION</span> for phone
- cart
- place-order
- key-list to obtain
- payment
  - check telegram payment https://core.telegram.org/bots/api#payments 
  - reserch which payment method are better (yandexpay)
  - payment implementation
  - reserch which crypt method are better
  - crypt payment method
- personal cabinet
- download app
- <span style="color:green">DONE</span> registration
  - <span style="color:green">DONE</span> email verification
    - <span style="color:lightblue">WONT DO</span> smtp server (decided to use yandex.ru smtp, google smtp etc. instead)
    ```
    https://dev.to/wneessen/sending-mails-in-go-the-easy-way-1lm7
    
    some info? https://habr.com/ru/articles/564750/
    
    smtp server: https://habr.com/ru/articles/673700/
    
    msmtp as alternative to sendmail https://askubuntu.com/questions/1363136/configuring-sendmail-on-ubuntu-20-04
    
    try as smtp https://smtp.bz/
    
    https://poste.io/doc/getting-started
    
    https://docs.mailcow.email/getstarted/prerequisite-dns/#dkim-spf-and-dmarc
    ```
    
- nginx for static
- kubernetes deploy
  - statefulset for postgress
  - deploy for app
  - volume

## backlog
- yandex metrics
- test cover
- kafka instead of postgress
- websockets (notify for some events) discovery a way how to handle communication
- code generation of decorators (find a package) to add metrics
- helm
- metrics:
  - db metrics should be done through context and some decarator to handle it
  - ?begin transaction should return "metric" tx (for metric decorator)
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
- vue js admin (https://demos.creative-tim.com/vue-argon-design-system/documentation/?_ga=2.135621303.708551717.1744408739-559679860.1744408739)

## dont forgot
to add traffic limit to jwt key