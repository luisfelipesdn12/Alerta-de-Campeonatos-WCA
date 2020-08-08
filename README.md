# Alerta-de-Campeonatos-WCA
Um script que manda um e-mail quando há um campeonato novo na WCA.
A script witch send an e-mail when there's a new WCA competition. 

![Header](https://raw.githubusercontent.com/luisfelipesdn12/Alerta-de-Campeonatos-WCA/golang/images/Email%20Header%20English.png)

## Ideia:
>"The World Cube Association governs competitions for mechanical puzzles that are operated by twisting groups of pieces, commonly known as 'twisty puzzles'. The most famous of these puzzles is the Rubik's Cube, invented by professor Rubik from Hungary. A selection of these puzzles are chosen as official events of the WCA.
As the WCA has evolved over the past decade, over 100,000 people have competed in our competitions."
>- Source: "[Who we are](https://www.worldcubeassociation.org/about)"  access in 2020 August, 08.

Me and my friends have the *speedcubing* as a hobbie, simplified: solve rubik's cube and other puzzles in the lowest time as possible.
There's official competitions all over the world, realized by the World Cube Association (WCA), as above-mentioned.
><img src="https://www.cps.sp.gov.br/wp-content/uploads/sites/1/2019/08/Etec-Jacare%C3%AD-4%C2%BA-campeonato-mundial-do-cubo.jpg" width="600">

We participate in them, and is very common we check the [WCA's site](https://www.worldcubeassociation.org/competitions) searching for nearly competitions. Sometimes, we check few times a week, and nothing; sometimes, we forgot to check and lost a long-awaited competition. 
To solve this issue, I had the idea of making a script that would check the site periodically and notify us when it identified a competition nearby that could be of interest to us.

## Usage:

Subscribe, inserting your name, email, language and the city you want to be notified by filling the form bellow:

[**Subscribe**](https://forms.gle/K6vW3YVAYp4d6nb97)

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

