package main

type Resultado struct {
	EscolhidoPorVoce     string          `json:"escolhido_por_voce"`
	EscolhidoPelaMaquina string          `json:"escolhido_pela_maquina"`
	Resultado            StatusResultado `json:"resultado"`
	ErrorMessage         string          `json:"mensagem_de_erro"`
}

type StatusResultado struct {
	Label string `json:"label"`
	Code  int    `json:"code"`
}

type OpcaoDeEscolha struct {
	Escolha string
	Vence   string
	Perde   string
}

func GetResultadoVencedor(jogador OpcaoDeEscolha, maquina OpcaoDeEscolha) Resultado {
	if jogador.Escolha == maquina.Escolha {
		return Resultado{
			Resultado: StatusResultado{
				Code:  3,
				Label: "empate",
			},
			EscolhidoPorVoce:     jogador.Escolha,
			EscolhidoPelaMaquina: maquina.Escolha,
		}
	}

	if jogador.Vence == maquina.Escolha {
		return Resultado{
			Resultado: StatusResultado{
				Code:  1,
				Label: "você ganhou",
			},
			EscolhidoPorVoce:     jogador.Escolha,
			EscolhidoPelaMaquina: maquina.Escolha,
		}
	}

	return Resultado{
		Resultado: StatusResultado{
			Code:  2,
			Label: "você perdeu",
		},
		EscolhidoPorVoce:     jogador.Escolha,
		EscolhidoPelaMaquina: maquina.Escolha,
	}
}

func GetEscolha(param string) (bool, OpcaoDeEscolha) {

	mapper := map[string]OpcaoDeEscolha{}

	for _, e := range OpcoesDeEscolha() {
		mapper[e.Escolha] = e
	}

	r, ok := mapper[param]

	return ok, r
}

func OpcoesDeEscolha() []OpcaoDeEscolha {

	return []OpcaoDeEscolha{
		{
			Escolha: "papel",
			Vence:   "pedra",
			Perde:   "tesoura",
		},
		{
			Escolha: "pedra",
			Vence:   "tesoura",
			Perde:   "papel",
		},
		{
			Escolha: "tesoura",
			Vence:   "papel",
			Perde:   "pedra",
		},
	}
}
