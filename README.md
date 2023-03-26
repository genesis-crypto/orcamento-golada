# __Calculadora de Orçamentos__
Este projeto consiste em uma aplicação de calculadora de orçamentos, que permite realizar cálculos de impostos e descontos sobre o valor dos produtos do orçamento em questão. O projeto foi desenvolvido como parte da disciplina de Design Patterns.

## __Requisitos Funcionais__
* Permitir a criação de orçamentos com produtos e quantidades
* Calcular o valor total do orçamento, incluindo impostos e descontos
* Permitir a aplicação de diferentes tipos de descontos no valor do orçamento
* Calcular o valor dos impostos sobre o valor do orçamento
* Permitir a visualização dos valores calculados de impostos e descontos

## __Requisitos Não-Funcionais__
* Utilizar Design Patterns para garantir uma arquitetura flexível e extensível
* Utilizar testes unitários para garantir a qualidade do código
* Utilizar boas práticas de programação, como SOLID e Clean Code

## __Estrutura de Arquivos__

O projeto está estruturado da seguinte forma:

```
orcamento-golada/
  ├── calculadora/
  │   ├── calculadora-cofins.decorator.go
  │   ├── calculadora-icms.decorator.go
  │   ├── calculadora-ipi.decorator.go
  │   ├── calculadora-ihit.decorator.go
  │   ├── calculadora-pis.decorator.go
  │   ├── calculadora.go  
  │   └── calculadora.interface.go
  ├── desconto/
  │   ├── desconto-quantia.chain.go
  │   ├── desconto-valor.go
  │   └── desconto.interface.go
  ├── orcamento/
  │   ├── aprovado.state.go
  │   ├── emaprovacao.state.go
  │   ├── finalizado.state.go
  │   ├── orcamento.go
  │   ├── orcamento.interface.go
  │   ├── reprovado.state.go
  │   └── state.go
  ├── go.mod
  ├── main.mod
  ├── orcamento-golada.exe
  └── README.md
```

A pasta calculadora contém a implementação da calculadora de impostos, com o uso do Decorator Pattern. A pasta orcamento contém a implementação do orçamento, com o uso do State Pattern. Já a pasta desconto contém a implementação do desconto, com o uso do Chain of Responsibility Pattern. Cada pasta contém um arquivo de implementação e um arquivo de testes.

## __Design Patterns utilizados__

### __State Pattern__
O State Pattern foi utilizado na implementação do orçamento do projeto. Esse padrão de projeto permitiu representar os diferentes estados de um orçamento de forma flexível e fácil de manter.

Foram criadas diferentes struct de estado, como `EmAprovacao`, `Aprovado`, `Reprovado` e `Finalizado`, que implementam uma interface comum IOrcamento. Cada struct de estado define o comportamento do objeto Orcamento para o estado correspondente.

Ao mudar o estado do objeto Orcamento, a referência para a struct de estado correspondente é atualizada, permitindo que o objeto mantenha seu estado atual e execute o comportamento apropriado.

### __Chain of Responsibility Pattern__
O Chain of Responsibility Pattern foi utilizado na implementação do desconto do orçamento. Esse padrão de projeto permitiu criar uma cadeia de responsabilidades para o cálculo de descontos em um orçamento, possibilitando que cada elo da cadeia pudesse decidir se aplicaria ou não um desconto e em qual valor.

Foram criadas diferentes structs de desconto, como `DescontoPorValor` e `DescontoPorQuantidade`, que implementam uma interface comum `IHDesconto`. Cada struct de desconto verifica se pode aplicar um desconto no objeto Orcamento e, caso contrário, repassa a solicitação para o próximo elo da cadeia.

Dessa forma, ao adicionar um novo tipo de desconto, basta criar uma nova classe e adicioná-la à cadeia de responsabilidades, sem afetar as outras structs já existentes. Além disso, o Chain of Responsibility Pattern tornou o cálculo de descontos mais flexível e extensível, permitindo que novos tipos de desconto possam ser adicionados à cadeia sem afetar o código existente.

### __Decorator Pattern__
O Decorator Pattern foi utilizado na implementação da calculadora de impostos do projeto. Esse padrão de projeto permitiu adicionar comportamentos de cálculo de impostos ao objeto Orcamento de forma flexível e dinâmica, sem alterar sua estrutura original.

Através da criação de "classes concretas ou struct concretas" de decoradores, como PIS, IPI, ICMS e Cofins, que implementam a interface ICalculadora e têm uma referência para o objeto Orcamento, foi possível criar uma "cadeia" de comportamentos que se encadeavam em tempo de execução.

Com o uso do Decorator Pattern, foi possível adicionar mais comportamentos de cálculo de impostos criando novas "classes ou structs" de decoradores, tornando a calculadora de impostos mais flexível e extensível.

## __Contato__
Pedro Cardozo - `p-cardozo@hotmail.com` ou `609455@univem.edu.br`