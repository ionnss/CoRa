# **CoRa**

CoRa é um projeto open source criado para facilitar a rotina de psicólogos, oferecendo uma alternativa simples e acessível para o gerenciamento de consultas e dados de pacientes.

---

## **Objetivo**
- **Acessibilidade Financeira:** Eliminar a necessidade de pagar por softwares que cobram assinaturas recorrentes, fornecendo uma solução totalmente gratuita.
- **Privacidade dos Dados:** Garantir que nenhum dado sensível dos pacientes seja enviado para servidores externos, evitando violações de privacidade. Os dados são armazenados localmente no computador do psicólogo, usando SQLite.
- **Simplicidade:** Simplificar processos do dia a dia sem comprometer a privacidade ou gerar custos adicionais com integrações automatizadas.

---

## **Principais Funcionalidades**
- **Gerenciamento Local de Dados:**
  - Armazene os dados dos pacientes no banco de dados SQLite.
- **Agendamento de Consultas:**
  - Organize horários e mantenha um histórico das sessões.
- **Notificações Manuais via WhatsApp:**
  - Gere links dinâmicos para enviar mensagens diretamente pelo WhatsApp.
- **Open Source e Gratuito:**
  - Código-fonte aberto, permitindo personalizações e contribuições da comunidade.

---

## **Requisitos**
- **Sistema Operacional:**
  - Windows, macOS ou Linux.
- **Dependências:**
  - [Go](https://go.dev/) versão 1.20 ou superior.
  - [SQLite](https://sqlite.org/).

---

## **Instalação**
1. Clone o repositório:
   ```bash
   git clone https://github.com/seuusuario/cora.git
   cd cora
   ```
2. Compile o programa:
   ```bash
   go build -o cora
   ```
3. Execute o programa:
   ```bash
   ./cora
   ```

---

## **Como Contribuir**
1. Faça um fork do repositório.
2. Crie uma branch para sua feature:
   ```bash
   git checkout -b minha-feature
   ```
3. Envie suas alterações:
   ```bash
   git add .
   git commit -m "Minha contribuição"
   git push origin minha-feature
   ```
4. Abra um Pull Request.

---

## **Licença**
CoRa é licenciado sob a [MIT License](https://opensource.org/licenses/MIT). Isso significa que você pode:
- Usar o software para qualquer finalidade.
- Modificar o código para suas necessidades.
- Compartilhar e distribuir versões modificadas.

### **Resumo da Licença:**
- **Responsabilidade:** Nenhuma garantia é fornecida; use por sua conta e risco.
- **Atribuição:** Dê os devidos créditos ao repositório original.

Para mais detalhes, consulte o arquivo [LICENSE](LICENSE) no repositório.

---

Com essas melhorias, o README.md está mais completo e profissional. Se precisar de ajuda com o arquivo de licença `LICENSE`, posso gerar um modelo para você!
