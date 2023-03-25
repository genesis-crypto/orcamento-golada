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
  │   ├── calculadora-pis.decorator.go
  │   ├── calculadora.go  
  │   └── calculadora.interface.go
  ├── desconto/
  ├── orcamento/
  │   ├── orcamento.go
  │   └── orcamento.interface.go
  ├── go.mod
  ├── main.mod
  └── README.md
```

A pasta calculadora contém a implementação da calculadora de impostos, com o uso do Decorator Pattern. A pasta orcamento contém a implementação do orçamento, com o uso do State Pattern. Já a pasta desconto contém a implementação do desconto, com o uso do Chain of Responsibility Pattern. Cada pasta contém um arquivo de implementação e um arquivo de testes.

## __Design Patterns utilizados__

### __State Pattern__
### __Chain of Responsibility Pattern__
### __Decorator Pattern__
O Decorator Pattern foi utilizado na implementação da calculadora de impostos do projeto. Esse padrão de projeto permitiu adicionar comportamentos de cálculo de impostos ao objeto Orcamento de forma flexível e dinâmica, sem alterar sua estrutura original.

Através da criação de "classes concretas ou struct concretas" de decoradores, como PIS, IPI, ICMS e Cofins, que implementam a interface ICalculadora e têm uma referência para o objeto Orcamento, foi possível criar uma "cadeia" de comportamentos que se encadeavam em tempo de execução.

Com o uso do Decorator Pattern, foi possível adicionar mais comportamentos de cálculo de impostos criando novas "classes ou structs" de decoradores, tornando a calculadora de impostos mais flexível e extensível.

## __Contato__
Pedro Cardozo - `p-cardozo@hotmail.com` ou `609455@univem.edu.br`