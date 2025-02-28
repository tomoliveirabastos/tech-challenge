package main

import "testing"

func TestIniciarUmaPartidaContraMaquina(t *testing.T) {

	_, maquinaEscolha := GetEscolha("tesoura")

	maquina := Maquina{
		OpcaoDeEscolha: maquinaEscolha,
	}

	j := Jogo{
		Maquina: maquina,
	}

	r := j.IniciarUmaPartidaContraMaquina("pedra")

	if r.Resultado.Code != 1 {
		t.Error("Resultado deveria ser 1", ", retornou:", r.Resultado.Code)
	}
}

func TestParametroPedraGanha(t *testing.T) {
	_, jogadorEscolha := GetEscolha("pedra")
	_, maquinaEscolha := GetEscolha("tesoura")

	maquina := Maquina{
		OpcaoDeEscolha: maquinaEscolha,
	}

	r := GetResultadoVencedor(jogadorEscolha, maquina.OpcaoDeEscolha)

	if r.Resultado.Code != 1 {
		t.Error("Resultado deveria ser 1", ", retornou:", r.Resultado.Code)
	}
}

func TestParametroPedraPerde(t *testing.T) {

	_, jogadorEscolha := GetEscolha("pedra")
	_, maquinaEscolha := GetEscolha("papel")

	maquina := Maquina{
		OpcaoDeEscolha: maquinaEscolha,
	}

	r := GetResultadoVencedor(jogadorEscolha, maquina.OpcaoDeEscolha)

	if r.Resultado.Code != 2 {
		t.Error("Resultado deveria ser 2", ", retornou:", r.Resultado.Code)
	}
}

func TestParametroPapelGanha(t *testing.T) {

	_, jogadorEscolha := GetEscolha("papel")
	_, maquinaEscolha := GetEscolha("pedra")

	maquina := Maquina{
		OpcaoDeEscolha: maquinaEscolha,
	}

	r := GetResultadoVencedor(jogadorEscolha, maquina.OpcaoDeEscolha)

	if r.Resultado.Code != 1 {
		t.Error("Resultado deveria ser 1", ", retornou:", r.Resultado.Code)
	}
}

func TestParametroPapelPerde(t *testing.T) {
	_, jogadorEscolha := GetEscolha("papel")
	_, maquinaEscolha := GetEscolha("tesoura")

	maquina := Maquina{
		OpcaoDeEscolha: maquinaEscolha,
	}

	r := GetResultadoVencedor(jogadorEscolha, maquina.OpcaoDeEscolha)

	if r.Resultado.Code != 2 {
		t.Error("Resultado deveria ser 2", ", retornou:", r.Resultado.Code)
	}
}

func TestParametroTesouraGanha(t *testing.T) {
	_, jogadorEscolha := GetEscolha("tesoura")
	_, maquinaEscolha := GetEscolha("papel")

	maquina := Maquina{
		OpcaoDeEscolha: maquinaEscolha,
	}

	r := GetResultadoVencedor(jogadorEscolha, maquina.OpcaoDeEscolha)

	if r.Resultado.Code != 1 {
		t.Error("Resultado deveria ser 1", ", retornou:", r.Resultado.Code)
	}
}

func TestParametroTesouraPerde(t *testing.T) {
	_, jogadorEscolha := GetEscolha("tesoura")
	_, maquinaEscolha := GetEscolha("pedra")

	maquina := Maquina{
		OpcaoDeEscolha: maquinaEscolha,
	}

	r := GetResultadoVencedor(jogadorEscolha, maquina.OpcaoDeEscolha)

	if r.Resultado.Code != 2 {
		t.Error("Resultado deveria ser 2", ", retornou:", r.Resultado.Code)
	}
}
