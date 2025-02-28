package main

type Jogo struct {
	Maquina Maquina
}

func NovoJogo() Jogo {

	return Jogo{
		Maquina: NovaMaquina(),
	}
}

func (j *Jogo) IniciarUmaPartidaContraMaquina(param string) Resultado {
	ok, jogador := GetEscolha(param)

	if !ok {
		return Resultado{
			ErrorMessage: "Necessário escolher um dos Parametros (pedra, papel, tesoura)",
		}
	}

	return GetResultadoVencedor(jogador, j.Maquina.OpcaoDeEscolha)
}
