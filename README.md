# Akiva
A ideia do akiva é ser um micro-serviço criado com o objetivo de captar comentários realizados pelo usuário na plataforma do Biblion e armazená-los no repositório de dados.

A estrutura inicial desse projeto envolve o estudo do RabbitMQ com Golang para processamento assíncrono dos comentários. Portanto, esta estrutura será simplificada futuramente.

## Estrutura Inicial
Uma vez que este projeto foi iniciado com o objetivo de entender melhor os conceitos do RabbitMQ e os diversos tipos de interação com este AMQP, a estrutura está organizada de modo geral em Consumers, Producers e MQ.

### Consumers
Neste package encontram-se os tipos de consumers que podem ser utilizados para tratar as mensagens processadas pelo RabbitMQ conforme hands-on do próprio Rabbit
- Receive: Go script simples que conecta ao Rabbit declara uma fila e aguarda mensagens simples para processamento.
- Subscriber: Go script que conecta ao Rabbit declara um exchange, uma fila e um tópico específico e aguarda publicações de _broadcast_ de mensagens.
- Worker: Go script que conecta ao Rabbit declara uma fila e trabalha em paralelo com multiplos workers para processamento de várias mensagens (_Task Queues_).

_A medida em que ainda estou estudando o produto, haverão outros tipos de consumers que serão desenvolvidos para estudo._

### Producers
Neste package encontram-se os tipos de producers que podem ser utilizados para publicar mensagens para o Rabbit.
- Basic: Go script que publica uma simples mensagem para o exchange default ("") e para uma fila específica.
- Publisher: Go script que publica uma mensagem para diversos subscribers, um broadcast de logs, conforme hands-on do RabbitMQ.
- Persistent: Go script que publica uma mensagem de modo persistent para ser trabalhada por vários workers (_Task Queues_)

### MQ
Neste package encontram-se scripts para facilitar a interação com o Rabbit encapsulando certas configurações redundantes para o trabalho com Exchanges, Queues, Publishs e Consumes.

## Executando Localmente
Em breve...

## Tecnologias
- Golang
- RabbitMQ

## Autor
Leonardo Felicissimo
