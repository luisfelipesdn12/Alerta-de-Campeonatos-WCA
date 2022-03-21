# Alerta-de-Campeonatos-WCA
Um script que manda um e-mail quando há um campeonato novo na WCA.
A script which send an e-mail when there's a new WCA competition. Integrates the WCA's and the Google Sheets API, the subscription is made by the Google Forms.

![Header](https://raw.githubusercontent.com/luisfelipesdn12/Alerta-de-Campeonatos-WCA/master/images/Email%20Header%20English.png)

[![Run - checking for new competitions](https://github.com/luisfelipesdn12/Alerta-de-Campeonatos-WCA/actions/workflows/run_check.yml/badge.svg)](https://github.com/luisfelipesdn12/Alerta-de-Campeonatos-WCA/actions/workflows/run_check.yml)
[![GoDoc](https://godoc.org/github.com/luisfelipesdn12/Alerta-de-Campeonatos-WCA?status.svg)](https://godoc.org/github.com/luisfelipesdn12/Alerta-de-Campeonatos-WCA)
[![GoReportCard](https://goreportcard.com/badge/github.com/luisfelipesdn12/Alerta-de-Campeonatos-WCA)](https://goreportcard.com/report/github.com/luisfelipesdn12/Alerta-de-Campeonatos-WCA)
[![License](https://img.shields.io/github/license/luisfelipesdn12/Alerta-de-Campeonatos-WCA)](https://github.com/luisfelipesdn12/Alerta-de-Campeonatos-WCA/blob/master/LICENSE)
[![](https://img.shields.io/badge/last%20runtime-information-informational)](https://luisfelipesdn12.github.io/Runtime-Information-WCA-Alert/)
[![CodeFactor](https://www.codefactor.io/repository/github/luisfelipesdn12/alerta-de-campeonatos-wca/badge)](https://www.codefactor.io/repository/github/luisfelipesdn12/alerta-de-campeonatos-wca)

## Ideia:
>"The World Cube Association governs competitions for mechanical puzzles that are operated by twisting groups of pieces, commonly known as 'twisty puzzles'. The most famous of these puzzles is the Rubik's Cube, invented by professor Rubik from Hungary. A selection of these puzzles are chosen as official events of the WCA.
As the WCA has evolved over the past decade, over 100,000 people have competed in our competitions."
>- Source: "[Who we are](https://www.worldcubeassociation.org/about)" access in 2020 August, 08.

Me and my friends have the *speedcubing* as a hobbie, simplified: solve rubik's cube and other puzzles in the lowest time as possible.
There's official competitions all over the world, realized by the World Cube Association (WCA), as above-mentioned.

<img src="https://www.cps.sp.gov.br/wp-content/uploads/sites/1/2019/08/Etec-Jacare%C3%AD-4%C2%BA-campeonato-mundial-do-cubo.jpg" width="600">

We participate in them, and is very common we check the [WCA's site](https://www.worldcubeassociation.org/competitions) searching for nearly competitions. Sometimes, we check few times a week, and nothing; sometimes, we forgot to check and lost a long-awaited competition. 
To solve this issue, I had the idea of making a script that would check the site periodically and notify us when it identified a competition nearby that could be of interest to us.

## Usage:

Subscribe, inserting your name, email, language and the city you want to be notified by filling the form bellow:


[![**Subscribscription Form**](https://img.shields.io/badge/subscribe%20me-I%20want%20to%20be%20notified-blue?style=for-the-badge&logo)](https://forms.gle/K6vW3YVAYp4d6nb97)

## Execution:
To execute this project, I've initially used Python with the libraries `requests` and `BeautifulSoup` for web scrapping in the site itself and `smtplib` for sending emails. 
But I made a migration to Golang, with the WCA's API instead of the front-end site. I studied modulation of code in local packages, the usage of libraries as `spreadsheet` for connect with Google Sheets API and the `gomail` to send the notifications.

The code works as follows:

- Fetch the data from a spreadsheet in my Google account (recipients data and credentials to emails sending);
- Verify the upcoming competitions of each recipient city;
- Update it in the spreadsheet;
- Compare the current verification with the last one;
- Send an email if this numbers are different;
> All this process is logged in `main.log` file.

In my Google account, the spreadsheet is organized in this format:

### Sheet 1 ("Recipients"):
> The data provided by the form and the past verifications.

|  Form was filled in  | Name |   Email    |    City    | Language | Upcoming Competitions |  Last Verification  |
| -------------------- | ---- | ---------- | ---------- | -------- | --------------------- | ------------------- |
| 00/00/0000 00:00:00  | anne | anne@e.com | New Jersey | English  | 7                     | 0000-00-00 00:00:00 |
| ...                  | ...  | ...        | ...        | ...      | ...                   | ...                 |

### Sheet 2 ("Betas"):
> My personal friends who agreed to be beta testers. When the code is in development it runs here first.

|          -           | Name |   Email    |    City    |  Language  | Upcoming Competitions |  Last Verification  |
| -------------------- | ---- | ---------- | ---------- | ---------- | --------------------- | ------------------- |
|          -           | tagu | tagu@u.com | São Paulo  | Português  | 2                     | 0000-00-00 00:00:00 |
| ...                  | ...  | ...        | ...        | ...        | ...                   | ...                 |

### Sheet 3 ("Credentials"):
> The email and password of the email sender account.

|        Email        |  Password  |
| ------------------- | ---------- |
| myaccount@gmail.com | my9a55w0rd |

## To do:
- [x] Add an runtime `map` with `{city : upcoming copetitions}` and if the city were already verificated in other recipient, do not verificate again and catch this data from the `map`.
- [ ] Add tests in the whole app.

> Suggest something to do in [issues](https://github.com/luisfelipesdn12/Alerta-de-Campeonatos-WCA/issues) :)

## LICENSE:
```LICENSE
Alerta-deCampeonatos-WCA - A script which send an e-mail when there's a new WCA competition. 
Copyright (C) 2020  Luis Felipe Santos do Nascimento

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
```
