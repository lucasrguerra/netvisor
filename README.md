# NetVisor

O NetVisor é uma ferramenta de análise da qualidade da rede de internet, que permite a visualização de dados de latência, jitter, perda de pacotes e largura de banda. A ferramenta foi desenvolvida para ser utilizada em conjunto com o [InfluxDB](https://www.influxdata.com/) e o [Grafana](https://grafana.com/), permitindo a visualização dos dados em tempo real.

## Instalação

Para instalar o NetVisor, basta clonar o repositório e instalar as dependências:

```bash
git clone https://github.com/lucasrguerra/netvisor.git
cd netvisor
pip install -r requirements.txt
```

## Utilização

Para utilizar o NetVisor, basta executar o script `netvisor.py`:

```bash
python netvisor.py
```