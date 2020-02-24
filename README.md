# Alerta-de-Campeonatos-WCA
Um script que manda um e-mail quando há um campeonato novo na WCA.

## Ideia:
>"A World Cube Association regula competicões de quebra-cabeças mecânicos que são operados girando-se os lados, comumente chamados de "twisty puzzles". O mais famoso deles é o "Rubik's Cube" (Cubo Mágico ou Cubo de Rubik), inventado pelo professor Rubik, da Hungria. Alguns destes quebra-cabeças são eventos oficiais da WCA.
À medida que a WCA evoluiu ao longo da última década, mais de 100.000 pessoas já participaram de nossas competições."
>- Fonte: "[Quem somos nós](https://www.worldcubeassociation.org/about)"  acessado em 23 de Fevereiro de 2020.

Eu e meus amigos temos como *hobbie* o *speedcubing*, simplificadamente: montar cubo-mágico e outros quebra-cabeças no menor tempo possível.  
Existem campeonatos oficiais por todo o Mundo, organizados pela Organização Mundial de Cubo Mágico (WCA), supracitada. 
![Imagem de um campeonato oficial.](https://www.cps.sp.gov.br/wp-content/uploads/sites/1/2019/08/Etec-Jacare%C3%AD-4%C2%BA-campeonato-mundial-do-cubo.jpg)
Nós participamos deles, e é bem comum consultarmos o [site da WCA](https://www.worldcubeassociation.org/competitions) em buscas de competições por perto. Às vezes, entrávamos algumas vezes na semana, e nada; às vezes, esquecíamos de entrar e perdíamos um campeonato tão aguardado. 
Para resolver esse problema, tive a ideia de fazer um script que verificasse o site periodicamente e nos notificasse quanto identificasse uma competição por perto que poderia ser de nosso interesse.  

## Execução:
Para executar esse projeto, estudei sobre Web Scrapping e envios de e-mails em Python. Utilizei bibliotecas como `requests` e `BeautifulSoup` para Web Scrapping e `smtplib` para o envio de e-mails.
