# Alerta-de-Campeonatos-WCA
Um script que manda um e-mail quando há um campeonato novo na WCA.

## Ideia:
>"A World Cube Association regula competicões de quebra-cabeças mecânicos que são operados girando-se os lados, comumente chamados de "twisty puzzles". O mais famoso deles é o "Rubik's Cube" (Cubo Mágico ou Cubo de Rubik), inventado pelo professor Rubik, da Hungria. Alguns destes quebra-cabeças são eventos oficiais da WCA.
À medida que a WCA evoluiu ao longo da última década, mais de 100.000 pessoas já participaram de nossas competições."
>- Fonte: "[Quem somos nós](https://www.worldcubeassociation.org/about)"  acessado em 23 de Fevereiro de 2020.

Eu e meus amigos temos como *hobbie* o *speedcubing*, simplificadamente: montar cubo-mágico e outros quebra-cabeças no menor tempo possível.  
Existem campeonatos oficiais por todo o Mundo, organizados pela Organização Mundial de Cubo Mágico (WCA), supracitada.
><img src="https://www.cps.sp.gov.br/wp-content/uploads/sites/1/2019/08/Etec-Jacare%C3%AD-4%C2%BA-campeonato-mundial-do-cubo.jpg" width="600">

Nós participamos deles, e é bem comum consultarmos o [site da WCA](https://www.worldcubeassociation.org/competitions) em buscas de competições por perto. Às vezes, entrávamos algumas vezes na semana, e nada; às vezes, esquecíamos de entrar e perdíamos um campeonato tão aguardado. 
Para resolver esse problema, tive a ideia de fazer um script que verificasse o site periodicamente e nos notificasse quanto identificasse uma competição por perto que poderia ser de nosso interesse.  

## Execução:
Para executar esse projeto, estudei sobre Web Scrapping e envios de e-mails em Python. Utilizei bibliotecas como `requests` e `BeautifulSoup` para Web Scrapping e `smtplib` para o envio de e-mails.

O funcionamento do código é simples:

- Conecta à uma planilha no Google Sheets, por meio de um `ARQUIVO.json`;
- Coleta os dados de `Destinatarios` atualizados;
><img src="https://raw.githubusercontent.com/luisfelipesdn12-email/Alerta-de-Campeonatos-WCA/master/demo_images/Sheet%20Dests%20Print.JPG" width="700">
- Coleta os dados de `Credenciais` atualizados, onde são armazenados as credenciais para o envio de emails;
><img src="https://raw.githubusercontent.com/luisfelipesdn12-email/Alerta-de-Campeonatos-WCA/master/demo_images/Sheet%20Creds%20Print.JPG" width="700">
- Verifica para cada destinatario, se houver alterações em `N de competições` desde a última verificação: manda um email.
- Atualiza os dados na planilha.
- Aguarda um tempo e repete o processo.

