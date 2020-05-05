# PubSub

## Design

![](https://user-images.githubusercontent.com/20987269/81062231-93ca9f80-8f10-11ea-9a17-4bd44b8e6aec.jpg)

## Message Type

| Name | Value | Description |  
 --- | --- | --- | 
| Connect                 | 0x0001 | request for confirmation of connection authority |
| ConnectAckSuccess       | 0x0002 | connection permission confirmation successful    |
| ConnectAckFailure       | 0x0003 | connection permission confirmation failed        |
| ReconnectRequest        | 0x0004 | reconnect request |
| Disconnect              | 0x0005 | disconnect request |
| Subscribe               | 0x0020 | subscribe request |
| UnSubscribe             | 0x0021 | subscription error |
| SubscribeAckSuccess     | 0x0022 | subscription error |
| SubscribePermisionError | 0x0023 | unsubscription request |
| Publish                 | 0x0040 | publish request |
| PublishAckSuccess       | 0x0041 | subscription error |
| PublishPermissionError  | 0x0042 | publish error |
| Ping                    | 0x0060 | ping |
| Pong                    | 0x0061 | pong |
| AddSubscriber           | 0x0100 | add subscriber |
| RemoveSubscriber        | 0x0101 | remove subscriber |
| GetSubscribersRequest   | 0x1000 | request subscribers list |
| SubscribersList         | 0x1001 | return subscribers list |
| CloseRequest            | 0x2000 | close topoc request (all subscribers will be unsubscribe) |

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
