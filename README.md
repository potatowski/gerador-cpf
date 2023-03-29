# Gerador de CPF

Durante o trabalho me desafiaram a montar um gerador de cpf, e como o calculo é simples aqui está.

Pra quem não sabe o calculo é assim
Vamos supor que o cpf é 475.426.123-28

Passo 1 - Somar todos os números até o digito verificador multiplicando cada por 10 e reduzindo sucessivamente.
Logo seria, (10x4)+(9x7)+(8x5)+(7x4)+(6x2)+(5x6)+(4x1)+(3x2)+(2x3) = x

Passo 2 - Após somar e multiplicar todos, temeros o valor "X", onde será dividido por 11 e pegaremos o resto da divisão, caso o resto seja, 0 ou 1, o primeiro digito será 0, caso não iremos subtrair esse valor de 11 e o resultado será primeiro digito do verificador.
Logo seria, 229%11 = 9, 11 - 9 = 2

Passo 3 - Repetir o mesmo passo 1 e 2, só que agora iniciando por 11 e adicionando no calculo valor obtido no passo 2.
Logo seria, (11x4)+(10x7)+(9x5)+(8x4)+(7x2)+(6x6)+(5x1)+(4x2)+(3x3)+(2x2) = 267, 267%11 = 3, 11 - 3 = 8

Passo 4 - Verificar se o digitos verificadores do cpf é igual ao valor obtido

Ah e vale comentar que o ante-penultimo digito do CPF (xxx.xxx.xx8-xx) é o estado que ele foi emitido conforme:
- 0 = Rio Grande do Sul
- 1 = Distrito Federal, Goiás, Mato Grosso, Mato Grosso do Sul e Tocantins
- 2 = Amazonas, Pará, Roraima, Amapá, Acre e Rondônia
- 3 = Ceará, Maranhão e Piauí
- 4 = Paraíba, Pernambuco, Alagoas e Rio Grande do Norte
- 5 = Bahia e Sergipe
- 6 = Minas Gerais
- 7 = Rio de Janeiro e Espírito Santo
- 8 = São Paulo
- 9 = Paraná e Santa Catarina
