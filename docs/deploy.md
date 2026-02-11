# Guia de Deploy: Contrato USP (Foundry)

Este documento descreve as etapas para realizar o deploy do contrato utilizando o script `DeployUSP.s.sol`.

## 1. Pré-requisitos

Antes de iniciar, certifique-se de ter o **Foundry** instalado e atualizado:
```bash
foundryup
```

## 2. Configuração do Ambiente

Crie ou edite o arquivo .env na raiz do projeto para armazenar suas credenciais de forma segura. Nunca versione este arquivo.

```bash
# URL da rede (ex: Alchemy, Infura ou QuickNode)
RPC_URL=[https://eth-sepolia.g.alchemy.com/v2/SUA_CHAVE](https://eth-sepolia.g.alchemy.com/v2/SUA_CHAVE)

# Chave privada da sua nova carteira MetaMask
PRIVATE_KEY=0x...

# Opcional: Para verificar o contrato no Etherscan
ETHERSCAN_API_KEY=SUA_CHAVE_ETHERSCAN
```

Para carregar as variáveis no terminal, execute:

```bash
source .env
```

## 3. Simulação de Deploy (Dry Run)

Sempre realize uma simulação local antes de gastar gás real. O comando abaixo executa o script localmente para garantir que não há erros de lógica:

```bash
forge script script/DeployUSP.s.sol --rpc-url $RPC_URL
```

## 4. Deploy Oficial
Para enviar a transação para a rede (on-chain), utilizamos a flag `--broadcast`. Se quiser verificar o contrato automaticamente no Etherscan, adicione `--verify`.

```bash
forge script script/DeployUSP.s.sol \
    --rpc-url $RPC_URL \
    --private-key $PRIVATE_KEY \
    --broadcast \
    --verify \
    -vvvv
```

**Explicação das Flags:**

-  `--broadcast`: Envia as transações para a rede real.

- `--verify`: Envia o código-fonte para o Etherscan para validação pública.

- `-vvvv`: Nível máximo de verbosidade (útil para debugar erros de gás ou revert).

## 5. Pós-Deploy

Após a execução, o Foundry criará uma pasta chamada broadcast. Nela, você encontrará um arquivo .json contendo o endereço do contrato (address) e o hash da transação.

**Dica:** Salve o endereço do contrato exibido no terminal para facilitar a interação posterior via cast ou no front-end.

