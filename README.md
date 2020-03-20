# PubSub

## Message Type

| Name | Value | Description |  
| --- | --- | --- | --- |
| Connect | 0x00001 | request for confirmation of connection authority |
| Connect Ack Succcess | 0x00001 | connection permission confirmation successful |
| Connect Ack Failed | 0x00001 | connection permission confirmation failed |
| Reconnect Request | 0x00002 | reconnect request |
| Disconnect | 0x00002 | disconnect request |
| Subscribe | 0x00002 | subscribe request |
| Subscribe Ack Success | 0x00002 | subscription error |
| Subscribe Permission Error | 0x00002 | subscription error |
| Unsubscribe | 0x00002 | unsubscription request |
| Publish | 0x00002 | publish request |
| Publish Ack Success | 0x00002 | subscription error |
| Publish Permission Error | 0x00002 | publish error |
| Resend Request | 0x00002 | resend request *(maybe data broken)|
| Ping | 0x00002 | ping request |
| Pong  | 0x00002 | pong |
| Add Subscriber | 0x00002 | add subscriber |
| Remove Subscriber | 0x00002 | remove subscriber |
| Get Subscribers Request | 0x00002 | request subscribers list |
| Subscribers List | 0x00002 | return subscribers list |
| Close Request | 0x00002 | close topoc request (all subscribers will be unsubscribe) |

## Data Type

| Name | value | Description | 
| --- | --- | --- | --- |
| Connect token | 0x00001 | request for confirmation of connection authority |
| Message ID | 0x00001 | request for confirmation of connection authority |
| Message | 0x00001 | request for confirmation of connection authority |
| topic | 0x00001 | request for confirmation of connection authority |
| subscriber id | 0x00001 | request for confirmation of connection authority |

## References

 * https://logmi.jp/tech/articles/322569
 * https://github.com/cskr/pubsub
