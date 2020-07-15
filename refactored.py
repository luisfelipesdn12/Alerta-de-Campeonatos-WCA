# IMPORTS:
## !pip install required:
### --> web scrapping:
import requests 
from bs4 import BeautifulSoup
from unidecode import unidecode
### --> google sheets manipulation:
import gspread
from oauth2client.service_account import ServiceAccountCredentials #autentica o acesso
## native libraries:
### --> send emails:
import smtplib
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText
### --> date/time manipulation:
from time import sleep
from datetime import datetime, date

json_credentials_file_name = 'alerta-de-campeonatos-wca-3cc3e6e36451.json'

def returnWorksheets():
    """
    Do the authentication on Google Sheets API and 
    returns a dictionary with data from recipients
    and the credentials used to send emails.
    """    

    scope = [
        "https://www.googleapis.com/auth/spreadsheets",
        "https://www.googleapis.com/auth/drive.file",
        "https://www.googleapis.com/auth/drive"
    ]

    # get credentials from a json file, from google console
    credentials = ServiceAccountCredentials
        .from_json_keyfile_name(
            json_credentials_file_name,
            scope
        )

    authorized_client = gspread.authorize(credentials)

    data = authorized_client.open('Dados')

    worksheets = dict()
    worksheets['recipients_sheet'] = sheet.worksheet('Destinatarios')
    worksheets['credentials_sheet'] = sheet.worksheet('Credenciais')

    return(worksheets)

