# Documentação dos contratos

Este documento descreve os contratos Solidity presentes em `packages/contracts/src/` do projeto USPToken. Para cada contrato há um resumo do papel desempenhado, lista de variáveis de estado, eventos e descrição dos métodos principais (parâmetros, comportamento e condições de falha).

---

## `IIdentityRegistry.sol` (interface)

Resumo
- Interface mínima para checar se um endereço representa um aluno ativo no registro.

Assinaturas
- `function isStudentActive(address studentAddress) external view returns (bool);`
  - Parâmetros: `studentAddress` — endereço do aluno a consultar.
  - Retorno: `bool` — `true` se o aluno estiver ativo, `false` caso contrário.
  - Visibilidade: `external view`.

Observações
- Usada por outros contratos (por exemplo `USPToken`, `USPCertificate`, `USPPaymaster`) para validar permissões baseadas no registro de identidade.

---

## `IdentityRegistry.sol`

Resumo
- Registro on-chain de alunos. Permite que contas com papel administrativo adicionem alunos, alterem seu status e consultem se um aluno está ativo.

Herança
- `AccessControl` (OpenZeppelin) — fornece gerenciamento de papéis/roles.

Variáveis de estado
- `bytes32 public constant ADMIN_ROLE` — identificador do papel administrativo.
- `struct Student { bool isActive; string studentId; }` — representa um aluno.
- `mapping(address => Student) public students` — mapeia endereço => informações do aluno.

Eventos
- `StudentAdded(address indexed studentAddress, string studentId)` — emitido quando um aluno é adicionado.
- `StudentStatusChanged(address indexed studentAddress, bool isActive)` — emitido quando o status ativo/inativo do aluno é alterado.

Construtor
- `constructor(address admin)`
  - Concede `DEFAULT_ADMIN_ROLE` e `ADMIN_ROLE` para `admin`.

Funções públicas
- `function addStudent(address studentAddress, string memory studentId) public onlyRole(ADMIN_ROLE)`
  - Ação: registra/atualiza o mapeamento `students` marcando `isActive = true` e definindo `studentId`.
  - Efeito: emite `StudentAdded`.
  - Restrição: apenas quem tem `ADMIN_ROLE` pode chamar.

- `function setStudentStatus(address studentAddress, bool status) public onlyRole(ADMIN_ROLE)`
  - Ação: altera o atributo `isActive` do aluno.
  - Requer: o aluno já deve ter um `studentId` salvo (verificação por `bytes(...).length > 0`).
  - Emite `StudentStatusChanged`.

- `function isStudentActive(address studentAddress) public view returns (bool)`
  - Retorna se o aluno está ativo.
  - Visibilidade: `public view`.

Observações
- O contrato centraliza a lógica de quem é considerado aluno ativo no sistema, servindo como fonte de verdade para os demais contratos.

---

## `USPToken.sol`

Resumo
- Token ERC20 utilizado como "moeda USP". Suporta mint por contas com papel `MINTER_ROLE`, e valida transferências apenas entre endereços que sejam alunos ativos (exceto mint/burn).

Herança
- `ERC20`, `ERC20Burnable`, `AccessControl`.

Variáveis de estado
- `bytes32 public constant MINTER_ROLE` — papel responsável por cunhar tokens.
- `IIdentityRegistry public identityRegistry` — referência ao contrato de registro de identidades.

Construtor
- `constructor(address _identityRegistry, address admin) ERC20("USP Token", "U$PT")`
  - Configura `identityRegistry` com a interface para `_identityRegistry`.
  - Concede `DEFAULT_ADMIN_ROLE` e `MINTER_ROLE` a `admin`.

Funções públicas
- `function mint(address to, uint256 amount) public onlyRole(MINTER_ROLE)`
  - Ação: cunha `amount` tokens para `to` usando `_mint`.
  - Requer: `identityRegistry.isStudentActive(to)` ser `true`.
  - Reverte com mensagem: `Destinatario nao e um estudante ativo` se o destinatário for inativo.

