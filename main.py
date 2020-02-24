#IMPORTAÇÕES:
#pip install necessário
import requests
from bs4 import BeautifulSoup
from unidecode import unidecode
#bibliotecas nativas
import smtplib
from time import sleep
import datetime
#dados
from dados import competicoes_num, email_dos_destinatarios

pesquisa = 'São Paulo' #equivalente à pesquisa no site, no caso, usado para delimitar a cidade
URL = f"https://www.worldcubeassociation.org/competitions?utf8=%E2%9C%93&event_ids%5B%5D=333&event_ids%5B%5D=222&region=Brazil&search={pesquisa.replace(' ', '+')}&state=present&year=all+years&from_date=&to_date=&delegate=&display=list" #fonte dos dados

#CONFIGURAÇõES:
headers = {"User-Agent" : 'YOUR_USER_AGENT'}
page = requests.get(URL, headers = headers)
soup = BeautifulSoup(page.content,  'html.parser')

#FUNÇÔES:
def manda_email(): #envia emails à todos os destinatários:
  #configurações:
  server = smtplib.SMTP('smtp.gmail.com', 587)
  server.ehlo()
  server.starttls()
  server.ehlo()

  #login:
  server.login('luisfelipesdn12@gmail.com', 'TOKEN_OR_PASSWORD')

  #conteúdo dos emails
  assunto = 'Teste do Web Scrapper WCA 1.0'
  corpo = f'O numero de campeonatos no Brasil,em {pesquisa} e em que ha as modalidades 2x2 e 3x3 -> {competicoes_num_atual} \nVerificacao feita em: {str(datetime.datetime.now())} \n\nVeja aqui: {URL}'
  msg = unidecode(f'Subject: {assunto}\n\n{corpo}') #(unidecode para evitar erros pelo uso de acentos)

  for email in email_dos_destinatarios:
    server.sendmail(
      'luisfelipesdn12@gmail.com',
      str(email),
      msg
    )
    print(f"O email foi mandado para {email}! \n {str(datetime.datetime.now())} \n---\n")

  server.quit()

#LOOPING DE VERIFICAÇÕES:
while True:
  competicoes_num_atual = int((str(soup.find(class_="list-group-item").get_text()))[-2]) #consulta número atual de competições

  #se o número coletado diferir do último dado coletado E o atual for maior; ou seja: se tiver aumentado o número de competições:
  if competicoes_num_atual != competicoes_num and competicoes_num_atual > competicoes_num:
    print(f'\n---\nVerificação feita. Emails enviados. \n\nDado antigo: {competicoes_num} competições futuras; \nDado atual: {competicoes_num_atual} competições futuras. \n\n{str(datetime.datetime.now())} \n---\n') #imprime na tela os dados
    manda_email() #aciona a função de mandar emails

  #caso contrário só imprime os dados:
  else: print(f'\n---\nVerificação feita. Emails não enviados. \n\nDado antigo: {competicoes_num} competições futuras; \nDado atual: {competicoes_num_atual} competições futuras. \n\n{str(datetime.datetime.now())} \n---\n')

  #atualiza os dados no arquivo dados.py:
  file = open('dados.py', 'w')
  file.writelines(f'competicoes_num = {competicoes_num_atual}\n\nemail_dos_destinatarios = {email_dos_destinatarios}')
  file.close()
  
  competicoes_num = competicoes_num_atual #atualiza os dados no script

  sleep(60) #tempo (segundos) entre uma verificação e outra

