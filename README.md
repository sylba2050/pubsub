# PubSub

## Design

![](https://user-images.githubusercontent.com/20987269/81062231-93ca9f80-8f10-11ea-9a17-4bd44b8e6aec.jpg)

* messageのlengthにはpayloadの長さが入る(bit)
* payloadのlengthにはdataのlengthに長さが入る(bit)
* dataは画像上は32bitだが実際は可変長
* ひとつのdata type に対してひとつの値がはいる(例えばsubscribers listを取得した場合, payloadにはsubscriber ID, length, dataがsubscriberの数だけ入る)
* サーバー側は内部的にクライアントのIDを保持(UUID?)
  * クライアントはそれを知るすべは今のところない

## Message Type

| Name | Value | Description |  
 --- | --- | --- | 
| Connect                          | 0x0001 | request for confirmation of connection authority |
| ConnectAckSuccess                | 0x0002 | connection permission confirmation successful    |
| ConnectAckFailure                | 0x0003 | connection permission confirmation failed        |
| ReconnectRequest                 | 0x0004 | reconnect request |
| Disconnect                       | 0x0005 | disconnect request |
| Subscribe                        | 0x0020 | subscribe request |
| UnSubscribe                      | 0x0021 | unsubscribe request |
| SubscribeAckSuccess              | 0x0022 | subscription successful |
| SubscribePermisionError          | 0x0023 | subscription failed |
| Publish                          | 0x0040 | publish request |
| PublishAckSuccess                | 0x0041 | publish successful |
| PublishPermissionError           | 0x0042 | publish failed |
| Ping                             | 0x0060 | ping |
| Pong                             | 0x0061 | pong |
| AddSubscriber ※                  | 0x0100 | add subscriber |
| AddSubscriberPermissionError     | 0x0101 | add subscriber failed |
| RemoveSubscriber ※               | 0x0102 | remove subscriber request |
| RemoveSubscriberPermissionError  | 0x0103 | remove subscriber failed |
| GetSubscribersRequest ※          | 0x1000 | subscribers list request |
| GetSubscribersPermissionError    | 0x1001 | get subscribers list failed |
| SubscribersList                  | 0x1002 | return subscribers list |
| CloseRequest ※                   | 0x2000 | close topic request (all subscribers will be unsubscribe) |

※ 要管理者権限

## Data Type

| Name | value | Description | 
| --- | --- | --- |
| Connect token   | 0x4001 | connection token (use check authority and reconnect) |
| Message ID      | 0x4002 | message id |
| Message         | 0x4003 | message |
| topic ID        | 0x4004 | topic id |
| subscriber ID   | 0x4005 | subscriber id |

## References

 * https://logmi.jp/tech/articles/322569
 * https://github.com/cskr/pubsub