Funções internas
- `function _update(address from, address to, uint256 value) internal override`
  - Observação: este contrato sobrescreve `_update` (implementação específica do projeto) para validar transferências.
  - Comportamento: se não for mint (`from == address(0)`) nem burn (`to == address(0)`), exige que tanto `from` quanto `to` sejam alunos ativos via `identityRegistry.isStudentActive(...)`.
  - Reverte com mensagens: `Remetente inativo no registro` ou `Destinatario inativo no registro` conforme o caso.
  - Em seguida chama `super._update(from, to, value)`.

Observações
- A proteção garante que tokens só circulem entre alunos registrados e ativos, enquanto permite mint/burn sem checagens.

---

## `USPCertificate.sol`

Resumo
- Contrato SBT (Soulbound Token) no padrão ERC-721 para emitir certificados universitários. Para resgatar um certificado o estudante deve queimar uma quantia definida de uma "moeda social" externa.

Herança
- `ERC721`, `ERC721URIStorage`, `AccessControl`.

Interfaces/Referências
- `ISocialCurrency` (interface local) com `function burnFrom(address account, uint256 amount) external;` — usada para queima da moeda social.
- `IIdentityRegistry` — usado para checar se o aluno está ativo.

Variáveis de estado
- `bytes32 public constant ISSUER_ROLE` — papel para emissores (setado no construtor como ROLE do admin).
- `ISocialCurrency public socialCurrency` — endereço/contrato da moeda social (contrato que permite `burnFrom`).
- `IIdentityRegistry public identityRegistry` — referência ao registro de identidades.
- `uint256 private _nextTokenId` — contador para o próximo tokenId a ser cunhado.
- `uint256 public constant CERTIFICATE_COST = 100 * 10**18` — custo em moeda social para resgatar um certificado.

Eventos
- `event Locked(uint256 tokenId)` — usado para indicar que o token é "prisão" (soulbound) conforme convenção ERC-5192.

Construtor
- `constructor(address _identityRegistry, address _socialCurrency, address admin)`
  - Configura `identityRegistry` e `socialCurrency`.
  - Concede `DEFAULT_ADMIN_ROLE` e `ISSUER_ROLE` para `admin`.

Funções públicas
- `function redeemCertificate(string memory metadataURI) public`
  - Parâmetros: `metadataURI` — URI do metadata a associar ao NFT.
  - Fluxo:
    1. Define `student = msg.sender`.
    2. Exige `identityRegistry.isStudentActive(student)` — caso contrário reverte com `Apenas alunos ativos podem resgatar`.
    3. Chama `socialCurrency.burnFrom(student, CERTIFICATE_COST)` — queima o montante exigido (o estudante precisa ter dado allowance previamente se aplicável).
    4. Faz mint do ERC721 para `student` usando `_safeMint` com um novo `tokenId` obtido de `_nextTokenId++`.
    5. Define o token URI com `_setTokenURI(tokenId, metadataURI)`.
    6. Emite `Locked(tokenId)` para indicar que o token é soulbound.

Lógica Soulbound
- `_update(address to, uint256 tokenId, address auth) internal override returns (address)`
  - Implementa bloqueio de transferência: obtém `from = _ownerOf(tokenId)`.
  - Se `from != address(0)` e `to != address(0)` (ou seja, nem mint nem burn), reverte com `Tokens Soulbound nao podem ser transferidos`.
  - Caso contrário, delega para `super._update(to, tokenId, auth)`.
  - Observação: essa função assume a existência de hooks internos (`_ownerOf`, `_update`) compatíveis com as bibliotecas usadas; seu objetivo é impedir transferências pós-mint.

Suporte a interfaces
- `function supportsInterface(bytes4 interfaceId)` — override para combinar `ERC721`, `ERC721URIStorage` e `AccessControl`.
- `function tokenURI(uint256 tokenId)` — override para delegar ao `ERC721URIStorage`.

Observações
- O padrão adotado é emitir o NFT no momento do resgate e impedir qualquer transferência subsequente. O evento `Locked` fornece compatibilidade com interfaces/consumos que esperem SBTs.

---

## `USPPaymaster.sol`

Resumo
- Paymaster que interage com um EntryPoint (Account Abstraction) para pagar por transações de estudantes ativos. Verifica se a operação é permitida e contabiliza gás gasto por estudante.

Herança / Interfaces
- Implementa `IPaymaster` (interface do pacote de account-abstraction).
- Herdado `Ownable` para funções administrativas.

