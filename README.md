# br-validator

## Descrição
O `br-validator` é uma biblioteca para validação de documentos brasileiros como CPF, CNPJ, e outros. Ele fornece funções simples e eficientes para garantir que os documentos estejam no formato correto e sejam válidos.

## Instalação
Para instalar o `br-validator`, você pode usar o npm ou yarn:

```bash
npm install br-validator
# ou
yarn add br-validator
```

## Uso
Aqui está um exemplo de como usar o `br-validator` para validar um CPF:

```javascript
const { validarCPF } = require('br-validator');

const cpf = '123.456.789-09';
const isValid = validarCPF(cpf);

console.log(`O CPF ${cpf} é válido? ${isValid}`);
```

## Contribuição
Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e pull requests. Para contribuir, siga estas etapas:

1. Faça um fork do repositório.
2. Crie uma nova branch (`git checkout -b minha-nova-feature`).
3. Faça suas alterações e commit (`git commit -am 'Adiciona nova feature'`).
4. Faça push para a branch (`git push origin minha-nova-feature`).
5. Abra um Pull Request.

## Licença
Este projeto está licenciado sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.