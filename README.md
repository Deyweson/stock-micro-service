# Micro-service de Estoque

Uma api para simular um micro serviço de controle de estoque, que o objetivo é poder ser consumido por varias lojas(da mesma franquia), para consultar e usar um mesmo estoque unificado com atualizações em tempo real.

## Banco de dados

|Campo|Valor|
| - | - |
| id | int |
| shopId | varchar(255) |
| prodId | varchar(255) |
| operation_type | "in" or "out" |
| amount | float |
|description| text |
|createdAt| timestamp |

## Rotas 

`GET /v1/stocks`
Rota de listagem das movimentações de estoque
