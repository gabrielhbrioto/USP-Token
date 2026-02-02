## 🏗️ Componentes da Arquitetura

### 1. Smart Contracts (Camada On-Chain)

#### Contrato `IdentityRegistry`

Gerencia a base de dados de identidade institucional dos alunos.


- **Papel**: Atua como o "Oráculo de Identidade" do sistema.


- **Funcionalidade**: Mapeia endereços Ethereum a uma `struct Student`, que contém o status de atividade (`isActive`) e o identificador acadêmico (`studentId` / Número USP).


- **Governança**: Apenas endereços com a `ADMIN_ROLE` podem adicionar alunos ou alterar seu status.


#### Contrato `USPToken` (ERC-20)

Implementa a moeda social utilizada para recompensar atividades acadêmicas.

- **Papel**: Motor econômico do ecossistema.

- **Restrições de Circulação**: A função `_update` verifica no `IdentityRegistry` se tanto o remetente quanto o destinatário são alunos ativos antes de permitir qualquer transferência.

- **Emissão**: Somente administradores com `MINTER_ROLE` podem gerar novos tokens para alunos ativos.


#### Contrato `USPCertificate` (ERC-5192 / SBT)

Gerencia a emissão de certificados digitais intransferíveis (Soulbound Tokens).

- **Papel**: Registro permanente de conquistas acadêmicas.

- **Mecânica de Troca**: Exige o pagamento de `CERTIFICATE_COST` (100 tokens), realizando a queima automática do saldo do aluno via `burnFrom` no contrato `USPToken`.

- **Intransferibilidade**: Bloqueia qualquer tentativa de transferência entre usuários após o minting, permitindo apenas a emissão e a queima (SBT).

#### Contrato USPPaymaster (ERC-4337)

Componente de Abstração de Conta que patrocina as taxas de rede (gás) para os alunos.

- **Papel**: Elimina a necessidade de o aluno possuir Ether para interagir com o sistema.

- **Validação**: Verifica se o aluno é ativo no `IdentityRegistry` e se o destino da transação está na `whitelistedTargets` antes de autorizar o pagamento.

- **Monitoramento**: Utiliza a função `postOp` para registrar o custo real de gás consumido por cada aluno para fins de auditoria e controle de cotas.


### 2. Middleware e Infraestrutura

#### DApp (Frontend)

Aplicação de página única (SPA) construída em React.js.

- **Papel**: Interface principal de interação para alunos e administradores.

- **Estudante**: Visualiza saldo de moedas, histórico de atividades e resgata certificados.

- **Administrador**: Painel para gerenciar o `IdentityRegistry` e realizar o mint de recompensas.

#### Servidor Relayer/Bundler (ERC-4337)

Serviço de infraestrutura em Node.js.

- **Papel**: Recebe `UserOperations` assinadas pelo DApp, agrupa-as (bundling) e as envia para o contrato `EntryPoint` na blockchain para execução patrocinada pelo `USPPaymaster`.

#### Serviço IPFS (Pinata/Infura)
Armazenamento descentralizado de arquivos.

- **Papel**: Hospeda os metadados JSON e as imagens dos certificados.

- **Integração**: O hash do arquivo no IPFS é salvo no contrato `USPCertificate` como a `tokenURI` do certificado emitido.

## 🔄 Fluxos Operacionais Previstos

### 1. Onboarding e Recompensa de Aluno

1. Administrador chama `addStudent` no `IdentityRegistry`, vinculando a carteira do aluno ao seu NUSP.

2. Após a realização de uma atividade, o Admin chama `mint` no `USPToken`.

3. O contrato `USPToken` consulta o `IdentityRegistry` para confirmar que o aluno está ativo.

4. O saldo do aluno é atualizado on-chain.

### 2. Resgate de Certificado (Fluxo Gasless)

1. Aluno seleciona um certificado no DApp.

2. O DApp gera uma `UserOperation` contendo o chamado para `redeemCertificate`.

3. O Bundler envia a operação para validação do `USPPaymaster`.

4. O `USPPaymaster` verifica se o aluno é ativo e se o contrato de certificados é autorizado.

5. O Contrato de Certificado:

    - Verifica o status do aluno no registro.

    - Executa o `burnFrom` de 100 tokens do saldo do aluno.

    - Emite o NFT Soulbound e bloqueia transferências.


6. A função `postOp` do Paymaster registra o custo de gás da operação no perfil do aluno.

### 3. Verificação de Autenticidade Off-Chain

1. Terceiro (ex: recrutador) scaneia o QR Code do certificado apresentado pelo aluno.

2. O sistema de verificação consulta as funções `ownerOf` e `tokenURI` no contrato `USPCertificate`.

3. O sistema cruza o endereço da carteira detentora com os dados no IPFS para validar que o nome e atividade correspondem ao registro oficial.