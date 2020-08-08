# Alerta-de-Campeonatos-WCA
Um script que manda um e-mail quando há um campeonato novo na WCA.
A script witch send an e-mail when there's a new WCA competition. 

![Header](https://raw.githubusercontent.com/luisfelipesdn12/Alerta-de-Campeonatos-WCA/golang/images/Email%20Header%20Portuguese.png)

## Idéia:
>"A World Cube Association regula competicões de quebra-cabeças mecânicos que são operados girando-se os lados, comumente chamados de "twisty puzzles". O mais famoso deles é o "Rubik's Cube" (Cubo Mágico ou Cubo de Rubik), inventado pelo professor Rubik, da Hungria. Alguns destes quebra-cabeças são eventos oficiais da WCA.
À medida que a WCA evoluiu ao longo da última década, mais de 100.000 pessoas já participaram de nossas competições."
>- Fonte: "[Quem somos nós](https://www.worldcubeassociation.org/about)"  acessado em 08 de Agosto de 2020.

Eu e meus amigos temos como *hobbie* o *speedcubing*, simplificadamente: montar cubo-mágico e outros quebra-cabeças no menor tempo possível.  
Existem campeonatos oficiais por todo o mundo, organizados pela Organização Mundial de Cubo Mágico (WCA), supracitada.
><img src="https://www.cps.sp.gov.br/wp-content/uploads/sites/1/2019/08/Etec-Jacare%C3%AD-4%C2%BA-campeonato-mundial-do-cubo.jpg" width="600">

Nós participamos deles, e é bem comum consultarmos o [site da WCA](https://www.worldcubeassociation.org/competitions) em buscas de competições por perto. Às vezes, entrávamos algumas vezes na semana, e nada; às vezes, esquecíamos de entrar e perdíamos um campeonato tão aguardado. 
Para resolver esse problema, tive a ideia de fazer um script que verificasse o site periodicamente e nos notificasse quanto identificasse uma competição por perto que poderia ser de nosso interesse.

## Uso:
Se inscreva, inserindo seu nome, e-mail, idioma e a cidade que deseja ser notificado preenchendo o seguinte formulário:

[**Inscreva-se**](https://forms.gle/K6vW3YVAYp4d6nb97)

## Execução:
Para executar o projeto, eu inicialmente usei Python com as bibliotecas `requests` e `BeaultifulSoup` para web scrapping no site em sí e `smtplib` para o envio de e-mails.
Mas eu fiz uma migração para a linguagem Go, com a API da WCA no lugar do front-end do site. Eu estudei modulação do código em pacotes locais, o uso de bibliotecas como `spreadsheet` para conectar-se com a API do Google Planilhas e o `gomail` para enviar as notificações.

O código funciona assim:

- Busca os dados da planilha na minha conta do Google (dados dos destinatários e as credenciais para o envio dos e-mails);
- Verifica as competições futuras na cidade de cada destinatário;
- Atualiza na planilha;
- Compara a verificação atual com a última;
- Manda um e-mail caso esses números sejam diferentes;
> O log de todo esse processo é armazenado num arquivo `main.log`.

Na minha conta do Google, a planilha está organizada nesse formato (com os nomes em inglês):

### Planilha 1 ("Destinatários"):
> Os dados providos pelo formulário e as verificações anteriores.

|  Formulário preenchido em  | Nome |   E-mail   |   Cidade   |  Idioma  |  Competições Futuras  | Última Verificação  |
| -------------------------- | ---- | ---------- | ---------- | -------- | --------------------- | ------------------- |
| 00/00/0000 00:00:00        | anne | anne@e.com | New Jersey | English  | 7                     | 0000-00-00 00:00:00 |
| ...                        | ...  | ...        | ...        | ...      | ...                   | ...                 |

### Planilha 2 ("Betas"):
> Meus amigos que concordaram em ser *beta testers*. Quando o código está em desenvolvimento, ele roda aqui primeiro.

|          -           | Nome |   E-mail   |   Cidade   |   Idioma   |  Competições Futuras  | Última Verificação  |
| -------------------- | ---- | ---------- | ---------- | ---------- | --------------------- | ------------------- |
|          -           | tagu | tagu@u.com | São Paulo  | Português  | 2                     | 0000-00-00 00:00:00 |
| ...                  | ...  | ...        | ...        | ...        | ...                   | ...                 |

### Planilha 3 ("Credenciais"):
> O e-mail e senha da conta que envia os e-mails.

|        E-mail        |   Senha    |
| -------------------- | ---------- |
| minhaconta@gmail.com | minh453nh4 |

