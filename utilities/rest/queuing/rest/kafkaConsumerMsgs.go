/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package rest

// KafkaConsumerMessages : messages to consume 5 second delay
//func KafkaConsumerMessages(cliCtx client.Context, kafkaState queuing.KafkaState) {
//	quit := make(chan bool)
//
//	var cliContextList []client.Context
//
//	var baseRequestList []rest.BaseReq
//
//	var ticketIDList []queuing.Ticket
//
//	var msgList []sdkTypes.Msg
//
//	go func() {
//		for {
//			select {
//			case <-quit:
//				return
//			default:
//				kafkaMsg := queuing.KafkaTopicConsumer("Topic", kafkaState.Consumers, cliCtx.LegacyAmino)
//				if kafkaMsg.Msg != nil {
//					cliContextList = append(cliContextList, queuing.CliCtxFromKafkaMsg(kafkaMsg, cliCtx))
//					baseRequestList = append(baseRequestList, kafkaMsg.BaseRequest)
//					ticketIDList = append(ticketIDList, kafkaMsg.TicketID)
//					msgList = append(msgList, kafkaMsg.Msg)
//				}
//			}
//		}
//	}()
//
//	time.Sleep(queuing.SleepTimer)
//	quit <- true
//
//	if len(msgList) == 0 {
//		return
//	}
//
//	output, err := SignAndBroadcastMultiple(baseRequestList, cliContextList, msgList)
//	if err != nil {
//		jsonError, e := cliCtx.LegacyAmino.MarshalJSON(struct {
//			Error string `json:"error"`
//		}{Error: err.Error()})
//		if e != nil {
//			panic(err)
//		}
//
//		for _, ticketID := range ticketIDList {
//			queuing.AddResponseToDB(ticketID, jsonError, kafkaState.KafkaDB, cliCtx.LegacyAmino)
//		}
//
//		return
//	}
//
//	for _, ticketID := range ticketIDList {
//		queuing.AddResponseToDB(ticketID, output, kafkaState.KafkaDB, cliCtx.LegacyAmino)
//	}
//}
