class recipient():
    def __init__(self, information:dict):
        self.name = information["Nome"]
        self.email = information["Email"]
        self.city = information['Cidade']
        self.n_competitions = information['N de competições']
        self.verification_date = information['Data de verificação']
        
    def 