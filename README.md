# WhatsApp Webhook Payload - Significado dos Campos

Este documento explica o significado de cada campo do objeto JSON recebido via webhook da API do WhatsApp Business (Meta).

## Exemplo de Objeto

```json
{
  "object": "whatsapp_business_account",
  "entry": [
    {
      "id": "xxxxxxxxxx",
      "changes": [
        {
          "value": {
            "messaging_product": "whatsapp",
            "metadata": {
              "display_phone_number": "xxxxxxxx",
              "phone_number_id": "xxxxxxxx"
            },
            "statuses": [
              {
                "id": "wamid.xxxxxxxxxx=",
                "status": "failed",
                "timestamp": "1750773570",
                "recipient_id": "xxxxxxxx",
                "errors": [
                  {
                    "code": 131047,
                    "title": "Re-engagement message",
                    "message": "Re-engagement message",
                    "error_data": {
                      "details": "Message failed to send because more than 24 hours have passed since the customer last replied to this number."
                    },
                    "href": "https://developers.facebook.com/docs/whatsapp/cloud-api/support/error-codes/"
                  }
                ]
              }
            ]
          },
          "field": "messages"
        }
      ]
    }
  ]
}
```

## Significado dos Campos

- **object**
  Sempre `"whatsapp_business_account"`. Indica que o evento é referente a uma conta de WhatsApp Business.

- **entry**
  Array de eventos recebidos. Cada item representa um evento para uma conta de negócio.

  - **id**
    **ID da conta de WhatsApp Business Account (WABA)** na Meta. Identifica a conta de negócio que recebeu o evento.

  - **changes**
    Array de mudanças/eventos ocorridos.

    - **field**
      Tipo de evento recebido. Exemplo: `"messages"`.

    - **value**
      Objeto com os detalhes do evento.

      - **messaging_product**
        Produto de mensagens, sempre `"whatsapp"`.

      - **metadata**
        Informações sobre o número de telefone do remetente.

        - **display_phone_number**
          Número de telefone formatado (exibido).

        - **phone_number_id**
          **ID do número de telefone** cadastrado na conta WABA. Usado para identificar o remetente nas APIs da Meta.

      - **statuses**
        Array de status de mensagens.

        - **id**
          **ID da mensagem** enviada/recebida (message ID).

        - **status**
          Status da mensagem (ex: `"sent"`, `"delivered"`, `"read"`, `"failed"`).

        - **timestamp**
          Momento do evento (em formato Unix timestamp).

        - **recipient_id**
          Número do destinatário da mensagem (em formato internacional).

        - **errors**
          Array de erros (se houver).

          - **code**
            Código do erro.

          - **title**
            Título do erro.

          - **message**
            Mensagem descritiva do erro.

          - **error_data.details**
            Detalhes adicionais sobre o erro.

          - **href**
            Link para documentação do erro.

---

## Resumo

- `"object"`: Tipo do objeto (sempre WhatsApp Business Account).
- `"entry[].id"`: ID da conta de negócio (WABA).
- `"metadata.phone_number_id"`: ID do número de telefone do remetente.
- `"statuses[].id"`: ID da mensagem.
- `"statuses[].status"`: Status da mensagem.
- `"statuses[].recipient_id"`: Número do destinatário.
- `"statuses[].errors"`: Detalhes de erro, se houver.

> Consulte a [documentação oficial da Meta](https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples) para
