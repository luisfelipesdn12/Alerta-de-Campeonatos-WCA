#IMPORTAÇÕES:
#pip install necessário:
import requests #para Web Scrapping
from bs4 import BeautifulSoup #para Web Scrapping
from unidecode import unidecode 
import gspread #manipula google sheets
from oauth2client.service_account import ServiceAccountCredentials #autentica o acesso
#bibliotecas nativas:
import smtplib #para o envio de e-mails
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText
#as duas acima são para mandar emails em UTF-8
from time import sleep
from datetime import datetime, date

#FUNCTIONS:
def retornaWorksheets():
  scope = [
    "https://www.googleapis.com/auth/spreadsheets", "https://www.googleapis.com/auth/drive.file", "https://www.googleapis.com/auth/drive"
    ]

  creds = ServiceAccountCredentials.from_json_keyfile_name('credencial_google.json', scope) #capitura credenciais de um arquivo .json

  client = gspread.authorize(creds) #usa essas credenciais

  sheet = client.open('Dados') #abre o arquivo

  worksheets = dict()
  worksheets['dest_sheet'] = sheet.worksheet('Destinatarios')
  worksheets['cred_sheet'] = sheet.worksheet('Credenciais')

  return(worksheets) #retorna um dicionário com as duas subplanilhas

worksheets = retornaWorksheets()

def retornaDados(): #retorn credenciais e lista de dests
  #usa o dicionário criado por retornaWorksheets()
  dest_sheet = worksheets['dest_sheet'] 
  cred_sheet = worksheets['cred_sheet']

  dados = dict()

  dados['list_of_dicts_dest_sheet'] = dest_sheet.get_all_records() #captura uma lista com dicionarios contendo infoemações de cada remetente

  list_of_dicts_cred_sheet = cred_sheet.get_all_records() #capitura as credenciais para envio de emails

  dados['credentials'] = {
    'email_remetente':f'{list_of_dicts_cred_sheet[0]["Email"]}',
    'senha':f'{list_of_dicts_cred_sheet[0]["Senha"]}'
    }

  return(dados) #retorna um dicionario com as duas infoemações

dados = retornaDados()

def retornaCompsFuturas(cidade):
  URL = f"https://www.worldcubeassociation.org/competitions?utf8=%E2%9C%93&event_ids%5B%5D=333&event_ids%5B%5D=222&region=Brazil&search={cidade.replace(' ', '+')}&state=present&year=all+years&from_date=&to_date=&delegate=&display=list" #fonte dos dados

  #CONFIGURAÇõES:
  headers = {"User-Agent" : 'YOUR_USER_AGENT'}
  page = requests.get(URL, headers = headers)
  soup = BeautifulSoup(page.content,  'html.parser')

  scrap_puro = soup.find(class_="list-group-item") #consulta número atual de competições

  #captura o número, o isolando do resto
  if scrap_puro != None:
    numero_isolado = int((str(scrap_puro.get_text()))[-2])
  #retorna 0 em vez de None
  else: numero_isolado = 0

  return(numero_isolado)

def manda_email(dest, competicoes_num_atual, data_de_verificacao_atual):
  credenciais_email = dados['credentials'] #pega as credenciais para o envio de emails

  #configurações:
  server = smtplib.SMTP('smtp.gmail.com', 587)
  server.ehlo()
  server.starttls()
  server.ehlo()

  #login:
  server.login(str(credenciais_email['email_remetente']), str(credenciais_email['senha']))

  #conteúdo dos emails
  assunto = f'{dest["Nome"]}! Temos uma atualização! - {date.today()}'
  corpo = f'''
Olá {dest["Nome"]}!
Este email é enviado automaticamente e tem informações sobre competições nas modalidades 2x2 e 3x3 na cidade {dest["Cidade"]}.

Há uma alteração de {dest["N de competições"]} competições futuras, para {competicoes_num_atual}.

##Informações Gerais:

Competições futuras: {dest["N de competições"]}
Data de verificação: {dest["Data de verificação"]}

Competições futuras: {competicoes_num_atual}
Data de verificação: {data_de_verificacao_atual}

##Veja aqui: https://www.worldcubeassociation.org/competitions?utf8=%E2%9C%93&event_ids%5B%5D=333&event_ids%5B%5D=222&region=Brazil&search={dest["Cidade"].replace(' ', '+')}&state=present&year=all+years&from_date=&to_date=&delegate=&display=list
  '''

  msg = MIMEMultipart("alternative")
  msg["Subject"] = f"{assunto}"
  msg.attach( MIMEText(f"{corpo}", "plain", "utf-8" ) )
  msg = msg.as_string().encode('ascii')

  server.sendmail(
    str(credenciais_email['email_remetente']),
    str(dest['Email']),
    msg
  )
  print(f'O email foi mandado para {dest["Email"]}! {str(datetime.now())} \n---\n')

#MAIN FUNCTION:
def main():
  lista_de_dests = dados['list_of_dicts_dest_sheet'] #pega uma lista com dicionarios contendo infoemações de cada remetente

  print("\nDADOS CAPTURADOS\n")

  c = 1 #variavel de controle das repetições abaixo, para saber em qual destinatario fazer alterações na planilha

  #verifica para cada destinatario
  for dest in lista_de_dests:
    competicoes_num_atual = retornaCompsFuturas(cidade = dest['Cidade']) #web scrap das competiÇões futuras

    data_de_verificacao_atual = datetime.now() #datetime da verificação

    #se houver alteração desde a última verificação, manda um email
    if dest['N de competições'] != '' and competicoes_num_atual != dest['N de competições']:
      try: manda_email(dest, competicoes_num_atual, data_de_verificacao_atual)
      except: print(f'\n ERRO AO ENVIO DE EMAIL: {dest}\n')

    #atualiza dados na planilha:
    retornaWorksheets()['dest_sheet'].update_acell(f'E{c+1}', f'{competicoes_num_atual}') #n de comps

    retornaWorksheets()['dest_sheet'].update_acell(f'F{c+1}', f'{data_de_verificacao_atual}') #data de verificação

    print(f'\n VERIFICAÇÃO FEITA: {dest} \n')

    c += 1


#MAIN LOOP:
main()