Variáveis de estado
- `IIdentityRegistry public identityRegistry` — referência ao registro de identidades para validar alunos.
- `IEntryPoint public immutable entryPoint` — referência ao EntryPoint (contrato de Account Abstraction).
- `mapping(address => uint256) public gasSpentByStudent` — registra gasto acumulado por estudante (campo usado para possíveis limites/controle).
- `uint256 public constant MAX_GAS_PER_STUDENT = 0.05 ether` — exemplo de constante (limite indicado no código).
- `mapping(address => bool) public whitelistedTargets` — contratos/endereços que o paymaster permite que os usuários invoquem.

Eventos
- `GasCharged(address indexed student, uint256 actualGasCost, PostOpMode mode)` — emite informações sobre cobrança de gás após execução.

Construtor
- `constructor(IEntryPoint _entryPoint, address _identityRegistry, address _admin) Ownable(_admin)`
  - Inicializa `entryPoint` e `identityRegistry`, e seta o dono via `Ownable`.

Funções externas (implementações do IPaymaster)
- `function validatePaymasterUserOp(PackedUserOperation calldata userOp, bytes32 /*userOpHash*/, uint256 maxCost) external view override returns (bytes memory context, uint256 validationData)`
  - Objetivo: decidir se o paymaster aceita pagar pela operação.
  - Verificações realizadas (na ordem):
    1. `identityRegistry.isStudentActive(userOp.sender)` — se falso, retorna `("", 1)` (negando a operação).
    2. Extrai `target` de `userOp.callData` via `_getTargetFromCallData` e checa `whitelistedTargets[target]` — se falso, retorna `("", 1)`.
    3. Verifica saldo do `entryPoint` para este paymaster (`entryPoint.balanceOf(address(this)) < maxCost`) — se insuficiente, retorna `("", 1)`.
  - Se tudo correto, retorna `context = abi.encode(userOp.sender)` (usado em `postOp`) e `validationData = 0`.
  - Observação: assinatura usa `PackedUserOperation`/`maxCost` de acordo com a versão da interface importada.

- `function postOp(PostOpMode mode, bytes calldata context, uint256 actualGasCost, uint256 /* actualUserOpFeePerGas */) external override`
  - Objetivo: ser chamado após execução da operação para registrar custos e possivelmente reembolsos/ajustes.
  - Decodifica `student = abi.decode(context, (address))` — valor retornado por `validatePaymasterUserOp`.
  - Atualiza `gasSpentByStudent[student] += actualGasCost`.
  - Emite `GasCharged(student, actualGasCost, mode)`.

Funções administrativas
- `function setTargetWhitelist(address target, bool status) external onlyOwner` — adiciona/retira um alvo da whitelist.
- `function make_deposit() external payable` — faz depositTo no `entryPoint` usando `msg.value`.
- `function withdrawTo(address payable withdrawAddress, uint256 amount) external onlyOwner` — retira fundos do `entryPoint` para `withdrawAddress`.

Funções utilitárias
- `function _getTargetFromCallData(bytes calldata data) internal pure returns (address target)`
  - Tentativa simples de extrair o endereço alvo dos primeiros 20 bytes do calldata; se `data.length < 20` retorna `address(0)`.

Observações
- O contrato atua como um gatekeeper: só paga por operações de estudantes ativos, para alvos aprovados e se houver saldo suficiente. Ele registra o gasto por estudante para controle/monitoramento futuro.
- Dependendo da versão exata das interfaces importadas (`IPaymaster`/`IEntryPoint`) as assinaturas exatas de `validatePaymasterUserOp` e `postOp` devem ser compatíveis; consultar os arquivos de interfaces usados no `lib/` do projeto.

---

### Notas finais
- A integração entre os contratos faz com que o `IdentityRegistry` seja a autoridade central que define quem é um "aluno ativo". `USPToken` e `USPCertificate` dependem dessa informação para permitir ações (mint, resgate) apenas para alunos ativos. `USPPaymaster` também depende do registro para decidir por quais operações ele paga.
- Ao alterar as regras do registro (por exemplo, lógica de ativação), considerar efeito imediato sobre os demais contratos.

---

Fim da documentação.
