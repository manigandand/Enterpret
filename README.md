### Ingestor

> Directory structure

```
- main.go
--- api: the push based ingestion, it expose a one generic webhook endpoint ,
where we can get all alerts sent from different sources.
--- config: app level config pkg
--- errors: custom app error pkg
--- ingestor: core ingestor pkg contains all the logics to ingest different types of data sources using gohtml templates.
--- manifests: conatains all the YAML configurations.
--- response: custom http response pkg
--- schmea
--- store: data store layer
--- task: the pull based ingestion.

```

> How to run

```
Just build and run.

go build && ./enterpret config.json

you can see the available alert configs test endpoints

/v1/alert/discourse/914cf64e-c1dc-4a52-93a9-b456042d3e21
/v1/alert/intercom/914cf64e-c1dc-4a52-93a9-b456042d3e21
```

> API's

> Push alert webhook
> `POST localhost:8080/v1/alert/intercom/f14d887c-48bd-4a09-b8bc-04ca1461657b`

```
{
    "data": {
        "id": 1,
        "created_at": "2022-03-23T04:41:17.600743+05:30",
        "alert_config_id": 1,
        "integration_id": 1,
        "organization_id": 1,
        "source": "intercom",
        "type": "conversation",
        "subject": "Twitter's terms and conditions change",
        "message": "We've removed this part of the conversation to comply with Twitter's terms and conditions. You can view the complete conversation in Intercom.",
        "language": "english",
        "metadata": {
            "conversation_url": "",
            "id": "1122334455",
            "type": "twitter"
        },
        "raw": {
            "conversation_message": {
                "attachments": [],
                "author": {
                    "email": "",
                    "id": "5310d8e7598c9a0b24000002",
                    "name": "",
                    "type": "user"
                },
                "body": "We've removed this part of the conversation to comply with Twitter's terms and conditions. You can view the complete conversation in Intercom.",
                "delivered_as": "customer_initiated",
                "id": "409820079",
                "subject": "Twitter's terms and conditions change",
                "type": "twitter",
                "url": ""
            },
            "id": "1122334455",
            "type": "conversation"
        }
    },
    "meta": {
        "status_code": 201
    }
}
```

> Get all alert configs
> `GET localhost:8080/v1/alert-sources`

```
{
    "data": {
        "discourse": {
            "id": 1,
            "version": "v1",
            "name": "Discourse",
            "slug": "discourse",
            "type": "feedback",
            "support_doc": "https://www.enterpret.com/integrations/discourse",
            "is_valid": true,
            "subject": "{{.topic_slug}}\n",
            "message": "{{.cooked}}\n",
            "language": "english",
            "metadata": {
                "id": "{{.id}}\n",
                "reads": "{{.reads}}\n",
                "score": "{{.score}}\n",
                "username": "{{.username}}\n"
            }
        },
        "intercom": {
            "id": 2,
            "version": "v1",
            "name": "Intercom",
            "slug": "intercom",
            "type": "conversation",
            "support_doc": "https://www.enterpret.com/integrations/intercom",
            "is_valid": true,
            "subject": "{{.conversation_message.subject}}\n",
            "message": "{{.conversation_message.body}}\n",
            "language": "english",
            "metadata": {
                "conversation_url": "{{.conversation_message.url}}\n",
                "id": "{{.id}}\n",
                "type": "{{.conversation_message.type}}\n"
            }
        }
    },
    "meta": {
        "status_code": 200
    }
}
```

> Get all alerts
> `GET localhost:8080/v1/alerts`

```
{
    "data": [
        {
            "id": 1,
            "created_at": "2022-03-23T05:46:41.616573+05:30",
            "alert_config_id": 1,
            "integration_id": 1,
            "organization_id": 1,
            "source": "discourse",
            "type": "feedback",
            "subject": "guest-gate-sign-up-popup-plugin",
            "message": "<p>Is there anything unique about the Guest Gate modal which will allow the “Please Sign up!” H3 to be changed without also changing the standard “Log in” and “Create New Account” H3s?</p>",
            "language": "english",
            "metadata": {
                "id": "872204",
                "reads": "105",
                "score": "71",
                "username": "Jonathan5"
            },
            "raw": {
                "avatar_template": "/user_avatar/meta.discourse.org/jonathan5/{size}/197134_2.png",
                "cooked": "<p>Is there anything unique about the Guest Gate modal which will allow the “Please Sign up!” H3 to be changed without also changing the standard “Log in” and “Create New Account” H3s?</p>",
                "created_at": "2021-01-11T00:20:49.378Z",
                "id": 872204,
                "incoming_link_count": 0,
                "name": "",
                "post_number": 64,
                "post_type": 1,
                "quote_count": 0,
                "readers_count": 104,
                "reads": 105,
                "reply_count": 1,
                "reply_to_post_number": 61,
                "score": 71,
                "topic_id": 56625,
                "topic_slug": "guest-gate-sign-up-popup-plugin",
                "updated_at": "2021-01-11T00:34:22.541Z",
                "username": "Jonathan5",
                "yours": false
            }
        }
    ],
    "meta": {
        "status_code": 200
    }
}
```

```
{
    "data": [
        {
            "id": 1,
            "created_at": "2022-03-23T04:41:17.600743+05:30",
            "alert_config_id": 1,
            "integration_id": 1,
            "organization_id": 1,
            "source": "intercom",
            "type": "conversation",
            "subject": "Twitter's terms and conditions change",
            "message": "We've removed this part of the conversation to comply with Twitter's terms and conditions. You can view the complete conversation in Intercom.",
            "language": "english",
            "metadata": {
                "conversation_url": "",
                "id": "1122334455",
                "type": "twitter"
            },
            "raw": {
                "conversation_message": {
                    "attachments": [],
                    "author": {
                        "email": "",
                        "id": "5310d8e7598c9a0b24000002",
                        "name": "",
                        "type": "user"
                    },
                    "body": "We've removed this part of the conversation to comply with Twitter's terms and conditions. You can view the complete conversation in Intercom.",
                    "delivered_as": "customer_initiated",
                    "id": "409820079",
                    "subject": "Twitter's terms and conditions change",
                    "type": "twitter",
                    "url": ""
                },
                "id": "1122334455",
                "type": "conversation"
            }
        }
    ],
    "meta": {
        "status_code": 200
    }
}
```
